package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v2/errors"
	faceauthentication "github.com/checkout/checkout-sdk-go/v2/identities/faceauthentication"
	"github.com/checkout/checkout-sdk-go/v2/identities"
)

// # tests

func TestCreateFaceAuthentication(t *testing.T) {
	t.Skip("Avoid creating identity resources all the time")
	cases := []struct {
		name    string
		request faceauthentication.CreateFaceAuthenticationRequest
		checker func(*faceauthentication.FaceAuthenticationResponse, error)
	}{
		{
			name:    "when request is valid then should return 201",
			request: createFaceAuthenticationRequest(),
			checker: func(response *faceauthentication.FaceAuthenticationResponse, err error) {
				assert.Nil(t, err)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
				validateFaceAuthenticationResponse(t, response)
			},
		},
		{
			name: "when applicant not found then should return error",
			request: faceauthentication.CreateFaceAuthenticationRequest{
				ApplicantId:   "aplt_not_found",
				UserJourneyId: "uj_test",
			},
			checker: func(response *faceauthentication.FaceAuthenticationResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := buildIdentitiesApi().FaceAuthentication

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.CreateFaceAuthentication(tc.request))
		})
	}
}

func TestGetFaceAuthentication(t *testing.T) {
	t.Skip("Avoid creating identity resources all the time")
	cases := []struct {
		name                 string
		faceAuthenticationId string
		checker              func(*faceauthentication.FaceAuthenticationResponse, error)
	}{
		{
			name:                 "when face authentication exists then should return 200",
			faceAuthenticationId: faceAuthenticationId,
			checker: func(response *faceauthentication.FaceAuthenticationResponse, err error) {
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, faceAuthenticationId, response.Id)
				validateFaceAuthenticationResponse(t, response)
			},
		},
		{
			name:                 "when face authentication not found then should return error",
			faceAuthenticationId: "face_auth_not_found",
			checker: func(response *faceauthentication.FaceAuthenticationResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := buildIdentitiesApi().FaceAuthentication

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetFaceAuthentication(tc.faceAuthenticationId))
		})
	}
}

func TestAnonymizeFaceAuthentication(t *testing.T) {
	t.Skip("Avoid creating identity resources all the time")
	cases := []struct {
		name                 string
		faceAuthenticationId string
		checker              func(*faceauthentication.FaceAuthenticationResponse, error)
	}{
		{
			name:                 "when face authentication exists then should return 200",
			faceAuthenticationId: faceAuthenticationId,
			checker: func(response *faceauthentication.FaceAuthenticationResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotEmpty(t, response.Id)
			},
		},
		{
			name:                 "when face authentication not found then should return error",
			faceAuthenticationId: "face_auth_not_found",
			checker: func(response *faceauthentication.FaceAuthenticationResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := buildIdentitiesApi().FaceAuthentication

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.AnonymizeFaceAuthentication(tc.faceAuthenticationId))
		})
	}
}

func TestCreateFaceAuthenticationAttempt(t *testing.T) {
	t.Skip("Avoid creating identity resources all the time")
	cases := []struct {
		name                 string
		faceAuthenticationId string
		request              faceauthentication.CreateFaceAuthenticationAttemptRequest
		checker              func(*faceauthentication.FaceAuthenticationAttemptResponse, error)
	}{
		{
			name:                 "when request is valid then should return 201",
			faceAuthenticationId: faceAuthenticationId,
			request:              createFaceAuthenticationAttemptRequest(),
			checker: func(response *faceauthentication.FaceAuthenticationAttemptResponse, err error) {
				assert.Nil(t, err)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
				validateFaceAuthenticationAttemptResponse(t, response)
			},
		},
		{
			name:                 "when face authentication not found then should return error",
			faceAuthenticationId: "face_auth_not_found",
			request:              createFaceAuthenticationAttemptRequest(),
			checker: func(response *faceauthentication.FaceAuthenticationAttemptResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := buildIdentitiesApi().FaceAuthentication

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.CreateFaceAuthenticationAttempt(tc.faceAuthenticationId, tc.request))
		})
	}
}

func TestGetFaceAuthenticationAttempts(t *testing.T) {
	t.Skip("Avoid creating identity resources all the time")
	cases := []struct {
		name                 string
		faceAuthenticationId string
		checker              func(*faceauthentication.FaceAuthenticationAttemptsResponse, error)
	}{
		{
			name:                 "when face authentication exists then should return 200",
			faceAuthenticationId: faceAuthenticationId,
			checker: func(response *faceauthentication.FaceAuthenticationAttemptsResponse, err error) {
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.Data)
			},
		},
		{
			name:                 "when face authentication not found then should return error",
			faceAuthenticationId: "face_auth_not_found",
			checker: func(response *faceauthentication.FaceAuthenticationAttemptsResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := buildIdentitiesApi().FaceAuthentication

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetFaceAuthenticationAttempts(tc.faceAuthenticationId))
		})
	}
}

func TestGetFaceAuthenticationAttempt(t *testing.T) {
	t.Skip("Avoid creating identity resources all the time")
	cases := []struct {
		name                 string
		faceAuthenticationId string
		attemptId            string
		checker              func(*faceauthentication.FaceAuthenticationAttemptResponse, error)
	}{
		{
			name:                 "when face authentication and attempt exist then should return 200",
			faceAuthenticationId: faceAuthenticationId,
			attemptId:            faceAuthAttemptId,
			checker: func(response *faceauthentication.FaceAuthenticationAttemptResponse, err error) {
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				validateFaceAuthenticationAttemptResponse(t, response)
			},
		},
		{
			name:                 "when face authentication not found then should return error",
			faceAuthenticationId: "face_auth_not_found",
			attemptId:            "atm_not_found",
			checker: func(response *faceauthentication.FaceAuthenticationAttemptResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := buildIdentitiesApi().FaceAuthentication

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetFaceAuthenticationAttempt(tc.faceAuthenticationId, tc.attemptId))
		})
	}
}

// # common methods

func createFaceAuthenticationRequest() faceauthentication.CreateFaceAuthenticationRequest {
	return faceauthentication.CreateFaceAuthenticationRequest{
		ApplicantId:   applicantResponse.Id,
		UserJourneyId: "uj_test_face_auth",
	}
}

func createFaceAuthenticationAttemptRequest() faceauthentication.CreateFaceAuthenticationAttemptRequest {
	return faceauthentication.CreateFaceAuthenticationAttemptRequest{
		RedirectUrl: SuccessUrl,
		ClientInformation: &identities.ClientInformation{
			PreSelectedResidenceCountry: "US",
			PreSelectedLanguage:         "en-US",
		},
	}
}

func validateFaceAuthenticationResponse(t *testing.T, response *faceauthentication.FaceAuthenticationResponse) {
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.Id)
	assert.NotEmpty(t, response.ApplicantId)
	assert.NotEmpty(t, response.UserJourneyId)
	assert.NotNil(t, response.Status)
}

func validateFaceAuthenticationAttemptResponse(t *testing.T, response *faceauthentication.FaceAuthenticationAttemptResponse) {
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.Id)
	assert.NotNil(t, response.Status)
}
