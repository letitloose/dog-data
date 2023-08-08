package models

import (
	"database/sql"
	"time"
)

type LegacyModel struct {
	DB *sql.DB
}

type Toller struct {
	Id             int
	Regnum         string
	Nsdtrcregnum   string
	Sequencenum    string
	Name           string
	Titlename      string
	Litterregnum   string
	Sex            string
	Color          string
	Pra            string
	HipClear       string
	EyeClear       string
	HeartClear     string
	ElbowClear     string
	Whelpdate      time.Time
	Nba            string
	Alive          string
	Owner          string
	Address        string
	Intact         string
	City           string
	State          string
	Zip            string
	Country        string
	Sireregnum     string
	Sirename       string
	Damregnum      string
	Damname        string
	BreederName    string
	BreederAddress string
	BreederCity    string
	BreederState   string
	BreederZip     string
	BreederCountry string
	Callname       string
}

func (model LegacyModel) GetTollers() ([]*Toller, error) {

	tollers := []*Toller{}

	query := `select a_registra, a_nsdtrc_r, a_seq, a_dogname, a_title_na, a_litterre, a_sex, a_color, 
		a_pra, a_hipclear, a_eyeclear, a_heart_cl, a_elbow_cl, a_whelpdat, a_nba, a_alive, a_owner, a_address1,
		a_intact, a_city, a_state, a_zip, a_country, a_sire_reg, a_sirename, a_dam_regn, a_damname,
		a_breeder, a_breedera, a_breederc, a_breeders, a_breederz, a_breeder0, a_callname from tollers;`

	//grab the row
	rows, err := model.DB.Query(query)
	if err != nil {
		return tollers, err
	}
	defer rows.Close()

	//prepare the return slice

	for rows.Next() {
		t := &Toller{}
		err = rows.Scan(&t.Regnum, &t.Nsdtrcregnum, &t.Sequencenum, &t.Name, &t.Titlename, &t.Litterregnum,
			&t.Sex, &t.Color, &t.Pra, &t.HipClear, &t.EyeClear, &t.HeartClear, &t.ElbowClear, &t.Whelpdate, 
			&t.Nba, &t.Alive,
			&t.Owner, &t.Address, &t.Intact, &t.City, &t.State, &t.Zip, &t.Country, &t.Sireregnum, &t.Sirename,
			&t.Damregnum, &t.Damname, &t.BreederName, &t.BreederAddress, &t.BreederCity, &t.BreederState,
			&t.BreederZip, &t.BreederCountry, &t.Callname)
		tollers = append(tollers, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tollers, nil

}
