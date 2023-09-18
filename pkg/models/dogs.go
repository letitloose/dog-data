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
			values (NULLIF(?, ''),?,?,?,?,?,?,?,?,?,?,NULLIF(?, 0),NULLIF(?, 0));`

	result, err := model.DB.Exec(statement, dog.Regnum, dog.Nsdtrcregnum,
		dog.Sequencenum, dog.Litterid, dog.Name, dog.Callname, dog.Whelpdate,
		dog.Sex, dog.Nba, dog.Alive, dog.Intact, dog.Sire, dog.Dam)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (model *DogModel) Update(dog Dog) error {

	statement := `update dogs set regnum = ?, 
		nsdtrcregnum = ?, sequencenum = ?, litterid = ?, name = ?, callname = ?,
		whelpdate = ?, sex = ?, nba = ?, alive = ?, intact = ?, sire = NULLIF(?, 0), 
		dam = NULLIF(?, 0)
		where id = ?;`

	result, err := model.DB.Exec(statement, dog.Regnum, dog.Nsdtrcregnum,
		dog.Sequencenum, dog.Litterid, dog.Name, dog.Callname, dog.Whelpdate,
		dog.Sex, dog.Nba, dog.Alive, dog.Intact, dog.Sire, dog.Dam, dog.Id)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (model *DogModel) GetByRegnum(regnum string) (Dog, error) {
	dog := Dog{}

	query := `select id, regnum, nsdtrcregnum, sequencenum, litterid, name, callname,
		whelpdate, sex, nba, alive, intact, sire, dam
		from dogs
		where regnum = ?`

	var sire, dam sql.NullInt16
	row := model.DB.QueryRow(query, regnum)
	err := row.Scan(&dog.Id, &dog.Regnum, &dog.Nsdtrcregnum, &dog.Sequencenum,
		&dog.Litterid, &dog.Name, &dog.Callname, &dog.Whelpdate, &dog.Sex,
		&dog.Nba, &dog.Alive, &dog.Intact, &sire, &dam)
	if err != nil {
		return dog, err
	}
	dog.Sire = int(sire.Int16)
	dog.Dam = int(dam.Int16)

	return dog, nil
}

func (model *DogModel) GetByName(name string) (Dog, error) {
	dog := Dog{}

	query := `select id, regnum, nsdtrcregnum, sequencenum, litterid, name, callname,
		whelpdate, sex, nba, alive, intact, sire, dam
		from dogs
		where UPPER(name) = UPPER(?)`

	var sire, dam sql.NullInt16
	row := model.DB.QueryRow(query, name)
	err := row.Scan(&dog.Id, &dog.Regnum, &dog.Nsdtrcregnum, &dog.Sequencenum,
		&dog.Litterid, &dog.Name, &dog.Callname, &dog.Whelpdate, &dog.Sex,
		&dog.Nba, &dog.Alive, &dog.Intact, &sire, &dam)
	if err != nil {
		return dog, err
	}
	dog.Sire = int(sire.Int16)
	dog.Dam = int(dam.Int16)

	return dog, nil
}
