package apm

import (
	"time"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/payments"
)

type (
	requestAlipaySource struct {
		Type payments.SourceType `json:"type,omitempty"`
	}

	requestBancontactSource struct {
		Type              payments.SourceType `json:"type,omitempty"`
		PaymentCountry    common.Country      `json:"payment_country,omitempty"`
		AccountHolderName string              `json:"account_holder_name,omitempty"`
		BillingDescriptor string              `json:"billing_descriptor,omitempty"`
		Language          string              `json:"language,omitempty"`
	}

	requestBenefitPaySource struct {
		Type            payments.SourceType `json:"type,omitempty"`
		IntegrationType IntegrationType     `json:"integration_type,omitempty"`
	}

	requestBoletoSource struct {
		Type            payments.SourceType `json:"type,omitempty"`
		IntegrationType IntegrationType     `json:"integration_type,omitempty"`
		Country         common.Country      `json:"country,omitempty"`
		Description     string              `json:"description,omitempty"`
		Payer           *payments.Payer     `json:"payer,omitempty"`
	}

	requestEpsSource struct {
		Type    payments.SourceType `json:"type,omitempty"`
		Purpose string              `json:"purpose,omitempty"`
		Bic     string              `json:"bic,omitempty"`
	}

	requestFawrySource struct {
		Type              payments.SourceType `json:"type,omitempty"`
		Description       string              `json:"description,omitempty"`
		CustomerProfileId string              `json:"customer_profile_id,omitempty"`
		CustomerEmail     string              `json:"customer_email,omitempty"`
		CustomerMobile    string              `json:"customer_mobile,omitempty"`
		ExpiresOn         *time.Time          `json:"expires_on,omitempty"`
		Products          []FawryProduct      `json:"products,omitempty"`
	}

	requestGiropaySource struct {
		Type       payments.SourceType `json:"type,omitempty"`
		Purpose    string              `json:"purpose,omitempty"`
		Bic        string              `json:"bic,omitempty"`
		InfoFields []InfoFields        `json:"info_fields,omitempty"`
	}

	requestIdealSource struct {
		Type        payments.SourceType `json:"type,omitempty"`
		Description string              `json:"description,omitempty"`
		Bic         string              `json:"bic,omitempty"`
		Language    string              `json:"language,omitempty"`
	}

	requestKlarnaSource struct {
		Type                   payments.SourceType      `json:"type,omitempty"`
		AuthorizationToken     string                   `json:"authorization_token,omitempty"`
		Locale                 string                   `json:"locale,omitempty"`
		PurchaseCountry        string                   `json:"purchase_country,omitempty"`
		AutoCapture            bool                     `json:"auto_capture,omitempty"`
		BillingAddress         *common.Address          `json:"billing_address,omitempty"`
		ShippingAddress        map[string]interface{}   `json:"shipping_address,omitempty"`
		TaxAmount              int64                    `json:"tax_amount,omitempty"`
		Products               []map[string]interface{} `json:"products,omitempty"`
		Customer               map[string]interface{}   `json:"customer,omitempty"`
		MerchantReference1     string                   `json:"merchant_reference1,omitempty"`
		MerchantReference2     string                   `json:"merchant_reference2,omitempty"`
		MerchantData           string                   `json:"merchant_data,omitempty"`
		Attachment             map[string]interface{}   `json:"attachment,omitempty"`
		CustomPaymentMethodIds []map[string]string      `json:"custom_payment_method_ids,omitempty"`
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

	requestMultiBancoSource struct {
		Type              payments.SourceType `json:"type,omitempty"`
		PaymentCountry    common.Country      `json:"payment_country,omitempty"`
		AccountHolderName string              `json:"account_holder_name,omitempty"`
		BillingDescriptor string              `json:"billing_descriptor,omitempty"`
	}

	requestOxxoSource struct {
		Type            payments.SourceType `json:"type,omitempty"`
		IntegrationType IntegrationType     `json:"integration_type,omitempty"`
		Country         common.Country      `json:"country,omitempty"`
		Description     string              `json:"description,omitempty"`
		Payer           *payments.Payer     `json:"payer,omitempty"`
	}

	requestP24Source struct {
		Type               payments.SourceType `json:"type,omitempty"`
		PaymentCountry     common.Country      `json:"payment_country,omitempty"`
		AccountHolderName  string              `json:"account_holder_name,omitempty"`
		AccountHolderEmail string              `json:"account_holder_email,omitempty"`
		BillingDescriptor  string              `json:"billing_descriptor,omitempty"`
	}

	requestPagoFacilSource struct {
		Type            payments.SourceType `json:"type,omitempty"`
		IntegrationType IntegrationType     `json:"integration_type,omitempty"`
		Country         common.Country      `json:"country,omitempty"`
		Description     string              `json:"description,omitempty"`
		Payer           *payments.Payer     `json:"payer,omitempty"`
	}

	requestPayPalSource struct {
		Type          payments.SourceType    `json:"type,omitempty"`
		InvoiceNumber string                 `json:"invoice_number,omitempty"`
		RecipientName string                 `json:"recipient_name,omitempty"`
		LogoUrl       string                 `json:"logo_url,omitempty"`
		Stc           map[string]interface{} `json:"stc,omitempty"`
	}

	requestPoliSource struct {
		Type payments.SourceType `json:"type,omitempty"`
	}

	requestQPaySource struct {
		Type        payments.SourceType `json:"type,omitempty"`
		Quantity    int                 `json:"quantity,omitempty"`
		Description string              `json:"description,omitempty"`
		Language    string              `json:"language,omitempty"`
		NationalId  string              `json:"national_id,omitempty"`
	}

	requestRapiPagoSource struct {
		Type            payments.SourceType `json:"type,omitempty"`
		IntegrationType IntegrationType     `json:"integration_type,omitempty"`
		Country         common.Country      `json:"country,omitempty"`
		Description     string              `json:"description,omitempty"`
		Payer           *payments.Payer     `json:"payer,omitempty"`
	}

	requestSepaSource struct {
		Type payments.SourceType `json:"type,omitempty"`
		Id   string              `json:"id,omitempty"`
	}

	requestSofortSource struct {
		Type         payments.SourceType `json:"type,omitempty"`
		CountryCode  common.Country      `json:"countryCode,omitempty"`
		LanguageCode string              `json:"languageCode,omitempty"`
	}
)

func NewRequestAlipaySource() *requestAlipaySource {
	return &requestAlipaySource{Type: payments.AlipaySource}
}

func NewRequestBancontactSource() *requestBancontactSource {
	return &requestBancontactSource{
		Type: payments.BancontactSource,
	}
}

func NewRequestBenefitPaySource() *requestBenefitPaySource {
	return &requestBenefitPaySource{
		Type:            payments.BenefitPaySource,
		IntegrationType: Mobile,
	}
}

func NewRequestBoletoSource() *requestBoletoSource {
	return &requestBoletoSource{
		Type: payments.BoletoSource,
	}
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

func NewRequestMultiBancoSource() *requestMultiBancoSource {
	return &requestMultiBancoSource{
		Type: payments.MultiBancoSource,
	}
}

func NewRequestOxxoSource() *requestOxxoSource {
	return &requestOxxoSource{
		Type: payments.OxxoSource,
	}
}

func NewRequestP24Source() *requestP24Source {
	return &requestP24Source{
		Type: payments.P24Source,
	}
}

func NewRequestPagoFacilSource() *requestPagoFacilSource {
	return &requestPagoFacilSource{
		Type:            payments.PagoFacilSource,
		IntegrationType: Redirect,
	}
}

func NewRequestPayPalSource() *requestPayPalSource {
	return &requestPayPalSource{Type: payments.PayPalSource}
}

func NewRequestPoliSource() *requestPoliSource {
	return &requestPoliSource{Type: payments.PoliSource}
}

func NewRequestQPaySource() *requestQPaySource {
	return &requestQPaySource{Type: payments.QPaySource}
}

func NewRequestRapiPagoSource() *requestRapiPagoSource {
	return &requestRapiPagoSource{
		Type:            payments.RapiPagoSource,
		IntegrationType: Redirect,
	}
}

func NewRequestSepaSource() *requestSepaSource {
	return &requestSepaSource{Type: payments.SepaSource}
}

func NewRequestSofortSource() *requestSofortSource {
	return &requestSofortSource{Type: payments.SofortSource}
}

func (s *requestAlipaySource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestBancontactSource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestBenefitPaySource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestBoletoSource) GetType() payments.SourceType {
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

func (s *requestMultiBancoSource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestOxxoSource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestP24Source) GetType() payments.SourceType {
	return s.Type
}

func (s *requestPagoFacilSource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestPayPalSource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestPoliSource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestQPaySource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestRapiPagoSource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestSepaSource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestSofortSource) GetType() payments.SourceType {
	return s.Type
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
