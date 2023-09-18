package models

import "database/sql"

type Health struct {
	Id         int
	Dogid      int
	HealthType string
	CertId     string
}

type HealthModel struct {
	DB *sql.DB
}

func (model *HealthModel) Insert(health Health) error {

	statement := `insert into health (dogid, healthtype, certid) values (?, ?, ?);`

	result, err := model.DB.Exec(statement, health.Dogid, health.HealthType, health.CertId)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (model *HealthModel) GetByDog(dogid int) ([]*Health, error) {
	healthList := []*Health{}

	query := `select id, dogid, healthtype, certid
		from health
		where dogid = ?`

	rows, err := model.DB.Query(query, dogid)
	if err != nil {
		return healthList, err
	}
	defer rows.Close()

	for rows.Next() {
		h := &Health{}
		err = rows.Scan(&h.Id, &h.Dogid, &h.HealthType, &h.CertId)
		healthList = append(healthList, h)
	}

	return healthList, nil
}
