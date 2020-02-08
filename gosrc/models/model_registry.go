package models

func GetModelSchemas() []string {
	return []string{
		chanteySchema,
	}
}

func GetModelConstraints() []string {
	return []string{
		chanteyConstraints,
	}
}
