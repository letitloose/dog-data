package models

import (
	"testing"
)

func TestInsertHealth(t *testing.T) {

	db := NewTestDB(t)

	healthModel := HealthModel{DB: db}
	health := Health{Dogid: 1, HealthType: "HHIP", CertId: "certid"}
	err := healthModel.Insert(health)
	if err != nil {
		t.Fatal(err)
	}

}

func TestGetByDog(t *testing.T) {

	db := NewTestDB(t)

	healthModel := HealthModel{DB: db}
	health := Health{Dogid: 1, HealthType: "HHIP", CertId: "certid"}
	err := healthModel.Insert(health)
	if err != nil {
		t.Fatal(err)
	}

	gotHealth, err := healthModel.GetByDog(1)
	if err != nil {
		t.Fatal(err)
	}

	if len(gotHealth) != 1 {
		t.Fatal("did not get correct number of health records")
	}

}
