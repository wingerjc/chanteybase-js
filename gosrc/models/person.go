package models

import "strings"

func LoadPersonConfig(dialect *SqlDialect) *DatabaseModel {
	conf := ModelConfig{
		Create: `CREATE TABLE IF NOT EXISTS person(
		id $TEXT PRIMARY KEY,
		first_name $TEXT NOT NULL,
		last_name $TEXT NOT NULL
		);`,
		Constraints: "",
	}
	return NewDatabaseModel(dialect, conf)
}

type Person struct {
	ID        string `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:last_name`
}

type PersonJson struct {
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
}

func (p *PersonJson) ToDBPerson() *Person {
	return &Person{
		ID:        p.ID(),
		FirstName: p.FirstName,
		LastName:  p.LastName,
	}
}

func (p *PersonJson) ID() string {
	b := strings.Builder{}
	b.WriteString(convertKeyString([]string{p.FirstName}, 7))
	b.WriteString(".")
	b.WriteString(convertKeyString([]string{p.LastName}, 7))
	return b.String()
}
