package models

import (
	"testing"
)

func TestInsertLitter(t *testing.T) {

	db := NewTestDB(t)

	litterModel := LitterModel{DB: db}

	litter := Litter{regnum: "litter2"}
	err := litterModel.Insert(litter)
	if err != nil {
		t.Fatal(err)
	}

}
