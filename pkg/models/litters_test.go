package models

import (
	"testing"
)

func TestInsertLitter(t *testing.T) {

	db := NewTestDB(t)

	litterModel := LitterModel{DB: db}

	litter := Litter{Regnum: "litter2"}
	err := litterModel.Insert(litter)
	if err != nil {
		t.Fatal(err)
	}

}

func TestGetLitter(t *testing.T) {

	db := NewTestDB(t)

	litterModel := LitterModel{DB: db}

	litter := Litter{Regnum: "litter2"}
	err := litterModel.Insert(litter)
	if err != nil {
		t.Fatal(err)
	}

	gotLitter, err := litterModel.GetByRegnum("litter2")
	if err != nil {
		t.Fatal(err)
	}

	if gotLitter.Id != 2 {
		t.Fatal("did not get litter")
	}

}
