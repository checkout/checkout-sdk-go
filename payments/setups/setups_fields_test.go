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
				NumberOfRooms: intPtr(2),
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

// NumberOfRooms is a *int precisely so these two cases differ. With a plain int and omitempty an
// explicit zero was dropped, and the swagger declares the field integer with no minimum.
func TestPaymentSetupAccommodation_ExplicitZeroRoomsSurvives(t *testing.T) {
	raw, err := json.Marshal(PaymentSetupAccommodation{
		Name:          "Checkout Lodge",
		NumberOfRooms: intPtr(0),
	})
	assert.Nil(t, err)

	var body map[string]interface{}
	assert.Nil(t, json.Unmarshal(raw, &body))

	value, present := body["number_of_rooms"]
	assert.True(t, present, "an explicit zero must survive serialization")
	assert.Equal(t, float64(0), value)
}

func TestPaymentSetupAccommodation_OmitsRoomsWhenUnset(t *testing.T) {
	raw, err := json.Marshal(PaymentSetupAccommodation{Name: "Checkout Lodge"})
	assert.Nil(t, err)

	var body map[string]interface{}
	assert.Nil(t, json.Unmarshal(raw, &body))

	_, present := body["number_of_rooms"]
	assert.False(t, present)
}

// Cash App Pay on Payment Setups: payment_methods.cashapp (schema CashApp) and the customer.device
// fields fingerprint, ipv4, ipv6, client and os. Fixture values are the public swagger examples.

const cashAppRedirectUrl = "https://sandbox.api.cash.app/customer-request/v1/requests/GRR_f5xg6wrxhtv3p4w24g0wrexa/interstitial?validity_token=bap03y"

func TestPaymentSetupRequest_CashAppWireKeyAndMerchantFields(t *testing.T) {
	request := PaymentSetupRequest{
		ProcessingChannelId: "pc_aaaaaaaaaaaaaaaaaaaaaaaaaa",
		Amount:              1000,
		Currency:            common.USD,
		PaymentMethods: &PaymentMethods{
			CashApp: &CashAppPaymentMethod{
				PaymentMethodBase:      PaymentMethodBase{Initialization: PaymentMethodInitializationEnabled},
				CustomerProfileSharing: boolPtr(true),
			},
		},
		Customer: &PaymentSetupCustomer{
			Device: &PaymentSetupCustomerDevice{
				Locale:      "en_US",
				Fingerprint: "fp_abc123xyz",
				Ipv4:        "203.0.113.0",
				Ipv6:        "2001:db8:85a3::8a2e:370:7334",
				Client:      PaymentSetupDeviceClientWeb,
				Os:          PaymentSetupDeviceOsAndroid,
			},
		},
	}

	marshalled, err := json.Marshal(request)
	assert.NoError(t, err)
	body := string(marshalled)
	assert.NotContains(t, body, "cash_app")
	assert.NotContains(t, body, "cashApp")
	assert.NotContains(t, body, "customerProfileSharing")

	var decoded map[string]interface{}
	assert.NoError(t, json.Unmarshal(marshalled, &decoded))
	methods := decoded["payment_methods"].(map[string]interface{})
	cashApp, present := methods["cashapp"].(map[string]interface{})
	assert.True(t, present, "payment_methods must carry the literal key cashapp")
	assert.Equal(t, map[string]interface{}{
		"initialization":           "enabled",
		"customer_profile_sharing": true,
	}, cashApp)

	device := decoded["customer"].(map[string]interface{})["device"]
	assert.Equal(t, map[string]interface{}{
		"locale":      "en_US",
		"fingerprint": "fp_abc123xyz",
		"ipv4":        "203.0.113.0",
		"ipv6":        "2001:db8:85a3::8a2e:370:7334",
		"client":      "web",
		"os":          "android",
	}, device)
}

// CustomerProfileSharing is a *bool so an explicit false reaches the API instead of being dropped.
func TestCashAppPaymentMethod_CustomerProfileSharingFalseIsSent(t *testing.T) {
	marshalled, err := json.Marshal(CashAppPaymentMethod{CustomerProfileSharing: boolPtr(false)})
	assert.NoError(t, err)
	assert.JSONEq(t, `{"customer_profile_sharing":false}`, string(marshalled))

	marshalled, err = json.Marshal(CashAppPaymentMethod{})
	assert.NoError(t, err)
	assert.JSONEq(t, `{}`, string(marshalled))
}

func TestPaymentSetupCustomerDevice_ClientAndOsValues(t *testing.T) {
	cases := []struct {
		device   PaymentSetupCustomerDevice
		expected string
	}{
		{PaymentSetupCustomerDevice{Client: PaymentSetupDeviceClientWeb}, `{"client":"web"}`},
		{PaymentSetupCustomerDevice{Client: PaymentSetupDeviceClientMobileWeb}, `{"client":"mobile_web"}`},
		{PaymentSetupCustomerDevice{Client: PaymentSetupDeviceClientApp}, `{"client":"app"}`},
		{PaymentSetupCustomerDevice{Os: PaymentSetupDeviceOsAndroid}, `{"os":"android"}`},
		{PaymentSetupCustomerDevice{Os: PaymentSetupDeviceOsIos}, `{"os":"ios"}`},
	}

	for _, tc := range cases {
		t.Run(tc.expected, func(t *testing.T) {
			marshalled, err := json.Marshal(tc.device)
			assert.NoError(t, err)
			assert.JSONEq(t, tc.expected, string(marshalled))
		})
	}
}

func TestPaymentSetupCustomerDevice_OnlyLocaleSerializesOnlyLocale(t *testing.T) {
	marshalled, err := json.Marshal(PaymentSetupCustomerDevice{Locale: "en_US"})
	assert.NoError(t, err)
	assert.Equal(t, `{"locale":"en_US"}`, string(marshalled))
}

// GetPaymentSetup and ConfirmPaymentSetup both return PaymentSetupResponse, so this read covers
// the get, create, update and confirm responses.
func TestPaymentSetupResponse_CashAppSwaggerExample(t *testing.T) {
	payload := `{
		"id": "ps_test_abcdefghijklmnopqr",
		"processing_channel_id": "pc_aaaaaaaaaaaaaaaaaaaaaaaaaa",
		"amount": 1000,
		"currency": "USD",
		"available_payment_methods": ["cashapp"],
		"customer": {
			"device": {
				"locale": "en_US",
				"fingerprint": "fp_abc123xyz",
				"ipv4": "203.0.113.0",
				"ipv6": "2001:db8:85a3::8a2e:370:7334",
				"client": "web",
				"os": "android"
			}
		},
		"payment_methods": {
			"cashapp": {
				"status": "action_required",
				"flags": [],
				"initialization": "enabled",
				"customer_profile_sharing": true,
				"reference": "ORDER-99",
				"action": {
					"type": "redirect",
					"redirect_url": "` + cashAppRedirectUrl + `"
				},
				"customer_profile": {
					"customer_id": "CST_AYVkuLzfsRqEhf4OyQFxQNv22m7IjNFjO6f2J5CDE2nxAC4-21wJ2H8_2kvsdIsDZMN4",
					"cashtag": "$CASHTAG_C_TOKEN",
					"reference_id": "value",
					"full_name": "John Middle Doe",
					"given_name": "John",
					"middle_name": "Middle",
					"family_name": "Doe",
					"suffix": "Jr.",
					"birth_date": "1990-01-01T00:00:00.0000000",
					"address": {
						"address_line_1": "123 Main St",
						"address_line_2": "Apt 2",
						"address_line_3": "Floor 3",
						"locality": "Springfield",
						"sublocality": "Downtown",
						"administrative_district_level_1": "IL",
						"postal_code": "62701",
						"country": "US"
					},
					"phone_number": "5555555555",
					"email_address": "cash@cash.com",
					"customer_since": "1970-01-18T12:46:04.8000000+00:00"
				}
			}
		}
	}`

	var response PaymentSetupResponse
	assert.NoError(t, json.Unmarshal([]byte(payload), &response))
	assert.Equal(t, []string{"cashapp"}, response.AvailablePaymentMethods)

	device := response.Customer.Device
	assert.Equal(t, "en_US", device.Locale)
	assert.Equal(t, "fp_abc123xyz", device.Fingerprint)
	assert.Equal(t, "203.0.113.0", device.Ipv4)
	assert.Equal(t, "2001:db8:85a3::8a2e:370:7334", device.Ipv6)
	assert.Equal(t, PaymentSetupDeviceClientWeb, device.Client)
	assert.Equal(t, PaymentSetupDeviceOsAndroid, device.Os)

	cashApp := response.PaymentMethods.CashApp
	assert.NotNil(t, cashApp)
	assert.Equal(t, "action_required", cashApp.Status)
	assert.NotNil(t, cashApp.Flags)
	assert.Empty(t, cashApp.Flags)
	assert.Equal(t, PaymentMethodInitializationEnabled, cashApp.Initialization)
	assert.True(t, *cashApp.CustomerProfileSharing)
	assert.Equal(t, "ORDER-99", cashApp.Reference)
	assert.Equal(t, "redirect", cashApp.Action.Type)
	assert.Equal(t, cashAppRedirectUrl, cashApp.Action.RedirectUrl)

	profile := cashApp.CustomerProfile
	assert.Equal(t, "CST_AYVkuLzfsRqEhf4OyQFxQNv22m7IjNFjO6f2J5CDE2nxAC4-21wJ2H8_2kvsdIsDZMN4", profile.CustomerId)
	assert.Equal(t, "$CASHTAG_C_TOKEN", profile.Cashtag)
	assert.Equal(t, "value", profile.ReferenceId)
	assert.Equal(t, "John Middle Doe", profile.FullName)
	assert.Equal(t, "John", profile.GivenName)
	assert.Equal(t, "Middle", profile.MiddleName)
	assert.Equal(t, "Doe", profile.FamilyName)
	assert.Equal(t, "Jr.", profile.Suffix)
	assert.Equal(t, "1990-01-01T00:00:00.0000000", profile.BirthDate)
	assert.Equal(t, "5555555555", profile.PhoneNumber)
	assert.Equal(t, "cash@cash.com", profile.EmailAddress)
	assert.Equal(t, "1970-01-18T12:46:04.8000000+00:00", profile.CustomerSince)

	address := profile.Address
	assert.Equal(t, "123 Main St", address.AddressLine1)
	assert.Equal(t, "Apt 2", address.AddressLine2)
	assert.Equal(t, "Floor 3", address.AddressLine3)
	assert.Equal(t, "Springfield", address.Locality)
	assert.Equal(t, "Downtown", address.Sublocality)
	assert.Equal(t, "IL", address.AdministrativeDistrictLevel1)
	assert.Equal(t, "62701", address.PostalCode)
	assert.Equal(t, common.US, address.Country)
}

func TestCashAppPaymentMethod_RoundTripKeepsEveryProperty(t *testing.T) {
	original := CashAppPaymentMethod{
		PaymentMethodBase: PaymentMethodBase{
			Status:         "action_required",
			Flags:          []string{"example_flag"},
			Initialization: PaymentMethodInitializationEnabled,
		},
		CustomerProfileSharing: boolPtr(true),
		Reference:              "ORDER-99",
		Action: &CashAppAction{
			Type:        "redirect",
			RedirectUrl: cashAppRedirectUrl,
		},
		CustomerProfile: &CashAppCustomerProfile{
			CustomerId:  "CST_AYVkuLzfsRqEhf4OyQFxQNv22m7IjNFjO6f2J5CDE2nxAC4-21wJ2H8_2kvsdIsDZMN4",
			Cashtag:     "$CASHTAG_C_TOKEN",
			ReferenceId: "value",
			FullName:    "John Middle Doe",
			GivenName:   "John",
			MiddleName:  "Middle",
			FamilyName:  "Doe",
			Suffix:      "Jr.",
			BirthDate:   "1990-01-01T00:00:00.0000000",
			Address: &CashAppAddress{
				AddressLine1:                 "123 Main St",
				AddressLine2:                 "Apt 2",
				AddressLine3:                 "Floor 3",
				Locality:                     "Springfield",
				Sublocality:                  "Downtown",
				AdministrativeDistrictLevel1: "IL",
				PostalCode:                   "62701",
				Country:                      common.US,
			},
			PhoneNumber:   "5555555555",
			EmailAddress:  "cash@cash.com",
			CustomerSince: "1970-01-18T12:46:04.8000000+00:00",
		},
	}

	marshalled, err := json.Marshal(original)
	assert.NoError(t, err)
	body := string(marshalled)
	assert.Contains(t, body, `"address_line_1":"123 Main St"`)
	assert.Contains(t, body, `"address_line_2":"Apt 2"`)
	assert.Contains(t, body, `"address_line_3":"Floor 3"`)
	assert.Contains(t, body, `"administrative_district_level_1":"IL"`)
	assert.NotContains(t, body, "address_line1")
	assert.NotContains(t, body, "administrative_district_level1")
	assert.Contains(t, body, `"redirect_url":"`)
	assert.Contains(t, body, `"customer_since":"1970-01-18T12:46:04.8000000+00:00"`)

	var decoded CashAppPaymentMethod
	assert.NoError(t, json.Unmarshal(marshalled, &decoded))
	assert.Equal(t, original, decoded)
}

// Every PaymentSetup.customer property, with the swagger examples.
func TestPaymentSetupCustomer_AllPropertiesRoundTripAndSwaggerExample(t *testing.T) {
	original := PaymentSetupCustomer{
		Id:      "cus_123456789",
		Country: common.GB,
		Email: &PaymentSetupCustomerEmail{
			Address:  "johnsmith@example.com",
			Verified: boolPtr(true),
		},
		Name:      "John Smith",
		TaxNumber: "GB123456789",
		Phone:     &common.Phone{CountryCode: "+44", Number: "207 946 0000"},
		Device: &PaymentSetupCustomerDevice{
			Locale:      "en_GB",
			Fingerprint: "fp_abc123xyz",
			Ipv4:        "203.0.113.0",
			Ipv6:        "2001:db8:85a3::8a2e:370:7334",
			Client:      PaymentSetupDeviceClientApp,
			Os:          PaymentSetupDeviceOsIos,
		},
		MerchantAccount: &CustomerMerchantAccount{
			Id:                   "acc_123",
			RegistrationDate:     shortDate(t, 2020, time.January, 1, 0, 0),
			LastModified:         shortDate(t, 2021, time.February, 2, 0, 0),
			ReturningCustomer:    boolPtr(true),
			FirstTransactionDate: shortDate(t, 2020, time.March, 3, 0, 0),
			LastTransactionDate:  shortDate(t, 2022, time.April, 4, 0, 0),
			TotalOrderCount:      7,
			LastPaymentAmount:    1500,
		},
	}

	marshalled, err := json.Marshal(original)
	assert.NoError(t, err)
	body := string(marshalled)
	assert.Contains(t, body, `"id":"cus_123456789"`)
	assert.Contains(t, body, `"country":"GB"`)
	assert.Contains(t, body, `"tax_number":"GB123456789"`)
	assert.NotContains(t, body, "taxNumber")

	var decoded PaymentSetupCustomer
	assert.NoError(t, json.Unmarshal(marshalled, &decoded))
	assert.Equal(t, original, decoded)

	example := `{
		"country": "GB",
		"id": "cus_123456789",
		"email": {"address": "johnsmith@example.com", "verified": true},
		"name": "John Smith",
		"tax_number": "GB123456789",
		"phone": {"country_code": "+44", "number": "207 946 0000"},
		"device": {"locale": "en_GB", "fingerprint": "fp_abc123xyz", "ipv4": "203.0.113.0",
			"ipv6": "2001:db8:85a3::8a2e:370:7334", "client": "web", "os": "android"}
	}`
	var fromExample PaymentSetupCustomer
	assert.NoError(t, json.Unmarshal([]byte(example), &fromExample))
	assert.Equal(t, "cus_123456789", fromExample.Id)
	assert.Equal(t, common.GB, fromExample.Country)
	assert.Equal(t, "johnsmith@example.com", fromExample.Email.Address)
	assert.True(t, *fromExample.Email.Verified)
	assert.Equal(t, "John Smith", fromExample.Name)
	assert.Equal(t, "GB123456789", fromExample.TaxNumber)
	assert.Equal(t, "+44", fromExample.Phone.CountryCode)
	assert.Equal(t, "207 946 0000", fromExample.Phone.Number)
	assert.Equal(t, "en_GB", fromExample.Device.Locale)
	assert.Equal(t, PaymentSetupDeviceClientWeb, fromExample.Device.Client)
	assert.Equal(t, PaymentSetupDeviceOsAndroid, fromExample.Device.Os)
}
