package models

import (
	"database/sql"
	"time"
)

type Dog struct {
	Id           int
	Regnum       string
	Nsdtrcregnum string
	Sequencenum  string
	Litterid     int
	Name         string
	Callname     string
	Whelpdate    time.Time
	Sex          string
	Nba          bool
	Alive        bool
	Intact       bool
	Sire         int
	Dam          int
}

type DogModel struct {
	DB *sql.DB
}

func (model *DogModel) Insert(dog Dog) error {

	statement := `insert into dogs (regnum, nsdtrcregnum, sequencenum, litterid, name, callname,
			whelpdate, sex, nba, alive, intact, sire, dam) 
			values (NULLIF(?, ''),?,?,?,?,?,?,?,?,?,?,?,?);`

	result, err := model.DB.Exec(statement, dog.Regnum, dog.Nsdtrcregnum, dog.Sequencenum, dog.Litterid, 
		dog.Name, dog.Callname, dog.Whelpdate, dog.Sex, dog.Nba, dog.Alive, dog.Intact, dog.Sire, dog.Dam)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}
