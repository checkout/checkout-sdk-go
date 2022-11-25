package apm

import (
	"time"

	"github.com/checkout/checkout-sdk-go-beta/common"
	"github.com/checkout/checkout-sdk-go-beta/payments"
)

type (
	RequestAlipaySource struct {
		Type payments.SourceType `json:"type,omitempty"`
	}

	RequestBancontactSource struct {
		Type              payments.SourceType `json:"type,omitempty"`
		PaymentCountry    common.Country      `json:"payment_country,omitempty"`
		AccountHolderName string              `json:"account_holder_name,omitempty"`
		BillingDescriptor string              `json:"billing_descriptor,omitempty"`
		Language          string              `json:"language,omitempty"`
	}

	RequestBenefitPaySource struct {
		Type            payments.SourceType `json:"type,omitempty"`
		IntegrationType IntegrationType     `json:"integration_type,omitempty"`
	}

	RequestBoletoSource struct {
		Type            payments.SourceType `json:"type,omitempty"`
		IntegrationType IntegrationType     `json:"integration_type,omitempty"`
		Country         common.Country      `json:"country,omitempty"`
		Description     string              `json:"description,omitempty"`
		Payer           *payments.Payer     `json:"payer,omitempty"`
	}

	RequestEpsSource struct {
		Type    payments.SourceType `json:"type,omitempty"`
		Purpose string              `json:"purpose,omitempty"`
		Bic     string              `json:"bic,omitempty"`
	}

	RequestFawrySource struct {
		Type              payments.SourceType `json:"type,omitempty"`
		Description       string              `json:"description,omitempty"`
		CustomerProfileId string              `json:"customer_profile_id,omitempty"`
		CustomerEmail     string              `json:"customer_email,omitempty"`
		CustomerMobile    string              `json:"customer_mobile,omitempty"`
		ExpiresOn         time.Time           `json:"expires_on,omitempty"`
		Products          []FawryProduct      `json:"products,omitempty"`
	}

	RequestGiropaySource struct {
		Type       payments.SourceType `json:"type,omitempty"`
		Purpose    string              `json:"purpose,omitempty"`
		Bic        string              `json:"bic,omitempty"`
		InfoFields []InfoFields        `json:"info_fields,omitempty"`
	}

	RequestIdealSource struct {
		Type        payments.SourceType `json:"type,omitempty"`
		Description string              `json:"description,omitempty"`
		Bic         string              `json:"bic,omitempty"`
		Language    string              `json:"language,omitempty"`
	}

	RequestKlarnaSource struct {
		Type                   payments.SourceType      `json:"type,omitempty"`
		AuthorizationToken     string                   `json:"authorization_token,omitempty"`
		Locale                 string                   `json:"locale,omitempty"`
		PurchaseCountry        string                   `json:"purchase_country,omitempty"`
		AutoCapture            bool                     `json:"auto_capture,omitempty"`
		BillingAddress         *common.Address          `json:"billing_address,omitempty"`
		ShippingAddress        map[string]interface{}   `json:"shipping_address,omitempty"`
		TaxAmount              int                      `json:"tax_amount,omitempty"`
		Products               []map[string]interface{} `json:"products,omitempty"`
		Customer               map[string]interface{}   `json:"customer,omitempty"`
		MerchantReference1     string                   `json:"merchant_reference1,omitempty"`
		MerchantReference2     string                   `json:"merchant_reference2,omitempty"`
		MerchantData           string                   `json:"merchant_data,omitempty"`
		Attachment             map[string]interface{}   `json:"attachment,omitempty"`
		CustomPaymentMethodIds []map[string]string      `json:"custom_payment_method_ids,omitempty"`
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

	RequestMultiBancoSource struct {
		Type              payments.SourceType `json:"type,omitempty"`
		PaymentCountry    common.Country      `json:"payment_country,omitempty"`
		AccountHolderName string              `json:"account_holder_name,omitempty"`
		BillingDescriptor string              `json:"billing_descriptor,omitempty"`
	}

	RequestOxxoSource struct {
		Type            payments.SourceType `json:"type,omitempty"`
		IntegrationType IntegrationType     `json:"integration_type,omitempty"`
		Country         common.Country      `json:"country,omitempty"`
		Description     string              `json:"description,omitempty"`
		Payer           *payments.Payer     `json:"payer,omitempty"`
	}

	RequestP24Source struct {
		Type               payments.SourceType `json:"type,omitempty"`
		PaymentCountry     common.Country      `json:"payment_country,omitempty"`
		AccountHolderName  string              `json:"account_holder_name,omitempty"`
		AccountHolderEmail string              `json:"account_holder_email,omitempty"`
		BillingDescriptor  string              `json:"billing_descriptor,omitempty"`
	}

	RequestPagoFacilSource struct {
		Type            payments.SourceType `json:"type,omitempty"`
		IntegrationType IntegrationType     `json:"integration_type,omitempty"`
		Country         common.Country      `json:"country,omitempty"`
		Description     string              `json:"description,omitempty"`
		Payer           *payments.Payer     `json:"payer,omitempty"`
	}

	RequestPayPalSource struct {
		Type          payments.SourceType    `json:"type,omitempty"`
		InvoiceNumber string                 `json:"invoice_number,omitempty"`
		RecipientName string                 `json:"recipient_name,omitempty"`
		LogoUrl       string                 `json:"logo_url,omitempty"`
		Stc           map[string]interface{} `json:"stc,omitempty"`
	}

	RequestPoliSource struct {
		Type payments.SourceType `json:"type,omitempty"`
	}

	RequestQPaySource struct {
		Type        payments.SourceType `json:"type,omitempty"`
		Quantity    int                 `json:"quantity,omitempty"`
		Description string              `json:"description,omitempty"`
		Language    string              `json:"language,omitempty"`
		NationalId  string              `json:"national_id,omitempty"`
	}

	RequestRapiPagoSource struct {
		Type            payments.SourceType `json:"type,omitempty"`
		IntegrationType IntegrationType     `json:"integration_type,omitempty"`
		Country         common.Country      `json:"country,omitempty"`
		Description     string              `json:"description,omitempty"`
		Payer           *payments.Payer     `json:"payer,omitempty"`
	}

	RequestSepaSource struct {
		Type payments.SourceType `json:"type,omitempty"`
		Id   string              `json:"id,omitempty"`
	}

	RequestSofortSource struct {
		Type         payments.SourceType `json:"type,omitempty"`
		CountryCode  common.Country      `json:"countryCode,omitempty"`
		LanguageCode string              `json:"languageCode,omitempty"`
	}
)

func NewRequestAlipaySource() *RequestAlipaySource {
	return &RequestAlipaySource{Type: payments.AlipaySource}
}

func NewRequestBancontactSource() *RequestBancontactSource {
	return &RequestBancontactSource{
		Type: payments.BancontactSource,
	}
}

func NewRequestBenefitPaySource() *RequestBenefitPaySource {
	return &RequestBenefitPaySource{
		Type:            payments.BenefitPaySource,
		IntegrationType: Mobile,
	}
}

func NewRequestBoletoSource() *RequestBoletoSource {
	return &RequestBoletoSource{
		Type: payments.BoletoSource,
	}
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

func NewRequestMultiBancoSource() *RequestMultiBancoSource {
	return &RequestMultiBancoSource{
		Type: payments.MultiBancoSource,
	}
}

func NewRequestOxxoSource() *RequestOxxoSource {
	return &RequestOxxoSource{
		Type: payments.OxxoSource,
	}
}

func NewRequestP24Source() *RequestP24Source {
	return &RequestP24Source{
		Type: payments.P24Source,
	}
}

func NewRequestPagoFacilSource() *RequestPagoFacilSource {
	return &RequestPagoFacilSource{
		Type:            payments.PagoFacilSource,
		IntegrationType: Redirect,
	}
}

func NewRequestPayPalSource() *RequestPayPalSource {
	return &RequestPayPalSource{Type: payments.PayPalSource}
}

func NewRequestPoliSource() *RequestPoliSource {
	return &RequestPoliSource{Type: payments.PoliSource}
}

func NewRequestQPaySource() *RequestQPaySource {
	return &RequestQPaySource{Type: payments.QPaySource}
}

func NewRequestRapiPagoSource() *RequestRapiPagoSource {
	return &RequestRapiPagoSource{
		Type:            payments.RapiPagoSource,
		IntegrationType: Redirect,
	}
}

func NewRequestSepaSource() *RequestSepaSource {
	return &RequestSepaSource{Type: payments.SepaSource}
}

func NewRequestSofortSource() *RequestSofortSource {
	return &RequestSofortSource{Type: payments.SofortSource}
}

type IntegrationType string

const (
	Direct   IntegrationType = "direct"
	Redirect IntegrationType = "redirect"
	Mobile   IntegrationType = "mobile"
)

type InfoFields struct {
	Label string `json:"label,omitempty"`
	Text  string `json:"text,omitempty"`
}

type (
	FawryProduct struct {
		ProductId   string `json:"product_id,omitempty"`
		Quantity    int    `json:"quantity,omitempty"`
		Price       int64  `json:"price,omitempty"`
		Description string `json:"description,omitempty"`
	}
)
