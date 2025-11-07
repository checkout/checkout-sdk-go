package payment_sessions

import (
	"time"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/payments"
	"github.com/checkout/checkout-sdk-go/payments/nas"
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
		Capture                    bool                                     `json:"capture"`
		CaptureOn                  *time.Time                               `json:"capture_on,omitempty"`
		ExpiresOn                  *time.Time                               `json:"expires_on,omitempty"`
		EnabledPaymentMethods      []PaymentMethodsType                     `json:"enabled_payment_methods,omitempty"`
		DisabledPaymentMethods     []PaymentMethodsType                     `json:"disabled_payment_methods,omitempty"`
		CustomerRetry              *payments.PaymentRetryRequest            `json:"customer_retry,omitempty"`
		IpAddress                  string                                   `json:"ip_address,omitempty"`
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
		Capture                    bool                                     `json:"capture"`
		CaptureOn                  *time.Time                               `json:"capture_on,omitempty"`
	}

	SubmitPaymentSessionRequest struct {
		SessionData    string                                  `json:"session_data"`
		Amount         int64                                   `json:"amount,omitempty"`
		Reference      string                                  `json:"reference,omitempty"`
		Items          []payments.Product                      `json:"items,omitempty"`
		ThreeDsRequest *payments.ThreeDsRequestFlowHostedLinks `json:"3ds,omitempty"`
		IpAddress      string                                  `json:"ip_address,omitempty"`
		PaymentType    payments.PaymentType                    `json:"payment_type,omitempty"`
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
