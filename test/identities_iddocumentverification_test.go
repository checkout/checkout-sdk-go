package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v2/errors"
	"github.com/checkout/checkout-sdk-go/v2/identities"
	iddocumentverification "github.com/checkout/checkout-sdk-go/v2/identities/iddocumentverification"
)

// # tests

func TestCreateIdDocumentVerification(t *testing.T) {
	t.Skip("Avoid creating identity resources all the time")
	cases := []struct {
		name    string
		request iddocumentverification.CreateIdDocumentVerificationRequest
		checker func(*iddocumentverification.IdDocumentVerificationResponse, error)
	}{
		{
			name:    "when request is valid then should return 201",
			request: createIdDocumentVerificationRequest(),
			checker: func(response *iddocumentverification.IdDocumentVerificationResponse, err error) {
				assert.Nil(t, err)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
				validateIdDocumentVerificationResponse(t, response)
			},
		},
		{
			name: "when applicant not found then should return error",
			request: iddocumentverification.CreateIdDocumentVerificationRequest{
				ApplicantId:   "aplt_not_found",
				UserJourneyId: "uj_test",
			},
			checker: func(response *iddocumentverification.IdDocumentVerificationResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := buildIdentitiesApi().IdDocumentVerification

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.CreateIdDocumentVerification(tc.request))
		})
	}
}

func TestGetIdDocumentVerification(t *testing.T) {
	t.Skip("Avoid creating identity resources all the time")
	cases := []struct {
		name           string
		verificationId string
		checker        func(*iddocumentverification.IdDocumentVerificationResponse, error)
	}{
		{
			name:           "when verification exists then should return 200",
			verificationId: idDocumentVerificationId,
			checker: func(response *iddocumentverification.IdDocumentVerificationResponse, err error) {
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, idDocumentVerificationId, response.Id)
				validateIdDocumentVerificationResponse(t, response)
			},
		},
		{
			name:           "when verification not found then should return error",
			verificationId: "iddoc_not_found",
			checker: func(response *iddocumentverification.IdDocumentVerificationResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := buildIdentitiesApi().IdDocumentVerification

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetIdDocumentVerification(tc.verificationId))
		})
	}
}

func TestAnonymizeIdDocumentVerification(t *testing.T) {
	t.Skip("Avoid creating identity resources all the time")
	cases := []struct {
		name           string
		verificationId string
		checker        func(*iddocumentverification.IdDocumentVerificationResponse, error)
	}{
		{
			name:           "when verification exists then should return 200",
			verificationId: idDocumentVerificationId,
			checker: func(response *iddocumentverification.IdDocumentVerificationResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotEmpty(t, response.Id)
			},
		},
		{
			name:           "when verification not found then should return error",
			verificationId: "iddoc_not_found",
			checker: func(response *iddocumentverification.IdDocumentVerificationResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := buildIdentitiesApi().IdDocumentVerification

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.AnonymizeIdDocumentVerification(tc.verificationId))
		})
	}
}

func TestCreateIdDocumentVerificationAttempt(t *testing.T) {
	t.Skip("Avoid creating identity resources all the time")
	cases := []struct {
		name           string
		verificationId string
		request        iddocumentverification.CreateIdDocumentVerificationAttemptRequest
		checker        func(*iddocumentverification.IdDocumentVerificationAttemptResponse, error)
	}{
		{
			name:           "when request is valid then should return 201",
			verificationId: idDocumentVerificationId,
			request:        createIdDocumentVerificationAttemptRequest(),
			checker: func(response *iddocumentverification.IdDocumentVerificationAttemptResponse, err error) {
				assert.Nil(t, err)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
				validateIdDocumentVerificationAttemptResponse(t, response)
			},
		},
		{
			name:           "when verification not found then should return error",
			verificationId: "iddoc_not_found",
			request:        createIdDocumentVerificationAttemptRequest(),
			checker: func(response *iddocumentverification.IdDocumentVerificationAttemptResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := buildIdentitiesApi().IdDocumentVerification

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.CreateIdDocumentVerificationAttempt(tc.verificationId, tc.request))
		})
	}
}

func TestGetIdDocumentVerificationAttempts(t *testing.T) {
	t.Skip("Avoid creating identity resources all the time")
	cases := []struct {
		name           string
		verificationId string
		checker        func(*iddocumentverification.IdDocumentVerificationAttemptsResponse, error)
	}{
		{
			name:           "when verification exists then should return 200",
			verificationId: idDocumentVerificationId,
			checker: func(response *iddocumentverification.IdDocumentVerificationAttemptsResponse, err error) {
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.Data)
			},
		},
		{
			name:           "when verification not found then should return error",
			verificationId: "iddoc_not_found",
			checker: func(response *iddocumentverification.IdDocumentVerificationAttemptsResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := buildIdentitiesApi().IdDocumentVerification

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetIdDocumentVerificationAttempts(tc.verificationId))
		})
	}
}

func TestGetIdDocumentVerificationAttempt(t *testing.T) {
	t.Skip("Avoid creating identity resources all the time")
	cases := []struct {
		name           string
		verificationId string
		attemptId      string
		checker        func(*iddocumentverification.IdDocumentVerificationAttemptResponse, error)
	}{
		{
			name:           "when verification and attempt exist then should return 200",
			verificationId: idDocumentVerificationId,
			attemptId:      idDocVerificationAttemptId,
			checker: func(response *iddocumentverification.IdDocumentVerificationAttemptResponse, err error) {
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				validateIdDocumentVerificationAttemptResponse(t, response)
			},
		},
		{
			name:           "when verification not found then should return error",
			verificationId: "iddoc_not_found",
			attemptId:      "atm_not_found",
			checker: func(response *iddocumentverification.IdDocumentVerificationAttemptResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := buildIdentitiesApi().IdDocumentVerification

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetIdDocumentVerificationAttempt(tc.verificationId, tc.attemptId))
		})
	}
}

func TestGetIdDocumentVerificationReport(t *testing.T) {
	t.Skip("Avoid creating identity resources all the time")
	cases := []struct {
		name           string
		verificationId string
		checker        func(*iddocumentverification.IdDocumentVerificationReportResponse, error)
	}{
		{
			name:           "when verification exists then should return 200",
			verificationId: idDocumentVerificationId,
			checker: func(response *iddocumentverification.IdDocumentVerificationReportResponse, err error) {
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotEmpty(t, response.SignedUrl)
			},
		},
		{
			name:           "when verification not found then should return error",
			verificationId: "iddoc_not_found",
			checker: func(response *iddocumentverification.IdDocumentVerificationReportResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := buildIdentitiesApi().IdDocumentVerification

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetIdDocumentVerificationReport(tc.verificationId))
		})
	}
}

// # common methods

func createIdDocumentVerificationRequest() iddocumentverification.CreateIdDocumentVerificationRequest {
	return iddocumentverification.CreateIdDocumentVerificationRequest{
		ApplicantId:   applicantResponse.Id,
		UserJourneyId: "uj_test_id_doc",
		DeclaredData: &identities.DeclaredData{
			Name: "John Doe",
		},
	}
}

func createIdDocumentVerificationAttemptRequest() iddocumentverification.CreateIdDocumentVerificationAttemptRequest {
	return iddocumentverification.CreateIdDocumentVerificationAttemptRequest{
		DocumentFront: "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mNk+M9QDwADhgGAWjR9awAAAABJRU5ErkJggg==",
	}
}

func validateIdDocumentVerificationResponse(t *testing.T, response *iddocumentverification.IdDocumentVerificationResponse) {
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.Id)
	assert.NotEmpty(t, response.ApplicantId)
	assert.NotEmpty(t, response.UserJourneyId)
	assert.NotNil(t, response.Status)
}

func validateIdDocumentVerificationAttemptResponse(t *testing.T, response *iddocumentverification.IdDocumentVerificationAttemptResponse) {
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.Id)
	assert.NotNil(t, response.Status)
}
