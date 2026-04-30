package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v2/errors"
	applicants "github.com/checkout/checkout-sdk-go/v2/identities/applicants"
)

// # tests

func TestCreateApplicant(t *testing.T) {
	t.Skip("Avoid creating identity resources all the time")
	cases := []struct {
		name    string
		request applicants.CreateApplicantRequest
		checker func(*applicants.ApplicantResponse, error)
	}{
		{
			name:    "when request is valid then should return 201",
			request: createApplicantRequest(),
			checker: func(response *applicants.ApplicantResponse, err error) {
				assert.Nil(t, err)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
				validateApplicantResponse(t, response)
			},
		},
	}

	client := buildIdentitiesApi().Applicants

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.CreateApplicant(tc.request))
		})
	}
}

func TestGetApplicant(t *testing.T) {
	t.Skip("Avoid creating identity resources all the time")
	cases := []struct {
		name        string
		applicantId string
		checker     func(*applicants.ApplicantResponse, error)
	}{
		{
			name:        "when applicant exists then should return 200",
			applicantId: applicantResponse.Id,
			checker: func(response *applicants.ApplicantResponse, err error) {
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, applicantResponse.Id, response.Id)
				validateApplicantResponse(t, response)
			},
		},
		{
			name:        "when applicant not found then should return error",
			applicantId: "aplt_not_found",
			checker: func(response *applicants.ApplicantResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := buildIdentitiesApi().Applicants

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetApplicant(tc.applicantId))
		})
	}
}

func TestUpdateApplicant(t *testing.T) {
	t.Skip("Avoid creating identity resources all the time")
	cases := []struct {
		name        string
		applicantId string
		request     applicants.UpdateApplicantRequest
		checker     func(*applicants.ApplicantResponse, error)
	}{
		{
			name:        "when applicant exists then should return 200",
			applicantId: applicantResponse.Id,
			request:     updateApplicantRequest(),
			checker: func(response *applicants.ApplicantResponse, err error) {
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, applicantResponse.Id, response.Id)
				validateApplicantResponse(t, response)
			},
		},
		{
			name:        "when applicant not found then should return error",
			applicantId: "aplt_not_found",
			request:     updateApplicantRequest(),
			checker: func(response *applicants.ApplicantResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := buildIdentitiesApi().Applicants

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.UpdateApplicant(tc.applicantId, tc.request))
		})
	}
}

func TestAnonymizeApplicant(t *testing.T) {
	t.Skip("Avoid creating identity resources all the time")
	cases := []struct {
		name        string
		applicantId string
		checker     func(*applicants.ApplicantResponse, error)
	}{
		{
			name:        "when applicant exists then should return 200",
			applicantId: applicantResponse.Id,
			checker: func(response *applicants.ApplicantResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotEmpty(t, response.Id)
			},
		},
		{
			name:        "when applicant not found then should return error",
			applicantId: "aplt_not_found",
			checker: func(response *applicants.ApplicantResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := buildIdentitiesApi().Applicants

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.AnonymizeApplicant(tc.applicantId))
		})
	}
}

// # common methods

func createApplicantRequest() applicants.CreateApplicantRequest {
	return applicants.CreateApplicantRequest{
		ExternalApplicantId:   "ext_applicant_test_123",
		Email:                 GenerateRandomEmail(),
		ExternalApplicantName: "John Doe",
	}
}

func updateApplicantRequest() applicants.UpdateApplicantRequest {
	return applicants.UpdateApplicantRequest{
		Email:                 GenerateRandomEmail(),
		ExternalApplicantName: "John Updated Doe",
	}
}

func validateApplicantResponse(t *testing.T, response *applicants.ApplicantResponse) {
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.Id)
	assert.NotEmpty(t, response.Email)
	assert.NotNil(t, response.CreatedOn)
	assert.NotNil(t, response.ModifiedOn)
}
