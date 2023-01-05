package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/apm/ideal"
)

/* TODO fix 404 not-found
func TestGetInfo(t *testing.T) {
	cases := []struct {
		name    string
		checker func(*ideal.IdealInfo, error)
	}{
		{
			name: "when auth is correct then return ideal info",
			checker: func(response *ideal.IdealInfo, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.IdealInfoLinks)
				assert.NotNil(t, response.IdealInfoLinks.Self)
				assert.NotNil(t, response.IdealInfoLinks.Curies)
				assert.NotNil(t, response.IdealInfoLinks.Issuers)
			},
		},
	}

	client := PreviousApi().Ideal

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetInfo())
		})
	}
}
*/

func TestGetIssuers(t *testing.T) {
	cases := []struct {
		name    string
		checker func(*ideal.IssuerResponse, error)
	}{
		{
			name: "when auth is correct then return issuers info",
			checker: func(response *ideal.IssuerResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.Countries)
				assert.NotNil(t, response.Links)
			},
		},
	}

	client := PreviousApi().Ideal

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetIssuers())
		})
	}
}
