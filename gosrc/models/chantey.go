package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

const (
	// Chantey types

	// TypeHalyard is the type definition for a halyard chantey.
	TypeHalyard = "HALYARD"
	// TypeCapstan is the type definition for a capstan chantey.
	TypeCapstan = "CAPSTAN"
	// TypeForecastle is the type definition for a forecastle chantey.
	TypeForecastle = "FORECASTLE"
	// TypeShortDrag is the type definition for a short drag chantey.
	TypeShortDrag = "SHORT_DRAG "
	// TypeBunting is the type definition for a bunting chantey.
	TypeBunting = "BUNTING"

	// Location types

	// LocationPage is a location type for a paged collection.
	LocationPage = "PAGE"
	// LocationSecond is a location type for a timed collection.
	LocationSecond = "SECOND"
	// LocationTrack is a location type for a tracked collection.
	LocationTrack = "TRACK"
)

// LoadChanteyConfig provides the database config for the chantey table.
func LoadChanteyConfig(dialect *SQLDialect) *DatabaseModel {
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
		);
		CREATE TABLE IF NOT EXISTS chantey_tune(
		chantey_id $TEXT REFERENCES chantey(id),
		tune_id $TEXT REFERENCES tune(id),
		CONSTRAINT chantey_tune_pk
		  PRIMARY KEY (chantey_id, tune_id)
		);`,
		Constraints: "",
	}
	return NewDatabaseModel(dialect, conf)
}

// Chantey encapsulates a Chantey entry for JSON and DB.
type Chantey struct {
	ID                 string `db:"id" json:"id"`
	TuneIDs            string `db:"tune_ids" json:"tune-ids"`
	CollectionID       string `db:"collection_id" json:"collection-id"`
	CollectionLocation int    `db:"collection_location" json:"collection-location"`
	LocationType       string `db:"location_type" json:"location-type"`
	Version            string `db:"version" json:"version"`
	PerformerID        string `db:"performer_id" json:"performer-id"`
	Title              string `db:"title" json:"title"`
	Themes             string `db:"themes" json:"themes"`
	Types              string `db:"types" json:"types"`
	Lyrics             string `db:"lyrics" json:"lyrics"`
	ABC                string `db:"abc" json:"abc"`
}

// Write a chantey into the given DB.
func (c *Chantey) Write(tx *sql.Tx, dialect SQLDialect) (sql.Result, error) {
	// Write Tune ID mappings for later searching.
	tuneIds := strings.Split(c.TuneIDs, "\n")
	for i, t := range tuneIds {
		tuneIds[i] = strings.ToUpper(nonIDCharRegex.ReplaceAllString(t, ""))
	}
	b := strings.Builder{}
	for i, t := range tuneIds {
		if i > 0 {
			b.WriteString(",")
		}
		b.WriteString("\n('" + t + "')")
	}
	b.WriteString(";")
	statement := dialect.replaceInsertPrefix + `INTO tune(id) VALUES` + b.String()
	log.Println(statement)
	if len(tuneIds) > 0 {
		res, err := tx.Exec(statement)
		if err != nil {
			return res, err
		}
	}

	b = strings.Builder{}
	for i, t := range tuneIds {
		if i > 0 {
			b.WriteString(",")
		}
		b.WriteString("('" + c.ID + "', '" + t + "')")
	}
	b.WriteString(";")

	// Write chantey type mappings for later searching.
	// TODO

	// Write root DB entry.
	statement = dialect.replaceInsertPrefix + `INTO chantey_tune(chantey_id, tune_id) VALUES` +
		b.String()
	log.Println(statement)
	if len(tuneIds) > 0 {
		res, err := tx.Exec(statement)
		if err != nil {
			return res, err
		}
	}

	statement = dialect.replaceInsertPrefix + `INTO
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
		c.PerformerID,
		c.Title,
		c.Themes,
		c.Types,
		c.Lyrics,
		c.ABC,
	)
}

// WriteChanteys writes all the chanteys into the database.
func WriteChanteys(db *sqlx.DB, chanteys []*Chantey, dialect SQLDialect) error {
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

// ChanteyJSON is the input format for a Chantey entry.
type ChanteyJSON struct {
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

// ToDBChantey converts the input format into a database entry.
func (c *ChanteyJSON) ToDBChantey() *Chantey {
	return &Chantey{
		ID:                 c.ID(),
		TuneIDs:            strings.Join(c.TuneIDs, "\n"),
		CollectionID:       c.CollectionID,
		CollectionLocation: c.CollectionLocation,
		LocationType:       c.LocationType,
		PerformerID:        c.PerfomerID,
		Title:              strings.Join(c.Title, "\n"),
		Themes:             strings.Join(c.Themes, "\n"),
		Types:              strings.Join(c.Types, "\n"),
		Lyrics:             strings.Join(c.Lyrics, "\n"),
		ABC:                strings.Join(c.ABC, "\n"),
	}
}

// ID creates the identifier for this chantey entry.
func (c *ChanteyJSON) ID() string {
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
