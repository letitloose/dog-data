package models

import (
	"testing"
)

func TestInsert(t *testing.T) {

	db := NewTestDB(t)

	dogModel := DogModel{DB: db}

	dog := Dog{regnum: "DOG7", name: "carl", litterid: 1, sex: "SD", sire: 1, dam: 1}
	err := dogModel.Insert(dog)
	if err != nil {
		t.Fatal(err)
	}

}
