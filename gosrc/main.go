package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	"local.dev/actions"
	"local.dev/models"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{"message":"hello world"}`))
}

func main() {
	runServer := flag.Bool("server", false, "Set true to run as an http server.")
	flag.Parse()

	config := loadConfig()
	sqlDB, err := sqlx.Open("sqlite3", config.DBFile)
	defer sqlDB.Close()
	if err != nil {
		log.Fatalf("Couldn't open DB: %s", err.Error())
	}

	dialect := models.Sqlite3Dialect()
	log.Printf("config dir %s", config.ConfigDirectory)
	modelDefs := models.GetModelDefinitions(dialect)
	for _, s := range modelDefs {
		log.Print(s.CreateScript())
		sqlDB.Exec(s.CreateScript())
	}
	for _, s := range modelDefs {
		log.Print(s.ConstraintScript())
		sqlDB.Exec(s.ConstraintScript())
	}
	for _, s := range modelDefs {
		log.Print(s.InsertScript())
		sqlDB.Exec(s.InsertScript())
	}

	data := models.GetDataFromJSON(config.DataDirectory, nil)
	err = models.WritePeople(sqlDB, data.People, *dialect)
	if err != nil {
		log.Fatalf("Couldn't insert people in DB: %s", err.Error())
	}

	err = models.WriteCollections(sqlDB, data.Collections, *dialect)
	if err != nil {
		log.Fatalf("Couldn't insert collections in DB: %s", err.Error())
	}

	err = models.WriteChanteys(sqlDB, data.Chanteys, *dialect)
	if err != nil {
		log.Fatalf("Couldn't insert chanteys in DB: %s", err.Error())
	}

	if *runServer {
		serverMain(sqlDB)
	}

}

func serverMain(db *sqlx.DB) {
	s := &server{}
	http.Handle("/", s)
	// person searches
	http.HandleFunc(actions.GetPersonByIDURL, actions.GetPersonByID(db))
	http.HandleFunc(actions.GetPersonIDsURL, actions.GetPersonIDs(db))
	// collection searches
	http.HandleFunc(actions.GetCollectionByIDURL, actions.CollectionByID(db))
	http.HandleFunc(actions.GetCollectionByTitleURL, actions.CollectionByTitle(db))
	// chantey searches
	http.HandleFunc(actions.GetChanteyByIDURL, actions.ChanteyByID(db))
	http.HandleFunc(actions.GetChanteyByCollectionIDURL, actions.ChanteyByCollectionID(db))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Config is the basic application configuration that can be read from JSON.
type Config struct {
	SQLDialect      string `json:"sql-dialect"`
	ConfigDirectory string `json:"config-dir"`
	DataDirectory   string `json:"data-dir"`
	DBFile          string `json:"db-file"`
}

// TODO: pass config filename.
func loadConfig() *Config {
	file, _ := ioutil.ReadFile("./config/config.json")
	config := Config{}
	_ = json.Unmarshal([]byte(file), &config)

	return &config
}
