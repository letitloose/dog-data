package models

import (
	"testing"
)

func TestInsertAddress(t *testing.T) {

	db := NewTestDB(t)

	addressModel := AddressModel{DB: db}

	address := Address{Address1: "1", Address2: "2", City: "troy",
		State: "SNY", Zip: "11111", Country: "COUSA"}
	err := addressModel.Insert(address)
	if err != nil {
		t.Fatal(err)
	}

}

func TestGet(t *testing.T) {

	db := NewTestDB(t)

	addressModel := AddressModel{DB: db}

	address := Address{Address1: "1", Address2: "2", City: "troy",
		State: "SNY", Zip: "11111", Country: "COUSA"}
	err := addressModel.Insert(address)
	if err != nil {
		t.Fatal(err)
	}

	gotAddress, err := addressModel.Get(1)
	if err != nil {
		t.Fatal(err)
	}

	if gotAddress.Address1 != "1" {
		t.Fatalf("did not get correct address, expected 1 got %s", gotAddress.Address1)
	}

}
