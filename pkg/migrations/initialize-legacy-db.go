package migrations

import (
	"bufio"
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
	"text/template"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/squeeze69/dbf"
)

// default values and other constants
const (
	defaultEngine      = "MyIsam"
	defaultCollation   = "utf8_general_ci"
	defaultRecordQueue = 100
	defaultGoroutines  = 2
	defaultFirstRecord = 0
	maxGoroutines      = 64
	minGoroutines      = 1
	minRecordQueue     = 1
)

//number of records in the queue
var recordQueue int

//number of goroutines spawned
var numGoroutines int

//global mysqlurl - see the go lang database/sql package
//sample url: "user:password@(127.0.0.1:3306)/database"
var mysqlurl string

//various flags, set by command line, default to false
var verbose, truncate, createtable, dumpcreatetable, insertignore, nobigint, droptable bool
var abortonsqlerror bool

//optional index
var index string

//max number of record to import, defaults to -1 (means no limit)
var maxrecord int

//first record to fetch
var firstRecord int

//read all dbf in memory
var readinmemory bool

//global variables for --create
var collate string
var engine string

//LockableCounter a simple counter with a Mutex
type LockableCounter struct {
	count int
	l     sync.Mutex
}

//Increment lockable counter by i items
func (lc *LockableCounter) Increment(i int) {
	lc.l.Lock()
	defer lc.l.Unlock()
	lc.count += i
}

//total number on insert errors (if any)
var ierror LockableCounter

//read profile, actually a fixed position file, first row it's a sql url
func readprofile(prfname string) error {
	f, err := os.Open(prfname)
	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	if scanner.Scan() {
		mysqlurl = scanner.Text()
	} else {
		return errors.New("no profile found")
	}
	return nil
}

//returns a "CREATE TABLE" string using templates
func createTableString(table string, collate string, engine string, dbr *dbf.Reader) string {
	var fieldtype string
	fields := dbr.FieldNames()
	//pre allocate
	arf := make([]string, 0, len(fields))
	for k := range fields {
		dbfld, _ := dbr.FieldInfo(k)
		switch dbfld.Type {
		case 'D': //Date field
			fieldtype = "DATE"
		case 'L': //logical
			fieldtype = "CHAR(1)"
		case 'C': //CHAR
			fieldtype = fmt.Sprintf("VARCHAR(%d)", dbfld.Len)
		case 'N': //Numeric could be either Int or fixed point decimal
			if dbfld.DecimalPlaces > 0 {
				//A VARCHAR will do it, +2 it's for sign and decimal sep.
				fieldtype = fmt.Sprintf("VARCHAR(%d)", dbfld.Len+2)
			} else {
				var inttype string
				switch {
				case dbfld.Len < 3:
					inttype = "TINYINT"
				case dbfld.Len >= 3 && dbfld.Len < 5:
					inttype = "SMALLINT"
				case dbfld.Len >= 5 && dbfld.Len < 7:
					inttype = "MEDIUMINT"
				case (dbfld.Len >= 7 && dbfld.Len < 10) || nobigint:
					inttype = "INT"
				case dbfld.Len >= 10:
					inttype = "BIGINT"
				}
				fieldtype = fmt.Sprintf("%s(%d)", inttype, dbfld.Len)
			}
		default:
			fieldtype = fmt.Sprintf("VARCHAR(%d)", dbfld.Len)
		}
		arf = append(arf, fmt.Sprintf("`%s` %s", dbf.Tillzero(dbfld.Name[:]), fieldtype))
	}

	//template for table's creation
	tmpl, err := template.New("table").Parse(
		`CREATE TABLE IF NOT EXISTS {{.Tablename}} (
{{range $i,$e := .Arf}}
{{- if $i}},
{{end}}{{$e}}{{end}}
{{- if .Index}},` +
			"\nINDEX `ndx` (`{{.Index}}`){{end}}" + `
){{if .Collate}} COLLATE='{{.Collate}}'{{end}}{{if .Engine}} ENGINE='{{.Engine}}'{{end}};`)
	if err != nil {
		log.Fatal(err)
	}
	var str string
	buf := bytes.NewBufferString(str)
	er1 := tmpl.Execute(buf, struct {
		Tablename, Collate, Engine, Index string
		Arf                               []string
	}{Tablename: "`" + table + "`", Collate: collate, Engine: engine, Arf: arf, Index: index})
	if er1 != nil {
		log.Fatal(er1)
	}
	return buf.String()
}

//insertRoutine goroutine to insert data in dbms
func insertRoutine(ch chan dbf.OrderedRecord, over *sync.WaitGroup, stmt *sql.Stmt) {
	defer over.Done()
	defer func() {
		//just respawning go routine in case of error - i.e. bad data are not inserted (i.e. slightly malformed dbf rows)
		if r := recover(); r != nil {
			err, ok := r.(error)
			if ok {
				ierror.Increment(1)
				if strings.Contains(err.Error(), "1114") {
					//table is full, no way to continue
					log.Fatal("1114 Table full")
				}
				if abortonsqlerror {
					log.Printf("%s\n", err)
				} else {
					fmt.Println("Recover:", err)
					over.Add(1)
					go insertRoutine(ch, over, stmt)
				}
			}
		}
	}()
	for i := range ch {
		_, err := stmt.Exec(i...)
		if err != nil {
			panic(err)
		}
	}
}

type InitLegacyModel struct {
	DB        *sql.DB
	filename  string
	tablename string
}

//workaround: os.Exit ignores deferred functions
func (init *InitLegacyModel) importLegacy() error {

	var start = time.Now()
	var qstring string
	var insertstatement = "INSERT"
	var skipped, inserted int

	var dbfile *dbf.Reader

	var allfile []byte
	// read the whole file in memory
	allfile, err := os.ReadFile(init.filename)
	if err != nil {
		return err
	}

	dbfile, err = dbf.NewReader(bytes.NewReader(allfile))
	if err != nil {
		return err
	}

	//Set some default flags, skips deleted and "weird" records (see dbf package)
	dbfile.SetFlags(dbf.FlagDateAssql | dbf.FlagSkipWeird | dbf.FlagSkipDeleted | dbf.FlagEmptyDateAsZero)

	//check if the table must be dropped before creation

	fmt.Println("Dropping table:", init.tablename)
	if _, erd := init.DB.Exec(fmt.Sprintf("DROP TABLE IF EXISTS `%s`;", init.tablename)); erd != nil {
		return erd
	}

	//create table section
	fmt.Println("Creating Table: ", init.tablename)
	ctstring := createTableString(init.tablename, collate, engine, dbfile)
	if _, erc := init.DB.Exec(ctstring); erc != nil {
		return erc
	}
	// fmt.Println("-- CREATE TABLE:\n", ctstring)

	//retrieve fields to build the query
	fields := dbfile.FieldNames()
	_, err = init.DB.Exec(fmt.Sprintf("TRUNCATE `%s`;", init.tablename))
	if err != nil {
		return err
	}

	//building the code for the prepared statement
	qstring = fmt.Sprintf("%s INTO %s (`%s`) VALUES (%s?);",
		insertstatement, init.tablename, strings.Join(fields, "`,`"), strings.Repeat("?,", len(fields)-1))
	// fmt.Println("QSTRING:", qstring)
	//it's using a prepared statement, much safer and faster
	stmt, err := init.DB.Prepare(qstring)
	if err != nil {
		return err
	}
	defer stmt.Close()

	fmt.Println("Number of dbf records:", dbfile.Length)

	recordQueue = minRecordQueue
	numGoroutines = 10
	maxrecord = -1
	firstRecord = defaultFirstRecord

	chord := make(chan dbf.OrderedRecord, recordQueue)

	wgroup := new(sync.WaitGroup)
	for i := 0; i < numGoroutines; i = i + 1 {
		wgroup.Add(1)
		go insertRoutine(chord, wgroup, stmt)
	}
	var lastrRecord int
	if maxrecord > 0 && firstRecord+maxrecord < dbfile.Length {
		lastrRecord = firstRecord + maxrecord
	} else {
		lastrRecord = dbfile.Length
	}
	for i := firstRecord; i < lastrRecord; i++ {
		runtime.Gosched()
		rec, err := dbfile.ReadOrdered(i)
		if err == nil {
			if verbose {
				fmt.Println(rec)
			}
			chord <- rec
			inserted++
		} else {
			if _, ok := err.(*dbf.SkipError); ok {
				skipped++
				continue
			}
			return err
		}
	}
	close(chord)
	//waiting for insertRoutine to end
	wgroup.Wait()
	//print some stats
	fmt.Printf("Records: Inserted: %d Skipped: %d\nElapsed Time: %s\n",
		inserted, skipped, time.Since(start))
	fmt.Printf("Queue capacity:%d,goroutines:%d\n",
		recordQueue, numGoroutines)
	if ierror.count > 0 {
		fmt.Printf("Insert Errors:%d\n", ierror.count)
	}
	return nil
}

func (init *InitLegacyModel) ImportDBFDir(directory string) error {

	items, err := ioutil.ReadDir(directory)
	if err != nil {
		return err
	}
	for _, item := range items {
		if strings.Contains(item.Name(), ".dbf") {
			init.filename = fmt.Sprintf("%s%c%s", directory, os.PathSeparator, item.Name())
			init.tablename = item.Name()[:len(item.Name())-4]
			err := init.importLegacy()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
