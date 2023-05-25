package issuing

import (
	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/configuration"
)

func (c *Client) CreateCard(request CardRequest) (*CardResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response CardResponse
	err = c.apiClient.Post(common.BuildPath(issuing, cards), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetCardDetails(cardId string) (*CardDetailsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response CardDetailsResponse
	err = c.apiClient.Get(common.BuildPath(issuing, cards, cardId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) EnrollThreeDS(
	cardId string,
	enrollmentRequest ThreeDSEnrollmentRequest,
) (*ThreeDSEnrollmentResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response ThreeDSEnrollmentResponse
	err = c.apiClient.Post(
		common.BuildPath(issuing, cards, cardId, threeDSEnrollment),
		auth,
		enrollmentRequest,
		&response,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UpdateThreeDS(
	cardId string,
	threeDSUpdateRequest ThreeDSUpdateRequest,
) (*ThreeDSUpdateResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response ThreeDSUpdateResponse
	err = c.apiClient.Patch(
		common.BuildPath(issuing, cards, cardId, threeDSEnrollment),
		auth,
		threeDSUpdateRequest,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetCardThreeDSDetails(cardId string) (*ThreeDSEnrollmentDetailsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response ThreeDSEnrollmentDetailsResponse
	err = c.apiClient.Get(common.BuildPath(issuing, cards, cardId, threeDSEnrollment), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) ActivateCard(cardId string) (*common.IdResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.IdResponse
	err = c.apiClient.Post(
		common.BuildPath(issuing, cards, cardId, activate),
		auth,
		nil,
		&response,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetCardCredentials(
	cardId string,
	credentialsQuery CardCredentialsQuery,
) (*CardCredentialsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response CardCredentialsResponse
	url, err := common.BuildQueryPath(common.BuildPath(issuing, cards, cardId, credentials), credentialsQuery)
	err = c.apiClient.Get(url, auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) RevokeCard(
	cardId string,
	revokeCardRequest RevokeCardRequest,
) (*common.IdResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.IdResponse
	err = c.apiClient.Post(
		common.BuildPath(issuing, cards, cardId, revoke),
		auth,
		revokeCardRequest,
		&response,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) SuspendCard(cardId string) (*common.IdResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.IdResponse
	err = c.apiClient.Post(
		common.BuildPath(issuing, cards, cardId, suspend),
		auth,
		nil,
		&response,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
