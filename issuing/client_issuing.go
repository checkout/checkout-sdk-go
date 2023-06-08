package issuing

import (
	"github.com/checkout/checkout-sdk-go/client"
	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/configuration"
	cardholders "github.com/checkout/checkout-sdk-go/issuing/cardholders"
	cards "github.com/checkout/checkout-sdk-go/issuing/cards"
	controls "github.com/checkout/checkout-sdk-go/issuing/controls"
	testing "github.com/checkout/checkout-sdk-go/issuing/testing"
)

const (
	issuingPath           = "issuing"
	cardholdersPath       = "cardholders"
	cardsPath             = "cards"
	threeDSEnrollmentPath = "3ds-enrollment"
	activatePath          = "activate"
	credentialsPath       = "credentials"
	revokePath            = "revoke"
	suspendPath           = "suspend"
	controlsPath          = "controls"
	simulatePath          = "simulate"
	authorizationsPath    = "authorizations"
	presentmentsPath      = "presentments"
	reversalsPath         = "reversals"
)

type Client struct {
	configuration *configuration.Configuration
	apiClient     client.HttpClient
}

func NewClient(configuration *configuration.Configuration, apiClient client.HttpClient) *Client {
	return &Client{
		configuration: configuration,
		apiClient:     apiClient,
	}
}

func (c *Client) CreateCardholder(request cardholders.CardholderRequest) (*cardholders.CardholderResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response cardholders.CardholderResponse
	err = c.apiClient.Post(common.BuildPath(issuingPath, cardholdersPath), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetCardholder(cardholderId string) (*cardholders.CardholderDetailsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response cardholders.CardholderDetailsResponse
	err = c.apiClient.Get(common.BuildPath(issuingPath, cardholdersPath, cardholderId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetCardholderCards(cardholderId string) (*cardholders.CardholderCardsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response cardholders.CardholderCardsResponse
	err = c.apiClient.Get(common.BuildPath(issuingPath, cardholdersPath, cardholderId, cardsPath), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) CreateCard(request cards.CardRequest) (*cards.CardResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response cards.CardResponse
	err = c.apiClient.Post(common.BuildPath(issuingPath, cardsPath), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetCardDetails(cardId string) (*cards.CardDetailsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response cards.CardDetailsResponse
	err = c.apiClient.Get(common.BuildPath(issuingPath, cardsPath, cardId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) EnrollThreeDS(
	cardId string,
	enrollmentRequest cards.ThreeDSEnrollmentRequest,
) (*cards.ThreeDSEnrollmentResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response cards.ThreeDSEnrollmentResponse
	err = c.apiClient.Post(
		common.BuildPath(issuingPath, cardsPath, cardId, threeDSEnrollmentPath),
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
	threeDSUpdateRequest cards.ThreeDSUpdateRequest,
) (*cards.ThreeDSUpdateResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response cards.ThreeDSUpdateResponse
	err = c.apiClient.Patch(
		common.BuildPath(issuingPath, cardsPath, cardId, threeDSEnrollmentPath),
		auth,
		threeDSUpdateRequest,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetCardThreeDSDetails(cardId string) (*cards.ThreeDSEnrollmentDetailsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response cards.ThreeDSEnrollmentDetailsResponse
	err = c.apiClient.Get(common.BuildPath(issuingPath, cardsPath, cardId, threeDSEnrollmentPath), auth, &response)
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
		common.BuildPath(issuingPath, cardsPath, cardId, activatePath),
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
	credentialsQuery cards.CardCredentialsQuery,
) (*cards.CardCredentialsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response cards.CardCredentialsResponse
	url, err := common.BuildQueryPath(
		common.BuildPath(issuingPath, cardsPath, cardId, credentialsPath),
		credentialsQuery,
	)
	err = c.apiClient.Get(url, auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) RevokeCard(
	cardId string,
	revokeCardRequest cards.RevokeCardRequest,
) (*common.IdResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.IdResponse
	err = c.apiClient.Post(
		common.BuildPath(issuingPath, cardsPath, cardId, revokePath),
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
		common.BuildPath(issuingPath, cardsPath, cardId, suspendPath),
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

func (c *Client) CreateControl(request controls.CardControlRequest) (*controls.CardControlResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response controls.CardControlResponse
	err = c.apiClient.Post(common.BuildPath(issuingPath, controlsPath), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetCardControls(
	query controls.CardControlsQuery,
) (*controls.CardControlsQueryResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response controls.CardControlsQueryResponse
	url, _ := common.BuildQueryPath(common.BuildPath(issuingPath, controlsPath), query)
	err = c.apiClient.Get(url, auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetCardControlDetails(controlId string) (*controls.CardControlResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response controls.CardControlResponse
	err = c.apiClient.Get(common.BuildPath(issuingPath, controlsPath, controlId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UpdateCardControl(
	controlId string,
	updateCardControlRequest controls.UpdateCardControlRequest,
) (*controls.CardControlResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response controls.CardControlResponse
	err = c.apiClient.Put(
		common.BuildPath(issuingPath, controlsPath, controlId),
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
	err = c.apiClient.Delete(common.BuildPath(issuingPath, controlsPath, controlId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) SimulateAuthorization(
	request testing.CardAuthorizationRequest,
) (*testing.CardAuthorizationResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response testing.CardAuthorizationResponse
	err = c.apiClient.Post(
		common.BuildPath(issuingPath, simulatePath, authorizationsPath),
		auth,
		request,
		&response,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) SimulateIncrement(
	transactionId string,
	request testing.CardSimulationRequest,
) (*testing.CardSimulationResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response testing.CardSimulationResponse
	err = c.apiClient.Post(
		common.BuildPath(issuingPath, simulatePath, authorizationsPath, transactionId, authorizationsPath),
		auth,
		request,
		&response,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) SimulateClearing(
	transactionId string,
	request testing.CardSimulationRequest,
) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.Post(
		common.BuildPath(issuingPath, simulatePath, authorizationsPath, transactionId, presentmentsPath),
		auth,
		request,
		&response,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) SimulateReversal(
	transactionId string,
	request testing.CardSimulationRequest,
) (*testing.CardSimulationResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response testing.CardSimulationResponse
	err = c.apiClient.Post(
		common.BuildPath(issuingPath, simulatePath, authorizationsPath, transactionId, reversalsPath),
		auth,
		request,
		&response,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
