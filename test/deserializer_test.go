package test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestCardsUnmarshallJson(t *testing.T) {
	physical, _ := ioutil.ReadFile("resources/card_details_physical_response.json")
	virtual, _ := ioutil.ReadFile("resources/card_details_virtual_response.json")

	cases := []struct {
		name    string
		json    []byte
		checker func(*bytes.Buffer, error)
	}{
		{
			name: "when deserializing card_details_physical_response type must be physical",
			json: physical,
			checker: func(serialized *bytes.Buffer, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, serialized)

				var deserialized map[string]interface{}
				unmErr := json.Unmarshal(serialized.Bytes(), &deserialized)

				assert.Nil(t, unmErr)
				assert.Equal(t, "physical", deserialized["type"])
				assert.Contains(t, deserialized, "type")
				assert.NotContains(t, deserialized, "is_single_use")
			},
		},
		{
			name: "when deserializing card_details_virtual_response type must be virtual",
			json: virtual,
			checker: func(serialized *bytes.Buffer, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, serialized)

				var deserialized map[string]interface{}
				unmErr := json.Unmarshal(serialized.Bytes(), &deserialized)

				assert.Nil(t, unmErr)
				assert.Equal(t, "virtual", deserialized["type"])
				assert.Contains(t, deserialized, "type")
				assert.Contains(t, deserialized, "is_single_use")
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(bytes.NewBuffer(tc.json), nil)
		})
	}
}

func TestControlsUnmarshallJson(t *testing.T) {
	mccLimit, _ := ioutil.ReadFile("resources/card_control_mcc_limit_response.json")
	velocityLimit, _ := ioutil.ReadFile("resources/card_control_velocity_limit_response.json")

	cases := []struct {
		name    string
		json    []byte
		checker func(*bytes.Buffer, error)
	}{
		{
			name: "when deserializing card_control_mcc_limit_response control_type must be mcc_limit",
			json: mccLimit,
			checker: func(serialized *bytes.Buffer, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, serialized)

				var deserialized map[string]interface{}
				unmErr := json.Unmarshal(serialized.Bytes(), &deserialized)

				assert.Nil(t, unmErr)
				assert.Equal(t, "mcc_limit", deserialized["control_type"])
				assert.Contains(t, deserialized, "mcc_limit")
				assert.NotContains(t, deserialized, "velocity_limit")
			},
		},
		{
			name: "when deserializing card_control_velocity_limit_response control_type must be velocity_limit",
			json: velocityLimit,
			checker: func(serialized *bytes.Buffer, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, serialized)

				var deserialized map[string]interface{}
				unmErr := json.Unmarshal(serialized.Bytes(), &deserialized)

				assert.Nil(t, unmErr)
				assert.Equal(t, "velocity_limit", deserialized["control_type"])
				assert.Contains(t, deserialized, "velocity_limit")
				assert.NotContains(t, deserialized, "mcc_limit")
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(bytes.NewBuffer(tc.json), nil)
		})
	}
}

func TestPaymentContextUnmarshallJson(t *testing.T) {
	paypalResponse, _ := ioutil.ReadFile("resources/payment_context_paypal_details_response.json")

	cases := []struct {
		name    string
		json    []byte
		checker func(*bytes.Buffer, error)
	}{
		{
			name: "when deserializing payment_context_paypal_details_response type must be PayPal",
			json: paypalResponse,
			checker: func(serialized *bytes.Buffer, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, serialized)

				var deserialized map[string]interface{}
				unmErr := json.Unmarshal(serialized.Bytes(), &deserialized)

				assert.Nil(t, unmErr)
				assert.NotNil(t, deserialized["payment_request"])
				paymentRequest := deserialized["payment_request"].(map[string]interface{})
				source := paymentRequest["source"].(map[string]interface{})

				assert.Equal(t, "paypal", source["type"])
				assert.Contains(t, source, "type")
				assert.Contains(t, source, "account_holder")

				accountHolder := source["account_holder"].(map[string]interface{})
				assert.Contains(t, accountHolder, "full_name")
				assert.Equal(t, "Andrey Young", accountHolder["full_name"])
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(bytes.NewBuffer(tc.json), nil)
		})
	}
}
