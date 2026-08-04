package sessions

import (
	"time"

	"github.com/checkout/checkout-sdk-go/v2/common"
	"github.com/checkout/checkout-sdk-go/v2/payments"
	"github.com/checkout/checkout-sdk-go/v2/sessions/channels"
	"github.com/checkout/checkout-sdk-go/v2/sessions/completion"
	"github.com/checkout/checkout-sdk-go/v2/sessions/sources"
)

// ShippingIndicator indicates the shipping method chosen for the transaction.
//
// Used by MerchantRiskInfo.ShippingIndicator. Choose the option that accurately describes the
// cardholder's specific transaction.
//
// [Optional]
//
// The constants are prefixed to avoid clashing with the scheme and payment-method constants of the
// same name elsewhere in the SDK.
type ShippingIndicator string

const (
	ShippingBillingAddress           ShippingIndicator = "billing_address"
	ShippingAnotherAddressOnFile     ShippingIndicator = "another_address_on_file"
	ShippingNotOnFile                ShippingIndicator = "not_on_file"
	ShippingStorePickUp              ShippingIndicator = "store_pick_up"
	ShippingDigitalGoods             ShippingIndicator = "digital_goods"
	ShippingTravelAndEventNoShipping ShippingIndicator = "travel_and_event_no_shipping"
	ShippingOther                    ShippingIndicator = "other"
)

type Experience string

const (
	ThreeDsExperience   Experience = "3ds"
	GoogleSpaExperience Experience = "google_spa"
)

type GoogleSpa struct {
	ContinueUrl string `json:"continue_url,omitempty"`
}

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
		ChallengeIndicator            SessionChallengeIndicator  `json:"challenge_indicator,omitempty"`
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
		GoogleSpa                     *GoogleSpa                 `json:"google_spa,omitempty"`
		PreferredExperiences          []Experience               `json:"preferred_experiences,omitempty"`
		DeviceInformation             *DeviceInformation         `json:"device_information,omitempty"`
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
		ChallengeIndicator:     SessionChallengeNoPreference,
		TransactionType:        GoodsService,
	}
}
