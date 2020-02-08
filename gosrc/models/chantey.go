package models

import (
	"database/sql"
	"path/filepath"
)

func LoadChanteyConfig(path string, dialect SqlDialect) *DatabaseModel {
	return NewDatabaseModel(filepath.Join(path, "chantey.json"), dialect)
}

type Chantey struct {
	ID                 string         `db:"id"`
	TuneIDs            string         `db:"tune_ids"`
	CollectionID       string         `db:"collection_id"`
	CollectionLocation sql.NullInt64  `db:"collection_location"`
	Title              string         `db:"title"`
	Themes             string         `db:"themes"`
	Lyrics             string         `db:"lyrics"`
	ABC                sql.NullString `db:"abc"`
}
