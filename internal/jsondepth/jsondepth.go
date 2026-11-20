// Package jsondepth checks how deeply a JSON document nests arrays and objects before it is
// decoded. It is internal to the SDK and not part of its public API.
package jsondepth

// MaxDepth is the deepest nesting of JSON arrays and objects the SDK decodes, the same limit
// encoding/json applies from Go 1.15.
const MaxDepth = 10000

// WithinLimit reports whether the arrays and objects of body, outside strings, are nested at most
// MaxDepth deep. It does not check that body is valid JSON.
func WithinLimit(body []byte) bool {
	depth := 0
	inString := false
	escaped := false
	for _, c := range body {
		switch {
		case escaped:
			escaped = false
		case inString:
			if c == '\\' {
				escaped = true
			} else if c == '"' {
				inString = false
			}
		case c == '"':
			inString = true
		case c == '[' || c == '{':
			depth++
			if depth > MaxDepth {
				return false
			}
		case c == ']' || c == '}':
			depth--
		}
	}
	return true
}
