package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path"
)

func GetModelDefinitions(dialect *SqlDialect) []*DatabaseModel {
	return []*DatabaseModel{
		LoadChanteyConfig(dialect),
		LoadCollectionConfig(dialect),
		LoadPersonConfig(dialect),
	}
}

type ProgressTracker struct {
	CurrentFile chan string
	Progress    chan int
}

type LoadedModelData struct {
	People      []*Person
	Collections []*Collection
	Chanteys    []*Chantey
}

func GetDataFromJson(dataPath string, progress *ProgressTracker) *LoadedModelData {
	filePath := path.Join(dataPath, "person.json")
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Error loading people file %s | %s", filePath, err.Error)
	}
	peopleJson := make([]PersonJson, 0)
	err = json.Unmarshal(data, &peopleJson)
	if err != nil {
		log.Printf("Error parsing people file %s", err.Error)
	}
	people := make([]*Person, 0, len(peopleJson))
	for _, p := range peopleJson {
		people = append(people, p.ToDBPerson())
	}

	filePath = path.Join(dataPath, "collection.json")
	data, err = ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Error loading collections file %s | %s", filePath, err.Error)
	}
	collectionsJson := make([]CollectionJson, 0)
	err = json.Unmarshal(data, &collectionsJson)
	if err != nil {
		log.Printf("Error parsing collections file %s", err.Error)
	}
	collections := make([]*Collection, 0, len(collectionsJson))
	for _, c := range collectionsJson {
		collections = append(collections, c.ToDBCollection())
	}

	filePath = path.Join(dataPath, "chantey.json")
	data, err = ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Error loading chantey file %s | %s", filePath, err.Error)
	}
	chanteysJson := make([]ChanteyJson, 0)
	err = json.Unmarshal(data, &chanteysJson)
	if err != nil {
		log.Printf("Error parsing chantey file %s", err.Error)
	}
	chanteys := make([]*Chantey, 0, len(chanteysJson))
	for _, p := range chanteysJson {
		chanteys = append(chanteys, p.ToDBChantey())
	}

	return &LoadedModelData{
		People:      people,
		Collections: collections,
		Chanteys:    chanteys,
	}
}
