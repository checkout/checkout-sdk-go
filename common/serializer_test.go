package common

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMarshal(t *testing.T) {
	cases := []struct {
		name    string
		request interface{}
		checker func(*bytes.Buffer, error)
	}{
		{
			name: "when serializing nil *time.Time field should not be included",
			request: struct {
				Time       *time.Time `json:"time,omitempty"`
				OtherField string     `json:"other_field,omitempty"`
			}{
				OtherField: "value",
			},
			checker: func(serialized *bytes.Buffer, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, serialized)

				var deserialized map[string]interface{}
				unmErr := json.Unmarshal(serialized.Bytes(), &deserialized)

				assert.Nil(t, unmErr)
				assert.Contains(t, deserialized, "other_field")
				assert.NotContains(t, deserialized, "time")
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(Marshal(tc.request))
		})
	}
}

func TestUnmarshalJSON(t *testing.T) {
	responseBody := []byte(`{"type":"example","control_type":"example_control"}`)
	headers := &Headers{Header: map[string][]string{"Content-Type": {"application/json"}}}
	metadata := &HttpMetadata{ResponseBody: responseBody, Headers: headers}

	var response TypeMapping
	err := Unmarshal(metadata, &response)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	assert.Equal(t, response, TypeMapping{Type: "example", ControlType: "example_control"})
}

func TestUnmarshalText(t *testing.T) {
	responseBody := []byte("name,age\nJohn,30")
	headers := &Headers{Header: map[string][]string{"Content-Type": {"text/csv"}}}
	metadata := &HttpMetadata{ResponseBody: responseBody, Headers: headers}

	type Response struct {
		Content string
	}

	var response Response
	err := Unmarshal(metadata, &response)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	assert.Equal(t, response.Content, "name,age\nJohn,30")
}

func TestUnmarshalPDF(t *testing.T) {
	responseBody := []byte("%PDF-1.4\n%...")
	headers := &Headers{Header: map[string][]string{"Content-Type": {"application/pdf"}}}
	metadata := &HttpMetadata{ResponseBody: responseBody, Headers: headers}

	type Response struct {
		Content []byte
	}

	var response Response
	err := Unmarshal(metadata, &response)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	assert.Equal(t, response.Content, responseBody)
}
