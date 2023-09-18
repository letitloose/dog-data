package models

import (
	"testing"
)

func TestInsert(t *testing.T) {

	db := NewTestDB(t)

	dogModel := DogModel{DB: db}

	dog := Dog{Regnum: "DOG7", Name: "carl", Litterid: 1, Sex: "SD"}
	err := dogModel.Insert(dog)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdate(t *testing.T) {

	db := NewTestDB(t)

	dogModel := DogModel{DB: db}

	dog := Dog{Regnum: "DOG7", Name: "carl", Litterid: 1, Sex: "SD"}
	err := dogModel.Insert(dog)
	if err != nil {
		t.Fatal(err)
	}

	dog1, err := dogModel.GetByRegnum("DOG7")
	if err != nil {
		t.Fatal(err)
	}

	dog.Id = dog1.Id
	dog.Regnum = "DOG8"
	err = dogModel.Update(dog)
	if err != nil {
		t.Fatal(err)
	}

	dog1, err = dogModel.GetByRegnum("DOG8")
	if err != nil {
		t.Fatal(err)
	}
	if dog1.Name != "carl" {
		t.Fatalf("got the wrong dog.  expectd %s, got: %s", "carl", dog1.Name)
	}
}

func TestGetByRegnum(t *testing.T) {
	db := NewTestDB(t)
	dogModel := DogModel{DB: db}

	dog := Dog{Regnum: "DOG7", Name: "carl", Litterid: 1, Sex: "SD", Sire: 1, Dam: 1}
	err := dogModel.Insert(dog)
	if err != nil {
		t.Fatal(err)
	}

	dog1, err := dogModel.GetByRegnum("DOG7")
	if err != nil {
		t.Fatal(err)
	}

	if dog1.Name != "carl" {
		t.Fatalf("got the wrong dog.  expectd %s, got: %s", "carl", dog1.Name)
	}
}

func TestGetByName(t *testing.T) {
	db := NewTestDB(t)
	dogModel := DogModel{DB: db}

	dog := Dog{Regnum: "DOG7", Name: "carl", Litterid: 1, Sex: "SD", Sire: 1, Dam: 1}
	err := dogModel.Insert(dog)
	if err != nil {
		t.Fatal(err)
	}

	dog1, err := dogModel.GetByName("carl")
	if err != nil {
		t.Fatal(err)
	}

	if dog1.Regnum != "DOG7" {
		t.Fatalf("got the wrong dog.  expectd %s, got: %s", "carl", dog1.Regnum)
	}
}
