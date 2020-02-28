package models

import (
	"strings"

	"github.com/jmoiron/sqlx"
)

func chanteyQuery(db *sqlx.DB, sql, searchString string) ([]Chantey, error) {
	result := []Chantey{}
	err := db.Select(&result, sql, searchString)
	return result, err
}

// ChanteyByID returns a list of chanteys that have IDs like the search term.
func ChanteyByID(db *sqlx.DB, idString string) ([]Chantey, error) {
	return chanteyQuery(
		db,
		`SELECT * FROM chantey WHERE id LIKE $1;`,
		dbSearchString(strings.ToUpper(NON_ID_CHAR_REGEX.ReplaceAllString(idString, ""))),
	)
}

// ChanteyByName returns a list of chanteys that have names like the search term.
func ChanteyByName(db *sqlx.DB, nameString string) ([]Chantey, error) {
	return chanteyQuery(
		db,
		`SELECT * FROM chantey WHERE name LIKE $1;`,
		dbSearchString(nameString),
	)
}

// ChanteyByCollectionID returns a list of chanteys that have an exact match on collection ID.
func ChanteyByCollectionID(db *sqlx.DB, collectionID string) ([]Chantey, error) {
	return chanteyQuery(
		db,
		`SELECT * FROM chantey WHERE collection_id = $1 ORDER BY collection_location;`,
		strings.ToUpper(collectionID),
	)
}
