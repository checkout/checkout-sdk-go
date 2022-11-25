package accounts

import (
	"github.com/checkout/checkout-sdk-go-beta/client"
	"github.com/checkout/checkout-sdk-go-beta/common"
	"github.com/checkout/checkout-sdk-go-beta/configuration"
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
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	var response OnboardEntityResponse
	err = c.apiClient.Post(
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

func (c *Client) GetEntity(entityId string) (*OnboardEntityDetails, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	var response OnboardEntityDetails
	err = c.apiClient.Get(
		common.BuildPath(accountsPath, entitiesPath, entityId),
		auth,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UpdateEntity(entityId string, request OnboardEntityRequest) (*OnboardEntityResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	var response OnboardEntityResponse
	err = c.apiClient.Put(
		common.BuildPath(accountsPath, entitiesPath, entityId),
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

func (c *Client) CreatePaymentInstrument(entityId string, request PaymentInstrument) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.Post(
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

func (c *Client) RetrievePayoutSchedule(entityId string) (*PayoutSchedule, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	var response PayoutSchedule
	err = c.apiClient.Get(
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
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	request := map[common.Currency]CurrencySchedule{
		currency: updateSchedule,
	}

	var response common.IdResponse
	err = c.apiClient.Put(
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

func (c *Client) UploadFile(file File) (*common.IdResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	req, err := common.BuildFileUploadRequest(&file)
	if err != nil {
		return nil, err
	}

	var response common.IdResponse
	err = c.filesClient.Upload(common.BuildPath(filesPath), auth, req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
