package apm

import (
	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/payments"
	"time"
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
type (
	requestAfterPaySource struct {
		Type          payments.SourceType   `json:"type,omitempty"`
		AccountHolder *common.AccountHolder `json:"account_holder,omitempty"`
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

	requestEpsSource struct {
		Type    payments.SourceType `json:"type,omitempty"`
		Purpose string              `json:"purpose,omitempty"`
	}

	requestFawrySource struct {
		Type              payments.SourceType `json:"type,omitempty"`
		Description       string              `json:"description,omitempty"`
		CustomerProfileId string              `json:"customer_profile_id,omitempty"`
		CustomerEmail     string              `json:"customer_email,omitempty"`
		CustomerMobile    string              `json:"customer_mobile,omitempty"`
		ExpiresOn         time.Time           `json:"expires_on,omitempty"`
		Products          []FawryProduct      `json:"Products,omitempty"`
	}

	requestGiropaySource struct {
		Type       payments.SourceType `json:"type,omitempty"`
		Purpose    string              `json:"purpose,omitempty"`
		InfoFields []InfoFields        `json:"info_fields,omitempty"`
	}

	requestIdealSource struct {
		Type        payments.SourceType `json:"type,omitempty"`
		Description string              `json:"description,omitempty"`
		Bic         string              `json:"bic,omitempty"`
		Language    string              `json:"language,omitempty"`
	}

	requestKlarnaSource struct {
		Type          payments.SourceType   `json:"type,omitempty"`
		AccountHolder *common.AccountHolder `json:"account_holder,omitempty"`
	}

	requestKnetSource struct {
		Type              payments.SourceType `json:"type,omitempty"`
		Language          string              `json:"language,omitempty"`
		UserDefinedField1 string              `json:"user_defined_field1,omitempty"`
		UserDefinedField2 string              `json:"user_defined_field2,omitempty"`
		UserDefinedField3 string              `json:"user_defined_field3,omitempty"`
		UserDefinedField4 string              `json:"user_defined_field4,omitempty"`
		UserDefinedField5 string              `json:"user_defined_field5,omitempty"`
		CardToken         string              `json:"card_token,omitempty"`
		Ptlf              string              `json:"ptlf,omitempty"`
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

	requestPostFinanceSource struct {
		Type              payments.SourceType `json:"type,omitempty"`
		PaymentCountry    common.Country      `json:"payment_country,omitempty"`
		AccountHolderName string              `json:"account_holder_name,omitempty"`
		BillingDescriptor string              `json:"billing_descriptor,omitempty"`
	}

	requestQPaySource struct {
		Type        payments.SourceType `json:"type,omitempty"`
		Quantity    int                 `json:"quantity,omitempty"`
		Description string              `json:"description,omitempty"`
		Language    string              `json:"language,omitempty"`
		NationalId  string              `json:"national_id,omitempty"`
	}

	requestSofortSource struct {
		Type         payments.SourceType `json:"type,omitempty"`
		CountryCode  common.Country      `json:"countryCode,omitempty"`
		LanguageCode string              `json:"languageCode,omitempty"`
	}

	requestStcPaySource struct {
		Type payments.SourceType `json:"type,omitempty"`
	}

	requestTamaraSource struct {
		Type           payments.SourceType `json:"type,omitempty"`
		BillingAddress *common.Address     `json:"billing_address,omitempty"`
	}

	requestWeChatPaySource struct {
		Type           payments.SourceType `json:"type,omitempty"`
		BillingAddress *common.Address     `json:"billing_address,omitempty"`
	}
)

func NewRequestAfterPaySource() *requestAfterPaySource {
	return &requestAfterPaySource{Type: payments.Afterpay}
}

func NewRequestAlipayPlusSource() *requestAlipayPlusSource {
	return &requestAlipayPlusSource{Type: payments.AlipayPlus}
}

func NewRequestAlipayPlusCNSource() *requestAlipayPlusSource {
	return &requestAlipayPlusSource{Type: payments.AlipayCn}
}

func NewRequestAlipayPlusGCashSource() *requestAlipayPlusSource {
	return &requestAlipayPlusSource{Type: payments.Gcash}
}

func NewRequestAlipayPlusHKSource() *requestAlipayPlusSource {
	return &requestAlipayPlusSource{Type: payments.AlipayHk}
}

func NewRequestAlipayPlusDanaSource() *requestAlipayPlusSource {
	return &requestAlipayPlusSource{Type: payments.Dana}
}

func NewRequestAlipayPlusKakaoPaySource() *requestAlipayPlusSource {
	return &requestAlipayPlusSource{Type: payments.Kakaopay}
}

func NewRequestAlipayPlusTrueMoneySource() *requestAlipayPlusSource {
	return &requestAlipayPlusSource{Type: payments.Truemoney}
}

func NewRequestAlipayPlusTNGSource() *requestAlipayPlusSource {
	return &requestAlipayPlusSource{Type: payments.Tng}
}

func NewRequestAlmaSource() *requestAlmaSource {
	return &requestAlmaSource{Type: payments.Alma}
}

func NewRequestBancontactSource() *requestBancontactSource {
	return &requestBancontactSource{Type: payments.BancontactSource}
}

func NewRequestBenefitSource() *requestBenefitSource {
	return &requestBenefitSource{Type: payments.Benefit}
}

func NewRequestEpsSource() *requestEpsSource {
	return &requestEpsSource{Type: payments.EpsSource}
}

func NewRequestFawrySource() *requestFawrySource {
	return &requestFawrySource{Type: payments.FawrySource}
}

func NewRequestGiropaySource() *requestGiropaySource {
	return &requestGiropaySource{Type: payments.GiropaySource}
}

func NewRequestIdealSource() *requestIdealSource {
	return &requestIdealSource{Type: payments.IdealSource}
}

func NewRequestKlarnaSource() *requestKlarnaSource {
	return &requestKlarnaSource{Type: payments.KlarnaSource}
}

func NewRequestKnetSource() *requestKnetSource {
	return &requestKnetSource{Type: payments.KnetSource}
}

func NewRequestMbwaySource() *requestMbwaySource {
	return &requestMbwaySource{Type: payments.Mbway}
}

func NewRequestMultiBancoSource() *requestMultiBancoSource {
	return &requestMultiBancoSource{Type: payments.MultiBancoSource}
}

func NewRequestP24Source() *requestP24Source {
	return &requestP24Source{Type: payments.P24Source}
}

func NewRequestPayPalSource() *requestPayPalSource {
	return &requestPayPalSource{Type: payments.PayPalSource}
}

func NewRequestPostFinanceSource() *requestPostFinanceSource {
	return &requestPostFinanceSource{Type: payments.Postfinance}
}

func NewRequestQPaySource() *requestQPaySource {
	return &requestQPaySource{Type: payments.QPaySource}
}

func NewRequestSofortSource() *requestSofortSource {
	return &requestSofortSource{Type: payments.SofortSource}
}

func NewRequestStcPaySource() *requestStcPaySource {
	return &requestStcPaySource{Type: payments.SofortSource}
}

func NewRequestTamaraSource() *requestTamaraSource {
	return &requestTamaraSource{Type: payments.TamaraSource}
}

func NewRequestWeChatPaySource() *requestWeChatPaySource {
	return &requestWeChatPaySource{Type: payments.Wechatpay}
}

func (s *requestAfterPaySource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestAlipayPlusSource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestAlmaSource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestBancontactSource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestBenefitSource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestEpsSource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestFawrySource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestGiropaySource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestIdealSource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestKlarnaSource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestKnetSource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestMbwaySource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestMultiBancoSource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestP24Source) GetType() payments.SourceType {
	return s.Type
}

func (s *requestPayPalSource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestPostFinanceSource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestQPaySource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestSofortSource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestStcPaySource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestTamaraSource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestWeChatPaySource) GetType() payments.SourceType {
	return s.Type
}
