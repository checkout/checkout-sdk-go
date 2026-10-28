package links

import (
	"time"

	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/payments"
	"github.com/checkout/checkout-sdk-go/v3/payments/nas"
)

const PaymentLinksPath = "payment-links"

type PaymentStatus string

const (
	Active          PaymentStatus = "Active"
	PaymentReceived PaymentStatus = "Payment Received"
	Expired         PaymentStatus = "Expired"
)

type (
	PaymentInterfacesCustomerSummary struct {
		RegistrationDate     string  `json:"registration_date,omitempty"`
		FirstTransactionDate string  `json:"first_transaction_date,omitempty"`
		LastPaymentDate      string  `json:"last_payment_date,omitempty"`
		TotalOrderCount      int     `json:"total_order_count,omitempty"`
		LastPaymentAmount    float64 `json:"last_payment_amount,omitempty"`
		IsPremiumCustomer    bool    `json:"is_premium_customer,omitempty"`
		IsReturningCustomer  bool    `json:"is_returning_customer,omitempty"`
		LifetimeValue        float64 `json:"lifetime_value,omitempty"`
	}

	PaymentLinkCustomer struct {
		Email   string                            `json:"email,omitempty"`
		Name    string                            `json:"name,omitempty"`
		Phone   *common.Phone                     `json:"phone,omitempty"`
		Summary *PaymentInterfacesCustomerSummary `json:"summary,omitempty"`
	}
)

type (
	PaymentLinkRequest struct {
		// Amount is the payment amount. The exact format depends on the currency.
		// [Required]
		// min 1
		Amount int `json:"amount,omitempty"`
		// Currency is the three-letter ISO currency code of the payment. The currency and billing.address.country fields
		// determine which payment methods are shown on the payment page.
		// [Required]
		// max 3 characters
		// min 3 characters
		Currency common.Currency `json:"currency,omitempty"`
		// Billing is the billing details.
		// [Required]
		Billing *payments.BillingInformation `json:"billing,omitempty"`
		// PaymentType: This must be specified for card payments where the cardholder is not present (i.e., recurring or
		// mail order / telephone order).
		// [Optional]
		// Default: Regular
		PaymentType payments.PaymentType `json:"payment_type,omitempty"`
		// PaymentIp is the Customers IP address. Only IPv4 and IPv6 addresses are accepted.
		// [Optional]
		// Format: ipv4
		// max 45 characters
		//
		// Deprecated: Use risk.device.network.ipv4 or risk.device.network.ipv6 instead.
		// This field will be removed in the next major version.
		PaymentIp string `json:"payment_ip,omitempty"`
		// BillingDescriptor is an optional description that is displayed on the customer's statement identifying a
		// purchase.
		// [Optional]
		BillingDescriptor *payments.BillingDescriptor `json:"billing_descriptor,omitempty"`
		// Reference is a reference you can use to identify the payment. For example, an order number. - For Amex
		// payments, this must be at most 30 characters. - For Benefit payments, the reference must be a unique
		// alphanumeric value. - For iDEAL payments, the reference is required and must be an alphanumeric value with a
		// 35-character limit.
		// [Optional]
		// max 50 characters
		Reference string `json:"reference,omitempty"`
		// Description is a description of the payment.
		// [Optional]
		// max 100 characters
		Description string `json:"description,omitempty"`
		// DisplayName is the merchant name to display to customers on the checkout page.
		// [Optional]
		DisplayName string `json:"display_name,omitempty"`
		// ProcessingChannelId is the processing channel to use for the payment. This field is required if you have more
		// than one processing channel associated with your API key.
		// [Optional]
		// ^(pc)_(\w{26})$
		ProcessingChannelId string `json:"processing_channel_id,omitempty"`
		// AmountAllocations is the sub-entities that the payment is being processed on behalf of.
		// [Optional]
		AmountAllocations []common.AmountAllocations `json:"amount_allocations,omitempty"`
		// ExpiresIn is the time for which the link remains valid, in seconds.
		// [Optional]
		// min 1
		// max 5184000
		// Default: 86400
		ExpiresIn int `json:"expires_in,omitempty"`
		// Customer is the customer's details.
		// [Optional]
		Customer *PaymentLinkCustomer `json:"customer,omitempty"`
		// Shipping is the address any products are being sent to.
		// [Optional]
		Shipping *payments.ShippingDetailsFlowHostedLinks `json:"shipping,omitempty"`
		// Recipient: Information about the recipient of the payment's funds. Applies to Account Funding Transactions,
		// and VISA or Mastercard domestic UK transactions processed by financial institutions.
		// [Optional]
		Recipient *payments.PaymentRecipient `json:"recipient,omitempty"`
		// Processing: Use the processing object to influence or override the data sent during card processing.
		// [Optional]
		Processing *payments.ProcessingSettings `json:"processing,omitempty"`
		// AllowPaymentMethods specifies which payment method options to present to the customer. The values in this
		// field override any equivalent values in disabled_payment_methods. Some payment methods have additional field
		// requirements. See the documentation for more information.
		// [Optional]
		AllowPaymentMethods []payments.SourceType `json:"allow_payment_methods,omitempty"`
		// DisabledPaymentMethods specifies which payment method options to not present to the customer. If you specify
		// the same payment method in this field and in allow_payment_methods, the disabled_payment_methods value will be
		// overridden. Any payment method options not explicitly specified in this field will be presented to the
		// customer by default.
		// [Optional]
		DisabledPaymentMethods []payments.SourceType `json:"disabled_payment_methods,omitempty"`
		// Products contains details about the products in the order.
		// [Optional]
		Products []payments.Product `json:"products,omitempty"`
		// Metadata allows you to store additional information about the transaction. This object only allows one level
		// of depth, so cannot accept non-primitive data types such as objects or arrays.
		// [Optional]
		Metadata map[string]interface{} `json:"metadata,omitempty"`
		// ThreeDs: Information required for 3D Secure payments.
		// [Optional]
		ThreeDs *payments.ThreeDsRequestFlowHostedLinks `json:"3ds,omitempty"`
		// Risk configures the risk assessment performed during payment processing.
		// [Optional]
		Risk *payments.RiskRequest `json:"risk,omitempty"`
		// CustomerRetry configures the parameters for customer retries. A customer retry is a payment attempt performed
		// by the customer on the checkout page. This differs from authorization request retries
		// (https://api-reference.checkout.com/#operation/requestAPaymentOrPayout!path=0/retry&t=request) performed in
		// the back end by Checkout.com.
		// [Optional]
		CustomerRetry *payments.CustomerRetry `json:"customer_retry,omitempty"`
		// Sender is the sender of the payment.
		// [Optional]
		Sender *nas.Sender `json:"sender,omitempty"`
		// ReturnUrl: If provided, the success page will include a button that redirects your customer to the provided
		// URL.
		// [Optional]
		// Format: uri
		// max 255 characters
		ReturnUrl string `json:"return_url,omitempty"`
		// Locale creates a translated version of the page in the specified language.
		// [Optional]
		// Default: en-GB
		// Enum: "ar" "da-DK" "de-DE" "el" "en-GB" "es-ES" "fi-FI" "fil-PH" "fr-FR" "hi-IN" "id-ID" "it-IT" "ja-JP" "ms-MY"
		// "nb-NO" "nl-NL" "pt-PT" "sv-SE" "th-TH" "vi-VN" "zh-CN" "zh-HK" "zh-TW"
		Locale string `json:"locale,omitempty"`
		// Capture: Whether to capture the payment (if applicable).
		// [Optional]
		Capture bool `json:"capture,omitempty"`
		// CaptureOn is a timestamp (ISO 8601 code) that determines when the payment should be captured. Providing this
		// field will automatically set capture to true.
		// [Optional]
		// Format: date-time
		CaptureOn *time.Time `json:"capture_on,omitempty"`
		// Instruction: Details about the payment instruction.
		// [Optional]
		Instruction *payments.PaymentInstruction `json:"instruction,omitempty"`
		// PaymentMethodConfiguration: Configuration options for specific payment methods.
		// [Optional]
		PaymentMethodConfiguration *payments.PaymentMethodConfiguration `json:"payment_method_configuration,omitempty"`
		// AuthorizationType is the authorization type.
		// [Optional]
		// Default: Final
		AuthorizationType nas.AuthorizationType `json:"authorization_type,omitempty" default:"Final"`
		// PaymentPlan is the information to process a recurring payment request. To be used when the payment_type is
		// Recurring.
		// [Optional]
		PaymentPlan *payments.PaymentPlan `json:"payment_plan,omitempty"`
	}
)

type (
	PaymentLinkResponse struct {
		HttpMetadata common.HttpMetadata
		Id           string                 `json:"id,omitempty"`
		Links        map[string]common.Link `json:"_links"`
		ExpiresOn    string                 `json:"expires_on,omitempty"`
		Reference    string                 `json:"reference,omitempty"`
		Warnings     []interface{}          `json:"warnings,omitempty"`
	}

	PaymentLinkDetails struct {
		HttpMetadata        common.HttpMetadata
		Id                  string                       `json:"id,omitempty"`
		Status              PaymentStatus                `json:"status,omitempty"`
		Amount              int                          `json:"amount,omitempty"`
		Currency            common.Currency              `json:"currency,omitempty"`
		ExpiresOn           string                       `json:"expires_on,omitempty"`
		CreatedOn           string                       `json:"created_on,omitempty"`
		Billing             *payments.BillingInformation `json:"billing,omitempty"`
		Links               map[string]common.Link       `json:"_links"`
		PaymentId           string                       `json:"payment_id,omitempty"`
		Reference           string                       `json:"reference,omitempty"`
		Description         string                       `json:"description,omitempty"`
		ProcessingChannelId string                       `json:"processing_channel_id,omitempty"`
		AmountAllocations   []common.AmountAllocations   `json:"amount_allocations,omitempty"`
		Customer            *common.CustomerRequest      `json:"customer,omitempty"`
		Shipping            *payments.ShippingDetails    `json:"shipping,omitempty"`
		Products            []payments.Product           `json:"products,omitempty"`
		Metadata            map[string]interface{}       `json:"metadata,omitempty"`
		Locale              string                       `json:"locale,omitempty"`
		ReturnUrl           string                       `json:"return_url,omitempty"`
	}
)
