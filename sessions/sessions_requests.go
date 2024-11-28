package sessions

import (
	"time"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/sessions/channels"
	"github.com/checkout/checkout-sdk-go/sessions/completion"
	"github.com/checkout/checkout-sdk-go/sessions/sources"
)

type AuthenticationMethod string

const (
	FederatedId              AuthenticationMethod = "federated_id"
	Fido                     AuthenticationMethod = "fido"
	IssuerCredentials        AuthenticationMethod = "issuer_credentials"
	NoAuthentication         AuthenticationMethod = "no_authentication"
	OwnCredentials           AuthenticationMethod = "own_credentials"
	ThirdPartyAuthentication AuthenticationMethod = "third_party_authentication"
)

type ShippingIndicator string

const (
	Visa ShippingIndicator = "visa"
)

type AccountTypeCardProductType string

const (
	Credit        AccountTypeCardProductType = "credit"
	Debit         AccountTypeCardProductType = "debit"
	NotApplicable AccountTypeCardProductType = "not_applicable"
)

type ThreeDsReqAuthMethodType string

const (
	ThreeDsFederatedId                              ThreeDsReqAuthMethodType = "federated_id"
	ThreeDsFidoAuthenticator                        ThreeDsReqAuthMethodType = "fido_authenticator"
	ThreeDsFidoAuthenticatorFidoAssuranceDataSigned ThreeDsReqAuthMethodType = "fido_authenticator_fido_assurance_data_signed"
	ThreeDsIssuerCredentials                        ThreeDsReqAuthMethodType = "issuer_credentials"
	ThreeDsNoAuthenticationOccurred                 ThreeDsReqAuthMethodType = "no_threeds_requestor_authentication_occurred"
	ThreeDsSrcAssuranceData                         ThreeDsReqAuthMethodType = "src_assurance_data"
	ThreeDsOwnCredentials                           ThreeDsReqAuthMethodType = "three3ds_requestor_own_credentials"
	ThreeDsThirdPartyAuthentication                 ThreeDsReqAuthMethodType = "third_party_authentication"
)

type (
	ThreeDsRequestorAuthenticationInfo struct {
		ThreeDsReqAuthMethod    *ThreeDsReqAuthMethodType `json:"three_ds_req_auth_method,omitempty"`
		ThreeDsReqAuthTimestamp *time.Time                `json:"three_ds_req_auth_timestamp,omitempty"`
		ThreeDsReqAuthData      string                    `json:"three_ds_req_auth_data,omitempty"`
	}

	CardholderAccountInfo struct {
		PurchaseCount                      int64                                     `json:"purchase_count,omitempty"`
		AccountAge                         string                                    `json:"account_age,omitempty"`
		AddCardAttempts                    int64                                     `json:"add_card_attempts,omitempty"`
		ShippingAddressAge                 string                                    `json:"shipping_address_age,omitempty"`
		AccountNameMatchesShippingName     bool                                      `json:"account_name_matches_shipping_name,omitempty"`
		SuspiciousAccountActivity          bool                                      `json:"suspicious_account_activity,omitempty"`
		TransactionsToday                  int64                                     `json:"transactions_today,omitempty"`
		AuthenticationMethod               *AuthenticationMethod                     `json:"authentication_method,omitempty"` // Deprecated field
		CardholderAccountAgeIndicator      common.CardholderAccountAgeIndicatorType  `json:"cardholder_account_age_indicator,omitempty"`
		AccountChange                      *time.Time                                `json:"account_change,omitempty"`
		AccountChangeIndicator             common.AccountChangeIndicatorType         `json:"account_change_indicator,omitempty"`
		AccountDate                        *time.Time                                `json:"account_date,omitempty"`
		AccountPasswordChange              string                                    `json:"account_password_change,omitempty"`
		AccountPasswordChangeIndicator     common.AccountPasswordChangeIndicatorType `json:"account_password_change_indicator,omitempty"`
		TransactionsPerYear                int                                       `json:"transactions_per_year,omitempty"`
		PaymentAccountAge                  *time.Time                                `json:"payment_account_age,omitempty"`
		ShippingAddressUsage               *time.Time                                `json:"shipping_address_usage,omitempty"`
		AccountType                        AccountTypeCardProductType                `json:"account_type,omitempty"`
		AccountId                          string                                    `json:"account_id,omitempty"`
		ThreeDsRequestorAuthenticationInfo *ThreeDsRequestorAuthenticationInfo       `json:"three_ds_requestor_authentication_info,omitempty"`
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
