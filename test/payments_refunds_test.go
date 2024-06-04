package test

import (
	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/payments/nas"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/payments"
)

func TestRefundCardPayment(t *testing.T) {
	paymentResponse := makeCardPayment(t, true, 10)

	order := payments.Order{
		Name:        "OrderTest",
		Quantity:    88,
		TotalAmount: 99,
	}

	bank := common.BankDetails{
		Name:    "Lloyds TSB",
		Branch:  "Bournemouth",
		Address: Address(),
	}

	destination := common.Destination{
		AccountType:   common.Savings,
		AccountNumber: "13654567455",
		BankCode:      "23-456",
		BranchCode:    "6443",
		Iban:          "HU93116000060000000012345676",
		Bban:          "3704 0044 0532 0130 00",
		SwiftBic:      "37040044",
		Country:       common.GB,
		AccountHolder: AccountHolder(),
		Bank:          &bank,
	}

	refundRequest := payments.RefundRequest{
		Amount:    paymentResponse.Amount,
		Reference: uuid.New().String(),
		Items: []payments.Order{
			order,
		},
		Destination: &destination,
	}

	cases := []struct {
		name          string
		paymentId     string
		refundRequest payments.RefundRequest
		checkerRefund func(*payments.RefundResponse, error)
		checkerGet    func(*nas.GetPaymentResponse, error)
	}{
		{
			name:          "when request is valid then return a refund response",
			paymentId:     paymentResponse.Id,
			refundRequest: refundRequest,
			checkerRefund: func(response *payments.RefundResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotEmpty(t, response.Reference)
				assert.NotEmpty(t, response.ActionId)
				assert.NotEmpty(t, response.Links)
				assert.NotEmpty(t, response.Links["payment"])
			},
			checkerGet: func(response *nas.GetPaymentResponse, err error) {
				assert.NotEmpty(t, response.Balances)
				assert.Equal(t, int64(10), response.Balances.TotalAuthorized)
				assert.Equal(t, int64(10), response.Balances.TotalCaptured)
				assert.Equal(t, int64(10), response.Balances.TotalRefunded)
				assert.Equal(t, int64(0), response.Balances.TotalVoided)
				assert.Equal(t, int64(0), response.Balances.AvailableToCapture)
				assert.Equal(t, int64(0), response.Balances.AvailableToRefund)
				assert.Equal(t, int64(0), response.Balances.AvailableToVoid)
			},
		},
	}

	client := DefaultApi().Payments

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			Wait(time.Duration(3))
			tc.checkerRefund(client.RefundPayment(tc.paymentId, &tc.refundRequest, nil))
			tc.checkerGet(client.GetPaymentDetails(tc.paymentId))
		})
	}
}

func TestRefundCardPaymentIdempotently(t *testing.T) {
	paymentResponse := makeCardPayment(t, true, 10)

	refundRequest := payments.RefundRequest{
		Amount:    5,
		Reference: uuid.New().String(),
		Metadata:  nil,
	}

	idempotencyKey := uuid.New().String()

	cases := []struct {
		name           string
		paymentId      string
		idempotencyKey string
		refundRequest  payments.RefundRequest
		checker        func(*payments.RefundResponse, error)
	}{
		{
			name:           "when request is valid with idempotencyKey request then return a refund response",
			paymentId:      paymentResponse.Id,
			refundRequest:  refundRequest,
			idempotencyKey: idempotencyKey,
			checker: func(response *payments.RefundResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)

				response2, err := DefaultApi().Payments.RefundPayment(paymentResponse.Id, &refundRequest, &idempotencyKey)
				assert.Nil(t, err)
				assert.NotNil(t, response2)

				assert.Equal(t, response.ActionId, response2.ActionId)
			},
		},
	}

	client := DefaultApi().Payments

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			Wait(time.Duration(3))
			tc.checker(client.RefundPayment(tc.paymentId, &tc.refundRequest, &tc.idempotencyKey))
		})
	}
}
