package issuing

import (
	"time"

	"github.com/checkout/checkout-sdk-go/v2/common"
)

type (
	IssuingDisputeResponse struct {
		HttpMetadata         common.HttpMetadata
		Id                   string                     `json:"id,omitempty"`
		Reason               string                     `json:"reason,omitempty"`
		DisputedAmount       *DisputeAmount             `json:"disputed_amount,omitempty"`
		Status               IssuingDisputeStatus       `json:"status,omitempty"`
		StatusReason         IssuingDisputeStatusReason `json:"status_reason,omitempty"`
		TransactionId        string                     `json:"transaction_id,omitempty"`
		PresentmentMessageId string                     `json:"presentment_message_id,omitempty"`
		Merchant             *DisputeMerchant           `json:"merchant,omitempty"`
		CreatedOn            *time.Time                 `json:"created_on,omitempty"`
		ModifiedOn           *time.Time                 `json:"modified_on,omitempty"`
		Chargeback           *DisputeChargeback         `json:"chargeback,omitempty"`
		Representment        *DisputeRepresentment      `json:"representment,omitempty"`
		PreArbitration       *DisputePreArbitration     `json:"pre_arbitration,omitempty"`
		Arbitration          *DisputeArbitration        `json:"arbitration,omitempty"`
		Links                map[string]common.Link     `json:"_links,omitempty"`
	}
)
