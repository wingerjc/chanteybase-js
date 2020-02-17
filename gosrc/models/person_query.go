package models

import "github.com/jmoiron/sqlx"

func GetPersonByID(db *sqlx.DB, id string) ([]Person, error) {
	result := []Person{}

	sql := `SELECT *  FROM person WHERE id = $1`
	if err := db.Select(&result, sql, id); err != nil {
		return nil, err
	}

	return result, nil
}
