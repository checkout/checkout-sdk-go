package applicants

import (
	"context"

	"github.com/checkout/checkout-sdk-go/v3/client"
	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/configuration"
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

// GetApplicant sends GET /applicants/{applicant_id}: get an applicant.
//
// applicantId is the {applicant_id} path value: the applicant profile's unique identifier.
//
// It returns the ApplicantResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." applicantId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) GetApplicant(applicantId string) (*ApplicantResponse, error) {
	return c.GetApplicantWithContext(context.Background(), applicantId)
}

// GetApplicantWithContext is GetApplicant with a caller-supplied context.
func (c *Client) GetApplicantWithContext(ctx context.Context, applicantId string) (*ApplicantResponse, error) {
	applicantId, err := common.EncodePathSegment(applicantId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response ApplicantResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(applicantsPath, applicantId), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdateApplicant sends PATCH /applicants/{applicant_id}: update an applicant.
//
// applicantId is the {applicant_id} path value: the applicant profile's unique identifier. request
// is sent as the JSON body.
//
// It returns the ApplicantResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." applicantId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) UpdateApplicant(applicantId string, request UpdateApplicantRequest) (*ApplicantResponse, error) {
	return c.UpdateApplicantWithContext(context.Background(), applicantId, request)
}

// UpdateApplicantWithContext is UpdateApplicant with a caller-supplied context.
func (c *Client) UpdateApplicantWithContext(ctx context.Context, applicantId string, request UpdateApplicantRequest) (*ApplicantResponse, error) {
	applicantId, err := common.EncodePathSegment(applicantId)
	if err != nil {
		return nil, err
	}

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

// AnonymizeApplicant sends POST /applicants/{applicant_id}/anonymize: anonymize an applicant.
//
// applicantId is the {applicant_id} path value: the applicant profile's unique identifier.
//
// It returns the ApplicantResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." applicantId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) AnonymizeApplicant(applicantId string) (*ApplicantResponse, error) {
	return c.AnonymizeApplicantWithContext(context.Background(), applicantId)
}

// AnonymizeApplicantWithContext is AnonymizeApplicant with a caller-supplied context.
func (c *Client) AnonymizeApplicantWithContext(ctx context.Context, applicantId string) (*ApplicantResponse, error) {
	applicantId, err := common.EncodePathSegment(applicantId)
	if err != nil {
		return nil, err
	}

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
