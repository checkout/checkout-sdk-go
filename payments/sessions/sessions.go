package payment_sessions

import (
	"time"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/payments"
	"github.com/checkout/checkout-sdk-go/payments/nas"
)

type PaymentMethodsType string

const (
	Applepay   PaymentMethodsType = "applepay"
	Bancontact PaymentMethodsType = "bancontact"
	Card       PaymentMethodsType = "card"
	EPS        PaymentMethodsType = "eps"
	Giropay    PaymentMethodsType = "giropay"
	Googlepay  PaymentMethodsType = "googlepay"
	Ideal      PaymentMethodsType = "ideal"
	KNet       PaymentMethodsType = "knet"
	Multibanco PaymentMethodsType = "multibanco"
	Przelewy24 PaymentMethodsType = "p24"
	Paypal     PaymentMethodsType = "paypal"
	Sofort     PaymentMethodsType = "sofort"
)

const PaymentSessionsPath = "payment-sessions"

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
		Amount                     int64                                    `json:"amount"`
		Currency                   common.Currency                          `json:"currency,omitempty"`
		PaymentType                payments.PaymentType                     `json:"payment_type,omitempty"`
		Billing                    *payments.BillingInformation             `json:"billing,omitempty"`
		BillingDescriptor          *payments.BillingDescriptor              `json:"billing_descriptor,omitempty"`
		Reference                  string                                   `json:"reference,omitempty"`
		Description                string                                   `json:"description,omitempty"`
		Customer                   *common.CustomerRequest                  `json:"customer,omitempty"`
		Shipping                   *payments.ShippingDetailsFlowHostedLinks `json:"shipping,omitempty"`
		Recipient                  *payments.PaymentRecipient               `json:"recipient,omitempty"`
		Processing                 *payments.ProcessingSettings             `json:"processing,omitempty"`
		ProcessingChannelId        string                                   `json:"processing_channel_id,omitempty"`
		ExpiresOn                  *time.Time                               `json:"expires_on,omitempty"`
		PaymentMethodConfiguration *payments.PaymentMethodConfiguration     `json:"payment_method_configuration,omitempty"`
		EnabledPaymentMethods      []PaymentMethodsType                     `json:"enabled_payment_methods,omitempty"`
		DisabledPaymentMethods     []PaymentMethodsType                     `json:"disabled_payment_methods,omitempty"`
		Items                      []payments.Product                       `json:"items,omitempty"`
		AmountAllocations          []common.AmountAllocations               `json:"amount_allocations,omitempty"`
		Risk                       *payments.RiskRequest                    `json:"risk,omitempty"`
		CustomerRetry              *payments.PaymentRetryRequest            `json:"customer_retry,omitempty"`
		DisplayName                string                                   `json:"display_name,omitempty"`
		SuccessUrl                 string                                   `json:"success_url,omitempty"`
		FailureUrl                 string                                   `json:"failure_url,omitempty"`
		Metadata                   map[string]interface{}                   `json:"metadata,omitempty"`
		Locale                     string                                   `json:"locale,omitempty"`
		ThreeDsRequest             *payments.ThreeDsRequestFlowHostedLinks  `json:"3ds,omitempty"`
		Sender                     *nas.Sender                              `json:"sender,omitempty"`
		Capture                    bool                                     `json:"capture"`
		CaptureOn                  *time.Time                               `json:"capture_on,omitempty"`
		IpAddress                  string                                   `json:"ip_address,omitempty"`
	}
)

type (
	PaymentMethods struct {
		Type        string   `json:"type,omitempty"`
		CardSchemes []string `json:"card_schemes,omitempty"`
	}

	PaymentSessionsResponse struct {
		HttpMetadata        common.HttpMetadata
		Id                  string                 `json:"id,omitempty"`
		PaymentSessionToken string                 `json:"payment_session_token,omitempty"`
		Links               map[string]common.Link `json:"links,omitempty"`
	}
)
