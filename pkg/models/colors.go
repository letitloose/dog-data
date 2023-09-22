package models

import "database/sql"

type Color struct {
	Id        int
	Dogid     int
	ColorCode string
	CertId    string
}

type ColorModel struct {
	DB *sql.DB
}

func (model *ColorModel) Insert(color Color) error {

	statement := `insert into colors (dogid, colorcode) values (?, ?);`

	result, err := model.DB.Exec(statement, color.Dogid, color.ColorCode)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (model *ColorModel) GetDogColors(dogid int) ([]*Color, error) {
	colorList := []*Color{}

	query := `select id, dogid, colorcode
		from colors
		where dogid = ?`

	rows, err := model.DB.Query(query, dogid)
	if err != nil {
		return colorList, err
	}
	defer rows.Close()

	for rows.Next() {
		c := &Color{}
		err = rows.Scan(&c.Id, &c.Dogid, &c.ColorCode)
		colorList = append(colorList, c)
	}

	return colorList, nil
}
