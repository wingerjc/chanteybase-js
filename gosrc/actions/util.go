package actions

import (
	"fmt"
	"net/http"
)

func parseParams(urlString, prefix string) []string {
	return []string{}
	// return NewURLParams(urlString, prefix, []string{})
	// p := strings.Split(, "/")
	// result := make([]string, 0, len(p))
	// for _, param := range p {
	// 	out, _ := url.PathUnescape(param)
	// 	result = append(result, out)
	// }
	// return result
}

func writeResp(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	fmt.Fprintf(w, msg)
}
