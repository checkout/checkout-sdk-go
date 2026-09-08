package balances

import (
	"time"

	"github.com/checkout/checkout-sdk-go/v3/common"
)

const (
	balances          = "balances"
	entities          = "entities"
	currencyAccounts  = "currency-accounts"
	topUpInstructions = "top-up-instructions"
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

// Response types for
// GET /entities/{entityId}/currency-accounts/{currencyAccountId}/top-up-instructions.
type (
	// TopUpFundingDetails holds the bank details for a single funding rail.
	//
	// BeneficiaryAccountName and BankName are the only fields always returned. The remaining
	// fields vary by rail and the receiving bank's jurisdiction, and are omitted when they do
	// not apply.
	TopUpFundingDetails struct {
		// BeneficiaryAccountName is the name of the account that receives the funds.
		// [Required]
		BeneficiaryAccountName string `json:"beneficiary_account_name,omitempty"`
		// BeneficiaryAddress is the address of the beneficiary, if the rail requires it.
		// [Optional]
		BeneficiaryAddress string `json:"beneficiary_address,omitempty"`
		// BankName is the name of the bank that receives the funds.
		// [Required]
		BankName string `json:"bank_name,omitempty"`
		// BankAddress is the address of the receiving bank, if the rail requires it.
		// [Optional]
		BankAddress string `json:"bank_address,omitempty"`
		// AccountNumber is the account number of the receiving account.
		// [Optional]
		AccountNumber string `json:"account_number,omitempty"`
		// SortCode is the sort code of the receiving bank. Returned for United Kingdom
		// domestic transfers.
		// [Optional]
		SortCode string `json:"sort_code,omitempty"`
		// RoutingNumber is the routing number of the receiving bank. Returned for United
		// States domestic transfers.
		// [Optional]
		RoutingNumber string `json:"routing_number,omitempty"`
		// Iban is the International Bank Account Number of the receiving account.
		// [Optional]
		Iban string `json:"iban,omitempty"`
		// SwiftCode is the SWIFT or BIC code of the receiving bank. Returned for
		// international transfers.
		// [Optional]
		SwiftCode string `json:"swift_code,omitempty"`
	}

	// TopUpBankDetails holds the bank details for each available funding rail.
	//
	// Both Domestic and International are optional, and their availability depends on the
	// sub-account's holding currency, jurisdiction, and banking partner. Do not assume that
	// both rails are always available. Both are pointers so an absent rail is distinguishable
	// from a rail returned with empty fields.
	TopUpBankDetails struct {
		// Domestic holds the bank details for the domestic funding rail.
		// [Optional]
		Domestic *TopUpFundingDetails `json:"domestic,omitempty"`
		// International holds the bank details for the international funding rail.
		// [Optional]
		International *TopUpFundingDetails `json:"international,omitempty"`
	}

	// TopUpInstructionsResponse is the response of
	// GET /entities/{entityId}/currency-accounts/{currencyAccountId}/top-up-instructions.
	//
	// It carries the bank details and payment reference used to top up a sub-account.
	TopUpInstructionsResponse struct {
		HttpMetadata common.HttpMetadata
		// CurrencyAccountId is the unique identifier of the sub-account that the instructions
		// apply to.
		// [Required]
		CurrencyAccountId string `json:"currency_account_id,omitempty"`
		// Currency is the currency that funds must be sent in, as a three-letter ISO 4217
		// currency code. This is the sub-account's holding currency, returned as
		// holding_currency by the Retrieve entity balances endpoint.
		// [Required]
		Currency common.Currency `json:"currency,omitempty"`
		// PaymentReference is the reference that must be quoted on the payment. It is how an
		// incoming payment is attributed to the sub-account. A payment sent without this
		// reference may not be credited.
		// [Required]
		PaymentReference string `json:"payment_reference,omitempty"`
		// BankDetails holds the bank details for each available funding rail.
		// [Required]
		BankDetails *TopUpBankDetails `json:"bank_details,omitempty"`
	}
)
