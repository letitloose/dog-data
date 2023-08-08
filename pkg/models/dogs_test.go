package models

import (
	"testing"
)

func TestInsert(t *testing.T) {

	db := NewTestDB(t)

	dogModel := DogModel{DB: db}

	dog := Dog{Regnum: "DOG7", Name: "carl", Litterid: 1, Sex: "SD", Sire: 1, Dam: 1}
	err := dogModel.Insert(dog)
	if err != nil {
		t.Fatal(err)
	}

}
