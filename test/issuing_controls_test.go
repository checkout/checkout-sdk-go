package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/issuing"
)

func TestCreateControl(t *testing.T) {
	cases := []struct {
		name    string
		checker func(*issuing.CardControlResponse, error)
	}{
		{
			name: "when create a card control and this request is correct then should return a response",
			checker: func(response *issuing.CardControlResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, 201, response.HttpMetadata.StatusCode)
				assert.Equal(t, "Max spend of 500€ per week for restaurants",
					response.VelocityLimitCardControlResponse.Description)
				assert.Equal(t, virtualCardId, response.VelocityLimitCardControlResponse.TargetId)
				assert.NotNil(t, response.VelocityLimitCardControlResponse.Id)
				assert.NotNil(t, response.VelocityLimitCardControlResponse.CreatedDate)
				assert.NotNil(t, response.VelocityLimitCardControlResponse.LastModifiedDate)
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
	t.Skip("Skipping tests because this suite is unstable")
	query := issuing.CardControlsQuery{
		TargetId: virtualCardId,
	}
	cases := []struct {
		name    string
		query   issuing.CardControlsQuery
		checker func(interface{}, error)
	}{
		{
			name:  "when get a card controls and this request is correct then should return a response",
			query: query,
			checker: func(response interface{}, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, 201, response.(*issuing.CardControlsQueryResponse).HttpMetadata.StatusCode)
				assert.Equal(t, "Max spend of 500€ per week for restaurants",
					response.(*issuing.CardControlsQueryResponse).Controls[0].VelocityLimitCardControlResponse.Description)
				assert.Equal(t, virtualCardId, response.(*issuing.CardControlsQueryResponse).Controls[0].VelocityLimitCardControlResponse.TargetId)
				assert.NotNil(t, response.(*issuing.CardControlsQueryResponse).Controls[0].VelocityLimitCardControlResponse.Id)
				assert.NotNil(t, response.(*issuing.CardControlsQueryResponse).Controls[0].VelocityLimitCardControlResponse.CreatedDate)
				assert.NotNil(t, response.(*issuing.CardControlsQueryResponse).Controls[0].VelocityLimitCardControlResponse.LastModifiedDate)
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
				response := data.(*issuing.CardControlsQueryResponse)
				return response.Controls != nil && len(response.Controls) >= 0
			}

			tc.checker(retriable(process, predicate, 5))
		})
	}
}

func TestGetCardControlDetails(t *testing.T) {
	cases := []struct {
		name      string
		controlId string
		checker   func(*issuing.CardControlResponse, error)
	}{
		{
			name:      "when get a card control details and this request is correct then should return a response",
			controlId: cardControlResponse.VelocityLimitCardControlResponse.Id,
			checker: func(response *issuing.CardControlResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, 200, response.HttpMetadata.StatusCode)
				assert.Equal(t, "Max spend of 500€ per week for restaurants",
					response.VelocityLimitCardControlResponse.Description)
				assert.Equal(t, virtualCardId, response.VelocityLimitCardControlResponse.TargetId)
				assert.NotNil(t, response.VelocityLimitCardControlResponse.Id)
				assert.NotNil(t, response.VelocityLimitCardControlResponse.CreatedDate)
				assert.NotNil(t, response.VelocityLimitCardControlResponse.LastModifiedDate)
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
	t.Skip("Skipping tests because this suite is unstable")
	request := issuing.UpdateCardControlRequest{
		Description: "New max spend of 1000€ per month for restaurants",
		VelocityLimit: issuing.VelocityLimit{
			AmountLimit: 1000,
			VelocityWindow: issuing.VelocityWindow{
				Type: issuing.Monthly,
			},
		},
	}
	cases := []struct {
		name      string
		controlId string
		request   issuing.UpdateCardControlRequest
		checker   func(*issuing.CardControlResponse, error)
	}{
		{
			name:      "when update a card control and this request is correct then should return a response",
			controlId: cardControlResponse.VelocityLimitCardControlResponse.Id,
			request:   request,
			checker: func(response *issuing.CardControlResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, 200, response.HttpMetadata.StatusCode)
				assert.Equal(t, "New max spend of 1000€ per month for restaurants",
					response.VelocityLimitCardControlResponse.Description)
				assert.Equal(t, virtualCardId, response.VelocityLimitCardControlResponse.TargetId)
				assert.NotNil(t, response.VelocityLimitCardControlResponse.Id)
				assert.NotNil(t, response.VelocityLimitCardControlResponse.CreatedDate)
				assert.NotNil(t, response.VelocityLimitCardControlResponse.LastModifiedDate)
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
	cases := []struct {
		name      string
		controlId string
		request   issuing.CardControlRequest
		checker   func(*common.IdResponse, error)
	}{
		{
			name:      "when remove a card control and this request is correct then should return a response",
			controlId: cardControlResponse.VelocityLimitCardControlResponse.Id,
			checker: func(response *common.IdResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, 200, response.HttpMetadata.StatusCode)
				assert.Equal(t, cardControlResponse.VelocityLimitCardControlResponse.Id, response.Id)
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
