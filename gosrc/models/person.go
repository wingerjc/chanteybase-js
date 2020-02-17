package models

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

func LoadPersonConfig(dialect *SqlDialect) *DatabaseModel {
	conf := ModelConfig{
		Create: `CREATE TABLE IF NOT EXISTS person(
		id $TEXT PRIMARY KEY,
		group_name $TEXT NOT NULL,
		first_name $TEXT NOT NULL,
		last_name $TEXT NOT NULL,
		clarifier $TEXT NOT NULL,
		note $TEXT NOT NULL
		);`,
		Constraints: "",
	}
	return NewDatabaseModel(dialect, conf)
}

type Person struct {
	ID        string `db:"id" json:"id"`
	GroupName string `db:"group_name" json:"group-name"`
	FirstName string `db:"first_name" json:"first-name"`
	LastName  string `db:"last_name" json:"last-name"`
	Clarifier string `db:"clarifier" json:"clarifier"`
	Note      string `db:"note" json:"note"`
}

func (p *Person) Write(tx *sql.Tx, dialect SqlDialect) (sql.Result, error) {
	statement := dialect.replaceInsertPrefix + `INTO
	person (id, group_name, first_name, last_name, clarifier, note)
	VALUES ($1, $2, $3, $4, $5, $6);`
	fmt.Println(p)
	return tx.Exec(
		statement,
		p.ID,
		p.GroupName,
		p.FirstName,
		p.LastName,
		p.Clarifier,
		p.Note,
	)
}

func WritePeople(db *sqlx.DB, people []*Person, dialect SqlDialect) error {
	tx, err := db.BeginTx(context.TODO(), &sql.TxOptions{ReadOnly: false})
	if err != nil {
		return err
	}
	for _, p := range people {
		if _, err := p.Write(tx, dialect); err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

type PersonJson struct {
	GroupName string `json:"group-name"`
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
	Clarifier string `json:"clarifier"`
	Note      string `json:"note"`
}

func (p *PersonJson) ToDBPerson() *Person {
	return &Person{
		ID:        p.ID(),
		GroupName: p.GroupName,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Clarifier: p.Clarifier,
		Note:      p.Note,
	}
}

func (p *PersonJson) ID() string {
	b := strings.Builder{}
	fmt.Printf("group name *%s*\n", p.GroupName)
	if len(p.GroupName) != 0 {
		b.WriteString(convertKeyString([]string{p.GroupName}, 12))
	} else {
		fmt.Printf("Last name -- %s\n", convertKeyString([]string{p.LastName}, 7))
		b.WriteString(convertKeyString([]string{p.LastName}, 7))
		b.WriteString(".")
		b.WriteString(convertKeyString([]string{p.FirstName}, 7))
	}
	if len(p.Clarifier) > 0 {
		b.WriteString(".")
		b.WriteString(convertKeyString([]string{p.Clarifier}, 5))
	}
	fmt.Printf("person id -- %s\n", b.String())
	return b.String()
}
