package models

import (
	"path/filepath"
)

func LoadCollectorConfig(path string, dialect SqlDialect) *DatabaseModel {
	return NewDatabaseModel(filepath.Join(path, "collector.json"), dialect)
}

type Collector struct {
}

type CollectorJson struct {
}
