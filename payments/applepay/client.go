package applepay

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

func (c *Client) UploadPaymentProcessingCertificate(request UploadCertificateRequest) (*UploadCertificateResponse, error) {
	return c.UploadPaymentProcessingCertificateWithContext(context.Background(), request)
}

func (c *Client) UploadPaymentProcessingCertificateWithContext(ctx context.Context, request UploadCertificateRequest) (*UploadCertificateResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.PublicKey)
	if err != nil {
		return nil, err
	}

	var response UploadCertificateResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(ApplePayCertificatesPath), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) EnrollDomain(request EnrollDomainRequest) (*common.MetadataResponse, error) {
	return c.EnrollDomainWithContext(context.Background(), request)
}

func (c *Client) EnrollDomainWithContext(ctx context.Context, request EnrollDomainRequest) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(ApplePayEnrollmentsPath), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GenerateCertificateSigningRequest(request GenerateCertificateSigningRequest) (*GenerateCertificateSigningRequestResponse, error) {
	return c.GenerateCertificateSigningRequestWithContext(context.Background(), request)
}

func (c *Client) GenerateCertificateSigningRequestWithContext(ctx context.Context, request GenerateCertificateSigningRequest) (*GenerateCertificateSigningRequestResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.PublicKey)
	if err != nil {
		return nil, err
	}

	var response GenerateCertificateSigningRequestResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(ApplePaySigningRequestsPath), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
