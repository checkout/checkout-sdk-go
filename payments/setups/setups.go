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
	HttpMetadata        common.HttpMetadata
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

// Payment Setup Confirm Response - Using common payment structs for consistency
type PaymentSetupConfirmResponse struct {
	HttpMetadata    common.HttpMetadata
	Id              string                      `json:"id,omitempty"`
	ActionId        string                      `json:"action_id,omitempty"`
	Amount          int64                       `json:"amount,omitempty"`
	Currency        common.Currency             `json:"currency,omitempty"`
	Approved        bool                        `json:"approved,omitempty"`
	Status          payments.PaymentStatus      `json:"status,omitempty"`
	AuthCode        string                      `json:"auth_code,omitempty"`
	ResponseCode    string                      `json:"response_code,omitempty"`
	ResponseSummary string                      `json:"response_summary,omitempty"`
	ThreeDs         *payments.ThreeDsEnrollment `json:"3ds,omitempty"`
	Risk            *payments.RiskAssessment    `json:"risk,omitempty"`
	Source          *PaymentSetupSource         `json:"source,omitempty"`
	Customer        *common.CustomerResponse    `json:"customer,omitempty"`
	ProcessedOn     *time.Time                  `json:"processed_on,omitempty"`
	Reference       string                      `json:"reference,omitempty"`
	Processing      *payments.PaymentProcessing `json:"processing,omitempty"`
	Eci             string                      `json:"eci,omitempty"`
	SchemeId        string                      `json:"scheme_id,omitempty"`
	Links           map[string]common.Link      `json:"_links,omitempty"`
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
	Id                   string `json:"id,omitempty"`
	RegistrationDate     string `json:"registration_date,omitempty"`
	LastModified         string `json:"last_modified,omitempty"`
	ReturningCustomer    bool   `json:"returning_customer,omitempty"`
	FirstTransactionDate string `json:"first_transaction_date,omitempty"`
	LastTransactionDate  string `json:"last_transaction_date,omitempty"`
	TotalOrderCount      int    `json:"total_order_count,omitempty"`
	LastPaymentAmount    int64  `json:"last_payment_amount,omitempty"`
}

// Payment Setup Shipping
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
	SuccessUrl string `json:"success_url,omitempty"`
	FailureUrl string `json:"failure_url,omitempty"`
}

// Payment Setup Order - Using common payment structs for DRY principles
type PaymentSetupOrder struct {
	Items          []payments.Product        `json:"items,omitempty"`
	Shipping       *payments.ShippingDetails `json:"shipping,omitempty"`
	SubMerchants   []OrderSubMerchant        `json:"sub_merchants,omitempty"`
	DiscountAmount int64                     `json:"discount_amount,omitempty"`
}

// Order Sub Merchant
type OrderSubMerchant struct {
	Id               string `json:"id,omitempty"`
	ProductCategory  string `json:"product_category,omitempty"`
	NumberOfTrades   int    `json:"number_of_trades,omitempty"`
	RegistrationDate string `json:"registration_date,omitempty"`
}

// Payment Setup Industry - Using common payment structs when possible for DRY principles
type PaymentSetupIndustry struct {
	AirlineData       AirlineData                  `json:"airline_data,omitempty"`
	AccommodationData []payments.AccommodationData `json:"accommodation_data,omitempty"`
}

type AirlineData struct {
	Ticket           *payments.Ticket            `json:"ticket,omitempty"`
	Passengers       []payments.Passenger        `json:"passengers,omitempty"`
	FlightLegDetails []payments.FlightLegDetails `json:"flight_leg_details,omitempty"`
}

// Note: Using common payment structures (payments.AccommodationData, etc.)
// to eliminate code duplication and maintain consistency across the SDK

// Payment Setup Source - Using common structs for consistency
type PaymentSetupSource struct {
	Type                    string              `json:"type,omitempty"`
	Id                      string              `json:"id,omitempty"`
	BillingAddress          *common.Address     `json:"billing_address,omitempty"`
	Phone                   *common.Phone       `json:"phone,omitempty"`
	Scheme                  string              `json:"scheme,omitempty"`
	Last4                   string              `json:"last4,omitempty"`
	Fingerprint             string              `json:"fingerprint,omitempty"`
	Bin                     string              `json:"bin,omitempty"`
	CardType                common.CardType     `json:"card_type,omitempty"`
	CardCategory            common.CardCategory `json:"card_category,omitempty"`
	Issuer                  string              `json:"issuer,omitempty"`
	IssuerCountry           common.Country      `json:"issuer_country,omitempty"`
	ProductType             string              `json:"product_type,omitempty"`
	AvsCheck                string              `json:"avs_check,omitempty"`
	CvvCheck                string              `json:"cvv_check,omitempty"`
	PaymentAccountReference string              `json:"payment_account_reference,omitempty"`
}
