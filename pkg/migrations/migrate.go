package migrations

import (
	"database/sql"
)

type MigrateModel struct {
	DB, LegacyDB *sql.DB
}

func (model *MigrateModel) migrateDogs() error {

	return nil
}
