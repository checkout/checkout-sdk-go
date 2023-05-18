package issuing

import (
	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/configuration"
)

func (c *Client) CreateCardholder(request CardholderRequest) (*CardholderResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response CardholderResponse
	err = c.apiClient.Post(common.BuildPath(issuing, cardholders), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetCardholder(cardholderId string) (*CardholderDetailsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response CardholderDetailsResponse
	err = c.apiClient.Get(common.BuildPath(issuing, cardholders, cardholderId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetCardholderCards(cardholderId string) (*CardholderCardsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response CardholderCardsResponse
	err = c.apiClient.Get(common.BuildPath(issuing, cardholders, cardholderId, cards), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
