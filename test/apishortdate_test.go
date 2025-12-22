package test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/stretchr/testify/assert"
)

func TestAPIShortDateUnmarshalling(t *testing.T) {
	cases := []struct {
		name         string
		jsonInput    string
		expectedDate time.Time
	}{
		{
			name:         "YYYY-MM-DD format (day > month)",
			jsonInput:    `"2023-03-15"`,
			expectedDate: time.Date(2023, 3, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			name:         "YYYY-MM-DD format (day < month)",
			jsonInput:    `"2023-12-05"`,
			expectedDate: time.Date(2023, 12, 5, 0, 0, 0, 0, time.UTC),
		},
		{
			name:         "YYYY-MM-DD leap year",
			jsonInput:    `"2024-02-29"`,
			expectedDate: time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC),
		},
		{
			name:         "YYYY-MM-DD single digits",
			jsonInput:    `"2023-01-09"`,
			expectedDate: time.Date(2023, 1, 9, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var apiDate common.APIShortDate

			err := json.Unmarshal([]byte(tc.jsonInput), &apiDate)
			assert.Nil(t, err, "Unmarshalling should not fail")

			actualTime := time.Time(apiDate)
			assert.Equal(t, tc.expectedDate.Year(), actualTime.Year(), "Year should match")
			assert.Equal(t, tc.expectedDate.Month(), actualTime.Month(), "Month should match")
			assert.Equal(t, tc.expectedDate.Day(), actualTime.Day(), "Day should match")
		})
	}
}

func TestAPIShortDateMarshalling(t *testing.T) {
	cases := []struct {
		name         string
		inputDate    time.Time
		expectedJSON string
	}{
		{
			name:         "Day > Month (15th of March)",
			inputDate:    time.Date(2023, 3, 15, 10, 30, 45, 0, time.UTC),
			expectedJSON: `"2023-03-15"`,
		},
		{
			name:         "Day < Month (5th of December)",
			inputDate:    time.Date(2023, 12, 5, 14, 20, 30, 0, time.UTC),
			expectedJSON: `"2023-12-05"`,
		},
		{
			name:         "Single digit month and day",
			inputDate:    time.Date(2023, 1, 9, 0, 0, 0, 0, time.UTC),
			expectedJSON: `"2023-01-09"`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiDate := common.APIShortDate(tc.inputDate)

			jsonBytes, err := json.Marshal(apiDate)
			assert.Nil(t, err, "Marshalling should not fail")

			actualJSON := string(jsonBytes)
			assert.Equal(t, tc.expectedJSON, actualJSON, "JSON output should match expected format")
		})
	}
}

func TestAPIShortDateFormatConfusion(t *testing.T) {
	cases := []struct {
		name          string
		jsonInput     string
		expectedDay   int
		expectedMonth time.Month
	}{
		{
			name:          "Day 15 Month 03",
			jsonInput:     `"2023-03-15"`,
			expectedDay:   15,
			expectedMonth: time.March,
		},
		{
			name:          "Day 05 Month 12",
			jsonInput:     `"2023-12-05"`,
			expectedDay:   5,
			expectedMonth: time.December,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var apiDate common.APIShortDate

			err := json.Unmarshal([]byte(tc.jsonInput), &apiDate)
			assert.Nil(t, err, "Unmarshalling should not fail")

			actualTime := time.Time(apiDate)
			assert.Equal(t, tc.expectedDay, actualTime.Day(), "Day should be correctly parsed")
			assert.Equal(t, tc.expectedMonth, actualTime.Month(), "Month should be correctly parsed")
		})
	}
}

func TestAPIShortDateInvalidFormats(t *testing.T) {
	cases := []struct {
		name      string
		jsonInput string
		errorMsg  string
	}{
		{
			name:      "ISO 8601 with timezone should fail",
			jsonInput: `"2023-06-20T14:30:45Z"`,
			errorMsg:  "should reject ISO format with time",
		},
		{
			name:      "ISO 8601 with milliseconds should fail",
			jsonInput: `"2023-09-12T09:15:30.123Z"`,
			errorMsg:  "should reject ISO format with milliseconds",
		},
		{
			name:      "Date without timezone should fail",
			jsonInput: `"2023-11-25T18:45:00"`,
			errorMsg:  "should reject datetime without timezone",
		},
		{
			name:      "Date with space should fail",
			jsonInput: `"2023-07-08 12:00:00"`,
			errorMsg:  "should reject date with space and time",
		},
		{
			name:      "Invalid date format should fail",
			jsonInput: `"not-a-date"`,
			errorMsg:  "should reject invalid date string",
		},
		{
			name:      "Wrong date format MM/DD/YYYY should fail",
			jsonInput: `"03/15/2023"`,
			errorMsg:  "should reject US date format",
		},
		{
			name:      "Wrong date format DD/MM/YYYY should fail",
			jsonInput: `"15/03/2023"`,
			errorMsg:  "should reject European date format",
		},
		{
			name:      "Invalid date values should fail",
			jsonInput: `"2023-13-45"`,
			errorMsg:  "should reject invalid month/day values",
		},
		{
			name:      "Partial date should fail",
			jsonInput: `"2023-03"`,
			errorMsg:  "should reject incomplete date",
		},
		{
			name:      "Date with extra characters should fail",
			jsonInput: `"2023-03-15extra"`,
			errorMsg:  "should reject date with extra characters",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var apiDate common.APIShortDate

			err := json.Unmarshal([]byte(tc.jsonInput), &apiDate)
			assert.NotNil(t, err, tc.errorMsg)
			assert.Contains(t, err.Error(), "APIShortDate only accepts", "Error should mention format restriction")
		})
	}
}

func TestAPIShortDateRoundTrip(t *testing.T) {
	cases := []struct {
		name      string
		inputJSON string
	}{
		{
			name:      "Day > Month case (March 25th)",
			inputJSON: `"2023-03-25"`,
		},
		{
			name:      "Day < Month case (December 8th)",
			inputJSON: `"2023-12-08"`,
		},
		{
			name:      "Leap year February 29th",
			inputJSON: `"2024-02-29"`,
		},
		{
			name:      "Year boundary December 31st",
			inputJSON: `"2023-12-31"`,
		},
		{
			name:      "Year boundary January 1st",
			inputJSON: `"2024-01-01"`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Step 1: Unmarshal input JSON
			var apiDate common.APIShortDate
			err := json.Unmarshal([]byte(tc.inputJSON), &apiDate)
			assert.Nil(t, err, "Initial unmarshalling should not fail")

			originalTime := time.Time(apiDate)

			// Step 2: Marshal back to JSON
			jsonBytes, err := json.Marshal(apiDate)
			assert.Nil(t, err, "Marshalling should not fail")

			// Step 3: Verify output format is yyyy-MM-dd
			outputJSON := string(jsonBytes)
			assert.Contains(t, outputJSON, "-", "Output should contain dashes (yyyy-MM-dd format)")
			assert.Equal(t, tc.inputJSON, outputJSON, "Round-trip should preserve exact format")

			// Step 4: Unmarshal the output back to verify round-trip integrity
			var roundTripDate common.APIShortDate
			err = json.Unmarshal(jsonBytes, &roundTripDate)
			assert.Nil(t, err, "Round-trip unmarshalling should work")

			// Step 5: Verify dates represent the same day
			roundTripTime := time.Time(roundTripDate)
			assert.Equal(t, originalTime.Year(), roundTripTime.Year(), "Year should be preserved in round-trip")
			assert.Equal(t, originalTime.Month(), roundTripTime.Month(), "Month should be preserved in round-trip")
			assert.Equal(t, originalTime.Day(), roundTripTime.Day(), "Day should be preserved in round-trip")

			// Step 6: Verify expected output format
			expectedOutput := originalTime.Format("2006-01-02")
			assert.Equal(t, `"`+expectedOutput+`"`, outputJSON, "Output should match yyyy-MM-dd format")
		})
	}
}
