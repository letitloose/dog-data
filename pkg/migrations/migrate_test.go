package migrations

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/letitloose/dog-data/pkg/models"
)

func TestMigrateDogs(t *testing.T) {

	db := NewTestDB(t)

	model := MigrateModel{
		LegacyModel: models.LegacyModel{
			DB: db,
		},
		DogModel: models.DogModel{
			DB: db,
		},
	}

	err := model.migrateDogs()
	if err != nil {
		t.Fatal(err)
	}

}

func NewTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("mysql", "test_web:pass@/test_dogdata?parseTime=true&multiStatements=true")
	if err != nil {
		t.Fatal(err)
	}

	// Read the setup SQL script from file and execute the statements.
	script, err := os.ReadFile("../models/testdata/setup.sql")
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec(string(script))
	if err != nil {
		t.Fatal(err)
	}

	// Use the t.Cleanup() to register a function *which will automatically be
	// called by Go when the current test (or sub-test) which calls newTestDB()
	// has finished*.
	t.Cleanup(func() {
		script, err := os.ReadFile("../models/testdata/teardown.sql")
		if err != nil {
			t.Fatal(err)
		}
		_, err = db.Exec(string(script))
		if err != nil {
			t.Fatal(err)
		}

		db.Close()
	})

	// Return the database connection pool.
	return db
}
