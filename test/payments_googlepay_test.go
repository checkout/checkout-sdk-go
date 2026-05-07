package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v2/common"
	"github.com/checkout/checkout-sdk-go/v2/payments/googlepay"
)

// tests

func TestCreateEnrollment(t *testing.T) {
	t.Skip("Requires OAuth credentials with vault:gpayme-enrollment scope and a valid entity")
	cases := []struct {
		name    string
		request googlepay.CreateEnrollmentRequest
		checker func(*googlepay.CreateEnrollmentResponse, error)
	}{
		{
			name:    "when request is valid then create enrollment",
			request: buildGooglePayCreateEnrollmentRequest(),
			checker: func(response *googlepay.CreateEnrollmentResponse, err error) {
				assert.Nil(t, err)
				assertCreateEnrollmentResponse(t, response)
			},
		},
	}

	client := OAuthApi().GooglePay

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.CreateEnrollment(tc.request))
		})
	}
}

func TestRegisterDomain(t *testing.T) {
	t.Skip("Requires an actively enrolled Google Pay entity")
	var (
		entityId = "ent_uzm3uxtssvmuxnyrfdffcyjxeu"
	)

	cases := []struct {
		name     string
		entityId string
		request  googlepay.RegisterDomainRequest
		checker  func(*common.MetadataResponse, error)
	}{
		{
			name:     "when domain is valid then register domain",
			entityId: entityId,
			request:  buildGooglePayRegisterDomainRequest(),
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusNoContent, response.HttpMetadata.StatusCode)
			},
		},
	}

	client := OAuthApi().GooglePay

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.RegisterDomain(tc.entityId, tc.request))
		})
	}
}

func TestGetRegisteredDomains(t *testing.T) {
	t.Skip("Requires an actively enrolled Google Pay entity with registered domains")
	var (
		entityId = "ent_uzm3uxtssvmuxnyrfdffcyjxeu"
	)

	cases := []struct {
		name     string
		entityId string
		checker  func(*googlepay.DomainListResponse, error)
	}{
		{
			name:     "when entity has registered domains then return domain list",
			entityId: entityId,
			checker: func(response *googlepay.DomainListResponse, err error) {
				assert.Nil(t, err)
				assertDomainListResponse(t, response)
			},
		},
		{
			name:     "when entity id is invalid then return error",
			entityId: "ent_invalid",
			checker: func(response *googlepay.DomainListResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
			},
		},
	}

	client := OAuthApi().GooglePay

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetRegisteredDomains(tc.entityId))
		})
	}
}

func TestGetEnrollmentState(t *testing.T) {
	t.Skip("Requires an actively enrolled Google Pay entity")
	var (
		entityId = "ent_uzm3uxtssvmuxnyrfdffcyjxeu"
	)

	cases := []struct {
		name     string
		entityId string
		checker  func(*googlepay.EnrollmentStateResponse, error)
	}{
		{
			name:     "when entity is enrolled then return enrollment state",
			entityId: entityId,
			checker: func(response *googlepay.EnrollmentStateResponse, err error) {
				assert.Nil(t, err)
				assertEnrollmentStateResponse(t, response)
			},
		},
		{
			name:     "when entity id is invalid then return error",
			entityId: "ent_invalid",
			checker: func(response *googlepay.EnrollmentStateResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
			},
		},
	}

	client := OAuthApi().GooglePay

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetEnrollmentState(tc.entityId))
		})
	}
}

// common methods

func buildGooglePayCreateEnrollmentRequest() googlepay.CreateEnrollmentRequest {
	return googlepay.CreateEnrollmentRequest{
		EntityId:             "ent_uzm3uxtssvmuxnyrfdffcyjxeu",
		EmailAddress:         "test@gmail.com",
		AcceptTermsOfService: true,
	}
}

func buildGooglePayRegisterDomainRequest() googlepay.RegisterDomainRequest {
	return googlepay.RegisterDomainRequest{
		WebDomain: "some.example.com",
	}
}

func assertCreateEnrollmentResponse(t *testing.T, response *googlepay.CreateEnrollmentResponse) {
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
	assert.NotNil(t, response.TosAcceptedTime)
	assert.NotEmpty(t, response.State)
}

func assertDomainListResponse(t *testing.T, response *googlepay.DomainListResponse) {
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
	assert.NotNil(t, response.Domains)
}

func assertEnrollmentStateResponse(t *testing.T, response *googlepay.EnrollmentStateResponse) {
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
	assert.NotEmpty(t, response.State)
}
