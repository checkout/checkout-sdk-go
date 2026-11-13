package common

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/checkout/checkout-sdk-go/v3/errors"
)

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

// BuildPath joins the segments with "/" and prefixes the result with "/". Segments are used as
// given: constants may contain "/", and every caller-supplied value must go through
// EncodePathSegment first.
func BuildPath(params ...string) string {
	var path string
	for _, s := range params {
		path += "/" + s
	}

	return path
}

// EncodePathSegment percent-encodes a caller-supplied value so that it is sent as exactly one
// path segment. An empty value, "." or "..", or a value containing "/", "\\" or an ASCII control
// character returns a CheckoutArgumentError. Separators are refused rather than encoded because
// the API decodes "%2F" and "%5C" back into separators.
func EncodePathSegment(segment string) (string, error) {
	if segment == "" || segment == "." || segment == ".." || hasForbiddenCharacter(segment) {
		return "", errors.CheckoutArgumentError("invalid path segment")
	}

	return url.PathEscape(segment), nil
}

func hasForbiddenCharacter(segment string) bool {
	for i := 0; i < len(segment); i++ {
		c := segment[i]
		if c == '/' || c == '\\' || c <= 0x1F || c == 0x7F {
			return true
		}
	}

	return false
}

// BuildQueryPath appends the encoded query values to path.
//
// An empty filter returns the path unchanged rather than appending a bare "?". That matters for
// the endpoints whose query filter is optional: before the 2026-09-02 pass the list-attempts
// endpoints used BuildPath and produced ".../attempts", and a caller passing a zero filter would
// otherwise have started getting ".../attempts?" instead. A trailing question mark with no query
// carries no information, so no caller loses anything.
func BuildQueryPath(path string, queryValues interface{}) (string, error) {
	values, err := query.Values(queryValues)
	if err != nil {
		return "", err
	}

	encoded := values.Encode()
	if encoded == "" {
		return path, nil
	}

	return fmt.Sprintf("%s?%s", path, encoded), nil
}

func EscapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

func BuildDefaultClient() *http.Client {
	return &http.Client{Timeout: time.Duration(10) * time.Second}
}
