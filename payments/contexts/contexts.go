package contexts

import (
	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/payments"

	"github.com/checkout/checkout-sdk-go/v3/payments/nas"
	"github.com/checkout/checkout-sdk-go/v3/payments/nas/sources/apm"
)

type PaymentContextDetailsStatusType string

const (
	Created  PaymentContextDetailsStatusType = "Created"
	Approved PaymentContextDetailsStatusType = "Approved"
)

const PaymentContextsPath = "payment-contexts"

type (
	PaymentContextsCustomerSummary struct {
		// RegistrationDate is the date the customer registered.
		// [Optional]
		// Format: yyyy-MM-dd
		RegistrationDate *common.APIShortDate `json:"registration_date,omitempty"`

		// FirstTransactionDate is the date of the customer's first transaction.
		// [Optional]
		// Format: yyyy-MM-dd
		FirstTransactionDate *common.APIShortDate `json:"first_transaction_date,omitempty"`

		// LastPaymentDate is the date of the customer's last payment.
		// [Optional]
		// Format: yyyy-MM-dd
		LastPaymentDate *common.APIShortDate `json:"last_payment_date,omitempty"`

		TotalOrderCount     int64   `json:"total_order_count,omitempty"`
		LastPaymentAmount   float64 `json:"last_payment_amount,omitempty"`
		IsPremiumCustomer   bool    `json:"is_premium_customer,omitempty"`
		IsReturningCustomer bool    `json:"is_returning_customer,omitempty"`
		LifetimeValue       float64 `json:"lifetime_value,omitempty"`
	}

	PaymentContextCustomerRequest struct {
		EmailVerified bool                            `json:"email_verified,omitempty"`
		Email         string                          `json:"email,omitempty"`
		Name          string                          `json:"name,omitempty"`
		Phone         *common.Phone                   `json:"phone,omitempty"`
		Summary       *PaymentContextsCustomerSummary `json:"summary,omitempty"`
	}

	PaymentContextsRequest struct {
		Source              payments.PaymentSource         `json:"source,omitempty"`
		Amount              int64                          `json:"amount,omitempty"`
		Currency            common.Currency                `json:"currency,omitempty"`
		PaymentType         payments.PaymentType           `json:"payment_type,omitempty"`
		AuthorizationType   string                         `json:"authorization_type,omitempty"`
		Capture             bool                           `json:"capture,omitempty"`
		Customer            *PaymentContextCustomerRequest `json:"customer,omitempty"`
		Shipping            *payments.ShippingDetails      `json:"shipping,omitempty"`
		Processing          *PaymentContextsProcessing     `json:"processing,omitempty"`
		ProcessingChannelId string                         `json:"processing_channel_id,omitempty"`
		Reference           string                         `json:"reference,omitempty"`
		Description         string                         `json:"description,omitempty"`
		SuccessUrl          string                         `json:"success_url,omitempty"`
		FailureUrl          string                         `json:"failure_url,omitempty"`
		Items               []PaymentContextsItems         `json:"items,omitempty"`
		Metadata            map[string]interface{}         `json:"metadata,omitempty"`
	}
)

type (
	PaymentContextsPartnerMetadata struct {
		OrderId     string `json:"order_id,omitempty"`
		CustomerId  string `json:"customer_id,omitempty"`
		SessionId   string `json:"session_id,omitempty"`
		ClientToken string `json:"client_token,omitempty"`
	}

	// PaymentContextsPartnerCustomerRiskData is a key-and-value pair with merchant-specific
	// data for the transaction.
	//
	// Deprecated: duplicates payments.PartnerCustomerRiskData, which maps the same
	// specification shape and is what PaymentContextsProcessing now uses. Retained for
	// backwards compatibility and will be removed in a future version.
	PaymentContextsPartnerCustomerRiskData struct {
		Key   string `json:"key,omitempty"`
		Value string `json:"value,omitempty"`
	}

	// PaymentContextsTicket contains information about the airline ticket.
	PaymentContextsTicket struct {
		// Number is the ticket's unique identifier.
		// [Optional]
		Number string `json:"number,omitempty"`

		// IssueDate is the date the airline ticket was issued.
		// [Optional]
		// Format: yyyy-MM-dd
		IssueDate *common.APIShortDate `json:"issue_date,omitempty"`

		// IssuingCarrierCode is the carrier code of the ticket issuer.
		// [Optional]
		IssuingCarrierCode string `json:"issuing_carrier_code,omitempty"`

		// TravelPackageIndicator is C = Car rental reservation, A = Airline flight reservation,
		// B = Both car rental and airline flight reservations included, N = Unknown.
		// [Optional]
		TravelPackageIndicator string `json:"travel_package_indicator,omitempty"`

		// TravelAgencyName is the name of the travel agency.
		// [Optional]
		TravelAgencyName string `json:"travel_agency_name,omitempty"`

		// TravelAgencyCode is the unique identifier from IATA or ARC for the travel agency that
		// issues the ticket.
		// [Optional]
		TravelAgencyCode string `json:"travel_agency_code,omitempty"`
	}

	// PaymentContextsPassenger contains information about a passenger on the flight.
	PaymentContextsPassenger struct {
		// FirstName is the passenger's first name.
		// [Optional]
		FirstName string `json:"first_name,omitempty"`

		// LastName is the passenger's last name.
		// [Optional]
		LastName string `json:"last_name,omitempty"`

		// DateOfBirth is the passenger's date of birth.
		// [Optional]
		// Format: yyyy-MM-dd
		DateOfBirth *common.APIShortDate `json:"date_of_birth,omitempty"`

		// Address contains information about the passenger's address.
		// [Optional]
		//
		// The specification defines exactly one property on this object, country. This shares
		// payments.PassengerAddress so that one specification shape maps to one type.
		Address *payments.PassengerAddress `json:"address,omitempty"`
	}

	// PaymentContextsFlightLegDetails contains information about a flight leg booked by the
	// customer.
	PaymentContextsFlightLegDetails struct {
		// FlightNumber is the flight identifier.
		// [Optional]
		FlightNumber string `json:"flight_number,omitempty"`

		// CarrierCode is the IATA 2-letter accounting code (PAX) that identifies the carrier.
		// This field is required if the airline data includes leg details.
		// [Optional]
		CarrierCode string `json:"carrier_code,omitempty"`

		// ClassOfTravelling is a one-letter travel class identifier. The following are common:
		// F = First class, J = Business class, Y = Economy class, W = Premium economy.
		// [Optional]
		ClassOfTravelling string `json:"class_of_travelling,omitempty"`

		// DepartureAirport is the IATA three-letter airport code of the departure airport.
		// This field is required if the airline data includes leg details.
		// [Optional]
		DepartureAirport string `json:"departure_airport,omitempty"`

		// DepartureDate is the date of the scheduled take off.
		// [Optional]
		// Format: yyyy-MM-dd
		DepartureDate *common.APIShortDate `json:"departure_date,omitempty"`

		// DepartureTime is the time of the scheduled take off.
		// [Optional]
		DepartureTime string `json:"departure_time,omitempty"`

		// ArrivalAirport is the IATA 3-letter airport code of the destination airport.
		// This field is required if the airline data includes leg details.
		// [Optional]
		ArrivalAirport string `json:"arrival_airport,omitempty"`

		// StopOverCode is a one-letter code that indicates whether the passenger is entitled to
		// make a stopover. Can be a space, O if the passenger is entitled to make a stopover, or
		// X if they are not.
		// [Optional]
		StopOverCode string `json:"stop_over_code,omitempty"`

		// FareBasisCode is the fare basis code, alphanumeric.
		// [Optional]
		FareBasisCode string `json:"fare_basis_code,omitempty"`
	}

	// PaymentContextsAirlineData contains information about the airline ticket and flights
	// booked by the customer.
	//
	// Passenger accepts both wire shapes; see PaymentContextsAirlineData.UnmarshalJSON.
	PaymentContextsAirlineData struct {
		// Ticket contains information about the airline ticket.
		// [Optional]
		//
		// The specification declares this as a single object. It was previously a slice, so the
		// SDK sent ticket as an array, a shape the API does not accept.
		Ticket *PaymentContextsTicket `json:"ticket,omitempty"`

		// Passenger contains information about the passenger(s) on the flight.
		// [Optional]
		//
		// Deserialization also accepts a single object, which PayPal returns in place of an
		// array; it is normalized to a one-element slice. Marshaling always emits an array.
		Passenger []PaymentContextsPassenger `json:"passenger,omitempty"`

		// FlightLegDetails contains information about the flight leg(s) booked by the customer.
		// [Optional]
		FlightLegDetails []PaymentContextsFlightLegDetails `json:"flight_leg_details,omitempty"`
	}

	// PaymentContextsProcessing holds settings that control how the payment context is
	// processed.
	PaymentContextsProcessing struct {
		// Plan is the plan details for a recurring payment with PayPal. Required when
		// payment_type is recurring.
		// [Optional]
		Plan *apm.BillingPlan `json:"plan,omitempty"`

		// DiscountAmount is the discount amount the merchant applied to the transaction.
		// [Optional]
		DiscountAmount int `json:"discount_amount,omitempty"`

		// ShippingAmount is the total freight or shipping and handling charges for the
		// transaction.
		// [Optional]
		ShippingAmount int `json:"shipping_amount,omitempty"`

		// TaxAmount is the total tax amount for the transaction, in the minor currency unit.
		// [Optional]
		TaxAmount int `json:"tax_amount,omitempty"`

		// InvoiceId is the invoice ID number.
		// [Optional]
		InvoiceId string `json:"invoice_id,omitempty"`

		// BrandName is the label that overrides the business name in the PayPal account on the
		// PayPal pages.
		// [Optional]
		BrandName string `json:"brand_name,omitempty"`

		// Locale is the language and region of the customer in ISO 639-2 language code; the
		// value consists of language-country.
		// [Optional]
		Locale string `json:"locale,omitempty"`

		// ShippingPreference is the shipping preference.
		// [Optional]
		// One of: no_shipping, set_provided_address, get_from_file
		ShippingPreference payments.ShippingPreference `json:"shipping_preference,omitempty"`

		// UserAction is a property required by PayPal to have an appropriate payment flow.
		// [Optional]
		// One of: pay_now, continue
		UserAction payments.UserAction `json:"user_action,omitempty"`

		// PartnerCustomerRiskData holds key-and-value pairs with merchant-specific data for
		// the transaction.
		// [Optional]
		//
		// Uses the shared payments.PartnerCustomerRiskData: payment contexts and the
		// payments request schemas resolve partner_customer_risk_data to the same
		// specification shape, and maintaining two identical structs for it invited drift.
		PartnerCustomerRiskData []payments.PartnerCustomerRiskData `json:"partner_customer_risk_data,omitempty"`

		// CustomPaymentMethodIds are promo codes. They define which of the configured payment
		// options within a payment category (pay_later, pay_over_time, and so on) are shown for
		// this purchase.
		// [Optional]
		CustomPaymentMethodIds []string `json:"custom_payment_method_ids,omitempty"`

		// AirlineData contains information about the airline ticket and flights booked by the
		// customer.
		// [Optional]
		AirlineData []PaymentContextsAirlineData `json:"airline_data,omitempty"`

		// AccommodationData contains information about the accommodation booked by the customer.
		// [Optional]
		//
		// Uses the shared payments.AccommodationData, because payment contexts, POST /payments
		// and the GET /payments/{id} response all resolve accommodation_data to the same
		// specification schema.
		AccommodationData []payments.AccommodationData `json:"accommodation_data,omitempty"`
	}

	PaymentContextsItems struct {
		Name           string `json:"name,omitempty"`
		Quantity       int    `json:"quantity,omitempty"`
		UnitPrice      int    `json:"unit_price,omitempty"`
		Reference      string `json:"reference,omitempty"`
		TotalAmount    int    `json:"total_amount,omitempty"`
		TaxAmount      int    `json:"tax_amount,omitempty"`
		DiscountAmount int    `json:"discount_amount,omitempty"`
		Url            string `json:"url,omitempty"`
		ImageUrl       string `json:"image_url,omitempty"`
	}

	PaymentContextsResponse struct {
		Source              *nas.SourceResponse        `json:"source,omitempty"`
		Amount              int64                      `json:"amount,omitempty"`
		Currency            common.Currency            `json:"currency,omitempty"`
		PaymentType         payments.PaymentType       `json:"payment_type,omitempty"`
		Capture             bool                       `json:"capture,omitempty"`
		Shipping            *payments.ShippingDetails  `json:"shipping"`
		Processing          *PaymentContextsProcessing `json:"processing"`
		ProcessingChannelId string                     `json:"processing_channel_id,omitempty"`
		Reference           string                     `json:"reference,omitempty"`
		Description         string                     `json:"description,omitempty"`
		SuccessUrl          string                     `json:"success_url,omitempty"`
		FailureUrl          string                     `json:"failure_url,omitempty"`
		Items               []PaymentContextsItems     `json:"items,omitempty"`
	}

	PaymentContextsRequestResponse struct {
		HttpMetadata    common.HttpMetadata
		Id              string                          `json:"id,omitempty"`
		PartnerMetadata *PaymentContextsPartnerMetadata `json:"partner_metadata,omitempty"`
		Links           map[string]common.Link          `json:"links,omitempty"`
	}

	PaymentContextDetailsResponse struct {
		HttpMetadata    common.HttpMetadata
		Id              string                          `json:"id,omitempty"`
		Status          PaymentContextDetailsStatusType `json:"status,omitempty"`
		PaymentRequest  *PaymentContextsResponse        `json:"payment_request,omitempty"`
		PartnerMetadata *PaymentContextsPartnerMetadata `json:"partner_metadata,omitempty"`
	}
)
