package common

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// APIShortDate is a date without a time component, for the fields the Checkout.com
// specification declares as "type": "string", "format": "date".
//
// It exists because encoding/json renders time.Time as RFC 3339 and offers no per-field
// override, so a time.Time on a date-only field puts a full timestamp on the wire. Some
// providers reject that outright -- Tamara returns a gateway error for a timestamp on
// processing.accommodation_data.check_in_date.
//
// Construct one by converting a time.Time; the time component is discarded on serialization:
//
//	d := common.APIShortDate(time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC))
//	request.CheckInDate = &d
type APIShortDate time.Time

// shortDateFormats are the layouts APIShortDate accepts when deserializing, in priority order.
//
// A date-time is deliberately NOT accepted. Every field typed APIShortDate is declared
// "format": "date" in the specification, so a timestamp on one of them is a server-side
// contract break; silently truncating it to a date would hide that. Java rejects date-times on
// LocalDate for the same reason, and test/apishortdate_test.go pins the rejection.
//
// time.Parse is strict about both layouts -- they are mutually exclusive, out-of-range dates and
// trailing characters are rejected -- so no length guard is needed around the second format.
var shortDateFormats = []string{
	"2006-01-02", // yyyy-MM-dd, the format the specification declares
	"20060102",   // yyyyMMdd, the compact form some endpoints return
}

// UnmarshalJSON parses a date-only JSON string into an APIShortDate. A JSON null or an empty
// string leaves the value at its zero time rather than erroring, so an absent optional date is
// not a parse failure.
func (t *APIShortDate) UnmarshalJSON(data []byte) error {
	// Remove quotes from JSON string
	str := strings.Trim(string(data), "\"")
	if str == "null" || str == "" {
		return nil
	}

	for _, format := range shortDateFormats {
		if parsed, err := time.Parse(format, str); err == nil {
			*t = APIShortDate(parsed)
			return nil
		}
	}

	return fmt.Errorf(
		"unable to parse time: %s, APIShortDate only accepts yyyy-MM-dd or yyyyMMdd format", str)
}

// MarshalJSON renders the date as yyyy-MM-dd, discarding any time component.
func (t APIShortDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Format("2006-01-02"))
}
