package models

func LoadPersonConfig(dialect *SqlDialect) *DatabaseModel {
	conf := ModelConfig{
		Create: `
		`,
		Constraints: "",
	}
	return NewDatabaseModel(dialect, conf)
}

type Person struct {
}

type PersonJson struct {
}
