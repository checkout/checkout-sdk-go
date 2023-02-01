package abc

import (
	"encoding/json"
	"github.com/checkout/checkout-sdk-go/common"
)

const (
	webhooks = "webhooks"
)

type ContentType string

const (
	Json ContentType = "json"
	Xml  ContentType = "xml"
)

type (
	WebhookRequest struct {
		Url         string                 `json:"url"`
		Active      bool                   `json:"active,omitempty" default:"true"`
		Headers     map[string]interface{} `json:"headers,omitempty"`
		ContentType ContentType            `json:"content_type,omitempty" default:"json"`
		EventTypes  []string               `json:"event_types"`
	}
)

type (
	WebhookResponse struct {
		HttpResponse common.HttpMetadata
		Id           string                 `json:"id,omitempty"`
		Url          string                 `json:"url,omitempty"`
		Active       bool                   `json:"active,omitempty"`
		Headers      interface{}            `json:"headers,omitempty"`
		ContentType  ContentType            `json:"content_type,omitempty"`
		EventTypes   []string               `json:"event_types,omitempty"`
		Links        map[string]common.Link `json:"_links" json:"links,omitempty"`
	}

	WebhooksResponse struct {
		HttpResponse common.HttpMetadata
		WebhookArray []WebhookResponse
	}
)

func (e *WebhooksResponse) UnmarshalJSON(data []byte) error {
	var webhookArray []WebhookResponse
	if err := json.Unmarshal(data, &webhookArray); err != nil {
		return err
	}
	e.WebhookArray = webhookArray
	return nil
}
