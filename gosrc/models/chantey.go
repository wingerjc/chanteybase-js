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
		);
		CREATE TABLE IF NOT EXISTS chantey_type_match(
		chantey_id $TEXT REFERENCES chantey(id),
		type_id $TEXT REFERENCES chantey_type(type),
		CONSTRAINT chantey_type_match_PK
		  PRIMARY KEY (chantey_id, type_id)
		);
		CREATE TABLE IF NOT EXISTS chantey_theme(
		chantey_id $TEXT REFERENCES chantey(id),
		theme_id $TEXT REFERENCES theme(id),
		CONSTRAINT chantey_theme_pk
		  PRIMARY KEY (chantey_id, theme_id)
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
	// Write root DB entry.
	statement := dialect.InsertStatement(`chantey (id, tune_ids, collection_id, collection_location, location_type, version, performer_id, title, themes, types, lyrics, abc)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`)
	fmt.Println(c)
	if res, err := tx.Exec(
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
	); err != nil {
		return res, err
	}
	// Write Tune ID mappings for later searching.

	tuneIds := strings.Split(c.TuneIDs, "\n")
	for i, t := range tuneIds {
		tuneIds[i] = strings.ToUpper(nonIDCharRegex.ReplaceAllString(t, ""))
	}
	statement = dialect.InsertStatement(`tune(id) VALUES` +
		formatList(tuneIds, ""))
	log.Println(statement)
	if len(tuneIds) > 0 {
		if res, err := tx.Exec(statement); err != nil {
			return res, err
		}
	}

	statement = dialect.InsertStatement(`chantey_tune(chantey_id, tune_id) VALUES` +
		formatList(tuneIds, c.ID))
	log.Println(statement)
	if len(tuneIds) > 0 {
		if res, err := tx.Exec(statement); err != nil {
			return res, err
		}
	}

	// Write chantey type mappings for later searching.
	types := strings.Split(c.Types, "\n")
	for i, t := range types {
		types[i] = strings.ToUpper(nonTypeCharRegex.ReplaceAllString(t, ""))
	}

	statement = dialect.InsertStatement("chantey_type_match(chantey_id, type_id) VALUES" +
		formatList(types, c.ID))
	log.Println(statement)
	if res, err := tx.Exec(statement); err != nil {
		return res, err
	}

	// Write theme mappings for later searching.
	themes := strings.Split(c.Themes, "\n")
	for i, t := range themes {
		themes[i] = nonThemeRegex.ReplaceAllString(t, "")
	}

	statement = dialect.InsertStatement("theme(id) VALUES" +
		formatList(themes, ""))
	log.Println(statement)
	if res, err := tx.Exec(statement); err != nil {
		return res, err
	}

	statement = dialect.InsertStatement("chantey_theme(chantey_id, theme_id) VALUES" +
		formatList(themes, c.ID))
	log.Println(statement)
	if res, err := tx.Exec(statement); err != nil {
		return res, err
	}

	return nil, nil
}

// SetCollectionID updates both the collectionID and the chantey ID.
func (c *Chantey) SetCollectionID(collectionID string) {
	c.CollectionID = collectionID
	c.ID = chanteyID(c.Title, c.CollectionID, c.Version, c.CollectionLocation)
}

func formatList(data []string, id string) string {
	if len(data) == 0 {
		return ""
	}
	b := strings.Builder{}
	for i, dat := range data {
		if i > 0 {
			b.WriteString(",")
		}
		b.WriteString("\n('")
		if len(id) > 0 {
			b.WriteString(id + "', '")
		}
		b.WriteString(dat + "')")
	}
	return b.String()
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
func (c *ChanteyJSON) ToDBChantey(collectionID string) *Chantey {
	return &Chantey{
		ID:                 c.ID(collectionID),
		TuneIDs:            strings.Join(c.TuneIDs, "\n"),
		CollectionID:       collectionID,
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
func (c *ChanteyJSON) ID(collectionID string) string {
	return chanteyID(strings.Join(c.Title, "\n"), collectionID, c.Version, c.CollectionLocation)
}

func chanteyID(title string, collectionID string, version string, location int) string {
	b := strings.Builder{}
	b.WriteString(convertKeyString(strings.Split(title, "\n"), 8))
	if location >= 0 {
		b.WriteString(".")
		b.WriteString(strconv.Itoa(location))
	}
	if len(version) > 0 {
		b.WriteString(".")
		b.WriteString(version)
	}
	b.WriteString(".")
	b.WriteString(collectionID)
	return b.String()
}
