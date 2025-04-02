package sessions

import (
	"time"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/payments"
	"github.com/checkout/checkout-sdk-go/sessions/channels"
	"github.com/checkout/checkout-sdk-go/sessions/completion"
	"github.com/checkout/checkout-sdk-go/sessions/sources"
)

type ShippingIndicator string

const (
	Visa ShippingIndicator = "visa"
)

type (
	ThreeDsRequestorAuthenticationInfo struct {
		ThreeDsReqAuthMethod    payments.ThreeDsReqAuthMethodType `json:"three_ds_req_auth_method,omitempty"`
		ThreeDsReqAuthTimestamp *time.Time                        `json:"three_ds_req_auth_timestamp,omitempty"`
		ThreeDsReqAuthData      string                            `json:"three_ds_req_auth_data,omitempty"`
	}

	CardholderAccountInfo struct {
		AccountInfo                        *payments.AccountInfo
		ThreeDsRequestorAuthenticationInfo *ThreeDsRequestorAuthenticationInfo `json:"three_ds_requestor_authentication_info,omitempty"`
	}

	SessionMarketplaceData struct {
		SubEntityId string `json:"sub_entity_id,omitempty"`
	}

	SessionsBillingDescriptor struct {
		Name string `json:"name,omitempty"`
	}
)

type (
	SessionRequest struct {
		Source                        sources.SessionSource      `json:"source,omitempty"`
		Amount                        int64                      `json:"amount,omitempty"`
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
		Optimization                  *Optimization              `json:"optimization,omitempty"`
		InitialTransaction            *InitialTransaction        `json:"initial_transaction,omitempty"`
	}

	ThreeDsMethodCompletionRequest struct {
		ThreeDsMethodCompletion common.ThreeDsMethodCompletion `json:"three_ds_method_completion,omitempty"`
	}
)

func NewSessionRequest() *SessionRequest {
	return &SessionRequest{
		Source:                 sources.NewSessionCardSource(),
		AuthenticationType:     RegularAuthType,
		AuthenticationCategory: Payment,
		ChallengeIndicator:     common.NoPreference,
		TransactionType:        GoodsService,
	}
}
