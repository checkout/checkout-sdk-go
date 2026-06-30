package balances

import (
	"time"

	"github.com/checkout/checkout-sdk-go/v2/common"
)

const (
	balances = "balances"
)

type (
	QueryFilter struct {
		Query                 string     `url:"query,omitempty"`
		WithCurrencyAccountId bool       `url:"withCurrencyAccountId,omitempty"`
		BalancesAt            *time.Time `url:"balancesAt,omitempty"`
	}
)

// QueryResponse
type (
	CollateralBreakdown struct {
		FixedReserve   int64 `json:"fixed_reserve,omitempty"`
		RollingReserve int64 `json:"rolling_reserve,omitempty"`
	}

	Balances struct {
		Pending             int64                `json:"pending,omitempty"`
		Available           int64                `json:"available,omitempty"`
		Payable             int64                `json:"payable,omitempty"`
		Collateral          int64                `json:"collateral,omitempty"`
		Operational         int64                `json:"operational,omitempty"`
		CollateralBreakdown *CollateralBreakdown `json:"collateral_breakdown,omitempty"`
	}

	AccountBalance struct {
		Descriptor        string          `json:"descriptor,omitempty"`
		CurrencyAccountId string          `json:"currency_account_id,omitempty"`
		HoldingCurrency   common.Currency `json:"holding_currency,omitempty"`
		BalancesAsOf      *time.Time      `json:"balances_as_of,omitempty"`
		Balances          Balances        `json:"balances,omitempty"`
	}

	QueryResponse struct {
		HttpMetadata common.HttpMetadata
		Data         []AccountBalance `json:"data,omitempty"`
	}
)
