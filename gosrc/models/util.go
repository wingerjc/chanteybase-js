package models

import (
	"database/sql"
	"regexp"
	"strings"
)

var nonAlphaRegex, _ = regexp.Compile("[[^:alnum:]]+")
var whitespaceRegex, _ = regexp.Compile("[[:space:]]+")
var nonIDCharRegex, _ = regexp.Compile(`[^A-Za-z0-9.]+`)

func convertKeyString(base []string, maxKeyLen int) string {
	title := strings.ToUpper(strings.Join(base, ""))
	title = whitespaceRegex.ReplaceAllString(nonAlphaRegex.ReplaceAllString(title, ""), "")
	keyLen := maxKeyLen
	if len(title) < keyLen {
		keyLen = len(title)
	}
	return title[:keyLen]
}

func toNullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: true}
}

func toNullInt(i int) sql.NullInt64 {
	return sql.NullInt64{Int64: int64(i), Valid: true}
}

func emptyToUnknown(s string) string {
	if len(s) == 0 {
		return "UNKNOWN"
	}
	return s
}

func emptyToNA(s string) string {
	if len(s) == 0 {
		return "N/A"
	}
	return s
}

func dbSearchString(s string) string {
	return "%" + s + "%"
}
