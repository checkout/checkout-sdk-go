package forex

import (
	"time"

	"github.com/checkout/checkout-sdk-go/common"
)

const (
	forex  = "forex"
	quotes = "quotes"
	rates  = "rates"
)

type Source string

const (
	Visa       Source = "visa"
	MasterCard Source = "mastercard"
)

type (
	QuoteRequest struct {
		SourceCurrency      common.Currency `json:"source_currency,omitempty"`
		SourceAmount        int64           `json:"source_amount,omitempty"`
		DestinationCurrency common.Currency `json:"destination_currency,omitempty"`
		DestinationAmount   int64           `json:"destination_amount,omitempty"`
		ProcessingChannelId string          `json:"processing_channel_id,omitempty"`
	}

	QuoteResponse struct {
		HttpMetadata        common.HttpMetadata `json:"http_metadata,omitempty"`
		Id                  string              `json:"id,omitempty"`
		SourceCurrency      common.Currency     `json:"source_currency,omitempty"`
		SourceAmount        int64               `json:"source_amount,omitempty"`
		DestinationCurrency common.Currency     `json:"destination_currency,omitempty"`
		DestinationAmount   int64               `json:"destination_amount,omitempty"`
		Rate                float64             `json:"rate,omitempty"`
		ExpiresOn           *time.Time          `json:"expires_on,omitempty"`
		IsSingleUse         bool                `json:"is_single_use,omitempty"`
	}
)

type (
	RatesQuery struct {
		Product             string `url:"product,omitempty"`
		Source              Source `url:"source,omitempty"`
		CurrencyPairs       string `url:"currency_pairs,omitempty"`
		ProcessingChannelId string `url:"processing_channel_id,omitempty"`
	}

	RatesResponse struct {
		Product              string   `json:"product,omitempty"`
		Source               Source   `json:"source,omitempty"`
		Rates                []Rate   `json:"rates,omitempty"`
		InvalidCurrencyPairs []string `json:"invalid_currency_pairs,omitempty"`
	}
)

type Rate struct {
	ExchangeRate float64 `json:"exchange_rate,omitempty"`
	CurrencyPair string  `json:"currency_pair,omitempty"`
}
