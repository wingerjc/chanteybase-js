package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path"
)

// GetModelDefinitions fetches all the definitions for models.
func GetModelDefinitions(dialect *SQLDialect) []*DatabaseModel {
	return []*DatabaseModel{
		LoadConstantsConfig(dialect),
		LoadChanteyConfig(dialect),
		LoadCollectionConfig(dialect),
		LoadPersonConfig(dialect),
	}
}

// ProgressTracker allows progress tracking for loading.
type ProgressTracker struct {
	CurrentFile chan string
	Progress    chan int
}

// LoadedModelData The data loaded from JSON.
type LoadedModelData struct {
	People      []*Person
	Collections []*Collection
	Chanteys    []*Chantey
}

// GetDataFromJSON loads all the data from JSON into writable models.
func GetDataFromJSON(dataPath string, progress *ProgressTracker) *LoadedModelData {
	filePath := path.Join(dataPath, "person.json")
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Error loading people file %s | %s", filePath, err.Error)
	}
	peopleJSON := make([]PersonJSON, 0)
	err = json.Unmarshal(data, &peopleJSON)
	if err != nil {
		log.Printf("Error parsing people file %s", err.Error)
	}
	people := make([]*Person, 0, len(peopleJSON))
	for _, p := range peopleJSON {
		people = append(people, p.ToDBPerson())
	}

	filePath = path.Join(dataPath, "collection.json")
	data, err = ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Error loading collections file %s | %s", filePath, err.Error())
	}
	collectionsJSON := make([]CollectionJSON, 0)
	err = json.Unmarshal(data, &collectionsJSON)
	if err != nil {
		log.Printf("Error parsing collections file %s", err.Error())
	}
	collections := make([]*Collection, 0, len(collectionsJSON))
	for _, c := range collectionsJSON {
		collections = append(collections, c.ToDBCollection())
	}

	filePath = path.Join(dataPath, "chantey.json")
	data, err = ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Error loading chantey file %s | %s", filePath, err.Error)
	}
	chanteysJSON := make([]ChanteyJSON, 0)
	err = json.Unmarshal(data, &chanteysJSON)
	if err != nil {
		log.Printf("Error parsing chantey file %s", err.Error)
	}
	chanteys := make([]*Chantey, 0, len(chanteysJSON))
	for _, p := range chanteysJSON {
		chanteys = append(chanteys, p.ToDBChantey())
	}

	return &LoadedModelData{
		People:      people,
		Collections: collections,
		Chanteys:    chanteys,
	}
}

// LoadConstantsConfig returns the config for constant tables.
func LoadConstantsConfig(dialect *SQLDialect) *DatabaseModel {
	conf := ModelConfig{
		Create: `CREATE TABLE IF NOT EXISTS location_type(type $TEXT PRIMARY KEY);
			CREATE TABLE IF NOT EXISTS chantey_type(type $TEXT PRIMARY KEY);`,
		Insert: `INSERT INTO TABLE location_type(type) VALUES
		  ('PAGE'),
		  ('SECONDS'),
		  ('TRACK');
		INSERT INTO TABLE chantey_type(type) VALUES
		  ('SHORT_DRAG'),
		  ('BUNTING'),
		  ('HALYARD'),
		  ('FORECASTLE'),
		  ('CAPSTAN');`,
	}

	return NewDatabaseModel(dialect, conf)
}
