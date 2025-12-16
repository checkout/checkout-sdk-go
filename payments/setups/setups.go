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

type PaymentMethodInitialization string

const (
	PaymentMethodInitializationDisabled PaymentMethodInitialization = "disabled"
	PaymentMethodInitializationEnabled  PaymentMethodInitialization = "enabled"
)

// ===== Main Request/Response Structs =====

type PaymentSetupRequest struct {
	ProcessingChannelId string                `json:"processing_channel_id"`
	Amount              int64                 `json:"amount"`
	Currency            common.Currency       `json:"currency"`
	PaymentType         payments.PaymentType  `json:"payment_type,omitempty"`
	Reference           string                `json:"reference,omitempty"`
	Description         string                `json:"description,omitempty"`
	PaymentMethods      *PaymentMethods       `json:"payment_methods,omitempty"`
	Settings            *PaymentSetupSettings `json:"settings,omitempty"`
	Customer            *PaymentSetupCustomer `json:"customer,omitempty"`
	Order               *PaymentSetupOrder    `json:"order,omitempty"`
	Industry            *PaymentSetupIndustry `json:"industry,omitempty"`
}

type PaymentSetupResponse struct {
	HttpMetadata        common.HttpMetadata
	Id                  string                `json:"id,omitempty"`
	ProcessingChannelId string                `json:"processing_channel_id"`
	Amount              int64                 `json:"amount"`
	Currency            common.Currency       `json:"currency"`
	PaymentType         payments.PaymentType  `json:"payment_type,omitempty"`
	Reference           string                `json:"reference,omitempty"`
	Description         string                `json:"description,omitempty"`
	PaymentMethods      *PaymentMethods       `json:"payment_methods,omitempty"`
	Settings            *PaymentSetupSettings `json:"settings,omitempty"`
	Customer            *PaymentSetupCustomer `json:"customer,omitempty"`
	Order               *PaymentSetupOrder    `json:"order,omitempty"`
	Industry            *PaymentSetupIndustry `json:"industry,omitempty"`
}

type PaymentSetupConfirmResponse struct {
	HttpMetadata    common.HttpMetadata
	Id              string                      `json:"id"`
	ActionId        string                      `json:"action_id"`
	Amount          int64                       `json:"amount"`
	Currency        common.Currency             `json:"currency"`
	Approved        bool                        `json:"approved"`
	Status          payments.PaymentStatus      `json:"status"`
	ResponseCode    string                      `json:"response_code"`
	ProcessedOn     *time.Time                  `json:"processed_on"`
	Links           map[string]common.Link      `json:"_links"`
	AmountRequested int                         `json:"amount_requested,omitempty"`
	AuthCode        string                      `json:"auth_code,omitempty"`
	ResponseSummary string                      `json:"response_summary,omitempty"`
	ExpiresOn       *time.Time                  `json:"expires_on,omitempty"`
	ThreeDs         *payments.ThreeDsEnrollment `json:"3ds,omitempty"`
	Risk            *payments.RiskAssessment    `json:"risk,omitempty"`
	Source          *PaymentSetupSource         `json:"source,omitempty"`
	Customer        *common.CustomerResponse    `json:"customer,omitempty"`
	Balances        *PaymentSetupBalances       `json:"balances,omitempty"`
	Reference       string                      `json:"reference,omitempty"`
	Subscription    *PaymentSetupSubscription   `json:"subscription,omitempty"`
	Processing      *payments.PaymentProcessing `json:"processing,omitempty"`
	Eci             string                      `json:"eci,omitempty"`
	SchemeId        string                      `json:"scheme_id,omitempty"`
	Retry           *PaymentSetupRetry          `json:"retry,omitempty"`
}

// ===== Customer Structs =====

type PaymentSetupCustomer struct {
	Email           *PaymentSetupCustomerEmail  `json:"email,omitempty"`
	Name            string                      `json:"name,omitempty"`
	Phone           *common.Phone               `json:"phone,omitempty"`
	Device          *PaymentSetupCustomerDevice `json:"device,omitempty"`
	MerchantAccount *CustomerMerchantAccount    `json:"merchant_account,omitempty"`
}

type PaymentSetupCustomerEmail struct {
	Address  string `json:"address,omitempty"`
	Verified *bool  `json:"verified,omitempty"`
}

type PaymentSetupCustomerDevice struct {
	Locale string `json:"locale,omitempty"`
}

type CustomerMerchantAccount struct {
	Id                   string     `json:"id,omitempty"`
	RegistrationDate     *time.Time `json:"registration_date,omitempty"`
	LastModified         *time.Time `json:"last_modified,omitempty"`
	ReturningCustomer    *bool      `json:"returning_customer,omitempty"`
	FirstTransactionDate *time.Time `json:"first_transaction_date,omitempty"`
	LastTransactionDate  *time.Time `json:"last_transaction_date,omitempty"`
	TotalOrderCount      int        `json:"total_order_count,omitempty"`
	LastPaymentAmount    int64      `json:"last_payment_amount,omitempty"`
}

// ===== Payment Methods Structs =====

type PaymentMethods struct {
	Klarna *KlarnaPaymentMethod `json:"klarna,omitempty"`
	Stcpay *StcpayPaymentMethod `json:"stcpay,omitempty"`
	Tabby  *TabbyPaymentMethod  `json:"tabby,omitempty"`
	Bizum  *BizumPaymentMethod  `json:"bizum,omitempty"`
}

type PaymentMethodBase struct {
	Status         string                      `json:"status,omitempty"`
	Flags          []string                    `json:"flags,omitempty"`
	Initialization PaymentMethodInitialization `json:"initialization,omitempty"`
}

type PaymentMethodOption struct {
	Id     string               `json:"id,omitempty"`
	Status string               `json:"status,omitempty"`
	Flags  []string             `json:"flags,omitempty"`
	Action *PaymentMethodAction `json:"action,omitempty"`
}

type PaymentMethodAction struct {
	Type        string `json:"type,omitempty"`
	ClientToken string `json:"client_token,omitempty"`
	SessionId   string `json:"session_id,omitempty"`
}

type PaymentMethodOptions struct {
	Sdk          *PaymentMethodOption `json:"sdk,omitempty"`
	PayInFull    *PaymentMethodOption `json:"pay_in_full,omitempty"`
	Installments *PaymentMethodOption `json:"installments,omitempty"`
	PayNow       *PaymentMethodOption `json:"pay_now,omitempty"`
}

type KlarnaPaymentMethod struct {
	PaymentMethodBase
	AccountHolder        *KlarnaAccountHolder  `json:"account_holder,omitempty"`
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}

type KlarnaAccountHolder struct {
	BillingAddress *common.Address `json:"billing_address,omitempty"`
}

type StcpayPaymentMethod struct {
	PaymentMethodBase
	Otp                  string                `json:"otp,omitempty"`
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}

type TabbyPaymentMethod struct {
	PaymentMethodBase
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}

type BizumPaymentMethod struct {
	PaymentMethodBase
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}

// ===== Support Structs =====

type PaymentSetupSettings struct {
	SuccessUrl string `json:"success_url,omitempty"`
	FailureUrl string `json:"failure_url,omitempty"`
}

type PaymentSetupOrder struct {
	Items          []payments.Product        `json:"items,omitempty"`
	Shipping       *payments.ShippingDetails `json:"shipping,omitempty"`
	SubMerchants   []OrderSubMerchant        `json:"sub_merchants,omitempty"`
	DiscountAmount int                       `json:"discount_amount,omitempty"`
}

type OrderSubMerchant struct {
	Id               string     `json:"id,omitempty"`
	ProductCategory  string     `json:"product_category,omitempty"`
	NumberOfTrades   int        `json:"number_of_trades,omitempty"`
	RegistrationDate *time.Time `json:"registration_date,omitempty"`
}

type PaymentSetupIndustry struct {
	AirlineData       *AirlineData                 `json:"airline_data,omitempty"`
	AccommodationData []payments.AccommodationData `json:"accommodation_data,omitempty"`
}

type AirlineData struct {
	Ticket           *payments.Ticket            `json:"ticket,omitempty"`
	Passengers       []payments.Passenger        `json:"passengers,omitempty"`
	FlightLegDetails []payments.FlightLegDetails `json:"flight_leg_details,omitempty"`
}

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

type PaymentSetupBalances struct {
	TotalAuthorized    int `json:"total_authorized,omitempty"`
	TotalVoided        int `json:"total_voided,omitempty"`
	AvailableToVoid    int `json:"available_to_void,omitempty"`
	TotalCaptured      int `json:"total_captured,omitempty"`
	AvailableToCapture int `json:"available_to_capture,omitempty"`
	TotalRefunded      int `json:"total_refunded,omitempty"`
	AvailableToRefund  int `json:"available_to_refund,omitempty"`
}

type PaymentSetupSubscription struct {
	Id string `json:"id,omitempty"`
}

type PaymentSetupRetry struct {
	MaxAttempts   int        `json:"max_attempts,omitempty"`
	EndsOn        *time.Time `json:"ends_on,omitempty"`
	NextAttemptOn *time.Time `json:"next_attempt_on,omitempty"`
}
