package accounts

import (
	"context"

	"github.com/checkout/checkout-sdk-go/client"
	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/configuration"
)

type Client struct {
	configuration *configuration.Configuration
	apiClient     client.HttpClient
	filesClient   client.HttpClient
}

func NewClient(
	configuration *configuration.Configuration,
	apiClient client.HttpClient,
	filesClient client.HttpClient,
) *Client {
	return &Client{
		configuration: configuration,
		apiClient:     apiClient,
		filesClient:   filesClient,
	}
}

func (c *Client) CreateEntity(request OnboardEntityRequest) (*OnboardEntityResponse, error) {
	return c.CreateEntityWithContext(context.Background(), request)
}

func (c *Client) CreateEntityWithContext(
	ctx context.Context,
	request OnboardEntityRequest,
) (*OnboardEntityResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response OnboardEntityResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(accountsPath, entitiesPath),
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

func (c *Client) GetSubEntityMembers(entityId string) (*OnboardSubEntityDetailsResponse, error) {
	return c.GetSubEntityMembersWithContext(context.Background(), entityId)
}

func (c *Client) GetSubEntityMembersWithContext(
	ctx context.Context,
	entityId string,
) (*OnboardSubEntityDetailsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response OnboardSubEntityDetailsResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(accountsPath, entitiesPath, entityId, membersPath), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) ReinviteSubEntityMember(
	entityId string,
	userId string,
	request OnboardSubEntityRequest,
) (*OnboardSubEntityResponse, error) {
	return c.ReinviteSubEntityMemberWithContext(context.Background(), entityId, userId, request)
}

func (c *Client) ReinviteSubEntityMemberWithContext(
	ctx context.Context,
	entityId string,
	userId string,
	request OnboardSubEntityRequest,
) (*OnboardSubEntityResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response OnboardSubEntityResponse
	err = c.apiClient.PutWithContext(ctx, common.BuildPath(
		accountsPath,
		entitiesPath,
		entityId,
		membersPath,
		userId,
	), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetEntity(entityId string) (*OnboardEntityDetails, error) {
	return c.GetEntityWithContext(context.Background(), entityId)
}

func (c *Client) GetEntityWithContext(ctx context.Context, entityId string) (*OnboardEntityDetails, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response OnboardEntityDetails
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(accountsPath, entitiesPath, entityId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UpdateEntity(entityId string, request OnboardEntityRequest) (*OnboardEntityResponse, error) {
	return c.UpdateEntityWithContext(context.Background(), entityId, request)
}

func (c *Client) UpdateEntityWithContext(
	ctx context.Context,
	entityId string,
	request OnboardEntityRequest,
) (*OnboardEntityResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response OnboardEntityResponse
	err = c.apiClient.PutWithContext(ctx, common.BuildPath(accountsPath, entitiesPath, entityId), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// Deprecated: Use CreatePaymentInstrument IdResponse for CreatePaymentInstrument instead.
func (c *Client) CreatePaymentInstruments(
	entityId string,
	request PaymentInstrument,
) (*common.MetadataResponse, error) {
	return c.CreatePaymentInstrumentsWithContext(context.Background(), entityId, request)
}

// Deprecated: Use CreatePaymentInstrumentWithContext IdResponse for CreatePaymentInstrumentsWithContext instead.
func (c *Client) CreatePaymentInstrumentsWithContext(
	ctx context.Context,
	entityId string,
	request PaymentInstrument,
) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(accountsPath, entitiesPath, entityId, instrumentsPath),
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

func (c *Client) CreatePaymentInstrument(
	entityId string,
	request PaymentInstrumentRequest,
) (*common.IdResponse, error) {
	return c.CreatePaymentInstrumentWithContext(context.Background(), entityId, request)
}

func (c *Client) CreatePaymentInstrumentWithContext(
	ctx context.Context,
	entityId string,
	request PaymentInstrumentRequest,
) (*common.IdResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.IdResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(accountsPath, entitiesPath, entityId, paymentInstrumentsPath),
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

func (c *Client) RetrievePaymentInstrumentDetails(
	entityId string,
	paymentInstrumentId string,
) (*PaymentInstrumentDetailsResponse, error) {
	return c.RetrievePaymentInstrumentDetailsWithContext(context.Background(), entityId, paymentInstrumentId)
}

func (c *Client) RetrievePaymentInstrumentDetailsWithContext(
	ctx context.Context,
	entityId string,
	paymentInstrumentId string,
) (*PaymentInstrumentDetailsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response PaymentInstrumentDetailsResponse
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(accountsPath, entitiesPath, entityId, paymentInstrumentsPath, paymentInstrumentId),
		auth,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UpdatePaymentInstrumentDetails(
	entityId, instrumentId string,
	request UpdatePaymentInstrumentRequest,
) (*common.IdResponse, error) {
	return c.UpdatePaymentInstrumentDetailsWithContext(context.Background(), entityId, instrumentId, request)
}

func (c *Client) UpdatePaymentInstrumentDetailsWithContext(
	ctx context.Context,
	entityId, instrumentId string,
	request UpdatePaymentInstrumentRequest,
) (*common.IdResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.IdResponse
	err = c.apiClient.PatchWithContext(
		ctx,
		common.BuildPath(accountsPath, entitiesPath, entityId, paymentInstrumentsPath, instrumentId),
		auth,
		request,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) QueryPaymentInstruments(
	entityId string,
	query PaymentInstrumentsQuery,
) (*PaymentInstrumentQueryResponse, error) {
	return c.QueryPaymentInstrumentsWithContext(context.Background(), entityId, query)
}

func (c *Client) QueryPaymentInstrumentsWithContext(
	ctx context.Context,
	entityId string,
	query PaymentInstrumentsQuery,
) (*PaymentInstrumentQueryResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	url, err := common.BuildQueryPath(
		common.BuildPath(accountsPath, entitiesPath, entityId, paymentInstrumentsPath),
		query,
	)
	if err != nil {
		return nil, err
	}

	var response PaymentInstrumentQueryResponse
	err = c.apiClient.GetWithContext(ctx, url, auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) RetrievePayoutSchedule(entityId string) (*PayoutSchedule, error) {
	return c.RetrievePayoutScheduleWithContext(context.Background(), entityId)
}

func (c *Client) RetrievePayoutScheduleWithContext(
	ctx context.Context,
	entityId string,
) (*PayoutSchedule, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response PayoutSchedule
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(accountsPath, entitiesPath, entityId, payoutSchedulesPath),
		auth,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UpdatePayoutSchedule(
	entityId string,
	currency common.Currency,
	updateSchedule CurrencySchedule,
) (*common.IdResponse, error) {
	return c.UpdatePayoutScheduleWithContext(context.Background(), entityId, currency, updateSchedule)
}

func (c *Client) UpdatePayoutScheduleWithContext(
	ctx context.Context,
	entityId string,
	currency common.Currency,
	updateSchedule CurrencySchedule,
) (*common.IdResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	request := map[common.Currency]CurrencySchedule{
		currency: updateSchedule,
	}

	var response common.IdResponse
	err = c.apiClient.PutWithContext(
		ctx,
		common.BuildPath(accountsPath, entitiesPath, entityId, payoutSchedulesPath),
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

func (c *Client) SubmitFile(file File) (*common.IdResponse, error) {
	return c.SubmitFileWithContext(context.Background(), file)
}

func (c *Client) SubmitFileWithContext(
	ctx context.Context,
	file File,
) (*common.IdResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	req, err := common.BuildFileUploadRequest(&file)
	if err != nil {
		return nil, err
	}

	var response common.IdResponse
	err = c.filesClient.UploadWithContext(ctx, common.BuildPath(filesPath), auth, req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UploadFile(entityId string, request File) (*UploadFileResponse, error) {
	return c.UploadFileWithContext(context.Background(), entityId, request)
}

func (c *Client) UploadFileWithContext(
	ctx context.Context,
	entityId string,
	request File,
) (*UploadFileResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response UploadFileResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(accountsPath, entityId, filesPath),
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

func (c *Client) RetrieveFile(entityId, fileId string) (*FileDetailsResponse, error) {
	return c.RetrieveFileWithContext(context.Background(), entityId, fileId)
}

func (c *Client) RetrieveFileWithContext(ctx context.Context, entityId, fileId string) (*FileDetailsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response FileDetailsResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(accountsPath, entityId, filesPath, fileId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
