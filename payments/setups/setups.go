package setups

import (
	"time"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/payments"
)

const (
	PaymentSetupsPath = "payments/setups"
	ConfirmPath       = "confirm"
)

// Payment Setup Request
type PaymentSetupRequest struct {
	ProcessingChannelId string                `json:"processing_channel_id,omitempty"`
	Amount              int64                 `json:"amount"`
	Currency            common.Currency       `json:"currency,omitempty"`
	PaymentType         payments.PaymentType  `json:"payment_type,omitempty"`
	Reference           string                `json:"reference,omitempty"`
	Description         string                `json:"description,omitempty"`
	PaymentMethods      *PaymentMethods       `json:"payment_methods,omitempty"`
	Settings            *PaymentSetupSettings `json:"settings,omitempty"`
	Customer            *PaymentSetupCustomer `json:"customer,omitempty"`
	Order               *PaymentSetupOrder    `json:"order,omitempty"`
	Industry            *PaymentSetupIndustry `json:"industry,omitempty"`
}

// Payment Setup Response
type PaymentSetupResponse struct {
	Id                  string                `json:"id,omitempty"`
	ProcessingChannelId string                `json:"processing_channel_id,omitempty"`
	Amount              int64                 `json:"amount,omitempty"`
	Currency            common.Currency       `json:"currency,omitempty"`
	PaymentType         payments.PaymentType  `json:"payment_type,omitempty"`
	Reference           string                `json:"reference,omitempty"`
	Description         string                `json:"description,omitempty"`
	PaymentMethods      *PaymentMethods       `json:"payment_methods,omitempty"`
	Settings            *PaymentSetupSettings `json:"settings,omitempty"`
	Customer            *PaymentSetupCustomer `json:"customer,omitempty"`
	Order               *PaymentSetupOrder    `json:"order,omitempty"`
	Industry            *PaymentSetupIndustry `json:"industry,omitempty"`
}

// Payment Setup Confirm Response
type PaymentSetupConfirmResponse struct {
	Id        string                       `json:"id,omitempty"`
	Status    PaymentSetupStatus           `json:"status,omitempty"`
	Source    *PaymentSetupSource          `json:"source,omitempty"`
	Customer  *PaymentSetupCustomer        `json:"customer,omitempty"`
	Actions   []common.AlternativeResponse `json:"actions,omitempty"`
	ExpiresOn *time.Time                   `json:"expires_on,omitempty"`
	Metadata  map[string]interface{}       `json:"metadata,omitempty"`
	Links     map[string]common.Link       `json:"_links,omitempty"`
}

// Payment Setup Status
type PaymentSetupStatus string

const (
	PaymentSetupStatusPending   PaymentSetupStatus = "Pending"
	PaymentSetupStatusConfirmed PaymentSetupStatus = "Confirmed"
	PaymentSetupStatusCompleted PaymentSetupStatus = "Completed"
	PaymentSetupStatusExpired   PaymentSetupStatus = "Expired"
	PaymentSetupStatusCancelled PaymentSetupStatus = "Cancelled"
)

// Payment Setup Customer
type PaymentSetupCustomer struct {
	Email           *PaymentSetupCustomerEmail  `json:"email,omitempty"`
	Name            string                      `json:"name,omitempty"`
	Phone           *common.Phone               `json:"phone,omitempty"`
	Device          *PaymentSetupCustomerDevice `json:"device,omitempty"`
	MerchantAccount *CustomerMerchantAccount    `json:"merchant_account,omitempty"`
	Id              string                      `json:"id,omitempty"`
	DateOfBirth     string                      `json:"date_of_birth,omitempty"`
	TaxNumber       string                      `json:"tax_number,omitempty"`
	Default         bool                        `json:"default,omitempty"`
}

// Payment Setup Customer Email
type PaymentSetupCustomerEmail struct {
	Address  string `json:"address,omitempty"`
	Verified bool   `json:"verified,omitempty"`
}

// Payment Setup Customer Device
type PaymentSetupCustomerDevice struct {
	Locale string `json:"locale,omitempty"`
}

// Customer Merchant Account
type CustomerMerchantAccount struct {
	Id                   string  `json:"id,omitempty"`
	RegistrationDate     string  `json:"registration_date,omitempty"`
	LastModified         string  `json:"last_modified,omitempty"`
	ReturningCustomer    bool    `json:"returning_customer,omitempty"`
	FirstTransactionDate string  `json:"first_transaction_date,omitempty"`
	LastTransactionDate  string  `json:"last_transaction_date,omitempty"`
	TotalOrderCount      int     `json:"total_order_count,omitempty"`
	LastPaymentAmount    float64 `json:"last_payment_amount,omitempty"`
}

// Payment Setup Billing
type PaymentSetupBilling struct {
	Address *common.Address `json:"address,omitempty"`
}

// Payment Setup Shipping
type PaymentSetupShipping struct {
	Address *common.Address `json:"address,omitempty"`
}

// Payment Methods
type PaymentMethods struct {
	Klarna *KlarnaPaymentMethod `json:"klarna,omitempty"`
	Stcpay *StcpayPaymentMethod `json:"stcpay,omitempty"`
	Tabby  *TabbyPaymentMethod  `json:"tabby,omitempty"`
	Bizum  *BizumPaymentMethod  `json:"bizum,omitempty"`
}

// Payment Method Initialization enum
type PaymentMethodInitialization string

const (
	PaymentMethodInitializationEnabled  PaymentMethodInitialization = "enabled"
	PaymentMethodInitializationDisabled PaymentMethodInitialization = "disabled"
)

// Payment Method Base - Common fields for all payment methods
type PaymentMethodBase struct {
	Status         string                      `json:"status,omitempty"`
	Flags          []string                    `json:"flags,omitempty"`
	Initialization PaymentMethodInitialization `json:"initialization,omitempty"`
}

// Payment Method Option - Unified option class like C#
type PaymentMethodOption struct {
	Id     string               `json:"id,omitempty"`
	Status string               `json:"status,omitempty"`
	Flags  []string             `json:"flags,omitempty"`
	Action *PaymentMethodAction `json:"action,omitempty"`
}

// Payment Method Action - Common action for all payment methods
type PaymentMethodAction struct {
	Type        string `json:"type,omitempty"`
	ClientToken string `json:"client_token,omitempty"`
	SessionId   string `json:"session_id,omitempty"`
}

// Payment Method Options - Unified options container like C#
type PaymentMethodOptions struct {
	Sdk          *PaymentMethodOption `json:"sdk,omitempty"`
	PayInFull    *PaymentMethodOption `json:"pay_in_full,omitempty"`
	Installments *PaymentMethodOption `json:"installments,omitempty"`
	PayNow       *PaymentMethodOption `json:"pay_now,omitempty"`
}

// Klarna Payment Method
type KlarnaPaymentMethod struct {
	PaymentMethodBase
	AccountHolder        *KlarnaAccountHolder  `json:"account_holder,omitempty"`
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}

// Klarna Account Holder
type KlarnaAccountHolder struct {
	BillingAddress *common.Address `json:"billing_address,omitempty"`
}

// Stcpay Payment Method
type StcpayPaymentMethod struct {
	PaymentMethodBase
	Otp                  string                `json:"otp,omitempty"`
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}

// Tabby Payment Method
type TabbyPaymentMethod struct {
	PaymentMethodBase
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}

// Bizum Payment Method
type BizumPaymentMethod struct {
	PaymentMethodBase
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}

// Payment Setup Settings
type PaymentSetupSettings struct {
	SuccessUrl              string `json:"success_url,omitempty"`
	FailureUrl              string `json:"failure_url,omitempty"`
	PaymentCollectionMethod string `json:"payment_collection_method,omitempty"`
}

// Payment Setup Order
type PaymentSetupOrder struct {
	Items          []PaymentSetupOrderItem    `json:"items,omitempty"`
	Shipping       *PaymentSetupOrderShipping `json:"shipping,omitempty"`
	SubMerchants   []OrderSubMerchant         `json:"sub_merchants,omitempty"`
	DiscountAmount int64                      `json:"discount_amount,omitempty"`
	Amount         int64                      `json:"amount,omitempty"`
	Currency       common.Currency            `json:"currency,omitempty"`
	Reference      string                     `json:"reference,omitempty"`
	Description    string                     `json:"description,omitempty"`
	TaxAmount      int64                      `json:"tax_amount,omitempty"`
}

// Payment Setup Order Item
type PaymentSetupOrderItem struct {
	Name           string `json:"name,omitempty"`
	Quantity       int    `json:"quantity,omitempty"`
	UnitPrice      int64  `json:"unit_price,omitempty"`
	TotalAmount    int64  `json:"total_amount,omitempty"`
	Reference      string `json:"reference,omitempty"`
	DiscountAmount int64  `json:"discount_amount,omitempty"`
	Url            string `json:"url,omitempty"`
	ImageUrl       string `json:"image_url,omitempty"`
	Type           string `json:"type,omitempty"`
}

// Payment Setup Order Shipping
type PaymentSetupOrderShipping struct {
	Address *common.Address `json:"address,omitempty"`
	Method  string          `json:"method,omitempty"`
}

// Order Sub Merchant
type OrderSubMerchant struct {
	Id               string          `json:"id,omitempty"`
	ProductCategory  string          `json:"product_category,omitempty"`
	NumberOfTrades   int             `json:"number_of_trades,omitempty"`
	RegistrationDate string          `json:"registration_date,omitempty"`
	Name             string          `json:"name,omitempty"`
	TaxId            string          `json:"tax_id,omitempty"`
	Address          *common.Address `json:"address,omitempty"`
}

// Payment Setup Industry
type PaymentSetupIndustry struct {
	AirlineData       *AirlineData        `json:"airline_data,omitempty"`
	AccommodationData []AccommodationData `json:"accommodation_data,omitempty"`
	Type              string              `json:"type,omitempty"`
}

// Airline Data
type AirlineData struct {
	Ticket           *AirlineTicket     `json:"ticket,omitempty"`
	Passengers       []AirlinePassenger `json:"passengers,omitempty"`
	FlightLegDetails []FlightLegDetail  `json:"flight_leg_details,omitempty"`
}

// Airline Ticket
type AirlineTicket struct {
	Number                 string `json:"number,omitempty"`
	IssueDate              string `json:"issue_date,omitempty"`
	IssuingCarrierCode     string `json:"issuing_carrier_code,omitempty"`
	TravelPackageIndicator string `json:"travel_package_indicator,omitempty"`
	TravelAgencyName       string `json:"travel_agency_name,omitempty"`
	TravelAgencyCode       string `json:"travel_agency_code,omitempty"`
}

// Airline Passenger
type AirlinePassenger struct {
	FirstName   string          `json:"first_name,omitempty"`
	LastName    string          `json:"last_name,omitempty"`
	DateOfBirth string          `json:"date_of_birth,omitempty"`
	Address     *common.Address `json:"address,omitempty"`
}

// Flight Leg Detail
type FlightLegDetail struct {
	FlightNumber      string `json:"flight_number,omitempty"`
	CarrierCode       string `json:"carrier_code,omitempty"`
	ClassOfTravelling string `json:"class_of_travelling,omitempty"`
	DepartureAirport  string `json:"departure_airport,omitempty"`
	DepartureDate     string `json:"departure_date,omitempty"`
	DepartureTime     string `json:"departure_time,omitempty"`
	ArrivalAirport    string `json:"arrival_airport,omitempty"`
	StopOverCode      string `json:"stop_over_code,omitempty"`
	FareBasisCode     string `json:"fare_basis_code,omitempty"`
	ServiceClass      string `json:"service_class,omitempty"`
	StopoverCode      string `json:"stopover_code,omitempty"`
}

// Accommodation Data
type AccommodationData struct {
	Name             string               `json:"name,omitempty"`
	BookingReference string               `json:"booking_reference,omitempty"`
	CheckInDate      string               `json:"check_in_date,omitempty"`
	CheckOutDate     string               `json:"check_out_date,omitempty"`
	Address          *common.Address      `json:"address,omitempty"`
	NumberOfRooms    int                  `json:"number_of_rooms,omitempty"`
	Guests           []AccommodationGuest `json:"guests,omitempty"`
	Room             []AccommodationRoom  `json:"room,omitempty"`
}

// Accommodation Guest
type AccommodationGuest struct {
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	DateOfBirth string `json:"date_of_birth,omitempty"`
}

// Accommodation Room
type AccommodationRoom struct {
	Rate           float64 `json:"rate,omitempty"`
	NumberOfNights int     `json:"number_of_nights,omitempty"`
}

// Payment Setup Source
type PaymentSetupSource struct {
	Type                    string              `json:"type,omitempty"`
	Id                      string              `json:"id,omitempty"`
	Fingerprint             string              `json:"fingerprint,omitempty"`
	Bin                     string              `json:"bin,omitempty"`
	CardType                common.CardType     `json:"card_type,omitempty"`
	CardCategory            common.CardCategory `json:"card_category,omitempty"`
	Issuer                  string              `json:"issuer,omitempty"`
	IssuerCountry           common.Country      `json:"issuer_country,omitempty"`
	ProductId               string              `json:"product_id,omitempty"`
	ProductType             string              `json:"product_type,omitempty"`
	Last4                   string              `json:"last4,omitempty"`
	ExpiryMonth             int                 `json:"expiry_month,omitempty"`
	ExpiryYear              int                 `json:"expiry_year,omitempty"`
	Name                    string              `json:"name,omitempty"`
	Scheme                  string              `json:"scheme,omitempty"`
	SchemeLocal             string              `json:"scheme_local,omitempty"`
	FastFunds               string              `json:"fast_funds,omitempty"`
	Payouts                 bool                `json:"payouts,omitempty"`
	PaymentAccountReference string              `json:"payment_account_reference,omitempty"`
}

// Payment Setup 3DS
type PaymentSetupThreeDs struct {
	Enabled            bool   `json:"enabled,omitempty"`
	AttemptN3d         bool   `json:"attempt_n3d,omitempty"`
	Eci                string `json:"eci,omitempty"`
	Cryptogram         string `json:"cryptogram,omitempty"`
	Xid                string `json:"xid,omitempty"`
	Version            string `json:"version,omitempty"`
	Exemption          string `json:"exemption,omitempty"`
	ChallengeIndicator string `json:"challenge_indicator,omitempty"`
	FlowType           string `json:"flow_type,omitempty"`
	AllowUpgrade       bool   `json:"allow_upgrade,omitempty"`
}
