package models

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

func LoadCollectionConfig(dialect *SqlDialect) *DatabaseModel {
	conf := ModelConfig{
		Create: `CREATE TABLE IF NOT EXISTS collection (
		id $TEXT PRIMARY KEY,
		title $TEXT NOT NULL,
		volume $INT NOT NULL,
		publication_year $INT NOT NULL,
		edition $INT NOT NULL,
		collector_id $TEXT NOT NULL
		CONSTRAINT collector_fk,
		  FOREIGN KEY (collector_id)
		  REFERENCES person(id)
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
	Edition         int    `db:"edition"`
	CollectorID     string `db:"collector_id"`
}

type CollectionJson struct {
	Title           []string `json:"title"`
	Volume          int      `json:"volume"`
	PublicationYear int      `json:"publication-year"`
	Edition         int      `json:"edition"`
	CollectorID     string   `json:"collector-id"`
}

func (c *CollectionJson) ToDBCollection() *Collection {
	return &Collection{
		ID:              c.ID(),
		Title:           strings.Join(c.Title, "\n"),
		Volume:          c.Volume,
		PublicationYear: c.PublicationYear,
		Edition:         c.Edition,
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

func (c *Collection) Write(tx *sql.Tx, dialect SqlDialect) (sql.Result, error) {
	statement := dialect.replaceInsertPrefix + `INTO
	collection (id, title, volume, publication_year, edition, collector_id)
	VALUES ($1, $2, $3, $4, $5, $6);`
	fmt.Println(c)
	return tx.Exec(
		statement,
		c.ID,
		c.Title,
		c.Volume,
		c.PublicationYear,
		c.Edition,
		c.CollectorID,
	)
}

func WriteCollections(db *sqlx.DB, collections []*Collection, dialect SqlDialect) error {
	tx, err := db.BeginTx(context.TODO(), &sql.TxOptions{ReadOnly: false})
	if err != nil {
		return err
	}
	for _, c := range collections {
		if _, err := c.Write(tx, dialect); err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}
