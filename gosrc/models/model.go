package models

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

const (
	DIALECT_SQLITE = iota
)

type SqlDialect int

type DatabaseModel struct {
	configFile       string
	createScript     string
	constraintScript string
}

type ModelConfig struct {
	Create      []string `json:"create"`
	Constraints []string `json:"constraints"`
}

func NewDatabaseModel(fileName string, dialect SqlDialect) *DatabaseModel {
	dialects := make(map[SqlDialect]func() map[string]string)
	dialects[DIALECT_SQLITE] = SQLITE_TEMPLATE

	file, _ := ioutil.ReadFile(fileName)
	config := ModelConfig{}
	_ = json.Unmarshal([]byte(file), &config)

	currentDialect := dialects[dialect]()
	return &DatabaseModel{
		createScript:     processScript(config.Create, currentDialect),
		constraintScript: processScript(config.Constraints, currentDialect),
	}
}

func (model *DatabaseModel) CreateScript() string {
	return model.createScript
}

func (model *DatabaseModel) ConstraintScript() string {
	return model.constraintScript
}

func processScript(script []string, dialect map[string]string) string {
	result := strings.Join(script, "\n")
	for k, v := range dialect {
		result = strings.ReplaceAll(result, k, v)
	}
	return result
}

func SQLITE_TEMPLATE() map[string]string {
	return map[string]string{
		"$TEXT": "TEXT",
	}
}
