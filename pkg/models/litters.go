package models

import "database/sql"

type Litter struct {
	id     int
	regnum string
}

type LitterModel struct {
	DB *sql.DB
}

func (model *LitterModel) Insert(litter Litter) error {

	statement := `insert into litters (regnum) values (?);`

	result, err := model.DB.Exec(statement, litter.regnum)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}
