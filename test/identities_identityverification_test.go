package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v2/errors"
	"github.com/checkout/checkout-sdk-go/v2/identities"
	identityverification "github.com/checkout/checkout-sdk-go/v2/identities/identityverification"
)

// # tests

func TestCreateIdentityVerificationAndAttempt(t *testing.T) {
	t.Skip("Avoid creating identity resources all the time")
	cases := []struct {
		name    string
		request identityverification.CreateIdentityVerificationAndAttemptRequest
		checker func(*identityverification.IdentityVerificationAndAttemptResponse, error)
	}{
		{
			name:    "when request is valid then should return 201",
			request: createIdentityVerificationAndAttemptRequest(),
			checker: func(response *identityverification.IdentityVerificationAndAttemptResponse, err error) {
				assert.Nil(t, err)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
				assert.NotEmpty(t, response.Id)
				assert.NotEmpty(t, response.ApplicantId)
				assert.NotEmpty(t, response.RedirectUrl)
				assert.NotNil(t, response.Status)
			},
		},
		{
			name: "when applicant not found then should return error",
			request: identityverification.CreateIdentityVerificationAndAttemptRequest{
				ApplicantId: "aplt_not_found",
				RedirectUrl: SuccessUrl,
			},
			checker: func(response *identityverification.IdentityVerificationAndAttemptResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := buildIdentitiesApi().IdentityVerification

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.CreateIdentityVerificationAndAttempt(tc.request))
		})
	}
}

func TestCreateIdentityVerification(t *testing.T) {
	t.Skip("Avoid creating identity resources all the time")
	cases := []struct {
		name    string
		request identityverification.CreateIdentityVerificationRequest
		checker func(*identityverification.IdentityVerificationResponse, error)
	}{
		{
			name:    "when request is valid then should return 201",
			request: createIdentityVerificationRequest(),
			checker: func(response *identityverification.IdentityVerificationResponse, err error) {
				assert.Nil(t, err)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
				validateIdentityVerificationResponse(t, response)
			},
		},
		{
			name: "when applicant not found then should return error",
			request: identityverification.CreateIdentityVerificationRequest{
				ApplicantId: "aplt_not_found",
			},
			checker: func(response *identityverification.IdentityVerificationResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := buildIdentitiesApi().IdentityVerification

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.CreateIdentityVerification(tc.request))
		})
	}
}

func TestGetIdentityVerification(t *testing.T) {
	t.Skip("Avoid creating identity resources all the time")
	cases := []struct {
		name           string
		verificationId string
		checker        func(*identityverification.IdentityVerificationResponse, error)
	}{
		{
			name:           "when verification exists then should return 200",
			verificationId: identityVerificationId,
			checker: func(response *identityverification.IdentityVerificationResponse, err error) {
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, identityVerificationId, response.Id)
				validateIdentityVerificationResponse(t, response)
			},
		},
		{
			name:           "when verification not found then should return error",
			verificationId: "idv_not_found",
			checker: func(response *identityverification.IdentityVerificationResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := buildIdentitiesApi().IdentityVerification

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetIdentityVerification(tc.verificationId))
		})
	}
}

func TestAnonymizeIdentityVerification(t *testing.T) {
	t.Skip("Avoid creating identity resources all the time")
	cases := []struct {
		name           string
		verificationId string
		checker        func(*identityverification.IdentityVerificationResponse, error)
	}{
		{
			name:           "when verification exists then should return 200",
			verificationId: identityVerificationId,
			checker: func(response *identityverification.IdentityVerificationResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotEmpty(t, response.Id)
			},
		},
		{
			name:           "when verification not found then should return error",
			verificationId: "idv_not_found",
			checker: func(response *identityverification.IdentityVerificationResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := buildIdentitiesApi().IdentityVerification

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.AnonymizeIdentityVerification(tc.verificationId))
		})
	}
}

func TestCreateIdentityVerificationAttempt(t *testing.T) {
	t.Skip("Avoid creating identity resources all the time")
	cases := []struct {
		name           string
		verificationId string
		request        identityverification.CreateIdentityVerificationAttemptRequest
		checker        func(*identityverification.IdentityVerificationAttemptResponse, error)
	}{
		{
			name:           "when request is valid then should return 201",
			verificationId: identityVerificationId,
			request:        createIdentityVerificationAttemptRequest(),
			checker: func(response *identityverification.IdentityVerificationAttemptResponse, err error) {
				assert.Nil(t, err)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
				validateIdentityVerificationAttemptResponse(t, response)
			},
		},
		{
			name:           "when verification not found then should return error",
			verificationId: "idv_not_found",
			request:        createIdentityVerificationAttemptRequest(),
			checker: func(response *identityverification.IdentityVerificationAttemptResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := buildIdentitiesApi().IdentityVerification

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.CreateIdentityVerificationAttempt(tc.verificationId, tc.request))
		})
	}
}

func TestGetIdentityVerificationAttempts(t *testing.T) {
	t.Skip("Avoid creating identity resources all the time")
	cases := []struct {
		name           string
		verificationId string
		checker        func(*identityverification.IdentityVerificationAttemptsResponse, error)
	}{
		{
			name:           "when verification exists then should return 200",
			verificationId: identityVerificationId,
			checker: func(response *identityverification.IdentityVerificationAttemptsResponse, err error) {
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.Data)
			},
		},
		{
			name:           "when verification not found then should return error",
			verificationId: "idv_not_found",
			checker: func(response *identityverification.IdentityVerificationAttemptsResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := buildIdentitiesApi().IdentityVerification

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetIdentityVerificationAttempts(tc.verificationId))
		})
	}
}

func TestGetIdentityVerificationAttempt(t *testing.T) {
	t.Skip("Avoid creating identity resources all the time")
	cases := []struct {
		name           string
		verificationId string
		attemptId      string
		checker        func(*identityverification.IdentityVerificationAttemptResponse, error)
	}{
		{
			name:           "when verification and attempt exist then should return 200",
			verificationId: identityVerificationId,
			attemptId:      idvAttemptId,
			checker: func(response *identityverification.IdentityVerificationAttemptResponse, err error) {
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				validateIdentityVerificationAttemptResponse(t, response)
			},
		},
		{
			name:           "when verification not found then should return error",
			verificationId: "idv_not_found",
			attemptId:      "atm_not_found",
			checker: func(response *identityverification.IdentityVerificationAttemptResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := buildIdentitiesApi().IdentityVerification

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetIdentityVerificationAttempt(tc.verificationId, tc.attemptId))
		})
	}
}

func TestGetIdentityVerificationReport(t *testing.T) {
	t.Skip("Avoid creating identity resources all the time")
	cases := []struct {
		name           string
		verificationId string
		checker        func(*identityverification.IdentityVerificationReportResponse, error)
	}{
		{
			name:           "when verification exists then should return 200",
			verificationId: identityVerificationId,
			checker: func(response *identityverification.IdentityVerificationReportResponse, err error) {
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotEmpty(t, response.SignedUrl)
			},
		},
		{
			name:           "when verification not found then should return error",
			verificationId: "idv_not_found",
			checker: func(response *identityverification.IdentityVerificationReportResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := buildIdentitiesApi().IdentityVerification

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetIdentityVerificationReport(tc.verificationId))
		})
	}
}

// # common methods

func createIdentityVerificationAndAttemptRequest() identityverification.CreateIdentityVerificationAndAttemptRequest {
	return identityverification.CreateIdentityVerificationAndAttemptRequest{
		ApplicantId: applicantResponse.Id,
		RedirectUrl: SuccessUrl,
		DeclaredData: &identities.DeclaredData{
			Name: "John Doe",
		},
	}
}

func createIdentityVerificationRequest() identityverification.CreateIdentityVerificationRequest {
	return identityverification.CreateIdentityVerificationRequest{
		ApplicantId: applicantResponse.Id,
		DeclaredData: &identities.DeclaredData{
			Name: "John Doe",
		},
	}
}

func createIdentityVerificationAttemptRequest() identityverification.CreateIdentityVerificationAttemptRequest {
	return identityverification.CreateIdentityVerificationAttemptRequest{
		RedirectUrl: SuccessUrl,
		ClientInformation: &identities.ClientInformation{
			PreSelectedResidenceCountry: "US",
			PreSelectedLanguage:         "en-US",
		},
	}
}

func validateIdentityVerificationResponse(t *testing.T, response *identityverification.IdentityVerificationResponse) {
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.Id)
	assert.NotEmpty(t, response.ApplicantId)
	assert.NotNil(t, response.Status)
}

func validateIdentityVerificationAttemptResponse(t *testing.T, response *identityverification.IdentityVerificationAttemptResponse) {
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.Id)
	assert.NotNil(t, response.Status)
}
