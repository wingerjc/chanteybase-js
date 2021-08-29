package models

import (
	"strings"
)

// DatabaseModel represent the config for a single model that can be written to the DB.
type DatabaseModel struct {
	configFile       string
	createScript     string
	constraintScript string
	insertScript     string
}

// ModelConfig Is the pre-dialect creation scripts for a given model.
type ModelConfig struct {
	Create      string
	Constraints string
	Insert      string
}

// SQLDialect has dialect replacements for create/insert/update scripts..
type SQLDialect struct {
	replaceInsertStatement string
	replacements           map[string]string
}

// InsertStatement creates an insert statement from the list of values.
func (dialect *SQLDialect) InsertStatement(valueStatement string) string {
	return strings.Replace(dialect.replaceInsertStatement, "$VALUES", valueStatement, 1)
}

// NewDatabaseModel converts a database config to the passed dialect.
func NewDatabaseModel(dialect *SQLDialect, config ModelConfig) *DatabaseModel {
	return &DatabaseModel{
		createScript:     processScript(config.Create, dialect.replacements),
		constraintScript: processScript(config.Constraints, dialect.replacements),
		insertScript:     processScript(config.Insert, dialect.replacements),
	}
}

// CreateScript returns the final creation script for the model tables.
func (model *DatabaseModel) CreateScript() string {
	return model.createScript
}

// ConstraintScript returns the final table creation scripts.
func (model *DatabaseModel) ConstraintScript() string {
	return model.constraintScript
}

// InsertScript returns the final single insertion script.
func (model *DatabaseModel) InsertScript() string {
	return model.insertScript
}

func processScript(script string, dialect map[string]string) string {
	result := script
	for k, v := range dialect {
		result = strings.ReplaceAll(result, k, v)
	}
	return result
}

// Sqlite3Dialect has the dialect definition for sqlite3.
func Sqlite3Dialect() *SQLDialect {
	return &SQLDialect{
		replaceInsertStatement: "INSERT OR REPLACE INTO $VALUES;",
		replacements: map[string]string{
			"$TEXT": "TEXT",
			"$INT":  "INTEGER",
		},
	}
}

// Postgres12Dialect is a SQLDialect for postgres 12
func Postgres12Dialect() *SQLDialect {
	return &SQLDialect{
		replaceInsertStatement: "INSERT INTO $VALUES ON CONFLICT DO NOTHING;",
		replacements: map[string]string{
			"$TEXT": "TEXT",
			"$INT":  "INTEGER",
		},
	}
}
