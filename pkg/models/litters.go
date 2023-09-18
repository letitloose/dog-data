package models

import "database/sql"

type Litter struct {
	Id     int
	Regnum string
}

type LitterModel struct {
	DB *sql.DB
}

func (model *LitterModel) Insert(litter Litter) error {

	statement := `insert into litters (regnum) values (?);`

	result, err := model.DB.Exec(statement, litter.Regnum)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (model *LitterModel) GetByRegnum(litternum string) (Litter, error) {
	litter := Litter{}

	query := `select id, regnum
		from litters
		where regnum = ?`

	row := model.DB.QueryRow(query, litternum)
	err := row.Scan(&litter.Id, &litter.Regnum)
	if err != nil {
		return litter, err
	}

	return litter, nil
}
