package actions

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"local.dev/models"
)

const (
	// GetCollectionByIDURL is the URL for searching collections by ID
	GetCollectionByIDURL = "/collection/"
	// GetCollectionByTitleURL is the URL for searching collections by title
	GetCollectionByTitleURL = "/collection-title/"
)

// CollectionByID is an HTTPFun for searching collections by ID
func CollectionByID(db *sqlx.DB) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		params := parseParams(req.URL.EscapedPath(), GetCollectionByIDURL)
		if len(params) == 0 {
			writeResp(w, 400, "Missing param for collection id search/fetch")
			return
		}
		data, err := models.CollectionByID(db, params[0])
		if err != nil {
			writeResp(w, 500, "Could not fetch collection data error")
			log.Printf("Error fetching data: %s", err.Error())
			return
		}
		var js []byte
		js, err = json.Marshal(data)
		if err != nil {
			writeResp(w, 500, "Couldn't format collections to JSON.")
			return
		}
		w.Write(js)
	}
}

// CollectionByTitle is an HTTPFunc for searching collections by title.
func CollectionByTitle(db *sqlx.DB) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		params := parseParams(req.URL.EscapedPath(), GetCollectionByTitleURL)
		if len(params) == 0 {
			writeResp(w, 400, "Missing param for collection id search/fetch")
			return
		}
		data, err := models.CollectionByTitle(db, params[0])
		if err != nil {
			writeResp(w, 500, "Could not fetch collection data error")
			log.Printf("Error fetching data: %s", err.Error())
			return
		}
		var js []byte
		js, err = json.Marshal(data)
		if err != nil {
			writeResp(w, 500, "Couldn't format collections to JSON.")
			return
		}
		w.Write(js)
	}
}
