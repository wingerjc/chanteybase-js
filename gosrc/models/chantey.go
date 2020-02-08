package models

import "database/sql"

const (
	chanteySchema = `CREATE TABLE IF NOT EXISTS chantey(
    id TEXT PRIMARY KEY,
    tune_ids TEXT NOT NULL,
    collection_id TEXT NOT NULL,
    collection_location INTEGER,
    title TEXT NOT NULL,
    themes TEXT NOT NULL,
    lyrics TEXT NOT NULL,
    abc TEXT
    );`
	chanteyConstraints = `-- fk constraints for the TABLE`
)

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
