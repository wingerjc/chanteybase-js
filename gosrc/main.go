package main

import (
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

	sqlDB, err := sqlx.Open("sqlite3", "./chanteys.db")
	defer sqlDB.Close()
	if err != nil {
		log.Fatalf("Couldn't open DB: %s", err.Error())
	}

	for _, s := range models.GetModelSchemas() {
		log.Print(s)
		sqlDB.Exec(s)
	}
}
