package issuing

import (
	"context"

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

func NewClient(
	configuration *configuration.Configuration,
	apiClient client.HttpClient,
) *Client {
	return &Client{
		configuration: configuration,
		apiClient:     apiClient,
	}
}

func (c *Client) CreateCardholder(request cardholders.CardholderRequest) (*cardholders.CardholderResponse, error) {
	return c.CreateCardholderWithContext(context.Background(), request)
}

func (c *Client) CreateCardholderWithContext(
	ctx context.Context,
	request cardholders.CardholderRequest,
) (*cardholders.CardholderResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response cardholders.CardholderResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(issuingPath, cardholdersPath),
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

func (c *Client) GetCardholder(cardholderId string) (*cardholders.CardholderDetailsResponse, error) {
	return c.GetCardholderWithContext(context.Background(), cardholderId)
}

func (c *Client) GetCardholderWithContext(
	ctx context.Context,
	cardholderId string,
) (*cardholders.CardholderDetailsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response cardholders.CardholderDetailsResponse
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(issuingPath, cardholdersPath, cardholderId),
		auth,
		&response,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetCardholderCards(cardholderId string) (*cardholders.CardholderCardsResponse, error) {
	return c.GetCardholderCardsWithContext(context.Background(), cardholderId)
}

func (c *Client) GetCardholderCardsWithContext(
	ctx context.Context,
	cardholderId string,
) (*cardholders.CardholderCardsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response cardholders.CardholderCardsResponse
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(issuingPath, cardholdersPath, cardholderId, cardsPath),
		auth,
		&response,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) CreateCard(request cards.CardRequest) (*cards.CardResponse, error) {
	return c.CreateCardWithContext(context.Background(), request)
}

func (c *Client) CreateCardWithContext(
	ctx context.Context,
	request cards.CardRequest,
) (*cards.CardResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response cards.CardResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(issuingPath, cardsPath),
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

func (c *Client) GetCardDetails(cardId string) (*cards.CardDetailsResponse, error) {
	return c.GetCardDetailsWithContext(context.Background(), cardId)
}

func (c *Client) GetCardDetailsWithContext(ctx context.Context, cardId string) (*cards.CardDetailsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response cards.CardDetailsResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(issuingPath, cardsPath, cardId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) EnrollThreeDS(
	cardId string,
	enrollmentRequest cards.ThreeDSEnrollmentRequest,
) (*cards.ThreeDSEnrollmentResponse, error) {
	return c.EnrollThreeDSWithContext(context.Background(), cardId, enrollmentRequest)
}

func (c *Client) EnrollThreeDSWithContext(
	ctx context.Context,
	cardId string,
	enrollmentRequest cards.ThreeDSEnrollmentRequest,
) (*cards.ThreeDSEnrollmentResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response cards.ThreeDSEnrollmentResponse
	err = c.apiClient.PostWithContext(
		ctx,
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
	return c.UpdateThreeDSWithContext(context.Background(), cardId, threeDSUpdateRequest)
}

func (c *Client) UpdateThreeDSWithContext(
	ctx context.Context,
	cardId string,
	threeDSUpdateRequest cards.ThreeDSUpdateRequest,
) (*cards.ThreeDSUpdateResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response cards.ThreeDSUpdateResponse
	err = c.apiClient.PatchWithContext(
		ctx,
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
	return c.GetCardThreeDSDetailsWithContext(context.Background(), cardId)
}

func (c *Client) GetCardThreeDSDetailsWithContext(
	ctx context.Context,
	cardId string,
) (*cards.ThreeDSEnrollmentDetailsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response cards.ThreeDSEnrollmentDetailsResponse
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(issuingPath, cardsPath, cardId, threeDSEnrollmentPath),
		auth,
		&response,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) ActivateCard(cardId string) (*common.IdResponse, error) {
	return c.ActivateCardWithContext(context.Background(), cardId)
}

func (c *Client) ActivateCardWithContext(
	ctx context.Context,
	cardId string,
) (*common.IdResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response common.IdResponse
	err = c.apiClient.PostWithContext(
		ctx,
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
	return c.GetCardCredentialsWithContext(context.Background(), cardId, credentialsQuery)
}

func (c *Client) GetCardCredentialsWithContext(
	ctx context.Context,
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
	if err != nil {
		return nil, err
	}
	err = c.apiClient.GetWithContext(ctx, url, auth, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) RevokeCard(
	cardId string,
	revokeCardRequest cards.RevokeCardRequest,
) (*common.IdResponse, error) {
	return c.RevokeCardWithContext(context.Background(), cardId, revokeCardRequest)
}

func (c *Client) RevokeCardWithContext(
	ctx context.Context,
	cardId string,
	revokeCardRequest cards.RevokeCardRequest,
) (*common.IdResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response common.IdResponse
	err = c.apiClient.PostWithContext(
		ctx,
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
	return c.SuspendCardWithContext(context.Background(), cardId)
}

func (c *Client) SuspendCardWithContext(
	ctx context.Context,
	cardId string,
) (*common.IdResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response common.IdResponse
	err = c.apiClient.PostWithContext(
		ctx,
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
	return c.CreateControlWithContext(context.Background(), request)
}

func (c *Client) CreateControlWithContext(
	ctx context.Context,
	request controls.CardControlRequest,
) (*controls.CardControlResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response controls.CardControlResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(issuingPath, controlsPath),
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

func (c *Client) GetCardControls(query controls.CardControlsQuery) (*controls.CardControlsQueryResponse, error) {
	return c.GetCardControlsWithContext(context.Background(), query)
}

func (c *Client) GetCardControlsWithContext(
	ctx context.Context,
	query controls.CardControlsQuery,
) (*controls.CardControlsQueryResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response controls.CardControlsQueryResponse
	url, _ := common.BuildQueryPath(common.BuildPath(issuingPath, controlsPath), query)
	err = c.apiClient.GetWithContext(ctx, url, auth, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetCardControlDetails(controlId string) (*controls.CardControlResponse, error) {
	return c.GetCardControlDetailsWithContext(context.Background(), controlId)
}

func (c *Client) GetCardControlDetailsWithContext(
	ctx context.Context,
	controlId string,
) (*controls.CardControlResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response controls.CardControlResponse
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(issuingPath, controlsPath, controlId),
		auth,
		&response,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) UpdateCardControl(
	controlId string,
	updateCardControlRequest controls.UpdateCardControlRequest,
) (*controls.CardControlResponse, error) {
	return c.UpdateCardControlWithContext(context.Background(), controlId, updateCardControlRequest)
}

func (c *Client) UpdateCardControlWithContext(
	ctx context.Context,
	controlId string,
	updateCardControlRequest controls.UpdateCardControlRequest,
) (*controls.CardControlResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response controls.CardControlResponse
	err = c.apiClient.PutWithContext(
		ctx,
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
	return c.RemoveCardControlWithContext(context.Background(), controlId)
}

func (c *Client) RemoveCardControlWithContext(
	ctx context.Context,
	controlId string,
) (*common.IdResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response common.IdResponse
	err = c.apiClient.DeleteWithContext(
		ctx,
		common.BuildPath(issuingPath, controlsPath, controlId),
		auth,
		&response,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) SimulateAuthorization(
	request testing.CardAuthorizationRequest,
) (*testing.CardAuthorizationResponse, error) {
	return c.SimulateAuthorizationWithContext(context.Background(), request)
}

func (c *Client) SimulateAuthorizationWithContext(
	ctx context.Context,
	request testing.CardAuthorizationRequest,
) (*testing.CardAuthorizationResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response testing.CardAuthorizationResponse
	err = c.apiClient.PostWithContext(
		ctx,
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
	return c.SimulateIncrementWithContext(context.Background(), transactionId, request)
}

func (c *Client) SimulateIncrementWithContext(
	ctx context.Context,
	transactionId string,
	request testing.CardSimulationRequest,
) (*testing.CardSimulationResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response testing.CardSimulationResponse
	err = c.apiClient.PostWithContext(
		ctx,
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
	return c.SimulateClearingWithContext(context.Background(), transactionId, request)
}

func (c *Client) SimulateClearingWithContext(
	ctx context.Context,
	transactionId string,
	request testing.CardSimulationRequest,
) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response common.MetadataResponse
	err = c.apiClient.PostWithContext(
		ctx,
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
	return c.SimulateReversalWithContext(context.Background(), transactionId, request)
}

func (c *Client) SimulateReversalWithContext(
	ctx context.Context,
	transactionId string,
	request testing.CardSimulationRequest,
) (*testing.CardSimulationResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response testing.CardSimulationResponse
	err = c.apiClient.PostWithContext(
		ctx,
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
