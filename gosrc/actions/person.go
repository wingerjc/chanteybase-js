package actions

import (
	"encoding/json"
	"net/http"

	"github.com/jmoiron/sqlx"
	"local.dev/models"
)

const (
	GetPersonByIDURL = "/person/"
)

func GetPersonByID(db *sqlx.DB) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		params := parseParams(req.URL.EscapedPath(), GetPersonByIDURL)
		if len(params) == 0 || len(params[0]) == 0 {
			writeResp(w, 400, "Missing ID param /person/:id")
			return
		}
		data, err := models.GetPersonByID(db, params[0])
		if err != nil {
			writeResp(w, 500, "Couldn't fetch person data.")
			return
		}
		var js []byte
		js, err = json.Marshal(data)
		if err != nil {
			writeResp(w, 500, "Could not format person JSON")
			return
		}
		w.Write(js)
	}
}
