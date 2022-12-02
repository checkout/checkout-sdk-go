package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/metadata"
	"github.com/checkout/checkout-sdk-go/metadata/sources"
)

func TestCardMetadataSources(t *testing.T) {
	cases := []struct {
		name    string
		request metadata.CardMetadataRequest
		checker func(*metadata.CardMetadataResponse, error)
	}{
		{
			name: "when card metadata requested by card number exist return card metadata info",
			request: metadata.CardMetadataRequest{
				Source: sources.NewRequestCardSource("4242424242424242"),
				Format: metadata.Basic,
			},
			checker: func(response *metadata.CardMetadataResponse, err error) {
				cardMetadataResponseAssertions(t, response, err)
			},
		},
		{
			name: "when card metadata requested by bin number exist return card metadata info",
			request: metadata.CardMetadataRequest{
				Source: sources.NewRequestBinSource("42424242"),
				Format: metadata.Basic,
			},
			checker: func(response *metadata.CardMetadataResponse, err error) {
				cardMetadataResponseAssertions(t, response, err)
			},
		},
		{
			name: "when card metadata requested by token exist return card metadata info",
			request: metadata.CardMetadataRequest{
				Source: sources.NewRequestTokenSource(RequestCardToken(t).Token),
				Format: metadata.Basic,
			},
			checker: func(response *metadata.CardMetadataResponse, err error) {
				cardMetadataResponseAssertions(t, response, err)
			},
		},
	}

	client := OAuthApi().Metadata

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			response, err := client.RequestCardMetadata(tc.request)

			tc.checker(response, err)
		})
	}
}

func cardMetadataResponseAssertions(t *testing.T, response *metadata.CardMetadataResponse, err error) {
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
	assert.NotNil(t, response.IssuerCountryName)
	assert.NotNil(t, response.ProductId)
	assert.NotNil(t, response.ProductType)
	assert.NotNil(t, response.Scheme)
	assert.Equal(t, "42424242", response.Bin)
	assert.IsType(t, metadata.CardMetadataResponse{}.CardType, response.CardType)
	assert.IsType(t, metadata.CardMetadataResponse{}.CardCategory, response.CardCategory)
	assert.IsType(t, metadata.CardMetadataResponse{}.IssuerCountry, response.IssuerCountry)
}
