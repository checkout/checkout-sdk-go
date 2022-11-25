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
	NativeUi SessionInterface = "native_ui"
	Html     SessionInterface = "html"
)

type ResponseCode string

const (
	Y ResponseCode = "Y"
	N ResponseCode = "N"
	U ResponseCode = "U"
	A ResponseCode = "A"
	C ResponseCode = "C"
	D ResponseCode = "D"
	R ResponseCode = "R"
	I ResponseCode = "I"
)

type (
	SessionResponse struct {
		Accepted *SessionDetails `json:"accepted,omitempty"`
		Created  *SessionDetails `json:"created,omitempty"`
	}

	SessionDetails struct {
		HttpMetadata           common.HttpMetadata
		Id                     string                    `json:"id,omitempty"`
		SessionSecret          string                    `json:"session_secret,omitempty"`
		TransactionId          string                    `json:"transaction_id,omitempty"`
		Scheme                 sources.SessionScheme     `json:"scheme,omitempty"`
		Amount                 int64                     `json:"amount,omitempty"`
		Currency               common.Currency           `json:"currency,omitempty"`
		Completed              bool                      `json:"completed,omitempty"`
		Challenged             bool                      `json:"challenged,omitempty"`
		AuthenticationType     AuthenticationType        `json:"authentication_type,omitempty"`
		AuthenticationCategory Category                  `json:"authentication_category,omitempty"`
		Certificates           *DsPublicKeys             `json:"certificates,omitempty"`
		Status                 SessionStatus             `json:"status,omitempty"`
		StatusReason           StatusReason              `json:"status_reason,omitempty"`
		Approved               bool                      `json:"approved,omitempty"`
		ProtocolVersion        string                    `json:"protocol_version,omitempty"`
		Reference              string                    `json:"reference,omitempty"`
		TransactionType        TransactionType           `json:"transaction_type,omitempty"`
		NextActions            []NextAction              `json:"next_actions,omitempty"`
		Ds                     *Ds                       `json:"ds,omitempty"`
		Acs                    *Acs                      `json:"acs,omitempty"`
		ResponseCode           ResponseCode              `json:"response_code,omitempty"`
		ResponseStatusReason   string                    `json:"response_status_reason,omitempty"`
		Pareq                  string                    `json:"pareq,omitempty"`
		Cryptogram             string                    `json:"cryptogram,omitempty"`
		Eci                    string                    `json:"eci,omitempty"`
		Xid                    string                    `json:"xid,omitempty"`
		CardholderInfo         string                    `json:"cardholder_info,omitempty"`
		Card                   *CardInfo                 `json:"card,omitempty"`
		Recurring              *Recurring                `json:"recurring,omitempty"`
		Installment            *Installment              `json:"installment,omitempty"`
		AuthenticationDate     time.Time                 `json:"authentication_date,omitempty"`
		Exemption              *ThreeDsExemption         `json:"exemption,omitempty"`
		FlowType               common.ThreeDsFlowType    `json:"flow_type,omitempty"`
		ChallengeIndicator     common.ChallengeIndicator `json:"challenge_indicator,omitempty"`
		SchemeInfo             *SchemeInfo               `json:"scheme_info,omitempty"`
		Links                  map[string]common.Link    `json:"_links"`
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
	CardInfo struct {
		InstrumentId string            `json:"instrument_id,omitempty"`
		Fingerprint  string            `json:"fingerprint,omitempty"`
		Metadata     map[string]string `json:"metadata,omitempty"`
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
		SignedContent             string                `json:"signed_content,omitempty"`
		ChallengeMandated         bool                  `json:"challenge_mandated,omitempty"`
		AuthenticationType        string                `json:"authentication_type,omitempty"`
		ChallengeCancelReason     ChallengeCancelReason `json:"challenge_cancel_reason,omitempty"`
		SessionInterface          SessionInterface      `json:"session_interface,omitempty"`
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
		Status string `json:"status,omitempty"`
		Source string `json:"source,omitempty"`
	}

	SchemeInfo struct {
		Name   sources.SessionScheme `json:"name,omitempty"`
		Score  string                `json:"score,omitempty"`
		Avalgo string                `json:"avalgo,omitempty"`
	}
)
