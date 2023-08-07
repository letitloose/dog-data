package migrations

import (
	"database/sql"
	"os"
)

type InitModel struct {
	DB *sql.DB
}

func (init *InitModel) RunScript(filename string) error {
	script, err := os.ReadFile(filename)
	if err != nil {
		return (err)
	}
	_, err = init.DB.Exec(string(script))
	if err != nil {
		return (err)
	}
	return nil
}
