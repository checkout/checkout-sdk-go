package payments

type SourceType string

const (
	Afterpay              SourceType = "afterpay"
	AlipayCn              SourceType = "alipay_cn"
	AlipayHk              SourceType = "alipay_hk"
	AlipayPlus            SourceType = "alipay_plus"
	AlipaySource          SourceType = "alipay"
	Alma                  SourceType = "alma"
	ApplepaySource        SourceType = "applepay"
	BancontactSource      SourceType = "bancontact"
	BankAccountSource     SourceType = "bank_account"
	Benefit               SourceType = "benefit"
	BenefitPaySource      SourceType = "benefitpay"
	BoletoSource          SourceType = "boleto"
	CardSource            SourceType = "card"
	CurrencyAccountSource SourceType = "currency_account"
	CustomerSource        SourceType = "customer"
	CvConnectSource       SourceType = "cvconnect"
	Dana                  SourceType = "dana"
	DLocalSource          SourceType = "dLocal"
	EntitySource          SourceType = "entity"
	EpsSource             SourceType = "eps"
	FawrySource           SourceType = "fawry"
	Gcash                 SourceType = "gcash"
	GiropaySource         SourceType = "giropay"
	GooglepaySource       SourceType = "googlepay"
	IdealSource           SourceType = "ideal"
	IdSource              SourceType = "id"
	IllicadoSource        SourceType = "illicado"
	Kakaopay              SourceType = "kakaopay"
	KlarnaSource          SourceType = "klarna"
	KnetSource            SourceType = "knet"
	Mbway                 SourceType = "mbway"
	MultiBancoSource      SourceType = "multibanco"
	NetworkTokenSource    SourceType = "network_token"
	OxxoSource            SourceType = "oxxo"
	P24Source             SourceType = "p24"
	PagoFacilSource       SourceType = "pagofacil"
	PayPalSource          SourceType = "paypal"
	PoliSource            SourceType = "poli"
	Postfinance           SourceType = "postfinance"
	ProviderTokenSource   SourceType = "provider_token"
	QPaySource            SourceType = "qpay"
	RapiPagoSource        SourceType = "rapipago"
	SepaSource            SourceType = "sepa"
	SofortSource          SourceType = "sofort"
	Stcpay                SourceType = "stcpay"
	TamaraSource          SourceType = "tamara"
	Tng                   SourceType = "tng"
	TokenSource           SourceType = "token"
	Truemoney             SourceType = "truemoney"
	TrustlySource         SourceType = "trustly"
	Wechatpay             SourceType = "wechatpay"
)

type (
	PaymentSource interface {
		GetType() SourceType
	}
)
