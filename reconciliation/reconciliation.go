package reconciliation

import (
	"time"

	"github.com/checkout/checkout-sdk-go/common"
)

const (
	reporting  = "reporting"
	payments   = "payments"
	statements = "statements"
	download   = "download"
)

type (
	PaymentReportsQuery struct {
		From      time.Time `url:"from,omitempty" layout:"2006-01-02T15:04:05Z"`
		To        time.Time `url:"to,omitempty" layout:"2006-01-02T15:04:05Z"`
		Reference string    `url:"reference,omitempty"`
		Limit     int       `url:"limit,omitempty"`
	}

	PaymentReportsResponse struct {
		HttpMetadata common.HttpMetadata
		Count        int                    `json:"count,omitempty"`
		Data         []PaymentReportData    `json:"data,omitempty"`
		Links        map[string]common.Link `json:"_links"`
	}

	PaymentReportData struct {
		Id                 string                 `json:"id,omitempty"`
		ProcessingCurrency common.Currency        `json:"processing_currency,omitempty"`
		PayoutCurrency     common.Currency        `json:"payout_currency,omitempty"`
		RequestedOn        string                 `json:"requested_on,omitempty"`
		ChannelName        string                 `json:"channel_name,omitempty"`
		Reference          string                 `json:"reference,omitempty"`
		PaymentMethod      string                 `json:"payment_method,omitempty"`
		CardType           string                 `json:"card_type,omitempty"`
		CardCategory       string                 `json:"card_category,omitempty"`
		IssuerCountry      common.Country         `json:"issuer_country,omitempty"`
		MerchantCountry    common.Country         `json:"merchant_country,omitempty"`
		Mid                string                 `json:"mid,omitempty"`
		Actions            []Action               `json:"actions,omitempty"`
		Links              map[string]common.Link `json:"_links"`
	}
)

type Breakdown struct {
	Type                     string  `json:"type,omitempty"`
	Date                     string  `json:"date,omitempty"`
	ProcessingCurrencyAmount float64 `json:"processing_currency_amount,omitempty"`
	PayoutCurrencyAmount     float64 `json:"payout_currency_amount,omitempty"`
}

type Action struct {
	Type                string      `json:"type,omitempty"`
	Id                  string      `json:"id,omitempty"`
	ProcessedOn         string      `json:"processed_on,omitempty"`
	ResponseCode        string      `json:"response_code,omitempty"`
	ResponseDescription string      `json:"response_description,omitempty"`
	Breakdown           []Breakdown `json:"breakdown,omitempty"`
}

type (
	StatementReportsResponse struct {
		HttpMetadata common.HttpMetadata
		Count        int                    `json:"count,omitempty"`
		Data         []StatementReportsData `json:"data,omitempty"`
		Links        map[string]common.Link `json:"_links"`
	}

	StatementReportsData struct {
		Id          string                 `json:"id,omitempty"`
		PeriodStart string                 `json:"period_start,omitempty"`
		PeriodEnd   string                 `json:"period_end,omitempty"`
		Date        string                 `json:"date,omitempty"`
		Payouts     []PayoutStatement      `json:"payouts,omitempty"`
		Links       map[string]common.Link `json:"_links"`
	}
)

type PayoutStatement struct {
	Id                   string                 `json:"id,omitempty"`
	Currency             common.Currency        `json:"currency,omitempty"`
	CarriedForwardAmount float64                `json:"carried_forward_amount,omitempty"`
	CurrentPeriodAmount  float64                `json:"current_period_amount,omitempty"`
	NetAmount            float64                `json:"net_amount,omitempty"`
	Date                 string                 `json:"date,omitempty"`
	PeriodStart          string                 `json:"period_start,omitempty"`
	PeriodEnd            string                 `json:"period_end,omitempty"`
	Status               string                 `json:"status,omitempty"`
	PayoutFee            float64                `json:"payout_fee,omitempty"`
	Links                map[string]common.Link `json:"_links"`
}
