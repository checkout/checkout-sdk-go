package transfers

import (
    "github.com/checkout/checkout-sdk-go/common"
)

type (
    // InitiateRequest ...
    InitiateRequest struct {
        Source      Source      `json:"source,omitempty"`
        Destination Destination `json:"destination,omitempty"`
        Reference   string      `json:"reference,omitempty"`

        TransferType common.TransferType `json:"transfer_type,omitempty"`
    }

    // Source ...
    Source struct {
        ID     string `json:"id"`
        Amount uint64 `json:"amount"`
    }

    // Destination ...
    Destination struct {
        ID string `json:"id"`
    }

    // Transfer ...
    Transfer struct {
        ID          string                 `json:"id"`
        Reference   string                 `json:"reference"`
        Status      common.TransferStatus  `json:"status,omitempty"`
        Type        common.TransferType    `json:"transfer_type,omitempty"`
        RequestedOn string                 `json:"requested_on,omitempty"`
        ReasonCodes []string               `json:"reason_codes,omitempty"`
        Source      *SourceResponse        `json:"source,omitempty"`
        Destination *DestinationResponse   `json:"destination,omitempty"`
        Links       map[string]common.Link `json:"_links,omitempty"`
    }

    // SourceResponse ...
    SourceResponse struct {
        EntityID string `json:"entity_id"`
        Amount   uint64 `json:"amount"`
        Currency string `json:"currency,omitempty"`
    }

    // DestinationResponse ...
    DestinationResponse struct {
        EntityID string `json:"entity_id"`
    }
)

type (
    // InitiateResponse ...
    InitiateResponse struct {
        ID     string                 `json:"id,omitempty"`
        Status common.TransferStatus  `json:"status,omitempty"`
        Links  map[string]common.Link `json:"_links,omitempty"`
    }
)
