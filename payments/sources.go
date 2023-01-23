package payments

type SourceType string

const (
	CardSource            SourceType = "card"
	IdSource              SourceType = "id"
	CustomerSource        SourceType = "customer"
	NetworkTokenSource    SourceType = "network_token"
	TokenSource           SourceType = "token"
	DLocalSource          SourceType = "dLocal"
	AlipaySource          SourceType = "alipay"
	BenefitPaySource      SourceType = "benefitpay"
	BoletoSource          SourceType = "boleto"
	EpsSource             SourceType = "eps"
	GiropaySource         SourceType = "giropay"
	IdealSource           SourceType = "ideal"
	KlarnaSource          SourceType = "klarna"
	KnetSource            SourceType = "knet"
	OxxoSource            SourceType = "oxxo"
	P24Source             SourceType = "p24"
	PagoFacilSource       SourceType = "pagofacil"
	PayPalSource          SourceType = "paypal"
	PoliSource            SourceType = "poli"
	RapiPagoSource        SourceType = "rapipago"
	BancontactSource      SourceType = "bancontact"
	FawrySource           SourceType = "fawry"
	QPaySource            SourceType = "qpay"
	MultiBancoSource      SourceType = "multibanco"
	SepaSource            SourceType = "sepa"
	SofortSource          SourceType = "sofort"
	AlipayHk              SourceType = "alipay_hk"
	AlipayCn              SourceType = "alipay_cn"
	AlipayPlus            SourceType = "alipay_plus"
	Gcash                 SourceType = "gcash"
	Wechatpay             SourceType = "wechatpay"
	Dana                  SourceType = "dana"
	Kakaopay              SourceType = "kakaopay"
	Truemoney             SourceType = "truemoney"
	Tng                   SourceType = "tng"
	Afterpay              SourceType = "afterpay"
	Benefit               SourceType = "benefit"
	Mbway                 SourceType = "mbway"
	Postfinance           SourceType = "postfinance"
	Stcpay                SourceType = "stcpay"
	Alma                  SourceType = "alma"
	BankAccountSource     SourceType = "bank_account"
	ProviderTokenSource   SourceType = "provider_token"
	CurrencyAccountSource SourceType = "currency_account"
	EntitySource          SourceType = "entity"
	TamaraSource          SourceType = "tamara"
)

type (
	PaymentSource interface {
		GetType() SourceType
	}
)
