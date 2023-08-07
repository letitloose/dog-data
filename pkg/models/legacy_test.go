package models

import (
	"testing"
)

func TestGetTollers(t *testing.T) {
	db := NewTestDB(t)

	model := LegacyModel{DB: db}
	tollers, err := model.GetTollers()
	if err != nil {
		t.Fatal(err)
	}

	if len(tollers) != 1 {
		t.Fatal("wrong number of tollers")
	}
}
