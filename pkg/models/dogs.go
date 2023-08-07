package models

import (
	"database/sql"
	"time"
)

type Dog struct {
	id           int
	regnum       string
	nsdtrcregnum string
	sequencenum  string
	litterid     int
	name         string
	callname     string
	whelpdate    time.Time
	sex          string
	nba          bool
	alive        bool
	intact       bool
	sire         int
	dam          int
}

type DogModel struct {
	DB *sql.DB
}

func (model *DogModel) Insert(dog Dog) error {

	statement := `insert into dogs (regnum, nsdtrcregnum, sequencenum, litterid, name, callname,
			whelpdate, sex, nba, alive, intact, sire, dam) 
			values (NULLIF(?, ''),?,?,?,?,?,?,?,?,?,?,?,?);`

	result, err := model.DB.Exec(statement, dog.regnum, dog.nsdtrcregnum, dog.sequencenum, dog.litterid, dog.name,
		dog.callname, dog.whelpdate, dog.sex, dog.nba, dog.alive, dog.intact, dog.sire, dog.dam)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}
