package apm

import (
	"time"

	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/payments"
	"github.com/checkout/checkout-sdk-go/v3/tokens"
)

type BillingPlanType string

const (
	MerchantInitiatedBilling                BillingPlanType = "MERCHANT_INITIATED_BILLING"
	MerchantInitiatedBillingSingleAgreement BillingPlanType = "MERCHANT_INITIATED_BILLING_SINGLE_AGREEMENT"
	ChannelInitiatedBilling                 BillingPlanType = "CHANNEL_INITIATED_BILLING"
	ChannelInitiatedBillingSingleAgreement  BillingPlanType = "CHANNEL_INITIATED_BILLING_SINGLE_AGREEMENT"
	RecurringPayments                       BillingPlanType = "RECURRING_PAYMENTS"
	PreApprovedPayments                     BillingPlanType = "PRE_APPROVED_PAYMENTS"
)

// Properties
type (
	FawryProduct struct {
		ProductId   string  `json:"product_id,omitempty"`
		Quantity    float64 `json:"quantity,omitempty"`
		Price       float64 `json:"price,omitempty"`
		Description string  `json:"description,omitempty"`
	}

	InfoFields struct {
		Label string `json:"label,omitempty"`
		Text  string `json:"text,omitempty"`
	}

	BillingPlan struct {
		Type                     BillingPlanType `json:"type,omitempty"`
		SkipShippingAddress      bool            `json:"skip_shipping_address,omitempty"`
		ImmutableShippingAddress bool            `json:"immutable_shipping_address,omitempty"`
	}
)

// Requests
// SepaMandateType is the type of SEPA mandate on a SEPA payment source.
//
// This deliberately does not reuse instruments/nas.SepaMandateType. The two carry the same two
// values today, but they belong to independent schemas: this one to
// PaymentRequestSEPAV4Source.mandate_type, the other to the SEPA instrument's instrument_data.type.
type SepaMandateType string

const (
	SepaMandateCore SepaMandateType = "Core"
	SepaMandateB2B  SepaMandateType = "B2B"
)

// SepaSourceAccountHolderType is the type of account holder on a SEPA payment source.
//
// This position declares two values only, unlike common.AccountHolderType which also declares
// "government". Send them lowercase: the specification declares them capitalized here, but every
// other account-holder-type position declares them lowercase and every other Checkout.com SDK sends
// lowercase. Pending confirmation from the API owners.
type SepaSourceAccountHolderType string

const (
	SepaSourceIndividual SepaSourceAccountHolderType = "individual"
	SepaSourceCorporate  SepaSourceAccountHolderType = "corporate"
)

// AchSourceAccountType is the type of Direct Debit account on an ACH payment source.
//
// PaymentRequestAchSource is the only schema declaring this set. common.AccountType is
// savings / current / cash and serves the bank-account positions, so it cannot express "checking".
// nas.AchAccountType is savings / checking and serves the stored ACH instrument positions, so it
// does not declare "cash".
type AchSourceAccountType string

const (
	AchSourceSavings  AchSourceAccountType = "savings"
	AchSourceChecking AchSourceAccountType = "checking"
	AchSourceCash     AchSourceAccountType = "cash"
)

type (
	// SepaSourceBillingAddress is the account holder's billing address on a SEPA payment source.
	//
	// Every property is required. Deliberately not common.Address, which also declares a State that
	// this position does not accept.
	SepaSourceBillingAddress struct {
		AddressLine1 string         `json:"address_line1,omitempty"`
		AddressLine2 string         `json:"address_line2,omitempty"`
		City         string         `json:"city,omitempty"`
		Zip          string         `json:"zip,omitempty"`
		Country      common.Country `json:"country,omitempty"`
	}

	// SepaSourceAccountHolder is the account holder's personal information on a SEPA payment source.
	//
	// Maps the account_holder object of PaymentRequestSEPAV4Source. Deliberately not
	// common.AccountHolder, which is an 18-property superset. Only BillingAddress is required here,
	// where the SEPA instrument requires the names too. Send Type lowercase (individual, corporate):
	// the specification declares it capitalized at this one position, but every other
	// account-holder-type position declares it lowercase and every other Checkout.com SDK sends
	// lowercase. Pending confirmation from the API owners.
	//
	// FirstName, LastName and CompanyName are each max 50 characters.
	SepaSourceAccountHolder struct {
		BillingAddress *SepaSourceBillingAddress   `json:"billing_address,omitempty"`
		FirstName      string                      `json:"first_name,omitempty"`
		LastName       string                      `json:"last_name,omitempty"`
		CompanyName    string                      `json:"company_name,omitempty"`
		Type           SepaSourceAccountHolderType `json:"type,omitempty"`
	}

	// AchSourceAccountHolder is the account holder's details on an ACH payment source.
	//
	// Maps the AccountHolderAch schema exactly. Deliberately not common.AccountHolder, which is an
	// 18-property superset, and distinct from the four-property ACH instrument account holder - the
	// instrument schema declares no billing address, date of birth or identification.
	//
	// Type, FirstName and LastName are required. BillingAddress reuses common.Address because that
	// schema's six properties are exactly what this position references. Identification reuses
	// common.AccountHolderIdentification, which carries one extra property, DateOfExpiry, that this
	// position does not declare - do not set it.
	AchSourceAccountHolder struct {
		Type           common.AccountHolderType            `json:"type,omitempty"`
		FirstName      string                              `json:"first_name,omitempty"`
		LastName       string                              `json:"last_name,omitempty"`
		CompanyName    string                              `json:"company_name,omitempty"`
		BillingAddress *common.Address                     `json:"billing_address,omitempty"`
		DateOfBirth    string                              `json:"date_of_birth,omitempty"`
		Identification *common.AccountHolderIdentification `json:"identification,omitempty"`
	}

	requestAchSource struct {
		Type payments.SourceType `json:"type,omitempty"`
		// AccountType is savings, checking or cash. Deliberately not common.AccountType, which
		// declares "current" instead of "checking" and is rejected at this position.
		AccountType   AchSourceAccountType    `json:"account_type,omitempty"`
		AccountHolder *AchSourceAccountHolder `json:"account_holder,omitempty"`
		AccountNumber string                  `json:"account_number,omitempty"`
		BankCode      string                  `json:"bank_code,omitempty"`
		Country       common.Country          `json:"country,omitempty"`
	}

	requestAfterPaySource struct {
		Type          payments.SourceType   `json:"type,omitempty"`
		AccountHolder *common.AccountHolder `json:"account_holder,omitempty"`
	}

	requestAlipayCnSource struct {
		Type payments.SourceType `json:"type,omitempty"`
	}

	requestAlipayHkSource struct {
		Type payments.SourceType `json:"type,omitempty"`
	}

	requestAlipayPlusSource struct {
		Type payments.SourceType `json:"type,omitempty"`
	}

	requestAlmaSource struct {
		Type           payments.SourceType `json:"type,omitempty"`
		BillingAddress *common.Address     `json:"billing_address,omitempty"`
	}

	requestBancontactSource struct {
		Type              payments.SourceType `json:"type,omitempty"`
		PaymentCountry    common.Country      `json:"payment_country,omitempty"`
		AccountHolderName string              `json:"account_holder_name,omitempty"`
		BillingDescriptor string              `json:"billing_descriptor,omitempty"`
		Language          string              `json:"language,omitempty"`
	}

	requestBenefitSource struct {
		Type payments.SourceType `json:"type,omitempty"`
	}

	requestBizumSource struct {
		Type payments.SourceType `json:"type,omitempty"`
	}

	requestCvConnectSource struct {
		Type           payments.SourceType `json:"type,omitempty"`
		BillingAddress *common.Address     `json:"billing_address,omitempty"`
	}

	requestDanaSource struct {
		Type payments.SourceType `json:"type,omitempty"`
	}

	requestEpsSource struct {
		Type          payments.SourceType   `json:"type,omitempty"`
		Purpose       string                `json:"purpose,omitempty"`
		AccountHolder *common.AccountHolder `json:"account_holder,omitempty"`
	}

	requestFawrySource struct {
		Type              payments.SourceType `json:"type,omitempty"`
		Description       string              `json:"description,omitempty"`
		Products          []FawryProduct      `json:"Products,omitempty"`
		CustomerEmail     string              `json:"customer_email,omitempty"`
		CustomerMobile    string              `json:"customer_mobile,omitempty"`
		CustomerProfileId string              `json:"customer_profile_id,omitempty"`
		ExpiresOn         *time.Time          `json:"expires_on,omitempty"`
	}

	requestGcashSource struct {
		Type payments.SourceType `json:"type,omitempty"`
	}

	requestGiropaySource struct {
		Type          payments.SourceType   `json:"type,omitempty"`
		AccountHolder *common.AccountHolder `json:"account_holder,omitempty"`
	}

	requestIdealSource struct {
		Type        payments.SourceType `json:"type,omitempty"`
		Description string              `json:"description,omitempty"`
		Language    string              `json:"language,omitempty"`
	}

	requestIllicadoSource struct {
		Type           payments.SourceType `json:"type,omitempty"`
		BillingAddress *common.Address     `json:"billing_address,omitempty"`
	}

	requestKakaopaySource struct {
		Type payments.SourceType `json:"type,omitempty"`
	}

	requestKlarnaSource struct {
		Type          payments.SourceType   `json:"type,omitempty"`
		AccountHolder *common.AccountHolder `json:"account_holder,omitempty"`
	}

	requestKnetSource struct {
		Type                 payments.SourceType            `json:"type,omitempty"`
		Language             string                         `json:"language,omitempty"`
		UserDefinedField1    string                         `json:"user_defined_field1,omitempty"`
		UserDefinedField2    string                         `json:"user_defined_field2,omitempty"`
		UserDefinedField3    string                         `json:"user_defined_field3,omitempty"`
		UserDefinedField4    string                         `json:"user_defined_field4,omitempty"`
		UserDefinedField5    string                         `json:"user_defined_field5,omitempty"`
		CardToken            string                         `json:"card_token,omitempty"`
		Ptlf                 string                         `json:"ptlf,omitempty"`
		TokenType            string                         `json:"token_type,omitempty"`
		TokenData            *tokens.ApplePayTokenData      `json:"token_data,omitempty"`
		PaymentMethodDetails *payments.PaymentMethodDetails `json:"payment_method_details,omitempty"`
	}

	requestMbwaySource struct {
		Type payments.SourceType `json:"type,omitempty"`
	}

	requestMultiBancoSource struct {
		Type              payments.SourceType `json:"type,omitempty"`
		PaymentCountry    common.Country      `json:"payment_country,omitempty"`
		AccountHolderName string              `json:"account_holder_name,omitempty"`
		BillingDescriptor string              `json:"billing_descriptor,omitempty"`
	}

	requestOctopusSource struct {
		Type payments.SourceType `json:"type,omitempty"`
	}

	requestPaynowSource struct {
		Type payments.SourceType `json:"type,omitempty"`
	}

	requestPlaidSource struct {
		Type          payments.SourceType   `json:"type,omitempty"`
		Token         string                `json:"token,omitempty"`
		AccountHolder *common.AccountHolder `json:"account_holder,omitempty"`
	}

	requestPostFinanceSource struct {
		Type              payments.SourceType `json:"type,omitempty"`
		PaymentCountry    common.Country      `json:"payment_country,omitempty"`
		AccountHolderName string              `json:"account_holder_name,omitempty"`
		BillingDescriptor string              `json:"billing_descriptor,omitempty"`
	}

	requestP24Source struct {
		Type               payments.SourceType `json:"type,omitempty"`
		PaymentCountry     common.Country      `json:"payment_country,omitempty"`
		AccountHolderName  string              `json:"account_holder_name,omitempty"`
		AccountHolderEmail string              `json:"account_holder_email,omitempty"`
		BillingDescriptor  string              `json:"billing_descriptor,omitempty"`
	}

	requestPayPalSource struct {
		Type payments.SourceType `json:"type,omitempty"`
		Plan *BillingPlan        `json:"plan,omitempty"`
	}

	requestQPaySource struct {
		Type        payments.SourceType `json:"type,omitempty"`
		Description string              `json:"description,omitempty"`
		Quantity    int                 `json:"quantity,omitempty"`
		Language    string              `json:"language,omitempty"`
		NationalId  string              `json:"national_id,omitempty"`
	}

	requestSofortSource struct {
		Type         payments.SourceType `json:"type,omitempty"`
		CountryCode  common.Country      `json:"countryCode,omitempty"`
		LanguageCode string              `json:"languageCode,omitempty"`
	}

	requestSequraSource struct {
		Type           payments.SourceType `json:"type,omitempty"`
		BillingAddress *common.Address     `json:"billing_address,omitempty"`
	}

	// requestSepaSource is the SEPA Direct Debit source.
	requestSepaSource struct {
		Type          payments.SourceType `json:"type,omitempty"`
		Country       common.Country      `json:"country,omitempty"`
		AccountNumber string              `json:"account_number,omitempty"`
		// BankCode is not declared by PaymentRequestSEPAV4Source. No SEPA schema in the
		// specification declares a bank code, and the SEPA source is identified by IBAN through
		// AccountNumber. Retained for retro-compatibility purposes only. Possibly an obsoleted
		// field.
		BankCode        string                   `json:"bank_code,omitempty"`
		Currency        common.Currency          `json:"currency,omitempty"`
		AccountHolder   *SepaSourceAccountHolder `json:"account_holder,omitempty"`
		MandateId       string                   `json:"mandate_id,omitempty"`
		MandateType     SepaMandateType          `json:"mandate_type,omitempty"`
		DateOfSignature string                   `json:"date_of_signature,omitempty"`
	}

	// requestBacsSource is the Bacs Direct Debit source.
	//
	// Id is the Bacs Direct Debit instrument ID and matches the pattern ^(src)_(\w{26})$.
	requestBacsSource struct {
		Type payments.SourceType `json:"type,omitempty"`
		Id   string              `json:"id,omitempty"`
	}

	requestStcPaySource struct {
		Type payments.SourceType `json:"type,omitempty"`
	}

	requestTamaraSource struct {
		Type           payments.SourceType `json:"type,omitempty"`
		BillingAddress *common.Address     `json:"billing_address,omitempty"`
	}

	requestTngSource struct {
		Type payments.SourceType `json:"type,omitempty"`
	}

	requestTruemoneySource struct {
		Type payments.SourceType `json:"type,omitempty"`
	}

	requestTwintSource struct {
		Type payments.SourceType `json:"type,omitempty"`
	}

	requestTrustlySource struct {
		Type           payments.SourceType `json:"type,omitempty"`
		BillingAddress *common.Address     `json:"billing_address,omitempty"`
	}

	requestWeChatPaySource struct {
		Type           payments.SourceType `json:"type,omitempty"`
		BillingAddress *common.Address     `json:"billing_address,omitempty"`
	}

	RequestBlikSource struct {
		Type               payments.SourceType `json:"type,omitempty"`
		PartnerAgreementId string              `json:"partner_agreement_id,omitempty"`
	}

	PaymentGetResponseBlikSource struct {
		Type payments.SourceType `json:"type,omitempty"`
		Id   string              `json:"id,omitempty"`
	}
)

func NewRequestAchSource() *requestAchSource {
	return &requestAchSource{Type: payments.AchSource}
}

func (s *requestAchSource) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestAfterPaySource() *requestAfterPaySource {
	return &requestAfterPaySource{Type: payments.Afterpay}
}

func (s *requestAfterPaySource) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestAlipayCnSource() *requestAlipayCnSource {
	return &requestAlipayCnSource{Type: payments.AlipayCn}
}

func (s *requestAlipayCnSource) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestAlipayHkSource() *requestAlipayHkSource {
	return &requestAlipayHkSource{Type: payments.AlipayHk}
}

func (s *requestAlipayHkSource) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestAlipayPlusSource() *requestAlipayPlusSource {
	return &requestAlipayPlusSource{Type: payments.AlipayPlus}
}

func (s *requestAlipayPlusSource) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestAlmaSource() *requestAlmaSource {
	return &requestAlmaSource{Type: payments.Alma}
}

func (s *requestAlmaSource) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestBancontactSource() *requestBancontactSource {
	return &requestBancontactSource{Type: payments.BancontactSource}
}

func (s *requestBancontactSource) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestBenefitSource() *requestBenefitSource {
	return &requestBenefitSource{Type: payments.Benefit}
}

func (s *requestBenefitSource) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestBizumSource() *requestBizumSource {
	return &requestBizumSource{Type: payments.BizumSource}
}

func (s *requestBizumSource) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestCvConnectSource() *requestCvConnectSource {
	return &requestCvConnectSource{Type: payments.CvConnectSource}
}

func (s *requestCvConnectSource) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestDanaSource() *requestDanaSource {
	return &requestDanaSource{Type: payments.Dana}
}

func (s *requestDanaSource) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestEpsSource() *requestEpsSource {
	return &requestEpsSource{Type: payments.EpsSource}
}

func (s *requestEpsSource) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestFawrySource() *requestFawrySource {
	return &requestFawrySource{Type: payments.FawrySource}
}

func (s *requestFawrySource) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestGcashSource() *requestGcashSource {
	return &requestGcashSource{Type: payments.Gcash}
}

func (s *requestGcashSource) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestGiropaySource() *requestGiropaySource {
	return &requestGiropaySource{Type: payments.GiropaySource}
}

func (s *requestGiropaySource) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestIdealSource() *requestIdealSource {
	return &requestIdealSource{Type: payments.IdealSource}
}

func (s *requestIdealSource) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestIllicadoSource() *requestIllicadoSource {
	return &requestIllicadoSource{Type: payments.IllicadoSource}
}

func (s *requestIllicadoSource) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestKakaopaySource() *requestKakaopaySource {
	return &requestKakaopaySource{Type: payments.Kakaopay}
}

func (s *requestKakaopaySource) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestKlarnaSource() *requestKlarnaSource {
	return &requestKlarnaSource{Type: payments.KlarnaSource}
}

func (s *requestKlarnaSource) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestKnetSource() *requestKnetSource {
	return &requestKnetSource{Type: payments.KnetSource}
}

func (s *requestKnetSource) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestMbwaySource() *requestMbwaySource {
	return &requestMbwaySource{Type: payments.Mbway}
}

func (s *requestMbwaySource) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestMultiBancoSource() *requestMultiBancoSource {
	return &requestMultiBancoSource{Type: payments.MultiBancoSource}
}

func (s *requestMultiBancoSource) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestOctopusSource() *requestOctopusSource {
	return &requestOctopusSource{Type: payments.OctopusSource}
}

func (s *requestOctopusSource) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestPaynowSource() *requestPaynowSource {
	return &requestPaynowSource{Type: payments.PaynowSource}
}

func (s *requestPaynowSource) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestPlaidSource() *requestPlaidSource {
	return &requestPlaidSource{Type: payments.PlaidSource}
}

func (s *requestPlaidSource) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestPostFinanceSource() *requestPostFinanceSource {
	return &requestPostFinanceSource{Type: payments.Postfinance}
}

func (s *requestPostFinanceSource) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestP24Source() *requestP24Source {
	return &requestP24Source{Type: payments.P24Source}
}

func (s *requestP24Source) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestPayPalSource() *requestPayPalSource {
	return &requestPayPalSource{Type: payments.PayPalSource}
}

func (s *requestPayPalSource) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestQPaySource() *requestQPaySource {
	return &requestQPaySource{Type: payments.QPaySource}
}

func (s *requestQPaySource) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestSofortSource() *requestSofortSource {
	return &requestSofortSource{Type: payments.SofortSource}
}

func (s *requestSofortSource) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestSequraSource() *requestSequraSource {
	return &requestSequraSource{Type: payments.SequraSource}
}

func (s *requestSequraSource) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestSepaSource() *requestSepaSource {
	return &requestSepaSource{Type: payments.SepaSource}
}

func (s *requestSepaSource) GetType() payments.SourceType {
	return s.Type
}

func NewRequestBacsSource() *requestBacsSource {
	return &requestBacsSource{Type: payments.BacsSource}
}

func (s *requestBacsSource) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestStcPaySource() *requestStcPaySource {
	return &requestStcPaySource{Type: payments.Stcpay}
}

func (s *requestStcPaySource) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestTamaraSource() *requestTamaraSource {
	return &requestTamaraSource{Type: payments.TamaraSource}
}

func (s *requestTamaraSource) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestTngSource() *requestTngSource {
	return &requestTngSource{Type: payments.Tng}
}

func (s *requestTngSource) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestTruemoneySource() *requestTruemoneySource {
	return &requestTruemoneySource{Type: payments.Truemoney}
}

func (s *requestTruemoneySource) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestTwintSource() *requestTwintSource {
	return &requestTwintSource{Type: payments.TwintSource}
}

func (s *requestTwintSource) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestTrustlySource() *requestTrustlySource {
	return &requestTrustlySource{Type: payments.TrustlySource}
}

func (s *requestTrustlySource) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestWeChatPaySource() *requestWeChatPaySource {
	return &requestWeChatPaySource{Type: payments.Wechatpay}
}

func (s *requestWeChatPaySource) GetType() payments.SourceType {
	return s.Type
}

//

func NewRequestBlikSource() *RequestBlikSource {
	return &RequestBlikSource{Type: payments.BlikSource}
}

func (s *RequestBlikSource) GetType() payments.SourceType {
	return s.Type
}
