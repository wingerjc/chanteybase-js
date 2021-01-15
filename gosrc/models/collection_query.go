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

// CollectionByID returns all collections where the id is like the search string.
func CollectionByID(db *sqlx.DB, id string) ([]Collection, error) {
	return collectionQuery(
		db,
		`SELECT * FROM collection WHERE id LIKE $1;`,
		dbSearchString(strings.ToUpper(nonIDCharRegex.ReplaceAllString(id, ""))),
	)
}

// CollectionByTitle returns all the collections where the title is like the search string.
func CollectionByTitle(db *sqlx.DB, name string) ([]Collection, error) {
	return collectionQuery(
		db,
		`SELECT * FROM collection WHERE title LIKE $1;`,
		dbSearchString(name),
	)
}
