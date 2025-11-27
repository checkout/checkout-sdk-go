package setups

import (
	"time"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/payments"
)

const (
	PaymentSetupsPath = "payment-setups"
	ConfirmPath       = "confirm"
)

// Payment Setup Request
type PaymentSetupRequest struct {
	Amount              int64                                `json:"amount"`
	Currency            common.Currency                      `json:"currency,omitempty"`
	AmountAllocations   []common.AmountAllocations           `json:"amount_allocations,omitempty"`
	Reference           string                               `json:"reference,omitempty"`
	Description         string                               `json:"description,omitempty"`
	Customer            *PaymentSetupCustomer                `json:"customer,omitempty"`
	BillingDescriptor   *payments.BillingDescriptor          `json:"billing_descriptor,omitempty"`
	Billing             *PaymentSetupBilling                 `json:"billing,omitempty"`
	Shipping            *PaymentSetupShipping                `json:"shipping,omitempty"`
	ProcessingChannelId string                               `json:"processing_channel_id,omitempty"`
	PaymentMethods      *PaymentMethods                      `json:"payment_methods,omitempty"`
	Settings            *PaymentSetupSettings                `json:"settings,omitempty"`
	Order               *PaymentSetupOrder                   `json:"order,omitempty"`
	Industry            *PaymentSetupIndustry                `json:"industry,omitempty"`
	Recipient           *payments.PaymentRecipient           `json:"recipient,omitempty"`
	Processing          *payments.ProcessingSettings         `json:"processing,omitempty"`
	ReturnUrl           string                               `json:"return_url,omitempty"`
	Metadata            map[string]interface{}               `json:"metadata,omitempty"`
}

// Payment Setup Response
type PaymentSetupResponse struct {
	Id                  string                              `json:"id,omitempty"`
	Status              PaymentSetupStatus                  `json:"status,omitempty"`
	Amount              int64                               `json:"amount,omitempty"`
	Currency            common.Currency                     `json:"currency,omitempty"`
	AmountAllocations   []common.AmountAllocations          `json:"amount_allocations,omitempty"`
	Reference           string                              `json:"reference,omitempty"`
	Description         string                              `json:"description,omitempty"`
	Customer            *PaymentSetupCustomer               `json:"customer,omitempty"`
	BillingDescriptor   *payments.BillingDescriptor         `json:"billing_descriptor,omitempty"`
	Billing             *PaymentSetupBilling                `json:"billing,omitempty"`
	Shipping            *PaymentSetupShipping               `json:"shipping,omitempty"`
	ProcessingChannelId string                              `json:"processing_channel_id,omitempty"`
	PaymentMethods      *PaymentMethods                     `json:"payment_methods,omitempty"`
	Settings            *PaymentSetupSettings               `json:"settings,omitempty"`
	Order               *PaymentSetupOrder                  `json:"order,omitempty"`
	Industry            *PaymentSetupIndustry               `json:"industry,omitempty"`
	Recipient           *payments.PaymentRecipient          `json:"recipient,omitempty"`
	Processing          *payments.ProcessingSettings        `json:"processing,omitempty"`
	ReturnUrl           string                              `json:"return_url,omitempty"`
	Source              *PaymentSetupSource                 `json:"source,omitempty"`
	Actions             []common.AlternativeResponse        `json:"actions,omitempty"`
	ExpiresOn           *time.Time                          `json:"expires_on,omitempty"`
	CreatedOn           *time.Time                          `json:"created_on,omitempty"`
	UpdatedOn           *time.Time                          `json:"updated_on,omitempty"`
	Metadata            map[string]interface{}              `json:"metadata,omitempty"`
	Links               map[string]common.Link              `json:"_links,omitempty"`
}

// Payment Setup Confirm Request
type PaymentSetupConfirmRequest struct {
	Source   *PaymentSetupSource   `json:"source,omitempty"`
	ThreeDs  *PaymentSetupThreeDs  `json:"3ds,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
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
	Id           string             `json:"id,omitempty"`
	Email        string             `json:"email,omitempty"`
	Name         string             `json:"name,omitempty"`
	Phone        *common.Phone      `json:"phone,omitempty"`
	DateOfBirth  string             `json:"date_of_birth,omitempty"`
	TaxNumber    string             `json:"tax_number,omitempty"`
	Default      bool               `json:"default,omitempty"`
	MerchantAccount *MerchantAccount `json:"merchant_account,omitempty"`
}

// Merchant Account
type MerchantAccount struct {
	Entity string `json:"entity,omitempty"`
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
	Bizum     *BizumPaymentMethod     `json:"bizum,omitempty"`
	Klarna    *KlarnaPaymentMethod    `json:"klarna,omitempty"`
	Stcpay    *StcpayPaymentMethod    `json:"stcpay,omitempty"`
	Tabby     *TabbyPaymentMethod     `json:"tabby,omitempty"`
}

// Bizum Payment Method
type BizumPaymentMethod struct {
	Options *PaymentMethodOptions `json:"options,omitempty"`
}

// Klarna Payment Method  
type KlarnaPaymentMethod struct {
	Options        *PaymentMethodOptions `json:"options,omitempty"`
	AccountHolder  *KlarnaAccountHolder  `json:"account_holder,omitempty"`
}

// Klarna Account Holder
type KlarnaAccountHolder struct {
	Type           string          `json:"type,omitempty"`
	TaxId          string          `json:"tax_id,omitempty"`
	DateOfBirth    string          `json:"date_of_birth,omitempty"`
	BillingAddress *common.Address `json:"billing_address,omitempty"`
}

// Stcpay Payment Method
type StcpayPaymentMethod struct {
	Options *PaymentMethodOptions `json:"options,omitempty"`
}

// Tabby Payment Method
type TabbyPaymentMethod struct {
	Options *PaymentMethodOptions `json:"options,omitempty"`
}

// Payment Method Options
type PaymentMethodOptions struct {
	Initialization *PaymentMethodInitialization `json:"initialization,omitempty"`
}

// Payment Method Initialization
type PaymentMethodInitialization struct {
	PaymentMethodActions []PaymentMethodAction `json:"payment_method_actions,omitempty"`
}

// Payment Method Action
type PaymentMethodAction struct {
	Type    string                 `json:"type,omitempty"`
	Options *PaymentMethodOption   `json:"options,omitempty"`
}

// Payment Method Option  
type PaymentMethodOption struct {
	RequiredDocuments []string `json:"required_documents,omitempty"`
}

// Payment Setup Settings
type PaymentSetupSettings struct {
	PaymentCollectionMethod string `json:"payment_collection_method,omitempty"`
}

// Payment Setup Order
type PaymentSetupOrder struct {
	Amount        int64                    `json:"amount,omitempty"`
	Currency      common.Currency          `json:"currency,omitempty"`
	Reference     string                   `json:"reference,omitempty"`
	Description   string                   `json:"description,omitempty"`
	TaxAmount     int64                    `json:"tax_amount,omitempty"`
	Items         []payments.Product       `json:"items,omitempty"`
	SubMerchant   *OrderSubMerchant        `json:"sub_merchant,omitempty"`
}

// Order Sub Merchant
type OrderSubMerchant struct {
	Name           string          `json:"name,omitempty"`
	TaxId          string          `json:"tax_id,omitempty"`
	Address        *common.Address `json:"address,omitempty"`
}

// Payment Setup Industry
type PaymentSetupIndustry struct {
	Type        string       `json:"type,omitempty"`
	AirlineData *AirlineData `json:"airline_data,omitempty"`
}

// Airline Data
type AirlineData struct {
	Ticket             *AirlineTicket     `json:"ticket,omitempty"`
	Passenger          *AirlinePassenger  `json:"passenger,omitempty"`
	FlightLegDetails   []FlightLegDetail  `json:"flight_leg_details,omitempty"`
}

// Airline Ticket
type AirlineTicket struct {
	Number            string `json:"number,omitempty"`
	IssueDate         string `json:"issue_date,omitempty"`
	IssuingCarrierCode string `json:"issuing_carrier_code,omitempty"`
	TravelAgencyName  string `json:"travel_agency_name,omitempty"`
	TravelAgencyCode  string `json:"travel_agency_code,omitempty"`
}

// Airline Passenger
type AirlinePassenger struct {
	Name *AirlinePassengerName `json:"name,omitempty"`
}

// Airline Passenger Name
type AirlinePassengerName struct {
	First  string `json:"first,omitempty"`
	Last   string `json:"last,omitempty"`
}

// Flight Leg Detail
type FlightLegDetail struct {
	FlightNumber         string `json:"flight_number,omitempty"`
	CarrierCode          string `json:"carrier_code,omitempty"`
	ServiceClass         string `json:"service_class,omitempty"`
	DepartureDate        string `json:"departure_date,omitempty"`
	DepartureTime        string `json:"departure_time,omitempty"`
	DepartureAirport     string `json:"departure_airport,omitempty"`
	ArrivalAirport       string `json:"arrival_airport,omitempty"`
	StopoverCode         string `json:"stopover_code,omitempty"`
	FareBasisCode        string `json:"fare_basis_code,omitempty"`
}

// Payment Setup Source
type PaymentSetupSource struct {
	Type                    string                 `json:"type,omitempty"`
	Id                      string                 `json:"id,omitempty"`
	Fingerprint             string                 `json:"fingerprint,omitempty"`
	Bin                     string                 `json:"bin,omitempty"`
	CardType                common.CardType        `json:"card_type,omitempty"`
	CardCategory            common.CardCategory    `json:"card_category,omitempty"`
	Issuer                  string                 `json:"issuer,omitempty"`
	IssuerCountry           common.Country         `json:"issuer_country,omitempty"`
	ProductId               string                 `json:"product_id,omitempty"`
	ProductType             string                 `json:"product_type,omitempty"`
	Last4                   string                 `json:"last4,omitempty"`
	ExpiryMonth             int                    `json:"expiry_month,omitempty"`
	ExpiryYear              int                    `json:"expiry_year,omitempty"`
	Name                    string                 `json:"name,omitempty"`
	Scheme                  string                 `json:"scheme,omitempty"`
	SchemeLocal             string                 `json:"scheme_local,omitempty"`
	FastFunds               string                 `json:"fast_funds,omitempty"`
	Payouts                 bool                   `json:"payouts,omitempty"`
	PaymentAccountReference string                 `json:"payment_account_reference,omitempty"`
}

// Payment Setup 3DS
type PaymentSetupThreeDs struct {
	Enabled                 bool   `json:"enabled,omitempty"`
	AttemptN3d              bool   `json:"attempt_n3d,omitempty"`
	Eci                     string `json:"eci,omitempty"`
	Cryptogram              string `json:"cryptogram,omitempty"`
	Xid                     string `json:"xid,omitempty"`
	Version                 string `json:"version,omitempty"`
	Exemption               string `json:"exemption,omitempty"`
	ChallengeIndicator      string `json:"challenge_indicator,omitempty"`
	FlowType                string `json:"flow_type,omitempty"`
	AllowUpgrade            bool   `json:"allow_upgrade,omitempty"`
}