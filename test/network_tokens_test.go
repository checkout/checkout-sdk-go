package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v2/common"
	"github.com/checkout/checkout-sdk-go/v2/networktokens"
)

// tests

func TestProvisionNetworkToken(t *testing.T) {
	t.Skip("use on demand")
	cases := []struct {
		name    string
		request networktokens.ProvisionNetworkTokenRequest
		checker func(*networktokens.NetworkTokenResponse, error)
	}{
		{
			name:    "when source is a valid instrument then provision network token",
			request: buildNetworkTokenProvisionRequest(),
			checker: func(response *networktokens.NetworkTokenResponse, err error) {
				assert.Nil(t, err)
				assertNetworkTokenResponse(t, response)
			},
		},
	}

	client := OAuthApi().NetworkTokens

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.ProvisionNetworkToken(tc.request))
		})
	}
}

func TestGetNetworkToken(t *testing.T) {
	t.Skip("use on demand")
	cases := []struct {
		name           string
		networkTokenId string
		checker        func(*networktokens.NetworkTokenResponse, error)
	}{
		{
			name:           "when token id is valid then return network token",
			networkTokenId: "nt_xgu3isllqfyu7ktpk5z2yxbwna",
			checker: func(response *networktokens.NetworkTokenResponse, err error) {
				assert.Nil(t, err)
				assertNetworkTokenResponse(t, response)
			},
		},
		{
			name:           "when token id is invalid then return error",
			networkTokenId: "nt_invalid",
			checker: func(response *networktokens.NetworkTokenResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
			},
		},
	}

	client := OAuthApi().NetworkTokens

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetNetworkToken(tc.networkTokenId))
		})
	}
}

func TestRequestCryptogram(t *testing.T) {
	t.Skip("use on demand")
	cases := []struct {
		name           string
		networkTokenId string
		request        networktokens.RequestCryptogramRequest
		checker        func(*networktokens.RequestCryptogramResponse, error)
	}{
		{
			name:           "when token is active then return cryptogram",
			networkTokenId: "nt_xgu3isllqfyu7ktpk5z2yxbwna",
			request:        buildCryptogramRequest(),
			checker: func(response *networktokens.RequestCryptogramResponse, err error) {
				assert.Nil(t, err)
				assertCryptogramResponse(t, response)
			},
		},
	}

	client := OAuthApi().NetworkTokens

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.RequestCryptogram(tc.networkTokenId, tc.request))
		})
	}
}

func TestDeleteNetworkToken(t *testing.T) {
	t.Skip("use on demand")
	cases := []struct {
		name           string
		networkTokenId string
		request        networktokens.DeleteNetworkTokenRequest
		checker        func(*common.MetadataResponse, error)
	}{
		{
			name:           "when token exists then delete network token",
			networkTokenId: "nt_xgu3isllqfyu7ktpk5z2yxbwna",
			request:        buildNetworkTokenDeleteRequest(),
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusNoContent, response.HttpMetadata.StatusCode)
			},
		},
	}

	client := OAuthApi().NetworkTokens

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.DeleteNetworkToken(tc.networkTokenId, tc.request))
		})
	}
}

// common methods

func buildNetworkTokenProvisionRequest() networktokens.ProvisionNetworkTokenRequest {
	src := networktokens.NewIdSource()
	src.Id = "src_wmlfc3zyhqzehihu7giusaaawu"
	return networktokens.ProvisionNetworkTokenRequest{Source: src}
}

func buildCryptogramRequest() networktokens.RequestCryptogramRequest {
	return networktokens.RequestCryptogramRequest{TransactionType: networktokens.Ecom}
}

func buildNetworkTokenDeleteRequest() networktokens.DeleteNetworkTokenRequest {
	return networktokens.DeleteNetworkTokenRequest{
		InitiatedBy: networktokens.TokenRequestor,
		Reason:      networktokens.Other,
	}
}

func assertNetworkTokenResponse(t *testing.T, response *networktokens.NetworkTokenResponse) {
	assert.NotNil(t, response)
	assert.NotNil(t, response.Card)
	assert.NotNil(t, response.NetworkToken)
	assert.NotEmpty(t, response.NetworkToken.Id)
	assert.NotEmpty(t, response.NetworkToken.State)
}

func assertCryptogramResponse(t *testing.T, response *networktokens.RequestCryptogramResponse) {
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
	assert.NotEmpty(t, response.Cryptogram)
}
