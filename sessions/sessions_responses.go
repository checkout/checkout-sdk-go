package sessions

import (
	"time"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/sessions/channels"
	"github.com/checkout/checkout-sdk-go/sessions/sources"
)

type ChallengeCancelReason string

const (
	CardHolderCancel    ChallengeCancelReason = "cardholder_cancel"
	TransactionTimedOut ChallengeCancelReason = "transaction_timed_out"
	ChallengeTimedOut   ChallengeCancelReason = "challenge_timed_out"
	TransactionError    ChallengeCancelReason = "transaction_error"
	SdkTimedOut         ChallengeCancelReason = "sdk_timed_out"
	Unknown             ChallengeCancelReason = "unknown"
)

type SessionInterface string

const (
	Html     SessionInterface = "html"
	NativeUi SessionInterface = "native_ui"
)

type ResponseCode string

const (
	A ResponseCode = "A"
	C ResponseCode = "C"
	D ResponseCode = "D"
	I ResponseCode = "I"
	N ResponseCode = "N"
	R ResponseCode = "R"
	U ResponseCode = "U"
	Y ResponseCode = "Y"
)

type TrustedBeneficiaryStatusType string

const (
	TrustedBeneficiaryE TrustedBeneficiaryStatusType = "E"
	TrustedBeneficiaryN TrustedBeneficiaryStatusType = "N"
	TrustedBeneficiaryP TrustedBeneficiaryStatusType = "P"
	TrustedBeneficiaryR TrustedBeneficiaryStatusType = "R"
	TrustedBeneficiaryU TrustedBeneficiaryStatusType = "U"
	TrustedBeneficiaryY TrustedBeneficiaryStatusType = "Y"
)

type (
	SessionResponse struct {
		Accepted *SessionDetails `json:"accepted,omitempty"`
		Created  *SessionDetails `json:"created,omitempty"`
	}

	OptimizedProperty struct {
		Field          string `json:"field,omitempty"`
		OriginalValue  string `json:"original_value,omitempty"`
		OptimizedValue string `json:"optimized_value,omitempty"`
	}

	Optimization struct {
		Optimized           bool                `json:"optimized,omitempty"`
		Framework           string              `json:"framework,omitempty"`
		OptimizedProperties []OptimizedProperty `json:"optimized_properties,omitempty"`
	}

	InitialTransaction struct {
		AcsTransactionId        string `json:"acs_transaction_id,omitempty"`
		AuthenticationMethod    string `json:"authentication_method,omitempty"`
		AuthenticationTimestamp string `json:"authentication_timestamp,omitempty"`
		AuthenticationData      string `json:"authentication_data,omitempty"`
		InitialSessionId        string `json:"initial_session_id,omitempty"`
	}

	SessionDetails struct {
		HttpMetadata           common.HttpMetadata
		Id                     string                     `json:"id,omitempty"`
		SessionSecret          string                     `json:"session_secret,omitempty"`
		TransactionId          string                     `json:"transaction_id,omitempty"`
		Scheme                 sources.SessionScheme      `json:"scheme,omitempty"`
		Amount                 int64                      `json:"amount,omitempty"`
		Currency               common.Currency            `json:"currency,omitempty"`
		Completed              bool                       `json:"completed,omitempty"`
		Challenged             bool                       `json:"challenged,omitempty"`
		AuthenticationType     AuthenticationType         `json:"authentication_type,omitempty"`
		AuthenticationCategory Category                   `json:"authentication_category,omitempty"`
		Status                 SessionStatus              `json:"status,omitempty"`
		StatusReason           StatusReason               `json:"status_reason,omitempty"`
		NextActions            []NextAction               `json:"next_actions,omitempty"`
		ProtocolVersion        string                     `json:"protocol_version,omitempty"`
		AccountInfo            *CardholderAccountInfo     `json:"account_info,omitempty"`
		MerchantRiskInfo       *MerchantRiskInfo          `json:"merchant_risk_info,omitempty"`
		Reference              string                     `json:"reference,omitempty"`
		Card                   *CardInfo                  `json:"card,omitempty"`
		Recurring              *Recurring                 `json:"recurring,omitempty"`
		Installment            *Installment               `json:"installment,omitempty"`
		InitialTransaction     *InitialTransaction        `json:"initial_transaction,omitempty"`
		AuthenticationDate     *time.Time                 `json:"authentication_date,omitempty"`
		ChallengeIndicator     *common.ChallengeIndicator `json:"challenge_indicator,omitempty"`
		Optimization           *Optimization              `json:"optimization,omitempty"`
		Certificates           *DsPublicKeys              `json:"certificates,omitempty"`
		Approved               bool                       `json:"approved,omitempty"`
		TransactionType        TransactionType            `json:"transaction_type,omitempty"`
		Ds                     *Ds                        `json:"ds,omitempty"`
		Acs                    *Acs                       `json:"acs,omitempty"`
		ResponseCode           ResponseCode               `json:"response_code,omitempty"`
		ResponseStatusReason   string                     `json:"response_status_reason,omitempty"`
		Pareq                  string                     `json:"pareq,omitempty"`
		Cryptogram             string                     `json:"cryptogram,omitempty"`
		Eci                    string                     `json:"eci,omitempty"`
		Xid                    string                     `json:"xid,omitempty"`
		CardholderInfo         string                     `json:"cardholder_info,omitempty"`
		CustomerIp             string                     `json:"customer_ip,omitempty"`
		Exemption              *ThreeDsExemption          `json:"exemption,omitempty"`
		FlowType               common.ThreeDsFlowType     `json:"flow_type,omitempty"`
		SchemeInfo             *SchemeInfo                `json:"scheme_info,omitempty"`
		Links                  map[string]common.Link     `json:"_links,omitempty"`
	}

	GetSessionResponse struct {
		SessionDetails
	}

	Update3dsMethodCompletionResponse struct {
		HttpMetadata           common.HttpMetadata
		Id                     string                 `json:"id,omitempty"`
		SessionSecret          string                 `json:"session_secret,omitempty"`
		TransactionId          string                 `json:"transaction_id,omitempty"`
		Scheme                 sources.SessionScheme  `json:"scheme,omitempty"`
		Amount                 int64                  `json:"amount,omitempty"`
		Currency               common.Currency        `json:"currency,omitempty"`
		Completed              bool                   `json:"completed,omitempty"`
		Challenged             bool                   `json:"challenged,omitempty"`
		AuthenticationType     AuthenticationType     `json:"authentication_type,omitempty"`
		AuthenticationCategory Category               `json:"authentication_category,omitempty"`
		Certificates           *DsPublicKeys          `json:"certificates,omitempty"`
		Status                 SessionStatus          `json:"status,omitempty"`
		StatusReason           StatusReason           `json:"status_reason,omitempty"`
		Approved               bool                   `json:"approved,omitempty"`
		ProtocolVersion        string                 `json:"protocol_version,omitempty"`
		AccountInfo            *CardholderAccountInfo `json:"account_info,omitempty"`
		MerchantRiskInfo       *MerchantRiskInfo      `json:"merchant_risk_info,omitempty"`
		Reference              string                 `json:"reference,omitempty"`
		TransactionType        TransactionType        `json:"transaction_type,omitempty"`
		NextActions            []NextAction           `json:"next_actions,omitempty"`
		Ds                     *Ds                    `json:"ds,omitempty"`
		Acs                    *Acs                   `json:"acs,omitempty"`
		ResponseCode           ResponseCode           `json:"response_code,omitempty"`
		ResponseStatusReason   string                 `json:"response_status_reason,omitempty"`
		Pareq                  string                 `json:"pareq,omitempty"`
		Cryptogram             string                 `json:"cryptogram,omitempty"`
		Eci                    string                 `json:"eci,omitempty"`
		Xid                    string                 `json:"xid,omitempty"`
		Card                   *CardInfo              `json:"card,omitempty"`
		Links                  map[string]common.Link `json:"_links"`
	}
)

func (r *SessionResponse) MapResponse(response *SessionDetails) {
	switch response.HttpMetadata.StatusCode {
	case 201:
		r.Created = response
	case 202:
		r.Accepted = response
	}
}

type (
	SessionsCardMetadataResponse struct {
		CardType      common.CardType     `json:"card_type"`
		CardCategory  common.CardCategory `json:"card_category"`
		IssuerName    string              `json:"issuer_name"`
		IssuerCountry common.Country      `json:"issuer_country"`
		ProductId     string              `json:"product_id"`
		ProductType   string              `json:"product_type"`
	}

	CardInfo struct {
		InstrumentId string                        `json:"instrument_id,omitempty"`
		Fingerprint  string                        `json:"fingerprint,omitempty"`
		Metadata     *SessionsCardMetadataResponse `json:"metadata,omitempty"`
	}

	DsPublicKeys struct {
		DsPublic string `json:"ds_public,omitempty"`
		CaPublic string `json:"ca_public,omitempty"`
	}

	Ds struct {
		DsId            string `json:"ds_id,omitempty"`
		ReferenceNumber string `json:"reference_number,omitempty"`
		TransactionId   string `json:"transaction_id,omitempty"`
	}

	Acs struct {
		ReferenceNumber           string                `json:"reference_number,omitempty"`
		TransactionId             string                `json:"transaction_id,omitempty"`
		OperatorId                string                `json:"operator_id,omitempty"`
		Url                       string                `json:"url,omitempty"`
		SignedContent             string                `json:"signed_content,omitempty"`
		ChallengeMandated         bool                  `json:"challenge_mandated,omitempty"`
		AuthenticationType        string                `json:"authentication_type,omitempty"`
		ChallengeCancelReason     ChallengeCancelReason `json:"challenge_cancel_reason,omitempty"`
		Interface                 SessionInterface      `json:"interface,omitempty"`
		UiTemplate                channels.UIElements   `json:"ui_template,omitempty"`
		ChallengeCancelReasonCode string                `json:"challenge_cancel_reason_code,omitempty"`
	}

	ThreeDsExemption struct {
		Requested          string              `json:"requested,omitempty"`
		Applied            common.Exemption    `json:"applied,omitempty"`
		Code               string              `json:"code,omitempty"`
		TrustedBeneficiary *TrustedBeneficiary `json:"trusted_beneficiary,omitempty"`
	}

	TrustedBeneficiary struct {
		Status TrustedBeneficiaryStatusType `json:"status,omitempty"`
		Source string                       `json:"source,omitempty"`
	}

	SchemeInfo struct {
		Name   sources.SessionScheme `json:"name,omitempty"`
		Score  string                `json:"score,omitempty"`
		Avalgo string                `json:"avalgo,omitempty"`
	}
)
