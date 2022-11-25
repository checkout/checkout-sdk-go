package sessions

import (
	"github.com/checkout/checkout-sdk-go-beta/common"
	"github.com/checkout/checkout-sdk-go-beta/sessions/channels"
	"github.com/checkout/checkout-sdk-go-beta/sessions/completion"
	"github.com/checkout/checkout-sdk-go-beta/sessions/sources"
)

type AuthenticationMethod string

const (
	NoAuthentication         AuthenticationMethod = "no_authentication"
	OwnCredentials           AuthenticationMethod = "own_credentials"
	FederatedId              AuthenticationMethod = "federated_id"
	IssuerCredentials        AuthenticationMethod = "issuer_credentials"
	ThirdPartyAuthentication AuthenticationMethod = "third_party_authentication"
	Fido                     AuthenticationMethod = "fido"
)

type DeliveryTimeframe string

const (
	ElectronicDelivery DeliveryTimeframe = "electronic_delivery"
	SameDay            DeliveryTimeframe = "same_day"
	Overnight          DeliveryTimeframe = "overnight"
	TwoDayOrMore       DeliveryTimeframe = "two_day_or_more"
)

type ShippingIndicator string

const (
	Visa ShippingIndicator = "visa"
)

type (
	SessionRequest struct {
		Source                        sources.SessionSource      `json:"source,omitempty"`
		Amount                        int                        `json:"amount,omitempty"`
		Currency                      common.Currency            `json:"currency,omitempty"`
		ProcessingChannelId           string                     `json:"processing_channel_id,omitempty"`
		Marketplace                   *SessionMarketplaceData    `json:"marketplace,omitempty"`
		AuthenticationType            AuthenticationType         `json:"authentication_type,omitempty"`
		AuthenticationCategory        Category                   `json:"authentication_category,omitempty"`
		AccountInfo                   *CardholderAccountInfo     `json:"account_info,omitempty"`
		ChallengeIndicator            common.ChallengeIndicator  `json:"challenge_indicator,omitempty"`
		BillingDescriptor             *SessionsBillingDescriptor `json:"billing_descriptor,omitempty"`
		Reference                     string                     `json:"reference,omitempty"`
		MerchantRiskInfo              *MerchantRiskInfo          `json:"merchant_risk_info,omitempty"`
		PriorTransactionReference     string                     `json:"prior_transaction_reference,omitempty"`
		TransactionType               TransactionType            `json:"transaction_type,omitempty"`
		ShippingAddress               *sources.SessionAddress    `json:"shipping_address,omitempty"`
		ShippingAddressMatchesBilling bool                       `json:"shipping_address_matches_billing,omitempty"`
		Completion                    completion.Completion      `json:"completion,omitempty"`
		ChannelData                   channels.Channel           `json:"channel_data,omitempty"`
		Recurring                     *Recurring                 `json:"recurring,omitempty"`
		Installment                   *Installment               `json:"installment,omitempty"`
	}

	ThreeDsMethodCompletionRequest struct {
		ThreeDsMethodCompletion common.ThreeDsMethodCompletion `json:"three_ds_method_completion,omitempty"`
	}
)

type SessionMarketplaceData struct {
	SubEntityId string `json:"sub_entity_id,omitempty"`
}

type CardholderAccountInfo struct {
	PurchaseCount                  int                  `json:"purchase_count,omitempty"`
	AccountAge                     string               `json:"account_age,omitempty"`
	AddCardAttempts                int                  `json:"add_card_attempts,omitempty"`
	ShippingAddressAge             string               `json:"shipping_address_age,omitempty"`
	AccountNameMatchesShippingName bool                 `json:"account_name_matches_shipping_name,omitempty"`
	SuspiciousAccountActivity      bool                 `json:"suspicious_account_activity,omitempty"`
	TransactionsToday              int                  `json:"transactions_today,omitempty"`
	AuthenticationMethod           AuthenticationMethod `json:"authentication_method,omitempty"`
}

type SessionsBillingDescriptor struct {
	Name string `json:"name,omitempty"`
}

type MerchantRiskInfo struct {
	DeliveryEmail     string            `json:"delivery_email,omitempty"`
	DeliveryTimeframe DeliveryTimeframe `json:"delivery_timeframe,omitempty"`
	IsPreorder        bool              `json:"is_preorder,omitempty"`
	IsReorder         bool              `json:"is_reorder,omitempty"`
	ShippingIndicator ShippingIndicator `json:"shipping_indicator,omitempty"`
}
