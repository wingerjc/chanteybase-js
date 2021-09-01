package actions

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/jmoiron/sqlx"
	"local.dev/models"
)

type URLParam struct {
	Name        string
	Description string
}

// PathSpec is a speck for http actions that includes, common name, URL, and handling function.
type PathSpec struct {
	Name           string
	Description    string
	URL            string
	Fn             func(url string, spec PathSpec, db *sqlx.DB) func(w http.ResponseWriter, req *http.Request)
	ReqPathParams  []URLParam
	OptPathParams  []URLParam
	ReqQueryParams []URLParam
	OptQueryParams []URLParam
}

// GetPathParams returns a list of all path param names for parsing paths.
func (s *PathSpec) GetPathParams() []string {
	reqLen := len(s.ReqPathParams)
	totalLen := reqLen + len(s.OptPathParams)
	res := make([]string, totalLen, totalLen)
	for i, req := range s.ReqPathParams {
		res[i] = req.Name
	}
	for i, opt := range s.OptPathParams {
		res[reqLen+i] = opt.Name
	}
	return res
}

// DisplayPath creates a string to explain how an API URL should be invoked.
func (s *PathSpec) DisplayPath() string {
	items := make([]string, 0, len(s.ReqPathParams))
	for _, p := range s.ReqPathParams {
		items = append(items, p.Name)
	}
	reqPath := strings.Join(items, "/")
	items = make([]string, 0, len(s.OptPathParams))
	for _, p := range s.OptPathParams {
		items = append(items, fmt.Sprintf("%s?", p.Name))
	}
	optPath := strings.Join(items, "/")
	// TODO: fix formatting for display path,and add query string data.
	return fmt.Sprintf("%s%s/%s", s.URL, reqPath, optPath)
}

func (s *PathSpec) ToJsonModel() models.PathSpecJSON {
	paramList := func(l []URLParam) []models.URLParamJSON {
		result := make([]models.URLParamJSON, len(l), len(l))
		for i, p := range l {
			result[i] = models.URLParamJSON{
				Name:        p.Name,
				Description: p.Description,
			}
		}
		return result
	}
	return models.PathSpecJSON{
		Name:           s.Name,
		Description:    s.Description,
		URL:            s.DisplayPath(),
		ReqPathParams:  paramList(s.ReqPathParams),
		OptPathParams:  paramList(s.OptPathParams),
		ReqQueryParams: paramList(s.ReqQueryParams),
		OptQueryParams: paramList(s.OptQueryParams),
	}
}

// URLParams contains path based parameters, the query fragment, and any query parameters passed.
type URLParams struct {
	PathParams    map[string]string
	PathFragment  string
	QueryParams   map[string][]string
	MissingParams map[string]bool
}

// NewUrlParams parses an action URL given the prefix and path parameters.
func NewURLParams(urlData *url.URL, prefix string, pathParams []string) (URLParams, error) {
	subPath := strings.TrimPrefix(urlData.Path, prefix)
	log.Printf("%s : %s : %s ", urlData.Path, prefix, subPath)
	subPathParams := strings.Split(subPath, "/")
	pathValues := make(map[string]string)
	missingValues := make(map[string]bool)
	for i, paramName := range pathParams {
		if len(subPath) > 0 && i < len(subPathParams) {
			pathValues[paramName] = subPathParams[i]
		} else {
			missingValues[paramName] = true
		}
	}

	return URLParams{
		PathParams:    pathValues,
		PathFragment:  urlData.Fragment,
		QueryParams:   urlData.Query(),
		MissingParams: missingValues,
	}, nil
}

func assertRequiredParams(spec PathSpec, params URLParams, w http.ResponseWriter) bool {
	if len(params.MissingParams) > 0 {
		missingList := make([]string, 0, len(params.MissingParams))
		for k, _ := range params.MissingParams {
			missingList = append(missingList, k)
		}
		writeResp(w, 400, fmt.Sprintf("Missing path params %s in api path '%s'", strings.Join(missingList, ", "), spec.DisplayPath()))
		return false
	}
	return true
}
