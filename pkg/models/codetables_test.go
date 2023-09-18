package models

import "testing"

func TestGetByCode(t *testing.T) {
	db := NewTestDB(t)

	model := CodetableModel{DB: db}

	value := model.GetByCode("D", "sex")

	if value.Display != "Dog" {
		t.Fatalf("wrong value, expected: %s, got : %s", "Dog", value.Display)
	}
}

func TestGetByCategory(t *testing.T) {
	db := NewTestDB(t)

	model := CodetableModel{DB: db}

	value, err := model.GetByCategory("sex")
	if err != nil {
		t.Fatal(err)
	}

	if len(value) != 2 {
		t.Fatal("wrong number of codetables returned")
	}
}
