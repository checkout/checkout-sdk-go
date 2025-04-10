package hosted

import (
	"time"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/payments"

	"github.com/checkout/checkout-sdk-go/payments/nas"
)

const HostedPaymentsPath = "hosted-payments"

type PaymentStatus string

const (
	PaymentPending  PaymentStatus = "Payment Pending"
	PaymentReceived PaymentStatus = "Payment Received"
	Expired         PaymentStatus = "Expired"
)

type (
	HostedPaymentRequest struct {
		Currency                   common.Currency                          `json:"currency,omitempty"`
		Billing                    *payments.BillingInformation             `json:"billing,omitempty"`
		SuccessUrl                 string                                   `json:"success_url,omitempty"`
		CancelUrl                  string                                   `json:"cancel_url,omitempty"`
		FailureUrl                 string                                   `json:"failure_url,omitempty"`
		Amount                     int                                      `json:"amount,omitempty"`
		PaymentType                payments.PaymentType                     `json:"payment_type,omitempty,omitempty"`
		PaymentIp                  string                                   `json:"payment_ip,omitempty"`
		BillingDescriptor          *payments.BillingDescriptor              `json:"billing_descriptor,omitempty"`
		Reference                  string                                   `json:"reference,omitempty"`
		Description                string                                   `json:"description,omitempty"`
		DisplayName                string                                   `json:"display_name,omitempty"`
		ProcessingChannelId        string                                   `json:"processing_channel_id,omitempty"`
		AmountAllocations          []common.AmountAllocations               `json:"amount_allocations,omitempty"`
		Customer                   *common.CustomerRequest                  `json:"customer,omitempty"`
		Shipping                   *payments.ShippingDetailsFlowHostedLinks `json:"shipping,omitempty"`
		Recipient                  *payments.PaymentRecipient               `json:"recipient,omitempty"`
		Processing                 *payments.ProcessingSettings             `json:"processing,omitempty"`
		AllowPaymentMethods        []payments.SourceType                    `json:"allow_payment_methods,omitempty"`
		DisabledPaymentMethods     []payments.SourceType                    `json:"disabled_payment_methods,omitempty"`
		Products                   []payments.Product                       `json:"products,omitempty"`
		Risk                       *payments.RiskRequest                    `json:"risk,omitempty"`
		CustomerRetry              *payments.PaymentRetryRequest            `json:"customer_retry,omitempty"`
		Sender                     *nas.Sender                              `json:"sender,omitempty"`
		Metadata                   map[string]interface{}                   `json:"metadata,omitempty"`
		Locale                     payments.LocalType                       `json:"locale,omitempty"`
		ThreeDs                    *payments.ThreeDsRequestFlowHostedLinks  `json:"3ds,omitempty"`
		Capture                    bool                                     `json:"capture,omitempty"`
		CaptureOn                  *time.Time                               `json:"capture_on,omitempty"`
		Instruction                *payments.PaymentInstruction             `json:"instruction,omitempty"`
		PaymentMethodConfiguration *payments.PaymentMethodConfiguration     `json:"payment_method_configuration,omitempty"`
	}
)

type (
	HostedPaymentResponse struct {
		HttpMetadata common.HttpMetadata
		Id           string                 `json:"id,omitempty"`
		Reference    string                 `json:"reference,omitempty"`
		Warnings     []interface{}          `json:"warnings,omitempty"`
		Links        map[string]common.Link `json:"_links"`
	}

	HostedPaymentDetails struct {
		HttpMetadata      common.HttpMetadata
		Id                string                       `json:"id,omitempty"`
		Status            PaymentStatus                `json:"status,omitempty"`
		Amount            int                          `json:"amount,omitempty"`
		Currency          common.Currency              `json:"currency,omitempty"`
		Billing           *payments.BillingInformation `json:"billing,omitempty"`
		SuccessUrl        string                       `json:"success_url,omitempty"`
		CancelUrl         string                       `json:"cancel_url,omitempty"`
		FailureUrl        string                       `json:"failure_url,omitempty"`
		PaymentId         string                       `json:"payment_id,omitempty"`
		Reference         string                       `json:"reference,omitempty"`
		Description       string                       `json:"description,omitempty"`
		Customer          *common.CustomerResponse     `json:"customer,omitempty"`
		Products          []payments.Product           `json:"products,omitempty"`
		Metadata          map[string]interface{}       `json:"metadata,omitempty"`
		AmountAllocations []common.AmountAllocations   `json:"amount_allocations,omitempty"`
		Links             map[string]common.Link       `json:"_links"`
	}
)
