package models

import (
	"testing"
)

func TestGetTollers(t *testing.T) {
	db := NewTestDB(t)

	model := LegacyModel{DB: db}
	tollers, err := model.GetTollers()
	if err != nil {
		t.Fatal(err)
	}

	if len(tollers) != 1 {
		t.Fatal("wrong number of tollers")
	}

	for _, toller := range tollers {
		if toller.Regnum != "regnum" {
			t.Fatalf("wrong regnum, expected: %s got: %s", "regnum", toller.Regnum)
		}
		if toller.Nsdtrcregnum != "nsdtrcregnum" {
			t.Fatalf("wrong nsdtrcregnum, expected: %s got: %s", "nsdtrcregnum", toller.Nsdtrcregnum)
		}
		if toller.Nba != "nba" {
			t.Fatalf("wrong nba, expected: %s got: %s", "nba", toller.Nba)
		}
		if toller.Sequencenum != "seqnum" {
			t.Fatalf("wrong seqnum, expected: %s got: %s", "seqnum", toller.Sequencenum)
		}
		if toller.Litterregnum != "litterrn" {
			t.Fatalf("wrong litterrn, expected: %s got: %s", "litterrn", toller.Litterregnum)
		}
		if toller.Callname != "callname" {
			t.Fatalf("wrong callname, expected: %s got: %s", "callname", toller.Callname)
		}
	}

	// ("regnum", "nsdtrcregnum","seqnum","name","titlename","litterrn","sex","color","pra",
	// "hipclear","eyeclear","heartclear","elbowclear",NOW(),"nba","Y","owner","address1","Y","city",
	// "state","zip", "country","sireregnum","sirename","damregnum","damname","breedername","breederaddress",
	// "breedercity","brst","breederzip","breedercountry","","callname","");
}
