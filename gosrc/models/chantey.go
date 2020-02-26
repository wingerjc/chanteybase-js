package models

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

const (
	// Chantey types
	TYPE_HALYARD    = "halyard"
	TYPE_CAPSTAN    = "capstan"
	TYPE_FOCSLE     = "focsle"
	TYPE_SHORT_HAUL = "short haul"
	TYPE_BUNTING    = "bunting"
	// Location types
	LOC_PAGE   = "page"
	LOC_SECONG = "second"
	LOC_TRACK  = "track"
)

func LoadChanteyConfig(dialect *SqlDialect) *DatabaseModel {
	conf := ModelConfig{
		Create: `CREATE TABLE IF NOT EXISTS chantey(
        id $TEXT PRIMARY KEY,
        tune_ids $TEXT NOT NULL,
	    collection_id $TEXT NOT NULL,
	    collection_location INTEGER NOT NULL,
	    location_type $TEXT NOT NULL,
	    version $TEXT NOT NULL,
	    performer_id $TEXT NOT NULL,
        title $TEXT NOT NULL,
	    themes $TEXT NOT NULL,
	    types $TEXT NOT NULL,
        lyrics $TEXT NOT NULL,
	    abc $TEXT NOT NULL,
	    CONSTRAINT perfomer_fk
		  FOREIGN KEY (performer_id)
		  REFERENCES person(id),
	    CONSTRAINT collection_fk
		  FOREIGN KEY (collection_id)
		  REFERENCES collection(id),
	    CONSTRAINT location_type_fk
		  FOREIGN KEY (location_type)
		  REFERENCES location_type(type)
    	);`,
		Constraints: "",
	}
	return NewDatabaseModel(dialect, conf)
}

type Chantey struct {
	ID                 string `db:"id" json:"id"`
	TuneIDs            string `db:"tune_ids" json:"tune-ids"`
	CollectionID       string `db:"collection_id" json:"collection-id"`
	CollectionLocation int    `db:"collection_location" json:"collection-location"`
	LocationType       string `db:"location_type" json:"location-type"`
	Version            string `db:"version" json:"version"`
	PerformerId        string `db:"performer_id" json:"performer-id"`
	Title              string `db:"title" json:"title"`
	Themes             string `db:"themes" json:"themes"`
	Types              string `db:"types" json:"types"`
	Lyrics             string `db:"lyrics" json:"lyrics"`
	ABC                string `db:"abc" json:"abc"`
}

func (c *Chantey) Write(tx *sql.Tx, dialect SqlDialect) (sql.Result, error) {
	statement := dialect.replaceInsertPrefix + `INTO
	chantey (id, tune_ids, collection_id, collection_location, location_type, version, performer_id, title, themes, types, lyrics, abc)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);`
	fmt.Println(c)
	return tx.Exec(
		statement,
		c.ID,
		c.TuneIDs,
		c.CollectionID,
		c.CollectionLocation,
		c.LocationType,
		c.Version,
		c.PerformerId,
		c.Title,
		c.Themes,
		c.Types,
		c.Lyrics,
		c.ABC,
	)
}

func WriteChanteys(db *sqlx.DB, chanteys []*Chantey, dialect SqlDialect) error {
	tx, err := db.BeginTx(context.TODO(), &sql.TxOptions{ReadOnly: false})
	if err != nil {
		return err
	}
	for _, c := range chanteys {
		if _, err := c.Write(tx, dialect); err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

type ChanteyJson struct {
	TuneIDs            []string `json:"tune-ids"`
	CollectionID       string   `json:"collection-id"`
	CollectionLocation int      `json:"collection-location"`
	LocationType       string   `json:"location-type"`
	Version            string   `json:"version"`
	PerfomerID         string   `json:"performer-id"`
	Title              []string `json:"title"`
	Themes             []string `json:"themes"`
	Types              []string `json:"types"`
	Lyrics             []string `json:"lyrics"`
	ABC                []string `json:"ABC"`
}

func (c *ChanteyJson) ToDBChantey() *Chantey {
	return &Chantey{
		ID:                 c.ID(),
		TuneIDs:            strings.Join(c.TuneIDs, "\n"),
		CollectionID:       c.CollectionID,
		CollectionLocation: c.CollectionLocation,
		LocationType:       c.LocationType,
		PerformerId:        c.PerfomerID,
		Title:              strings.Join(c.Title, "\n"),
		Themes:             strings.Join(c.Themes, "\n"),
		Types:              strings.Join(c.Types, "\n"),
		Lyrics:             strings.Join(c.Lyrics, "\n"),
		ABC:                strings.Join(c.ABC, "\n"),
	}
}

func (c *ChanteyJson) ID() string {
	b := strings.Builder{}
	b.WriteString(convertKeyString(c.Title, 8))
	if c.CollectionLocation >= 0 {
		b.WriteString(".")
		b.WriteString(strconv.Itoa(c.CollectionLocation))
	}
	if len(c.Version) > 0 {
		b.WriteString(".")
		b.WriteString(c.Version)
	}
	b.WriteString(".")
	b.WriteString(c.CollectionID)
	return b.String()
}
