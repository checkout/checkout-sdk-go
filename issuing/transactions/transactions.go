package issuing

import (
	"time"

	"github.com/checkout/checkout-sdk-go/v2/common"
)

type TransactionStatusType string

const (
	TransactionAuthorized TransactionStatusType = "authorized"
	TransactionDeclined   TransactionStatusType = "declined"
	TransactionCanceled   TransactionStatusType = "canceled"
	TransactionCleared    TransactionStatusType = "cleared"
	TransactionRefunded   TransactionStatusType = "refunded"
	TransactionDisputed   TransactionStatusType = "disputed"
)

type WalletType string

const (
	WalletGooglePay              WalletType = "googlepay"
	WalletApplePay               WalletType = "applepay"
	WalletRemoteCommercePrograms WalletType = "remote_commerce_programs"
)

type ReferenceType string

const (
	ReferenceOriginalMit       ReferenceType = "original_mit"
	ReferenceOriginalRecurring ReferenceType = "original_recurring"
)

type TransactionType string

const (
	AccountFunding         TransactionType = "account_funding"
	AccountTransfer        TransactionType = "account_transfer"
	AtmInstallment         TransactionType = "atm_installment"
	BalanceInquiry         TransactionType = "balance_inquiry"
	BillPayment            TransactionType = "bill_payment"
	CashAdvance            TransactionType = "cash_advance"
	Cashback               TransactionType = "cashback"
	CreditAdjustment       TransactionType = "credit_adjustment"
	DebitAdjustment        TransactionType = "debit_adjustment"
	OriginalCredit         TransactionType = "original_credit"
	PaymentAccountInquiry  TransactionType = "payment_account_inquiry"
	Payment                TransactionType = "payment"
	PinChange              TransactionType = "pin_change"
	PinUnblock             TransactionType = "pin_unblock"
	PurchaseAccountInquiry TransactionType = "purchase_account_inquiry"
	Purchase               TransactionType = "purchase"
	QuasiCash              TransactionType = "quasi_cash"
	RemittanceFunding      TransactionType = "remittance_funding"
	RemittancePayment      TransactionType = "remittance_payment"
	Unknown                TransactionType = "unknown"
	Withdrawal             TransactionType = "withdrawal"
	Refund                 TransactionType = "refund"
)

type (
	TransactionsQuery struct {
		Limit        int                   `url:"limit,omitempty"`
		Skip         int                   `url:"skip,omitempty"`
		CardholderId string                `url:"cardholder_id,omitempty"`
		CardId       string                `url:"card_id,omitempty"`
		EntityId     string                `url:"entity_id,omitempty"`
		Status       TransactionStatusType `url:"status,omitempty"`
		From         string                `url:"from,omitempty"`
		To           string                `url:"to,omitempty"`
	}
)

type (
	TransactionClient struct {
		Id string `json:"id,omitempty"`
	}

	TransactionEntity struct {
		Id string `json:"id,omitempty"`
	}

	TransactionCard struct {
		Id      string `json:"id,omitempty"`
		Network string `json:"network,omitempty"`
	}

	TransactionDigitalCard struct {
		Id         string     `json:"id,omitempty"`
		WalletType WalletType `json:"wallet_type,omitempty"`
	}

	TransactionCardholder struct {
		Id string `json:"id,omitempty"`
	}

	TransactionMerchant struct {
		Id           string `json:"id,omitempty"`
		Name         string `json:"name,omitempty"`
		City         string `json:"city,omitempty"`
		State        string `json:"state,omitempty"`
		CountryCode  string `json:"country_code,omitempty"`
		CategoryCode string `json:"category_code,omitempty"`
	}

	TransactionAmountDetail struct {
		Amount   *int64          `json:"amount,omitempty"`
		Currency common.Currency `json:"currency,omitempty"`
	}

	TransactionAmounts struct {
		TotalHeld      *TransactionAmountDetail `json:"total_held,omitempty"`
		TotalAuthorized *TransactionAmountDetail `json:"total_authorized,omitempty"`
		TotalReversed  *TransactionAmountDetail `json:"total_reversed,omitempty"`
		TotalCleared   *TransactionAmountDetail `json:"total_cleared,omitempty"`
		TotalRefunded  *TransactionAmountDetail `json:"total_refunded,omitempty"`
	}

	TransactionReferenceTransaction struct {
		TransactionId string        `json:"transaction_id,omitempty"`
		ReferenceType ReferenceType `json:"reference_type,omitempty"`
	}

	TransactionMessage struct {
		Id             string          `json:"id,omitempty"`
		Initiator      string          `json:"initiator,omitempty"`
		Type           string          `json:"type,omitempty"`
		Result         string          `json:"result,omitempty"`
		IsRelayed      *bool           `json:"is_relayed,omitempty"`
		Indicator      string          `json:"indicator,omitempty"`
		DeclineReason     string          `json:"decline_reason,omitempty"`
		AuthorizationCode string          `json:"authorization_code,omitempty"`
		BillingAmount     *int64          `json:"billing_amount,omitempty"`
		BillingCurrency common.Currency `json:"billing_currency,omitempty"`
		CreatedOn      *time.Time      `json:"created_on,omitempty"`
	}

	TransactionResponse struct {
		HttpMetadata         common.HttpMetadata
		Id                   string                            `json:"id,omitempty"`
		CreatedOn            *time.Time                        `json:"created_on,omitempty"`
		Status               TransactionStatusType             `json:"status,omitempty"`
		TransactionType      TransactionType                   `json:"transaction_type,omitempty"`
		Client               *TransactionClient                `json:"client,omitempty"`
		Entity               *TransactionEntity                `json:"entity,omitempty"`
		Card                 *TransactionCard                  `json:"card,omitempty"`
		DigitalCard          *TransactionDigitalCard           `json:"digital_card,omitempty"`
		Cardholder           *TransactionCardholder            `json:"cardholder,omitempty"`
		Amounts              *TransactionAmounts               `json:"amounts,omitempty"`
		Merchant             *TransactionMerchant              `json:"merchant,omitempty"`
		ReferenceTransaction *TransactionReferenceTransaction  `json:"reference_transaction,omitempty"`
		Messages             []TransactionMessage              `json:"messages,omitempty"`
		Links                map[string]common.Link            `json:"_links,omitempty"`
	}

	TransactionsListResponse struct {
		HttpMetadata common.HttpMetadata
		Limit        *int                  `json:"limit,omitempty"`
		Skip         *int                  `json:"skip,omitempty"`
		TotalCount   *int                  `json:"total_count,omitempty"`
		Data         []TransactionResponse `json:"data,omitempty"`
	}
)
