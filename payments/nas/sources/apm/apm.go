package apm

import (
	"github.com/checkout/checkout-sdk-go-beta/common"
	"github.com/checkout/checkout-sdk-go-beta/payments"
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
	RequestAfterPaySource struct {
		Type          payments.SourceType   `json:"type,omitempty"`
		AccountHolder *common.AccountHolder `json:"account_holder,omitempty"`
	}

	RequestAlipayPlusSource struct {
		Type payments.SourceType `json:"type,omitempty"`
	}

	RequestAlmaSource struct {
		Type           payments.SourceType `json:"type,omitempty"`
		BillingAddress *common.Address     `json:"billing_address,omitempty"`
	}

	RequestBancontactSource struct {
		Type              payments.SourceType `json:"type,omitempty"`
		PaymentCountry    common.Country      `json:"payment_country,omitempty"`
		AccountHolderName string              `json:"account_holder_name,omitempty"`
		BillingDescriptor string              `json:"billing_descriptor,omitempty"`
		Language          string              `json:"language,omitempty"`
	}

	RequestBenefitSource struct {
		Type payments.SourceType `json:"type,omitempty"`
	}

	RequestEpsSource struct {
		Type    payments.SourceType `json:"type,omitempty"`
		Purpose string              `json:"purpose,omitempty"`
	}

	RequestFawrySource struct {
		Type              payments.SourceType `json:"type,omitempty"`
		Description       string              `json:"description,omitempty"`
		CustomerProfileId string              `json:"customer_profile_id,omitempty"`
		CustomerEmail     string              `json:"customer_email,omitempty"`
		CustomerMobile    string              `json:"customer_mobile,omitempty"`
		ExpiresOn         time.Time           `json:"expires_on,omitempty"`
		Products          []FawryProduct      `json:"Products,omitempty"`
	}

	RequestGiropaySource struct {
		Type       payments.SourceType `json:"type,omitempty"`
		Purpose    string              `json:"purpose,omitempty"`
		InfoFields []InfoFields        `json:"info_fields,omitempty"`
	}

	RequestIdealSource struct {
		Type        payments.SourceType `json:"type,omitempty"`
		Description string              `json:"description,omitempty"`
		Bic         string              `json:"bic,omitempty"`
		Language    string              `json:"language,omitempty"`
	}

	RequestKlarnaSource struct {
		Type          payments.SourceType   `json:"type,omitempty"`
		AccountHolder *common.AccountHolder `json:"account_holder,omitempty"`
	}

	RequestKnetSource struct {
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

	RequestMbwaySource struct {
		Type payments.SourceType `json:"type,omitempty"`
	}

	RequestMultiBancoSource struct {
		Type              payments.SourceType `json:"type,omitempty"`
		PaymentCountry    common.Country      `json:"payment_country,omitempty"`
		AccountHolderName string              `json:"account_holder_name,omitempty"`
		BillingDescriptor string              `json:"billing_descriptor,omitempty"`
	}

	RequestP24Source struct {
		Type               payments.SourceType `json:"type,omitempty"`
		PaymentCountry     common.Country      `json:"payment_country,omitempty"`
		AccountHolderName  string              `json:"account_holder_name,omitempty"`
		AccountHolderEmail string              `json:"account_holder_email,omitempty"`
		BillingDescriptor  string              `json:"billing_descriptor,omitempty"`
	}

	RequestPayPalSource struct {
		Type payments.SourceType `json:"type,omitempty"`
		Plan *BillingPlan        `json:"plan,omitempty"`
	}

	RequestPostFinanceSource struct {
		Type              payments.SourceType `json:"type,omitempty"`
		PaymentCountry    common.Country      `json:"payment_country,omitempty"`
		AccountHolderName string              `json:"account_holder_name,omitempty"`
		BillingDescriptor string              `json:"billing_descriptor,omitempty"`
	}

	RequestQPaySource struct {
		Type        payments.SourceType `json:"type,omitempty"`
		Quantity    int                 `json:"quantity,omitempty"`
		Description string              `json:"description,omitempty"`
		Language    string              `json:"language,omitempty"`
		NationalId  string              `json:"national_id,omitempty"`
	}

	RequestSofortSource struct {
		Type         payments.SourceType `json:"type,omitempty"`
		CountryCode  common.Country      `json:"countryCode,omitempty"`
		LanguageCode string              `json:"languageCode,omitempty"`
	}

	RequestStcPaySource struct {
		Type payments.SourceType `json:"type,omitempty"`
	}

	RequestTamaraSource struct {
		Type           payments.SourceType `json:"type,omitempty"`
		BillingAddress *common.Address     `json:"billing_address,omitempty"`
	}

	RequestWeChatPaySource struct {
		Type           payments.SourceType `json:"type,omitempty"`
		BillingAddress *common.Address     `json:"billing_address,omitempty"`
	}
)

func NewRequestAfterPaySource() *RequestAfterPaySource {
	return &RequestAfterPaySource{Type: payments.Afterpay}
}

func NewRequestAlipayPlusSource() *RequestAlipayPlusSource {
	return &RequestAlipayPlusSource{Type: payments.AlipayPlus}
}

func NewRequestAlipayPlusCNSource() *RequestAlipayPlusSource {
	return &RequestAlipayPlusSource{Type: payments.AlipayCn}
}

func NewRequestAlipayPlusGCashSource() *RequestAlipayPlusSource {
	return &RequestAlipayPlusSource{Type: payments.Gcash}
}

func NewRequestAlipayPlusHKSource() *RequestAlipayPlusSource {
	return &RequestAlipayPlusSource{Type: payments.AlipayHk}
}

func NewRequestAlipayPlusDanaSource() *RequestAlipayPlusSource {
	return &RequestAlipayPlusSource{Type: payments.Dana}
}

func NewRequestAlipayPlusKakaoPaySource() *RequestAlipayPlusSource {
	return &RequestAlipayPlusSource{Type: payments.Kakaopay}
}

func NewRequestAlipayPlusTrueMoneySource() *RequestAlipayPlusSource {
	return &RequestAlipayPlusSource{Type: payments.Truemoney}
}

func NewRequestAlipayPlusTNGSource() *RequestAlipayPlusSource {
	return &RequestAlipayPlusSource{Type: payments.Tng}
}

func NewRequestAlmaSource() *RequestAlmaSource {
	return &RequestAlmaSource{Type: payments.Alma}
}

func NewRequestBancontactSource() *RequestBancontactSource {
	return &RequestBancontactSource{Type: payments.BancontactSource}
}

func NewRequestBenefitSource() *RequestBenefitSource {
	return &RequestBenefitSource{Type: payments.Benefit}
}

func NewRequestEpsSource() *RequestEpsSource {
	return &RequestEpsSource{Type: payments.EpsSource}
}

func NewRequestFawrySource() *RequestFawrySource {
	return &RequestFawrySource{Type: payments.FawrySource}
}

func NewRequestGiropaySource() *RequestGiropaySource {
	return &RequestGiropaySource{Type: payments.GiropaySource}
}

func NewRequestIdealSource() *RequestIdealSource {
	return &RequestIdealSource{Type: payments.IdealSource}
}

func NewRequestKlarnaSource() *RequestKlarnaSource {
	return &RequestKlarnaSource{Type: payments.KlarnaSource}
}

func NewRequestKnetSource() *RequestKnetSource {
	return &RequestKnetSource{Type: payments.KnetSource}
}

func NewRequestMbwaySource() *RequestMbwaySource {
	return &RequestMbwaySource{Type: payments.Mbway}
}

func NewRequestMultiBancoSource() *RequestMultiBancoSource {
	return &RequestMultiBancoSource{Type: payments.MultiBancoSource}
}

func NewRequestP24Source() *RequestP24Source {
	return &RequestP24Source{Type: payments.P24Source}
}

func NewRequestPayPalSource() *RequestPayPalSource {
	return &RequestPayPalSource{Type: payments.PayPalSource}
}

func NewRequestPostFinanceSource() *RequestPostFinanceSource {
	return &RequestPostFinanceSource{Type: payments.Postfinance}
}

func NewRequestQPaySource() *RequestQPaySource {
	return &RequestQPaySource{Type: payments.QPaySource}
}

func NewRequestSofortSource() *RequestSofortSource {
	return &RequestSofortSource{Type: payments.SofortSource}
}

func NewRequestStcPaySource() *RequestStcPaySource {
	return &RequestStcPaySource{Type: payments.SofortSource}
}

func NewRequestTamaraSource() *RequestTamaraSource {
	return &RequestTamaraSource{Type: payments.TamaraSource}
}

func NewRequestWeChatPaySource() *RequestWeChatPaySource {
	return &RequestWeChatPaySource{Type: payments.Wechatpay}
}
