package webhooks

import (
	"github.com/shiuh-yaw-cko/checkout"
	"github.com/shiuh-yaw-cko/checkout/common"
)

type (
	// Request -
	Request struct {
		*Webhook
	}
	// Webhook ...
	Webhook struct {
		URL         string                    `json:"url,omitempty"`
		Active      *bool                     `json:"active,omitempty"`
		Headers     *Headers                  `json:"headers,omitempty"`
		ContentType common.WebhookContentType `json:"content_type,omitempty"`
		EventTypes  []string                  `json:"event_types,omitempty"`
	}
	// Headers ...
	Headers struct {
		Authorization string `json:"Authorization,omitempty"`
	}
)

type (
	// Response -
	Response struct {
		StatusResponse     *checkout.StatusResponse `json:"api_response,omitempty"`
		ConfiguredWebhooks []WebhookResponse        `json:"webhooks,omitempty"`
		Webhook            *WebhookResponse         `json:"webhook,omitempty"`
	}

	// WebhookResponse -
	WebhookResponse struct {
		ID          string                    `json:"id,omitempty"`
		URL         string                    `json:"url,omitempty"`
		Active      *bool                     `json:"active,omitempty"`
		Headers     *Headers                  `json:"headers,omitempty"`
		ContentType common.WebhookContentType `json:"content_type,omitempty"`
		EventTypes  []string                  `json:"event_types,omitempty"`
		Links       map[string]common.Link    `json:"_links,omitempty"`
		Version     string                    `json:"version,omitempty"`
	}
)
