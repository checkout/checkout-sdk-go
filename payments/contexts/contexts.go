package contexts

import (
	"time"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/payments"

	"github.com/checkout/checkout-sdk-go/payments/nas"
	"github.com/checkout/checkout-sdk-go/payments/nas/sources/apm"
)

const PaymentContextsPath = "payment-contexts"

type (
	PaymentContextsRequest struct {
		Source              payments.PaymentSource     `json:"source,omitempty"`
		Amount              int64                      `json:"amount,omitempty"`
		Currency            common.Currency            `json:"currency,omitempty"`
		PaymentType         payments.PaymentType       `json:"payment_type,omitempty"`
		Capture             bool                       `json:"capture,omitempty"`
		Shipping            *payments.ShippingDetails  `json:"shipping,omitempty"`
		Processing          *PaymentContextsProcessing `json:"processing,omitempty"`
		ProcessingChannelId string                     `json:"processing_channel_id,omitempty"`
		Reference           string                     `json:"reference,omitempty"`
		Description         string                     `json:"description,omitempty"`
		SuccessUrl          string                     `json:"success_url,omitempty"`
		FailureUrl          string                     `json:"failure_url,omitempty"`
		Items               []PaymentContextsItems     `json:"items,omitempty"`
	}
)

type (
	PaymentContextsPartnerMetadata struct {
		OrderId    string `json:"order_id,omitempty"`
		CustomerId string `json:"customer_id,omitempty"`
	}

	PaymentContextsPartnerCustomerRiskData struct {
		Key   string `json:"key,omitempty"`
		Value string `json:"value,omitempty"`
	}

	PaymentContextsTicket struct {
		Number                 string     `json:"number,omitempty"`
		IssueDate              *time.Time `json:"issue_date,omitempty"`
		IssuingCarrierCode     string     `json:"issuing_carrier_code,omitempty"`
		TravelPackageIndicator string     `json:"travel_package_indicator,omitempty"`
		TravelAgencyName       string     `json:"travel_agency_name,omitempty"`
		TravelAgencyCode       string     `json:"travel_agency_code,omitempty"`
	}

	PaymentContextsPassenger struct {
		FirstName   string          `json:"first_name,omitempty"`
		LastName    string          `json:"last_name,omitempty"`
		DateOfBirth *time.Time      `json:"date_of_birth,omitempty"`
		Address     *common.Address `json:"address,omitempty"`
	}

	PaymentContextsFlightLegDetails struct {
		FlightNumber      string     `json:"flight_number,omitempty"`
		CarrierCode       string     `json:"carrier_code,omitempty"`
		ClassOfTravelling string     `json:"class_of_travelling,omitempty"`
		DepartureAirport  string     `json:"departure_airport,omitempty"`
		DepartureDate     *time.Time `json:"departure_date,omitempty"`
		DepartureTime     string     `json:"departure_time,omitempty"`
		ArrivalAirport    string     `json:"arrival_airport,omitempty"`
		StopOverCode      string     `json:"stop_over_code,omitempty"`
		FareBasisCode     string     `json:"fare_basis_code,omitempty"`
	}

	PaymentContextsAirlineData struct {
		Ticket           []PaymentContextsTicket           `json:"ticket,omitempty"`
		Passenger        []PaymentContextsPassenger        `json:"passenger,omitempty"`
		FlightLegDetails []PaymentContextsFlightLegDetails `json:"flight_leg_details,omitempty"`
	}

	PaymentContextsProcessing struct {
		Plan                    *apm.BillingPlan                        `json:"plan,omitempty"`
		ShippingAmount          int                                     `json:"shipping_amount,omitempty"`
		InvoiceId               string                                  `json:"invoice_id,omitempty"`
		BrandName               string                                  `json:"brand_name,omitempty"`
		Locale                  string                                  `json:"locale,omitempty"`
		ShippingPreference      payments.ShippingPreference             `json:"shipping_preference,omitempty"`
		UserAction              payments.UserAction                     `json:"user_action,omitempty"`
		PartnerCustomerRiskData *PaymentContextsPartnerCustomerRiskData `json:"partner_customer_risk_data,omitempty"`
		AirlineData             []PaymentContextsAirlineData            `json:"airline_data,omitempty"`
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
		PaymentRequest  *PaymentContextsResponse        `json:"payment_request,omitempty"`
		PartnerMetadata *PaymentContextsPartnerMetadata `json:"partner_metadata,omitempty"`
		Customer        map[string]interface{}          `json:"customer,omitempty"`
	}
)
