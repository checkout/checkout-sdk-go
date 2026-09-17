package setups

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/common"
)

// Verifies the Payment Setups schemas aligned with the 2026-06-29 Checkout.com swagger delta:
//   - PaymentSetup: billing_descriptor, presentment_details, terminal, latest_payment
//   - payment_methods: bacs, card_present, pay_by_bank, stablecoin
//   - order.amount_allocations[] (+ commission)
//   - KlarnaAccountHolder.name

func TestPaymentSetupRequest_NewTopLevelFields(t *testing.T) {
	request := PaymentSetupRequest{
		ProcessingChannelId: "pc_test_abcdefghijklmnopqrstuvw",
		Amount:              1000,
		Currency:            common.GBP,
		BillingDescriptor: &PaymentSetupBillingDescriptor{
			Name:      "SDK Test",
			City:      "London",
			Reference: "order-123",
		},
		PresentmentDetails: &PaymentSetupPresentmentDetails{
			Amount:   1200,
			Currency: common.USD,
		},
		Terminal: &PaymentSetupTerminal{
			Id:            "trm12345",
			LocalDateTime: timePtr(t, "2026-06-01T10:00:00Z"),
		},
	}

	marshalled, err := json.Marshal(request)
	assert.NoError(t, err)
	body := string(marshalled)
	assert.Contains(t, body, `"billing_descriptor":`)
	assert.Contains(t, body, `"name":"SDK Test"`)
	assert.Contains(t, body, `"city":"London"`)
	assert.Contains(t, body, `"reference":"order-123"`)
	assert.Contains(t, body, `"presentment_details":`)
	assert.Contains(t, body, `"amount":1200`)
	assert.Contains(t, body, `"currency":"USD"`)
	assert.Contains(t, body, `"terminal":`)
	assert.Contains(t, body, `"id":"trm12345"`)
	assert.Contains(t, body, `"local_date_time":"2026-06-01T10:00:00Z"`)
}

func TestPaymentSetupResponse_NewFieldsAndLatestPayment(t *testing.T) {
	payload := `{
		"id": "psp_test_abcdefghijklmnopqr",
		"billing_descriptor": {"name": "SDK Test", "city": "London", "reference": "order-123"},
		"presentment_details": {"amount": 1200, "currency": "USD"},
		"terminal": {"id": "trm12345", "local_date_time": "2026-06-01T10:00:00Z"},
		"latest_payment": {"id": "pay_test_abcdefghijklmnopqr", "status": "Authorized"}
	}`

	var response PaymentSetupResponse
	assert.NoError(t, json.Unmarshal([]byte(payload), &response))
	assert.Equal(t, "SDK Test", response.BillingDescriptor.Name)
	assert.Equal(t, "London", response.BillingDescriptor.City)
	assert.Equal(t, int64(1200), response.PresentmentDetails.Amount)
	assert.Equal(t, common.USD, response.PresentmentDetails.Currency)
	assert.Equal(t, "trm12345", response.Terminal.Id)
	assert.NotNil(t, response.Terminal.LocalDateTime)
	assert.Equal(t, "pay_test_abcdefghijklmnopqr", response.LatestPayment["id"])
	assert.Equal(t, "Authorized", response.LatestPayment["status"])
}

func TestPaymentMethods_NewConfigs(t *testing.T) {
	methods := PaymentMethods{
		Bacs: &BacsPaymentMethod{
			InstrumentId:      "src_test_abcdefghijklmnopqr",
			AccountNumber:     "12345678",
			BankCode:          "200000",
			Country:           common.GB,
			Currency:          "GBP",
			AllowPartialMatch: true,
			AccountHolder: &BacsAccountHolder{
				Type:      BacsAccountHolderIndividual,
				FirstName: "John",
				LastName:  "Smith",
				Email:     "john.smith@example.com",
			},
		},
		CardPresent: &CardPresentPaymentMethod{
			Track2:    "track2data",
			Emv:       "emvdata",
			EntryMode: "chip",
			Name:      "John Smith",
			Pin: &CardPresentPin{
				KeySetId:    "kset_123",
				Block:       "block_data",
				BlockFormat: "ISO-0",
			},
		},
		PayByBank: &PayByBankPaymentMethod{
			BankId: "bank_123",
			Action: &PayByBankAction{
				Type: "select_bank",
				Banks: []PayByBankBank{
					{BankId: "bank_123", DisplayName: "Test Bank", LogoUrl: "https://example.com/logo.png", Available: true},
				},
			},
		},
		Stablecoin: &StablecoinPaymentMethod{},
	}

	marshalled, err := json.Marshal(methods)
	assert.NoError(t, err)
	body := string(marshalled)
	// wiring keys
	assert.Contains(t, body, `"bacs":`)
	assert.Contains(t, body, `"card_present":`)
	assert.Contains(t, body, `"pay_by_bank":`)
	assert.Contains(t, body, `"stablecoin":`)
	// bacs
	assert.Contains(t, body, `"instrument_id":"src_test_abcdefghijklmnopqr"`)
	assert.Contains(t, body, `"account_number":"12345678"`)
	assert.Contains(t, body, `"bank_code":"200000"`)
	assert.Contains(t, body, `"allow_partial_match":true`)
	assert.Contains(t, body, `"type":"individual"`)
	// card_present
	assert.Contains(t, body, `"track2":"track2data"`)
	assert.Contains(t, body, `"entry_mode":"chip"`)
	assert.Contains(t, body, `"key_set_id":"kset_123"`)
	assert.Contains(t, body, `"block_format":"ISO-0"`)
	// pay_by_bank
	assert.Contains(t, body, `"bank_id":"bank_123"`)
	assert.Contains(t, body, `"type":"select_bank"`)
	assert.Contains(t, body, `"display_name":"Test Bank"`)
	assert.Contains(t, body, `"logo_url":"https://example.com/logo.png"`)
}

func TestBacsAccountHolderType_Values(t *testing.T) {
	assert.Equal(t, BacsAccountHolderType("individual"), BacsAccountHolderIndividual)
	assert.Equal(t, BacsAccountHolderType("corporate"), BacsAccountHolderCorporate)
}

func TestPaymentSetupOrder_AmountAllocations(t *testing.T) {
	order := PaymentSetupOrder{
		AmountAllocations: []PaymentSetupAmountAllocation{
			{
				Id:        "ent_test_abcdefghijklmnopqr",
				Amount:    750,
				Reference: "split-1",
				Commission: &AmountAllocationCommission{
					Amount:     50,
					Percentage: 2.5,
				},
			},
		},
	}

	marshalled, err := json.Marshal(order)
	assert.NoError(t, err)
	body := string(marshalled)
	assert.Contains(t, body, `"amount_allocations":`)
	assert.Contains(t, body, `"id":"ent_test_abcdefghijklmnopqr"`)
	assert.Contains(t, body, `"amount":750`)
	assert.Contains(t, body, `"reference":"split-1"`)
	assert.Contains(t, body, `"commission":`)
	assert.Contains(t, body, `"percentage":2.5`)

	var decoded PaymentSetupOrder
	assert.NoError(t, json.Unmarshal(marshalled, &decoded))
	assert.Len(t, decoded.AmountAllocations, 1)
	assert.Equal(t, int64(750), decoded.AmountAllocations[0].Amount)
	assert.Equal(t, int64(50), decoded.AmountAllocations[0].Commission.Amount)
	assert.Equal(t, 2.5, decoded.AmountAllocations[0].Commission.Percentage)
}

func TestKlarnaAccountHolder_Name(t *testing.T) {
	marshalled, err := json.Marshal(KlarnaAccountHolder{Name: "John Smith"})
	assert.NoError(t, err)
	assert.Contains(t, string(marshalled), `"name":"John Smith"`)

	var decoded KlarnaAccountHolder
	assert.NoError(t, json.Unmarshal([]byte(`{"name":"Jane Doe"}`), &decoded))
	assert.Equal(t, "Jane Doe", decoded.Name)
}

func boolPtr(value bool) *bool {
	return &value
}

// Verifies PaymentSetupAirline and PaymentSetupAccommodation aligned with the 2026-09-08 Checkout.com
// swagger delta: total_number_of_passengers, travel_type, trip_type, refundable, delivery_recipient,
// ancillaries, insurance (airline) and total_number_of_guests, refundable, delivery_recipient, host
// (accommodation). Both schemas are nested under PaymentSetup.industry.airline[] / .accommodation[].
func TestPaymentSetupIndustry_AirlineAllFields(t *testing.T) {
	industry := PaymentSetupIndustry{
		Airline: []PaymentSetupAirline{
			{
				Ticket: &PaymentSetupAirlineTicket{
					Number:                 "0742464639523",
					IssueDate:              shortDate(t, 2025, time.May, 1, 0, 0),
					IssuingCarrierCode:     "042",
					TravelPackageIndicator: "A",
					TravelAgencyName:       "Checkout Travel Agents",
					TravelAgencyCode:       "91114362",
				},
				Passengers: []PaymentSetupAirlinePassenger{
					{
						FirstName:   "John",
						LastName:    "Smith",
						DateOfBirth: shortDate(t, 1990, time.October, 31, 0, 0),
						Address:     &PaymentSetupAirlinePassengerAddress{Country: common.GB},
					},
				},
				FlightLegDetails: []PaymentSetupFlightLegDetails{
					{
						FlightNumber:      "BA1483",
						CarrierCode:       "BA",
						ClassOfTravelling: "W",
						DepartureAirport:  "LHR",
						DepartureDate:     shortDate(t, 2025, time.October, 13, 0, 0),
						DepartureTime:     "18:30",
						ArrivalAirport:    "JFK",
						StopOverCode:      "X",
						FareBasisCode:     "WUP14B",
					},
				},
				TotalNumberOfPassengers: 1,
				TravelType:              "international",
				TripType:                "one_way",
				Refundable:              boolPtr(true),
				DeliveryRecipient:       "jane.smith@example.com",
				Ancillaries:             "extra_baggage",
				Insurance: &PaymentSetupAirlineInsurance{
					Type:    "travel",
					Company: "AXA",
					Price: &PaymentSetupAirlineInsurancePrice{
						Amount:   500,
						Currency: "SAR",
					},
				},
			},
		},
	}

	marshalled, err := json.Marshal(industry)
	assert.NoError(t, err)
	body := string(marshalled)
	assert.Contains(t, body, `"airline":`)
	assert.Contains(t, body, `"issue_date":"2025-05-01"`)
	assert.Contains(t, body, `"date_of_birth":"1990-10-31"`)
	assert.Contains(t, body, `"departure_date":"2025-10-13"`)
	assert.NotContains(t, body, "T00:00:00")
	assert.Contains(t, body, `"issuing_carrier_code":"042"`)
	assert.Contains(t, body, `"class_of_travelling":"W"`)
	assert.Contains(t, body, `"stop_over_code":"X"`)
	assert.Contains(t, body, `"total_number_of_passengers":1`)
	assert.Contains(t, body, `"travel_type":"international"`)
	assert.Contains(t, body, `"trip_type":"one_way"`)
	assert.Contains(t, body, `"refundable":true`)
	assert.Contains(t, body, `"delivery_recipient":"jane.smith@example.com"`)
	assert.Contains(t, body, `"ancillaries":"extra_baggage"`)
	assert.Contains(t, body, `"insurance":`)
	assert.Contains(t, body, `"company":"AXA"`)
	assert.Contains(t, body, `"currency":"SAR"`)

	var decoded PaymentSetupIndustry
	assert.NoError(t, json.Unmarshal(marshalled, &decoded))
	assert.Len(t, decoded.Airline, 1)
	assert.Equal(t, 1, decoded.Airline[0].TotalNumberOfPassengers)
	assert.True(t, *decoded.Airline[0].Refundable)
	assert.Equal(t, "extra_baggage", decoded.Airline[0].Ancillaries)
	assert.Equal(t, float64(500), decoded.Airline[0].Insurance.Price.Amount)
	assert.Equal(t, common.GB, decoded.Airline[0].Passengers[0].Address.Country)
}

func TestPaymentSetupIndustry_AccommodationAllFields(t *testing.T) {
	industry := PaymentSetupIndustry{
		Accommodation: []PaymentSetupAccommodation{
			{
				Name:             "Checkout Lodge",
				BookingReference: "REF9083748",
				CheckInDate:      shortDate(t, 2025, time.April, 11, 0, 0),
				CheckOutDate:     shortDate(t, 2025, time.April, 18, 0, 0),
				Address: &common.Address{
					AddressLine1: "123 High Street",
					City:         "London",
					State:        "Greater London",
					Country:      common.GB,
					Zip:          "NE1 1CK",
				},
				NumberOfRooms: 2,
				Guests: []PaymentSetupAccommodationGuest{
					{FirstName: "John", LastName: "Smith", DateOfBirth: shortDate(t, 1970, time.March, 19, 0, 0)},
				},
				Room: []PaymentSetupAccommodationRoom{
					{Rate: 42.3, NumberOfNights: 5, Type: "deluxe"},
				},
				TotalNumberOfGuests: 2,
				Refundable:          boolPtr(true),
				DeliveryRecipient:   "jane.smith@example.com",
				Host: &PaymentSetupAccommodationHost{
					RegistrationDate:      shortDate(t, 2020, time.January, 1, 0, 0),
					TotalReservationCount: 150,
				},
			},
		},
	}

	marshalled, err := json.Marshal(industry)
	assert.NoError(t, err)
	body := string(marshalled)
	assert.Contains(t, body, `"accommodation":`)
	assert.Contains(t, body, `"check_in_date":"2025-04-11"`)
	assert.Contains(t, body, `"check_out_date":"2025-04-18"`)
	assert.Contains(t, body, `"date_of_birth":"1970-03-19"`)
	assert.Contains(t, body, `"registration_date":"2020-01-01"`)
	assert.NotContains(t, body, "T00:00:00")
	assert.Contains(t, body, `"total_number_of_guests":2`)
	assert.Contains(t, body, `"refundable":true`)
	assert.Contains(t, body, `"delivery_recipient":"jane.smith@example.com"`)
	assert.Contains(t, body, `"host":`)
	assert.Contains(t, body, `"total_reservation_count":150`)

	var decoded PaymentSetupIndustry
	assert.NoError(t, json.Unmarshal(marshalled, &decoded))
	assert.Len(t, decoded.Accommodation, 1)
	assert.Equal(t, 2, decoded.Accommodation[0].TotalNumberOfGuests)
	assert.True(t, *decoded.Accommodation[0].Refundable)
	assert.Equal(t, "jane.smith@example.com", decoded.Accommodation[0].DeliveryRecipient)
	assert.Equal(t, 150, decoded.Accommodation[0].Host.TotalReservationCount)
}

func timePtr(t *testing.T, value string) *time.Time {
	parsed, err := time.Parse(time.RFC3339, value)
	assert.NoError(t, err)
	return &parsed
}
