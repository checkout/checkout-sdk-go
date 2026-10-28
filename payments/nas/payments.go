package nas

import (
	"encoding/json"
	"time"

	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/payments"
	"github.com/checkout/checkout-sdk-go/v3/payments/nas/sources"
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

type PreferredExperiencesType string

const (
	GoogleSpaPET PreferredExperiencesType = "google_spa"
	ThreeDsPET   PreferredExperiencesType = "3ds"
)

type PaymentPurposeType string

const (
	DonationsPPT         PaymentPurposeType = "donations"
	EducationPPT         PaymentPurposeType = "education"
	EmergencyNeedPPT     PaymentPurposeType = "emergency_need"
	ExpatriationPPT      PaymentPurposeType = "expatriation"
	FamilySupportPPT     PaymentPurposeType = "family_support"
	FinancialServicesPPT PaymentPurposeType = "financial_services"
	GiftsPPT             PaymentPurposeType = "gifts"
	IncomePPT            PaymentPurposeType = "income"
	InsurancePPT         PaymentPurposeType = "insurance"
	InvestmentPPT        PaymentPurposeType = "investment"
	ItServicesPPT        PaymentPurposeType = "it_services"
	LeisurePPT           PaymentPurposeType = "leisure"
	LoanPaymentPPT       PaymentPurposeType = "loan_payment"
	MedicalTreatmentPPT  PaymentPurposeType = "medical_treatment"
	OtherPPT             PaymentPurposeType = "other"
	PensionPPT           PaymentPurposeType = "pension"
	RoyaltiesPPT         PaymentPurposeType = "royalties"
	SavingsPPT           PaymentPurposeType = "savings"
	TravelAndTourismPPT  PaymentPurposeType = "travel_and_tourism"
)

type (
	PayoutBillingDescriptor struct {
		Reference string `json:"reference,omitempty"`
	}

	CardPayoutItem struct {
		Type    payments.ItemType    `json:"type,omitempty"`
		SubType payments.ItemSubType `json:"sub_type,omitempty"`
	}

	CardPayoutProcessing struct {
		AffiliateId     string         `json:"affiliate_id,omitempty"`
		AffiliateUrl    string         `json:"affiliate_url,omitempty"`
		ProcessingSpeed string         `json:"processing_speed,omitempty"`
		PurchaseCountry common.Country `json:"purchase_country,omitempty"`
	}

	PaymentInstruction struct {
		Purpose           PaymentPurposeType        `json:"purpose,omitempty"`
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
		ValueDate *time.Time `json:"value_date,omitempty"`
		// FundsTransferType is the scheme's categorisation of the client,
		// for example FD, MT or AA.
		FundsTransferType string                 `json:"funds_transfer_type,omitempty"`
		Links             map[string]common.Link `json:"_links"`
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

	Authentication struct {
		PreferredExperiences []PreferredExperiencesType `json:"preferred_experiences,omitempty"`
	}

	PaymentRoutingAttempt struct {
		Scheme string `json:"scheme,omitempty"`
	}

	PaymentRouting struct {
		Attempts []PaymentRoutingAttempt `json:"attempts,omitempty"`
	}
)

// Request
type (
	PaymentRequest struct {
		// PaymentContextId is the Payment Context's unique identifier. Required when requesting a payment from a payment
		// context.
		// [Optional]
		// ^(pct)_(\w{26})$
		// max 30 characters
		// min 30 characters
		PaymentContextId string `json:"payment_context_id,omitempty"`
		// Currency is the three-letter ISO currency code.
		// [Required]
		// max 3 characters
		// min 3 characters
		Currency common.Currency `json:"currency,omitempty"`
		// Source is the source of the payment. You can encrypt this field by using the Cko-Encryption-Key-Id and
		// Cko-Encrypted-Properties headers
		// (https://api-reference.checkout.com/tag/payments/POST/payments.headers.Cko-Encryption-Key-Id).
		// [Optional]
		Source payments.PaymentSource `json:"source,omitempty"`
		// FallbackSource is a card payment source.
		// [Optional]
		FallbackSource payments.PaymentSource `json:"fallback_source,omitempty"`
		// Amount is the payment amount. To perform a card verification, do not provide the amount or provide a value of
		// 0. The amount must be provided in the minor currency unit.
		// [Optional]
		// Format: int64
		// min 0
		// max 9999999999
		Amount int64 `json:"amount"`
		// PaymentType is the type of payment. This field is required for card payments in which the cardholder is not
		// present. For example, mail orders, telephone orders, or merchant-initiated transactions (MITs) in a recurring
		// payment
		// (https://www.checkout.com/docs/payments/accept-payments/pay-with-stored-card-details/recurring-payments-with-stored-card-details)
		// series. For MITs, this field must not be set to Regular.
		// [Optional]
		// Default: Regular
		PaymentType payments.PaymentType `json:"payment_type,omitempty" default:"Regular"`
		// PaymentPlan is the details of a recurring subscription or installment.
		// [Optional]
		PaymentPlan payments.PaymentPlan `json:"payment_plan,omitempty"`
		// MerchantInitiated: Whether the payment is a merchant-initiated transaction (MIT). Must be set to true for all
		// MITs. If you set this field to true, the value for payment_type must not be set to Regular.
		// [Optional]
		MerchantInitiated bool `json:"merchant_initiated"`
		// Reference is a reference you can use to identify the payment. For example, an order number. - For Amex
		// payments, the reference has a 30-character limit. - For Benefit payments, the reference must be a unique
		// alphanumeric value. - For Blik payments, the reference has a 35-character limit. - For Discover payments, the
		// reference has a 60-character limit. - For iDEAL payments, the reference is required and must be an
		// alphanumeric value with a 35-character limit. - For OmanNet payments, the reference must be a unique
		// alphanumeric value between 12 and 30 characters long, generated for every payment action. - For seQura
		// payments, the reference is required and must be a unique value. - For TWINT payments, the reference has a 50
		// character limit.
		// [Optional]
		// max 80 characters
		Reference string `json:"reference,omitempty"`
		// Description is a description of the payment.
		// [Optional]
		// max 100 characters
		Description string `json:"description,omitempty"`
		// AuthorizationType is the authorization type.
		// [Optional]
		// Default: Final
		AuthorizationType AuthorizationType `json:"authorization_type,omitempty" default:"Final"`
		// PartialAuthorization: Required information to allow partial authorization.
		// [Optional]
		PartialAuthorization *PartialAuthorization `json:"partial_authorization,omitempty"`
		// Capture: Whether to capture the payment (if applicable).
		// [Optional]
		// Default: true
		Capture bool `json:"capture" default:"true"`
		// CaptureOn is a timestamp (ISO 8601 code) that determines when the payment should be captured. Providing this
		// field will automatically set capture to true.
		// [Optional]
		// Format: date-time
		CaptureOn *time.Time `json:"capture_on,omitempty"`
		// ExpireOn is the date and time when the Multibanco payment expires in UTC. Format: ISO 8601.
		// Example: 2025-01-31T10:20:30.456.
		// [Optional]
		// Format: date-time
		ExpireOn *time.Time `json:"expire_on,omitempty"`
		// Customer is the customer's details. Required if source.type is tamara.
		// [Optional]
		Customer *common.CustomerRequest `json:"customer,omitempty"`
		// BillingDescriptor is an optional description that is displayed on the customer's statement identifying a
		// purchase.
		// [Optional]
		BillingDescriptor *payments.BillingDescriptor `json:"billing_descriptor,omitempty"`
		// ShippingDetails is the shipping details.
		// [Optional]
		ShippingDetails *payments.ShippingDetails `json:"shipping,omitempty"`
		// ThreeDsRequest: Information required for 3D Secure authentication payments.
		// [Optional]
		ThreeDsRequest *payments.ThreeDsRequest `json:"3ds,omitempty"`
		// Authentication provides information required to authenticate payments.
		// [Optional]
		Authentication *Authentication `json:"authentication,omitempty"`
		// ProcessingChannelId is the processing channel to be used for the payment.
		// [Optional]
		// ^(pc)_(\w{26})$
		ProcessingChannelId string `json:"processing_channel_id,omitempty"`
		// PreviousPaymentId is an identifier that links the payment to an existing series of payments. You must only
		// pass this field if the transaction is a merchant-initiated transaction (MIT) in a recurring payment series. To
		// link the payment, pass one of the following as its value: - a payment identifier (for example,
		// pay_cr4hxwizzp6k7biycuk2ibltnm) from the recurring series, or - the scheme transaction ID.
		// [Optional]
		// max 100 characters
		PreviousPaymentId string `json:"previous_payment_id,omitempty"`
		// Risk configures the risk assessment performed during the processing of the payment.
		// [Optional]
		Risk *payments.RiskRequest `json:"risk,omitempty"`
		// SuccessUrl: For redirect payment methods, this overrides the default success redirect URL configured on your
		// account.
		// [Optional]
		// Format: uri
		// max 1024 characters
		SuccessUrl string `json:"success_url,omitempty"`
		// FailureUrl: For redirect payment methods, this overrides the default failure redirect URL configured on your
		// account.
		// [Optional]
		// Format: uri
		// max 1024 characters
		FailureUrl string `json:"failure_url,omitempty"`
		// PaymentIp is the IP address used to make the payment. Used by Checkout.com's risk engine to check the customer's
		// IP address, only accepts IPv4 and IPv6 addresses. Required if source.type is sequra.
		// [Optional]
		// Format: ipv4
		// max 45 characters
		//
		// Deprecated: Use risk.device.network.ipv4 or risk.device.network.ipv6 instead.
		// This field will be removed in the next major version.
		PaymentIp string `json:"payment_ip,omitempty"`
		// Sender: Information about the sender of the payment's funds.
		// [Optional]
		Sender Sender `json:"sender,omitempty"`
		// Recipient: Information about the recipient of the payment's funds. Relevant for Account Funding Transactions
		// (https://www.checkout.com/docs/payments/manage-payments/account-funding-transactions) and VISA or Mastercard
		// domestic UK transactions processed by financial institutions
		// (https://www.checkout.com/docs/payments/regulation-support/requirements-for-financial-institutions).
		// [Optional]
		Recipient *payments.PaymentRecipient `json:"recipient,omitempty"`
		// Marketplace holds the marketplace sub-entity data. Not documented in the current API reference.
		// [Optional]
		Marketplace *common.MarketplaceData `json:"marketplace,omitempty"`
		// AmountAllocations is the sub-entities that the payment is being processed on behalf of.
		// [Optional]
		AmountAllocations []common.AmountAllocations `json:"amount_allocations,omitempty"`
		// Processing: Use the processing object to influence or override the data sent during card processing.
		// [Optional]
		Processing *payments.ProcessingSettings `json:"processing,omitempty"`
		// Items is the order's line items.
		// [Optional]
		Items []payments.Product `json:"items,omitempty"`
		// Retry: Configuration of asynchronous retries. For more information about asynchronous retries and supported
		// response codes, see Configure scheduled retries
		// (https://www.checkout.com/docs/payments/manage-payments/configure-scheduled-retries).
		// [Optional]
		Retry *payments.PaymentRetryRequest `json:"retry,omitempty"`
		// Metadata stores additional information about a transaction with custom fields and up to five user-defined
		// fields, which you can use for reporting purposes. The object supports string, number, and boolean fields, but
		// not arrays or objects. You can provide up to 20 metadata fields per API call, but the value of each field must
		// not exceed 255 characters in length. You can also reference metadata properties
		// (https://www.checkout.com/docs/business-operations/prevent-fraud/understand-fraud-detection#Metadata) in your
		// custom rules for Fraud Detection. For example, $coupon_code = '1234’.
		// [Optional]
		Metadata map[string]interface{} `json:"metadata,omitempty"`
		// Segment is the dimension details about business segment for payment request. At least one dimension required.
		// [Optional]
		Segment *payments.PaymentSegment `json:"segment,omitempty"`
		// Instruction: Details about the payment instruction.
		// [Optional]
		Instruction *PaymentInstruction `json:"instruction,omitempty"`
		// Routing controls processor attempts at the payment level.
		// [Optional]
		Routing *PaymentRouting `json:"routing,omitempty"`
		// Subscription is the details of the subscription.
		// [Optional]
		Subscription *payments.Subscription `json:"subscription,omitempty"`
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
		PreviousPaymentId   string                   `json:"previous_payment_id,omitempty"`
		Items               []CardPayoutItem         `json:"items,omitempty"`
		Segment             *payments.PaymentSegment `json:"segment,omitempty"`
		Processing          *CardPayoutProcessing    `json:"processing,omitempty"`
		Metadata            map[string]interface{}   `json:"metadata,omitempty"`
	}

	IncrementAuthorizationRequest struct {
		Amount    int64                  `json:"amount,omitempty"`
		Reference string                 `json:"reference,omitempty"`
		Metadata  map[string]interface{} `json:"metadata,omitempty"`
	}

	CaptureRequest struct {
		Amount            int64                               `json:"amount,omitempty"`
		CaptureType       CaptureType                         `json:"capture_type,omitempty"`
		Reference         string                              `json:"reference,omitempty"`
		Customer          *common.CustomerRequest             `json:"customer,omitempty"`
		Description       string                              `json:"description,omitempty"`
		BillingDescriptor *payments.BillingDescriptor         `json:"billing_descriptor,omitempty"`
		Shipping          *payments.CaptureShipping           `json:"shipping,omitempty"`
		Items             []payments.Product                  `json:"items,omitempty"`
		Marketplace       *common.MarketplaceData             `json:"marketplace,omitempty"`
		AmountAllocations []common.AmountAllocations          `json:"amount_allocations,omitempty"`
		Processing        *payments.CaptureProcessingSettings `json:"processing,omitempty"`
		Metadata          map[string]interface{}              `json:"metadata,omitempty"`
	}

	SearchPaymentsRequest struct {
		Query string     `json:"query,omitempty"`
		Limit int        `json:"limit,omitempty"`
		From  *time.Time `json:"from,omitempty"`
		To    *time.Time `json:"to,omitempty"`
	}
)

// Response
type (
	PaymentResponse struct {
		HttpMetadata    common.HttpMetadata
		Id              string                         `json:"id,omitempty"`
		ActionId        string                         `json:"action_id,omitempty"`
		Amount          int64                          `json:"amount,omitempty"`
		Currency        common.Currency                `json:"currency,omitempty"`
		Approved        bool                           `json:"approved,omitempty"`
		Status          payments.PaymentStatus         `json:"status,omitempty"`
		ResponseCode    string                         `json:"response_code,omitempty"`
		ProcessedOn     *time.Time                     `json:"processed_on,omitempty"`
		Links           map[string]common.Link         `json:"_links"`
		PaymentType     payments.PaymentType           `json:"payment_type,omitempty" default:"Regular"`
		PaymentPlan     payments.PaymentPlan           `json:"payment_plan,omitempty"`
		AmountRequested int                            `json:"amount_requested,omitempty"`
		AuthCode        string                         `json:"auth_code,omitempty"`
		ResponseSummary string                         `json:"response_summary,omitempty"`
		ExpiresOn       *time.Time                     `json:"expires_on,omitempty"`
		ThreeDs         *payments.ThreeDsEnrollment    `json:"3ds,omitempty"`
		Risk            *payments.RiskAssessment       `json:"risk,omitempty"`
		Source          *SourceResponse                `json:"source,omitempty"`
		Customer        *common.CustomerResponse       `json:"customer,omitempty"`
		Balances        *PaymentResponseBalances       `json:"balances,omitempty"`
		Reference       string                         `json:"reference,omitempty"`
		Processing      *payments.PaymentProcessing    `json:"processing,omitempty"`
		Eci             string                         `json:"eci,omitempty"`
		SchemeId        string                         `json:"scheme_id,omitempty"`
		Retry           *payments.PaymentRetryResponse `json:"retry,omitempty"`
	}

	PayoutResponse struct {
		HttpMetadata common.HttpMetadata
		Id           string                      `json:"id,omitempty"`
		Status       payments.PaymentStatus      `json:"status,omitempty"`
		Reference    string                      `json:"reference,omitempty"`
		Destination  *DestinationResponse        `json:"destination,omitempty"`
		Instruction  *PaymentInstructionResponse `json:"instruction,omitempty"`
	}

	GetPaymentResponse struct {
		// HttpMetadata holds the HTTP status code and headers of the response. Not part of the JSON body.
		HttpMetadata common.HttpMetadata
		// Id: Payment unique identifier.
		// [Required]
		// ^(pay)_(\w{26})$
		// max 30 characters
		// min 30 characters
		Id string `json:"id,omitempty"`
		// RequestedOn is the date/time the payment was requested.
		// [Required]
		// Format: date-time
		RequestedOn *time.Time `json:"requested_on,omitempty"`
		// Source is the source of the payment.
		// [Optional]
		Source *SourceResponse `json:"source,omitempty"`
		// Destination is the destination of the payment.
		// [Optional]
		Destination *DestinationResponse `json:"destination,omitempty"`
		// Sender: Information about the sender of the payment's funds.
		// [Optional]
		Sender *SenderResponse `json:"sender,omitempty"`
		// Amount is the original payment amount.
		// [Required]
		Amount int64 `json:"amount,omitempty"`
		// Currency is the three-letter ISO currency code of the payment.
		// [Required]
		// max 3 characters
		// min 3 characters
		Currency common.Currency `json:"currency,omitempty"`
		// PaymentType: This must be specified for card payments where the cardholder is not present (For example,
		// recurring or mail order / telephone order).
		// [Optional]
		// Default: Regular
		PaymentType payments.PaymentType `json:"payment_type,omitempty"`
		// Reference: Your reference for the payment.
		// [Optional]
		Reference string `json:"reference,omitempty"`
		// Description is a description of the payment.
		// [Optional]
		Description string `json:"description,omitempty"`
		// Approved: Whether the payment was successful.
		// [Required]
		Approved bool `json:"approved,omitempty"`
		// ExpiresOn is the timestamp (ISO 8601 code) for when the authorization's validity period expires.
		// [Optional]
		ExpiresOn *time.Time `json:"expires_on,omitempty"`
		// Status is the status of the payment.
		// [Required]
		Status payments.PaymentStatus `json:"status,omitempty"`
		// Balances is the payment balances.
		// [Optional]
		Balances *PaymentResponseBalances `json:"balances,omitempty"`
		// ThreeDs provides information relating to the processing of 3D Secure payments.
		// [Optional]
		ThreeDs *payments.ThreeDsData `json:"3ds,omitempty"`
		// Risk returns the payments risk assessment results.
		// [Optional]
		Risk *payments.RiskAssessment `json:"risk,omitempty"`
		// Customer is the customer to which this payment is linked.
		// [Optional]
		Customer *common.CustomerResponse `json:"customer,omitempty"`
		// BillingDescriptor is an optional description that is displayed on the customer's statement identifying a
		// purchase.
		// [Optional]
		BillingDescriptor *payments.BillingDescriptor `json:"billing_descriptor,omitempty"`
		// ShippingDetails is the payment shipping details.
		// [Optional]
		ShippingDetails *payments.ShippingDetails `json:"shipping,omitempty"`
		// PaymentIp is the IP address used to make the payment. Used by Checkout.com's risk engine to check the
		// customer's IP address, only accepts IPv4 and IPv6 addresses.
		// [Optional]
		// Format: ipv4
		// max 45 characters
		PaymentIp string `json:"payment_ip,omitempty"`
		// Marketplace holds the marketplace sub-entity data. Not documented in the current API reference.
		// [Optional]
		Marketplace *common.MarketplaceData `json:"marketplace,omitempty"`
		// AmountAllocations is the sub-entities that the payment is being processed on behalf of.
		// [Optional]
		AmountAllocations []common.AmountAllocations `json:"amount_allocations,omitempty"`
		// Recipient: Information about the recipient of the payment's funds. Relevant for Account Funding Transactions
		// (https://www.checkout.com/docs/payments/manage-payments/account-funding-transactions) and VISA or Mastercard
		// domestic UK transactions processed by financial institutions
		// (https://www.checkout.com/docs/payments/regulation-support/requirements-for-financial-institutions).
		// [Optional]
		Recipient *payments.PaymentRecipient `json:"recipient,omitempty"`
		// ProcessingData returns information related to the processing of the payment.
		// [Optional]
		ProcessingData *payments.ProcessingData `json:"processing,omitempty"`
		// Items is the order's line items.
		// [Optional]
		Items []payments.Product `json:"items,omitempty"`
		// Metadata is a set of key-value pairs that you can attach to a payment. It can be useful for storing additional
		// information in a structured format.
		// [Optional]
		Metadata map[string]interface{} `json:"metadata,omitempty"`
		// Eci is the final Electronic Commerce Indicator (ECI) security level used to authorize the payment. Applicable
		// for 3D Secure and network token payments.
		// [Optional]
		Eci string `json:"eci,omitempty"`
		// SchemeId is the scheme transaction identifier.
		// [Optional]
		SchemeId string `json:"scheme_id,omitempty"`
		// Actions is a summary of the payment's actions, returned when a session ID is used to get the payment details.
		// [Optional]
		Actions []payments.PaymentActionSummary `json:"actions,omitempty"`
		// Retry: Configuration relating to asynchronous retries.
		// [Optional]
		Retry *payments.PaymentRetryResponse `json:"retry,omitempty"`
		// PanTypeProcessed is the type of Primary Account Number (PAN) used for the payment. DPAN indicates network
		// token was used, FPAN indicates the full card was used. The API reference documents it under
		// processing.pan_type_processed.
		// [Optional]
		PanTypeProcessed payments.PanProcessedType `json:"pan_type_processed,omitempty"`
		// CkoNetworkTokenAvailable is the flag indicating if Checkout Network Token was available for the payment.
		// [Optional]
		CkoNetworkTokenAvailable bool `json:"cko_network_token_available,omitempty"`
		// ProcessedOn: If the scheme has finished processing the payout request, the UTC date and time the payout was
		// approved or declined. Returned for payouts.
		// [Optional]
		// Format: date-time
		ProcessedOn *time.Time `json:"processed_on,omitempty"`
		// Instruction: Details about the payment instruction.
		// [Optional]
		Instruction *PaymentInstruction `json:"instruction,omitempty"`
		// Links is the links related to the payment.
		// [Required]
		Links map[string]common.Link `json:"_links"`
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

	SearchPaymentsResponse struct {
		HttpMetadata common.HttpMetadata
		Data         []GetPaymentResponse   `json:"data,omitempty"`
		Links        map[string]common.Link `json:"_links"`
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
