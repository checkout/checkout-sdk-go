package applicants

import (
	"context"

	"github.com/checkout/checkout-sdk-go/v2/client"
	"github.com/checkout/checkout-sdk-go/v2/common"
	"github.com/checkout/checkout-sdk-go/v2/configuration"
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

func (c *Client) CreateApplicant(request CreateApplicantRequest) (*ApplicantResponse, error) {
	return c.CreateApplicantWithContext(context.Background(), request)
}

func (c *Client) CreateApplicantWithContext(ctx context.Context, request CreateApplicantRequest) (*ApplicantResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response ApplicantResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(applicantsPath), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetApplicant(applicantId string) (*ApplicantResponse, error) {
	return c.GetApplicantWithContext(context.Background(), applicantId)
}

func (c *Client) GetApplicantWithContext(ctx context.Context, applicantId string) (*ApplicantResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response ApplicantResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(applicantsPath, applicantId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UpdateApplicant(applicantId string, request UpdateApplicantRequest) (*ApplicantResponse, error) {
	return c.UpdateApplicantWithContext(context.Background(), applicantId, request)
}

func (c *Client) UpdateApplicantWithContext(ctx context.Context, applicantId string, request UpdateApplicantRequest) (*ApplicantResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response ApplicantResponse
	err = c.apiClient.PatchWithContext(ctx, common.BuildPath(applicantsPath, applicantId), auth, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) AnonymizeApplicant(applicantId string) (*ApplicantResponse, error) {
	return c.AnonymizeApplicantWithContext(context.Background(), applicantId)
}

func (c *Client) AnonymizeApplicantWithContext(ctx context.Context, applicantId string) (*ApplicantResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response ApplicantResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(applicantsPath, applicantId, anonymizePath), auth, nil, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
