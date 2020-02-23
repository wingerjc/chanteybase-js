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

func ChanteyByID(db *sqlx.DB, idString string) ([]Chantey, error) {
	return chanteyQuery(
		db,
		`SELECT * FROM chantey WHERE id LIKE $1;`,
		dbSearchString(strings.ToUpper(NON_ID_CHAR_REGEX.ReplaceAllString(idString, ""))),
	)
}

func ChanteyByName(db *sqlx.DB, nameString string) ([]Chantey, error) {
	return chanteyQuery(
		db,
		`SELECT * FROM chantey WHERE name LIKE $1;`,
		dbSearchString(nameString),
	)
}

func ChanteyByCollectionID(db *sqlx.DB, collectionID string) ([]Chantey, error) {
	return chanteyQuery(
		db,
		`SELECT * FROM chantey WHERE collection_id = $1 ORDER BY collection_location;`,
		strings.ToUpper(collectionID),
	)
}
