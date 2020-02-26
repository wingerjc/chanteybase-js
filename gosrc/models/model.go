package models

import (
	"strings"
)

type DatabaseModel struct {
	configFile       string
	createScript     string
	constraintScript string
	insertScript     string
}

type ModelConfig struct {
	Create      string
	Constraints string
	Insert      string
}

type SqlDialect struct {
	replaceInsertPrefix string
	replacements        map[string]string
}

func NewDatabaseModel(dialect *SqlDialect, config ModelConfig) *DatabaseModel {
	return &DatabaseModel{
		createScript:     processScript(config.Create, dialect.replacements),
		constraintScript: processScript(config.Constraints, dialect.replacements),
		insertScript:     processScript(config.Insert, dialect.replacements),
	}
}

func (model *DatabaseModel) CreateScript() string {
	return model.createScript
}

func (model *DatabaseModel) ConstraintScript() string {
	return model.constraintScript
}

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

func SQLITE3_DIALECT() *SqlDialect {
	return &SqlDialect{
		replaceInsertPrefix: "INSERT OR REPLACE ",
		replacements: map[string]string{
			"$TEXT": "TEXT",
			"$INT":  "INTEGER",
		},
	}
}
