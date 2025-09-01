package links

import (
	"time"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/payments"
	"github.com/checkout/checkout-sdk-go/payments/nas"
)

const PaymentLinksPath = "payment-links"

type PaymentStatus string

const (
	Active          PaymentStatus = "Active"
	PaymentReceived PaymentStatus = "Payment Received"
	Expired         PaymentStatus = "Expired"
)

type (
	PaymentLinkRequest struct {
		Amount                     int                                      `json:"amount,omitempty"`
		Currency                   common.Currency                          `json:"currency,omitempty"`
		Billing                    *payments.BillingInformation             `json:"billing,omitempty"`
		PaymentType                payments.PaymentType                     `json:"payment_type,omitempty,omitempty"`
		PaymentIp                  string                                   `json:"payment_ip,omitempty"`
		BillingDescriptor          *payments.BillingDescriptor              `json:"billing_descriptor,omitempty"`
		Reference                  string                                   `json:"reference,omitempty"`
		Description                string                                   `json:"description,omitempty"`
		DisplayName                string                                   `json:"display_name,omitempty"`
		ProcessingChannelId        string                                   `json:"processing_channel_id,omitempty"`
		AmountAllocations          []common.AmountAllocations               `json:"amount_allocations,omitempty"`
		ExpiresIn                  int                                      `json:"expires_in,omitempty"`
		Customer                   *common.CustomerRequest                  `json:"customer,omitempty"`
		Shipping                   *payments.ShippingDetailsFlowHostedLinks `json:"shipping,omitempty"`
		Recipient                  *payments.PaymentRecipient               `json:"recipient,omitempty"`
		Processing                 *payments.ProcessingSettings             `json:"processing,omitempty"`
		AllowPaymentMethods        []payments.SourceType                    `json:"allow_payment_methods,omitempty"`
		DisabledPaymentMethods     []payments.SourceType                    `json:"disabled_payment_methods,omitempty"`
		Products                   []payments.Product                       `json:"products,omitempty"`
		Metadata                   map[string]interface{}                   `json:"metadata,omitempty"`
		ThreeDs                    *payments.ThreeDsRequestFlowHostedLinks  `json:"3ds,omitempty"`
		Risk                       *payments.RiskRequest                    `json:"risk,omitempty"`
		CustomerRetry              *payments.PaymentRetryRequest            `json:"customer_retry,omitempty"`
		Sender                     *nas.Sender                              `json:"sender,omitempty"`
		ReturnUrl                  string                                   `json:"return_url,omitempty"`
		Locale                     string                                   `json:"locale,omitempty"`
		Capture                    bool                                     `json:"capture,omitempty"`
		CaptureOn                  *time.Time                               `json:"capture_on,omitempty"`
		Instruction                *payments.PaymentInstruction             `json:"instruction,omitempty"`
		PaymentMethodConfiguration *payments.PaymentMethodConfiguration     `json:"payment_method_configuration,omitempty"`
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
