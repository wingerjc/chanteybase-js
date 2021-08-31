package models

import (
	"strings"

	"github.com/jmoiron/sqlx"
)

// GetPersonByID fetches a person, by exact ID match.
func GetPersonByID(db *sqlx.DB, id string) ([]Person, error) {
	result := []Person{}

	sql := `SELECT * FROM person WHERE id = $1`
	if err := db.Select(&result, sql, strings.ToUpper(id)); err != nil {
		return nil, err
	}

	return result, nil
}

// GetPersonIDs fetches a list of person IDs by similar id match.
func GetPersonIDs(db *sqlx.DB, pattern string) ([]string, error) {
	result := []string{}
	searchString := "%" + strings.ToUpper(nonIDCharRegex.ReplaceAllString(pattern, "")) + "%"
	sql := `SELECT id FROM person WHERE id like $1;`
	err := db.Select(&result, sql, searchString)
	return result, err
}
