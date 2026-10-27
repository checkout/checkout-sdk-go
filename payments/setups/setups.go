package setups

import (
	"time"

	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/payments"
)

const (
	// PaymentSetupsPath is the base path of the Payment Setups API.
	PaymentSetupsPath = "payments/setups"
	// ConfirmPath is the path segment used to confirm a Payment Setup.
	ConfirmPath = "confirm"
)

// PaymentMethodInitialization is the initialization state of a payment method.
type PaymentMethodInitialization string

const (
	PaymentMethodInitializationDisabled PaymentMethodInitialization = "disabled"
	PaymentMethodInitializationEnabled  PaymentMethodInitialization = "enabled"
)

// PaymentSetupDeviceClient is the type of client the customer uses to initiate the payment.
type PaymentSetupDeviceClient string

const (
	PaymentSetupDeviceClientWeb       PaymentSetupDeviceClient = "web"
	PaymentSetupDeviceClientMobileWeb PaymentSetupDeviceClient = "mobile_web"
	PaymentSetupDeviceClientApp       PaymentSetupDeviceClient = "app"
)

// PaymentSetupDeviceOs is the operating system of the customer's device.
type PaymentSetupDeviceOs string

const (
	PaymentSetupDeviceOsAndroid PaymentSetupDeviceOs = "android"
	PaymentSetupDeviceOsIos     PaymentSetupDeviceOs = "ios"
)

// PaymentSetupPaymentType is the type of payment of a Payment Setup. The Payment Setups API uses
// lowercase values, unlike payments.PaymentType, which belongs to the payments API.
type PaymentSetupPaymentType string

const (
	PaymentSetupPaymentTypeRegular     PaymentSetupPaymentType = "regular"
	PaymentSetupPaymentTypeRecurring   PaymentSetupPaymentType = "recurring"
	PaymentSetupPaymentTypeMoto        PaymentSetupPaymentType = "moto"
	PaymentSetupPaymentTypeInstallment PaymentSetupPaymentType = "installment"
	PaymentSetupPaymentTypePayLater    PaymentSetupPaymentType = "pay_later"
	PaymentSetupPaymentTypeUnscheduled PaymentSetupPaymentType = "unscheduled"
)

// ===== Main Request/Response Structs =====

// PaymentSetupRequest represents the request body for POST /payments/setups and
// PUT /payments/setups/{id} (schema PaymentSetup).
type PaymentSetupRequest struct {
	// ProcessingChannelId is the processing channel to use for the payment.
	// [Required]
	// ^(pc)_(\w{26})$
	ProcessingChannelId string `json:"processing_channel_id"`

	// Amount is the payment amount, in the minor currency unit.
	// [Required]
	Amount int64 `json:"amount"`

	// Currency is the currency of the payment, as a three-letter ISO currency code.
	// [Required]
	Currency common.Currency `json:"currency"`

	// PaymentType is the type of payment. You must provide this field for card payments in which
	// the cardholder is not present. For example, if the transaction is a recurring payment, or a
	// mail order/telephone order (MOTO) payment.
	// [Optional]
	// Default: "regular"
	PaymentType PaymentSetupPaymentType `json:"payment_type,omitempty"`

	// Reference is a reference you can use to identify the payment. For example, an order number.
	// [Optional]
	// max 80 characters
	Reference string `json:"reference,omitempty"`

	// Description is a description of the payment.
	// [Optional]
	// max 100 characters
	Description string `json:"description,omitempty"`

	// PaymentMethods holds the payment methods that are enabled on your account and available for use.
	// [Optional]
	PaymentMethods *PaymentMethods `json:"payment_methods,omitempty"`

	// Settings holds the settings for the Payment Setup.
	// [Optional]
	Settings *PaymentSetupSettings `json:"settings,omitempty"`

	// Customer holds the customer's details.
	// [Optional]
	Customer *PaymentSetupCustomer `json:"customer,omitempty"`

	// Order holds the customer's order details.
	// [Optional]
	Order *PaymentSetupOrder `json:"order,omitempty"`

	// Billing holds the billing details for the payment.
	// [Optional]
	Billing *PaymentSetupBilling `json:"billing,omitempty"`

	// Industry holds industry-specific information.
	// [Optional]
	Industry *PaymentSetupIndustry `json:"industry,omitempty"`

	// AccountFundingTransaction holds the account funding transaction details for the payment.
	// [Optional]
	AccountFundingTransaction *PaymentSetupAccountFundingTransaction `json:"account_funding_transaction,omitempty"`

	// BillingDescriptor is the billing descriptor for the payment.
	// [Optional]
	BillingDescriptor *PaymentSetupBillingDescriptor `json:"billing_descriptor,omitempty"`

	// PresentmentDetails is the amount and currency to present to the customer, when the
	// settlement currency differs from the customer-facing currency.
	// [Optional]
	PresentmentDetails *PaymentSetupPresentmentDetails `json:"presentment_details,omitempty"`

	// Terminal holds the terminal details.
	// [Optional]
	Terminal *PaymentSetupTerminal `json:"terminal,omitempty"`

	// Risk holds the risk assessment settings for the Payment Setup.
	// [Optional]
	Risk *PaymentSetupRisk `json:"risk,omitempty"`

	// Metadata is a set of key/value pairs you can attach to a payment. It can be useful for
	// storing additional information about the payment. Each value is a string (max 255
	// characters), a number or a boolean.
	// [Optional]
	// max 20 properties
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// PaymentSetupResponse represents the 200 response body of POST /payments/setups,
// PUT /payments/setups/{id}, GET /payments/setups/{id} and
// POST /payments/setups/{id}/confirm/{payment_method_name} (schema PaymentSetup).
type PaymentSetupResponse struct {
	// HttpMetadata holds the HTTP status code and headers of the response. Not part of the body.
	HttpMetadata common.HttpMetadata

	// Id is the unique identifier of the Payment Setup.
	// [Optional]
	// Read only
	Id string `json:"id,omitempty"`

	// ProcessingChannelId is the processing channel to use for the payment.
	// [Required]
	// ^(pc)_(\w{26})$
	ProcessingChannelId string `json:"processing_channel_id"`

	// Amount is the payment amount, in the minor currency unit.
	// [Required]
	Amount int64 `json:"amount"`

	// Currency is the currency of the payment, as a three-letter ISO currency code.
	// [Required]
	Currency common.Currency `json:"currency"`

	// PaymentType is the type of payment. You must provide this field for card payments in which
	// the cardholder is not present. For example, if the transaction is a recurring payment, or a
	// mail order/telephone order (MOTO) payment.
	// [Optional]
	// Default: "regular"
	PaymentType PaymentSetupPaymentType `json:"payment_type,omitempty"`

	// Reference is a reference you can use to identify the payment. For example, an order number.
	// [Optional]
	// max 80 characters
	Reference string `json:"reference,omitempty"`

	// Description is a description of the payment.
	// [Optional]
	// max 100 characters
	Description string `json:"description,omitempty"`

	// PaymentMethods holds the payment methods that are enabled on your account and available for use.
	// [Optional]
	PaymentMethods *PaymentMethods `json:"payment_methods,omitempty"`

	// AvailablePaymentMethods is an ordered list of available payment method names. The order
	// indicates the recommended presentation priority, with the first item being the highest priority.
	// [Optional]
	// Read only
	AvailablePaymentMethods []string `json:"available_payment_methods,omitempty"`

	// Settings holds the settings for the Payment Setup.
	// [Optional]
	Settings *PaymentSetupSettings `json:"settings,omitempty"`

	// Customer holds the customer's details.
	// [Optional]
	Customer *PaymentSetupCustomer `json:"customer,omitempty"`

	// Order holds the customer's order details.
	// [Optional]
	Order *PaymentSetupOrder `json:"order,omitempty"`

	// Billing holds the billing details for the payment.
	// [Optional]
	Billing *PaymentSetupBilling `json:"billing,omitempty"`

	// Industry holds industry-specific information.
	// [Optional]
	Industry *PaymentSetupIndustry `json:"industry,omitempty"`

	// AccountFundingTransaction holds the account funding transaction details for the payment.
	// [Optional]
	AccountFundingTransaction *PaymentSetupAccountFundingTransaction `json:"account_funding_transaction,omitempty"`

	// BillingDescriptor is the billing descriptor for the payment.
	// [Optional]
	BillingDescriptor *PaymentSetupBillingDescriptor `json:"billing_descriptor,omitempty"`

	// PresentmentDetails is the amount and currency to present to the customer, when the
	// settlement currency differs from the customer-facing currency.
	// [Optional]
	PresentmentDetails *PaymentSetupPresentmentDetails `json:"presentment_details,omitempty"`

	// Terminal holds the terminal details.
	// [Optional]
	Terminal *PaymentSetupTerminal `json:"terminal,omitempty"`

	// Risk holds the risk assessment settings for the Payment Setup.
	// [Optional]
	Risk *PaymentSetupRisk `json:"risk,omitempty"`

	// LatestPayment is the latest payment for this setup. Returned after a successful
	// confirmation, including auto-confirm during setup creation.
	// [Optional]
	// Read only
	LatestPayment map[string]interface{} `json:"latest_payment,omitempty"`

	// Metadata is a set of key/value pairs you can attach to a payment. It can be useful for
	// storing additional information about the payment. Each value is a string (max 255
	// characters), a number or a boolean.
	// [Optional]
	// max 20 properties
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// PaymentSetupBilling is the billing details for the payment.
type PaymentSetupBilling struct {
	// Address is a physical address.
	// [Optional]
	Address *common.Address `json:"address,omitempty"`
}

// PaymentSetupBillingDescriptor is the billing descriptor for the payment.
type PaymentSetupBillingDescriptor struct {
	// Name is a dynamic description of the payment.
	// [Optional]
	// max 25 characters
	Name string `json:"name,omitempty"`

	// City is the city from which the payment was made.
	// [Optional]
	// max 13 characters
	City string `json:"city,omitempty"`

	// Reference is the reference shown on the statement.
	// [Optional]
	// max 50 characters
	Reference string `json:"reference,omitempty"`
}

// PaymentSetupPresentmentDetails is the amount and currency to present to the customer, when the
// settlement currency differs from the customer-facing currency.
type PaymentSetupPresentmentDetails struct {
	// Amount is the presentment amount, in the minor currency unit.
	// [Optional]
	// Format: int64
	Amount int64 `json:"amount,omitempty"`

	// Currency is the presentment currency, as a three-letter ISO currency code.
	// [Optional]
	Currency common.Currency `json:"currency,omitempty"`
}

// PaymentSetupRisk holds the risk assessment settings for the Payment Setup.
type PaymentSetupRisk struct {
	// Enabled indicates whether a risk assessment should be performed. The API defaults to true when omitted.
	// [Optional]
	Enabled *bool `json:"enabled,omitempty"`
}

// PaymentSetupTerminal holds terminal details.
type PaymentSetupTerminal struct {
	// Id is the terminal identifier.
	// [Optional]
	// min 8 characters, max 8 characters
	Id string `json:"id,omitempty"`

	// LocalDateTime is the local date and time on the terminal, in ISO 8601 format.
	// [Optional]
	// Format: date-time (RFC 3339)
	LocalDateTime *time.Time `json:"local_date_time,omitempty"`
}

// ===== Customer Structs =====

// PaymentSetupCustomer holds the customer's details.
type PaymentSetupCustomer struct {
	// Id is the unique identifier of the customer.
	// [Optional]
	Id string `json:"id,omitempty"`

	// Country is the two-letter ISO country code of the customer for this payment.
	// [Optional]
	// min 2 characters, max 2 characters
	Country common.Country `json:"country,omitempty"`

	// Email holds the details of the customer's email.
	// [Optional]
	Email *PaymentSetupCustomerEmail `json:"email,omitempty"`

	// Name is the customer's full name.
	// [Optional]
	// max 100 characters
	Name string `json:"name,omitempty"`

	// TaxNumber is the customer's tax identification number.
	// [Optional]
	TaxNumber string `json:"tax_number,omitempty"`

	// Phone is the customer's phone number.
	// [Optional]
	Phone *common.Phone `json:"phone,omitempty"`

	// Device holds the details of the customer's device.
	// [Optional]
	Device *PaymentSetupCustomerDevice `json:"device,omitempty"`

	// MerchantAccount holds the details of the account the customer holds with the merchant.
	// [Optional]
	MerchantAccount *CustomerMerchantAccount `json:"merchant_account,omitempty"`
}

// PaymentSetupCustomerEmail holds the details of the customer's email.
type PaymentSetupCustomerEmail struct {
	// Address is the customer's email address.
	// [Optional]
	Address string `json:"address,omitempty"`

	// Verified specifies whether the customer's email address is verified.
	// [Optional]
	Verified *bool `json:"verified,omitempty"`
}

// PaymentSetupCustomerDevice holds the details of the customer's device.
type PaymentSetupCustomerDevice struct {
	// Locale is the locale of the device.
	// [Optional]
	Locale string `json:"locale,omitempty"`

	// Fingerprint is a unique identifier for the customer's device.
	// [Optional]
	Fingerprint string `json:"fingerprint,omitempty"`

	// Ipv4 is the customer's device IPv4 address, used by some payment methods for risk and
	// eligibility checks.
	// [Optional]
	Ipv4 string `json:"ipv4,omitempty"`

	// Ipv6 is the customer's device IPv6 address, used by some payment methods for risk and
	// eligibility checks.
	// [Optional]
	Ipv6 string `json:"ipv6,omitempty"`

	// Client is the type of client the customer uses to initiate the payment. Required when using
	// Cash App Pay.
	// [Optional]
	Client PaymentSetupDeviceClient `json:"client,omitempty"`

	// Os is the operating system of the customer's device.
	// [Optional]
	Os PaymentSetupDeviceOs `json:"os,omitempty"`
}

// CustomerMerchantAccount holds the details of the account the customer holds with the merchant.
type CustomerMerchantAccount struct {
	// Id is the merchant's unique identifier for the customer's account.
	// [Optional]
	Id string `json:"id,omitempty"`

	// RegistrationDate is the date the customer registered their account with the merchant.
	// [Optional]
	// Format: yyyy-MM-dd
	RegistrationDate *common.APIShortDate `json:"registration_date,omitempty"`

	// LastModified is the date the customer's account with the merchant was last modified.
	// [Optional]
	// Format: yyyy-MM-dd
	LastModified *common.APIShortDate `json:"last_modified,omitempty"`

	// ReturningCustomer specifies if the customer is a returning customer.
	// [Optional]
	ReturningCustomer *bool `json:"returning_customer,omitempty"`

	// FirstTransactionDate is the date of the customer's first transaction.
	// [Optional]
	// Format: yyyy-MM-dd
	FirstTransactionDate *common.APIShortDate `json:"first_transaction_date,omitempty"`

	// LastTransactionDate is the date of the customer's most recent transaction.
	// [Optional]
	// Format: yyyy-MM-dd
	LastTransactionDate *common.APIShortDate `json:"last_transaction_date,omitempty"`

	// TotalOrderCount is the total number of orders made by the customer.
	// [Optional]
	TotalOrderCount int `json:"total_order_count,omitempty"`

	// LastPaymentAmount is the payment amount of the customer's most recent transaction. The
	// specification types this as a number (example 55.99); the API currently stores it as an
	// integer.
	// [Optional]
	LastPaymentAmount float64 `json:"last_payment_amount,omitempty"`
}

// ===== Payment Methods Structs =====

// PaymentMethods holds the payment methods that are enabled on your account and available for use.
type PaymentMethods struct {
	// AutoConfirm is the payment method that automatically confirms the payment when it is ready,
	// for example "card".
	// [Optional]
	AutoConfirm string `json:"auto_confirm,omitempty"`

	// Instrument is the instrument payment method's details and configuration.
	// [Optional]
	Instrument *InstrumentPaymentMethod `json:"instrument,omitempty"`

	// Klarna is the Klarna payment method's details and configuration.
	// [Optional]
	Klarna *KlarnaPaymentMethod `json:"klarna,omitempty"`

	// Stcpay is the stc pay payment method's details and configuration.
	// [Optional]
	Stcpay *StcpayPaymentMethod `json:"stcpay,omitempty"`

	// Tabby is the Tabby payment method's details and configuration.
	// [Optional]
	Tabby *TabbyPaymentMethod `json:"tabby,omitempty"`

	// Bizum is the Bizum payment method's details and configuration.
	// [Optional]
	// Read only
	Bizum *BizumPaymentMethod `json:"bizum,omitempty"`

	// Paynow is the PayNow payment method's details and configuration.
	// [Optional]
	// Read only
	Paynow *SimplePaymentMethod `json:"paynow,omitempty"`

	// Qpay is the QPay payment method's details and configuration.
	// [Optional]
	Qpay *SimplePaymentMethod `json:"qpay,omitempty"`

	// Eps is the EPS payment method's details and configuration.
	// [Optional]
	// Read only
	Eps *SimplePaymentMethod `json:"eps,omitempty"`

	// Ideal is the iDEAL payment method's details and configuration.
	// [Optional]
	Ideal *IdealPaymentMethod `json:"ideal,omitempty"`

	// Knet is the KNET payment method's details and configuration.
	// [Optional]
	Knet *KnetPaymentMethod `json:"knet,omitempty"`

	// Bancontact is the Bancontact payment method's details and configuration.
	// [Optional]
	Bancontact *BancontactPaymentMethod `json:"bancontact,omitempty"`

	// Benefit is the Benefit payment method's details and configuration.
	// [Optional]
	// Read only
	Benefit *SimplePaymentMethod `json:"benefit,omitempty"`

	// Vipps is the Vipps payment method's details and configuration.
	// [Optional]
	// Read only
	Vipps *SimplePaymentMethod `json:"vipps,omitempty"`

	// Twint is the Twint payment method's details and configuration.
	// [Optional]
	// Read only
	Twint *SimplePaymentMethod `json:"twint,omitempty"`

	// AlipayCn is the Alipay CN payment method's details and configuration.
	// [Optional]
	AlipayCn *SimplePaymentMethod `json:"alipay_cn,omitempty"`

	// AlipayHk is the Alipay HK payment method's details and configuration.
	// [Optional]
	AlipayHk *SimplePaymentMethod `json:"alipay_hk,omitempty"`

	// Gcash is the GCash payment method's details and configuration.
	// [Optional]
	Gcash *SimplePaymentMethod `json:"gcash,omitempty"`

	// Tng is the TNG payment method's details and configuration.
	// [Optional]
	Tng *SimplePaymentMethod `json:"tng,omitempty"`

	// Dana is the Dana payment method's details and configuration.
	// [Optional]
	Dana *SimplePaymentMethod `json:"dana,omitempty"`

	// Mobilepay is the MobilePay payment method's details and configuration.
	// [Optional]
	// Read only
	Mobilepay *SimplePaymentMethod `json:"mobilepay,omitempty"`

	// Tamara is the Tamara payment method's details and configuration.
	// [Optional]
	Tamara *SimplePaymentMethod `json:"tamara,omitempty"`

	// Mbway is the MBWay payment method's details and configuration.
	// [Optional]
	// Read only
	Mbway *SimplePaymentMethod `json:"mbway,omitempty"`

	// Multibanco is the Multibanco payment method's details and configuration.
	// [Optional]
	Multibanco *MultibancoPaymentMethod `json:"multibanco,omitempty"`

	// Wechatpay is the WeChatPay payment method's details and configuration.
	// [Optional]
	// Read only
	Wechatpay *SimplePaymentMethod `json:"wechatpay,omitempty"`

	// Kakaopay is the KakaoPay payment method's details and configuration.
	// [Optional]
	Kakaopay *SimplePaymentMethod `json:"kakaopay,omitempty"`

	// Truemoney is the TrueMoney payment method's details and configuration.
	// [Optional]
	Truemoney *SimplePaymentMethod `json:"truemoney,omitempty"`

	// Octopus is the Octopus payment method's details and configuration.
	// [Optional]
	// Read only
	Octopus *SimplePaymentMethod `json:"octopus,omitempty"`

	// P24 is the P24 (Przelewy24) payment method's details and configuration.
	// [Optional]
	P24 *P24PaymentMethod `json:"p24,omitempty"`

	// Alma is the Alma payment method's details and configuration.
	// [Optional]
	Alma *SimplePaymentMethod `json:"alma,omitempty"`

	// Swish is the Swish payment method's details and configuration.
	// [Optional]
	Swish *SwishPaymentMethod `json:"swish,omitempty"`

	// Sequra is the Sequra payment method's details and configuration.
	// [Optional]
	Sequra *SimplePaymentMethod `json:"sequra,omitempty"`

	// Ach is the ACH payment method's details and configuration.
	// [Optional]
	Ach *AchPaymentMethod `json:"ach,omitempty"`

	// Sepa is the SEPA payment method's details and configuration.
	// [Optional]
	Sepa *SepaPaymentMethod `json:"sepa,omitempty"`

	// Paypal is the PayPal payment method's details and configuration.
	// [Optional]
	Paypal *PaypalPaymentMethod `json:"paypal,omitempty"`

	// CashApp is the Cash App payment method's details and configuration. The wire key is
	// cashapp, one lowercase word.
	// [Optional]
	CashApp *CashAppPaymentMethod `json:"cashapp,omitempty"`

	// Googlepay is the Google Pay payment method's details and configuration.
	// [Optional]
	Googlepay *GooglePayPaymentMethod `json:"googlepay,omitempty"`

	// Applepay is the Apple Pay payment method's details and configuration.
	// [Optional]
	Applepay *ApplePayPaymentMethod `json:"applepay,omitempty"`

	// Card is the Card payment method's details and configuration.
	// [Optional]
	Card *CardPaymentMethod `json:"card,omitempty"`

	// Blik is the Blik payment method's details and configuration.
	// [Optional]
	Blik *BlikPaymentMethod `json:"blik,omitempty"`

	// Bacs is the Bacs payment method's details and configuration.
	// [Optional]
	Bacs *BacsPaymentMethod `json:"bacs,omitempty"`

	// CardPresent is the Card Present payment method's details and configuration. Not part of the
	// current Payment Setup schema.
	// [Optional]
	CardPresent *CardPresentPaymentMethod `json:"card_present,omitempty"`

	// PayByBank is the Pay by Bank (Open Banking) payment method's details and configuration.
	// [Optional]
	PayByBank *PayByBankPaymentMethod `json:"paybybank,omitempty"`

	// Stablecoin is the Stablecoin payment method's details and configuration.
	// [Optional]
	Stablecoin *StablecoinPaymentMethod `json:"stablecoin,omitempty"`
}

// PaymentMethodBase holds the status, flags and initialization of a payment method (schemas
// PaymentSetupPaymentMethod and PaymentMethodInitialization). Some types that embed it have no
// initialization in their schema; they declare their own deprecated Initialization, which shadows
// this one (see deprecated_initialization.go).
type PaymentMethodBase struct {
	// Status is the payment method status.
	// [Optional]
	// Read only
	// Enum: "unavailable" "action_required" "ready" "initialization_required" "invalid"
	Status string `json:"status,omitempty"`

	// Flags is the list of error codes or indicators that highlight missing or invalid information.
	// [Optional]
	// Read only
	Flags []string `json:"flags,omitempty"`

	// Initialization is the initialization state of the payment method. When you create a
	// Payment Setup, this defaults to disabled.
	// [Optional]
	// Default: "disabled"
	Initialization PaymentMethodInitialization `json:"initialization,omitempty"`
}

// PaymentSetupPaymentMethod holds the status and flags of a payment method whose schema does not
// support initialization.
type PaymentSetupPaymentMethod struct {
	// Status is the payment method status.
	// [Optional]
	// Read only
	// Enum: "unavailable" "action_required" "ready" "initialization_required" "invalid"
	Status string `json:"status,omitempty"`

	// Flags is the list of error codes or indicators that highlight missing or invalid information.
	// [Optional]
	// Read only
	Flags []string `json:"flags,omitempty"`
}

// PaymentMethodOption is a payment method option. Not part of the current Payment Setup schema.
type PaymentMethodOption struct {
	// Id is the unique identifier of the payment method option.
	// [Optional]
	Id string `json:"id,omitempty"`

	// Status is the payment method option status.
	// [Optional]
	// Read only
	Status string `json:"status,omitempty"`

	// Flags is the list of error codes or indicators that highlight missing or invalid information.
	// [Optional]
	// Read only
	Flags []string `json:"flags,omitempty"`

	// Action is the next available action for the payment method option.
	// [Optional]
	// Read only
	Action *PaymentMethodAction `json:"action,omitempty"`
}

// PaymentMethodAction is the next available action for a payment method option.
type PaymentMethodAction struct {
	// Type is the type of action.
	// [Optional]
	Type string `json:"type,omitempty"`

	// ClientToken is the client token for the provider SDK.
	// [Optional]
	ClientToken string `json:"client_token,omitempty"`

	// SessionId is the session ID.
	// [Optional]
	SessionId string `json:"session_id,omitempty"`
}

// PaymentMethodOptions holds the payment method options. Not part of the current Payment Setup schema.
type PaymentMethodOptions struct {
	// Sdk is the SDK payment method option.
	// [Optional]
	Sdk *PaymentMethodOption `json:"sdk,omitempty"`

	// PayInFull is the pay in full payment method option.
	// [Optional]
	PayInFull *PaymentMethodOption `json:"pay_in_full,omitempty"`

	// Installments is the installments payment method option.
	// [Optional]
	Installments *PaymentMethodOption `json:"installments,omitempty"`

	// PayNow is the pay now payment method option.
	// [Optional]
	PayNow *PaymentMethodOption `json:"pay_now,omitempty"`
}

// KlarnaPaymentMethod is the Klarna payment method's details and configuration.
type KlarnaPaymentMethod struct {
	PaymentMethodBase

	// Action is the next available action for the payment method.
	// [Optional]
	// Read only
	Action *KlarnaAction `json:"action,omitempty"`

	// AccountHolder is the account holder details returned by Klarna after the shopper completes
	// verification.
	// [Optional]
	// Read only
	AccountHolder *KlarnaAccountHolder `json:"account_holder,omitempty"`

	// Iban is the IBAN of the account used to complete Klarna verification.
	// [Optional]
	// Read only
	Iban string `json:"iban,omitempty"`

	// PaymentMethodOptions holds the payment method options.
	// [Optional]
	//
	// Deprecated: not part of the API spec; the API no longer sends it and ignores it on
	// requests. Use Action instead. It will be removed in the next major version.
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}

// KlarnaActionType is the type of the next available action for the Klarna payment method.
type KlarnaActionType string

const (
	// KlarnaActionTypeSdk completes the payment with the Klarna SDK, using the client_token and
	// session_id.
	KlarnaActionTypeSdk KlarnaActionType = "sdk"
)

// KlarnaAction is the next available action for the Klarna payment method.
type KlarnaAction struct {
	// Type is the type of action.
	// [Optional]
	// Read only
	Type KlarnaActionType `json:"type,omitempty"`

	// ClientToken is the client token for the Klarna SDK.
	// [Optional]
	// Read only
	ClientToken string `json:"client_token,omitempty"`

	// SessionId is the session ID.
	// [Optional]
	// Read only
	SessionId string `json:"session_id,omitempty"`
}

// KlarnaAccountHolder is the account holder details returned by Klarna after the shopper
// completes verification.
type KlarnaAccountHolder struct {
	// Name is the full name of the account holder.
	// [Optional]
	Name string `json:"name,omitempty"`
}

// StcpayPaymentMethod is the stc pay payment method's details and configuration.
type StcpayPaymentMethod struct {
	PaymentMethodBase

	// Otp is the one-time password (OTP) for stc pay.
	// [Optional]
	// Write only
	Otp string `json:"otp,omitempty"`

	// Action is the next available action for the payment method.
	// [Optional]
	// Read only
	Action *StcpayAction `json:"action,omitempty"`

	// PaymentMethodOptions holds the payment method options.
	// [Optional]
	//
	// Deprecated: not part of the API spec; the API no longer sends it and ignores it on
	// requests. Use Action instead. It will be removed in the next major version.
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}

// StcpayActionType is the type of the next available action for the stc pay payment method.
type StcpayActionType string

const (
	// StcpayActionTypeOtp asks for the one-time password the customer received.
	StcpayActionTypeOtp StcpayActionType = "otp"
)

// StcpayAction is the next available action for the stc pay payment method.
type StcpayAction struct {
	// Type is the type of action.
	// [Optional]
	// Read only
	Type StcpayActionType `json:"type,omitempty"`
}

// TabbyPaymentMethod is the Tabby payment method's details and configuration.
type TabbyPaymentMethod struct {
	PaymentMethodBase

	// PaymentTypes is the available payment types for Tabby (for example, "installments").
	// [Optional]
	PaymentTypes []string `json:"payment_types,omitempty"`

	// PaymentMethodOptions holds the payment method options.
	// [Optional]
	//
	// Deprecated: not part of the API spec; the API no longer sends it and ignores it on
	// requests. There is no replacement. It will be removed in the next major version.
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}

// BizumPaymentMethod is the Bizum payment method's details and configuration.
type BizumPaymentMethod struct {
	PaymentMethodBase

	// Initialization is the initialization state of the payment method. It shadows
	// PaymentMethodBase.Initialization so the deprecation is reported; a value set on either one is
	// sent, and a received value is set on both.
	// [Optional]
	//
	// Deprecated: Not part of the Bizum schema in the Payment Setups API specification; the API
	// ignores it on requests. There is no replacement. Will be removed in the next major version.
	Initialization PaymentMethodInitialization `json:"initialization,omitempty"`

	// PaymentMethodOptions holds the payment method options.
	// [Optional]
	//
	// Deprecated: Not part of the Bizum schema in the Payment Setups API specification; the API
	// ignores it on requests. There is no replacement. Will be removed in the next major version.
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}

// SimplePaymentMethod is the details and configuration of a payment method that carries only the
// shared payment method members: PayNow, EPS, Benefit, Vipps, Twint, MobilePay, MB WAY, WeChat
// Pay, Octopus, Alma, Sequra, QPay, Alipay CN, Alipay HK, GCash, TNG, Dana, KakaoPay, TrueMoney
// and Tamara.
//
// Initialization, promoted from PaymentMethodBase, is deprecated on this type except for Tamara,
// whose schema has it. Not part of the other payment method schemas in the Payment Setups API
// specification; the API ignores it on requests. There is no replacement. Will be removed in the
// next major version.
type SimplePaymentMethod struct {
	PaymentMethodBase

	// PaymentMethodOptions holds the payment method options.
	// [Optional]
	//
	// Deprecated: Not part of these payment method schemas in the Payment Setups API
	// specification; the API ignores it on requests. There is no replacement. Will be removed in
	// the next major version.
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}

// IdealPaymentMethod is the iDEAL payment method's details and configuration.
type IdealPaymentMethod struct {
	PaymentMethodBase

	// Initialization is the initialization state of the payment method. It shadows
	// PaymentMethodBase.Initialization so the deprecation is reported; a value set on either one is
	// sent, and a received value is set on both.
	// [Optional]
	//
	// Deprecated: Not part of the Ideal schema in the Payment Setups API specification; the API
	// ignores it on requests. There is no replacement. Will be removed in the next major version.
	Initialization PaymentMethodInitialization `json:"initialization,omitempty"`

	// Description is a description of the products or services being paid for. Do not include
	// HTML tags or special characters, as these are rejected by iDEAL.
	// [Optional]
	// max 35 characters
	Description string `json:"description,omitempty"`

	// PaymentMethodOptions holds the payment method options.
	// [Optional]
	//
	// Deprecated: Not part of the Ideal schema in the Payment Setups API specification; the API
	// ignores it on requests. There is no replacement. Will be removed in the next major version.
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}

// KnetLanguage is the customer's preferred language for KNET, as an ISO 639-1 code.
type KnetLanguage string

const (
	// KnetLanguageEn is English.
	KnetLanguageEn KnetLanguage = "en"
	// KnetLanguageAr is Arabic.
	KnetLanguageAr KnetLanguage = "ar"
)

// KnetPaymentMethod is the KNET payment method's details and configuration.
type KnetPaymentMethod struct {
	PaymentMethodBase

	// Initialization is the initialization state of the payment method. It shadows
	// PaymentMethodBase.Initialization so the deprecation is reported; a value set on either one is
	// sent, and a received value is set on both.
	// [Optional]
	//
	// Deprecated: Not part of the Knet schema in the Payment Setups API specification; the API
	// ignores it on requests. There is no replacement. Will be removed in the next major version.
	Initialization PaymentMethodInitialization `json:"initialization,omitempty"`

	// Language is the customer's preferred language, as an ISO 639-1 code. For example, the
	// language selected on your site, if the issuer's site supports it.
	// - ar: Arabic
	// - en: English
	// [Optional]
	Language KnetLanguage `json:"language,omitempty"`

	// PaymentMethodOptions holds the payment method options.
	// [Optional]
	//
	// Deprecated: Not part of the Knet schema in the Payment Setups API specification; the API
	// ignores it on requests. There is no replacement. Will be removed in the next major version.
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}

// BancontactPaymentMethod is the Bancontact payment method's details and configuration.
type BancontactPaymentMethod struct {
	PaymentMethodBase

	// Initialization is the initialization state of the payment method. It shadows
	// PaymentMethodBase.Initialization so the deprecation is reported; a value set on either one is
	// sent, and a received value is set on both.
	// [Optional]
	//
	// Deprecated: Not part of the Bancontact schema in the Payment Setups API specification; the
	// API ignores it on requests. There is no replacement. Will be removed in the next major
	// version.
	Initialization PaymentMethodInitialization `json:"initialization,omitempty"`

	// AccountHolderName is the account holder's name.
	// [Optional]
	// min 3 characters, max 100 characters
	AccountHolderName string `json:"account_holder_name,omitempty"`

	// PaymentMethodOptions holds the payment method options.
	// [Optional]
	//
	// Deprecated: Not part of the Bancontact schema in the Payment Setups API specification; the
	// API ignores it on requests. There is no replacement. Will be removed in the next major
	// version.
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}

// PaymentSetupScheme is a card scheme name, as listed in the schemes of a Payment Setup card,
// Apple Pay or Google Pay payment method.
type PaymentSetupScheme string

const (
	PaymentSetupSchemeAccel           PaymentSetupScheme = "accel"
	PaymentSetupSchemeAmex            PaymentSetupScheme = "amex"
	PaymentSetupSchemeCartesBancaires PaymentSetupScheme = "cartes_bancaires"
	PaymentSetupSchemeChinaUnionPay   PaymentSetupScheme = "china_union_pay"
	PaymentSetupSchemeDiners          PaymentSetupScheme = "diners"
	PaymentSetupSchemeDiscover        PaymentSetupScheme = "discover"
	PaymentSetupSchemeJcb             PaymentSetupScheme = "jcb"
	PaymentSetupSchemeMada            PaymentSetupScheme = "mada"
	PaymentSetupSchemeMaestro         PaymentSetupScheme = "maestro"
	PaymentSetupSchemeMastercard      PaymentSetupScheme = "mastercard"
	PaymentSetupSchemeNyce            PaymentSetupScheme = "nyce"
	PaymentSetupSchemeOmannet         PaymentSetupScheme = "omannet"
	PaymentSetupSchemePulse           PaymentSetupScheme = "pulse"
	PaymentSetupSchemeShazam          PaymentSetupScheme = "shazam"
	PaymentSetupSchemeStar            PaymentSetupScheme = "star"
	PaymentSetupSchemeUpi             PaymentSetupScheme = "upi"
	PaymentSetupSchemeVisa            PaymentSetupScheme = "visa"
)

// PaymentSetupCardAccountHolder is the account holder of a card, Apple Pay or Google Pay payment method.
type PaymentSetupCardAccountHolder struct {
	// Type is the card account holder type.
	// [Optional]
	Type common.AccountHolderType `json:"type,omitempty"`

	// FirstName is the card account holder's first name.
	// [Optional]
	FirstName string `json:"first_name,omitempty"`

	// LastName is the card account holder's last name.
	// [Optional]
	LastName string `json:"last_name,omitempty"`

	// MiddleName is the card account holder's middle name.
	// [Optional]
	MiddleName string `json:"middle_name,omitempty"`

	// CompanyName is the card account holder's company name.
	// [Optional]
	CompanyName string `json:"company_name,omitempty"`

	// AccountNameInquiry indicates whether the ANI (account name inquiry) check is performed with
	// the card scheme (Visa or Mastercard). A pointer so an explicit false is sent.
	// [Optional]
	AccountNameInquiry *bool `json:"account_name_inquiry,omitempty"`
}

// CardToken is the Checkout.com card token of the card payment method.
type CardToken struct {
	// Id is the Checkout.com card token identifier.
	// [Optional]
	Id string `json:"id,omitempty"`
}

// CardDetails holds the card details of the card payment method.
type CardDetails struct {
	// Number is the card number (without separators).
	// [Optional]
	// Write only
	Number string `json:"number,omitempty"`

	// Last4 is the last four digits of the card number.
	// [Optional]
	// Read only
	Last4 string `json:"last4,omitempty"`

	// Bin is the card issuer's Bank Identification Number (BIN).
	// [Optional]
	// Read only
	Bin string `json:"bin,omitempty"`

	// Scheme is the card scheme.
	// [Optional]
	// Read only
	Scheme string `json:"scheme,omitempty"`

	// ExpiryMonth is the expiry month of the card.
	// [Optional]
	ExpiryMonth int `json:"expiry_month,omitempty"`

	// ExpiryYear is the expiry year of the card.
	// [Optional]
	ExpiryYear int `json:"expiry_year,omitempty"`

	// Name is the cardholder's name.
	// [Optional]
	// Write only
	Name string `json:"name,omitempty"`

	// Cvv is the card verification value/code. 3 digits, except for American Express (4 digits).
	// [Optional]
	// Write only
	Cvv string `json:"cvv,omitempty"`

	// Stored indicates whether this card is being submitted from your own stored card-on-file
	// system, rather than being entered by the customer at the time of payment. Set to true when
	// you provide card details that you already have on file for the customer. A pointer so an
	// explicit false is sent.
	// [Optional]
	// Write only
	Stored *bool `json:"stored,omitempty"`

	// ExpiresOn is the time by which the card details must be confirmed.
	// [Optional]
	// Read only
	// Format: date-time (RFC 3339)
	ExpiresOn *time.Time `json:"expires_on,omitempty"`

	// AllowUpdate indicates whether to use the Real-Time Account Updater to update the card
	// information. A pointer so an explicit false is sent.
	// [Optional]
	AllowUpdate *bool `json:"allow_update,omitempty"`
}

// CardPaymentMethod is the card payment method's details and configuration.
type CardPaymentMethod struct {
	PaymentSetupPaymentMethod

	// Token is the Checkout.com card token to use instead of the raw card details.
	// [Optional]
	Token *CardToken `json:"token,omitempty"`

	// Details holds the card details.
	// [Optional]
	Details *CardDetails `json:"details,omitempty"`

	// StoreForFutureUse is set to true if you intend to reuse the payment credentials in subsequent
	// payments. If set to false, a payment instrument will not be included in the payment response.
	// A pointer so an explicit false is sent.
	// [Optional]
	// Write only
	StoreForFutureUse *bool `json:"store_for_future_use,omitempty"`

	// Phone is the customer's phone number.
	// [Optional]
	Phone *common.Phone `json:"phone,omitempty"`

	// AccountHolder holds the card account holder details.
	// [Optional]
	AccountHolder *PaymentSetupCardAccountHolder `json:"account_holder,omitempty"`

	// Schemes is the list of supported card scheme names.
	// [Optional]
	// Read only
	Schemes []PaymentSetupScheme `json:"schemes,omitempty"`

	// PreferredScheme is the preferred card scheme to process the payment.
	// [Optional]
	PreferredScheme payments.PreferredSchema `json:"preferred_scheme,omitempty"`
}

// WalletToken is the Checkout.com token of an Apple Pay or Google Pay payment method.
type WalletToken struct {
	// Id is the token identifier, prefixed by tok_.
	// [Required]
	// ^(tok)_(\w{26})$
	Id string `json:"id"`

	// ExpiresOn is the date and time the token expires.
	// [Optional]
	// Format: date-time (RFC 3339)
	ExpiresOn *time.Time `json:"expires_on,omitempty"`
}

// ApplePayTokenDataHeader holds additional version-dependent information used to decrypt and
// verify the payment.
type ApplePayTokenDataHeader struct {
	// EphemeralPublicKey is the ephemeral public key bytes, Base64-encoded.
	// [Optional]
	// Write only
	EphemeralPublicKey string `json:"ephemeral_public_key,omitempty"`

	// PublicKeyHash is the hash of the X.509 encoded public key bytes of the merchant's
	// certificate, Base64-encoded.
	// [Optional]
	// Write only
	PublicKeyHash string `json:"public_key_hash,omitempty"`

	// TransactionId is the transaction identifier, generated on the device, hexadecimal.
	// [Optional]
	// Write only
	TransactionId string `json:"transaction_id,omitempty"`
}

// ApplePayTokenData is the Apple Pay encrypted payment token data.
type ApplePayTokenData struct {
	// Version is the version of the payment token. The token uses EC_v1 for ECC-encrypted data,
	// and RSA_v1 for RSA-encrypted data.
	// [Optional]
	// Write only
	Version string `json:"version,omitempty"`

	// Data is the encrypted payment data, Base64-encoded.
	// [Optional]
	// Write only
	Data string `json:"data,omitempty"`

	// Signature is the signature of the payment and header data. The signature includes the
	// signing certificate, its intermediate CA certificate, and information about the signing
	// algorithm.
	// [Optional]
	// Write only
	Signature string `json:"signature,omitempty"`

	// Header holds additional version-dependent information used to decrypt and verify the payment.
	// [Optional]
	// Write only
	Header *ApplePayTokenDataHeader `json:"header,omitempty"`
}

// ApplePayPaymentMethod is the Apple Pay payment method's details and configuration.
type ApplePayPaymentMethod struct {
	PaymentSetupPaymentMethod

	// TokenData is the Apple Pay encrypted payment token data.
	// [Optional]
	// Write only
	TokenData *ApplePayTokenData `json:"token_data,omitempty"`

	// Token is the Checkout.com Apple Pay token.
	// [Optional]
	Token *WalletToken `json:"token,omitempty"`

	// StoreForFutureUse is set to true if you intend to reuse the payment credentials in subsequent
	// payments. If set to false, a payment instrument will not be included in the payment response.
	// A pointer so an explicit false is sent.
	// [Optional]
	// Write only
	StoreForFutureUse *bool `json:"store_for_future_use,omitempty"`

	// Phone is the customer's phone number.
	// [Optional]
	Phone *common.Phone `json:"phone,omitempty"`

	// AccountHolder holds the card account holder details.
	// [Optional]
	AccountHolder *PaymentSetupCardAccountHolder `json:"account_holder,omitempty"`

	// Schemes is the list of supported card scheme names.
	// [Optional]
	// Read only
	Schemes []PaymentSetupScheme `json:"schemes,omitempty"`

	// PreferredScheme is the preferred card scheme to process the payment.
	// [Optional]
	PreferredScheme payments.PreferredSchema `json:"preferred_scheme,omitempty"`
}

// GooglePayTokenData is the Google Pay encrypted payment token data.
type GooglePayTokenData struct {
	// ProtocolVersion is the encryption and signing scheme used to create this message. If not
	// set, the default is ECv0.
	// [Optional]
	// Write only
	ProtocolVersion string `json:"protocol_version,omitempty"`

	// Signature is the signature that verifies the message came from Google, created using ECDSA.
	// [Optional]
	// Write only
	Signature string `json:"signature,omitempty"`

	// SignedMessage is a serialized JSON string containing the encryptedMessage,
	// ephemeralPublicKey, and tag. To simplify the signature verification process, this value is
	// serialized.
	// [Optional]
	// Write only
	SignedMessage string `json:"signed_message,omitempty"`

	// TokenizationKey is the public key provided to the Google API in the
	// tokenizationSpecification.parameters.gatewayMerchantId field.
	// [Optional]
	// Write only
	TokenizationKey string `json:"tokenization_key,omitempty"`
}

// GooglePayPaymentMethod is the Google Pay payment method's details and configuration.
type GooglePayPaymentMethod struct {
	PaymentSetupPaymentMethod

	// TokenData is the Google Pay encrypted payment token data.
	// [Optional]
	// Write only
	TokenData *GooglePayTokenData `json:"token_data,omitempty"`

	// Token is the Checkout.com Google Pay token.
	// [Optional]
	Token *WalletToken `json:"token,omitempty"`

	// StoreForFutureUse is set to true if you intend to reuse the payment credentials in subsequent
	// payments. If set to false, a payment instrument will not be included in the payment response.
	// A pointer so an explicit false is sent.
	// [Optional]
	// Write only
	StoreForFutureUse *bool `json:"store_for_future_use,omitempty"`

	// Phone is the customer's phone number.
	// [Optional]
	Phone *common.Phone `json:"phone,omitempty"`

	// AccountHolder holds the card account holder details.
	// [Optional]
	AccountHolder *PaymentSetupCardAccountHolder `json:"account_holder,omitempty"`

	// Schemes is the list of supported card scheme names. Google Pay does not list cartes_bancaires.
	// [Optional]
	// Read only
	Schemes []PaymentSetupScheme `json:"schemes,omitempty"`

	// PreferredScheme is the preferred card scheme to process the payment.
	// [Optional]
	// Enum: "visa" "mastercard"
	PreferredScheme GooglePayPreferredScheme `json:"preferred_scheme,omitempty"`
}

// GooglePayPreferredScheme is a card scheme Google Pay accepts as the preferred scheme.
type GooglePayPreferredScheme string

const (
	GooglePayPreferredSchemeVisa       GooglePayPreferredScheme = "visa"
	GooglePayPreferredSchemeMastercard GooglePayPreferredScheme = "mastercard"
)

// InstrumentPaymentMethod is the instrument payment method's details and configuration.
type InstrumentPaymentMethod struct {
	PaymentMethodBase

	// Initialization is the initialization state of the payment method. It shadows
	// PaymentMethodBase.Initialization so the deprecation is reported; a value set on either one is
	// sent, and a received value is set on both.
	// [Optional]
	//
	// Deprecated: Not part of the PaymentSetupInstrument schema in the Payment Setups API
	// specification; the API ignores it on requests. There is no replacement. Will be removed in
	// the next major version.
	Initialization PaymentMethodInitialization `json:"initialization,omitempty"`

	// Id is the unique identifier of the selected instrument.
	// [Optional]
	Id string `json:"id,omitempty"`

	// Type is the type of the instrument, for example "bacs".
	// [Optional]
	// Read only
	Type string `json:"type,omitempty"`

	// Action is the next available action for the payment method.
	// [Optional]
	// Read only
	Action *InstrumentAction `json:"action,omitempty"`

	// PaymentMethodOptions holds the payment method options.
	// [Optional]
	//
	// Deprecated: Not part of the PaymentSetupInstrument schema in the Payment Setups API
	// specification; the API ignores it on requests. There is no replacement. Will be removed in
	// the next major version.
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}

// InstrumentActionType is the type of the next available action for the instrument payment method.
type InstrumentActionType string

const (
	// InstrumentActionTypeSetInstrument sets the instrument on the Payment Setup.
	InstrumentActionTypeSetInstrument InstrumentActionType = "set_instrument"
)

// InstrumentAction is the next available action for the instrument payment method.
type InstrumentAction struct {
	// Type is the type of action.
	// [Optional]
	// Read only
	Type InstrumentActionType `json:"type,omitempty"`
}

// MultibancoPaymentMethod is the Multibanco payment method's details and configuration.
type MultibancoPaymentMethod struct {
	PaymentMethodBase

	// Initialization is the initialization state of the payment method. It shadows
	// PaymentMethodBase.Initialization so the deprecation is reported; a value set on either one is
	// sent, and a received value is set on both.
	// [Optional]
	//
	// Deprecated: Not part of the Multibanco schema in the Payment Setups API specification; the
	// API ignores it on requests. There is no replacement. Will be removed in the next major
	// version.
	Initialization PaymentMethodInitialization `json:"initialization,omitempty"`

	// AccountHolderName is the account holder's name.
	// [Optional]
	// min 3 characters, max 100 characters
	AccountHolderName string `json:"account_holder_name,omitempty"`

	// PaymentMethodOptions holds the payment method options.
	// [Optional]
	//
	// Deprecated: Not part of the Multibanco schema in the Payment Setups API specification; the API
	// ignores it on requests. There is no replacement. Will be removed in the next major version.
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}

// P24AccountHolder is the P24 account holder's details.
type P24AccountHolder struct {
	// Name is the account holder's name.
	// [Optional]
	// min 3 characters, max 100 characters
	Name string `json:"name,omitempty"`

	// Email is the account holder's email address.
	// [Optional]
	// Format: email
	// max 254 characters
	Email string `json:"email,omitempty"`
}

// P24PaymentMethod is the P24 (Przelewy24) payment method's details and configuration.
type P24PaymentMethod struct {
	PaymentMethodBase

	// Initialization is the initialization state of the payment method. It shadows
	// PaymentMethodBase.Initialization so the deprecation is reported; a value set on either one is
	// sent, and a received value is set on both.
	// [Optional]
	//
	// Deprecated: Not part of the P24 schema in the Payment Setups API specification; the API
	// ignores it on requests. There is no replacement. Will be removed in the next major version.
	Initialization PaymentMethodInitialization `json:"initialization,omitempty"`

	// AccountHolder is the account holder's details.
	// [Optional]
	AccountHolder *P24AccountHolder `json:"account_holder,omitempty"`

	// PaymentMethodOptions holds the payment method options.
	// [Optional]
	//
	// Deprecated: Not part of the P24 schema in the Payment Setups API specification; the API
	// ignores it on requests. There is no replacement. Will be removed in the next major version.
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}

// SwishAccountHolder is the Swish account holder's details.
type SwishAccountHolder struct {
	// FirstName is the account holder's first name.
	// [Optional]
	// min 1 characters
	FirstName string `json:"first_name,omitempty"`

	// LastName is the account holder's last name.
	// [Optional]
	// min 1 characters
	LastName string `json:"last_name,omitempty"`
}

// SwishPaymentMethod is the Swish payment method's details and configuration.
type SwishPaymentMethod struct {
	PaymentMethodBase

	// Initialization is the initialization state of the payment method. It shadows
	// PaymentMethodBase.Initialization so the deprecation is reported; a value set on either one is
	// sent, and a received value is set on both.
	// [Optional]
	//
	// Deprecated: Not part of the Swish schema in the Payment Setups API specification; the API
	// ignores it on requests. There is no replacement. Will be removed in the next major version.
	Initialization PaymentMethodInitialization `json:"initialization,omitempty"`

	// BillingDescriptor is a description that appears on the customer's billing statement.
	// [Optional]
	BillingDescriptor string `json:"billing_descriptor,omitempty"`

	// AccountHolder is the account holder's details.
	// [Optional]
	AccountHolder *SwishAccountHolder `json:"account_holder,omitempty"`

	// PaymentMethodOptions holds the payment method options.
	// [Optional]
	//
	// Deprecated: Not part of the Swish schema in the Payment Setups API specification; the API
	// ignores it on requests. There is no replacement. Will be removed in the next major version.
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}

// AchAccountHolder is the ACH account holder details.
type AchAccountHolder struct {
	// Type is the type of account holder.
	// [Optional]
	// Enum: "individual" "corporate" "government"
	Type string `json:"type,omitempty"`

	// FirstName is the first name of the account holder.
	// [Optional]
	FirstName string `json:"first_name,omitempty"`

	// LastName is the last name of the account holder.
	// [Optional]
	LastName string `json:"last_name,omitempty"`

	// CompanyName is the legal name of a registered company that holds the account.
	// [Optional]
	CompanyName string `json:"company_name,omitempty"`

	// DateOfBirth is the account holder's date of birth. A plain string, as the specification
	// gives it no format.
	// [Optional]
	DateOfBirth string `json:"date_of_birth,omitempty"`

	// Identification is the account holder's government document identification, for example a
	// Social Security Number (SSN).
	// [Optional]
	// Write only
	Identification *AchAccountHolderIdentification `json:"identification,omitempty"`
}

// AchAccountHolderIdentification is the ACH account holder's government document identification,
// for example a Social Security Number (SSN).
type AchAccountHolderIdentification struct {
	// Type is the document type.
	// [Optional]
	Type string `json:"type,omitempty"`

	// IssuingCountry is the country where the document was issued.
	// [Optional]
	IssuingCountry string `json:"issuing_country,omitempty"`

	// Number is the document number.
	// [Optional]
	Number string `json:"number,omitempty"`
}

// AchPaymentMethod is the ACH payment method's details and configuration.
type AchPaymentMethod struct {
	PaymentMethodBase

	// Initialization is the initialization state of the payment method. It shadows
	// PaymentMethodBase.Initialization so the deprecation is reported; a value set on either one is
	// sent, and a received value is set on both.
	// [Optional]
	//
	// Deprecated: Not part of the Ach schema in the Payment Setups API specification; the API
	// ignores it on requests. There is no replacement. Will be removed in the next major version.
	Initialization PaymentMethodInitialization `json:"initialization,omitempty"`

	// AccountType is the type of Direct Debit account.
	// [Optional]
	// Enum: "savings" "current" "cash"
	AccountType string `json:"account_type,omitempty"`

	// AccountHolder is the account holder details.
	// [Optional]
	AccountHolder *AchAccountHolder `json:"account_holder,omitempty"`

	// AccountNumber is the account number of the Direct Debit account.
	// [Optional]
	// Write only
	// min 4 characters, max 17 characters
	AccountNumber string `json:"account_number,omitempty"`

	// BankCode is the bank code of the Direct Debit account.
	// [Optional]
	// Write only
	// min 8 characters, max 9 characters
	BankCode string `json:"bank_code,omitempty"`

	// Country is the two-letter ISO country code of the bank account.
	// [Optional]
	// min 2 characters, max 2 characters
	Country common.Country `json:"country,omitempty"`

	// PaymentMethodOptions holds the payment method options.
	// [Optional]
	//
	// Deprecated: Not part of the Ach schema in the Payment Setups API specification; the API
	// ignores it on requests. There is no replacement. Will be removed in the next major version.
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}

// SepaAccountHolder is the SEPA account holder details.
type SepaAccountHolder struct {
	// Type is the type of account holder.
	// [Optional]
	// Enum: "individual" "corporate"
	Type string `json:"type,omitempty"`

	// FirstName is the first name of the account holder.
	// [Optional]
	FirstName string `json:"first_name,omitempty"`

	// LastName is the last name of the account holder.
	// [Optional]
	LastName string `json:"last_name,omitempty"`

	// CompanyName is the legal name of a registered company that holds the account.
	// [Optional]
	CompanyName string `json:"company_name,omitempty"`
}

// SepaMandateType is the type of a SEPA mandate.
type SepaMandateType string

const (
	SepaMandateTypeCore SepaMandateType = "core"
	SepaMandateTypeB2b  SepaMandateType = "b2b"
)

// SepaMandate holds the SEPA mandate details.
type SepaMandate struct {
	// Id is the ID of the mandate.
	// [Optional]
	Id string `json:"id,omitempty"`

	// Type is the type of mandate.
	// [Optional]
	Type SepaMandateType `json:"type,omitempty"`

	// DateOfSignature is the date the mandate was signed.
	// [Optional]
	// Format: yyyy-MM-dd
	DateOfSignature *common.APIShortDate `json:"date_of_signature,omitempty"`
}

// SepaPaymentMethod is the SEPA payment method's details and configuration.
type SepaPaymentMethod struct {
	PaymentMethodBase

	// Initialization is the initialization state of the payment method. It shadows
	// PaymentMethodBase.Initialization so the deprecation is reported; a value set on either one is
	// sent, and a received value is set on both.
	// [Optional]
	//
	// Deprecated: Not part of the Sepa schema in the Payment Setups API specification; the API
	// ignores it on requests. There is no replacement. Will be removed in the next major version.
	Initialization PaymentMethodInitialization `json:"initialization,omitempty"`

	// AccountHolder is the account holder details.
	// [Optional]
	AccountHolder *SepaAccountHolder `json:"account_holder,omitempty"`

	// AccountNumber is the account holder's IBAN.
	// [Optional]
	// Write only
	AccountNumber string `json:"account_number,omitempty"`

	// Country is the account's country, as an ISO 3166-1 alpha-2 code.
	// [Optional]
	// min 2 characters, max 2 characters
	Country common.Country `json:"country,omitempty"`

	// Currency is the account holder's account currency.
	// [Optional]
	Currency string `json:"currency,omitempty"`

	// Mandate holds the mandate details.
	// [Optional]
	Mandate *SepaMandate `json:"mandate,omitempty"`

	// PaymentMethodOptions holds the payment method options.
	// [Optional]
	//
	// Deprecated: Not part of the Sepa schema in the Payment Setups API specification; the API
	// ignores it on requests. There is no replacement. Will be removed in the next major version.
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}

// PaypalShippingPreference is where PayPal obtains the shipping information.
type PaypalShippingPreference string

const (
	// PaypalShippingPreferenceNoShipping redacts the shipping address from the PayPal pages.
	PaypalShippingPreferenceNoShipping PaypalShippingPreference = "no_shipping"
	// PaypalShippingPreferenceGetFromFile obtains the shipping address from the customer's PayPal
	// account.
	PaypalShippingPreferenceGetFromFile PaypalShippingPreference = "get_from_file"
	// PaypalShippingPreferenceSetProvidedAddress uses the shipping address provided by the merchant.
	PaypalShippingPreferenceSetProvidedAddress PaypalShippingPreference = "set_provided_address"
)

// PaypalPaymentMethod is the PayPal payment method's details and configuration.
type PaypalPaymentMethod struct {
	PaymentMethodBase

	// UserAction is the user action for the PayPal widget.
	// - pay_now: after the customer clicks the PayPal button, the customer is redirected to a page
	// to enter payment details and immediately directed to finalize the payment.
	// - continue: after the customer clicks the PayPal button, the customer is redirected to a
	// page to enter payment details and then redirected back to the merchant's site to review and
	// finalize the payment.
	// [Optional]
	// Enum: "pay_now" "continue"
	UserAction string `json:"user_action,omitempty"`

	// BrandName is the brand name to display in the PayPal checkout experience.
	// [Optional]
	// max 127 characters
	BrandName string `json:"brand_name,omitempty"`

	// ShippingPreference is where to obtain the shipping information.
	// - no_shipping: redacts the shipping address from the PayPal pages.
	// - get_from_file: obtains the shipping address from the customer's PayPal account.
	// - set_provided_address: uses the shipping address provided by the merchant.
	// [Optional]
	ShippingPreference PaypalShippingPreference `json:"shipping_preference,omitempty"`

	// Action is the next available action for the payment method.
	// [Optional]
	// Read only
	Action *PaypalAction `json:"action,omitempty"`

	// PaymentMethodOptions holds the payment method options.
	// [Optional]
	//
	// Deprecated: not part of the API spec; the API no longer sends it and ignores it on
	// requests. Use Action instead. It will be removed in the next major version.
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}

// PaypalActionType is the type of the next available action for the PayPal payment method.
type PaypalActionType string

const (
	// PaypalActionTypeSdk completes the payment with the PayPal SDK, using the order_id.
	PaypalActionTypeSdk PaypalActionType = "sdk"
)

// PaypalAction is the next available action for the PayPal payment method.
type PaypalAction struct {
	// Type is the type of action.
	// [Optional]
	// Read only
	Type PaypalActionType `json:"type,omitempty"`

	// OrderId is the PayPal order ID to use with the PayPal SDK.
	// [Optional]
	// Read only
	OrderId string `json:"order_id,omitempty"`
}

// BacsAccountHolderType is the type of Bacs account holder.
type BacsAccountHolderType string

const (
	BacsAccountHolderIndividual BacsAccountHolderType = "individual"
	BacsAccountHolderCorporate  BacsAccountHolderType = "corporate"
)

// BacsAccountHolder is the Bacs account holder details.
type BacsAccountHolder struct {
	// Type is the type of account holder.
	// [Optional]
	Type BacsAccountHolderType `json:"type,omitempty"`

	// FirstName is the first name of the account holder.
	// [Optional]
	FirstName string `json:"first_name,omitempty"`

	// LastName is the last name of the account holder.
	// [Optional]
	LastName string `json:"last_name,omitempty"`

	// CompanyName is the legal name of a registered company that holds the account.
	// [Optional]
	CompanyName string `json:"company_name,omitempty"`

	// Email is the email address of the account holder.
	// [Optional]
	Email string `json:"email,omitempty"`
}

// BacsPaymentMethod is the Bacs payment method's details and configuration.
type BacsPaymentMethod struct {
	PaymentMethodBase

	// InstrumentId is the ID of the Bacs instrument used for the payment.
	// [Optional]
	// Read only
	InstrumentId string `json:"instrument_id,omitempty"`

	// AccountHolder is the account holder details.
	// [Optional]
	AccountHolder *BacsAccountHolder `json:"account_holder,omitempty"`

	// AccountNumber is the account number of the Bacs Direct Debit account.
	// [Optional]
	// Write only
	AccountNumber string `json:"account_number,omitempty"`

	// BankCode is the sort code of the Bacs Direct Debit account.
	// [Optional]
	// Write only
	BankCode string `json:"bank_code,omitempty"`

	// Country is the account's country, as an ISO 3166-1 alpha-2 code.
	// [Optional]
	// min 2 characters, max 2 characters
	Country common.Country `json:"country,omitempty"`

	// Currency is the account holder's account currency.
	// [Optional]
	Currency string `json:"currency,omitempty"`

	// AllowPartialMatch indicates whether the Bacs instrument is created when account validation
	// returns a partial match. When true, the instrument is created on a partial match; when
	// false, instrument creation fails on a partial match.
	// [Optional]
	// Default: false
	AllowPartialMatch bool `json:"allow_partial_match,omitempty"`

	// PaymentMethodOptions holds the payment method options.
	// [Optional]
	//
	// Deprecated: Not part of the Bacs schema in the Payment Setups API specification; the API
	// ignores it on requests. There is no replacement. Will be removed in the next major version.
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}

// CardPresentPin holds the PIN details of a card present payment. Not part of the current
// Payment Setup schema.
type CardPresentPin struct {
	// KeySetId is the identifier of the key set used to encrypt the PIN block.
	// [Optional]
	KeySetId string `json:"key_set_id,omitempty"`

	// Block is the encrypted PIN block.
	// [Optional]
	Block string `json:"block,omitempty"`

	// BlockFormat is the format of the PIN block.
	// [Optional]
	BlockFormat string `json:"block_format,omitempty"`
}

// CardPresentPaymentMethod is the Card Present payment method's details and configuration. Not
// part of the current Payment Setup schema.
type CardPresentPaymentMethod struct {
	PaymentMethodBase

	// Track2 is the card's track 2 data.
	// [Optional]
	Track2 string `json:"track2,omitempty"`

	// Emv is the card's EMV data.
	// [Optional]
	Emv string `json:"emv,omitempty"`

	// EntryMode is the mode in which the card details were entered.
	// [Optional]
	EntryMode string `json:"entry_mode,omitempty"`

	// Pin holds the PIN details.
	// [Optional]
	Pin *CardPresentPin `json:"pin,omitempty"`

	// StoreForFutureUse indicates whether to store the payment credentials for future use.
	// [Optional]
	StoreForFutureUse bool `json:"store_for_future_use,omitempty"`

	// Name is the cardholder's name.
	// [Optional]
	Name string `json:"name,omitempty"`
}

// PayByBankBank is a bank available for the customer to select.
type PayByBankBank struct {
	// BankId is the unique identifier of the bank.
	// [Optional]
	BankId string `json:"bank_id,omitempty"`

	// DisplayName is the display name of the bank.
	// [Optional]
	DisplayName string `json:"display_name,omitempty"`

	// LogoUrl is the URL of the bank's logo.
	// [Optional]
	LogoUrl string `json:"logo_url,omitempty"`

	// Available is whether the bank is currently available for selection.
	// [Optional]
	Available bool `json:"available,omitempty"`
}

// PayByBankAction is the next available action for the Pay by Bank payment method.
type PayByBankAction struct {
	// Type is the type of action.
	// [Optional]
	// Enum: "select_bank"
	Type string `json:"type,omitempty"`

	// Banks is the list of banks available for the customer to select.
	// [Optional]
	Banks []PayByBankBank `json:"banks,omitempty"`
}

// PayByBankPaymentMethod is the Pay by Bank (Open Banking) payment method's details and configuration.
type PayByBankPaymentMethod struct {
	PaymentMethodBase

	// Initialization is the initialization state of the payment method. It shadows
	// PaymentMethodBase.Initialization so the deprecation is reported; a value set on either one is
	// sent, and a received value is set on both.
	// [Optional]
	//
	// Deprecated: Not part of the PayByBank schema in the Payment Setups API specification; the API
	// ignores it on requests. There is no replacement. Will be removed in the next major version.
	Initialization PaymentMethodInitialization `json:"initialization,omitempty"`

	// BankId is the identifier of the bank the customer has selected for the payment.
	// [Optional]
	BankId string `json:"bank_id,omitempty"`

	// Action is the next available action for the payment method.
	// [Optional]
	// Read only
	Action *PayByBankAction `json:"action,omitempty"`
}

// CashAppPaymentMethod is the Cash App payment method's details and configuration (schema CashApp).
type CashAppPaymentMethod struct {
	PaymentMethodBase

	// CustomerProfileSharing indicates whether the customer consents to share their Cash App
	// customer profile with Checkout.com. A pointer so an explicit false is sent.
	// [Optional]
	CustomerProfileSharing *bool `json:"customer_profile_sharing,omitempty"`

	// CustomerProfile is the customer's Cash App profile that they consented to share. Included in
	// the response when customer_profile_sharing is enabled.
	//
	// Cash App releases this profile only once. It's present in the first successful response when
	// you get the Payment Setup after the customer authorizes the payment. Every subsequent
	// response omits it.
	// [Optional]
	// Read only
	CustomerProfile *CashAppCustomerProfile `json:"customer_profile,omitempty"`

	// Reference is a reference for the Cash App Pay transaction, returned by the provider.
	// [Optional]
	// Read only
	// max 80 characters
	Reference string `json:"reference,omitempty"`

	// Action is the next available action for the payment method.
	// [Optional]
	// Read only
	Action *CashAppAction `json:"action,omitempty"`
}

// CashAppActionType is the type of the next available action for the Cash App payment method.
type CashAppActionType string

const (
	// CashAppActionTypeRedirect redirects the customer to Cash App to authorize the payment.
	CashAppActionTypeRedirect CashAppActionType = "redirect"
)

// CashAppAction is the next available action for the Cash App payment method.
type CashAppAction struct {
	// Type is the type of action.
	// [Optional]
	// Read only
	Type CashAppActionType `json:"type,omitempty"`

	// RedirectUrl is the URL to redirect the customer to so they can authorize the payment with
	// Cash App.
	// [Optional]
	// Read only
	// Format: uri
	RedirectUrl string `json:"redirect_url,omitempty"`
}

// CashAppCustomerProfile is the customer's Cash App profile that they consented to share.
type CashAppCustomerProfile struct {
	// CustomerId is Cash App's identifier for the customer. This is not a Checkout.com customer
	// identifier.
	// [Optional]
	// Read only
	CustomerId string `json:"customer_id,omitempty"`

	// Cashtag is the customer's $Cashtag.
	// [Optional]
	// Read only
	Cashtag string `json:"cashtag,omitempty"`

	// ReferenceId is Cash App's reference for the customer profile.
	// [Optional]
	// Read only
	ReferenceId string `json:"reference_id,omitempty"`

	// FullName is the customer's full name.
	// [Optional]
	// Read only
	FullName string `json:"full_name,omitempty"`

	// GivenName is the customer's given name.
	// [Optional]
	// Read only
	GivenName string `json:"given_name,omitempty"`

	// MiddleName is the customer's middle name.
	// [Optional]
	// Read only
	MiddleName string `json:"middle_name,omitempty"`

	// FamilyName is the customer's family name.
	// [Optional]
	// Read only
	FamilyName string `json:"family_name,omitempty"`

	// Suffix is the suffix of the customer's name.
	// [Optional]
	// Read only
	Suffix string `json:"suffix,omitempty"`

	// BirthDate is the customer's date of birth. Kept as a string because the provider's format
	// varies (the documented example is a date-time).
	// [Optional]
	// Read only
	// Format: date
	BirthDate string `json:"birth_date,omitempty"`

	// Address is the customer's address.
	// [Optional]
	// Read only
	Address *CashAppAddress `json:"address,omitempty"`

	// PhoneNumber is the customer's phone number.
	// [Optional]
	// Read only
	PhoneNumber string `json:"phone_number,omitempty"`

	// EmailAddress is the customer's email address.
	// [Optional]
	// Read only
	EmailAddress string `json:"email_address,omitempty"`

	// CustomerSince is the date and time the customer's Cash App account was created.
	// [Optional]
	// Read only
	// Format: date-time (RFC 3339)
	CustomerSince *time.Time `json:"customer_since,omitempty"`
}

// CashAppAddress is the customer's address in a Cash App customer profile. It uses Cash App's
// field names, which differ from the Checkout.com common address.
type CashAppAddress struct {
	// AddressLine1 is the first line of the address.
	// [Optional]
	// Read only
	AddressLine1 string `json:"address_line_1,omitempty"`

	// AddressLine2 is the second line of the address.
	// [Optional]
	// Read only
	AddressLine2 string `json:"address_line_2,omitempty"`

	// AddressLine3 is the third line of the address.
	// [Optional]
	// Read only
	AddressLine3 string `json:"address_line_3,omitempty"`

	// Locality is the address locality, such as the city or town.
	// [Optional]
	// Read only
	Locality string `json:"locality,omitempty"`

	// Sublocality is the address sublocality, such as the district or neighborhood.
	// [Optional]
	// Read only
	Sublocality string `json:"sublocality,omitempty"`

	// AdministrativeDistrictLevel1 is the address's top-level administrative district, such as the
	// state or province.
	// [Optional]
	// Read only
	AdministrativeDistrictLevel1 string `json:"administrative_district_level_1,omitempty"`

	// PostalCode is the postal or zip code.
	// [Optional]
	// Read only
	PostalCode string `json:"postal_code,omitempty"`

	// Country is the address country, in ISO 3166-1 alpha-2 format.
	// [Optional]
	// Read only
	// max 2 characters
	Country common.Country `json:"country,omitempty"`
}

// StablecoinPaymentMethod is the Stablecoin payment method's details and configuration.
type StablecoinPaymentMethod struct {
	PaymentMethodBase

	// Initialization is the initialization state of the payment method. It shadows
	// PaymentMethodBase.Initialization so the deprecation is reported; a value set on either one is
	// sent, and a received value is set on both.
	// [Optional]
	//
	// Deprecated: Not part of the Stablecoin schema in the Payment Setups API specification; the
	// API ignores it on requests. There is no replacement. Will be removed in the next major
	// version.
	Initialization PaymentMethodInitialization `json:"initialization,omitempty"`
}

// ===== Support Structs =====

// PaymentSetupSettings holds the settings for the Payment Setup.
type PaymentSetupSettings struct {
	// SuccessUrl is the URL to redirect the customer to, if the payment is successful. For payment
	// methods with a redirect, this value overrides the default success redirect URL configured
	// on your account.
	// [Optional]
	// Format: uri
	// max 255 characters
	SuccessUrl string `json:"success_url,omitempty"`

	// FailureUrl is the URL to redirect the customer to, if the payment is unsuccessful. For
	// payment methods with a redirect, this value overrides the default failure redirect URL
	// configured on your account.
	// [Optional]
	// Format: uri
	// max 255 characters
	FailureUrl string `json:"failure_url,omitempty"`

	// Capture indicates whether to capture the payment immediately.
	// [Optional]
	// Default: true
	//
	// A pointer so an explicit false reaches the API. With a plain bool and omitempty, false was
	// dropped from the payload and the API applied its default of true, capturing the payment.
	Capture *bool `json:"capture,omitempty"`

	// ExcludedPaymentMethods is the list of payment methods excluded from the Payment Setup.
	// [Optional]
	ExcludedPaymentMethods []string `json:"excluded_payment_methods,omitempty"`
}

// PaymentSetupOrder holds the customer's order details.
type PaymentSetupOrder struct {
	// Items is a list of items in the order.
	// [Optional]
	Items []payments.Product `json:"items,omitempty"`

	// Shipping holds the customer's shipping details.
	// [Optional]
	Shipping *payments.ShippingDetails `json:"shipping,omitempty"`

	// SubMerchants holds the details of the sub-merchants.
	// [Optional]
	SubMerchants []OrderSubMerchant `json:"sub_merchants,omitempty"`

	// InvoiceId is the unique identifier for the invoice.
	// [Optional]
	InvoiceId string `json:"invoice_id,omitempty"`

	// ShippingAmount is the total shipping amount for the order.
	// [Optional]
	// min 0
	ShippingAmount int `json:"shipping_amount,omitempty"`

	// DiscountAmount is the discount amount the merchant applied to the transaction.
	// [Optional]
	// min 0
	DiscountAmount int `json:"discount_amount,omitempty"`

	// SurchargeAmount is the total surcharge amount for the order.
	// [Optional]
	// Format: int64
	// min 0
	SurchargeAmount int `json:"surcharge_amount,omitempty"`

	// TaxAmount is the total tax amount for the order.
	// [Optional]
	// min 0
	TaxAmount int `json:"tax_amount,omitempty"`

	// TippingAmount is the total tipping amount for the order.
	// [Optional]
	// Format: int64
	// min 0
	TippingAmount int `json:"tipping_amount,omitempty"`

	// AmountAllocations holds the sub-entities that the payment is being processed on behalf of.
	// [Optional]
	AmountAllocations []PaymentSetupAmountAllocation `json:"amount_allocations,omitempty"`
}

// AmountAllocationCommission is the commission to collect from an amount allocation split.
type AmountAllocationCommission struct {
	// Amount is the optional fixed amount of commission to collect. The amount must be provided
	// in the minor currency unit.
	// [Optional]
	// Format: int64
	// min 0
	Amount int64 `json:"amount,omitempty"`

	// Percentage is the optional percentage of commission to collect. Supports up to 8 decimal places.
	// [Optional]
	// min 0, max 100
	Percentage float64 `json:"percentage,omitempty"`
}

// PaymentSetupAmountAllocation represents a sub-entity on whose behalf the payment is processed.
type PaymentSetupAmountAllocation struct {
	// Id is the id of the sub-entity.
	// [Required]
	Id string `json:"id,omitempty"`

	// Amount is the split amount, credited to your sub-entity's currency account. The sum of all
	// split amounts must be equal to the payment amount. The amount must be provided in the minor
	// currency unit.
	// [Required]
	// Format: int64
	// min 0, max 9999999999
	Amount int64 `json:"amount,omitempty"`

	// Reference is a reference you can later use to identify this split, such as an order number.
	// [Optional]
	// max 50 characters
	Reference string `json:"reference,omitempty"`

	// Commission is the commission you'd like to collect from this split, credited to your
	// currency account. The commission cannot exceed the split amount.
	// Commission = (amount * commission.percentage) + commission.amount
	// [Optional]
	Commission *AmountAllocationCommission `json:"commission,omitempty"`
}

// OrderSubMerchant holds the details of a sub-merchant.
type OrderSubMerchant struct {
	// Id is the unique identifier for the sub-merchant.
	// [Optional]
	Id string `json:"id,omitempty"`

	// ProductCategory is the product category for the sub-merchant.
	// [Optional]
	ProductCategory string `json:"product_category,omitempty"`

	// NumberOfSales is the number of orders the sub-merchant has processed.
	// [Optional]
	NumberOfSales int `json:"number_of_sales,omitempty"`

	// RegistrationDate is the date the sub-merchant was registered.
	// [Optional]
	// Format: yyyy-MM-dd
	RegistrationDate *common.APIShortDate `json:"registration_date,omitempty"`
}

// PaymentSetupIndustry holds industry-specific information for the payment setup.
type PaymentSetupIndustry struct {
	// Airline holds the airline booking details.
	// [Optional]
	Airline []PaymentSetupAirline `json:"airline,omitempty"`

	// Accommodation holds the accommodation or cruise booking details.
	// [Optional]
	Accommodation []PaymentSetupAccommodation `json:"accommodation,omitempty"`
}

// PaymentSetupAirline holds details about the airline ticket and flights the customer booked.
type PaymentSetupAirline struct {
	// Ticket holds the details about the airline ticket.
	// [Optional]
	Ticket *PaymentSetupAirlineTicket `json:"ticket,omitempty"`

	// Passengers is the list of passengers on the flight.
	// [Optional]
	Passengers []PaymentSetupAirlinePassenger `json:"passengers,omitempty"`

	// FlightLegDetails is the list of flight legs booked by the customer.
	// [Optional]
	FlightLegDetails []PaymentSetupFlightLegDetails `json:"flight_leg_details,omitempty"`

	// TotalNumberOfPassengers is the total number of passengers on the booking.
	// [Optional]
	TotalNumberOfPassengers int `json:"total_number_of_passengers,omitempty"`

	// TravelType is the type of travel. For example, domestic or international.
	// [Optional]
	TravelType string `json:"travel_type,omitempty"`

	// TripType is the type of trip. For example, one_way or round_trip.
	// [Optional]
	TripType string `json:"trip_type,omitempty"`

	// Refundable specifies whether the booking is refundable.
	// [Optional]
	Refundable *bool `json:"refundable,omitempty"`

	// DeliveryRecipient is the recipient the ticket is delivered to.
	// [Optional]
	DeliveryRecipient string `json:"delivery_recipient,omitempty"`

	// Ancillaries is any additional add-ons purchased with the booking. For example, extra_baggage.
	// [Optional]
	Ancillaries string `json:"ancillaries,omitempty"`

	// Insurance holds details about the travel insurance purchased with the booking.
	// [Optional]
	Insurance *PaymentSetupAirlineInsurance `json:"insurance,omitempty"`
}

// PaymentSetupAirlineTicket holds details about the airline ticket.
type PaymentSetupAirlineTicket struct {
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

	// TravelPackageIndicator is the travel package indicator, as one of the following codes:
	// - A: airline flight reservation
	// - B: car rental and airline reservation
	// - C: car rental reservation
	// - N: unknown
	// [Optional]
	TravelPackageIndicator string `json:"travel_package_indicator,omitempty"`

	// TravelAgencyName is the name of the travel agency.
	// [Optional]
	TravelAgencyName string `json:"travel_agency_name,omitempty"`

	// TravelAgencyCode is the IATA or ARC unique identifier for the travel agency that issues the ticket.
	// [Optional]
	TravelAgencyCode string `json:"travel_agency_code,omitempty"`
}

// PaymentSetupAirlinePassenger is a passenger on the flight.
type PaymentSetupAirlinePassenger struct {
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

	// Address holds details about the passenger's address.
	// [Optional]
	Address *PaymentSetupAirlinePassengerAddress `json:"address,omitempty"`
}

// PaymentSetupAirlinePassengerAddress holds the passenger's country of residence.
type PaymentSetupAirlinePassengerAddress struct {
	// Country is the two-letter ISO country code of the passenger's country of residence.
	// [Optional]
	Country common.Country `json:"country,omitempty"`
}

// PaymentSetupFlightLegDetails is a flight leg booked by the customer.
type PaymentSetupFlightLegDetails struct {
	// FlightNumber is the flight identifier.
	// [Optional]
	FlightNumber string `json:"flight_number,omitempty"`

	// CarrierCode is the IATA 2-letter accounting code (PAX) that identifies the carrier. Required
	// if the airline data includes leg details.
	// [Optional]
	CarrierCode string `json:"carrier_code,omitempty"`

	// ClassOfTravelling is a one-letter identifier for the travel class. For example:
	// - F: first class
	// - J: business class
	// - W: premium economy class
	// - Y: economy class
	// [Optional]
	ClassOfTravelling string `json:"class_of_travelling,omitempty"`

	// DepartureAirport is the IATA three-letter airport code for the departure airport. Required
	// if the airline data includes leg details.
	// [Optional]
	DepartureAirport string `json:"departure_airport,omitempty"`

	// DepartureDate is the date of the scheduled take-off.
	// [Optional]
	// Format: yyyy-MM-dd
	DepartureDate *common.APIShortDate `json:"departure_date,omitempty"`

	// DepartureTime is the time of the scheduled take-off.
	// [Optional]
	DepartureTime string `json:"departure_time,omitempty"`

	// ArrivalAirport is the IATA three-letter airport code for the destination airport. Required
	// if the airline data includes leg details.
	// [Optional]
	ArrivalAirport string `json:"arrival_airport,omitempty"`

	// StopOverCode is a one-letter code that indicates whether the passenger is entitled to make a
	// stopover. Specify O or a blank space if the passenger is entitled. Specify X if the
	// passenger is not entitled.
	// [Optional]
	StopOverCode string `json:"stop_over_code,omitempty"`

	// FareBasisCode is the alphanumeric fare basis code.
	// [Optional]
	FareBasisCode string `json:"fare_basis_code,omitempty"`
}

// PaymentSetupAirlineInsurance holds details about the travel insurance purchased with the booking.
type PaymentSetupAirlineInsurance struct {
	// Type is the type of insurance.
	// [Optional]
	Type string `json:"type,omitempty"`

	// Company is the name of the insurance company.
	// [Optional]
	Company string `json:"company,omitempty"`

	// Price is the price of the insurance.
	// [Optional]
	Price *PaymentSetupAirlineInsurancePrice `json:"price,omitempty"`
}

// PaymentSetupAirlineInsurancePrice is the price of the travel insurance.
type PaymentSetupAirlineInsurancePrice struct {
	// Amount is the insurance amount, in the minor currency unit.
	// [Optional]
	Amount float64 `json:"amount,omitempty"`

	// Currency is the currency of the insurance amount, as a three-letter ISO currency code.
	// [Optional]
	Currency string `json:"currency,omitempty"`
}

// PaymentSetupAccommodation holds details about the accommodation or cruise booked by the customer.
type PaymentSetupAccommodation struct {
	// Name is, for lodging, the lodging name that appears on the storefront and customer receipts.
	// For cruise, the ship name booked for the cruise.
	// [Optional]
	Name string `json:"name,omitempty"`

	// BookingReference is a unique identifier for the booking.
	// [Optional]
	BookingReference string `json:"booking_reference,omitempty"`

	// CheckInDate is, for lodging bookings, the customer's check-in date. For cruise bookings,
	// the cruise departure date (also referred to as the sail date).
	// [Optional]
	// Format: yyyy-MM-dd
	CheckInDate *common.APIShortDate `json:"check_in_date,omitempty"`

	// CheckOutDate is, for lodging bookings, the customer's check-out date. For cruise bookings,
	// the cruise return date.
	// [Optional]
	// Format: yyyy-MM-dd
	CheckOutDate *common.APIShortDate `json:"check_out_date,omitempty"`

	// Address is the accommodation's address.
	// [Optional]
	Address *common.Address `json:"address,omitempty"`

	// NumberOfRooms is the total number of rooms booked for the accommodation.
	// [Optional]
	//
	// A pointer so an explicit zero is distinguishable from "not set". With a plain int and
	// omitempty, number_of_rooms: 0 was dropped from the payload entirely, and the swagger
	// declares the field integer with no minimum, so 0 is representable.
	NumberOfRooms *int `json:"number_of_rooms,omitempty"`

	// Guests is the list of guests staying at the accommodation.
	// [Optional]
	Guests []PaymentSetupAccommodationGuest `json:"guests,omitempty"`

	// Room is the list of rooms booked by the customer.
	// [Optional]
	Room []PaymentSetupAccommodationRoom `json:"room,omitempty"`

	// TotalNumberOfGuests is the total number of guests on the booking.
	// [Optional]
	TotalNumberOfGuests int `json:"total_number_of_guests,omitempty"`

	// Refundable specifies whether the booking is refundable.
	// [Optional]
	Refundable *bool `json:"refundable,omitempty"`

	// DeliveryRecipient is the recipient the booking confirmation is delivered to.
	// [Optional]
	DeliveryRecipient string `json:"delivery_recipient,omitempty"`

	// Host holds details about the host of the accommodation.
	// [Optional]
	Host *PaymentSetupAccommodationHost `json:"host,omitempty"`
}

// PaymentSetupAccommodationGuest is a guest staying at the accommodation.
type PaymentSetupAccommodationGuest struct {
	// FirstName is the guest's first name.
	// [Optional]
	FirstName string `json:"first_name,omitempty"`

	// LastName is the guest's last name.
	// [Optional]
	LastName string `json:"last_name,omitempty"`

	// DateOfBirth is the guest's date of birth.
	// [Optional]
	// Format: yyyy-MM-dd
	DateOfBirth *common.APIShortDate `json:"date_of_birth,omitempty"`
}

// PaymentSetupAccommodationRoom is a room booked by the customer.
type PaymentSetupAccommodationRoom struct {
	// Rate is the rate or cost of the room per day.
	// [Optional]
	Rate float64 `json:"rate,omitempty"`

	// NumberOfNights is the number of nights the room is booked for.
	// [Optional]
	NumberOfNights int `json:"number_of_nights,omitempty"`

	// Type is the room class or type booked. For example, deluxe.
	// [Optional]
	Type string `json:"type,omitempty"`
}

// PaymentSetupAccommodationHost holds details about the host of the accommodation.
type PaymentSetupAccommodationHost struct {
	// RegistrationDate is the date the host registered.
	// [Optional]
	// Format: yyyy-MM-dd
	RegistrationDate *common.APIShortDate `json:"registration_date,omitempty"`

	// TotalReservationCount is the total number of reservations the host has received.
	// [Optional]
	TotalReservationCount int `json:"total_reservation_count,omitempty"`
}

// AccountFundingTransactionPurpose is the purpose of the account funding transaction.
type AccountFundingTransactionPurpose string

const (
	AFTPurposeDonations         AccountFundingTransactionPurpose = "donations"
	AFTPurposeEducation         AccountFundingTransactionPurpose = "education"
	AFTPurposeEmergencyNeed     AccountFundingTransactionPurpose = "emergency_need"
	AFTPurposeExpatriation      AccountFundingTransactionPurpose = "expatriation"
	AFTPurposeFamilySupport     AccountFundingTransactionPurpose = "family_support"
	AFTPurposeFinancialServices AccountFundingTransactionPurpose = "financial_services"
	AFTPurposeGifts             AccountFundingTransactionPurpose = "gifts"
	AFTPurposeIncome            AccountFundingTransactionPurpose = "income"
	AFTPurposeInsurance         AccountFundingTransactionPurpose = "insurance"
	AFTPurposeInvestment        AccountFundingTransactionPurpose = "investment"
	AFTPurposeItServices        AccountFundingTransactionPurpose = "it_services"
	AFTPurposeLeisure           AccountFundingTransactionPurpose = "leisure"
	AFTPurposeLoanPayment       AccountFundingTransactionPurpose = "loan_payment"
	AFTPurposeMedicalTreatment  AccountFundingTransactionPurpose = "medical_treatment"
	AFTPurposeOther             AccountFundingTransactionPurpose = "other"
	AFTPurposePension           AccountFundingTransactionPurpose = "pension"
	AFTPurposeRoyalties         AccountFundingTransactionPurpose = "royalties"
	AFTPurposeSavings           AccountFundingTransactionPurpose = "savings"
	AFTPurposeTravelAndTourism  AccountFundingTransactionPurpose = "travel_and_tourism"
)

// AccountFundingTransactionIdentificationType is the type of identification used to identify the sender.
type AccountFundingTransactionIdentificationType string

const (
	AFTIdentificationPassport       AccountFundingTransactionIdentificationType = "passport"
	AFTIdentificationDrivingLicense AccountFundingTransactionIdentificationType = "driving_license"
	AFTIdentificationNationalId     AccountFundingTransactionIdentificationType = "national_id"
)

// AccountFundingTransactionIdentification holds the sender identification details.
type AccountFundingTransactionIdentification struct {
	// Type is the type of identification used to identify the sender.
	// [Optional]
	Type AccountFundingTransactionIdentificationType `json:"type,omitempty"`

	// Number is the identification number.
	// [Optional]
	Number string `json:"number,omitempty"`

	// IssuingCountry is the two-letter ISO country code of the country that issued the identification.
	// [Optional]
	IssuingCountry string `json:"issuing_country,omitempty"`
}

// AccountFundingTransactionSender holds the account funding transaction sender details.
type AccountFundingTransactionSender struct {
	// DateOfBirth is the date of birth of the sender.
	// [Optional]
	// Format: yyyy-MM-dd
	DateOfBirth *common.APIShortDate `json:"date_of_birth,omitempty"`

	// Reference is the unique reference for the sender of the payment.
	// [Optional]
	Reference string `json:"reference,omitempty"`

	// Identification holds the sender identification details.
	// [Optional]
	Identification *AccountFundingTransactionIdentification `json:"identification,omitempty"`
}

// AccountFundingTransactionRecipient holds the account funding transaction recipient details.
type AccountFundingTransactionRecipient struct {
	// DateOfBirth is the date of birth of the recipient.
	// [Optional]
	// Format: yyyy-MM-dd
	DateOfBirth *common.APIShortDate `json:"date_of_birth,omitempty"`

	// AccountNumber is any identifier like part of the PAN (first six digits and last four
	// digits), an IBAN, an internal account number, or a phone number related to the primary
	// recipient's account.
	// [Optional]
	// max 34 characters
	AccountNumber string `json:"account_number,omitempty"`

	// FirstName is the recipient's first name.
	// [Optional]
	// max 50 characters
	FirstName string `json:"first_name,omitempty"`

	// LastName is the recipient's last name.
	// [Optional]
	// max 50 characters
	LastName string `json:"last_name,omitempty"`

	// Address is a physical address.
	// [Optional]
	Address *common.Address `json:"address,omitempty"`
}

// PaymentSetupAccountFundingTransaction holds the account funding transaction details for the payment.
type PaymentSetupAccountFundingTransaction struct {
	// Enabled is whether to process this payment as an account funding transaction.
	// [Optional]
	Enabled *bool `json:"enabled,omitempty"`

	// Purpose specifies the purpose of the account funding transaction.
	// [Optional]
	Purpose AccountFundingTransactionPurpose `json:"purpose,omitempty"`

	// Sender holds the account funding transaction sender details.
	// [Optional]
	Sender *AccountFundingTransactionSender `json:"sender,omitempty"`

	// Recipient holds the account funding transaction recipient details.
	// [Optional]
	Recipient *AccountFundingTransactionRecipient `json:"recipient,omitempty"`
}

// BlikPaymentMethod is the Blik payment method's details and configuration.
type BlikPaymentMethod struct {
	PaymentMethodBase

	// Initialization is the initialization state of the payment method. It shadows
	// PaymentMethodBase.Initialization so the deprecation is reported; a value set on either one is
	// sent, and a received value is set on both.
	// [Optional]
	//
	// Deprecated: Not part of the Blik schema in the Payment Setups API specification; the API
	// ignores it on requests. There is no replacement. Will be removed in the next major version.
	Initialization PaymentMethodInitialization `json:"initialization,omitempty"`

	// PartnerCode is the 6-digit BLIK code generated by the customer's banking app.
	// [Optional]
	// Write only
	// ^[0-9]{6}$
	PartnerCode string `json:"partner_code,omitempty"`

	// PaymentMethodOptions holds the payment method options.
	// [Optional]
	//
	// Deprecated: Not part of the Blik schema in the Payment Setups API specification; the API
	// ignores it on requests. There is no replacement. Will be removed in the next major version.
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}
