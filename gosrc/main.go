package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	//"database/sql"

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
	s := &server{}
	http.Handle("/", s)
	//log.Fatal(http.ListenAndServe(":8080",nil))

	config := loadConfig()
	sqlDB, err := sqlx.Open("sqlite3", "./chanteys.db")
	defer sqlDB.Close()
	if err != nil {
		log.Fatalf("Couldn't open DB: %s", err.Error())
	}

	log.Printf("config dir %s", config.ConfigDirectory)
	modelDefs := models.GetModelDefinitions(models.SQLITE3_DIALECT())
	for _, s := range modelDefs {
		log.Print(s.CreateScript())
		sqlDB.Exec(s.CreateScript())
	}
	for _, s := range modelDefs {
		log.Print(s.ConstraintScript())
		sqlDB.Exec(s.ConstraintScript())
	}
}

type Config struct {
	SqlDialect      string `json:"sql-dialect"`
	ConfigDirectory string `json:"config-dir"`
}

// TODO: pass config filename.
func loadConfig() *Config {
	file, _ := ioutil.ReadFile("./config/config.json")
	config := Config{}
	_ = json.Unmarshal([]byte(file), &config)

	return &config
}
