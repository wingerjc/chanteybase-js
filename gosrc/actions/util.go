package actions

import (
	"net/http"

	"local.dev/models"
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
	w.Write(models.CreateErrorJson(msg))
	//fmt.Fprintf(w, msg)
}
