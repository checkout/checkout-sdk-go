package test

import (
	"github.com/checkout/checkout-sdk-go/forward"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func getForwardRequestIdSource() forward.ForwardRequest {
	IdSource := forward.NewIdSource()
	IdSource.Id = "src_v5rgkf3gdtpuzjqesyxmyodnya"

	request := forward.ForwardRequest{
		Source: IdSource,
		DestinationRequest: &forward.DestinationRequest{
			Url:    "https://example.com/payments",
			Method: forward.PostMT,
			Headers: &forward.Headers{
				Encrypted: "<JWE encrypted JSON object with string values>",
				Raw: map[string]string{
					"Idempotency-Key": "xe4fad12367dfgrds",
					"Content-Type":    "application/json",
				},
			},
			Body: `{\"amount\": 1000, \"currency\": \"USD\", \"reference\": \"some_reference\", \"source\": {\"type\": \"card\", \"number\": \"{{card_number}}\", \"expiry_month\": \"{{card_expiry_month}}\", \"expiry_year\": \"{{card_expiry_year_yyyy}}\", \"name\": \"Ali Farid\"}, \"payment_type\": \"Regular\", \"authorization_type\": \"Final\", \"capture\": true, \"processing_channel_id\": \"pc_xxxxxxxxxxx\", \"risk\": {\"enabled\": false}, \"merchant_initiated\": true}`,
		},
		Reference:           "ORD-5023-4E89",
		ProcessingChannelId: "pc_azsiyswl7bwe2ynjzujy7lcjca",
		NetworkToken: &forward.NetworkToken{
			Enabled:           true,
			RequestCryptogram: false,
		},
	}
	return request
}

func TestForwardAnApiRequest(t *testing.T) {
	t.Skip("unavailable")
	cases := []struct {
		name    string
		request forward.ForwardRequest
		checker func(*forward.ForwardAnApiResponse, error)
	}{
		{
			name:    "when request is correct then should forward an api request",
			request: getForwardRequestIdSource(),
			checker: func(response *forward.ForwardAnApiResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
			},
		},
	}

	client := OAuthApi().Forward

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.ForwardAnApiRequest(tc.request))
		})
	}
}

func TestGetForwardRequest(t *testing.T) {
	t.Skip("unavailable")
	cases := []struct {
		name      string
		forwardId string
		checker   func(*forward.GetForwardResponse, error)
	}{
		{
			name:      "when forward request exists then should return forward request",
			forwardId: "fwd_01HK153X00VZ1K15Z3HYC0QGPN",
			checker: func(response *forward.GetForwardResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, "fwd_01HK153X00VZ1K15Z3HYC0QGPN", response.RequestId)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotEmpty(t, response.DestinationResponse)
				assert.Equal(t, 201, response.DestinationResponse.Status)
				assert.NotEmpty(t, response.DestinationResponse.Headers)
			},
		},
	}

	client := OAuthApi().Forward

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetForwardRequest(tc.forwardId))
		})
	}
}
