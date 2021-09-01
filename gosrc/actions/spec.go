package actions

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"local.dev/models"
)

const (
	// GetAllSpecURL is the URL for listing all api specs.
	GetAllSpecURL = "/all-spec/"
	// GetDocSpecURL is the URL for all the spec describing endpoints.
	GetDocSpecURL = "/doc-spec/"
	// GetPersonSpecURL is the URL for all the person endpoints.
	GetPersonSpecURL = "/person-spec/"
	// GetEndpointSpecURL is the URL for getting specific endpoint data.
	GetEndpointSpecURL = "/spec/"
)

var (
	// SpecActions is a list of all actions that mainly fetch spec data.
	SpecActions = []PathSpec{
		{Name: "GetAllSpec", URL: GetAllSpecURL, Fn: GetAllSpecFn,
			Description: "Return a list of all Endpoint names and display URLs",
		},
		{Name: "GetEndpointsSpec", URL: GetEndpointSpecURL, Fn: GetEndpointSpecFn,
			Description: "Return data on a given endpoint.",
			ReqPathParams: []URLParam{
				{Name: ":id", Description: "Exact string ID of the endpoint spec to fetch."},
			},
		},
		{Name: "GetDocSpec", URL: GetDocSpecURL, Fn: GetDocSpecFn,
			Description: "Return a list of all api spec endpoints.",
		},
		{Name: "GetPersonSpecURL", URL: GetPersonSpecURL, Fn: GetPersonSpecFn,
			Description: "Return a list of all Person fetching endpoints.",
		},
	}
	// Used to avoid circular init dependency.
	internalSpecActions []PathSpec
	allSpecs            []PathSpec
	allSpecsJSON        map[string]models.PathSpecJSON
)

func InitSpecEndpoints(specs []PathSpec) {
	internalSpecActions = SpecActions
	allSpecs = specs
	allSpecsJSON = make(map[string]models.PathSpecJSON)
	for _, s := range allSpecs {
		allSpecsJSON[s.Name] = s.ToJsonModel()
	}
}

// GetFn is an HTTPFunc for
func GetAllSpecFn(url string, spec PathSpec, db *sqlx.DB) func(w http.ResponseWriter, req *http.Request) {
	return sectionSpecFn(allSpecs, url, spec, db)
}

// GetFn is an HTTPFunc for
func GetEndpointSpecFn(url string, spec PathSpec, db *sqlx.DB) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		params, err := NewURLParams(req.URL, url, spec.GetPathParams())
		if !assertRequiredParams(spec, params, w) {
			return
		}
		data, ok := allSpecsJSON[params.PathParams[":id"]]
		if !ok {
			writeResp(w, 500, fmt.Sprintf("Couldn't fetch spec named '%s'", params.PathParams[":id"]))
			return
		}
		var js []byte
		js, err = json.Marshal(data)
		if err != nil {
			writeResp(w, 500, "Could not format path spec JSON")
			return
		}
		w.Write(js)
	}
}

// GetDocSpecFn is an HTTPFunc for listing doc/spec endpoints
func GetDocSpecFn(url string, spec PathSpec, db *sqlx.DB) func(w http.ResponseWriter, req *http.Request) {
	return sectionSpecFn(internalSpecActions, url, spec, db)
}

// GetPersonSpecFn is an HTTPFunc for listing all person based endpoints.
func GetPersonSpecFn(url string, spec PathSpec, db *sqlx.DB) func(w http.ResponseWriter, req *http.Request) {
	return sectionSpecFn(PersonActions, url, spec, db)
}

// sectionSpecFn returns an HTTPFunc for a list of specs.
func sectionSpecFn(specList []PathSpec, url string, spec PathSpec, db *sqlx.DB) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		data := make([][]string, 0, len(specList))
		for _, s := range specList {
			data = append(data, []string{s.Name, s.Description, s.DisplayPath()})
		}
		var js []byte
		js, err := json.Marshal(data)
		if err != nil {
			writeResp(w, 500, "Could not format spec list JSON")
			return
		}
		w.Write(js)
	}

}
