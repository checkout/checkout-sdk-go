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
