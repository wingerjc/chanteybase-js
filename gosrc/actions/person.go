package actions

import (
	"encoding/json"
	"net/http"

	"github.com/jmoiron/sqlx"
	"local.dev/models"
)

const (
	// GetPersonByIDURL is the URL for searching people by ID
	GetPersonByIDURL = "/person/"
	// GetPersonIDsURL is the URL for getting all person IDs by partial match
	GetPersonIDsURL = "/person-ids/"
)

var (
	// PersonActions is a list of all actions that mainly fetch person objects.
	PersonActions = []PathSpec{
		{Name: "GetPersonByID", URL: GetPersonByIDURL, Fn: GetPersonByIDFn,
			ReqPathParams: []URLParam{
				{Name: ":id", Comment: "Exact ID of the person record to fetch"},
			},
		},
		{Name: "GetPersonIDs", URL: GetPersonIDsURL, Fn: GetPersonIDsFn,
			OptPathParams: []URLParam{
				{Name: ":target", Comment: "Plain text search target (contains)"},
			},
		},
	}
)

// GetPersonByIDFn is an HTTPFunc for searching all people by partial ID.
func GetPersonByIDFn(url string, spec PathSpec, db *sqlx.DB) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		params, err := NewURLParams(req.URL, url, spec.GetPathParams())
		if !AssertRequiredParams(spec, params, w) {
			return
		}
		data, err := models.GetPersonByID(db, params.PathParams[":id"])
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

// GetPersonIDsFn is an HTTPFunc for getting all person ID's by partial match.
func GetPersonIDsFn(url string, spec PathSpec, db *sqlx.DB) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		params, err := NewURLParams(req.URL, url, spec.GetPathParams())
		if err != nil {
			writeResp(w, 500, "Error Fetching person object IDs, could not parse path")
			return
		}
		searchString := ""
		if target, ok := params.PathParams[":target"]; ok {
			searchString = target
		}
		data, err := models.GetPersonIDs(db, searchString)
		if err != nil {
			writeResp(w, 500, "Error Fetching person object IDs, could not fetch data")
			return
		}
		var js []byte
		js, err = json.Marshal(data)
		if err != nil {
			writeResp(w, 500, "Could not convert ID list to JSON")
			return
		}
		w.Write(js)
	}
}
