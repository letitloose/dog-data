package models

import (
	"database/sql"
	"time"
)

type LegacyModel struct {
	DB *sql.DB
}

type Toller struct {
	id             int
	regnum         string
	nsdtrcregnum   string
	sequencenum    string
	name           string
	titlename      string
	litterregnum   string
	sex            string
	color          string
	pra            string
	hipClear       string
	eyeClear       string
	heartClear     string
	elbowClear     string
	whelpdate      time.Time
	nba            string
	alive          string
	owner          string
	address        string
	intact         string
	city           string
	state          string
	zip            string
	country        string
	sireregnum     string
	sirename       string
	damregnum      string
	damname        string
	breederName    string
	breederAddress string
	breederCity    string
	breederState   string
	breederZip     string
	breederCountry string
	callname       string
}

func (model LegacyModel) GetTollers() ([]*Toller, error) {

	tollers := []*Toller{}

	query := `select a_registra, a_nsdtrc_r, a_seq, a_dogname, a_title_na, a_litterre, a_sex, a_color, 
		a_pra, a_hipclear, a_eyeclear, a_heart_cl, a_elbow_cl, a_nba, a_alive, a_owner, a_address1,
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
		err = rows.Scan(&t.regnum, &t.nsdtrcregnum, &t.sequencenum, &t.name, &t.titlename, &t.litterregnum,
			&t.sex, &t.color, &t.pra, &t.hipClear, &t.eyeClear, &t.heartClear, &t.elbowClear, &t.nba, &t.alive,
			&t.owner, &t.address, &t.intact, &t.city, &t.state, &t.zip, &t.country, &t.sireregnum, &t.sirename,
			&t.damregnum, &t.damname, &t.breederName, &t.breederAddress, &t.breederCity, &t.breederState,
			&t.breederZip, &t.breederCountry, &t.callname)
		tollers = append(tollers, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tollers, nil

}
