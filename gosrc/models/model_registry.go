package models

func GetModelDefinitions(dialect *SqlDialect) []*DatabaseModel {
	return []*DatabaseModel{
		LoadChanteyConfig(dialect),
		LoadCollectionConfig(dialect),
		LoadPersonConfig(dialect),
	}
}
