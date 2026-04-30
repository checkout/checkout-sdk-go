package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v2/errors"
	amlscreening "github.com/checkout/checkout-sdk-go/v2/identities/amlscreening"
	"github.com/checkout/checkout-sdk-go/v2/identities"
)

// # tests

func TestCreateAmlScreening(t *testing.T) {
	t.Skip("Avoid creating identity resources all the time")
	cases := []struct {
		name    string
		request amlscreening.CreateAmlScreeningRequest
		checker func(*amlscreening.AmlScreeningResponse, error)
	}{
		{
			name:    "when request is valid then should return 201",
			request: createAmlScreeningRequest(),
			checker: func(response *amlscreening.AmlScreeningResponse, err error) {
				assert.Nil(t, err)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
				validateAmlScreeningResponse(t, response)
			},
		},
		{
			name:    "when applicant not found then should return error",
			request: amlscreening.CreateAmlScreeningRequest{
				ApplicantId:      "aplt_not_found",
				SearchParameters: identities.SearchParameters{ConfigurationIdentifier: "config_test_123"},
				Monitored:        Bool(true),
			},
			checker: func(response *amlscreening.AmlScreeningResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := buildIdentitiesApi().AmlScreening

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.CreateAmlScreening(tc.request))
		})
	}
}

func TestGetAmlScreening(t *testing.T) {
	t.Skip("Avoid creating identity resources all the time")
	cases := []struct {
		name        string
		screeningId string
		checker     func(*amlscreening.AmlScreeningResponse, error)
	}{
		{
			name:        "when screening exists then should return 200",
			screeningId: "scr_existing_test_id",
			checker: func(response *amlscreening.AmlScreeningResponse, err error) {
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				validateAmlScreeningResponse(t, response)
			},
		},
		{
			name:        "when screening not found then should return error",
			screeningId: "scr_not_found",
			checker: func(response *amlscreening.AmlScreeningResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := buildIdentitiesApi().AmlScreening

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetAmlScreening(tc.screeningId))
		})
	}
}

// # common methods

func createAmlScreeningRequest() amlscreening.CreateAmlScreeningRequest {
	return amlscreening.CreateAmlScreeningRequest{
		ApplicantId: applicantResponse.Id,
		SearchParameters: identities.SearchParameters{
			ConfigurationIdentifier: "config_test_123",
		},
		Monitored: Bool(true),
	}
}

func validateAmlScreeningResponse(t *testing.T, response *amlscreening.AmlScreeningResponse) {
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.Id)
	assert.NotEmpty(t, response.ApplicantId)
	assert.NotNil(t, response.Status)
	assert.NotNil(t, response.SearchParameters)
	assert.NotNil(t, response.CreatedOn)
	assert.NotNil(t, response.ModifiedOn)
}
