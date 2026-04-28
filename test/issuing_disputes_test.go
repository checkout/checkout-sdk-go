package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v2/common"
	"github.com/checkout/checkout-sdk-go/v2/errors"

	disputes "github.com/checkout/checkout-sdk-go/v2/issuing/disputes"
)

// # tests

func TestCreateDispute(t *testing.T) {
	t.Skip("Avoid creating cards all the time")
	idempotencyKey := "test-idempotency-key-create"
	cases := []struct {
		name           string
		request        disputes.CreateDisputeRequest
		idempotencyKey *string
		checker        func(*disputes.IssuingDisputeResponse, error)
	}{
		{
			name:           "when request is correct then should return 201",
			request:        createDisputeRequest(),
			idempotencyKey: &idempotencyKey,
			checker: func(response *disputes.IssuingDisputeResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.Status)
				assert.NotNil(t, response.TransactionId)
			},
		},
		{
			name:           "when request is invalid then return error",
			request:        disputes.CreateDisputeRequest{},
			idempotencyKey: nil,
			checker: func(response *disputes.IssuingDisputeResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
			},
		},
	}

	client := buildIssuingClientApi().Issuing

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.CreateDispute(tc.request, tc.idempotencyKey))
		})
	}
}

func TestGetDispute(t *testing.T) {
	t.Skip("Avoid creating cards all the time")
	cases := []struct {
		name      string
		disputeId string
		checker   func(*disputes.IssuingDisputeResponse, error)
	}{
		{
			name:      "when request is correct then should return 200",
			disputeId: disputeResponse.Id,
			checker: func(response *disputes.IssuingDisputeResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, disputeResponse.Id, response.Id)
				assert.NotNil(t, response.Status)
				assert.NotNil(t, response.TransactionId)
			},
		},
		{
			name:      "when dispute id not found then return error",
			disputeId: "idsp_not_found",
			checker: func(response *disputes.IssuingDisputeResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := buildIssuingClientApi().Issuing

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetDispute(tc.disputeId))
		})
	}
}

func TestCancelDispute(t *testing.T) {
	t.Skip("Avoid creating cards all the time")
	idempotencyKey := "test-idempotency-key-cancel"
	cases := []struct {
		name           string
		disputeId      string
		idempotencyKey *string
		checker        func(*common.MetadataResponse, error)
	}{
		{
			name:           "when request is correct then should return 200",
			disputeId:      disputeResponse.Id,
			idempotencyKey: &idempotencyKey,
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
			},
		},
		{
			name:           "when dispute id not found then return error",
			disputeId:      "idsp_not_found",
			idempotencyKey: nil,
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := buildIssuingClientApi().Issuing

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.CancelDispute(tc.disputeId, tc.idempotencyKey))
		})
	}
}

func TestEscalateDispute(t *testing.T) {
	t.Skip("Avoid creating cards all the time")
	idempotencyKey := "test-idempotency-key-escalate"
	cases := []struct {
		name           string
		disputeId      string
		request        disputes.EscalateDisputeRequest
		idempotencyKey *string
		checker        func(*common.MetadataResponse, error)
	}{
		{
			name:           "when request is correct then should return 200",
			disputeId:      disputeResponse.Id,
			request:        escalateDisputeRequest(),
			idempotencyKey: &idempotencyKey,
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
			},
		},
		{
			name:           "when dispute id not found then return error",
			disputeId:      "idsp_not_found",
			request:        escalateDisputeRequest(),
			idempotencyKey: nil,
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := buildIssuingClientApi().Issuing

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.EscalateDispute(tc.disputeId, tc.request, tc.idempotencyKey))
		})
	}
}

// # common methods

func createDisputeRequest() disputes.CreateDisputeRequest {
	amount := int64(1000)
	return disputes.CreateDisputeRequest{
		TransactionId: virtualCardResponse.Id,
		Reason:        "4837",
		Amount:        &amount,
		Justification: "Customer dispute",
		Evidence: []disputes.DisputeEvidence{
			{
				Name:        "receipt.pdf",
				Content:     "SGVsbG8gV29ybGQ=",
				Description: "Transaction receipt",
			},
		},
	}
}

func escalateDisputeRequest() disputes.EscalateDisputeRequest {
	amount := int64(500)
	return disputes.EscalateDisputeRequest{
		Justification: "Escalating due to additional evidence",
		Amount:        &amount,
		AdditionalEvidence: []disputes.DisputeEvidence{
			{
				Name:        "additional_evidence.pdf",
				Content:     "QWRkaXRpb25hbCBFdmlkZW5jZQ==",
				Description: "Additional supporting documentation",
			},
		},
	}
}

func assertIssuingDisputeResponse(t *testing.T, response *disputes.IssuingDisputeResponse) {
	assert.NotNil(t, response)
	assert.NotNil(t, response.Id)
	assert.NotNil(t, response.Status)
	assert.NotNil(t, response.TransactionId)
}
