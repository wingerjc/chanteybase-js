package models

import (
	"regexp"
	"strings"

	"github.com/jmoiron/sqlx"
)

func GetPersonByID(db *sqlx.DB, id string) ([]Person, error) {
	result := []Person{}

	sql := `SELECT * FROM person WHERE id = $1`
	if err := db.Select(&result, sql, id); err != nil {
		return nil, err
	}

	return result, nil
}

func GetPersonIDs(db *sqlx.DB, pattern string) ([]string, error) {
	replacePattern, err := regexp.Compile(`[^A-Za-z0-9.]+`)
	if err != nil {
		return nil, err
	}
	result := []string{}

	searchString := "%" + strings.ToUpper(replacePattern.ReplaceAllString(pattern, "")) + "%"
	sql := `SELECT id FROM person WHERE id like $1;`
	err = db.Select(&result, sql, searchString)
	return result, err
}
