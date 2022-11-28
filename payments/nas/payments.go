package nas

import (
	"encoding/json"
	"time"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/payments"
)

type AuthorizationType string

const (
	FinalAuthorizationType       AuthorizationType = "Final"
	EstimatedAuthorizationType   AuthorizationType = "Estimated"
	IncrementalAuthorizationType AuthorizationType = "Incremental"
)

type CaptureType string

const (
	NonFinalCaptureType CaptureType = "NonFinal"
	FinalCaptureType    CaptureType = "Final"
)

type InstructionScheme string

const (
	SwiftInstructionScheme   InstructionScheme = "swift"
	LocalInstructionScheme   InstructionScheme = "local"
	InstantInstructionScheme InstructionScheme = "instant"
)

type IdentificationType string

const (
	Passport       IdentificationType = "passport"
	DrivingLicence IdentificationType = "driving_licence"
	NationalId     IdentificationType = "national_id"
)

type (
	Product struct {
		Name           string `json:"name,omitempty"`
		Quantity       int    `json:"quantity,omitempty"`
		UnitPrice      int    `json:"unit_price,omitempty"`
		Reference      string `json:"reference,omitempty"`
		CommodityCode  string `json:"commodity_code,omitempty"`
		UnitOfMeasure  string `json:"unit_of_measure,omitempty"`
		TotalAmount    int64  `json:"total_amount,omitempty"`
		TaxAmount      int64  `json:"tax_amount,omitempty"`
		DiscountAmount int64  `json:"discount_amount,omitempty"`
		WxpayGoodsId   string `json:"wxpay_goods_id,omitempty"`
		ImageUrl       string `json:"image_url,omitempty"`
		Url            string `json:"url,omitempty"`
		Sku            string `json:"sku,omitempty"`
	}

	PayoutBillingDescriptor struct {
		Reference string `json:"reference,omitempty"`
	}

	PaymentInstruction struct {
		Purpose           string             `json:"purpose,omitempty"`
		ChargeBearer      string             `json:"charge_bearer,omitempty"`
		Repair            bool               `json:"repair"`
		Scheme            *InstructionScheme `json:"scheme,omitempty"`
		QuoteId           string             `json:"quote_id,omitempty"`
		SkipExpiry        string             `json:"skip_expiry,omitempty"`
		FundsTransferType string             `json:"funds_transfer_type,omitempty"`
		Mvv               string             `json:"mvv,omitempty"`
	}

	PaymentResponseBalances struct {
		TotalAuthorized    int `json:"total_authorized,omitempty"`
		TotalVoided        int `json:"total_voided,omitempty"`
		AvailableToVoid    int `json:"available_to_void"`
		TotalCaptured      int `json:"total_captured,omitempty"`
		AvailableToCapture int `json:"available_to_capture,omitempty"`
		TotalRefunded      int `json:"total_refunded,omitempty"`
		AvailableToRefund  int `json:"available_to_refund,omitempty"`
	}

	PaymentInstructionResponse struct {
		ValueDate *time.Time             `json:"value_date,omitempty"`
		Links     map[string]common.Link `json:"_links"`
	}

	Identification struct {
		Type           IdentificationType `json:"type,omitempty"`
		Number         string             `json:"number,omitempty"`
		IssuingCountry common.Country     `json:"issuing_country,omitempty"`
		DateOfExpiry   string             `json:"date_of_expiry,omitempty"`
	}
)

//Request
type (
	PaymentRequest struct {
		Source              interface{}                  `json:"source,omitempty"`
		Amount              int                          `json:"amount,omitempty"`
		Currency            common.Currency              `json:"currency,omitempty"`
		PaymentType         payments.PaymentType         `json:"payment_type,omitempty"`
		MerchantInitiated   bool                         `json:"merchant_initiated"`
		Reference           string                       `json:"reference,omitempty"`
		Description         string                       `json:"description,omitempty"`
		AuthorizationType   AuthorizationType            `json:"authorization_type,omitempty"`
		Capture             bool                         `json:"capture"`
		CaptureOn           time.Time                    `json:"capture_on,omitempty"`
		Customer            *common.CustomerRequest      `json:"customer,omitempty"`
		BillingDescriptor   *payments.BillingDescriptor  `json:"billing_descriptor,omitempty"`
		ShippingDetails     *payments.ShippingDetails    `json:"shipping,omitempty"`
		ThreeDsRequest      *payments.ThreeDsRequest     `json:"3ds,omitempty"`
		PreviousPaymentId   string                       `json:"previous_payment_id,omitempty"`
		ProcessingChannelId string                       `json:"processing_channel_id,omitempty"`
		Risk                *payments.RiskRequest        `json:"risk,omitempty"`
		SuccessUrl          string                       `json:"success_url,omitempty"`
		FailureUrl          string                       `json:"failure_url,omitempty"`
		PaymentIp           string                       `json:"payment_ip,omitempty"`
		Sender              interface{}                  `json:"sender,omitempty"`
		Recipient           *payments.PaymentRecipient   `json:"recipient,omitempty"`
		Marketplace         *common.MarketplaceData      `json:"marketplace,omitempty"`
		Processing          *payments.ProcessingSettings `json:"processing,omitempty"`
		Items               []Product                    `json:"items,omitempty"`
		Metadata            map[string]interface{}       `json:"metadata,omitempty"`
	}

	PayoutRequest struct {
		Source              interface{}              `json:"source,omitempty"`
		Destination         interface{}              `json:"destination,omitempty"`
		Amount              int                      `json:"amount,omitempty"`
		Currency            common.Currency          `json:"currency,omitempty"`
		Reference           string                   `json:"reference,omitempty"`
		BillingDescriptor   *PayoutBillingDescriptor `json:"billing_descriptor,omitempty"`
		Sender              interface{}              `json:"sender,omitempty"`
		Instruction         *PaymentInstruction      `json:"instruction,omitempty"`
		ProcessingChannelId string                   `json:"processing_channel_id,omitempty"`
		Metadata            map[string]interface{}   `json:"metadata,omitempty"`
	}

	CaptureRequest struct {
		Amount            int                          `json:"amount,omitempty"`
		CaptureType       CaptureType                  `json:"capture_type,omitempty"`
		Reference         string                       `json:"reference,omitempty"`
		Customer          *common.CustomerRequest      `json:"customer,omitempty"`
		Description       string                       `json:"description,omitempty"`
		BillingDescriptor *payments.BillingDescriptor  `json:"billing_descriptor,omitempty"`
		Shipping          *payments.ShippingDetails    `json:"shipping,omitempty"`
		Items             []Product                    `json:"items,omitempty"`
		Marketplace       *common.MarketplaceData      `json:"marketplace,omitempty"`
		AmountAllocations []common.AmountAllocations   `json:"amount_allocations,omitempty"`
		Processing        *payments.ProcessingSettings `json:"processing,omitempty"`
		Metadata          map[string]interface{}       `json:"metadata,omitempty"`
	}
)

//Response
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
		ExpiresOn       time.Time                   `json:"expires_on,omitempty"`
		Balances        *PaymentResponseBalances    `json:"balances,omitempty"`
		Processing      *payments.PaymentProcessing `json:"processing,omitempty"`
		Eci             string                      `json:"eci,omitempty"`
		SchemeId        string                      `json:"scheme_id,omitempty"`
		Links           map[string]common.Link      `json:"_links"`
	}

	PayoutResponse struct {
		HttpMetadata common.HttpMetadata
		Id           string                      `json:"id,omitempty"`
		Status       payments.PaymentStatus      `json:"status,omitempty"`
		Reference    string                      `json:"reference,omitempty"`
		Instruction  *PaymentInstructionResponse `json:"instruction,omitempty"`
	}

	GetPaymentResponse struct {
		HttpMetadata      common.HttpMetadata
		Id                string                          `json:"id,omitempty"`
		RequestedOn       time.Time                       `json:"requested_on,omitempty"`
		Source            *SourceResponse                 `json:"source,omitempty"`
		Destination       *DestinationResponse            `json:"destination,omitempty"`
		Sender            *SenderResponse                 `json:"sender,omitempty"`
		Amount            int                             `json:"amount,omitempty"`
		Currency          common.Currency                 `json:"currency,omitempty"`
		PaymentType       payments.PaymentType            `json:"payment_type,omitempty"`
		Reference         string                          `json:"reference,omitempty"`
		Description       string                          `json:"description,omitempty"`
		Approved          bool                            `json:"approved,omitempty"`
		ExpiresOn         time.Time                       `json:"expires_on,omitempty"`
		Status            payments.PaymentStatus          `json:"status,omitempty"`
		Balances          *PaymentResponseBalances        `json:"balances,omitempty"`
		ThreeDs           *payments.ThreeDsData           `json:"3ds,omitempty"`
		Risk              *payments.RiskAssessment        `json:"risk,omitempty"`
		Customer          *common.CustomerResponse        `json:"customer,omitempty"`
		BillingDescriptor *payments.BillingDescriptor     `json:"billing_descriptor,omitempty"`
		ShippingDetails   *payments.ShippingDetails       `json:"shipping,omitempty"`
		PaymentIp         string                          `json:"payment_ip,omitempty"`
		Marketplace       *common.MarketplaceData         `json:"marketplace,omitempty"`
		AmountAllocations []common.AmountAllocations      `json:"amount_allocations,omitempty"`
		Recipient         *payments.PaymentRecipient      `json:"recipient,omitempty"`
		ProcessingData    *payments.ProcessingData        `json:"processing,omitempty"`
		Items             []Product                       `json:"items,omitempty"`
		Metadata          map[string]interface{}          `json:"metadata,omitempty"`
		Eci               string                          `json:"eci,omitempty"`
		SchemeId          string                          `json:"scheme_id,omitempty"`
		Actions           []payments.PaymentActionSummary `json:"actions,omitempty"`
		ProcessedOn       time.Time                       `json:"processed_on,omitempty"`
		Links             map[string]common.Link          `json:"_links"`
	}

	GetPaymentActionsResponse struct {
		HttpMetadata common.HttpMetadata
		Actions      []PaymentAction `json:"actions,omitempty"`
	}

	PaymentAction struct {
		Id                string                       `json:"id,omitempty"`
		Type              payments.ActionType          `json:"type,omitempty"`
		ProcessedOn       time.Time                    `json:"processed_on,omitempty"`
		Amount            int                          `json:"amount,omitempty"`
		Approved          bool                         `json:"approved,omitempty"`
		AuthCode          string                       `json:"auth_code,omitempty"`
		ResponseCode      string                       `json:"response_code,omitempty"`
		ResponseSummary   string                       `json:"response_summary,omitempty"`
		AuthorizationType AuthorizationType            `json:"authorization_type,omitempty"`
		Reference         string                       `json:"reference,omitempty"`
		Processing        *payments.ProcessingSettings `json:"processing,omitempty"`
		Metadata          map[string]interface{}       `json:"metadata,omitempty"`
		AmountAllocations []common.AmountAllocations   `json:"amount_allocations,omitempty"`
		Links             map[string]common.Link       `json:"_links"`
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
