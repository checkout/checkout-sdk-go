package balances

import "github.com/checkout/checkout-sdk-go/common"

const (
	balances = "balances"
)

type (
	QueryFilter struct {
		Query string `url:"query,omitempty"`
	}
)

// QueryResponse
type (
	Balances struct {
		Pending    int64 `json:"pending,omitempty"`
		Available  int64 `json:"available,omitempty"`
		Payable    int64 `json:"payable,omitempty"`
		Collateral int64 `json:"collateral,omitempty"`
	}

	AccountBalance struct {
		Descriptor      string   `json:"descriptor,omitempty"`
		HoldingCurrency string   `json:"holding_currency,omitempty"`
		Balances        Balances `json:"balances,omitempty"`
	}

	QueryResponse struct {
		HttpMetadata common.HttpMetadata
		Data         []AccountBalance `json:"data,omitempty"`
	}
)
