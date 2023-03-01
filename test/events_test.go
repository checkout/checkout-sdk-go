package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	events "github.com/checkout/checkout-sdk-go/events/abc"
)

func TestShouldGetEventTypes(t *testing.T) {
	cases := []struct {
		name    string
		query   events.QueryRetrieveAllEventType
		checker func(response *events.EventTypesResponse, err error)
	}{
		{
			name:  "when get event types by one version then return this version",
			query: events.QueryRetrieveAllEventType{Version: "1.0"},
			checker: func(response *events.EventTypesResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, "1.0", response.EventTypes[0].Version)
			},
		},
		{
			name:  "when get event types by one version then return this version",
			query: events.QueryRetrieveAllEventType{Version: "2.0"},
			checker: func(response *events.EventTypesResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, "2.0", response.EventTypes[0].Version)
			},
		},
		{
			name: "when get event types then return all version",
			checker: func(response *events.EventTypesResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, true, len(response.EventTypes) > 0)
			},
		},
	}

	client := PreviousApi().Events

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.RetrieveAllEventTypesQuery(tc.query))
		})
	}
}

func TestRetrieveEvents(t *testing.T) {
	makeCardPaymentPrevious(t, false, 10)

	cases := []struct {
		name    string
		checker func(interface{}, error)
	}{
		{
			name: "when retrieve events then return events",
			checker: func(response interface{}, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, true, response.(*events.EventsPageResponse).TotalCount > 0)
				assert.Equal(t, 10, response.(*events.EventsPageResponse).Limit)
				assert.Equal(t, true, len(response.(*events.EventsPageResponse).Data) > 0)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(retriableEventsCallback(
				func() (interface{}, error) {
					return PreviousApi().Events.RetrieveEventsQuery(
						events.QueryRetrieveEvents{},
					)
				}),
			)
		})
	}
}

func TestRetrieveEvent(t *testing.T) {
	makeCardPaymentPrevious(t, false, 10)

	allEvents, _ := PreviousApi().Events.RetrieveEvents()

	cases := []struct {
		name    string
		eventId string
		checker func(interface{}, error)
	}{
		{
			name:    "when retrieve events then return events",
			eventId: allEvents.Data[0].Id,
			checker: func(response interface{}, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.(*events.EventResponse).Id)
				assert.NotNil(t, response.(*events.EventResponse).Type)
				assert.NotNil(t, response.(*events.EventResponse).Version)
				assert.NotNil(t, response.(*events.EventResponse).CreatedOn)
			},
		},
	}

	for _, tc := range cases {
		tc.checker(PreviousApi().Events.RetrieveEvent(tc.eventId))
	}
}

func TestRetrieveEventNotification(t *testing.T) {
	t.Skip("Skipping tests because this suite is unstable")
	makeCardPaymentPrevious(t, false, 20)

	allEvents, _ := PreviousApi().Events.RetrieveEventsQuery(events.QueryRetrieveEvents{
		Limit: 20,
	})

	var (
		eventId        = "evt_zzzzzzzzzzzzzzzzzzzzzzzz"
		notificationId = "ntf_zzzzzzzzzzzzzzzzzzzzzzzz"
	)

	for _, event := range allEvents.Data {
		eventNotifications, _ := PreviousApi().Events.RetrieveEvent(event.Id)
		if len(eventNotifications.Notifications) > 0 {
			eventId = event.Id
			notificationId = eventNotifications.Notifications[0].Id
			break
		}
	}

	cases := []struct {
		name           string
		eventId        string
		notificationId string
		checker        func(interface{}, error)
	}{
		{
			name:           "when retrieve event notification then return event notification data",
			eventId:        eventId,
			notificationId: notificationId,
			checker: func(response interface{}, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(
					t,
					response.(*events.EventNotificationResponse).Id,
					notificationId,
				)
				assert.NotNil(t, response.(*events.EventNotificationResponse).Attempts)
			},
		},
	}

	for _, tc := range cases {
		tc.checker(
			retriableEventsCallback(
				func() (interface{}, error) {
					return PreviousApi().Events.RetrieveEventNotification(
						tc.eventId,
						tc.notificationId,
					)
				},
			),
		)
	}
}

func retriableEventsCallback(callback func() (interface{}, error)) (interface{}, error) {
	process := func() (interface{}, error) {
		return callback()
	}

	predicate := func(data interface{}) bool {
		switch data.(type) {
		case *events.EventsPageResponse:
			response := data.(*events.EventsPageResponse)
			return response.Data != nil && len(response.Data) > 0
		case *events.EventResponse:
			response := data.(*events.EventResponse)
			return response.Notifications != nil && len(response.Notifications) > 0
		case *events.EventNotificationResponse:
			response := data.(*events.EventNotificationResponse)
			return response.Attempts != nil && len(response.Attempts) > 0
		default:
			return false
		}
	}

	callbackResponse, err := retriable(process, predicate, 1)

	return callbackResponse, err
}
