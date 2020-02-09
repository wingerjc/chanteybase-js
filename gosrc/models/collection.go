package models

import (
	"path/filepath"
	"strings"
)

func LoadCollectionConfig(path string, dialect SqlDialect) *DatabaseModel {
	return NewDatabaseModel(filepath.Join(path, "collection.json"), dialect)
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
