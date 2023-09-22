package models

import (
	"testing"
)

func TestInsertColor(t *testing.T) {

	db := NewTestDB(t)

	colorModel := ColorModel{DB: db}
	health := Color{Dogid: 1, ColorCode: "CR"}
	err := colorModel.Insert(health)
	if err != nil {
		t.Fatal(err)
	}

}

func TestGetDogColors(t *testing.T) {

	db := NewTestDB(t)

	colorModel := ColorModel{DB: db}
	health := Color{Dogid: 1, ColorCode: "CR"}
	err := colorModel.Insert(health)
	if err != nil {
		t.Fatal(err)
	}

	gotColor, err := colorModel.GetDogColors(1)
	if err != nil {
		t.Fatal(err)
	}

	if len(gotColor) != 1 {
		t.Fatal("did not get correct number of health records")
	}

}
