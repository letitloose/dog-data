package models

import "database/sql"

type Address struct {
	Id       int
	Address1 string
	Address2 string
	City     string
	State    string
	Zip      string
	Country  string
}

type AddressModel struct {
	DB *sql.DB
}

func (model *AddressModel) Insert(address Address) error {

	statement := `insert into addresses (address1, address2, city, state, zip, country) 
		values (?, ?, ?, ?, ?, ?);`

	result, err := model.DB.Exec(statement, address.Address1, address.Address2,
		address.City, address.State, address.Zip, address.Country)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (model *AddressModel) Get(id int) (Address, error) {
	address := Address{}

	query := `select id, address1, address2, city, state, zip, country
		from addresses
		where id = ?`

	row := model.DB.QueryRow(query, id)
	err := row.Scan(&address.Id, &address.Address1, &address.Address2, &address.City,
		&address.State, &address.Zip, &address.Country)
	if err != nil {
		return address, err
	}

	return address, nil
}
