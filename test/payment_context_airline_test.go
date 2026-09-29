package test

import (
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/payments"
	"github.com/checkout/checkout-sdk-go/v3/payments/contexts"
	sources "github.com/checkout/checkout-sdk-go/v3/payments/nas/sources/contexts"
)

// TestRequestPaymentContextWithAirlineData exercises processing.airline_data against the live API
// on the surface that is strictest about passenger cardinality.
//
// POST /payment-contexts rejects the array form of passenger with 422 passenger_required and
// accepts a single object, which is the opposite of what the specification declares. Nothing else
// in the suite sends airline data to this endpoint, so a regression in
// PaymentContextsAirlineData.MarshalJSON would otherwise ship unnoticed; the equivalent Go test on
// hosted payments is what caught exactly that mistake in the other SDKs.
func TestRequestPaymentContextWithAirlineData(t *testing.T) {
	request := contexts.PaymentContextsRequest{
		Source:              sources.NewPaymentContextsPayPalSource(),
		Amount:              1000,
		Currency:            common.EUR,
		PaymentType:         payments.Regular,
		Capture:             true,
		ProcessingChannelId: os.Getenv("CHECKOUT_PROCESSING_CHANNEL_ID"),
		SuccessUrl:          "https://example.com/payments/success",
		FailureUrl:          "https://example.com/payments/fail",
		Items: []contexts.PaymentContextsItems{
			{Name: "flight", Quantity: 1, UnitPrice: 1000, TotalAmount: 1000},
		},
		Processing: &contexts.PaymentContextsProcessing{
			AirlineData: []contexts.PaymentContextsAirlineData{
				{
					Ticket: &contexts.PaymentContextsTicket{
						Number:                 "045-21351455613",
						IssuingCarrierCode:     "AI",
						TravelPackageIndicator: "B",
					},
					Passenger: []contexts.PaymentContextsPassenger{
						{
							FirstName: "John",
							LastName:  "White",
							Address:   &payments.PassengerAddress{Country: common.GB},
						},
					},
					FlightLegDetails: []contexts.PaymentContextsFlightLegDetails{
						{
							FlightNumber:      "101",
							CarrierCode:       "BA",
							ClassOfTravelling: "J",
							DepartureAirport:  "LHR",
							ArrivalAirport:    "LAX",
							// StopOverCode is deliberately omitted. POST /payment-contexts
							// rejects it with flight_leg_detail_stop_over_code_invalid for
							// every value tried, including the three its own description
							// names (a space, O and X) and the "x" in the swagger example.
							// POST /payments accepts "X" and "x" without complaint. Raised
							// as a spec/API defect; omitting it keeps this test focused on
							// the passenger cardinality it exists to guard.
						},
					},
				},
			},
		},
	}

	response, err := DefaultApi().Contexts.RequestPaymentContexts(request)

	// A 422 here means the serialized passenger cardinality no longer matches what the endpoint
	// accepts. That is the regression this test exists to catch.
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
	assert.NotNil(t, response.Id)
}
