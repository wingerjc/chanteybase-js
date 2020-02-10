package models

import (
	"strings"
)

func LoadCollectionConfig(dialect *SqlDialect) *DatabaseModel {
	conf := ModelConfig{
		Create: `CREATE TABLE IF NOT EXISTS collection (
		id $TEXT PRIMARY KEY,
		title $TEXT NOT NULL,
		volume $INT NOT NULL,
		publication_year $INT NOT NULL,
		collector_id $TEXT NOT NULL
		);`,
		Constraints: "",
	}
	return NewDatabaseModel(dialect, conf)
}

type Collection struct {
	ID              string `db:"id"`
	Title           string `db:"title"`
	Volume          int    `db:Volume`
	PublicationYear int    `db:"publication_year"`
	CollectorID     string `db:"collector_id"`
}

type CollectionJson struct {
	Title           []string `json:"title"`
	Volume          int      `json:"volume"`
	PublicationYear int      `json:"publication-year"`
	CollectorID     string   `json:"collector-id"`
}

func (c *CollectionJson) ToDBCollection() *Collection {
	return &Collection{
		ID:              c.ID(),
		Title:           strings.Join(c.Title, "\n"),
		Volume:          c.Volume,
		PublicationYear: c.PublicationYear,
		CollectorID:     c.CollectorID,
	}
}

func (c *CollectionJson) ID() string {
	b := strings.Builder{}
	b.WriteString(convertKeyString(c.Title, 8))
	b.WriteString(".")
	b.WriteString(string(c.PublicationYear))
	if c.Volume != 0 {
		b.WriteString(".")
		b.WriteString(string(c.Volume))
	}
	return b.String()
}
