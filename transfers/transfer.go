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
)

type (
    // InitiateResponse ...
    InitiateResponse struct {
        ID     string                 `json:"id,omitempty"`
        Status common.TransferStatus  `json:"status,omitempty"`
        Links  map[string]common.Link `json:"_links,omitempty"`
    }
)
