package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v2/errors"
	"github.com/checkout/checkout-sdk-go/v2/identities"
	addressdocumentverification "github.com/checkout/checkout-sdk-go/v2/identities/addressdocumentverification"
)

func createAddressDocumentVerificationRequest() addressdocumentverification.CreateAddressDocumentVerificationRequest {
	return addressdocumentverification.CreateAddressDocumentVerificationRequest{
		ApplicantId:   "aplt_tkoi5db4hryu5cei5vwoabr7we",
		UserJourneyId: "usj_tkoi5db4hryu5cei5vwoabr7we",
		DeclaredData: &identities.DeclaredData{
			Name: "Hannah Bret",
		},
	}
}

func TestCreateAddressDocumentVerification(t *testing.T) {
	t.Skip("Avoid creating identity resources all the time")
	cases := []struct {
		name    string
		request addressdocumentverification.CreateAddressDocumentVerificationRequest
		checker func(*addressdocumentverification.AddressDocumentVerificationResponse, error)
	}{
		{
			name:    "when request is valid then should return 201",
			request: createAddressDocumentVerificationRequest(),
			checker: func(response *addressdocumentverification.AddressDocumentVerificationResponse, err error) {
				assert.Nil(t, err)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
				assert.NotEmpty(t, response.Id)
			},
		},
		{
			name: "when applicant not found then should return error",
			request: addressdocumentverification.CreateAddressDocumentVerificationRequest{
				ApplicantId:   "aplt_not_found",
				UserJourneyId: "usj_test",
			},
			checker: func(response *addressdocumentverification.AddressDocumentVerificationResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := buildIdentitiesApi().AddressDocumentVerification

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.CreateAddressDocumentVerification(tc.request))
		})
	}
}

func TestGetAddressDocumentVerification(t *testing.T) {
	t.Skip("Avoid creating identity resources all the time")

	client := buildIdentitiesApi().AddressDocumentVerification
	response, err := client.GetAddressDocumentVerification("adv_tkoi5db4hryu5cei5vwoabr7we")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
}

func TestAnonymizeAddressDocumentVerification(t *testing.T) {
	t.Skip("Avoid creating identity resources all the time")

	client := buildIdentitiesApi().AddressDocumentVerification
	_, err := client.AnonymizeAddressDocumentVerification("adv_tkoi5db4hryu5cei5vwoabr7we")
	assert.Nil(t, err)
}

func TestCreateAddressDocumentVerificationAttempt(t *testing.T) {
	t.Skip("Avoid creating identity resources all the time")

	client := buildIdentitiesApi().AddressDocumentVerification
	request := addressdocumentverification.CreateAddressDocumentVerificationAttemptRequest{
		Document: "base64-encoded-document-image-data",
	}
	_, err := client.CreateAddressDocumentVerificationAttempt("adv_tkoi5db4hryu5cei5vwoabr7we", request)
	assert.Nil(t, err)
}

func TestGetAddressDocumentVerificationAttempts(t *testing.T) {
	t.Skip("Avoid creating identity resources all the time")

	client := buildIdentitiesApi().AddressDocumentVerification
	_, err := client.GetAddressDocumentVerificationAttempts("adv_tkoi5db4hryu5cei5vwoabr7we")
	assert.Nil(t, err)
}

func TestGetAddressDocumentVerificationAttempt(t *testing.T) {
	t.Skip("Avoid creating identity resources all the time")

	client := buildIdentitiesApi().AddressDocumentVerification
	_, err := client.GetAddressDocumentVerificationAttempt("adv_tkoi5db4hryu5cei5vwoabr7we", "adva_tkoi5db4hryu5cei5vwoabr7we")
	assert.Nil(t, err)
}

func TestGetAddressDocumentVerificationReport(t *testing.T) {
	t.Skip("Avoid creating identity resources all the time")

	client := buildIdentitiesApi().AddressDocumentVerification
	_, err := client.GetAddressDocumentVerificationReport("adv_tkoi5db4hryu5cei5vwoabr7we")
	assert.Nil(t, err)
}
