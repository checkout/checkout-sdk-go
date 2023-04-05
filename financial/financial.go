package financial

import (
	"time"

	"github.com/checkout/checkout-sdk-go/common"
)

const (
	FinancialActionsPath = "financial-actions"
)

type Region string

const (
	Domestic      Region = "Domestic"
	EEA           Region = "EEA"
	International Region = "International"
)

type (
	QueryFilter struct {
		PaymentId       string `url:"payment_id,omitempty"`
		ActionId        string `url:"action_id,omitempty"`
		Reference       string `json:"reference,omitempty"`
		Limit           int    `url:"limit,omitempty"`
		PaginationToken string `url:"pagination_token,omitempty"`
	}
)

type (
	QueryResponse struct {
		HttpMetadata common.HttpMetadata
		Count        int                    `json:"count,omitempty"`
		Limit        int                    `json:"limit,omitempty"`
		Data         []FinancialAction      `json:"data,omitempty"`
		Links        map[string]common.Link `json:"_links"`
	}

	FinancialAction struct {
		PaymentId            string              `json:"payment_id,omitempty"`
		ActionId             string              `json:"action_id,omitempty"`
		ActionType           string              `json:"action_type,omitempty"`
		EntityId             string              `json:"entity_id,omitempty"`
		SubEntityId          string              `json:"sub_entity_id,omitempty"`
		CurrencyAccountId    string              `json:"currency_account_id,omitempty"`
		PaymentMethod        string              `json:"payment_method,omitempty"`
		ProcessingChannelId  string              `json:"processing_channel_id,omitempty"`
		Reference            string              `json:"reference,omitempty"`
		Mid                  string              `json:"mid,omitempty"`
		ResponseCode         string              `json:"response_code,omitempty"`
		ResponseDescription  string              `json:"response_description,omitempty"`
		Region               Region              `json:"region,omitempty"`
		CardType             common.CardType     `json:"card_type,omitempty"`
		CardCategory         common.CardCategory `json:"card_category,omitempty"`
		IssuerCountry        common.Country      `json:"issuer_country,omitempty"`
		MerchantCategoryCode string              `json:"merchant_category_code,omitempty"`
		FxTradeId            string              `json:"fx_trade_id,omitempty"`
		ProcessedOn          time.Time           `json:"processed_on,omitempty"`
		RequestedOn          time.Time           `json:"requested_on,omitempty"`
		Breakdown            []ActionBreakdown   `json:"breakdown,omitempty"`
	}

	ActionBreakdown struct {
		BreakdownType                         string          `json:"breakdown_type,omitempty"`
		FxRateApplied                         float64         `json:"fx_rate_applied,omitempty"`
		HoldingCurrency                       common.Currency `json:"holding_currency,omitempty"`
		HoldingCurrencyAmount                 float64         `json:"holding_currency_amount,omitempty"`
		ProcessingCurrency                    common.Currency `json:"processing_currency,omitempty"`
		ProcessingCurrencyAmount              float64         `json:"processing_currency_amount,omitempty"`
		TransactionCurrency                   common.Currency `json:"transaction_currency,omitempty"`
		TransactionCurrencyAccount            float64         `json:"transaction_currency_account,omitempty"`
		ProcessingToTransactionCurrencyFxRate float64         `json:"processing_to_transaction_currency_fx_rate,omitempty"`
		TransactionToHoldingCurrencyFxRate    float64         `json:"transaction_to_holding_currency_fx_rate,omitempty"`
		FeeDetail                             string          `json:"fee_detail,omitempty"`
		ReserveRate                           string          `json:"reserve_rate,omitempty"`
		ReserveReleaseDate                    time.Time       `json:"reserve_release_date,omitempty"`
		ReserveDeductedDate                   time.Time       `json:"reserve_deducted_date,omitempty"`
		TaxFxRate                             float64         `json:"tax_fx_rate,omitempty"`
		EntityCountryTaxCurrency              common.Currency `json:"entity_country_tax_currency,omitempty"`
		TaxCurrencyAmount                     float64         `json:"tax_currency_amount,omitempty"`
	}
)
