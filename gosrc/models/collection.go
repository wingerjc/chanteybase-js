package models

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
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
		collector_id $TEXT NOT NULL,
		oclc $TEXT NOT NULL,
		lccn $TEXT NOT NULL,
		isbn $TEXT NOT NULL,
		CONSTRAINT collector_fk,
		  FOREIGN KEY (collector_id)
		  REFERENCES person(id)
		);`,
		Constraints: "",
	}
	return NewDatabaseModel(dialect, conf)
}

type Collection struct {
	ID              string `db:"id" json:"id"`
	Title           string `db:"title" json:"title"`
	Volume          int    `db:"volume" json:"volume"`
	PublicationYear int    `db:"publication_year" json:"publication-year"`
	Edition         int    `db:"edition" json:"edition"`
	CollectorID     string `db:"collector_id" json:"collector-id"`
	OCLC            string `db:"oclc" json:"oclc"`
	LCCN            string `db:"lccn" json:"lccn"`
	ISBN            string `db:"isbn" json:"isbn"`
}

type CollectionJson struct {
	Title           []string `json:"title"`
	Volume          int      `json:"volume"`
	PublicationYear int      `json:"publication-year"`
	Edition         int      `json:"edition"`
	CollectorID     string   `json:"collector-id"`
	OCLC            string   `json:"oclc"`
	LCCN            string   `json:"lccn"`
	ISBN            string   `json:"isbn"`
}

func (c *CollectionJson) ToDBCollection() *Collection {
	return &Collection{
		ID:              c.ID(),
		Title:           strings.Join(c.Title, "\n"),
		Volume:          c.Volume,
		PublicationYear: c.PublicationYear,
		Edition:         c.Edition,
		CollectorID:     c.CollectorID,
		OCLC:            c.OCLC,
		LCCN:            c.LCCN,
		ISBN:            c.ISBN,
	}
}

func (c *CollectionJson) ID() string {
	b := strings.Builder{}
	b.WriteString(convertKeyString(c.Title, 8))
	b.WriteString(".")
	b.WriteString(strconv.Itoa(c.PublicationYear))
	if c.Volume != 0 {
		b.WriteString(".")
		b.WriteString(string(c.Volume))
	}
	return b.String()
}

func (c *Collection) Write(tx *sql.Tx, dialect SqlDialect) (sql.Result, error) {
	statement := dialect.replaceInsertPrefix + `INTO
	collection (id, title, volume, publication_year, edition, collector_id, oclc, lccn, isbn)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);`
	fmt.Println(c)
	return tx.Exec(
		statement,
		c.ID,
		c.Title,
		c.Volume,
		c.PublicationYear,
		c.Edition,
		c.CollectorID,
		c.OCLC,
		c.LCCN,
		c.ISBN,
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
