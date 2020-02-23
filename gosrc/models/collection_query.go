package models

import (
	"strings"

	"github.com/jmoiron/sqlx"
)

func collectionQuery(db *sqlx.DB, sql, searchString string) ([]Collection, error) {
	result := []Collection{}
	err := db.Select(&result, sql, searchString)
	return result, err
}

func CollectionByID(db *sqlx.DB, id string) ([]Collection, error) {
	return collectionQuery(
		db,
		`SELECT * FROM collection WHERE id LIKE $1;`,
		dbSearchString(strings.ToUpper(NON_ID_CHAR_REGEX.ReplaceAllString(id, ""))),
	)
}

func CollectionByTitle(db *sqlx.DB, name string) ([]Collection, error) {
	return collectionQuery(
		db,
		`SELECT * FROM collection WHERE title LIKE $1;`,
		dbSearchString(name),
	)
}
