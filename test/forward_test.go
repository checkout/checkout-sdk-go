package test

import (
	"net/http"
	"testing"

	"github.com/checkout/checkout-sdk-go/v2/common"
	"github.com/checkout/checkout-sdk-go/v2/forward"
	"github.com/stretchr/testify/assert"
)

// tests

func TestForwardAnApiRequest(t *testing.T) {
	t.Skip("unavailable")
	cases := []struct {
		name    string
		request forward.ForwardRequest
		checker func(*forward.ForwardAnApiResponse, error)
	}{
		{
			name:    "when request is correct then should forward an api request",
			request: buildForwardRequest(),
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

func TestCreateSecret(t *testing.T) {
	t.Skip("This test requires forward secrets scopes and valid credentials")
	cases := []struct {
		name    string
		request forward.CreateSecretRequest
		checker func(*forward.SingleSecretResponse, error)
	}{
		{
			name:    "when request is valid then create secret",
			request: buildCreateSecretRequest(),
			checker: func(response *forward.SingleSecretResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
				assert.NotEmpty(t, response.Name)
				assert.Equal(t, 1, response.Version)
				assert.NotZero(t, response.CreatedAt)
				assert.NotZero(t, response.UpdatedAt)
			},
		},
	}

	client := OAuthApi().Forward

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.CreateSecret(tc.request))
		})
	}
}

func TestListSecrets(t *testing.T) {
	t.Skip("This test requires forward secrets scopes and valid credentials")
	cases := []struct {
		name    string
		checker func(*forward.ListSecretsResponse, error)
	}{
		{
			name: "when secrets exist then return list of secrets",
			checker: func(response *forward.ListSecretsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.Data)
			},
		},
	}

	client := OAuthApi().Forward

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.ListSecrets())
		})
	}
}

func TestUpdateSecret(t *testing.T) {
	t.Skip("This test requires forward secrets scopes and valid credentials")
	cases := []struct {
		name       string
		secretName string
		request    forward.UpdateSecretRequest
		checker    func(*forward.SingleSecretResponse, error)
	}{
		{
			name:       "when request is valid then update secret",
			secretName: "secret_name",
			request:    buildUpdateSecretRequest(),
			checker: func(response *forward.SingleSecretResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotEmpty(t, response.Name)
				assert.NotZero(t, response.UpdatedAt)
			},
		},
	}

	client := OAuthApi().Forward

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.UpdateSecret(tc.secretName, tc.request))
		})
	}
}

func TestDeleteSecret(t *testing.T) {
	t.Skip("This test requires forward secrets scopes and valid credentials")
	cases := []struct {
		name       string
		secretName string
		checker    func(*common.MetadataResponse, error)
	}{
		{
			name:       "when secret exists then delete secret",
			secretName: "secret_name",
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusNoContent, response.HttpMetadata.StatusCode)
			},
		},
	}

	client := OAuthApi().Forward

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.DeleteSecret(tc.secretName))
		})
	}
}

// common methods

func buildForwardRequest() forward.ForwardRequest {
	idSource := forward.NewIdSource()
	idSource.Id = "src_v5rgkf3gdtpuzjqesyxmyodnya"

	dlocalSignature := forward.NewDlocalSignature()
	dlocalSignature.DlocalParameters = forward.DlocalParameters{
		SecretKey: "9f439fe1a9f96e67b047d3c1a28c33a2e",
	}

	return forward.ForwardRequest{
		Source: idSource,
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
			Body:      `{\"amount\": 1000, \"currency\": \"USD\", \"reference\": \"some_reference\", \"source\": {\"type\": \"card\", \"number\": \"{{card_number}}\", \"expiry_month\": \"{{card_expiry_month}}\", \"expiry_year\": \"{{card_expiry_year_yyyy}}\", \"name\": \"Ali Farid\"}, \"payment_type\": \"Regular\", \"authorization_type\": \"Final\", \"capture\": true, \"processing_channel_id\": \"pc_xxxxxxxxxxx\", \"risk\": {\"enabled\": false}, \"merchant_initiated\": true}`,
			Signature: dlocalSignature,
		},
		Reference:           "ORD-5023-4E89",
		ProcessingChannelId: "pc_azsiyswl7bwe2ynjzujy7lcjca",
		NetworkToken: &forward.NetworkToken{
			Enabled:           true,
			RequestCryptogram: false,
		},
	}
}

func buildCreateSecretRequest() forward.CreateSecretRequest {
	return forward.CreateSecretRequest{
		Name:  "test_secret_" + GenerateRandomString(8, "abcdefghijklmnopqrstuvwxyz0123456789"),
		Value: "plaintext_value",
	}
}

func buildUpdateSecretRequest() forward.UpdateSecretRequest {
	return forward.UpdateSecretRequest{
		Value: "updated_value",
	}
}
