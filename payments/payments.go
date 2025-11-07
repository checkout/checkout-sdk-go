package payments

import (
	"time"

	"github.com/checkout/checkout-sdk-go/common"
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
	QrCode           ProductType = "QR Code"
	InApp            ProductType = "In-App"
	OfficialAccount  ProductType = "Official Account"
	MiniProgram      ProductType = "Mini Program"
	PayInFull        ProductType = "pay_in_full"
	PayByInstalment  ProductType = "pay_by_instalment"
	PayByInstalment2 ProductType = "pay_by_instalment2"
	PayByInstalment3 ProductType = "pay_by_instalment3"
	PayByInstalment4 ProductType = "pay_by_instalment4"
	PayByInstalment6 ProductType = "pay_by_instalment6"
	Invoice          ProductType = "invoice"
	PayLater         ProductType = "pay_later"
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

type StorePaymentDetailsType string

const (
	Disabled StorePaymentDetailsType = "disabled"
	Enabled  StorePaymentDetailsType = "enabled"
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

type LocalType string

const (
	ArLT    LocalType = "ar"
	DaDKLT  LocalType = "da-DK"
	DeDELT  LocalType = "de-DE"
	ElELLT  LocalType = "el"
	EnGBLT  LocalType = "en-GB"
	EsESLT  LocalType = "es-ES"
	FiFILT  LocalType = "fi-FI"
	FilPhLT LocalType = "fil-PH"
	FrFRLT  LocalType = "fr-FR"
	HiInLT  LocalType = "hi-IN"
	IdIDLT  LocalType = "id-ID"
	ItITLT  LocalType = "it-IT"
	JaJpLT  LocalType = "ja-JP"
	KoKRLT  LocalType = "ko-KR"
	MsMYLT  LocalType = "ms-MY"
	NbNOLT  LocalType = "nb-NO"
	NlNLLT  LocalType = "nl-NL"
	PtPTLT  LocalType = "pt-PT"
	SvSELT  LocalType = "sv-SE"
	ThTHLT  LocalType = "th-TH"
	ViVNLT  LocalType = "vi-VN"
	ZhCNLT  LocalType = "zh-CN"
	ZhHkLT  LocalType = "zh-HK"
	ZhTWLT  LocalType = "zh-TW"
)

type AmountVariabilityType string

const (
	FixedAVT    AmountVariabilityType = "Fixed"
	VariableAVT AmountVariabilityType = "Variable"
)

type CharacterSetType string

const (
	KANJICST    CharacterSetType = "kanji"
	KATAKANACST CharacterSetType = "katakana"
)

type ThreeDsReqAuthMethodType string

const (
	ThreeDsFederatedId                              ThreeDsReqAuthMethodType = "federated_id"
	ThreeDsFidoAuthenticator                        ThreeDsReqAuthMethodType = "fido_authenticator"
	ThreeDsFidoAuthenticatorFidoAssuranceDataSigned ThreeDsReqAuthMethodType = "fido_authenticator_fido_assurance_data_signed"
	ThreeDsIssuerCredentials                        ThreeDsReqAuthMethodType = "issuer_credentials"
	ThreeDsNoAuthenticationOccurred                 ThreeDsReqAuthMethodType = "no_threeds_requestor_authentication_occurred"
	ThreeDsSrcAssuranceData                         ThreeDsReqAuthMethodType = "src_assurance_data"
	ThreeDsOwnCredentials                           ThreeDsReqAuthMethodType = "three3ds_requestor_own_credentials"
	ThreeDsThirdPartyAuthentication                 ThreeDsReqAuthMethodType = "third_party_authentication"
)

type AuthenticationMethodType string

const (
	FederatedId              AuthenticationMethodType = "federated_id"
	Fido                     AuthenticationMethodType = "fido"
	IssuerCredentials        AuthenticationMethodType = "issuer_credentials"
	NoAuthentication         AuthenticationMethodType = "no_authentication"
	OwnCredentials           AuthenticationMethodType = "own_credentials"
	ThirdPartyAuthentication AuthenticationMethodType = "third_party_authentication"
)

type AccountTypeCardProductType string

const (
	Credit        AccountTypeCardProductType = "credit"
	Debit         AccountTypeCardProductType = "debit"
	NotApplicable AccountTypeCardProductType = "not_applicable"
)

type ServiceType string

const (
	SameDayST  ServiceType = "same_day"
	StandardST ServiceType = "standard"
)

type AccountUpdateStatusType string

const (
	CardUpdatedAUST       AccountUpdateStatusType = "card_updated"
	CardExpiryUpdatedAUST AccountUpdateStatusType = "card_expiry_updated"
	CardClosedAUST        AccountUpdateStatusType = "card_closed"
	ContactHolderAUST     AccountUpdateStatusType = "contact_holder"
)

type DeliveryTimeframe string

const (
	ElectronicDelivery DeliveryTimeframe = "electronic_delivery"
	SameDay            DeliveryTimeframe = "same_day"
	Overnight          DeliveryTimeframe = "overnight"
	TwoDayOrMore       DeliveryTimeframe = "two_day_or_more"
)

type (
	AirlineData struct {
		Ticket           *Ticket            `json:"ticket,omitempty"`
		Passenger        *Passenger         `json:"passenger,omitempty"`
		FlightLegDetails []FlightLegDetails `json:"flight_leg_details,omitempty"`
	}

	Ticket struct {
		Number                 string `json:"number,omitempty"`
		IssueDate              string `json:"issue_date,omitempty"`
		IssuingCarrierCode     string `json:"issuing_carrier_code,omitempty"`
		TravelPackageIndicator string `json:"travel_package_indicator,omitempty"`
		TravelAgencyName       string `json:"travel_agency_name,omitempty"`
		TravelAgencyCode       string `json:"travel_agency_code,omitempty"`
	}

	Passenger struct {
		FirstName   string          `json:"first_name,omitempty"`
		LastName    string          `json:"last_name,omitempty"`
		DateOfBirth string          `json:"date_of_birth,omitempty"`
		Address     *common.Address `json:"address,omitempty"`
	}

	FlightLegDetails struct {
		FlightNumber     string `json:"flight_number,omitempty"`
		CarrierCode      string `json:"carrier_code,omitempty"`
		ClassOfTraveling string `json:"class_of_traveling,omitempty"`
		DepartureAirport string `json:"departure_airport,omitempty"`
		DepartureDate    string `json:"departure_date,omitempty"`
		DepartureTime    string `json:"departure_time,omitempty"`
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

	ShippingDetailsFlowHostedLinks struct {
		Address *common.Address `json:"address,omitempty"`
		Phone   *common.Phone   `json:"phone,omitempty"`
	}

	ShippingDetails struct {
		FirstName      string                        `json:"first_name,omitempty"`
		LastName       string                        `json:"last_name,omitempty"`
		Email          string                        `json:"email,omitempty"`
		Address        *common.Address               `json:"address,omitempty"`
		Phone          *common.Phone                 `json:"phone,omitempty"`
		FromAddressZip string                        `json:"from_address_zip,omitempty"`
		Timeframe      DeliveryTimeframe             `json:"time_frame,omitempty"`
		Method         PaymentContextsShippingMethod `json:"method,omitempty"`
		Delay          int                           `json:"delay"`
	}

	PaymentSegment struct {
		Brand            string `json:"brand,omitempty"`
		BusinessCategory string `json:"business_category,omitempty"`
		Market           string `json:"market,omitempty"`
	}

	Dunning struct {
		Enabled      bool `json:"enabled,omitempty"`
		MaxAttempts  int  `json:"max_attempts,omitempty" default:"6"`
		EndAfterDays int  `json:"end_after_days,omitempty" default:"30"`
	}

	Downtime struct {
		Enabled bool `json:"enabled,omitempty"`
	}

	PaymentRetryRequest struct {
		Dunning  *Dunning  `json:"dunning,omitempty"`
		Downtime *Downtime `json:"downtime,omitempty"`

		Enabled      bool `json:"enabled,omitempty"`
		MaxAttempts  int  `json:"max_attempts,omitempty"`
		EndAfterDays int  `json:"end_after_days,omitempty"`
	}

	LocalDescriptor struct {
		Name         string           `json:"name,omitempty"`
		CharacterSet CharacterSetType `json:"character_set,omitempty"`
	}

	BillingDescriptor struct {
		Name            string            `json:"name,omitempty"`
		City            string            `json:"city,omitempty"`
		Reference       string            `json:"reference,omitempty"`
		LocalDescriptor []LocalDescriptor `json:"local_descriptor,omitempty"`
	}

	MerchantAuthenticationInfo struct {
		ThreeDsReqAuthMethod    ThreeDsReqAuthMethodType `json:"three_ds_req_auth_method,omitempty"`
		ThreeDSReqAuthTimestamp *time.Time               `json:"three_ds_req_auth_timestamp,omitempty"`
		ThreeDsReqAuthData      string                   `json:"three_ds_req_auth_data,omitempty"`
	}

	AccountInfo struct {
		PurchaseCount                  int64                                     `json:"purchase_count,omitempty"`
		AccountAge                     string                                    `json:"account_age,omitempty"`
		AddCardAttempts                int64                                     `json:"add_card_attempts,omitempty"`
		ShippingAddressAge             string                                    `json:"shipping_address_age,omitempty"`
		AccountNameMatchesShippingName bool                                      `json:"account_name_matches_shipping_name,omitempty"`
		SuspiciousAccountActivity      bool                                      `json:"suspicious_account_activity,omitempty"`
		TransactionsToday              int64                                     `json:"transactions_today,omitempty"`
		AuthenticationMethod           AuthenticationMethodType                  `json:"authentication_method,omitempty"` // Deprecated field
		CardholderAccountAgeIndicator  common.CardholderAccountAgeIndicatorType  `json:"cardholder_account_age_indicator,omitempty"`
		AccountChange                  *time.Time                                `json:"account_change,omitempty"`
		AccountChangeIndicator         common.AccountChangeIndicatorType         `json:"account_change_indicator,omitempty"`
		AccountDate                    *time.Time                                `json:"account_date,omitempty"`
		AccountPasswordChange          string                                    `json:"account_password_change,omitempty"`
		AccountPasswordChangeIndicator common.AccountPasswordChangeIndicatorType `json:"account_password_change_indicator,omitempty"`
		TransactionsPerYear            int                                       `json:"transactions_per_year,omitempty"`
		PaymentAccountAge              *time.Time                                `json:"payment_account_age,omitempty"`
		ShippingAddressUsage           *time.Time                                `json:"shipping_address_usage,omitempty"`
		AccountType                    AccountTypeCardProductType                `json:"account_type,omitempty"`
		AccountId                      string                                    `json:"account_id,omitempty"`
	}

	ThreeDsRequestFlowHostedLinks struct {
		Enabled            bool                      `json:"enabled"`
		AttemptN3D         bool                      `json:"attempt_n3d"`
		ChallengeIndicator common.ChallengeIndicator `json:"challenge_indicator,omitempty" default:"NoPreference"`
		Exemption          Exemption                 `json:"exemption,omitempty"`
		AllowUpgrade       bool                      `json:"allow_upgrade,omitempty"`
	}

	ThreeDsRequest struct {
		Enabled                    bool                        `json:"enabled"`
		AttemptN3D                 bool                        `json:"attempt_n3d"`
		Eci                        string                      `json:"eci,omitempty"`
		Cryptogram                 string                      `json:"cryptogram,omitempty"`
		Xid                        string                      `json:"xid,omitempty"`
		Version                    string                      `json:"version,omitempty"`
		Exemption                  Exemption                   `json:"exemption,omitempty"`
		ChallengeIndicator         common.ChallengeIndicator   `json:"challenge_indicator,omitempty" default:"NoPreference"`
		AllowUpgrade               bool                        `json:"allow_upgrade,omitempty"`
		Status                     string                      `json:"status,omitempty"`
		AuthenticationDate         *time.Time                  `json:"authentication_date,omitempty"`
		AuthenticationAmount       float64                     `json:"authentication_amount"`
		FlowType                   common.ThreeDsFlowType      `json:"flow_type,omitempty"`
		StatusReasonCode           string                      `json:"status_reason_code,omitempty"`
		ChallengeCancelReason      string                      `json:"challenge_cancel_reason,omitempty"`
		Score                      string                      `json:"score,omitempty"`
		CryptogramAlgorithm        string                      `json:"cryptogram_algorithm,omitempty"`
		AuthenticationId           string                      `json:"authentication_id,omitempty"`
		MerchantAuthenticationInfo *MerchantAuthenticationInfo `json:"merchant_authentication_info,omitempty"`
		AccountInfo                *AccountInfo                `json:"account_info,omitempty"`
	}

	Network struct {
		Ipv4  string `json:"ipv4,omitempty"`
		Ipv6  string `json:"ipv6,omitempty"`
		Tor   bool   `json:"tor,omitempty"`
		Vpn   bool   `json:"vpn,omitempty"`
		Proxy bool   `json:"proxy,omitempty"`
	}

	Provider struct {
		Id   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	}

	DeviceDetails struct {
		UserAgent         string    `json:"user_agent,omitempty"`
		Network           *Network  `json:"network,omitempty"`
		Provider          *Provider `json:"provider,omitempty"`
		Timestamp         string    `json:"timestamp,omitempty"`
		Timezone          string    `json:"timezone,omitempty"`
		VirtualMachine    bool      `json:"virtual_machine,omitempty"`
		Incognito         bool      `json:"incognito,omitempty"`
		Jailbroken        bool      `json:"jailbroken,omitempty"`
		Rooted            bool      `json:"rooted,omitempty"`
		JavaEnabled       bool      `json:"java_enabled,omitempty"`
		JavascriptEnabled bool      `json:"javascript_enabled,omitempty"`
		Language          string    `json:"language,omitempty"`
		ColorDepth        string    `json:"color_depth,omitempty"`
		ScreenHeight      string    `json:"screen_height,omitempty"`
		ScreenWidth       string    `json:"screen_width,omitempty"`
	}

	RiskRequest struct {
		Enabled         bool           `json:"enabled" default:"true"`
		DeviceSessionId string         `json:"device_session_id,omitempty"`
		Device          *DeviceDetails `json:"device,omitempty"`
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

	Guest struct {
		FirstName   string     `json:"first_name,omitempty"`
		LastName    string     `json:"last_name,omitempty"`
		DateOfBirth *time.Time `json:"date_of_birth,omitempty"`
	}

	Room struct {
		Rate                     string `json:"rate,omitempty"`
		NumberOfNightsAtRoomRate string `json:"number_of_nights_at_room_rate,omitempty"`
	}

	AccommodationData struct {
		Name             string          `json:"name,omitempty"`
		BookingReference string          `json:"booking_reference,omitempty"`
		CheckInDate      *time.Time      `json:"check_in_date,omitempty"`
		CheckOutDate     *time.Time      `json:"check_out_date,omitempty"`
		Address          *common.Address `json:"address,omitempty"`
		State            string          `json:"state,omitempty"`
		Country          common.Country  `json:"country,omitempty"`
		City             string          `json:"city,omitempty"`
		NumberOfRooms    int             `json:"number_of_rooms,omitempty"`
		Guests           []Guest         `json:"guests,omitempty"`
		Room             []Room          `json:"room,omitempty"`
	}
	PartnerCustomerRiskData struct {
		Key   string `json:"key,omitempty"`
		Value string `json:"value,omitempty"`
	}

	ProcessingSettings struct {
		OrderId                 string                    `json:"order_id,omitempty"`
		TaxAmount               int64                     `json:"tax_amount"`
		SurchargeAmount         int64                     `json:"surcharge_amount,omitempty"`
		DiscountAmount          int64                     `json:"discount_amount"`
		DutyAmount              int64                     `json:"duty_amount"`
		ShippingAmount          int64                     `json:"shipping_amount"`
		ShippingTaxAmount       int64                     `json:"shipping_tax_amount"`
		Aft                     bool                      `json:"aft,omitempty"`
		PreferredScheme         PreferredSchema           `json:"preferred_scheme,omitempty"`
		MerchantInitiatedReason MerchantInitiatedReason   `json:"merchant_initiated_reason,omitempty"`
		CampaignId              int64                     `json:"campaign_id,omitempty"`
		ProductType             ProductType               `json:"product_type,omitempty"`
		OpenId                  string                    `json:"open_id,omitempty"`
		OriginalOrderAmount     int64                     `json:"original_order_amount"`
		ReceiptId               string                    `json:"receipt_id,omitempty"`
		TerminalType            TerminalType              `json:"terminal_type,omitempty" default:"WEB"`
		OsType                  OsType                    `json:"os_type,omitempty"`
		InvoiceId               string                    `json:"invoice_id,omitempty"`
		BrandName               string                    `json:"brand_name,omitempty"`
		Locale                  string                    `json:"locale,omitempty"`
		ShippingPreference      ShippingPreference        `json:"shipping_preference,omitempty"`
		UserAction              UserAction                `json:"user_action,omitempty"`
		SetTransactionContext   []map[string]string       `json:"set_transaction_context,omitempty"`
		AirlineData             []AirlineData             `json:"airline_data,omitempty"`
		AccommodationData       []AccommodationData       `json:"accommodation_data,omitempty"`
		OtpValue                string                    `json:"otp_value,omitempty"`
		PurchaseCountry         common.Country            `json:"purchase_country,omitempty"`
		CustomPaymentMethodIds  []string                  `json:"custom_payment_method_ids,omitempty"`
		MerchantCallbackUrl     string                    `json:"merchant_callback_url,omitempty"`
		ShippingDelay           int64                     `json:"shipping_delay,omitempty"`
		ShippingInfo            string                    `json:"shipping_info,omitempty"`
		LineOfBusiness          string                    `json:"line_of_business,omitempty"`
		PanPreference           PanProcessedType          `json:"pan_preference,omitempty"`
		ServiceType             ServiceType               `json:"service_type,omitempty"`
		ProvisionNetworkToken   bool                      `json:"provision_network_token,omitempty" default:"ture"`
		SenderInformation       *SenderInformation        `json:"senderInformation,omitempty"`
		Purpose                 string                    `json:"purpose,omitempty"`
		Dlocal                  *DLocalProcessingSettings `json:"dlocal,omitempty"`
		PartnerCustomerRiskData *PartnerCustomerRiskData  `json:"partner_customer_risk_data,omitempty"`
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
		PartnerPaymentId                 string           `json:"partner_payment_id,omitempty"`
		PartnerStatus                    string           `json:"partner_status,omitempty"`
		PartnerTransactionId             string           `json:"partner_transaction_id,omitempty"`
		PartnerSessionId                 string           `json:"partner_session_id,omitempty"`
		PartnerErrorCodes                []string         `json:"partner_error_codes,omitempty"`
		PartnerErrorMessage              string           `json:"partner_error_message,omitempty"`
		PartnerAuthorizationCode         string           `json:"partner_authorization_code,omitempty"`
		PartnerAuthorizationResponseCode string           `json:"partner_authorization_response_code,omitempty"`
		SurchargeAmount                  int64            `json:"surcharge_amount,omitempty"`
		PanTypeProcessed                 PanProcessedType `json:"pan_type_processed,omitempty"`
		CkoNetworkTokenAvailable         bool             `json:"cko_network_token_available"`
		PartnerClientToken               string           `json:"partner_client_token,omitempty"`
		ContinuationPayload              string           `json:"continuation_payload,omitempty"`
		Pun                              string           `json:"pun,omitempty"`
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
		Type           string     `json:"type,omitempty"`
		Name           string     `json:"name,omitempty"`
		Quantity       int        `json:"quantity,omitempty"`
		UnitPrice      int        `json:"unit_price"`
		Price          int        `json:"price"`
		Reference      string     `json:"reference,omitempty"`
		CommodityCode  string     `json:"commodity_code,omitempty"`
		UnitOfMeasure  string     `json:"unit_of_measure,omitempty"`
		TotalAmount    int64      `json:"total_amount,omitempty"`
		TaxRate        int64      `json:"tax_rate,omitempty"`
		TaxAmount      int64      `json:"tax_amount,omitempty"`
		DiscountAmount int64      `json:"discount_amount,omitempty"`
		WxpayGoodsId   string     `json:"wxpay_goods_id,omitempty"`
		Url            string     `json:"url,omitempty"`
		ImageUrl       string     `json:"image_url,omitempty"`
		ServiceEndsOn  *time.Time `json:"service_ends_on,omitempty"`
		Sku            string     `json:"sku,omitempty"`
	}

	BillingInformation struct {
		Address *common.Address `json:"address,omitempty"`
		Phone   *common.Phone   `json:"phone,omitempty"`
	}

	RefundOrder struct {
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

	Applepay struct {
		AccountHolder       *common.AccountHolder   `json:"account_holder,omitempty"`
		StorePaymentDetails StorePaymentDetailsType `json:"store_payment_details,omitempty"`
	}

	Card struct {
		AccountHolder       *common.AccountHolder   `json:"account_holder,omitempty"`
		StorePaymentDetails StorePaymentDetailsType `json:"store_payment_details,omitempty"`
	}

	Googlepay struct {
		AccountHolder       *common.AccountHolder   `json:"account_holder,omitempty"`
		StorePaymentDetails StorePaymentDetailsType `json:"store_payment_details,omitempty"`
	}

	StoredCard struct {
		AccountHolder       *common.AccountHolder   `json:"account_holder,omitempty"`
		StorePaymentDetails StorePaymentDetailsType `json:"store_payment_details,omitempty"`
	}

	PaymentMethodConfiguration struct {
		Applepay   *Applepay   `json:"applepay,omitempty"`
		Card       *Card       `json:"card,omitempty"`
		Googlepay  *Googlepay  `json:"googlepay,omitempty"`
		StoredCard *StoredCard `json:"stored_card,omitempty"`
	}

	PaymentInstruction struct {
		Purpose PaymentPurposeType `json:"purpose,omitempty"`
	}

	PaymentPlan struct {
		//Recurring
		AmountVariability AmountVariabilityType `json:"amount_variability,omitempty"`

		//Installment
		Financing bool  `json:"financing,omitempty" default:"false"`
		Amount    int64 `json:"amount,omitempty"`

		//Common
		DaysBetweenPayments   int        `json:"days_between_payments,omitempty"`
		TotalNumberOfPayments int        `json:"total_number_of_payments,omitempty"`
		CurrentPaymentNumber  int        `json:"current_payment_number,omitempty"`
		Expiry                *time.Time `json:"expiry,omitempty"`
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
		Items             []RefundOrder              `json:"items,omitempty"`
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
