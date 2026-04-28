package issuing

import (
	"context"

	"github.com/checkout/checkout-sdk-go/v2/client"
	"github.com/checkout/checkout-sdk-go/v2/common"
	"github.com/checkout/checkout-sdk-go/v2/configuration"
	cardholders "github.com/checkout/checkout-sdk-go/v2/issuing/cardholders"
	cards "github.com/checkout/checkout-sdk-go/v2/issuing/cards"
	controlgroups "github.com/checkout/checkout-sdk-go/v2/issuing/controlgroups"
	controlprofiles "github.com/checkout/checkout-sdk-go/v2/issuing/controlprofiles"
	controls "github.com/checkout/checkout-sdk-go/v2/issuing/controls"
	digitalcards "github.com/checkout/checkout-sdk-go/v2/issuing/digitalcards"
	disputes "github.com/checkout/checkout-sdk-go/v2/issuing/disputes"
	testing "github.com/checkout/checkout-sdk-go/v2/issuing/testing"
	transactions "github.com/checkout/checkout-sdk-go/v2/issuing/transactions"
)

const (
	issuingPath            = "issuing"
	cardholdersPath        = "cardholders"
	cardsPath              = "cards"
	threeDSEnrollmentPath  = "3ds-enrollment"
	activatePath           = "activate"
	credentialsPath        = "credentials"
	revokePath             = "revoke"
	renewPath              = "renew"
	scheduleRevocationPath = "schedule-revocation"
	suspendPath            = "suspend"
	controlsPath           = "controls"
	controlGroupsPath      = "control-groups"
	controlProfilesPath    = "control-profiles"
	addPath                = "add"
	removePath             = "remove"
	digitalCardsPath       = "digital-cards"
	disputesPath           = "disputes"
	cancelPath             = "cancel"
	escalatePath           = "escalate"
	simulatePath           = "simulate"
	authorizationsPath     = "authorizations"
	presentmentsPath       = "presentments"
	refundsPath            = "refunds"
	reversalsPath          = "reversals"
	oobPath                = "oob"
	authenticationPath     = "authentication"
	transactionsPath       = "transactions"
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

func (c *Client) UpdateCardholder(
	cardholderId string,
	request cardholders.CardholderRequest,
) (*cardholders.CardholderUpdateResponse, error) {
	return c.UpdateCardholderWithContext(context.Background(), cardholderId, request)
}

func (c *Client) UpdateCardholderWithContext(
	ctx context.Context,
	cardholderId string,
	request cardholders.CardholderRequest,
) (*cardholders.CardholderUpdateResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response cardholders.CardholderUpdateResponse
	err = c.apiClient.PatchWithContext(
		ctx,
		common.BuildPath(issuingPath, cardholdersPath, cardholderId),
		auth,
		request,
		&response,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) RenewCard(cardId string, request cards.RenewCardRequest) (*cards.RenewCardResponse, error) {
	return c.RenewCardWithContext(context.Background(), cardId, request)
}

func (c *Client) RenewCardWithContext(
	ctx context.Context,
	cardId string,
	request cards.RenewCardRequest,
) (*cards.RenewCardResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response cards.RenewCardResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(issuingPath, cardsPath, cardId, renewPath),
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

func (c *Client) ScheduleCardRevocation(
	cardId string,
	request cards.ScheduleRevocationRequest,
) (*cards.ScheduleRevocationResponse, error) {
	return c.ScheduleCardRevocationWithContext(context.Background(), cardId, request)
}

func (c *Client) ScheduleCardRevocationWithContext(
	ctx context.Context,
	cardId string,
	request cards.ScheduleRevocationRequest,
) (*cards.ScheduleRevocationResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response cards.ScheduleRevocationResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(issuingPath, cardsPath, cardId, scheduleRevocationPath),
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

func (c *Client) DeleteScheduledRevocation(cardId string) (*cards.ScheduleRevocationResponse, error) {
	return c.DeleteScheduledRevocationWithContext(context.Background(), cardId)
}

func (c *Client) DeleteScheduledRevocationWithContext(
	ctx context.Context,
	cardId string,
) (*cards.ScheduleRevocationResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response cards.ScheduleRevocationResponse
	err = c.apiClient.DeleteWithContext(
		ctx,
		common.BuildPath(issuingPath, cardsPath, cardId, scheduleRevocationPath),
		auth,
		&response,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetControlGroups(
	query controlgroups.ControlGroupsQuery,
) (*controlgroups.ControlGroupsResponse, error) {
	return c.GetControlGroupsWithContext(context.Background(), query)
}

func (c *Client) GetControlGroupsWithContext(
	ctx context.Context,
	query controlgroups.ControlGroupsQuery,
) (*controlgroups.ControlGroupsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response controlgroups.ControlGroupsResponse
	url, err := common.BuildQueryPath(common.BuildPath(issuingPath, controlsPath, controlGroupsPath), query)
	if err != nil {
		return nil, err
	}
	err = c.apiClient.GetWithContext(ctx, url, auth, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) CreateControlGroup(
	request controlgroups.CreateControlGroupRequest,
) (*controlgroups.ControlGroupResponse, error) {
	return c.CreateControlGroupWithContext(context.Background(), request)
}

func (c *Client) CreateControlGroupWithContext(
	ctx context.Context,
	request controlgroups.CreateControlGroupRequest,
) (*controlgroups.ControlGroupResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response controlgroups.ControlGroupResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(issuingPath, controlsPath, controlGroupsPath),
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

func (c *Client) GetControlGroupDetails(controlGroupId string) (*controlgroups.ControlGroupResponse, error) {
	return c.GetControlGroupDetailsWithContext(context.Background(), controlGroupId)
}

func (c *Client) GetControlGroupDetailsWithContext(
	ctx context.Context,
	controlGroupId string,
) (*controlgroups.ControlGroupResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response controlgroups.ControlGroupResponse
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(issuingPath, controlsPath, controlGroupsPath, controlGroupId),
		auth,
		&response,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) RemoveControlGroup(controlGroupId string) (*common.MetadataResponse, error) {
	return c.RemoveControlGroupWithContext(context.Background(), controlGroupId)
}

func (c *Client) RemoveControlGroupWithContext(
	ctx context.Context,
	controlGroupId string,
) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response common.MetadataResponse
	err = c.apiClient.DeleteWithContext(
		ctx,
		common.BuildPath(issuingPath, controlsPath, controlGroupsPath, controlGroupId),
		auth,
		&response,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetAllControlProfiles(
	query controlprofiles.ControlProfilesQuery,
) (*controlprofiles.ControlProfilesResponse, error) {
	return c.GetAllControlProfilesWithContext(context.Background(), query)
}

func (c *Client) GetAllControlProfilesWithContext(
	ctx context.Context,
	query controlprofiles.ControlProfilesQuery,
) (*controlprofiles.ControlProfilesResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response controlprofiles.ControlProfilesResponse
	url, err := common.BuildQueryPath(common.BuildPath(issuingPath, controlsPath, controlProfilesPath), query)
	if err != nil {
		return nil, err
	}
	err = c.apiClient.GetWithContext(ctx, url, auth, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) CreateControlProfile(
	request controlprofiles.ControlProfileRequest,
) (*controlprofiles.ControlProfileResponse, error) {
	return c.CreateControlProfileWithContext(context.Background(), request)
}

func (c *Client) CreateControlProfileWithContext(
	ctx context.Context,
	request controlprofiles.ControlProfileRequest,
) (*controlprofiles.ControlProfileResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response controlprofiles.ControlProfileResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(issuingPath, controlsPath, controlProfilesPath),
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

func (c *Client) GetControlProfileDetails(
	controlProfileId string,
) (*controlprofiles.ControlProfileResponse, error) {
	return c.GetControlProfileDetailsWithContext(context.Background(), controlProfileId)
}

func (c *Client) GetControlProfileDetailsWithContext(
	ctx context.Context,
	controlProfileId string,
) (*controlprofiles.ControlProfileResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response controlprofiles.ControlProfileResponse
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(issuingPath, controlsPath, controlProfilesPath, controlProfileId),
		auth,
		&response,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) UpdateControlProfile(
	controlProfileId string,
	request controlprofiles.ControlProfileRequest,
) (*controlprofiles.ControlProfileResponse, error) {
	return c.UpdateControlProfileWithContext(context.Background(), controlProfileId, request)
}

func (c *Client) UpdateControlProfileWithContext(
	ctx context.Context,
	controlProfileId string,
	request controlprofiles.ControlProfileRequest,
) (*controlprofiles.ControlProfileResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response controlprofiles.ControlProfileResponse
	err = c.apiClient.PatchWithContext(
		ctx,
		common.BuildPath(issuingPath, controlsPath, controlProfilesPath, controlProfileId),
		auth,
		request,
		&response,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) RemoveControlProfile(controlProfileId string) (*common.MetadataResponse, error) {
	return c.RemoveControlProfileWithContext(context.Background(), controlProfileId)
}

func (c *Client) RemoveControlProfileWithContext(
	ctx context.Context,
	controlProfileId string,
) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response common.MetadataResponse
	err = c.apiClient.DeleteWithContext(
		ctx,
		common.BuildPath(issuingPath, controlsPath, controlProfilesPath, controlProfileId),
		auth,
		&response,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) AddTargetToControlProfile(
	controlProfileId string,
	targetId string,
) (*common.MetadataResponse, error) {
	return c.AddTargetToControlProfileWithContext(context.Background(), controlProfileId, targetId)
}

func (c *Client) AddTargetToControlProfileWithContext(
	ctx context.Context,
	controlProfileId string,
	targetId string,
) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response common.MetadataResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(issuingPath, controlsPath, controlProfilesPath, controlProfileId, addPath, targetId),
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

func (c *Client) RemoveTargetFromControlProfile(
	controlProfileId string,
	targetId string,
) (*common.MetadataResponse, error) {
	return c.RemoveTargetFromControlProfileWithContext(context.Background(), controlProfileId, targetId)
}

func (c *Client) RemoveTargetFromControlProfileWithContext(
	ctx context.Context,
	controlProfileId string,
	targetId string,
) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response common.MetadataResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(issuingPath, controlsPath, controlProfilesPath, controlProfileId, removePath, targetId),
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

func (c *Client) GetDigitalCard(digitalCardId string) (*digitalcards.GetDigitalCardResponse, error) {
	return c.GetDigitalCardWithContext(context.Background(), digitalCardId)
}

func (c *Client) GetDigitalCardWithContext(
	ctx context.Context,
	digitalCardId string,
) (*digitalcards.GetDigitalCardResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response digitalcards.GetDigitalCardResponse
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(issuingPath, digitalCardsPath, digitalCardId),
		auth,
		&response,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) CreateDispute(
	request disputes.CreateDisputeRequest,
	idempotencyKey *string,
) (*disputes.IssuingDisputeResponse, error) {
	return c.CreateDisputeWithContext(context.Background(), request, idempotencyKey)
}

func (c *Client) CreateDisputeWithContext(
	ctx context.Context,
	request disputes.CreateDisputeRequest,
	idempotencyKey *string,
) (*disputes.IssuingDisputeResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response disputes.IssuingDisputeResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(issuingPath, disputesPath),
		auth,
		request,
		&response,
		idempotencyKey,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetDispute(disputeId string) (*disputes.IssuingDisputeResponse, error) {
	return c.GetDisputeWithContext(context.Background(), disputeId)
}

func (c *Client) GetDisputeWithContext(
	ctx context.Context,
	disputeId string,
) (*disputes.IssuingDisputeResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response disputes.IssuingDisputeResponse
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(issuingPath, disputesPath, disputeId),
		auth,
		&response,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) CancelDispute(disputeId string, idempotencyKey *string) (*common.MetadataResponse, error) {
	return c.CancelDisputeWithContext(context.Background(), disputeId, idempotencyKey)
}

func (c *Client) CancelDisputeWithContext(
	ctx context.Context,
	disputeId string,
	idempotencyKey *string,
) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response common.MetadataResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(issuingPath, disputesPath, disputeId, cancelPath),
		auth,
		nil,
		&response,
		idempotencyKey,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) EscalateDispute(
	disputeId string,
	request disputes.EscalateDisputeRequest,
	idempotencyKey *string,
) (*common.MetadataResponse, error) {
	return c.EscalateDisputeWithContext(context.Background(), disputeId, request, idempotencyKey)
}

func (c *Client) EscalateDisputeWithContext(
	ctx context.Context,
	disputeId string,
	request disputes.EscalateDisputeRequest,
	idempotencyKey *string,
) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response common.MetadataResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(issuingPath, disputesPath, disputeId, escalatePath),
		auth,
		request,
		&response,
		idempotencyKey,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) SimulateRefund(
	transactionId string,
	request testing.SimulateRefundRequest,
) (*common.MetadataResponse, error) {
	return c.SimulateRefundWithContext(context.Background(), transactionId, request)
}

func (c *Client) SimulateRefundWithContext(
	ctx context.Context,
	transactionId string,
	request testing.SimulateRefundRequest,
) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response common.MetadataResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(issuingPath, simulatePath, authorizationsPath, transactionId, refundsPath),
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

func (c *Client) SimulateOobAuthentication(
	request testing.SimulateOobAuthenticationRequest,
) (*common.MetadataResponse, error) {
	return c.SimulateOobAuthenticationWithContext(context.Background(), request)
}

func (c *Client) SimulateOobAuthenticationWithContext(
	ctx context.Context,
	request testing.SimulateOobAuthenticationRequest,
) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response common.MetadataResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(issuingPath, simulatePath, oobPath, authenticationPath),
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

func (c *Client) GetListTransactions(
	query transactions.TransactionsQuery,
) (*transactions.TransactionsListResponse, error) {
	return c.GetListTransactionsWithContext(context.Background(), query)
}

func (c *Client) GetListTransactionsWithContext(
	ctx context.Context,
	query transactions.TransactionsQuery,
) (*transactions.TransactionsListResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response transactions.TransactionsListResponse
	url, err := common.BuildQueryPath(common.BuildPath(issuingPath, transactionsPath), query)
	if err != nil {
		return nil, err
	}
	err = c.apiClient.GetWithContext(ctx, url, auth, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetSingleTransaction(transactionId string) (*transactions.TransactionResponse, error) {
	return c.GetSingleTransactionWithContext(context.Background(), transactionId)
}

func (c *Client) GetSingleTransactionWithContext(
	ctx context.Context,
	transactionId string,
) (*transactions.TransactionResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response transactions.TransactionResponse
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(issuingPath, transactionsPath, transactionId),
		auth,
		&response,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
