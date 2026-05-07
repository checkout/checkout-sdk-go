package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v2/nas"

	applicants "github.com/checkout/checkout-sdk-go/v2/identities/applicants"
	faceauthentication "github.com/checkout/checkout-sdk-go/v2/identities/faceauthentication"
	iddocumentverification "github.com/checkout/checkout-sdk-go/v2/identities/iddocumentverification"
	identityverification "github.com/checkout/checkout-sdk-go/v2/identities/identityverification"
)

var (
	identitiesApi *nas.Api

	applicantResponse *applicants.ApplicantResponse

	identityVerificationId      string
	idvAttemptId                string
	faceAuthenticationId        string
	faceAuthAttemptId           string
	idDocumentVerificationId    string
	idDocVerificationAttemptId  string
)

func TestSetupIdentities(t *testing.T) {
	t.Skip("Avoid creating identity resources all the time")

	applicantResponse = setupApplicant(t)

	idvResponse := setupIdentityVerification(t, applicantResponse.Id)
	identityVerificationId = idvResponse.Id

	idvAttemptResponse := setupIdentityVerificationAttempt(t, identityVerificationId)
	idvAttemptId = idvAttemptResponse.Id

	faceAuthResponse := setupFaceAuthentication(t, applicantResponse.Id)
	faceAuthenticationId = faceAuthResponse.Id

	faceAuthAttemptResponse := setupFaceAuthenticationAttempt(t, faceAuthenticationId)
	faceAuthAttemptId = faceAuthAttemptResponse.Id

	idDocResponse := setupIdDocumentVerification(t, applicantResponse.Id)
	idDocumentVerificationId = idDocResponse.Id

	idDocAttemptResponse := setupIdDocumentVerificationAttempt(t, idDocumentVerificationId)
	idDocVerificationAttemptId = idDocAttemptResponse.Id
}

func buildIdentitiesApi() *nas.Api {
	if identitiesApi == nil {
		identitiesApi = DefaultApi()
	}
	return identitiesApi
}

func setupApplicant(t *testing.T) *applicants.ApplicantResponse {
	request := applicants.CreateApplicantRequest{
		ExternalApplicantId:   "ext_applicant_setup",
		Email:                 GenerateRandomEmail(),
		ExternalApplicantName: "John Doe",
	}
	response, err := buildIdentitiesApi().Applicants.CreateApplicant(request)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("error creating applicant - %s", err.Error()))
	}
	return response
}

func setupIdentityVerification(t *testing.T, applicantId string) *identityverification.IdentityVerificationResponse {
	request := identityverification.CreateIdentityVerificationRequest{
		ApplicantId: applicantId,
	}
	response, err := buildIdentitiesApi().IdentityVerification.CreateIdentityVerification(request)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("error creating identity verification - %s", err.Error()))
	}
	return response
}

func setupIdentityVerificationAttempt(t *testing.T, verificationId string) *identityverification.IdentityVerificationAttemptResponse {
	request := identityverification.CreateIdentityVerificationAttemptRequest{
		RedirectUrl: SuccessUrl,
	}
	response, err := buildIdentitiesApi().IdentityVerification.CreateIdentityVerificationAttempt(verificationId, request)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("error creating identity verification attempt - %s", err.Error()))
	}
	return response
}

func setupFaceAuthentication(t *testing.T, applicantId string) *faceauthentication.FaceAuthenticationResponse {
	request := faceauthentication.CreateFaceAuthenticationRequest{
		ApplicantId:   applicantId,
		UserJourneyId: "uj_test_setup",
	}
	response, err := buildIdentitiesApi().FaceAuthentication.CreateFaceAuthentication(request)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("error creating face authentication - %s", err.Error()))
	}
	return response
}

func setupFaceAuthenticationAttempt(t *testing.T, faceAuthId string) *faceauthentication.FaceAuthenticationAttemptResponse {
	request := faceauthentication.CreateFaceAuthenticationAttemptRequest{
		RedirectUrl: SuccessUrl,
	}
	response, err := buildIdentitiesApi().FaceAuthentication.CreateFaceAuthenticationAttempt(faceAuthId, request)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("error creating face authentication attempt - %s", err.Error()))
	}
	return response
}

func setupIdDocumentVerification(t *testing.T, applicantId string) *iddocumentverification.IdDocumentVerificationResponse {
	request := iddocumentverification.CreateIdDocumentVerificationRequest{
		ApplicantId:   applicantId,
		UserJourneyId: "uj_test_setup",
	}
	response, err := buildIdentitiesApi().IdDocumentVerification.CreateIdDocumentVerification(request)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("error creating id document verification - %s", err.Error()))
	}
	return response
}

func setupIdDocumentVerificationAttempt(t *testing.T, verificationId string) *iddocumentverification.IdDocumentVerificationAttemptResponse {
	request := iddocumentverification.CreateIdDocumentVerificationAttemptRequest{
		DocumentFront: "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mNk+M9QDwADhgGAWjR9awAAAABJRU5ErkJggg==",
	}
	response, err := buildIdentitiesApi().IdDocumentVerification.CreateIdDocumentVerificationAttempt(verificationId, request)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("error creating id document verification attempt - %s", err.Error()))
	}
	return response
}
