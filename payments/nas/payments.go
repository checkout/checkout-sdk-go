package nas

import (
	"encoding/json"
	"time"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/payments"
	"github.com/checkout/checkout-sdk-go/payments/nas/sources"
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
	PayoutBillingDescriptor struct {
		Reference string `json:"reference,omitempty"`
	}

	PaymentInstruction struct {
		Purpose           string                    `json:"purpose,omitempty"`
		ChargeBearer      string                    `json:"charge_bearer,omitempty"`
		Repair            bool                      `json:"repair"`
		Scheme            *InstructionScheme        `json:"scheme,omitempty"`
		QuoteId           string                    `json:"quote_id,omitempty"`
		SkipExpiry        string                    `json:"skip_expiry,omitempty"`
		FundsTransferType payments.FundTransferType `json:"funds_transfer_type,omitempty"`
		Mvv               string                    `json:"mvv,omitempty"`
	}

	PaymentResponseBalances struct {
		TotalAuthorized    int64 `json:"total_authorized,omitempty"`
		TotalVoided        int64 `json:"total_voided,omitempty"`
		AvailableToVoid    int64 `json:"available_to_void"`
		TotalCaptured      int64 `json:"total_captured,omitempty"`
		AvailableToCapture int64 `json:"available_to_capture,omitempty"`
		TotalRefunded      int64 `json:"total_refunded,omitempty"`
		AvailableToRefund  int64 `json:"available_to_refund,omitempty"`
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

	PartialAuthorization struct {
		Enabled bool `json:"enabled,omitempty"`
	}
)

// Request
type (
	PaymentRequest struct {
		PaymentContextId     string                        `json:"payment_context_id,omitempty"`
		Source               payments.PaymentSource        `json:"source,omitempty"`
		Amount               int64                         `json:"amount,omitempty"`
		Currency             common.Currency               `json:"currency,omitempty"`
		PaymentType          payments.PaymentType          `json:"payment_type,omitempty"`
		MerchantInitiated    bool                          `json:"merchant_initiated"`
		Reference            string                        `json:"reference,omitempty"`
		Description          string                        `json:"description,omitempty"`
		AuthorizationType    AuthorizationType             `json:"authorization_type,omitempty"`
		PartialAuthorization *PartialAuthorization         `json:"partial_authorization,omitempty"`
		Capture              bool                          `json:"capture"`
		CaptureOn            *time.Time                    `json:"capture_on,omitempty"`
		Customer             *common.CustomerRequest       `json:"customer,omitempty"`
		BillingDescriptor    *payments.BillingDescriptor   `json:"billing_descriptor,omitempty"`
		ShippingDetails      *payments.ShippingDetails     `json:"shipping,omitempty"`
		Segment              *payments.PaymentSegment      `json:"segment,omitempty"`
		ThreeDsRequest       *payments.ThreeDsRequest      `json:"3ds,omitempty"`
		PreviousPaymentId    string                        `json:"previous_payment_id,omitempty"`
		ProcessingChannelId  string                        `json:"processing_channel_id,omitempty"`
		Risk                 *payments.RiskRequest         `json:"risk,omitempty"`
		SuccessUrl           string                        `json:"success_url,omitempty"`
		FailureUrl           string                        `json:"failure_url,omitempty"`
		PaymentIp            string                        `json:"payment_ip,omitempty"`
		Sender               Sender                        `json:"sender,omitempty"`
		Recipient            *payments.PaymentRecipient    `json:"recipient,omitempty"`
		Marketplace          *common.MarketplaceData       `json:"marketplace,omitempty"`
		AmountAllocations    []common.AmountAllocations    `json:"amount_allocations,omitempty"`
		Processing           *payments.ProcessingSettings  `json:"processing,omitempty"`
		Items                []payments.Product            `json:"items,omitempty"`
		Retry                *payments.PaymentRetryRequest `json:"retry,omitempty"`
		Metadata             map[string]interface{}        `json:"metadata,omitempty"`
		Instruction          *PaymentInstruction           `json:"instruction,omitempty"`
	}

	PayoutRequest struct {
		Source              sources.PayoutSource     `json:"source,omitempty"`
		Destination         payments.Destination     `json:"destination,omitempty"`
		Amount              int64                    `json:"amount,omitempty"`
		Currency            common.Currency          `json:"currency,omitempty"`
		Reference           string                   `json:"reference,omitempty"`
		BillingDescriptor   *PayoutBillingDescriptor `json:"billing_descriptor,omitempty"`
		Sender              Sender                   `json:"sender,omitempty"`
		Instruction         *PaymentInstruction      `json:"instruction,omitempty"`
		ProcessingChannelId string                   `json:"processing_channel_id,omitempty"`
		Metadata            map[string]interface{}   `json:"metadata,omitempty"`
	}

	IncrementAuthorizationRequest struct {
		Amount    int64                  `json:"amount,omitempty"`
		Reference string                 `json:"reference,omitempty"`
		Metadata  map[string]interface{} `json:"metadata,omitempty"`
	}

	CaptureRequest struct {
		Amount            int64                        `json:"amount,omitempty"`
		CaptureType       CaptureType                  `json:"capture_type,omitempty"`
		Reference         string                       `json:"reference,omitempty"`
		Customer          *common.CustomerRequest      `json:"customer,omitempty"`
		Description       string                       `json:"description,omitempty"`
		BillingDescriptor *payments.BillingDescriptor  `json:"billing_descriptor,omitempty"`
		Shipping          *payments.ShippingDetails    `json:"shipping,omitempty"`
		Items             []payments.Product           `json:"items,omitempty"`
		Marketplace       *common.MarketplaceData      `json:"marketplace,omitempty"`
		AmountAllocations []common.AmountAllocations   `json:"amount_allocations,omitempty"`
		Processing        *payments.ProcessingSettings `json:"processing,omitempty"`
		Metadata          map[string]interface{}       `json:"metadata,omitempty"`
	}
)

// Response
type (
	PaymentResponse struct {
		HttpMetadata    common.HttpMetadata
		ActionId        string                         `json:"action_id,omitempty"`
		Amount          int64                          `json:"amount,omitempty"`
		Approved        bool                           `json:"approved,omitempty"`
		AuthCode        string                         `json:"auth_code,omitempty"`
		Id              string                         `json:"id,omitempty"`
		Currency        common.Currency                `json:"currency,omitempty"`
		Customer        *common.CustomerResponse       `json:"customer,omitempty"`
		Source          *SourceResponse                `json:"source,omitempty"`
		Status          payments.PaymentStatus         `json:"status,omitempty"`
		ThreeDs         *payments.ThreeDsEnrollment    `json:"3ds,omitempty"`
		Reference       string                         `json:"reference,omitempty"`
		ResponseCode    string                         `json:"response_code,omitempty"`
		ResponseSummary string                         `json:"response_summary,omitempty"`
		Risk            *payments.RiskAssessment       `json:"risk,omitempty"`
		ProcessedOn     *time.Time                     `json:"processed_on,omitempty"`
		ExpiresOn       *time.Time                     `json:"expires_on,omitempty"`
		Balances        *PaymentResponseBalances       `json:"balances,omitempty"`
		Processing      *payments.PaymentProcessing    `json:"processing,omitempty"`
		Eci             string                         `json:"eci,omitempty"`
		SchemeId        string                         `json:"scheme_id,omitempty"`
		Retry           *payments.PaymentRetryResponse `json:"retry,omitempty"`
		Links           map[string]common.Link         `json:"_links"`
	}

	PayoutResponse struct {
		HttpMetadata common.HttpMetadata
		Id           string                      `json:"id,omitempty"`
		Status       payments.PaymentStatus      `json:"status,omitempty"`
		Reference    string                      `json:"reference,omitempty"`
		Instruction  *PaymentInstructionResponse `json:"instruction,omitempty"`
	}

	GetPaymentResponse struct {
		HttpMetadata             common.HttpMetadata
		Id                       string                          `json:"id,omitempty"`
		RequestedOn              *time.Time                      `json:"requested_on,omitempty"`
		Source                   *SourceResponse                 `json:"source,omitempty"`
		Destination              *DestinationResponse            `json:"destination,omitempty"`
		Sender                   *SenderResponse                 `json:"sender,omitempty"`
		Amount                   int64                           `json:"amount,omitempty"`
		Currency                 common.Currency                 `json:"currency,omitempty"`
		PaymentType              payments.PaymentType            `json:"payment_type,omitempty"`
		Reference                string                          `json:"reference,omitempty"`
		Description              string                          `json:"description,omitempty"`
		Approved                 bool                            `json:"approved,omitempty"`
		ExpiresOn                *time.Time                      `json:"expires_on,omitempty"`
		Status                   payments.PaymentStatus          `json:"status,omitempty"`
		Balances                 *PaymentResponseBalances        `json:"balances,omitempty"`
		ThreeDs                  *payments.ThreeDsData           `json:"3ds,omitempty"`
		Risk                     *payments.RiskAssessment        `json:"risk,omitempty"`
		Customer                 *common.CustomerResponse        `json:"customer,omitempty"`
		BillingDescriptor        *payments.BillingDescriptor     `json:"billing_descriptor,omitempty"`
		ShippingDetails          *payments.ShippingDetails       `json:"shipping,omitempty"`
		PaymentIp                string                          `json:"payment_ip,omitempty"`
		Marketplace              *common.MarketplaceData         `json:"marketplace,omitempty"`
		AmountAllocations        []common.AmountAllocations      `json:"amount_allocations,omitempty"`
		Recipient                *payments.PaymentRecipient      `json:"recipient,omitempty"`
		ProcessingData           *payments.ProcessingData        `json:"processing,omitempty"`
		Items                    []payments.Product              `json:"items,omitempty"`
		Metadata                 map[string]interface{}          `json:"metadata,omitempty"`
		Eci                      string                          `json:"eci,omitempty"`
		SchemeId                 string                          `json:"scheme_id,omitempty"`
		Actions                  []payments.PaymentActionSummary `json:"actions,omitempty"`
		Retry                    *payments.PaymentRetryResponse  `json:"retry,omitempty"`
		PanTypeProcessed         payments.PanProcessedType       `json:"pan_type_processed,omitempty"`
		CkoNetworkTokenAvailable bool                            `json:"cko_network_token_available,omitempty"`
		ProcessedOn              *time.Time                      `json:"processed_on,omitempty"`
		Instruction              *PaymentInstruction             `json:"instruction,omitempty"`
		Links                    map[string]common.Link          `json:"_links"`
	}

	GetPaymentActionsResponse struct {
		HttpMetadata common.HttpMetadata
		Actions      []PaymentAction `json:"actions,omitempty"`
	}

	PaymentAction struct {
		Id                string                       `json:"id,omitempty"`
		Type              payments.ActionType          `json:"type,omitempty"`
		ProcessedOn       *time.Time                   `json:"processed_on,omitempty"`
		Amount            int64                        `json:"amount,omitempty"`
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

	IncrementAuthorizationResponse struct {
		HttpMetadata    common.HttpMetadata
		ActionId        string                      `json:"action_id,omitempty"`
		Amount          int64                       `json:"amount,omitempty"`
		Currency        common.Currency             `json:"currency,omitempty"`
		Approved        bool                        `json:"approved,omitempty"`
		Status          payments.PaymentStatus      `json:"status,omitempty"`
		AuthCode        string                      `json:"auth_code,omitempty"`
		ResponseCode    string                      `json:"response_code,omitempty"`
		ResponseSummary string                      `json:"response_summary,omitempty"`
		ExpiresOn       *time.Time                  `json:"expires_on,omitempty"`
		Balances        *PaymentResponseBalances    `json:"balances,omitempty"`
		ProcessedOn     *time.Time                  `json:"processed_on,omitempty"`
		Reference       string                      `json:"reference,omitempty"`
		Processing      *payments.PaymentProcessing `json:"processing,omitempty"`
		Eci             string                      `json:"eci,omitempty"`
		SchemeId        string                      `json:"scheme_id,omitempty"`
		Links           map[string]common.Link      `json:"_links"`
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
