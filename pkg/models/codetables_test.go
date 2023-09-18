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
