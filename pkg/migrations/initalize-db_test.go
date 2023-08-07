package migrations

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestInitialize(t *testing.T) {
	t.Run("testing initialize", func(t *testing.T) {

		db, err := openDB("lougar:thewarrior@/tollerdata?parseTime=true&multiStatements=true")
		defer db.Close()
		if err != nil {
			t.Fatalf("failed to open connection to DB: %s", err)
		}

		initDB := InitModel{
			DB: db,
		}

		initDB.RunScript("./sql/teardown.sql")
		err = initDB.RunScript("./sql/setup.sql")

		if err != nil {
			t.Fatalf("failed to execute db script: %s", err)
		}
	})
}

func TestTeardown(t *testing.T) {
	t.Run("testing initialize", func(t *testing.T) {

		db, err := openDB("lougar:thewarrior@/tollerdata?parseTime=true&multiStatements=true")
		defer db.Close()
		if err != nil {
			t.Fatalf("failed to open connection to DB: %s", err)
		}

		initDB := InitModel{
			DB: db,
		}

		err = initDB.RunScript("./sql/teardown.sql")

		if err != nil {
			t.Fatalf("failed to execute db script: %s", err)
		}
	})
}
