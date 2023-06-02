package test

import (
	"fmt"
	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	webhooks "github.com/checkout/checkout-sdk-go/webhooks/abc"
)

func TestRetrieveWebhook(t *testing.T) {
	t.Skip("Skipping tests because this suite is unstable ")
	registerWebhook := registerWebHook(t, "https://checkout.com/webhooks")

	cases := []struct {
		name      string
		webhookId string
		checker   func(interface{}, error)
	}{
		{
			name:      "when register webhook then return a webhook",
			webhookId: registerWebhook.Id,
			checker: func(webhookResponse interface{}, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, webhookResponse)
				assert.Equal(t, registerWebhook.Url, webhookResponse.(*webhooks.WebhookResponse).Url)
				assert.Equal(t, registerWebhook.Active, webhookResponse.(*webhooks.WebhookResponse).Active)
				assert.Equal(t, registerWebhook.ContentType, webhookResponse.(*webhooks.WebhookResponse).ContentType)
				assert.Equal(t, registerWebhook.EventTypes, webhookResponse.(*webhooks.WebhookResponse).EventTypes)
			},
		},
		{
			name:      "when register webhook empty then return an error",
			webhookId: "wh_zzzzzzzzzzzzzzzzzzzzzzz",
			checker: func(response interface{}, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(
				retriableWebhookCallback(
					func() (interface{}, error) {
						return PreviousApi().Webhooks.RetrieveWebhook(tc.webhookId)
					},
				),
			)
		})
	}
}

func TestUpdateWebhook(t *testing.T) {
	t.Skip("Skipping tests because this suite is unstable ")
	var (
		webhookUpdateRequest = webhooks.WebhookRequest{
			Url:    "https://example.com/webhooks",
			Active: true,
			Headers: map[string]interface{}{
				"authorization": "1234",
			},
			ContentType: webhooks.Json,
			EventTypes: []string{
				"source_updated",
			},
		}
	)

	registerWebhook := registerWebHook(t, "https://checkout.com/webhooks")

	cases := []struct {
		name                 string
		webhookId            string
		webhookUpdateRequest webhooks.WebhookRequest
		updateWebhook        func() (interface{}, error)
		checker              func(interface{}, error)
	}{
		{
			name:                 "when update a webhook then return this webhook updated",
			webhookId:            registerWebhook.Id,
			webhookUpdateRequest: webhookUpdateRequest,
			checker: func(webhookUpdateResponse interface{}, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, webhookUpdateResponse)
				assert.Equal(t, webhookUpdateRequest.Url, webhookUpdateResponse.(*webhooks.WebhookResponse).Url)
				assert.Equal(t, webhookUpdateRequest.Active, webhookUpdateResponse.(*webhooks.WebhookResponse).Active)
				assert.Equal(t, webhookUpdateRequest.ContentType, webhookUpdateResponse.(*webhooks.WebhookResponse).ContentType)
				assert.Equal(t, webhookUpdateRequest.EventTypes, webhookUpdateResponse.(*webhooks.WebhookResponse).EventTypes)
			},
		},
		{
			name:                 "when update a not found webhook then return 404",
			webhookId:            "wh_zzzzzzzzzzzzzzzzzzzzzzz",
			webhookUpdateRequest: webhookUpdateRequest,
			checker: func(response interface{}, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
		{
			name:                 "when update a webhook with empty request then return 422",
			webhookId:            registerWebhook.Id,
			webhookUpdateRequest: webhooks.WebhookRequest{},
			checker: func(response interface{}, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(retriableWebhookCallback(
				func() (interface{}, error) {
					return PreviousApi().Webhooks.UpdateWebhook(tc.webhookId, tc.webhookUpdateRequest)
				},
			))
		})
	}
}

func TestPatchWebhook(t *testing.T) {
	t.Skip("Skipping tests because this suite is unstable ")
	var (
		webhookUpdateRequest = webhooks.WebhookRequest{
			Url:    "https://example.com/webhooks",
			Active: true,
			Headers: map[string]interface{}{
				"authorization": "1234",
			},
			ContentType: webhooks.Json,
			EventTypes: []string{
				"source_updated",
			},
		}
	)

	registerWebhook := registerWebHook(t, "https://checkout.com/webhooks")

	cases := []struct {
		name                 string
		webhookId            string
		webhookUpdateRequest webhooks.WebhookRequest
		updateWebhook        func() (interface{}, error)
		checker              func(interface{}, error)
	}{
		{
			name:                 "when patch a webhook then return this webhook updated",
			webhookId:            registerWebhook.Id,
			webhookUpdateRequest: webhookUpdateRequest,
			checker: func(webhookUpdateResponse interface{}, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, webhookUpdateResponse)
				assert.Equal(t, webhookUpdateRequest.Url, webhookUpdateResponse.(*webhooks.WebhookResponse).Url)
				assert.Equal(t, webhookUpdateRequest.Active, webhookUpdateResponse.(*webhooks.WebhookResponse).Active)
				assert.Equal(t, webhookUpdateRequest.ContentType, webhookUpdateResponse.(*webhooks.WebhookResponse).ContentType)
				assert.Equal(t, webhookUpdateRequest.EventTypes, webhookUpdateResponse.(*webhooks.WebhookResponse).EventTypes)
			},
		},
		{
			name:                 "when patch a not found webhook then return an error",
			webhookId:            "wh_zzzzzzzzzzzzzzzzzzzzzzz",
			webhookUpdateRequest: webhookUpdateRequest,
			checker: func(response interface{}, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(retriableWebhookCallback(
				func() (interface{}, error) {
					return PreviousApi().Webhooks.PartiallyUpdateWebhook(tc.webhookId, tc.webhookUpdateRequest)
				},
			))
		})
	}
}

func TestRemoveWebhooks(t *testing.T) {
	t.Skip("Skipping tests because this suite is unstable ")
	var (
		eventTypes = []string{"payment_approved", "payment_pending",
			"payment_declined", "payment_expired", "payment_canceled", "payment_voided", "payment_void_declined",
			"payment_captured", "payment_capture_declined", "payment_capture_pending", "payment_refunded",
			"payment_refund_declined", "payment_refund_pending"}

		webhookRequest = webhooks.WebhookRequest{
			Url:    "https://checkout.com/webhooks",
			Active: true,
			Headers: map[string]interface{}{
				"authorization": "1234",
			},
			ContentType: webhooks.Json,
			EventTypes:  eventTypes,
		}

		webhookResponse, _ = PreviousApi().Webhooks.RegisterWebhook(webhookRequest)
	)

	cases := []struct {
		name      string
		webhookId string
		checker   func(*common.MetadataResponse, error)
	}{
		{
			name:      "when delete webhook then return 200",
			webhookId: webhookResponse.Id,
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
			},
		},
		{
			name:      "when delete not found webhook then return an error",
			webhookId: "wh_zzzzzzzzzzzzzzzzzzzzzzz",
			checker: func(response *common.MetadataResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := PreviousApi().Webhooks

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.RemoveWebhook(tc.webhookId))
		})
	}
}

func registerWebHook(t *testing.T, url string) *webhooks.WebhookResponse {
	var (
		eventTypes = []string{"payment_approved", "payment_canceled", "payment_capture_declined",
			"payment_capture_pending", "payment_captured", "payment_declined", "payment_expired",
			"payment_pending", "payment_refund_declined", "payment_refund_pending", "payment_refunded",
			"payment_void_declined", "payment_voided"}
	)

	webhookRequest := webhooks.WebhookRequest{
		Url:    url,
		Active: true,
		Headers: map[string]interface{}{
			"authorization": "1234",
		},
		ContentType: webhooks.Json,
		EventTypes:  eventTypes,
	}

	webhookResponse, err := PreviousApi().Webhooks.RegisterWebhook(webhookRequest)
	assert.Nil(t, err)
	assert.NotNil(t, webhookResponse)
	assert.Equal(t, url, webhookResponse.Url)
	assert.Equal(t, webhooks.Json, webhookResponse.ContentType)
	assert.Equal(t, eventTypes, webhookResponse.EventTypes)

	t.Cleanup(func() {
		if _, err = PreviousApi().Webhooks.RemoveWebhook(webhookResponse.Id); err != nil {
			assert.Fail(t, fmt.Sprintf("Failed to remove webhook: %s", err.Error()))
		}
	})

	return webhookResponse
}

func retriableWebhookCallback(callback func() (interface{}, error)) (interface{}, error) {
	processWebhook := func() (interface{}, error) {
		return callback()
	}

	predicateWebhook := func(data interface{}) bool {
		response := data.(*webhooks.WebhookResponse)
		return response.EventTypes != nil && len(response.EventTypes) >= 0
	}

	callbackResponse, err := retriable(processWebhook, predicateWebhook, 1)

	return callbackResponse, err
}
