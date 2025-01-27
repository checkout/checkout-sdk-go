package payments

import (
	"time"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/sessions"
)

const PathPayments = "payments"

type PaymentType string

const (
	Regular     PaymentType = "Regular"
	Recurring   PaymentType = "Recurring"
	MOTO        PaymentType = "MOTO"
	Installment PaymentType = "Installment"
	Unscheduled PaymentType = "Unscheduled"
)

type PaymentStatus string

const (
	Active            PaymentStatus = "Active"
	Pending           PaymentStatus = "Pending"
	Authorized        PaymentStatus = "Authorized"
	CardVerified      PaymentStatus = "Card Verified"
	Voided            PaymentStatus = "Voided"
	PartiallyCaptured PaymentStatus = "Partially Captured"
	Captured          PaymentStatus = "Captured"
	PartiallyRefunded PaymentStatus = "Partially Refunded"
	Refunded          PaymentStatus = "Refunded"
	Declined          PaymentStatus = "Declined"
	Canceled          PaymentStatus = "Canceled"
	Expired           PaymentStatus = "Expired"
	Requested         PaymentStatus = "Requested"
	Paid              PaymentStatus = "Paid"
)

type Exemption string

const (
	None                      Exemption = "none"
	LowValue                  Exemption = "low_value"
	RecurringOperation        Exemption = "recurring_operation"
	TransactionRiskAssessment Exemption = "transaction_risk_assessment"
	SecureCorporatePayment    Exemption = "secure_corporate_payment"
	TrustedListing            Exemption = "trusted_listing"
	ThreeDsOutage             Exemption = "3ds_outage"
	ScaDelegation             Exemption = "sca_delegation"
	OutOfScaScope             Exemption = "out_of_sca_scope"
	Other                     Exemption = "other"
	LowRiskProgram            Exemption = "low_risk_program"
)

type ThreeDsEnrollmentStatus string

const (
	Yes ThreeDsEnrollmentStatus = "Y"
	No  ThreeDsEnrollmentStatus = "N"
	U   ThreeDsEnrollmentStatus = "U"
)

type ActionType string

const (
	AuthorizationYes ActionType = "Authorization"
	CardVerification ActionType = "Card Verification"
	Void             ActionType = "Void"
	Capture          ActionType = "Capture"
	Refund           ActionType = "Refund"
	Payout           ActionType = "Payout"
	Return           ActionType = "Return"
)

type NetworkTokenType string

const (
	Vts       NetworkTokenType = "vts"
	Mdes      NetworkTokenType = "mdes"
	ApplePay  NetworkTokenType = "applepay"
	GooglePay NetworkTokenType = "googlepay"
)

type PreferredSchema string

const (
	Visa            PreferredSchema = "visa"
	Mastercard      PreferredSchema = "mastercard"
	CartesBancaires PreferredSchema = "cartes_bancaires"
)

type ProductType string

const (
	QrCode          ProductType = "QR Code"
	InApp           ProductType = "In-App"
	OfficialAccount ProductType = "Official Account"
	MiniProgram     ProductType = "Mini Program"
)

type MerchantInitiatedReason string

const (
	DelayedCharge   MerchantInitiatedReason = "Delayed_charge"
	Resubmission    MerchantInitiatedReason = "Resubmission"
	NoShow          MerchantInitiatedReason = "No_show"
	Reauthorization MerchantInitiatedReason = "Reauthorization"
)

type TerminalType string

const (
	App TerminalType = "APP"
	Wap TerminalType = "WAP"
	Web TerminalType = "WEB"
)

type OsType string

const (
	Android OsType = "ANDROID"
	Ios     OsType = "IOS"
)

type ShippingPreference string

const (
	NoShipping         ShippingPreference = "NO_SHIPPING"
	SetProvidedAddress ShippingPreference = "SET_PROVIDED_ADDRESS"
	GetFromFile        ShippingPreference = "GET_FROM_FILE"
)

type UserAction string

const (
	PayNow   UserAction = "PAY_NOW"
	Continue UserAction = "CONTINUE"
)

type FundTransferType string

const (
	AA  FundTransferType = "AA"
	PP  FundTransferType = "PP"
	FT  FundTransferType = "FT"
	FD  FundTransferType = "FD"
	PD  FundTransferType = "PD"
	LO  FundTransferType = "LO"
	OG  FundTransferType = "OG"
	CO4 FundTransferType = "CO4"
	CO7 FundTransferType = "CO7"
	C52 FundTransferType = "C52"
	C55 FundTransferType = "C55"
)

type PaymentContextsShippingMethod string

const (
	Digital        PaymentContextsShippingMethod = "Digital"
	PickUp         PaymentContextsShippingMethod = "PickUp"
	BillingAddress PaymentContextsShippingMethod = "BillingAddress"
	OtherAddress   PaymentContextsShippingMethod = "OtherAddress"
)

type PanProcessedType string

const (
	FPAN PanProcessedType = "fpan"
	DPAN PanProcessedType = "dpan"
)

type (
	AirlineData struct {
		Ticket           *Ticket            `json:"ticket,omitempty"`
		Passenger        *Passenger         `json:"passenger,omitempty"`
		FlightLegDetails []FlightLegDetails `json:"flight_leg_details,omitempty"`
	}

	Ticket struct {
		Number             string `json:"number,omitempty"`
		IssueDate          string `json:"issue_date,omitempty"`
		IssuingCarrierCode string `json:"issuing_carrier_code,omitempty"`
		TravelAgencyName   string `json:"travel_agency_name,omitempty"`
		TravelAgencyCode   string `json:"travel_agency_code,omitempty"`
	}

	Passenger struct {
		Name        *PassengerName `json:"name,omitempty"`
		DateOfBirth string         `json:"date_of_birth,omitempty"`
		CountryCode common.Country `json:"country_code,omitempty"`
	}

	PassengerName struct {
		FullName string `json:"full_name,omitempty"`
	}

	FlightLegDetails struct {
		FlightNumber     int64  `json:"flight_number,omitempty"`
		CarrierCode      string `json:"carrier_code,omitempty"`
		ServiceClass     string `json:"service_class,omitempty"`
		DepartureDate    string `json:"departure_date,omitempty"`
		DepartureTime    string `json:"departure_time,omitempty"`
		DepartureAirport string `json:"departure_airport,omitempty"`
		ArrivalAirport   string `json:"arrival_airport,omitempty"`
		StopoverCode     string `json:"stopover_code,omitempty"`
		FareBasisCode    string `json:"fare_basis_code,omitempty"`
	}

	ShippingInfo struct {
		ShippingCompany       string `json:"shipping_company,omitempty"`
		ShippingMethod        string `json:"shipping_method,omitempty"`
		TrackingNumber        string `json:"tracking_number,omitempty"`
		TrackingUri           string `json:"tracking_uri,omitempty"`
		ReturnShippingCompany string `json:"return_shipping_company,omitempty"`
		ReturnTrackingNumber  string `json:"return_tracking_number,omitempty"`
		ReturnTrackingUri     string `json:"return_tracking_uri,omitempty"`
	}

	ShippingDetails struct {
		FirstName      string                        `json:"first_name,omitempty"`
		LastName       string                        `json:"last_name,omitempty"`
		Email          string                        `json:"email,omitempty"`
		Address        *common.Address               `json:"address,omitempty"`
		Phone          *common.Phone                 `json:"phone,omitempty"`
		FromAddressZip string                        `json:"from_address_zip,omitempty"`
		Timeframe      sessions.DeliveryTimeframe    `json:"time_frame,omitempty"`
		Method         PaymentContextsShippingMethod `json:"method,omitempty"`
		Delay          int                           `json:"delay,omitempty"`
	}

	PaymentSegment struct {
		Brand            string `json:"brand,omitempty"`
		BusinessCategory string `json:"business_category,omitempty"`
		Market           string `json:"market,omitempty"`
	}

	PaymentRetryRequest struct {
		Enabled      bool `json:"enabled,omitempty"`
		MaxAttempts  int  `json:"max_attempts,omitempty"`
		EndAfterDays int  `json:"end_after_days,omitempty"`
	}

	BillingDescriptor struct {
		Name string `json:"name,omitempty"`
		City string `json:"city,omitempty"`
		// Not available on Previous
		Reference string `json:"reference,omitempty"`
	}

	ThreeDsRequest struct {
		Enabled            bool                      `json:"enabled"`
		AttemptN3D         bool                      `json:"attempt_n3d"`
		Eci                string                    `json:"eci,omitempty"`
		Cryptogram         string                    `json:"cryptogram,omitempty"`
		Xid                string                    `json:"xid,omitempty"`
		Version            string                    `json:"version,omitempty"`
		Exemption          Exemption                 `json:"exemption,omitempty"`
		ChallengeIndicator common.ChallengeIndicator `json:"challenge_indicator,omitempty"`
		AllowUpgrade       bool                      `json:"allow_upgrade,omitempty"`
		// Not available on Previous
		Status                string                 `json:"status,omitempty"`
		AuthenticationDate    *time.Time             `json:"authentication_date,omitempty"`
		AuthenticationAmount  float64                `json:"authentication_amount,omitempty"`
		FlowType              common.ThreeDsFlowType `json:"flow_type,omitempty"`
		StatusReasonCode      string                 `json:"status_reason_code,omitempty"`
		ChallengeCancelReason string                 `json:"challenge_cancel_reason,omitempty"`
		Score                 string                 `json:"score,omitempty"`
		CryptogramAlgorithm   string                 `json:"cryptogram_algorithm,omitempty"`
		AuthenticationId      string                 `json:"authentication_id,omitempty"`
	}

	RiskRequest struct {
		Enabled         bool   `json:"enabled"`
		DeviceSessionId string `json:"device_session_id,omitempty"`
	}

	PaymentRecipient struct {
		DateOfBirth   string          `json:"dob,omitempty"`
		AccountNumber string          `json:"account_number,omitempty"`
		Address       *common.Address `json:"address,omitempty"`
		CountryCode   common.Country  `json:"country_code,omitempty"`
		Zip           string          `json:"zip,omitempty"`
		FirstName     string          `json:"first_name,omitempty"`
		LastName      string          `json:"last_name,omitempty"`
	}

	ProcessingSettings struct {
		OrderId                 string                  `json:"order_id,omitempty"`
		TaxAmount               int64                   `json:"tax_amount,omitempty"`
		DiscountAmount          int64                   `json:"discount_amount,omitempty"`
		DutyAmount              int64                   `json:"duty_amount,omitempty"`
		ShippingAmount          int64                   `json:"shipping_amount,omitempty"`
		ShippingTaxAmount       int64                   `json:"shipping_tax_amount,omitempty"`
		Aft                     bool                    `json:"aft,omitempty"`
		PreferredScheme         PreferredSchema         `json:"preferred_scheme,omitempty"`
		MerchantInitiatedReason MerchantInitiatedReason `json:"merchant_initiated_reason,omitempty"`
		CampaignId              int64                   `json:"campaign_id,omitempty"`
		ProductType             ProductType             `json:"product_type,omitempty"`
		OpenId                  string                  `json:"open_id,omitempty"`
		OriginalOrderAmount     int64                   `json:"original_order_amount,omitempty"`
		ReceiptId               string                  `json:"receipt_id,omitempty"`
		TerminalType            TerminalType            `json:"terminal_type,omitempty"`
		OsType                  OsType                  `json:"os_type,omitempty"`
		InvoiceId               string                  `json:"invoice_id,omitempty"`
		BrandName               string                  `json:"brand_name,omitempty"`
		Locale                  string                  `json:"locale,omitempty"`
		ShippingPreference      ShippingPreference      `json:"shipping_preference,omitempty"`
		UserAction              UserAction              `json:"user_action,omitempty"`
		SetTransactionContext   []map[string]string     `json:"set_transaction_context,omitempty"`
		AirlineData             []AirlineData           `json:"airline_data,omitempty"`
		OtpValue                string                  `json:"otp_value,omitempty"`
		PurchaseCountry         common.Country          `json:"purchase_country,omitempty"`
		CustomPaymentMethodIds  []string                `json:"custom_payment_method_ids,omitempty"`
		MerchantCallbackUrl     string                  `json:"merchant_callback_url,omitempty"`
		ShippingDelay           int64                   `json:"shipping_delay,omitempty"`
		ShippingInfo            string                  `json:"shipping_info,omitempty"`
		LineOfBusiness          string                  `json:"line_of_business,omitempty"`
		// Only available on Previous
		SenderInformation *SenderInformation        `json:"senderInformation,omitempty"`
		Purpose           string                    `json:"purpose,omitempty"`
		Dlocal            *DLocalProcessingSettings `json:"dlocal,omitempty"`
	}

	ThreeDsEnrollment struct {
		Downgraded    bool                    `json:"downgraded,omitempty"`
		Enrolled      ThreeDsEnrollmentStatus `json:"enrolled,omitempty"`
		UpgradeReason string                  `json:"upgrade_reason,omitempty"`
	}

	RiskAssessment struct {
		Flagged bool    `json:"flagged,omitempty"`
		Score   float64 `json:"score,omitempty"`
	}

	PaymentProcessing struct {
		RetrievalReferenceNumber         string           `json:"retrieval_reference_number,omitempty"`
		AcquirerTransactionId            string           `json:"acquirer_transaction_id,omitempty"`
		RecommendationCode               string           `json:"recommendation_code,omitempty"`
		Scheme                           string           `json:"scheme,omitempty"`
		PartnerMerchantAdviceCode        string           `json:"partner_merchant_advice_code,omitempty"`
		PartnerResponseCode              string           `json:"partner_response_code,omitempty"`
		PartnerOrderId                   string           `json:"partner_order_id,omitempty"`
		PartnerSessionId                 string           `json:"partner_session_id,omitempty"`
		PartnerClientToken               string           `json:"partner_client_token,omitempty"`
		PartnerPaymentId                 string           `json:"partner_payment_id,omitempty"`
		PanTypeProcessed                 PanProcessedType `json:"pan_type_processed,omitempty"`
		ContinuationPayload              string           `json:"continuation_payload,omitempty"`
		Pun                              string           `json:"pun,omitempty"`
		PartnerStatus                    string           `json:"partner_status,omitempty"`
		PartnerTransactionId             string           `json:"partner_transaction_id,omitempty"`
		PartnerErrorCodes                []string         `json:"partner_error_codes,omitempty"`
		PartnerErrorMessage              string           `json:"partner_error_message,omitempty"`
		PartnerAuthorizationCode         string           `json:"partner_authorization_code,omitempty"`
		PartnerAuthorizationResponseCode string           `json:"partner_authorization_response_code,omitempty"`
		SurchargeAmount                  int64            `json:"surcharge_amount,omitempty"`
		CkoNetworkTokenAvailable         bool             `json:"cko_network_token_available"`
		MerchantCategoryCode             string           `json:"merchant_category_code,omitempty"`
	}

	PaymentRetryResponse struct {
		MaxAttempts   int        `json:"max_attempts,omitempty"`
		EndsOn        *time.Time `json:"ends_on,omitempty"`
		NextAttemptOn *time.Time `json:"next_attempt_on,omitempty"`
	}

	ThreeDsData struct {
		Downgraded                 bool      `json:"downgraded,omitempty"`
		Enrolled                   string    `json:"enrolled,omitempty"`
		UpgradeReason              string    `json:"upgrade_reason,omitempty"`
		SignatureValid             string    `json:"signature_valid,omitempty"`
		AuthenticationResponse     string    `json:"authentication_response,omitempty"`
		AuthenticationStatusReason string    `json:"authentication_status_reason,omitempty"`
		Cryptogram                 string    `json:"cryptogram,omitempty"`
		Xid                        string    `json:"xid,omitempty"`
		Version                    string    `json:"version,omitempty"`
		Exemption                  Exemption `json:"exemption,omitempty"`
		ExemptionApplied           string    `json:"exemption_applied,omitempty"`
		Challenged                 bool      `json:"challenged,omitempty"`
	}

	PaymentActionSummary struct {
		Id              string     `json:"id,omitempty"`
		Type            ActionType `json:"type,omitempty"`
		ResponseCode    string     `json:"response_code,omitempty"`
		ResponseSummary string     `json:"response_summary,omitempty"`
	}

	Processing struct {
		AcquirerReferenceNumber  string     `json:"acquirer_reference_number,omitempty"`
		RetrievalReferenceNumber ActionType `json:"retrieval_reference_number,omitempty"`
		AcquirerTransactionId    string     `json:"acquirer_transaction_id,omitempty"`
	}

	Installments struct {
		Count string `json:"count,omitempty"`
	}

	SenderInformation struct {
		Reference     string         `json:"reference,omitempty"`
		FirstName     string         `json:"firstName,omitempty"`
		LastName      string         `json:"lastName,omitempty"`
		Dob           string         `json:"dob,omitempty"`
		Address       string         `json:"address,omitempty"`
		City          string         `json:"city,omitempty"`
		State         string         `json:"state,omitempty"`
		Country       common.Country `json:"country,omitempty"`
		PostalCode    string         `json:"postalCode,omitempty"`
		SourceOfFunds string         `json:"sourceOfFunds,omitempty"`
		Purpose       string         `json:"purpose,omitempty"`
	}

	DLocalProcessingSettings struct {
		Country      common.Country `json:"country,omitempty"`
		Payer        Payer          `json:"payer,omitempty"`
		Installments *Installments  `json:"installments,omitempty"`
	}

	Product struct {
		Name           string `json:"name,omitempty"`
		Quantity       int    `json:"quantity,omitempty"`
		UnitPrice      int    `json:"unit_price,omitempty"`
		Price          int    `json:"price,omitempty"`
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

	BillingInformation struct {
		Address *common.Address `json:"address,omitempty"`
		Phone   *common.Phone   `json:"phone,omitempty"`
	}

	Order struct {
		Name           string     `json:"name,omitempty"`
		Quantity       int64      `json:"quantity,omitempty"`
		UnitPrice      int64      `json:"unit_price,omitempty"`
		Reference      string     `json:"reference,omitempty"`
		CommodityCode  string     `json:"commodity_code,omitempty"`
		UnitOfMeasure  string     `json:"unit_of_measure,omitempty"`
		TotalAmount    int64      `json:"total_amount,omitempty"`
		TaxAmount      int64      `json:"tax_amount,omitempty"`
		DiscountAmount int64      `json:"discount_amount,omitempty"`
		WxpayGoodsId   string     `json:"wxpay_goods_id,omitempty"`
		ImageUrl       string     `json:"image_url,omitempty"`
		Url            string     `json:"url,omitempty"`
		Type           string     `json:"type,omitempty"`
		ServiceEndsOn  *time.Time `json:"service_ends_on,omitempty"`
	}

	PaymentMethodDetails struct {
		DisplayName string `json:"display_name,omitempty"`
		Type        string `json:"type,omitempty"`
		Network     string `json:"network,omitempty"`
	}
)

// Request
type (
	RefundRequest struct {
		Amount    int64                  `json:"amount,omitempty"`
		Reference string                 `json:"reference,omitempty"`
		Metadata  map[string]interface{} `json:"metadata,omitempty"`
		// Not available on Previous
		AmountAllocations []common.AmountAllocations `json:"amount_allocations,omitempty"`
		CaptureActionId   string                     `json:"capture_action_id,omitempty"`
		Destination       *common.Destination        `json:"destination,omitempty"`
		Items             []Order                    `json:"items,omitempty"`
	}

	VoidRequest struct {
		Reference string                 `json:"reference,omitempty"`
		Metadata  map[string]interface{} `json:"metadata,omitempty"`
	}

	QueryRequest struct {
		Limit     int    `url:"limit,10"`
		Skip      int    `url:"skip,0"`
		Reference string `url:"reference,omitempty"`
	}
)

// Response
type (
	CaptureResponse struct {
		HttpMetadata common.HttpMetadata
		ActionId     string                 `json:"action_id,omitempty"`
		Reference    string                 `json:"reference,omitempty"`
		Links        map[string]common.Link `json:"_links"`
	}

	RefundResponse struct {
		HttpMetadata common.HttpMetadata
		ActionId     string                 `json:"action_id,omitempty"`
		Reference    string                 `json:"reference,omitempty"`
		Links        map[string]common.Link `json:"_links"`
	}

	VoidResponse struct {
		HttpMetadata common.HttpMetadata
		ActionId     string                 `json:"action_id,omitempty"`
		Reference    string                 `json:"reference,omitempty"`
		Links        map[string]common.Link `json:"_links"`
	}

	ProcessingData struct {
		PreferredScheme                  PreferredSchema                  `json:"preferred_scheme,omitempty"`
		AppId                            string                           `json:"app_id,omitempty"`
		PartnerCustomerId                string                           `json:"partner_customer_id,omitempty"`
		PartnerPaymentId                 string                           `json:"partner_payment_id,omitempty"`
		TaxAmount                        int64                            `json:"tax_amount,omitempty"`
		PurchaseCountry                  common.Country                   `json:"purchase_country,omitempty"`
		Locale                           string                           `json:"locale,omitempty"`
		RetrievalReferenceNumber         string                           `json:"retrieval_reference_number,omitempty"`
		PartnerOrderId                   string                           `json:"partner_order_id,omitempty"`
		PartnerStatus                    string                           `json:"partner_status,omitempty"`
		PartnerTransactionId             string                           `json:"partner_transaction_id,omitempty"`
		PartnerErrorCodes                string                           `json:"partner_error_codes,omitempty"`
		PartnerErrorMessage              string                           `json:"partner_error_message,omitempty"`
		PartnerAuthorizationCode         string                           `json:"partner_authorization_code,omitempty"`
		PartnerAuthorizationResponseCode string                           `json:"partner_authorization_response_code,omitempty"`
		FraudStatus                      string                           `json:"fraud_status,omitempty"`
		ProviderAuthorizedPaymentMethod  *ProviderAuthorizedPaymentMethod `json:"provider_authorized_payment_method,omitempty"`
		CustomPaymentMethodIds           []string                         `json:"custom_payment_method_ids,omitempty"`
		Aft                              bool                             `json:"aft,omitempty"`
		MerchantCategoryCode             string                           `json:"merchant_category_code,omitempty"`
		SchemeMerchantId                 string                           `json:"scheme_merchant_id,omitempty"`
		PanTypeProcessed                 PanProcessedType                 `json:"pan_type_processed,omitempty"`
		CkoNetworkTokenAvailable         bool                             `json:"cko_network_token_available,omitempty"`
	}

	ProviderAuthorizedPaymentMethod struct {
		Type                 string `json:"type,omitempty"`
		Description          string `json:"description,omitempty"`
		NumberOfInstallments int64  `json:"number_of_installments,omitempty"`
		NumberOfDays         int64  `json:"number_of_days,omitempty"`
	}
)

type (
	DestinationTypeMapping struct {
		Destination string `json:"destination"`
	}

	SenderTypeMapping struct {
		Sender string `json:"sender"`
	}
)
