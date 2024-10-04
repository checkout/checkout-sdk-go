package nas

import (
	"context"

	"github.com/checkout/checkout-sdk-go/client"
	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/checkout/checkout-sdk-go/payments"
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

func (c *Client) RequestPayment(request PaymentRequest, idempotencyKey *string) (*PaymentResponse, error) {
	return c.RequestPaymentWithContext(context.Background(), request, idempotencyKey)
}

func (c *Client) RequestPaymentWithContext(ctx context.Context, request PaymentRequest, idempotencyKey *string) (*PaymentResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response PaymentResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(payments.PathPayments),
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

func (c *Client) RequestPaymentList(request payments.QueryRequest) (*GetPaymentListResponse, error) {
	return c.RequestPaymentListWithContext(context.Background(), request)
}

func (c *Client) RequestPaymentListWithContext(ctx context.Context, request payments.QueryRequest) (*GetPaymentListResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	url, err := common.BuildQueryPath(common.BuildPath(payments.PathPayments), request)
	if err != nil {
		return nil, err
	}

	var response GetPaymentListResponse
	err = c.apiClient.GetWithContext(ctx, url, auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) RequestPayout(request PayoutRequest, idempotencyKey *string) (*PayoutResponse, error) {
	return c.RequestPayoutWithContext(context.Background(), request, idempotencyKey)
}

func (c *Client) RequestPayoutWithContext(ctx context.Context, request PayoutRequest, idempotencyKey *string) (*PayoutResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response PayoutResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(payments.PathPayments),
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

func (c *Client) GetPaymentDetails(paymentId string) (*GetPaymentResponse, error) {
	return c.GetPaymentDetailsWithContext(context.Background(), paymentId)
}

func (c *Client) GetPaymentDetailsWithContext(ctx context.Context, paymentId string) (*GetPaymentResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response GetPaymentResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(payments.PathPayments, paymentId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetPaymentActions(paymentId string) (*GetPaymentActionsResponse, error) {
	return c.GetPaymentActionsWithContext(context.Background(), paymentId)
}

func (c *Client) GetPaymentActionsWithContext(ctx context.Context, paymentId string) (*GetPaymentActionsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response GetPaymentActionsResponse
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(payments.PathPayments, paymentId, "actions"),
		auth,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) IncrementAuthorization(
	paymentId string,
	incrementAuthorizationRequest IncrementAuthorizationRequest,
	idempotencyKey *string,
) (*IncrementAuthorizationResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response IncrementAuthorizationResponse
	err = c.apiClient.Post(
		common.BuildPath(payments.PathPayments, paymentId, "authorizations"),
		auth,
		incrementAuthorizationRequest,
		&response,
		idempotencyKey,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) CapturePayment(
	paymentId string,
	captureRequest CaptureRequest,
	idempotencyKey *string,
) (*payments.CaptureResponse, error) {
	return c.CapturePaymentWithContext(context.Background(), paymentId, captureRequest, idempotencyKey)
}

func (c *Client) CapturePaymentWithContext(
	ctx context.Context,
	paymentId string,
	captureRequest CaptureRequest,
	idempotencyKey *string,
) (*payments.CaptureResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response payments.CaptureResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(payments.PathPayments, paymentId, "captures"),
		auth,
		captureRequest,
		&response,
		idempotencyKey,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) CapturePaymentWithoutRequest(
	paymentId string,
	idempotencyKey *string,
) (*payments.CaptureResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response payments.CaptureResponse
	err = c.apiClient.Post(
		common.BuildPath(payments.PathPayments, paymentId, "captures"),
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

func (c *Client) RefundPayment(
	paymentId string,
	refundRequest *payments.RefundRequest,
	idempotencyKey *string,
) (*payments.RefundResponse, error) {
	return c.RefundPaymentWithContext(context.Background(), paymentId, refundRequest, idempotencyKey)
}

func (c *Client) RefundPaymentWithContext(
	ctx context.Context,
	paymentId string,
	refundRequest *payments.RefundRequest,
	idempotencyKey *string,
) (*payments.RefundResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response payments.RefundResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(payments.PathPayments, paymentId, "refunds"),
		auth,
		refundRequest,
		&response,
		idempotencyKey,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) VoidPayment(
	paymentId string,
	voidRequest *payments.VoidRequest,
	idempotencyKey *string,
) (*payments.VoidResponse, error) {
	return c.VoidPaymentWithContext(context.Background(), paymentId, voidRequest, idempotencyKey)
}

func (c *Client) VoidPaymentWithContext(
	ctx context.Context,
	paymentId string,
	voidRequest *payments.VoidRequest,
	idempotencyKey *string,
) (*payments.VoidResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response payments.VoidResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(payments.PathPayments, paymentId, "voids"),
		auth,
		voidRequest,
		&response,
		idempotencyKey,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
