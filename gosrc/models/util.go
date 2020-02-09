package models

import (
	"regexp"
	"strings"
	"database/sql"
)

var NON_ALPHA_REGEX, _ = regexp.Compile("[[^:alnum:]]+")

func convertKeyString(base []string, maxKeyLen int) string {
	title := strings.ToUpper(strings.Join(base, ""))
	title = NON_ALPHA_REGEX.ReplaceAllString(title, "")
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