package migrations

import (
	"testing"
)

func TestImportLegacy(t *testing.T) {
	t.Run("testing initialize legacy db", func(t *testing.T) {

		db, err := openDB("lougar:thewarrior@/legacytollerdata?parseTime=true&multiStatements=true")
		defer db.Close()
		if err != nil {
			t.Fatalf("failed to open connection to DB: %s", err)
		}

		initDB := &InitLegacyModel{
			DB:        db,
			filename:  "./dbf/toller.dbf",
			tablename: "tollers",
		}

		err = initDB.importLegacy()

		if err != nil {
			t.Fatalf("failed to execute db script: %s", err)
		}
	})
}

func TestImportDBFDir(t *testing.T) {
	t.Run("testing importing all DBFs", func(t *testing.T) {

		db, err := openDB("lougar:thewarrior@/legacytollerdata?parseTime=true&multiStatements=true")
		defer db.Close()
		if err != nil {
			t.Fatalf("failed to open connection to DB: %s", err)
		}

		initDB := &InitLegacyModel{
			DB: db,
		}

		err = initDB.ImportDBFDir("./dbf")

		if err != nil {
			t.Fatalf("failed to execute db script: %s", err)
		}
	})
}
