package hosted

import (
	"time"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/payments"
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
		Amount              int                          `json:"amount,omitempty"`
		Currency            common.Currency              `json:"currency,omitempty"`
		PaymentType         payments.PaymentType         `json:"payment_type,omitempty,omitempty"`
		PaymentIp           string                       `json:"payment_ip,omitempty"`
		BillingDescriptor   *payments.BillingDescriptor  `json:"billing_descriptor,omitempty"`
		Reference           string                       `json:"reference,omitempty"`
		Description         string                       `json:"description,omitempty"`
		Customer            *common.CustomerRequest      `json:"customer,omitempty"`
		Shipping            *payments.ShippingDetails    `json:"shipping,omitempty"`
		Billing             *payments.BillingInformation `json:"billing,omitempty"`
		Recipient           *payments.PaymentRecipient   `json:"recipient,omitempty"`
		Processing          *payments.ProcessingSettings `json:"processing,omitempty"`
		AllowPaymentMethods []payments.SourceType        `json:"allow_payment_methods,omitempty"`
		Products            []payments.Product           `json:"products,omitempty"`
		Risk                *payments.RiskRequest        `json:"risk,omitempty"`
		SuccessUrl          string                       `json:"success_url,omitempty"`
		CancelUrl           string                       `json:"cancel_url,omitempty"`
		FailureUrl          string                       `json:"failure_url,omitempty"`
		Metadata            map[string]interface{}       `json:"metadata,omitempty"`
		Locale              string                       `json:"locale,omitempty"`
		ThreeDs             *payments.ThreeDsRequest     `json:"3ds,omitempty"`
		Capture             bool                         `json:"capture,omitempty"`
		CaptureOn           *time.Time                   `json:"capture_on,omitempty"`
		//Not available on previous
		ProcessingChannelId string                     `json:"processing_channel_id,omitempty"`
		AmountAllocations   []common.AmountAllocations `json:"amount_allocations,omitempty"`
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
		HttpMetadata common.HttpMetadata
		Id           string                       `json:"id,omitempty"`
		Status       PaymentStatus                `json:"status,omitempty"`
		PaymentId    string                       `json:"payment_id,omitempty"`
		Amount       int                          `json:"amount,omitempty"`
		Currency     common.Currency              `json:"currency,omitempty"`
		Reference    string                       `json:"reference,omitempty"`
		Description  string                       `json:"description,omitempty"`
		Customer     *common.CustomerRequest      `json:"customer,omitempty"`
		Billing      *payments.BillingInformation `json:"billing,omitempty"`
		Products     []payments.Product           `json:"products,omitempty"`
		Metadata     map[string]interface{}       `json:"metadata,omitempty"`
		SuccessUrl   string                       `json:"success_url,omitempty"`
		CancelUrl    string                       `json:"cancel_url,omitempty"`
		FailureUrl   string                       `json:"failure_url,omitempty"`
		Links        map[string]common.Link       `json:"_links"`
		//Not available on previous
		AmountAllocations []common.AmountAllocations `json:"amount_allocations,omitempty"`
	}
)
