package actions

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"local.dev/models"
)

const (
	// GetChanteyByIDURL is the URL prefix for searching chanteys by ID
	GetChanteyByIDURL = "/chantey/"
	// GetChanteyByCollectionIDURL is the URL for searching chanteys by collection ID
	GetChanteyByCollectionIDURL = "/chantey-collection/"
)

// ChanteyByID is an HTTPFunc for searching chanteys by ID.
func ChanteyByID(db *sqlx.DB) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		params := parseParams(req.URL.EscapedPath(), GetChanteyByIDURL)
		if len(params) == 0 || len(params[0]) == 0 {
			writeResp(w, 400, "Missing param for chantey search: id")
			return
		}
		data, err := models.ChanteyByID(db, params[0])
		if err != nil {
			writeResp(w, 500, "Could not fetch chanteys from data layer.")
			log.Printf("Error fetching chantey from DB: %s", err.Error())
			return
		}
		var js []byte
		js, err = json.Marshal(data)
		if err != nil {
			writeResp(w, 500, "Couldn't format chanteys in JSON.")
			return
		}
		w.Write(js)
	}
}

// ChanteyByCollectionID is an HTTPFunc for searching chanteys by collection ID.
func ChanteyByCollectionID(db *sqlx.DB) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		params := parseParams(req.URL.EscapedPath(), GetChanteyByCollectionIDURL)
		if len(params) == 0 || len(params[0]) == 0 {
			writeResp(w, 400, "Missing param for chantey: Collection ID ")
			return
		}
		data, err := models.ChanteyByCollectionID(db, params[0])
		if err != nil {
			writeResp(w, 500, "Couldn't fetch chanteys from data layer.")
			log.Printf("Error fetching chantey by collection ID: %s", err.Error())
			return
		}
		var js []byte
		js, err = json.Marshal(data)
		if err != nil {
			writeResp(w, 500, "Couldn't format chanteys in JSON.")
			return
		}
		w.Write(js)
	}
}
