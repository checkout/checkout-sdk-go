package abc

import (
	"encoding/json"
	"time"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/payments"
)

type FundTransferType string

const (
	AA FundTransferType = "AA"
	PP FundTransferType = "PP"
	FT FundTransferType = "FT"
	FD FundTransferType = "FD"
	PD FundTransferType = "PD"
	LO FundTransferType = "LO"
	OG FundTransferType = "OG"
)

// Request
type (
	PaymentRequest struct {
		Source            interface{}                 `json:"source,omitempty"`
		Amount            int                         `json:"amount,omitempty"`
		Currency          common.Currency             `json:"currency,omitempty"`
		PaymentType       payments.PaymentType        `json:"payment_type,omitempty"`
		MerchantInitiated bool                        `json:"merchant_initiated"`
		Reference         string                      `json:"reference,omitempty"`
		Description       string                      `json:"description,omitempty"`
		Capture           bool                        `json:"capture"`
		CaptureOn         time.Time                   `json:"capture_on,omitempty"`
		Customer          *common.CustomerRequest     `json:"customer,omitempty"`
		BillingDescriptor *payments.BillingDescriptor `json:"billing_descriptor,omitempty"`
		ShippingDetails   *payments.ShippingDetails   `json:"shipping,omitempty"`
		PreviousPaymentId string                      `json:"previous_payment_id,omitempty"`
		Risk              *payments.RiskRequest       `json:"risk,omitempty"`
		SuccessUrl        string                      `json:"success_url,omitempty"`
		FailureUrl        string                      `json:"failure_url,omitempty"`
		PaymentIp         string                      `json:"payment_ip,omitempty"`
		ThreeDsRequest    *payments.ThreeDsRequest    `json:"3ds,omitempty"`
		PaymentRecipient  *payments.PaymentRecipient  `json:"recipient,omitempty"`
		Metadata          map[string]interface{}      `json:"metadata,omitempty"`
		Processing        map[string]interface{}      `json:"processing,omitempty"`
	}

	PayoutRequest struct {
		Destination       interface{}                 `json:"destination,omitempty"`
		Amount            int                         `json:"amount,omitempty"`
		FundTransferType  FundTransferType            `json:"fund_transfer_type,omitempty"`
		Currency          common.Currency             `json:"currency,omitempty"`
		PaymentType       payments.PaymentType        `json:"payment_type,omitempty"`
		Reference         string                      `json:"reference,omitempty"`
		Description       string                      `json:"description,omitempty"`
		Capture           bool                        `json:"capture"`
		CaptureOn         time.Time                   `json:"capture_on,omitempty"`
		Customer          *common.CustomerRequest     `json:"customer,omitempty"`
		BillingDescriptor *payments.BillingDescriptor `json:"billing_descriptor,omitempty"`
		ShippingDetails   *payments.ShippingDetails   `json:"shipping,omitempty"`
		PreviousPaymentId string                      `json:"previous_payment_id,omitempty"`
		Risk              *payments.RiskRequest       `json:"risk,omitempty"`
		SuccessUrl        string                      `json:"success_url,omitempty"`
		FailureUrl        string                      `json:"failure_url,omitempty"`
		PaymentIp         string                      `json:"payment_ip,omitempty"`
		PaymentRecipient  *payments.PaymentRecipient  `json:"recipient,omitempty"`
		Metadata          map[string]interface{}      `json:"metadata,omitempty"`
		Processing        map[string]interface{}      `json:"processing,omitempty"`
	}

	CaptureRequest struct {
		Amount    int                    `json:"amount,omitempty"`
		Reference string                 `json:"reference,omitempty"`
		Metadata  map[string]interface{} `json:"metadata,omitempty"`
	}
)

// Response
type (
	PaymentResponse struct {
		HttpMetadata    common.HttpMetadata
		ActionId        string                      `json:"action_id,omitempty"`
		Amount          int                         `json:"amount,omitempty"`
		Approved        bool                        `json:"approved,omitempty"`
		AuthCode        string                      `json:"auth_code,omitempty"`
		Id              string                      `json:"id,omitempty"`
		Currency        common.Currency             `json:"currency,omitempty"`
		Customer        *common.CustomerResponse    `json:"customer,omitempty"`
		Source          *SourceResponse             `json:"source,omitempty"`
		Status          payments.PaymentStatus      `json:"status,omitempty"`
		ThreeDs         *payments.ThreeDsEnrollment `json:"3ds,omitempty"`
		Reference       string                      `json:"reference,omitempty"`
		ResponseCode    string                      `json:"response_code,omitempty"`
		ResponseSummary string                      `json:"response_summary,omitempty"`
		Risk            *payments.RiskAssessment    `json:"risk,omitempty"`
		ProcessedOn     time.Time                   `json:"processed_on,omitempty"`
		Processing      *payments.PaymentProcessing `json:"processing,omitempty"`
		Eci             string                      `json:"eci,omitempty"`
		SchemeId        string                      `json:"scheme_id,omitempty"`
		Links           map[string]common.Link      `json:"_links"`
	}

	GetPaymentResponse struct {
		HttpMetadata      common.HttpMetadata
		Id                string                          `json:"id,omitempty"`
		RequestedOn       time.Time                       `json:"requested_on,omitempty"`
		Source            interface{}                     `json:"source,omitempty"`
		Destination       interface{}                     `json:"destination,omitempty"`
		Amount            int                             `json:"amount,omitempty"`
		Currency          common.Currency                 `json:"currency,omitempty"`
		PaymentType       payments.PaymentType            `json:"payment_type,omitempty"`
		Reference         string                          `json:"reference,omitempty"`
		Description       string                          `json:"description,omitempty"`
		Approved          bool                            `json:"approved,omitempty"`
		Status            payments.PaymentStatus          `json:"status,omitempty"`
		ThreeDs           *payments.ThreeDsData           `json:"3ds,omitempty"`
		Risk              *payments.RiskAssessment        `json:"risk,omitempty"`
		Customer          *common.CustomerResponse        `json:"customer,omitempty"`
		BillingDescriptor *payments.BillingDescriptor     `json:"billing_descriptor,omitempty"`
		ShippingDetails   *payments.ShippingDetails       `json:"shipping,omitempty"`
		PaymentIp         string                          `json:"payment_ip,omitempty"`
		PaymentRecipient  *payments.PaymentRecipient      `json:"recipient,omitempty"`
		Metadata          map[string]interface{}          `json:"metadata,omitempty"`
		Eci               string                          `json:"eci,omitempty"`
		SchemeId          string                          `json:"scheme_id,omitempty"`
		Actions           []payments.PaymentActionSummary `json:"actions,omitempty"`
		Links             map[string]common.Link          `json:"_links"`
	}

	GetPaymentActionsResponse struct {
		HttpMetadata common.HttpMetadata
		Actions      []PaymentAction `json:"actions,omitempty"`
	}

	PaymentAction struct {
		Id              string                 `json:"id,omitempty"`
		Type            payments.ActionType    `json:"type,omitempty"`
		ProcessedOn     time.Time              `json:"processed_on,omitempty"`
		Amount          int                    `json:"amount,omitempty"`
		Approved        bool                   `json:"approved,omitempty"`
		AuthCode        string                 `json:"auth_code,omitempty"`
		ResponseCode    string                 `json:"response_code,omitempty"`
		ResponseSummary string                 `json:"response_summary,omitempty"`
		Reference       string                 `json:"reference,omitempty"`
		Processing      *payments.Processing   `json:"processing,omitempty"`
		Metadata        map[string]interface{} `json:"metadata,omitempty"`
		Links           map[string]common.Link `json:"_links"`
	}

	GetPaymentListResponse struct {
		HttpMetadata common.HttpMetadata
		Limit        int                  `json:"limit,omitempty"`
		Skip         int                  `json:"skip,omitempty"`
		TotalCount   int                  `json:"total_count,omitempty"`
		Data         []GetPaymentResponse `json:"data,omitempty"`
	}
)

func (p *GetPaymentActionsResponse) UnmarshalJSON(data []byte) error {
	var actions []PaymentAction
	if err := json.Unmarshal(data, &actions); err != nil {
		return err
	}
	p.Actions = actions
	return nil
}
