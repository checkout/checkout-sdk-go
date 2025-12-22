package common

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type APIShortDate time.Time

func (t *APIShortDate) UnmarshalJSON(data []byte) error {
	// Remove quotes from JSON string
	str := strings.Trim(string(data), "\"")
	if str == "null" || str == "" {
		return nil
	}

	// Try multiple date formats that your API might return
	formats := []string{
		"2006-01-02", // Specific yyyy-MM-dd Date only
	}

	for _, format := range formats {
		if parsed, err := time.Parse(format, str); err == nil {
			*t = APIShortDate(parsed)
			return nil
		}
	}

	return fmt.Errorf("unable to parse time: %s, APIShortDate only accepts yyyy-MM-dd format", str)
}

func (t APIShortDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Format("2006-01-02"))
}
