package forex

import (
	"time"

	"github.com/checkout/checkout-sdk-go-beta/common"
)

const (
	forex  = "forex"
	quotes = "quotes"
)

type (
	QuoteRequest struct {
		SourceCurrency      common.Currency `json:"source_currency,omitempty"`
		SourceAmount        int             `json:"source_amount,omitempty"`
		DestinationCurrency common.Currency `json:"destination_currency,omitempty"`
		DestinationAmount   int             `json:"destination_amount,omitempty"`
		ProcessingChannelId string          `json:"processing_channel_id,omitempty"`
	}

	QuoteResponse struct {
		HttpMetadata        common.HttpMetadata `json:"http_metadata,omitempty"`
		Id                  string              `json:"id,omitempty"`
		SourceCurrency      common.Currency     `json:"source_currency,omitempty"`
		SourceAmount        int                 `json:"source_amount,omitempty"`
		DestinationCurrency common.Currency     `json:"destination_currency,omitempty"`
		DestinationAmount   int                 `json:"destination_amount,omitempty"`
		Rate                float64             `json:"rate,omitempty"`
		ExpiresOn           time.Time           `json:"expires_on,omitempty"`
		IsSingleUse         bool                `json:"is_single_use,omitempty"`
	}
)
