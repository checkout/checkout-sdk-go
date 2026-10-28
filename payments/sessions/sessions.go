package payment_sessions

import (
	"time"

	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/payments"
	"github.com/checkout/checkout-sdk-go/v3/payments/nas"
)

type PaymentMethodsType string

const (
	AlipayCnPMT   PaymentMethodsType = "alipay_cn"
	AlipayHkPMT   PaymentMethodsType = "alipay_hk"
	AlmaPMT       PaymentMethodsType = "alma"
	ApplepayPMT   PaymentMethodsType = "applepay"
	BancontactPMT PaymentMethodsType = "bancontact"
	BenefitPMT    PaymentMethodsType = "benefit"
	BizumPMT      PaymentMethodsType = "bizum"
	CardPMT       PaymentMethodsType = "card"
	DanaPMT       PaymentMethodsType = "dana"
	EPSPMT        PaymentMethodsType = "eps"
	GcashPMT      PaymentMethodsType = "gcash"
	GooglepayPMT  PaymentMethodsType = "googlepay"
	IdealPMT      PaymentMethodsType = "ideal"
	KakaopayPMT   PaymentMethodsType = "kakaopay"
	KlarnaPMT     PaymentMethodsType = "klarna"
	KNetPMT       PaymentMethodsType = "knet"
	MbwayPMT      PaymentMethodsType = "mbway"
	MobilepayPMT  PaymentMethodsType = "mobilepay"
	MultibancoPMT PaymentMethodsType = "multibanco"
	P24PMT        PaymentMethodsType = "p24"
	PaypalPMT     PaymentMethodsType = "paypal"
	PlaidPMT      PaymentMethodsType = "plaid"
	QpayPMT       PaymentMethodsType = "qpay"
	RememberMePMT PaymentMethodsType = "remember_me"
	SepaPMT       PaymentMethodsType = "sepa"
	StcpayPMT     PaymentMethodsType = "stcpay"
	StoredCardPMT PaymentMethodsType = "stored_card"
	TabbyPMT      PaymentMethodsType = "tabby"
	TamaraPMT     PaymentMethodsType = "tamara"
	TngPMT        PaymentMethodsType = "tng"
	TruemoneyPMT  PaymentMethodsType = "truemoney"
	TwintPMT      PaymentMethodsType = "twint"
	VippsPMT      PaymentMethodsType = "vipps"
	WechatpayPMT  PaymentMethodsType = "wechatpay"
	PaynowPMT     PaymentMethodsType = "paynow"
	OctopusPMT    PaymentMethodsType = "octopus"
)

const (
	PaymentSessionsPath         = "payment-sessions"
	PaymentSessionsCompletePath = "payment-sessions/complete"
	PaymentSessionsSubmitPath   = "submit"
)

type (
	Billing struct {
		Address *common.Address `json:"address,omitempty"`
	}

	PaymentCustomerRequest struct {
		Id        string        `json:"id,omitempty"`
		Email     string        `json:"email,omitempty"`
		Name      string        `json:"name,omitempty"`
		TaxNumber string        `json:"tax_number,omitempty"`
		Phone     *common.Phone `json:"phone,omitempty"`
		Default   bool          `json:"default,omitempty"`
	}

	PaymentSessionsRequest struct {
		// Amount is the payment amount. Provide a value of 0 to perform a card verification. The amount must be provided
		// in the minor currency unit.
		// [Required]
		// min 0
		Amount int64 `json:"amount"`
		// Currency is the three-letter ISO currency code.
		// [Required]
		// max 3 characters
		// min 3 characters
		Currency common.Currency `json:"currency,omitempty"`
		// Billing is the billing details.
		// [Required]
		Billing *payments.BillingInformation `json:"billing,omitempty"`
		// SuccessUrl overrides the default success redirect URL configured on your account, for payment methods that
		// require a redirect.
		// [Required]
		// Format: uri
		// max 1024 characters
		SuccessUrl string `json:"success_url,omitempty"`
		// FailureUrl overrides the default failure redirect URL configured on your account, for payment methods that
		// require a redirect.
		// [Required]
		// Format: uri
		// max 1024 characters
		FailureUrl string `json:"failure_url,omitempty"`
		// PaymentType: Must be specified for card-not-present (CNP) payments. For example, a recurring mail order /
		// telephone order (MOTO) payment.
		// [Optional]
		// Default: Regular
		PaymentType payments.PaymentType `json:"payment_type,omitempty"`
		// BillingDescriptor is a description of the purchase, which is displayed on the customer's statement.
		// [Optional]
		BillingDescriptor *payments.BillingDescriptor `json:"billing_descriptor,omitempty"`
		// Reference is a reference you can use to identify the payment. For example, an order number. - For Amex
		// payments, this must be at most 30 characters. - For Benefit payments, the reference must be a unique
		// alphanumeric value. - For iDEAL payments, the reference is required and must be an alphanumeric value with a
		// 35-character limit.
		// [Optional]
		// max 50 characters
		Reference string `json:"reference,omitempty"`
		// Description is a description for the payment.
		// [Optional]
		// max 100 characters
		Description string `json:"description,omitempty"`
		// Customer is the customer's details. Required if source.type is tamara.
		// [Optional]
		Customer *common.CustomerRequest `json:"customer,omitempty"`
		// Shipping is the shipping details.
		// [Optional]
		Shipping *payments.ShippingDetailsFlowHostedLinks `json:"shipping,omitempty"`
		// Recipient: Information about the recipient of the payment's funds. Applies to Account Funding Transactions,
		// and VISA or Mastercard domestic UK transactions processed by financial institutions.
		// [Optional]
		Recipient *payments.PaymentRecipient `json:"recipient,omitempty"`
		// Processing: Use the processing object to influence or override the data sent during card processing.
		// [Optional]
		Processing *payments.ProcessingSettings `json:"processing,omitempty"`
		// Instruction: Details about the payment instruction.
		// [Optional]
		Instruction *payments.PaymentInstruction `json:"instruction,omitempty"`
		// ProcessingChannelId is the processing channel to use for the payment.
		// [Optional]
		// ^(pc)_(\w{26})$
		ProcessingChannelId string `json:"processing_channel_id,omitempty"`
		// PaymentMethodConfiguration: Configurations for payment method-specific settings.
		// [Optional]
		PaymentMethodConfiguration *payments.PaymentMethodConfiguration `json:"payment_method_configuration,omitempty"`
		// Items is the line items in the order.
		// [Optional]
		Items []payments.Product `json:"items,omitempty"`
		// AmountAllocations is the sub-entities that the payment is being processed on behalf of.
		// [Optional]
		AmountAllocations []common.AmountAllocations `json:"amount_allocations,omitempty"`
		// Risk configures the risk assessment performed during payment processing.
		// [Optional]
		Risk *payments.RiskRequest `json:"risk,omitempty"`
		// DisplayName is the merchant's display name.
		// [Optional]
		// max 255 characters
		DisplayName string `json:"display_name,omitempty"`
		// Metadata allows you to store additional information about a transaction with custom fields and up to five
		// user-defined fields, which can be used for reporting purposes. You can supply fields of type string, number,
		// and boolean within the metadata object. Arrays and objects are not supported. You can provide up to 18
		// metadata fields per API call, but the value of each field must not exceed 255 characters in length. You can
		// also reference metadata properties
		// (https://www.checkout.com/docs/business-operations/prevent-fraud/understand-fraud-detection#Metadata) in your
		// custom rules for Fraud Detection. For example, $coupon_code = '1234'.
		// [Optional]
		Metadata map[string]interface{} `json:"metadata,omitempty"`
		// Locale creates a translated version of the page in the specified language.
		// [Optional]
		// Default: en-GB
		Locale payments.LocalType `json:"locale,omitempty"`
		// ThreeDsRequest is the 3D Secure configuration for the payment session.
		// [Optional]
		ThreeDsRequest *payments.ThreeDsRequestFlowHostedLinks `json:"3ds,omitempty"`
		// Sender is the sender of the payment.
		// [Optional]
		Sender *nas.Sender `json:"sender,omitempty"`
		// Capture specifies whether to capture the payment, if applicable.
		// [Optional]
		// Default: true
		Capture *bool `json:"capture,omitempty"`
		// CaptureOn is a timestamp specifying when to capture the payment, as an ISO 8601 code. If a value is provided,
		// capture is automatically set to true.
		// [Optional]
		// Format: date-time
		CaptureOn *time.Time `json:"capture_on,omitempty"`
		// ExpiresOn is a timestamp specifying when the PaymentSession should expire, as an ISO 8601 code. If no value is
		// provided, expiry is set to 24 hours after the PaymentSession is created. You cannot set the session expiry to
		// more than 60 days after the PaymentSession is created.
		// [Optional]
		// Format: date-time
		ExpiresOn *time.Time `json:"expires_on,omitempty"`
		// EnabledPaymentMethods specifies which payment method options to present to the customer. The values in this
		// field override any equivalent values in disabled_payment_methods.
		// [Optional]
		EnabledPaymentMethods []PaymentMethodsType `json:"enabled_payment_methods,omitempty"`
		// DisabledPaymentMethods specifies which payment method options to not present to the customer. If you specify
		// the same payment method in this field and in enabled_payment_methods, the disabled_payment_methods value will
		// be overridden. Any payment method options not explicitly specified in this field will be presented to the
		// customer by default.
		// [Optional]
		DisabledPaymentMethods []PaymentMethodsType `json:"disabled_payment_methods,omitempty"`
		// CustomerRetry is the configuration for asynchronous retries. When set, CustomerRetry.MaxAttempts is required.
		// [Optional]
		CustomerRetry *payments.CustomerRetry `json:"customer_retry,omitempty"`
		// AuthorizationType is the authorization type.
		// [Optional]
		// Default: Final
		AuthorizationType nas.AuthorizationType `json:"authorization_type,omitempty" default:"Final"`
		// PaymentPlan is the information to process a recurring payment request. To be used when the payment_type is
		// Recurring.
		// [Optional]
		PaymentPlan *payments.PaymentPlan `json:"payment_plan,omitempty"`
		// IpAddress is the Customers IP address. Only IPv4 and IPv6 addresses are accepted.
		// [Optional]
		// Format: ipv4
		// max 45 characters
		//
		// Deprecated: IpAddress is deprecated. Use alternative fields for client identification. This field will be removed
		// in the next major version.
		IpAddress string `json:"ip_address,omitempty"`
	}
)

type (
	PaymentMethods struct {
		Type        string   `json:"type,omitempty"`
		CardSchemes []string `json:"card_schemes,omitempty"`
	}

	PaymentSessionsResponse struct {
		HttpMetadata         common.HttpMetadata
		Id                   string                 `json:"id,omitempty"`
		PaymentSessionToken  string                 `json:"payment_session_token,omitempty"`
		PaymentSessionSecret string                 `json:"payment_session_secret,omitempty"`
		Links                map[string]common.Link `json:"_links,omitempty"`
	}

	PaymentSessionsWithPaymentRequest struct {
		SessionData                string                                   `json:"session_data"`
		Amount                     int64                                    `json:"amount"`
		Currency                   common.Currency                          `json:"currency,omitempty"`
		Billing                    *payments.BillingInformation             `json:"billing,omitempty"`
		SuccessUrl                 string                                   `json:"success_url,omitempty"`
		FailureUrl                 string                                   `json:"failure_url,omitempty"`
		PaymentType                payments.PaymentType                     `json:"payment_type,omitempty"`
		BillingDescriptor          *payments.BillingDescriptor              `json:"billing_descriptor,omitempty"`
		Reference                  string                                   `json:"reference,omitempty"`
		Description                string                                   `json:"description,omitempty"`
		Customer                   *common.CustomerRequest                  `json:"customer,omitempty"`
		Shipping                   *payments.ShippingDetailsFlowHostedLinks `json:"shipping,omitempty"`
		Recipient                  *payments.PaymentRecipient               `json:"recipient,omitempty"`
		Processing                 *payments.ProcessingSettings             `json:"processing,omitempty"`
		Instruction                *payments.PaymentInstruction             `json:"instruction,omitempty"`
		ProcessingChannelId        string                                   `json:"processing_channel_id,omitempty"`
		PaymentMethodConfiguration *payments.PaymentMethodConfiguration     `json:"payment_method_configuration,omitempty"`
		Items                      []payments.Product                       `json:"items,omitempty"`
		AmountAllocations          []common.AmountAllocations               `json:"amount_allocations,omitempty"`
		Risk                       *payments.RiskRequest                    `json:"risk,omitempty"`
		DisplayName                string                                   `json:"display_name,omitempty"`
		Metadata                   map[string]interface{}                   `json:"metadata,omitempty"`
		Locale                     payments.LocalType                       `json:"locale,omitempty"`
		ThreeDsRequest             *payments.ThreeDsRequestFlowHostedLinks  `json:"3ds,omitempty"`
		Sender                     *nas.Sender                              `json:"sender,omitempty"`
		Capture                    *bool                                    `json:"capture,omitempty"`
		CaptureOn                  *time.Time                               `json:"capture_on,omitempty"`
		AuthorizationType          nas.AuthorizationType                    `json:"authorization_type,omitempty"`
		PaymentPlan                *payments.PaymentPlan                    `json:"payment_plan,omitempty"`
	}

	SubmitPaymentSessionRequest struct {
		SessionData       string                                   `json:"session_data"`
		Amount            int64                                    `json:"amount,omitempty"`
		Currency          common.Currency                          `json:"currency,omitempty"`
		Billing           *payments.BillingInformation             `json:"billing,omitempty"`
		SuccessUrl        string                                   `json:"success_url,omitempty"`
		FailureUrl        string                                   `json:"failure_url,omitempty"`
		PaymentType       payments.PaymentType                     `json:"payment_type,omitempty"`
		BillingDescriptor *payments.BillingDescriptor              `json:"billing_descriptor,omitempty"`
		Reference         string                                   `json:"reference,omitempty"`
		Customer          *common.CustomerRequest                  `json:"customer,omitempty"`
		Shipping          *payments.ShippingDetailsFlowHostedLinks `json:"shipping,omitempty"`
		Items             []payments.Product                       `json:"items,omitempty"`
		ThreeDsRequest    *payments.ThreeDsRequestFlowHostedLinks  `json:"3ds,omitempty"`
		// Deprecated: IpAddress is deprecated. Use alternative fields for client identification.
		IpAddress                  string                               `json:"ip_address,omitempty"`
		PaymentMethodConfiguration *payments.PaymentMethodConfiguration `json:"payment_method_configuration,omitempty"`
		Recipient                  *payments.PaymentRecipient           `json:"recipient,omitempty"`
		Instruction                *payments.PaymentInstruction         `json:"instruction,omitempty"`
		ProcessingChannelId        string                               `json:"processing_channel_id,omitempty"`
		Metadata                   map[string]interface{}               `json:"metadata,omitempty"`
		Sender                     *nas.Sender                          `json:"sender,omitempty"`
		Capture                    *bool                                `json:"capture,omitempty"`
		CaptureOn                  *time.Time                           `json:"capture_on,omitempty"`
		Processing                 *payments.ProcessingSettings         `json:"processing,omitempty"`
		AmountAllocations          []common.AmountAllocations           `json:"amount_allocations,omitempty"`
	}

	// Response structures for payment session submit/complete endpoints
	PaymentSessionPaymentResponse struct {
		HttpMetadata         common.HttpMetadata
		Id                   string             `json:"id,omitempty"`
		Status               string             `json:"status,omitempty"`
		Type                 PaymentMethodsType `json:"type,omitempty"`
		Action               interface{}        `json:"action,omitempty"`
		PaymentSessionId     string             `json:"payment_session_id,omitempty"`
		PaymentSessionSecret string             `json:"payment_session_secret,omitempty"`
		DeclineReason        string             `json:"decline_reason,omitempty"`
	}
)
