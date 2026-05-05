package issuing

import "github.com/checkout/checkout-sdk-go/v2/common"

type TransactionType string

const (
	Purchase TransactionType = "purchase"
)

type AuthorizationType string

const (
	Authorization    AuthorizationType = "authorization"
	PreAuthorization AuthorizationType = "pre_authorization"
)

type (
	CardSimulation struct {
		Id          string `json:"id,omitempty"`
		ExpiryMonth int    `json:"expiry_month,omitempty"`
		ExpiryYear  int    `json:"expiry_year,omitempty"`
	}

	SimulationMerchant struct {
		CategoryCode string `json:"category_code,omitempty"`
	}

	TransactionSimulation struct {
		Type              TransactionType   `json:"type,omitempty"`
		Amount            int               `json:"amount,omitempty"`
		Currency          common.Currency   `json:"currency,omitempty"`
		Merchant          *SimulationMerchant `json:"merchant,omitempty"`
		AuthorizationType AuthorizationType `json:"authorization_type,omitempty"`
	}

	CardAuthorizationRequest struct {
		Card        CardSimulation        `json:"card,omitempty"`
		Transaction TransactionSimulation `json:"transaction,omitempty"`
	}

	CardSimulationRequest struct {
		Amount int `json:"amount,omitempty"`
	}

	SimulateRefundRequest struct {
		Amount int `json:"amount,omitempty"`
	}

	OobTransactionDetails struct {
		LastFour         string          `json:"last_four,omitempty"`
		MerchantName     string          `json:"merchant_name,omitempty"`
		PurchaseAmount   float64         `json:"purchase_amount,omitempty"`
		PurchaseCurrency common.Currency `json:"purchase_currency,omitempty"`
	}

	SimulateOobAuthenticationRequest struct {
		CardId             string                 `json:"card_id,omitempty"`
		TransactionDetails *OobTransactionDetails `json:"transaction_details,omitempty"`
	}
)
