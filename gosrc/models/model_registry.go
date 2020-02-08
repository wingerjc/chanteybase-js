package models

func GetModelDefinitions(path string, dialect SqlDialect) []*DatabaseModel {
	return []*DatabaseModel{
		LoadChanteyConfig(path, dialect),
	}
}
