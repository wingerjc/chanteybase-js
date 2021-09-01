package models

type URLParamJSON struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type PathSpecJSON struct {
	Name           string         `json:"name"`
	Description    string         `json:"description"`
	URL            string         `json:"url-path"`
	ReqPathParams  []URLParamJSON `json:"required-path-parameters"`
	OptPathParams  []URLParamJSON `json:"optional-path-parameters"`
	ReqQueryParams []URLParamJSON `json:"required-query-parameters"`
	OptQueryParams []URLParamJSON `json:"optional-query-parameters"`
}
