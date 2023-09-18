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
	CategorySex = "sex"
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
