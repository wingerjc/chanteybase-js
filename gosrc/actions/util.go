package actions

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/jmoiron/sqlx"
)

func parseParams(urlString, prefix string) []string {
	p := strings.Split(strings.TrimPrefix(urlString, prefix), "/")
	result := make([]string, 0, len(p))
	for _, param := range p {
		out, _ := url.PathUnescape(param)
		result = append(result, out)
	}
	return result
}

func writeResp(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	fmt.Fprintf(w, msg)
}

// PathSpec is a speck for http actions that includes, common name, URL, and handling function.
type PathSpec struct {
	Name string
	URL  string
	Fn   func(db *sqlx.DB) func(w http.ResponseWriter, req *http.Request)
}
