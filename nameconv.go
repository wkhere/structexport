package structexport

import (
	"strings"
	"unicode"
)

// nameconv converts FieldName to FIELD_NAME.
// In other words, pascal case to upper sname case.
func nameconv(s string) string {
	if s == "" {
		return ""
	}

	buf := new(strings.Builder)
	lastUpper := true

	for _, c := range s {
		upper := unicode.IsUpper(c)
		switch {
		case upper && !lastUpper:
			buf.WriteByte('_')
			buf.WriteRune(c)
		case upper:
			buf.WriteRune(c)
		default:
			buf.WriteRune(unicode.ToUpper(c))
		}
		lastUpper = upper
	}
	return buf.String()
}
