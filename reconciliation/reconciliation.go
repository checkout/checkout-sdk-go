package reconciliation

import (
	"time"

	"github.com/shiuh-yaw-cko/checkout"
	"github.com/shiuh-yaw-cko/checkout/common"
)

type (
	// Request -
	Request struct {
		*PaymentsParameter
		*PaymentParameter
		*StatementsParameter
		*StatementParameter
	}

	// PaymentsParameter -
	PaymentsParameter struct {
		From      time.Time `url:"from,omitempty"`
		To        time.Time `url:"to,omitempty"`
		Reference string    `url:"reference,omitempty"`
		Limit     uint64    `url:"limit,omitempty"`
	}
	// PaymentParameter -
	PaymentParameter struct {
		Reference string `url:"reference,omitempty"`
	}

	// StatementsParameter -
	StatementsParameter struct {
		From    time.Time `url:"from,omitempty"`
		To      time.Time `url:"to,omitempty"`
		Include string    `url:"include,omitempty"`
	}

	// StatementParameter -
	StatementParameter struct {
		PayoutID       string `url:"payout_id,omitempty"`
		PayoutCurrency string `url:"payout_currency,omitempty"`
		Limit          uint64 `url:"limit,omitempty"`
	}
)

type (
	// Response -
	Response struct {
		StatusResponse   *checkout.StatusResponse `json:"api_response,omitempty"`
		PaymentsReport   *PaymentsReport          `json:"payments_report,omitempty"`
		StatementsReport *StatementsReport        `json:"statements_report,omitempty"`
		CSV              [][]string               `json:"csv,omitempty"`
	}

	// PaymentsReport -
	PaymentsReport struct {
		Count uint64                 `json:"count,omitempty"`
		Data  []Payment              `json:"data,omitempty"`
		Links map[string]common.Link `json:"_links"`
	}

	// Payment -
	Payment struct {
		ID                 string                 `json:"id,omitempty"`
		ProcessingCurrency string                 `json:"processing_currency,omitempty"`
		PayoutCurrency     string                 `json:"payout_currency,omitempty"`
		RequestedOn        string                 `json:"requested_on,omitempty"`
		ChannelName        string                 `json:"channel_name,omitempty"`
		Reference          string                 `json:"reference,omitempty"`
		PaymentMethod      string                 `json:"payment_method,omitempty"`
		CardType           common.CardType        `json:"card_type,omitempty"`
		CardCategory       common.CardCategory    `json:"card_category,omitempty"`
		IssuerCountry      string                 `json:"issuer_country,omitempty"`
		MerchantCountry    string                 `json:"merchant_country,omitempty"`
		MID                string                 `json:"mid,omitempty"`
		Actions            []Action               `json:"actions,omitempty"`
		Links              map[string]common.Link `json:"_links"`
	}

	// Action -
	Action struct {
		Type                string      `json:"type,omitempty"`
		ID                  string      `json:"id,omitempty"`
		ProcessedOn         string      `json:"processed_on,omitempty"`
		ResponseCode        uint64      `json:"response_code,omitempty"`
		ResponseDescription string      `json:"response_description,omitempty"`
		Breakdown           []Breakdown `json:"breakdown,omitempty"`
	}

	// Breakdown -
	Breakdown struct {
		Type                     string   `json:"type,omitempty"`
		Date                     string   `json:"date,omitempty"`
		ProcessingCurrencyAmount *float64 `json:"processing_currency_amount,omitempty"`
		PayoutCurrencyAmount     *float64 `json:"payout_currency_amount,omitempty"`
	}

	// StatementsReport -
	StatementsReport struct {
		Count uint64                 `json:"count,omitempty"`
		Data  []Statement            `json:"data,omitempty"`
		Links map[string]common.Link `json:"_links"`
	}

	// Statement -
	Statement struct {
		ID          string                 `json:"id,omitempty"`
		PeriodStart string                 `json:"period_start,omitempty"`
		PeriodEnd   string                 `json:"period_end,omitempty"`
		Date        string                 `json:"date,omitempty"`
		Payouts     *[]Payout              `json:"payouts,omitempty"`
		Links       map[string]common.Link `json:"_links"`
	}

	// Payout -
	Payout struct {
		Currency               string                  `json:"currency,omitempty"`
		CarriedForwardAmount   *float64                `json:"carried_forward_amount,omitempty"`
		CurrentPeriodAmount    *float64                `json:"current_period_amount,omitempty"`
		NetAmount              *float64                `json:"net_amount,omitempty"`
		Date                   string                  `json:"date,omitempty"`
		PeriodStart            string                  `json:"period_start,omitempty"`
		PeriodEnd              string                  `json:"period_end,omitempty"`
		ID                     string                  `json:"id,omitempty"`
		Status                 string                  `json:"status,omitempty"`
		PayoutFee              *float64                `json:"payout_fee,omitempty"`
		CurrentPeriodBreakdown *CurrentPeriodBreakdown `json:"current_period_breakdown,omitempty"`
		Links                  map[string]common.Link  `json:"_links"`
	}

	// CurrentPeriodBreakdown -
	CurrentPeriodBreakdown struct {
		ProcessedAmount         *float64                 `json:"processed_amount,omitempty"`
		RefundAmount            *float64                 `json:"refund_amount,omitempty"`
		ChargebackAmount        *float64                 `json:"chargeback_amount,omitempty"`
		ProcessingFees          *float64                 `json:"processing_fees,omitempty"`
		ProcessingFeesBreakdown *ProcessingFeesBreakdown `json:"processing_fees_breakdown,omitempty"`
		RollingReserveAmount    *float64                 `json:"rolling_reserve_amount,omitempty"`
		Tax                     *float64                 `json:"tax,omitempty"`
		AdminFees               *float64                 `json:"admin_fees,omitempty"`
		GeneralAdjustments      *float64                 `json:"general_adjustments,omitempty"`
	}

	// ProcessingFeesBreakdown -
	ProcessingFeesBreakdown struct {
		InterchangeFees           *float64 `json:"interchange_fees,omitempty"`
		SchemeAndOtherNetworkFees *float64 `json:"scheme_and_other_network_fees,omitempty"`
		PremiumAndAPMFees         *float64 `json:"premium_and_apm_fees,omitempty"`
		ChargebackFees            *float64 `json:"chargeback_fees,omitempty"`
		PaymentGatewayFees        *float64 `json:"payment_gateway_fees,omitempty"`
		AccountUpdaterFees        *float64 `json:"account_updater_fees,omitempty"`
	}
)
