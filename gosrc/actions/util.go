package actions

import (
	"fmt"
	"net/http"
	"strings"
)

func parseParams(url, prefix string) []string {
	return strings.Split(strings.TrimPrefix(url, prefix), "/")
}

func writeResp(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	fmt.Fprintf(w, msg)
}
