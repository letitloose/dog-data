package models

import (
	"database/sql"
)

type Codetable struct {
	Id       string
	Category string
	Code     string
	Display  string
}

type CodetableModel struct {
	DB *sql.DB
}

const (
	CategorySex    = "sex"
	CategoryHealth = "health"
)

func (model *CodetableModel) GetByCode(code, category string) Codetable {

	codetable := Codetable{}

	query := `select id, category, code, display
		from codetables
		where code = ? 
		and category = ?`

	row := model.DB.QueryRow(query, code, category)
	err := row.Scan(&codetable.Id, &codetable.Category, &codetable.Code, &codetable.Display)
	if err != nil {
		return codetable
	}

	return codetable
}

func (model *CodetableModel) GetByCategory(category string) ([]*Codetable, error) {

	codetables := []*Codetable{}

	query := `select id, category, code, display
		from codetables
		where category = ?`

	rows, err := model.DB.Query(query, category)
	if err != nil {
		return codetables, err
	}
	defer rows.Close()

	for rows.Next() {
		ct := &Codetable{}
		err = rows.Scan(&ct.Id, &ct.Category, &ct.Code, &ct.Display)
		codetables = append(codetables, ct)
	}

	return codetables, nil
}
