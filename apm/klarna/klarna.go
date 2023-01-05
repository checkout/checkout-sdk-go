package klarna

import (
	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/payments"
)

const (
	creditSessionPath = "credit-sessions"
	ordersPath        = "orders"
	capturesPath      = "captures"
	voidsPath         = "voids"
)

type (
	CreditSessionRequest struct {
		PurchaseCountry common.Country           `json:"purchase_country,omitempty"`
		Currency        common.Currency          `json:"currency,omitempty"`
		Locale          string                   `json:"locale,omitempty"`
		Amount          int64                    `json:"amount,omitempty"`
		TaxAmount       int                      `json:"tax_amount,omitempty"`
		Products        []map[string]interface{} `json:"products,omitempty"`
	}

	CreditSessionResponse struct {
		HttpMetadata            common.HttpMetadata `json:"http_metadata,omitempty"`
		SessionId               string              `json:"session_id,omitempty"`
		ClientToken             string              `json:"client_token,omitempty"`
		PaymentMethodCategories []PaymentMethod     `json:"payment_method_categories,omitempty"`
	}

	PaymentMethod struct {
		Identifier string    `json:"identifier,omitempty"`
		Name       string    `json:"name,omitempty"`
		AssetUrls  *AssetUrl `json:"asset_urls,omitempty"`
	}

	AssetUrl struct {
		Descriptive string `json:"descriptive,omitempty"`
		Standard    string `json:"standard,omitempty"`
	}
)

type (
	CreditSession struct {
		HttpMetadata    common.HttpMetadata      `json:"http_metadata,omitempty"`
		ClientToken     string                   `json:"client_token,omitempty"`
		PurchaseCountry string                   `json:"purchase_country,omitempty"`
		Currency        string                   `json:"currency,omitempty"`
		Locale          string                   `json:"locale,omitempty"`
		Amount          int64                    `json:"amount,omitempty"`
		TaxAmount       int                      `json:"tax_amount,omitempty"`
		Products        []map[string]interface{} `json:"products,omitempty"`
	}
)

type (
	OrderCaptureRequest struct {
		Type          payments.SourceType    `json:"type,omitempty"`
		Amount        int64                  `json:"amount,omitempty"`
		Reference     string                 `json:"reference,omitempty"`
		Metadata      map[string]interface{} `json:"metadata,omitempty"`
		Klarna        *Klarna                `json:"klarna,omitempty"`
		ShippingInfo  map[string]interface{} `json:"shipping_info,omitempty"`
		ShippingDelay int                    `json:"shipping_delay,omitempty"`
	}

	CaptureResponse struct {
		HttpMetadata common.HttpMetadata `json:"http_metadata,omitempty"`
		ActionId     string              `json:"action_id,omitempty"`
		Reference    string              `json:"reference,omitempty"`
	}

	Klarna struct {
		Description   string                   `json:"description,omitempty"`
		Products      []map[string]interface{} `json:"products,omitempty"`
		ShippingInfo  []map[string]interface{} `json:"shipping_info,omitempty"`
		ShippingDelay int                      `json:"shipping_delay,omitempty"`
	}
)
