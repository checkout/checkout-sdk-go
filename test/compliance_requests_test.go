package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v2/common"
	"github.com/checkout/checkout-sdk-go/v2/compliancerequests"
)

// tests

func TestGetComplianceRequest(t *testing.T) {
	t.Skip("Requires a payment ID associated with an active compliance request")
	cases := []struct {
		name      string
		paymentId string
		checker   func(*compliancerequests.GetComplianceRequestResponse, error)
	}{
		{
			name:      "when payment id is valid then return compliance request details",
			paymentId: "pay_fun26akvvjjerahhctaq2uzhu4",
			checker: func(response *compliancerequests.GetComplianceRequestResponse, err error) {
				assert.Nil(t, err)
				assertComplianceRequestResponse(t, response)
			},
		},
		{
			name:      "when payment id is invalid then return error",
			paymentId: "pay_invalid",
			checker: func(response *compliancerequests.GetComplianceRequestResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
			},
		},
	}

	client := OAuthApi().ComplianceRequests

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetComplianceRequest(tc.paymentId))
		})
	}
}

func TestRespondToComplianceRequest(t *testing.T) {
	t.Skip("Requires a payment ID associated with an active compliance request")
	cases := []struct {
		name      string
		paymentId string
		request   compliancerequests.RespondToComplianceRequestRequest
		checker   func(*common.MetadataResponse, error)
	}{
		{
			name:      "when request is valid then respond to compliance request",
			paymentId: "pay_fun26akvvjjerahhctaq2uzhu4",
			request:   buildRespondToComplianceRequestIntegrationRequest(),
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusAccepted, response.HttpMetadata.StatusCode)
			},
		},
	}

	client := OAuthApi().ComplianceRequests

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.RespondToComplianceRequest(tc.paymentId, tc.request))
		})
	}
}

// common methods

func buildRespondToComplianceRequestIntegrationRequest() compliancerequests.RespondToComplianceRequestRequest {
	return compliancerequests.RespondToComplianceRequestRequest{
		Fields: compliancerequests.ComplianceRespondedFields{
			Sender: []compliancerequests.ComplianceRespondedField{
				{Name: "date_of_birth", Value: "2000-01-01", NotAvailable: false},
			},
		},
		Comments: "Providing the requested compliance information",
	}
}

func assertComplianceRequestResponse(t *testing.T, response *compliancerequests.GetComplianceRequestResponse) {
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.PaymentId)
	assert.NotEmpty(t, response.Status)
}
