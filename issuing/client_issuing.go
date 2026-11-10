package issuing

import (
	"context"

	"github.com/checkout/checkout-sdk-go/v3/client"
	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/configuration"
	cardholders "github.com/checkout/checkout-sdk-go/v3/issuing/cardholders"
	cards "github.com/checkout/checkout-sdk-go/v3/issuing/cards"
	controlgroups "github.com/checkout/checkout-sdk-go/v3/issuing/controlgroups"
	controlprofiles "github.com/checkout/checkout-sdk-go/v3/issuing/controlprofiles"
	controls "github.com/checkout/checkout-sdk-go/v3/issuing/controls"
	digitalcards "github.com/checkout/checkout-sdk-go/v3/issuing/digitalcards"
	disputes "github.com/checkout/checkout-sdk-go/v3/issuing/disputes"
	testing "github.com/checkout/checkout-sdk-go/v3/issuing/testing"
	transactions "github.com/checkout/checkout-sdk-go/v3/issuing/transactions"
)

const (
	issuingPath           = "issuing"
	cardholdersPath       = "cardholders"
	cardsPath             = "cards"
	threeDSEnrollmentPath = "3ds-enrollment"
	activatePath          = "activate"
	credentialsPath       = "credentials"
	revokePath            = "revoke"
	renewPath             = "renew"
	suspendPath           = "suspend"
	controlsPath          = "controls"
	controlGroupsPath     = "control-groups"
	controlProfilesPath   = "control-profiles"
	addPath               = "add"
	removePath            = "remove"
	digitalCardsPath      = "digital-cards"
	disputesPath          = "disputes"
	cancelPath            = "cancel"
	escalatePath          = "escalate"
	amendPath             = "amend"
	simulatePath          = "simulate"
	authorizationsPath    = "authorizations"
	presentmentsPath      = "presentments"
	refundsPath           = "refunds"
	reversalsPath         = "reversals"
	oobPath               = "oob"
	authenticationPath    = "authentication"
	transactionsPath      = "transactions"
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

// GetCardholder sends GET /issuing/cardholders/{cardholderId}: get cardholder details.
//
// cardholderId is the {cardholderId} path value: the cardholder's unique identifier.
//
// It returns the cardholders.CardholderDetailsResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." cardholderId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) GetCardholder(cardholderId string) (*cardholders.CardholderDetailsResponse, error) {
	return c.GetCardholderWithContext(context.Background(), cardholderId)
}

// GetCardholderWithContext is GetCardholder with a caller-supplied context.
func (c *Client) GetCardholderWithContext(
	ctx context.Context,
	cardholderId string,
) (*cardholders.CardholderDetailsResponse, error) {
	cardholderId, err := common.EncodePathSegment(cardholderId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response cardholders.CardholderDetailsResponse
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(issuingPath, cardholdersPath, cardholderId),
		auth,
		nil,
		&response)

	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetCardholderCards sends GET /issuing/cardholders/{cardholderId}/cards: get a cardholder's cards.
//
// cardholderId is the {cardholderId} path value: the cardholder's unique identifier.
//
// It returns the cardholders.CardholderCardsResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." cardholderId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) GetCardholderCards(cardholderId string) (*cardholders.CardholderCardsResponse, error) {
	return c.GetCardholderCardsWithContext(context.Background(), cardholderId)
}

// GetCardholderCardsWithContext is GetCardholderCards with a caller-supplied context.
func (c *Client) GetCardholderCardsWithContext(
	ctx context.Context,
	cardholderId string,
) (*cardholders.CardholderCardsResponse, error) {
	cardholderId, err := common.EncodePathSegment(cardholderId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response cardholders.CardholderCardsResponse
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(issuingPath, cardholdersPath, cardholderId, cardsPath),
		auth,
		nil,
		&response)

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

// GetCardDetails sends GET /issuing/cards/{cardId}: get card details.
//
// cardId is the {cardId} path value: the card's unique identifier.
//
// It returns the cards.CardDetailsResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." cardId returns a
// CheckoutArgumentError and sends no request.
func (c *Client) GetCardDetails(cardId string) (*cards.CardDetailsResponse, error) {
	return c.GetCardDetailsWithContext(context.Background(), cardId)
}

// GetCardDetailsWithContext is GetCardDetails with a caller-supplied context.
func (c *Client) GetCardDetailsWithContext(ctx context.Context, cardId string) (*cards.CardDetailsResponse, error) {
	cardId, err := common.EncodePathSegment(cardId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response cards.CardDetailsResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(issuingPath, cardsPath, cardId), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// EnrollThreeDS sends POST /issuing/cards/{cardId}/3ds-enrollment: enroll a card in 3DS.
//
// cardId is the {cardId} path value: the card's unique identifier. enrollmentRequest is sent as the
// JSON body.
//
// It returns the cards.ThreeDSEnrollmentResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." cardId returns a
// CheckoutArgumentError and sends no request.
func (c *Client) EnrollThreeDS(
	cardId string,
	enrollmentRequest cards.ThreeDSEnrollmentRequest,
) (*cards.ThreeDSEnrollmentResponse, error) {
	return c.EnrollThreeDSWithContext(context.Background(), cardId, enrollmentRequest)
}

// EnrollThreeDSWithContext is EnrollThreeDS with a caller-supplied context.
func (c *Client) EnrollThreeDSWithContext(
	ctx context.Context,
	cardId string,
	enrollmentRequest cards.ThreeDSEnrollmentRequest,
) (*cards.ThreeDSEnrollmentResponse, error) {
	cardId, err := common.EncodePathSegment(cardId)
	if err != nil {
		return nil, err
	}

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

// UpdateThreeDS sends PATCH /issuing/cards/{cardId}/3ds-enrollment: update a card's 3DS details.
//
// cardId is the {cardId} path value: the card's unique identifier. threeDSUpdateRequest is sent as
// the JSON body.
//
// It returns the cards.ThreeDSUpdateResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." cardId returns a
// CheckoutArgumentError and sends no request.
func (c *Client) UpdateThreeDS(
	cardId string,
	threeDSUpdateRequest cards.ThreeDSUpdateRequest,
) (*cards.ThreeDSUpdateResponse, error) {
	return c.UpdateThreeDSWithContext(context.Background(), cardId, threeDSUpdateRequest)
}

// UpdateThreeDSWithContext is UpdateThreeDS with a caller-supplied context.
func (c *Client) UpdateThreeDSWithContext(
	ctx context.Context,
	cardId string,
	threeDSUpdateRequest cards.ThreeDSUpdateRequest,
) (*cards.ThreeDSUpdateResponse, error) {
	cardId, err := common.EncodePathSegment(cardId)
	if err != nil {
		return nil, err
	}

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

// GetCardThreeDSDetails sends GET /issuing/cards/{cardId}/3ds-enrollment: get a card's 3DS
// enrollment details.
//
// cardId is the {cardId} path value: the card's unique identifier.
//
// It returns the cards.ThreeDSEnrollmentDetailsResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." cardId returns a
// CheckoutArgumentError and sends no request.
func (c *Client) GetCardThreeDSDetails(cardId string) (*cards.ThreeDSEnrollmentDetailsResponse, error) {
	return c.GetCardThreeDSDetailsWithContext(context.Background(), cardId)
}

// GetCardThreeDSDetailsWithContext is GetCardThreeDSDetails with a caller-supplied context.
func (c *Client) GetCardThreeDSDetailsWithContext(
	ctx context.Context,
	cardId string,
) (*cards.ThreeDSEnrollmentDetailsResponse, error) {
	cardId, err := common.EncodePathSegment(cardId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response cards.ThreeDSEnrollmentDetailsResponse
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(issuingPath, cardsPath, cardId, threeDSEnrollmentPath),
		auth,
		nil,
		&response)

	if err != nil {
		return nil, err
	}
	return &response, nil
}

// ActivateCard sends POST /issuing/cards/{cardId}/activate: activate a card.
//
// cardId is the {cardId} path value: the card's unique identifier.
//
// It returns the cards.ActivateCardResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." cardId returns a
// CheckoutArgumentError and sends no request.
func (c *Client) ActivateCard(cardId string) (*cards.ActivateCardResponse, error) {
	return c.ActivateCardWithContext(context.Background(), cardId)
}

// ActivateCardWithContext is ActivateCard with a caller-supplied context.
func (c *Client) ActivateCardWithContext(
	ctx context.Context,
	cardId string,
) (*cards.ActivateCardResponse, error) {
	cardId, err := common.EncodePathSegment(cardId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response cards.ActivateCardResponse
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

// GetCardCredentials sends GET /issuing/cards/{cardId}/credentials: get the card credentials.
//
// cardId is the {cardId} path value: the card's unique identifier. credentialsQuery is sent as the
// query string.
//
// It returns the cards.CardCredentialsResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." cardId returns a
// CheckoutArgumentError and sends no request.
func (c *Client) GetCardCredentials(
	cardId string,
	credentialsQuery cards.CardCredentialsQuery,
) (*cards.CardCredentialsResponse, error) {
	return c.GetCardCredentialsWithContext(context.Background(), cardId, credentialsQuery)
}

// GetCardCredentialsWithContext is GetCardCredentials with a caller-supplied context.
func (c *Client) GetCardCredentialsWithContext(
	ctx context.Context,
	cardId string,
	credentialsQuery cards.CardCredentialsQuery,
) (*cards.CardCredentialsResponse, error) {
	cardId, err := common.EncodePathSegment(cardId)
	if err != nil {
		return nil, err
	}

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
	err = c.apiClient.GetWithContext(ctx, url, auth, nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// RevokeCard revokes a card permanently; a revoked card cannot be reactivated.
// The request body is optional: pass cards.RevokeCardRequest{} to send no reason, which sends {}.
//
// cardId is the {cardId} path value: the card's unique identifier. revokeCardRequest is sent as the
// JSON body.
//
// It returns the cards.RevokeCardResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." cardId returns a
// CheckoutArgumentError and sends no request.
func (c *Client) RevokeCard(
	cardId string,
	revokeCardRequest cards.RevokeCardRequest,
) (*cards.RevokeCardResponse, error) {
	return c.RevokeCardWithContext(context.Background(), cardId, revokeCardRequest)
}

// RevokeCardWithContext is the context-aware variant of RevokeCard. The request is optional: a zero
// cards.RevokeCardRequest{} sends {}.
func (c *Client) RevokeCardWithContext(
	ctx context.Context,
	cardId string,
	revokeCardRequest cards.RevokeCardRequest,
) (*cards.RevokeCardResponse, error) {
	cardId, err := common.EncodePathSegment(cardId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response cards.RevokeCardResponse
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

// SuspendCard suspends a card; a suspended card declines authorizations until it is reactivated.
// The request body is optional: pass cards.SuspendCardRequest{} to send no reason, which sends {}.
//
// cardId is the {cardId} path value: the card's unique identifier. suspendCardRequest is sent as
// the JSON body.
//
// It returns the cards.SuspendCardResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." cardId returns a
// CheckoutArgumentError and sends no request.
func (c *Client) SuspendCard(
	cardId string,
	suspendCardRequest cards.SuspendCardRequest,
) (*cards.SuspendCardResponse, error) {
	return c.SuspendCardWithContext(context.Background(), cardId, suspendCardRequest)
}

// SuspendCardWithContext is the context-aware variant of SuspendCard. The request is optional: a zero
// cards.SuspendCardRequest{} sends {}.
func (c *Client) SuspendCardWithContext(
	ctx context.Context,
	cardId string,
	suspendCardRequest cards.SuspendCardRequest,
) (*cards.SuspendCardResponse, error) {
	cardId, err := common.EncodePathSegment(cardId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response cards.SuspendCardResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(issuingPath, cardsPath, cardId, suspendPath),
		auth,
		suspendCardRequest,
		&response,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// UpdateCard updates the details of an issued card. Only the fields for which you provide values
// are updated.
//
// PATCH /issuing/cards/{cardId}
//
// cardId is the {cardId} path value: the card's unique identifier. request is sent as the JSON
// body.
//
// It returns the cards.CardUpdateResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." cardId returns
// a CheckoutArgumentError and sends no request.
func (c *Client) UpdateCard(
	cardId string,
	request cards.CardUpdateRequest,
) (*cards.CardUpdateResponse, error) {
	return c.UpdateCardWithContext(context.Background(), cardId, request)
}

// UpdateCardWithContext is the context-aware variant of UpdateCard.
func (c *Client) UpdateCardWithContext(
	ctx context.Context,
	cardId string,
	request cards.CardUpdateRequest,
) (*cards.CardUpdateResponse, error) {
	return c.updateCard(ctx, cardId, request, nil)
}

// UpdateCardHeaders updates a card like UpdateCard and also sends the return-encrypted-cvv and
// Encryption-Key headers from headers when they are set. The API no longer documents these
// headers; they are still sent when set, and the response type is the same as for UpdateCard.
//
// cardId, request, the return value and the CheckoutArgumentError for an empty, "." or ".."
// cardId are as for UpdateCard.
//
// Deprecated: The API no longer documents these headers; use UpdateCard without headers.
// Will be removed in the next major version.
func (c *Client) UpdateCardHeaders(
	cardId string,
	request cards.CardUpdateRequest,
	headers *cards.CardUpdateHeaders,
) (*cards.CardUpdateResponse, error) {
	return c.UpdateCardHeadersWithContext(context.Background(), cardId, request, headers)
}

// UpdateCardHeadersWithContext is the context-aware variant of UpdateCardHeaders.
//
// Deprecated: The API no longer documents these headers; use UpdateCardWithContext without
// headers. Will be removed in the next major version.
func (c *Client) UpdateCardHeadersWithContext(
	ctx context.Context,
	cardId string,
	request cards.CardUpdateRequest,
	headers *cards.CardUpdateHeaders,
) (*cards.CardUpdateResponse, error) {
	return c.updateCard(ctx, cardId, request, headers)
}

func (c *Client) updateCard(
	ctx context.Context,
	cardId string,
	request cards.CardUpdateRequest,
	headers *cards.CardUpdateHeaders,
) (*cards.CardUpdateResponse, error) {
	cardId, err := common.EncodePathSegment(cardId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	// The ApiClient resolves per-request headers by reflecting over a field named Headers, and
	// json:"-" keeps the wrapper out of the request body. Same shape the accounts client uses.
	body := struct {
		cards.CardUpdateRequest
		Headers *cards.CardUpdateHeaders `json:"-"`
	}{request, headers}

	var response cards.CardUpdateResponse
	err = c.apiClient.PatchWithContext(
		ctx,
		common.BuildPath(issuingPath, cardsPath, cardId),
		auth,
		body,
		&response,
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
	err = c.apiClient.GetWithContext(ctx, url, auth, nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetCardControlDetails sends GET /issuing/controls/{controlId}: get control details.
//
// controlId is the {controlId} path value: the control's unique identifier.
//
// It returns the controls.CardControlResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." controlId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) GetCardControlDetails(controlId string) (*controls.CardControlResponse, error) {
	return c.GetCardControlDetailsWithContext(context.Background(), controlId)
}

// GetCardControlDetailsWithContext is GetCardControlDetails with a caller-supplied context.
func (c *Client) GetCardControlDetailsWithContext(
	ctx context.Context,
	controlId string,
) (*controls.CardControlResponse, error) {
	controlId, err := common.EncodePathSegment(controlId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response controls.CardControlResponse
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(issuingPath, controlsPath, controlId),
		auth,
		nil,
		&response)

	if err != nil {
		return nil, err
	}
	return &response, nil
}

// UpdateCardControl sends PUT /issuing/controls/{controlId}: update a control.
//
// controlId is the {controlId} path value: the control's unique identifier.
// updateCardControlRequest is sent as the JSON body.
//
// It returns the controls.CardControlResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." controlId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) UpdateCardControl(
	controlId string,
	updateCardControlRequest controls.UpdateCardControlRequest,
) (*controls.CardControlResponse, error) {
	return c.UpdateCardControlWithContext(context.Background(), controlId, updateCardControlRequest)
}

// UpdateCardControlWithContext is UpdateCardControl with a caller-supplied context.
func (c *Client) UpdateCardControlWithContext(
	ctx context.Context,
	controlId string,
	updateCardControlRequest controls.UpdateCardControlRequest,
) (*controls.CardControlResponse, error) {
	controlId, err := common.EncodePathSegment(controlId)
	if err != nil {
		return nil, err
	}

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

// RemoveCardControl sends DELETE /issuing/controls/{controlId}: remove a control.
//
// controlId is the {controlId} path value: the control's unique identifier.
//
// It returns the common.IdResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." controlId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) RemoveCardControl(controlId string) (*common.IdResponse, error) {
	return c.RemoveCardControlWithContext(context.Background(), controlId)
}

// RemoveCardControlWithContext is RemoveCardControl with a caller-supplied context.
func (c *Client) RemoveCardControlWithContext(
	ctx context.Context,
	controlId string,
) (*common.IdResponse, error) {
	controlId, err := common.EncodePathSegment(controlId)
	if err != nil {
		return nil, err
	}

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

// SimulateIncrement sends POST /issuing/simulate/authorizations/{id}/authorizations: simulate
// incrementing an authorization.
//
// transactionId is the {id} path value: the transaction's unique identifier. request is sent as the
// JSON body.
//
// It returns the testing.CardSimulationResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." transactionId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) SimulateIncrement(
	transactionId string,
	request testing.CardSimulationRequest,
) (*testing.CardSimulationResponse, error) {
	return c.SimulateIncrementWithContext(context.Background(), transactionId, request)
}

// SimulateIncrementWithContext is SimulateIncrement with a caller-supplied context.
func (c *Client) SimulateIncrementWithContext(
	ctx context.Context,
	transactionId string,
	request testing.CardSimulationRequest,
) (*testing.CardSimulationResponse, error) {
	transactionId, err := common.EncodePathSegment(transactionId)
	if err != nil {
		return nil, err
	}

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

// SimulateClearing sends POST /issuing/simulate/authorizations/{id}/presentments: simulate
// clearing.
//
// transactionId is the {id} path value: the transaction's unique identifier. request is sent as the
// JSON body.
//
// It returns the common.MetadataResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." transactionId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) SimulateClearing(
	transactionId string,
	request testing.CardSimulationRequest,
) (*common.MetadataResponse, error) {
	return c.SimulateClearingWithContext(context.Background(), transactionId, request)
}

// SimulateClearingWithContext is SimulateClearing with a caller-supplied context.
func (c *Client) SimulateClearingWithContext(
	ctx context.Context,
	transactionId string,
	request testing.CardSimulationRequest,
) (*common.MetadataResponse, error) {
	transactionId, err := common.EncodePathSegment(transactionId)
	if err != nil {
		return nil, err
	}

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

// SimulateReversal sends POST /issuing/simulate/authorizations/{id}/reversals: simulate reversal.
//
// transactionId is the {id} path value: the transaction's unique identifier. request is sent as the
// JSON body.
//
// It returns the testing.CardSimulationResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." transactionId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) SimulateReversal(
	transactionId string,
	request testing.CardSimulationRequest,
) (*testing.CardSimulationResponse, error) {
	return c.SimulateReversalWithContext(context.Background(), transactionId, request)
}

// SimulateReversalWithContext is SimulateReversal with a caller-supplied context.
func (c *Client) SimulateReversalWithContext(
	ctx context.Context,
	transactionId string,
	request testing.CardSimulationRequest,
) (*testing.CardSimulationResponse, error) {
	transactionId, err := common.EncodePathSegment(transactionId)
	if err != nil {
		return nil, err
	}

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

// UpdateCardholder sends PATCH /issuing/cardholders/{cardholderId}: update a cardholder.
//
// cardholderId is the {cardholderId} path value: the cardholder's unique identifier. request is
// sent as the JSON body.
//
// It returns the cardholders.CardholderUpdateResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." cardholderId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) UpdateCardholder(
	cardholderId string,
	request cardholders.CardholderRequest,
) (*cardholders.CardholderUpdateResponse, error) {
	return c.UpdateCardholderWithContext(context.Background(), cardholderId, request)
}

// UpdateCardholderWithContext is UpdateCardholder with a caller-supplied context.
func (c *Client) UpdateCardholderWithContext(
	ctx context.Context,
	cardholderId string,
	request cardholders.CardholderRequest,
) (*cardholders.CardholderUpdateResponse, error) {
	cardholderId, err := common.EncodePathSegment(cardholderId)
	if err != nil {
		return nil, err
	}

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

// RenewCard sends POST /issuing/cards/{cardId}/renew: renew a card.
//
// cardId is the {cardId} path value: the card's unique identifier. request is sent as the JSON
// body.
//
// It returns the cards.RenewCardResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." cardId returns a
// CheckoutArgumentError and sends no request.
func (c *Client) RenewCard(cardId string, request cards.RenewCardRequest) (*cards.RenewCardResponse, error) {
	return c.RenewCardWithContext(context.Background(), cardId, request)
}

// RenewCardWithContext is RenewCard with a caller-supplied context.
func (c *Client) RenewCardWithContext(
	ctx context.Context,
	cardId string,
	request cards.RenewCardRequest,
) (*cards.RenewCardResponse, error) {
	cardId, err := common.EncodePathSegment(cardId)
	if err != nil {
		return nil, err
	}

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
	err = c.apiClient.GetWithContext(ctx, url, auth, nil, &response)
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

// GetControlGroupDetails sends GET /issuing/controls/control-groups/{controlGroupId}: get control
// group details.
//
// controlGroupId is the {controlGroupId} path value: the control group's unique identifier.
//
// It returns the controlgroups.ControlGroupResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." controlGroupId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) GetControlGroupDetails(controlGroupId string) (*controlgroups.ControlGroupResponse, error) {
	return c.GetControlGroupDetailsWithContext(context.Background(), controlGroupId)
}

// GetControlGroupDetailsWithContext is GetControlGroupDetails with a caller-supplied context.
func (c *Client) GetControlGroupDetailsWithContext(
	ctx context.Context,
	controlGroupId string,
) (*controlgroups.ControlGroupResponse, error) {
	controlGroupId, err := common.EncodePathSegment(controlGroupId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response controlgroups.ControlGroupResponse
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(issuingPath, controlsPath, controlGroupsPath, controlGroupId),
		auth,
		nil,
		&response)

	if err != nil {
		return nil, err
	}
	return &response, nil
}

// RemoveControlGroup sends DELETE /issuing/controls/control-groups/{controlGroupId}: remove a
// control group.
//
// controlGroupId is the {controlGroupId} path value: the control group's unique identifier.
//
// It returns the common.MetadataResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." controlGroupId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) RemoveControlGroup(controlGroupId string) (*common.MetadataResponse, error) {
	return c.RemoveControlGroupWithContext(context.Background(), controlGroupId)
}

// RemoveControlGroupWithContext is RemoveControlGroup with a caller-supplied context.
func (c *Client) RemoveControlGroupWithContext(
	ctx context.Context,
	controlGroupId string,
) (*common.MetadataResponse, error) {
	controlGroupId, err := common.EncodePathSegment(controlGroupId)
	if err != nil {
		return nil, err
	}

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
	err = c.apiClient.GetWithContext(ctx, url, auth, nil, &response)
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

// GetControlProfileDetails sends GET /issuing/controls/control-profiles/{controlProfileId}: get
// control profile details.
//
// controlProfileId is the {controlProfileId} path value: the control profile's unique identifier.
//
// It returns the controlprofiles.ControlProfileResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." controlProfileId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) GetControlProfileDetails(
	controlProfileId string,
) (*controlprofiles.ControlProfileResponse, error) {
	return c.GetControlProfileDetailsWithContext(context.Background(), controlProfileId)
}

// GetControlProfileDetailsWithContext is GetControlProfileDetails with a caller-supplied context.
func (c *Client) GetControlProfileDetailsWithContext(
	ctx context.Context,
	controlProfileId string,
) (*controlprofiles.ControlProfileResponse, error) {
	controlProfileId, err := common.EncodePathSegment(controlProfileId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response controlprofiles.ControlProfileResponse
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(issuingPath, controlsPath, controlProfilesPath, controlProfileId),
		auth,
		nil,
		&response)

	if err != nil {
		return nil, err
	}
	return &response, nil
}

// UpdateControlProfile sends PATCH /issuing/controls/control-profiles/{controlProfileId}: update a
// control profile.
//
// controlProfileId is the {controlProfileId} path value: the control profile's unique identifier.
// request is sent as the JSON body.
//
// It returns the controlprofiles.ControlProfileResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." controlProfileId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) UpdateControlProfile(
	controlProfileId string,
	request controlprofiles.ControlProfileRequest,
) (*controlprofiles.ControlProfileResponse, error) {
	return c.UpdateControlProfileWithContext(context.Background(), controlProfileId, request)
}

// UpdateControlProfileWithContext is UpdateControlProfile with a caller-supplied context.
func (c *Client) UpdateControlProfileWithContext(
	ctx context.Context,
	controlProfileId string,
	request controlprofiles.ControlProfileRequest,
) (*controlprofiles.ControlProfileResponse, error) {
	controlProfileId, err := common.EncodePathSegment(controlProfileId)
	if err != nil {
		return nil, err
	}

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

// RemoveControlProfile sends DELETE /issuing/controls/control-profiles/{controlProfileId}: remove a
// control profile.
//
// controlProfileId is the {controlProfileId} path value: the control profile's unique identifier.
//
// It returns the common.MetadataResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." controlProfileId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) RemoveControlProfile(controlProfileId string) (*common.MetadataResponse, error) {
	return c.RemoveControlProfileWithContext(context.Background(), controlProfileId)
}

// RemoveControlProfileWithContext is RemoveControlProfile with a caller-supplied context.
func (c *Client) RemoveControlProfileWithContext(
	ctx context.Context,
	controlProfileId string,
) (*common.MetadataResponse, error) {
	controlProfileId, err := common.EncodePathSegment(controlProfileId)
	if err != nil {
		return nil, err
	}

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

// AddTargetToControlProfile sends POST
// /issuing/controls/control-profiles/{controlProfileId}/add/{targetId}: add target to control
// profile.
//
// controlProfileId is the {controlProfileId} path value: the control profile's unique identifier.
// targetId is the {targetId} path value: the ID of the card to apply the control profile to.
//
// It returns the common.MetadataResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." controlProfileId
// or targetId returns a CheckoutArgumentError and sends no request.
func (c *Client) AddTargetToControlProfile(
	controlProfileId, targetId string,
) (*common.MetadataResponse, error) {
	return c.AddTargetToControlProfileWithContext(context.Background(), controlProfileId, targetId)
}

// AddTargetToControlProfileWithContext is AddTargetToControlProfile with a caller-supplied context.
func (c *Client) AddTargetToControlProfileWithContext(
	ctx context.Context,
	controlProfileId, targetId string,
) (*common.MetadataResponse, error) {
	controlProfileId, err := common.EncodePathSegment(controlProfileId)
	if err != nil {
		return nil, err
	}

	targetId, err = common.EncodePathSegment(targetId)
	if err != nil {
		return nil, err
	}

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

// RemoveTargetFromControlProfile sends POST
// /issuing/controls/control-profiles/{controlProfileId}/remove/{targetId}: remove target from
// control profile.
//
// controlProfileId is the {controlProfileId} path value: the control profile's unique identifier.
// targetId is the {targetId} path value: the ID of the card the control profile should no longer
// apply to.
//
// It returns the common.MetadataResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." controlProfileId
// or targetId returns a CheckoutArgumentError and sends no request.
func (c *Client) RemoveTargetFromControlProfile(
	controlProfileId, targetId string,
) (*common.MetadataResponse, error) {
	return c.RemoveTargetFromControlProfileWithContext(context.Background(), controlProfileId, targetId)
}

// RemoveTargetFromControlProfileWithContext is RemoveTargetFromControlProfile with a caller-supplied context.
func (c *Client) RemoveTargetFromControlProfileWithContext(
	ctx context.Context,
	controlProfileId, targetId string,
) (*common.MetadataResponse, error) {
	controlProfileId, err := common.EncodePathSegment(controlProfileId)
	if err != nil {
		return nil, err
	}

	targetId, err = common.EncodePathSegment(targetId)
	if err != nil {
		return nil, err
	}

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

// GetDigitalCard sends GET /issuing/digital-cards/{digitalCardId}: get digital card details.
//
// digitalCardId is the {digitalCardId} path value: the digital card's unique identifier.
//
// It returns the digitalcards.GetDigitalCardResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." digitalCardId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) GetDigitalCard(digitalCardId string) (*digitalcards.GetDigitalCardResponse, error) {
	return c.GetDigitalCardWithContext(context.Background(), digitalCardId)
}

// GetDigitalCardWithContext is GetDigitalCard with a caller-supplied context.
func (c *Client) GetDigitalCardWithContext(
	ctx context.Context,
	digitalCardId string,
) (*digitalcards.GetDigitalCardResponse, error) {
	digitalCardId, err := common.EncodePathSegment(digitalCardId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response digitalcards.GetDigitalCardResponse
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(issuingPath, digitalCardsPath, digitalCardId),
		auth,
		nil,
		&response)

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

// GetDispute sends GET /issuing/disputes/{disputeId}: get an Issuing dispute.
//
// disputeId is the {disputeId} path value: the unique identifier for the Issuing dispute.
//
// It returns the disputes.IssuingDisputeResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." disputeId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) GetDispute(disputeId string) (*disputes.IssuingDisputeResponse, error) {
	return c.GetDisputeWithContext(context.Background(), disputeId)
}

// GetDisputeWithContext is GetDispute with a caller-supplied context.
func (c *Client) GetDisputeWithContext(
	ctx context.Context,
	disputeId string,
) (*disputes.IssuingDisputeResponse, error) {
	disputeId, err := common.EncodePathSegment(disputeId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response disputes.IssuingDisputeResponse
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(issuingPath, disputesPath, disputeId),
		auth,
		nil,
		&response)

	if err != nil {
		return nil, err
	}
	return &response, nil
}

// CancelDispute sends POST /issuing/disputes/{disputeId}/cancel: cancel an Issuing dispute.
//
// disputeId is the {disputeId} path value: the unique identifier for the Issuing dispute.
// idempotencyKey, when not nil, is sent as the Cko-Idempotency-Key header.
//
// It returns the common.MetadataResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." disputeId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) CancelDispute(disputeId string, idempotencyKey *string) (*common.MetadataResponse, error) {
	return c.CancelDisputeWithContext(context.Background(), disputeId, idempotencyKey)
}

// CancelDisputeWithContext is CancelDispute with a caller-supplied context.
func (c *Client) CancelDisputeWithContext(
	ctx context.Context,
	disputeId string,
	idempotencyKey *string,
) (*common.MetadataResponse, error) {
	disputeId, err := common.EncodePathSegment(disputeId)
	if err != nil {
		return nil, err
	}

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

// EscalateDispute sends POST /issuing/disputes/{disputeId}/escalate: escalate an Issuing dispute.
//
// disputeId is the {disputeId} path value: the unique identifier for the Issuing dispute. request
// is sent as the JSON body. idempotencyKey, when not nil, is sent as the Cko-Idempotency-Key
// header.
//
// It returns the common.MetadataResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." disputeId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) EscalateDispute(
	disputeId string,
	request disputes.EscalateDisputeRequest,
	idempotencyKey *string,
) (*common.MetadataResponse, error) {
	return c.EscalateDisputeWithContext(context.Background(), disputeId, request, idempotencyKey)
}

// EscalateDisputeWithContext is EscalateDispute with a caller-supplied context.
func (c *Client) EscalateDisputeWithContext(
	ctx context.Context,
	disputeId string,
	request disputes.EscalateDisputeRequest,
	idempotencyKey *string,
) (*common.MetadataResponse, error) {
	disputeId, err := common.EncodePathSegment(disputeId)
	if err != nil {
		return nil, err
	}

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

// AmendDispute submits an amendment to a dispute that is currently blocked from proceeding.
// It handles both chargeback-stage and escalation-stage amendments.
//
// disputeId is the {disputeId} path value: the unique identifier for the Issuing dispute. request
// is sent as the JSON body. idempotencyKey, when not nil, is sent as the Cko-Idempotency-Key
// header.
//
// It returns the disputes.IssuingDisputeResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." disputeId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) AmendDispute(
	disputeId string,
	request disputes.AmendDisputeRequest,
	idempotencyKey *string,
) (*disputes.IssuingDisputeResponse, error) {
	return c.AmendDisputeWithContext(context.Background(), disputeId, request, idempotencyKey)
}

// AmendDisputeWithContext is AmendDispute with a caller-supplied context.
func (c *Client) AmendDisputeWithContext(
	ctx context.Context,
	disputeId string,
	request disputes.AmendDisputeRequest,
	idempotencyKey *string,
) (*disputes.IssuingDisputeResponse, error) {
	disputeId, err := common.EncodePathSegment(disputeId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response disputes.IssuingDisputeResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(issuingPath, disputesPath, disputeId, amendPath),
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

// SimulateRefund sends POST /issuing/simulate/authorizations/{id}/refunds: simulate refund.
//
// transactionId is the {id} path value: the transaction's unique identifier. request is sent as the
// JSON body.
//
// It returns the common.MetadataResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." transactionId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) SimulateRefund(
	transactionId string,
	request testing.SimulateRefundRequest,
) (*common.MetadataResponse, error) {
	return c.SimulateRefundWithContext(context.Background(), transactionId, request)
}

// SimulateRefundWithContext is SimulateRefund with a caller-supplied context.
func (c *Client) SimulateRefundWithContext(
	ctx context.Context,
	transactionId string,
	request testing.SimulateRefundRequest,
) (*common.MetadataResponse, error) {
	transactionId, err := common.EncodePathSegment(transactionId)
	if err != nil {
		return nil, err
	}

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
	err = c.apiClient.GetWithContext(ctx, url, auth, nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetSingleTransaction sends GET /issuing/transactions/{transactionId}: get a single transaction.
//
// transactionId is the {transactionId} path value: the transaction's unique identifier.
//
// It returns the transactions.TransactionResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." transactionId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) GetSingleTransaction(transactionId string) (*transactions.TransactionResponse, error) {
	return c.GetSingleTransactionWithContext(context.Background(), transactionId)
}

// GetSingleTransactionWithContext is GetSingleTransaction with a caller-supplied context.
func (c *Client) GetSingleTransactionWithContext(
	ctx context.Context,
	transactionId string,
) (*transactions.TransactionResponse, error) {
	transactionId, err := common.EncodePathSegment(transactionId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response transactions.TransactionResponse
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(issuingPath, transactionsPath, transactionId),
		auth,
		nil,
		&response)

	if err != nil {
		return nil, err
	}
	return &response, nil
}
