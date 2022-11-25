package common

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func BuildPath(params ...string) string {
	var path string
	for _, s := range params {
		path += "/" + s
	}

	return path
}

func BuildQueryPath(path string, queryValues interface{}) (string, error) {
	values, err := query.Values(queryValues)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("/%s?%s", path, values.Encode()), nil
}

func EscapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

func BuildDefaultClient() *http.Client {
	return &http.Client{Timeout: time.Duration(10) * time.Second}
}
