package models

import (
	"database/sql"
	"path/filepath"
	"strings"
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

type ChanteyJson struct {
	TuneIDs            []string `json:"tune-ids"`
	CollectionID       string   `json:"collection-id"`
	CollectionLocation int      `json:"collection-location"`
	Title              []string `json:"title"`
	Themes             []string `json:"themes"`
	Lyrics             []string `json:"lyrics"`
	ABC                []string `json:"ABC"`
}

func (c *ChanteyJson) ToDBChantey() *Chantey {
	location := sql.NullInt64{}
	if c.CollectionLocation < 0 {
		location = toNullInt(c.CollectionLocation)
	}
	abcStr := strings.Join(c.ABC, "\n")
	abc := sql.NullString{}
	if len(abcStr) > 0 {
		abc = toNullString(abcStr)
	}
	return &Chantey{
		ID:                 c.ID(),
		TuneIDs:            strings.Join(c.TuneIDs, "\n"),
		CollectionID:       c.CollectionID,
		CollectionLocation: location,
		Title:              strings.Join(c.Title, "\n"),
		Themes:             strings.Join(c.Themes, "\n"),
		Lyrics:             strings.Join(c.Lyrics, "\n"),
		ABC:                abc,
	}
}

func (c *ChanteyJson) ID() string {
	b := strings.Builder{}
	b.WriteString(convertKeyString(c.Title, 8))
	b.WriteString(".")
	b.WriteString(c.CollectionID)
	return b.String()
}
