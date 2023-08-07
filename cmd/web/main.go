package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/letitloose/dog-data/pkg/migrations"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	//initilize app
	app := application{
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}

	app.infoLog.Println("Welcome to Dog Data")

	//initialize the legacy DB
	app.infoLog.Println("Initializing the legacy DB from the DBF files")
	db, err := openDB("lougar:thewarrior@/legacytollerdata?parseTime=true&multiStatements=true")
	defer db.Close()
	if err != nil {
		app.errorLog.Fatal(err)
	}

	initLegacyDB := &migrations.InitLegacyModel{
		DB: db,
	}

	err = initLegacyDB.ImportDBFDir("pkg/migrations/dbf")
	if err != nil {
		app.errorLog.Fatal(err)
	}

	//initialize the maria DB
	app.infoLog.Println("Initializing the dogdata DB from SQL scripts")
	ndb, err := openDB("lougar:thewarrior@/tollerdata?parseTime=true&multiStatements=true")
	defer db.Close()
	if err != nil {
		app.errorLog.Fatalf("failed to open connection to DB: %s", err)
	}

	initDB := migrations.InitModel{
		DB: ndb,
	}

	err = initDB.RunScript("pkg/migrations/sql/setup.sql")
	if err != nil {
		app.errorLog.Fatalf("failed to initialize new DB: %s", err)
	}

	app.infoLog.Println("Ready to migrate!")

}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
