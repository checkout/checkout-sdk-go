package paymentmethods

import "github.com/checkout/checkout-sdk-go/v2/common"

const paymentMethodsPath = "payment-methods"

type PaymentMethodType string

const (
	Accel           PaymentMethodType = "accel"
	Ach             PaymentMethodType = "ach"
	AlipayCn        PaymentMethodType = "alipay_cn"
	AlipayHk        PaymentMethodType = "alipay_hk"
	AlipayPlus      PaymentMethodType = "alipay_plus"
	Alma            PaymentMethodType = "alma"
	Amex            PaymentMethodType = "amex"
	Bancontact      PaymentMethodType = "bancontact"
	BankRedirects   PaymentMethodType = "bank_redirects"
	Bnpl            PaymentMethodType = "bnpl"
	Boost           PaymentMethodType = "boost"
	Bpi             PaymentMethodType = "bpi"
	CardScheme      PaymentMethodType = "card_scheme"
	CartesBancaires PaymentMethodType = "cartes_bancaires"
	ChinaUnionPay   PaymentMethodType = "china_union_pay"
	ConnectWallet   PaymentMethodType = "connect_wallet"
	Dana            PaymentMethodType = "dana"
	Dci             PaymentMethodType = "dci"
	Diners          PaymentMethodType = "diners"
	Discover        PaymentMethodType = "discover"
	Eps             PaymentMethodType = "eps"
	Gcash           PaymentMethodType = "gcash"
	Ideal           PaymentMethodType = "ideal"
	Jcb             PaymentMethodType = "jcb"
	Kakaopay        PaymentMethodType = "kakaopay"
	Klarna          PaymentMethodType = "klarna"
	Knet            PaymentMethodType = "knet"
	Mada            PaymentMethodType = "mada"
	Mastercard      PaymentMethodType = "mastercard"
	Mbway           PaymentMethodType = "mbway"
	Multibanco      PaymentMethodType = "multibanco"
	Nyce            PaymentMethodType = "nyce"
	Omannet         PaymentMethodType = "omannet"
	P24             PaymentMethodType = "p24"
	Paypal          PaymentMethodType = "paypal"
	Paypay          PaymentMethodType = "paypay"
	Pulse           PaymentMethodType = "pulse"
	Qpay            PaymentMethodType = "qpay"
	RabbitLinePay   PaymentMethodType = "rabbit_line_pay"
	Sepa            PaymentMethodType = "sepa"
	Sequra          PaymentMethodType = "sequra"
	Shazam          PaymentMethodType = "shazam"
	Star            PaymentMethodType = "star"
	Stcpay          PaymentMethodType = "stcpay"
	Swish           PaymentMethodType = "swish"
	Tamara          PaymentMethodType = "tamara"
	Tng             PaymentMethodType = "tng"
	Truemoney       PaymentMethodType = "truemoney"
	Upi             PaymentMethodType = "upi"
	Visa            PaymentMethodType = "visa"
	Wallet          PaymentMethodType = "wallet"
	Wechatpay       PaymentMethodType = "wechatpay"
)

type GetPaymentMethodsQuery struct {
	ProcessingChannelId string `url:"processing_channel_id"`
}

type PaymentMethod struct {
	Type              PaymentMethodType `json:"type"`
	Name              string            `json:"name,omitempty"`
	PartnerMerchantId string            `json:"partner_merchant_id,omitempty"`
}

type GetPaymentMethodsResponse struct {
	HttpMetadata common.HttpMetadata
	Methods      []PaymentMethod `json:"methods"`
}
