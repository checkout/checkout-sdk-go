package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/common"

	controls "github.com/checkout/checkout-sdk-go/issuing/controls"
)

func TestCreateControl(t *testing.T) {
	t.Skip("Avoid creating cards all the time")
	cases := []struct {
		name    string
		checker func(*controls.CardControlResponse, error)
	}{
		{
			name: "when create a card control and this request is correct then should return a response",
			checker: func(response *controls.CardControlResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, 201, response.HttpMetadata.StatusCode)
				assert.Equal(t, "Max spend of 500€ per week for restaurants",
					response.Description)
				assert.Equal(t, virtualCardId, response.TargetId)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.CreatedDate)
				assert.NotNil(t, response.LastModifiedDate)
				assert.NotNil(t, response.Limit.(*controls.VelocityLimit).AmountLimit)
				assert.NotNil(t, response.Limit.(*controls.VelocityLimit).VelocityWindow)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(cardControlRequest(virtualCardId), nil)
		})
	}
}

func TestGetCardControls(t *testing.T) {
	t.Skip("Avoid creating cards all the time")
	query := controls.CardControlsQuery{
		TargetId: virtualCardId,
	}
	cases := []struct {
		name    string
		query   controls.CardControlsQuery
		checker func(interface{}, error)
	}{
		{
			name:  "when get a card controls and this request is correct then should return a response",
			query: query,
			checker: func(response interface{}, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, 200, response.(*controls.CardControlsQueryResponse).HttpMetadata.StatusCode)
				assert.Equal(t, "Max spend of 500€ per week for restaurants",
					response.(*controls.CardControlsQueryResponse).Controls[0].Description)
				assert.Equal(t, virtualCardId, response.(*controls.CardControlsQueryResponse).Controls[0].TargetId)
				assert.NotNil(t, response.(*controls.CardControlsQueryResponse).Controls[0].Id)
				assert.NotNil(t, response.(*controls.CardControlsQueryResponse).Controls[0].CreatedDate)
				assert.NotNil(t, response.(*controls.CardControlsQueryResponse).Controls[0].LastModifiedDate)
			},
		},
	}

	client := buildIssuingClientApi().Issuing

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			process := func() (interface{}, error) {
				return client.GetCardControls(tc.query)
			}
			predicate := func(data interface{}) bool {
				response := data.(*controls.CardControlsQueryResponse)
				return response.Controls != nil && len(response.Controls) >= 0
			}

			tc.checker(retriable(process, predicate, 5))
		})
	}
}

func TestGetCardControlDetails(t *testing.T) {
	t.Skip("Avoid creating cards all the time")
	cases := []struct {
		name      string
		controlId string
		checker   func(*controls.CardControlResponse, error)
	}{
		{
			name:      "when get a card control details and this request is correct then should return a response",
			controlId: cardControlResponse.Id,
			checker: func(response *controls.CardControlResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, 200, response.HttpMetadata.StatusCode)
				assert.Equal(t, "Max spend of 500€ per week for restaurants",
					response.Description)
				assert.Equal(t, virtualCardId, response.TargetId)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.CreatedDate)
				assert.NotNil(t, response.LastModifiedDate)
				assert.NotNil(t, response.Limit.(*controls.VelocityLimit).AmountLimit)
				assert.NotNil(t, response.Limit.(*controls.VelocityLimit).VelocityWindow)
			},
		},
	}

	client := buildIssuingClientApi().Issuing

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetCardControlDetails(tc.controlId))
		})
	}
}

func TestUpdateCardControl(t *testing.T) {
	t.Skip("Avoid creating cards all the time")
	request := controls.UpdateCardControlRequest{
		Description: "New max spend of 1000€ per month for restaurants",
		VelocityLimit: &controls.VelocityLimit{
			AmountLimit: 1000,
			VelocityWindow: controls.VelocityWindow{
				Type: controls.Monthly,
			},
		},
	}
	cases := []struct {
		name      string
		controlId string
		request   controls.UpdateCardControlRequest
		checker   func(*controls.CardControlResponse, error)
	}{
		{
			name:      "when update a card control and this request is correct then should return a response",
			controlId: cardControlResponse.Id,
			request:   request,
			checker: func(response *controls.CardControlResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, 200, response.HttpMetadata.StatusCode)
				assert.Equal(t, "New max spend of 1000€ per month for restaurants",
					response.Description)
				assert.Equal(t, virtualCardId, response.TargetId)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.CreatedDate)
				assert.NotNil(t, response.LastModifiedDate)
			},
		},
	}

	client := buildIssuingClientApi().Issuing

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.UpdateCardControl(tc.controlId, tc.request))
		})
	}
}

func TestRemoveCardControl(t *testing.T) {
	t.Skip("Avoid creating cards all the time")
	cases := []struct {
		name      string
		controlId string
		request   controls.CardControlRequest
		checker   func(*common.IdResponse, error)
	}{
		{
			name:      "when remove a card control and this request is correct then should return a response",
			controlId: cardControlResponse.Id,
			checker: func(response *common.IdResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, 200, response.HttpMetadata.StatusCode)
				assert.Equal(t, cardControlResponse.Id, response.Id)
			},
		},
	}

	client := buildIssuingClientApi().Issuing

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.RemoveCardControl(tc.controlId))
		})
	}
}
