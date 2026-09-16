package hosted

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/payments"
	"github.com/checkout/checkout-sdk-go/v3/payments/nas"
)

// Verifies the fields recently aligned with the Checkout.com swagger spec
// (HostedPaymentsRequest) serialize to their snake_case wire names.
//   - authorization_type
//   - payment_plan
func TestPaymentHostedRequest_AuthorizationTypeAndPaymentPlan(t *testing.T) {
	request := PaymentHostedRequest{
		AuthorizationType: nas.EstimatedAuthorizationType,
		PaymentPlan: &payments.PaymentPlan{
			AmountVariability:   payments.FixedAVT,
			DaysBetweenPayments: 28,
		},
	}

	marshalled, err := json.Marshal(request)

	assert.NoError(t, err)
	assert.Contains(t, string(marshalled), `"authorization_type":"Estimated"`)
	assert.Contains(t, string(marshalled), `"payment_plan":`)
	assert.Contains(t, string(marshalled), `"days_between_payments":28`)
}

// Verifies the new fields are omitted when unset.
func TestPaymentHostedRequest_OmitsNewFieldsWhenUnset(t *testing.T) {
	marshalled, err := json.Marshal(PaymentHostedRequest{})

	assert.NoError(t, err)
	assert.NotContains(t, string(marshalled), "authorization_type")
	assert.NotContains(t, string(marshalled), "payment_plan")
}

// The merchant-reported defect, end to end on the surface it was reported on.
//
// A Hosted Payments Page session carrying processing.accommodation_data must put
// yyyy-MM-dd on the wire for check_in_date, check_out_date and guests[].date_of_birth.
// The specification declares all three as "type": "string", "format": "date".
//
// While they were *time.Time the SDK emitted "2026-10-01T00:00:00Z"; Tamara rejects that
// and the merchant saw a gateway 500. This test is the regression guard for the exact
// payload from the report.
func TestPaymentHostedRequest_AccommodationDatesUseShortDateFormat(t *testing.T) {
	checkIn := common.APIShortDate(time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC))
	checkOut := common.APIShortDate(time.Date(2026, 10, 5, 0, 0, 0, 0, time.UTC))
	dateOfBirth := common.APIShortDate(time.Date(1985, 7, 14, 0, 0, 0, 0, time.UTC))

	request := PaymentHostedRequest{
		Amount:   1000,
		Currency: common.GBP,
		Processing: &payments.ProcessingSettings{
			AccommodationData: []payments.AccommodationData{{
				Name:             "Grand Hotel",
				BookingReference: "HOTEL123",
				CheckInDate:      &checkIn,
				CheckOutDate:     &checkOut,
				Guests: []payments.Guest{{
					FirstName:   "Jane",
					LastName:    "Doe",
					DateOfBirth: &dateOfBirth,
				}},
			}},
		},
	}

	raw, err := json.Marshal(request)
	assert.Nil(t, err)

	var body map[string]interface{}
	assert.Nil(t, json.Unmarshal(raw, &body))

	processing := body["processing"].(map[string]interface{})
	accommodation := processing["accommodation_data"].([]interface{})
	assert.Len(t, accommodation, 1)

	first := accommodation[0].(map[string]interface{})
	assert.Equal(t, "Grand Hotel", first["name"])
	assert.Equal(t, "HOTEL123", first["booking_reference"])
	assert.Equal(t, "2026-10-01", first["check_in_date"])
	assert.Equal(t, "2026-10-05", first["check_out_date"])

	guests := first["guests"].([]interface{})
	assert.Len(t, guests, 1)
	assert.Equal(t, "1985-07-14", guests[0].(map[string]interface{})["date_of_birth"])

	// No field anywhere in the HPP body may carry a time component on a date-only field.
	assert.NotContains(t, string(raw), "T00:00:00")
}
