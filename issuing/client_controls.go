package issuing

import (
	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/configuration"
)

func (c *Client) CreateControl(request CardControlRequest) (*CardControlResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response CardControlResponse
	err = c.apiClient.Post(common.BuildPath(issuing, controls), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetCardControls(query CardControlsQuery) (*CardControlsQueryResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response CardControlsQueryResponse
	url, _ := common.BuildQueryPath(common.BuildPath(issuing, controls), query)
	err = c.apiClient.Get(url, auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetCardControlDetails(controlId string) (*CardControlResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response CardControlResponse
	err = c.apiClient.Get(common.BuildPath(issuing, controls, controlId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UpdateCardControl(
	controlId string,
	updateCardControlRequest UpdateCardControlRequest,
) (*CardControlResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response CardControlResponse
	err = c.apiClient.Put(
		common.BuildPath(issuing, controls, controlId),
		auth,
		updateCardControlRequest,
		&response,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) RemoveCardControl(controlId string) (*common.IdResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.IdResponse
	err = c.apiClient.Delete(common.BuildPath(issuing, controls, controlId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
