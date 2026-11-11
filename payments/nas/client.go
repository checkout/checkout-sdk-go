package nas

import (
	"context"

	"github.com/checkout/checkout-sdk-go/v3/client"
	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/configuration"
	"github.com/checkout/checkout-sdk-go/v3/payments"
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
	err = c.apiClient.GetWithContext(ctx, url, auth, nil, &response)
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

// GetPaymentDetails sends GET /payments/{id}: get payment details.
//
// paymentId is the {id} path value: the payment or payment session identifier.
//
// It returns the GetPaymentResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." paymentId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) GetPaymentDetails(paymentId string) (*GetPaymentResponse, error) {
	return c.GetPaymentDetailsWithContext(context.Background(), paymentId)
}

// GetPaymentDetailsWithContext is GetPaymentDetails with a caller-supplied context.
func (c *Client) GetPaymentDetailsWithContext(ctx context.Context, paymentId string) (*GetPaymentResponse, error) {
	paymentId, err := common.EncodePathSegment(paymentId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response GetPaymentResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(payments.PathPayments, paymentId), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetPaymentActions sends GET /payments/{id}/actions: get payment actions.
//
// paymentId is the {id} path value: the payment identifier.
//
// It returns the GetPaymentActionsResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." paymentId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) GetPaymentActions(paymentId string) (*GetPaymentActionsResponse, error) {
	return c.GetPaymentActionsWithContext(context.Background(), paymentId)
}

// GetPaymentActionsWithContext is GetPaymentActions with a caller-supplied context.
func (c *Client) GetPaymentActionsWithContext(ctx context.Context, paymentId string) (*GetPaymentActionsResponse, error) {
	paymentId, err := common.EncodePathSegment(paymentId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response GetPaymentActionsResponse
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(payments.PathPayments, paymentId, "actions"),
		auth,
		nil,
		&response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

// IncrementAuthorization sends POST /payments/{id}/authorizations: increment authorization.
//
// paymentId is the {id} path value: the payment identifier. incrementAuthorizationRequest is sent
// as the JSON body. idempotencyKey, when not nil, is sent as the Cko-Idempotency-Key header.
//
// It returns the IncrementAuthorizationResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." paymentId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) IncrementAuthorization(
	paymentId string,
	incrementAuthorizationRequest IncrementAuthorizationRequest,
	idempotencyKey *string,
) (*IncrementAuthorizationResponse, error) {
	return c.IncrementAuthorizationWithContext(context.Background(), paymentId, incrementAuthorizationRequest, idempotencyKey)
}

// IncrementAuthorizationWithContext is IncrementAuthorization with a caller-supplied context.
func (c *Client) IncrementAuthorizationWithContext(
	ctx context.Context,
	paymentId string,
	incrementAuthorizationRequest IncrementAuthorizationRequest,
	idempotencyKey *string,
) (*IncrementAuthorizationResponse, error) {
	paymentId, err := common.EncodePathSegment(paymentId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response IncrementAuthorizationResponse
	err = c.apiClient.PostWithContext(
		ctx,
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

// CapturePayment sends POST /payments/{id}/captures: capture a payment.
//
// paymentId is the {id} path value: the payment identifier. captureRequest is sent as the JSON
// body. idempotencyKey, when not nil, is sent as the Cko-Idempotency-Key header.
//
// It returns the payments.CaptureResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." paymentId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) CapturePayment(
	paymentId string,
	captureRequest CaptureRequest,
	idempotencyKey *string,
) (*payments.CaptureResponse, error) {
	return c.CapturePaymentWithContext(context.Background(), paymentId, captureRequest, idempotencyKey)
}

// CapturePaymentWithContext is CapturePayment with a caller-supplied context.
func (c *Client) CapturePaymentWithContext(
	ctx context.Context,
	paymentId string,
	captureRequest CaptureRequest,
	idempotencyKey *string,
) (*payments.CaptureResponse, error) {
	paymentId, err := common.EncodePathSegment(paymentId)
	if err != nil {
		return nil, err
	}

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

// CapturePaymentWithoutRequest sends POST /payments/{id}/captures: capture a payment.
//
// paymentId is the {id} path value: the payment identifier. idempotencyKey, when not nil, is sent
// as the Cko-Idempotency-Key header.
//
// It returns the payments.CaptureResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." paymentId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) CapturePaymentWithoutRequest(
	paymentId string,
	idempotencyKey *string,
) (*payments.CaptureResponse, error) {
	return c.CapturePaymentWithoutRequestWithContext(context.Background(), paymentId, idempotencyKey)
}

// CapturePaymentWithoutRequestWithContext is CapturePaymentWithoutRequest with a caller-supplied context.
func (c *Client) CapturePaymentWithoutRequestWithContext(
	ctx context.Context,
	paymentId string,
	idempotencyKey *string,
) (*payments.CaptureResponse, error) {
	paymentId, err := common.EncodePathSegment(paymentId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response payments.CaptureResponse
	err = c.apiClient.PostWithContext(
		ctx,
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

// RefundPayment sends POST /payments/{id}/refunds: refund a payment.
//
// paymentId is the {id} path value: the payment identifier. refundRequest is sent as the JSON body.
// idempotencyKey, when not nil, is sent as the Cko-Idempotency-Key header.
//
// It returns the payments.RefundResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." paymentId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) RefundPayment(
	paymentId string,
	refundRequest *payments.RefundRequest,
	idempotencyKey *string,
) (*payments.RefundResponse, error) {
	return c.RefundPaymentWithContext(context.Background(), paymentId, refundRequest, idempotencyKey)
}

// RefundPaymentWithContext is RefundPayment with a caller-supplied context.
func (c *Client) RefundPaymentWithContext(
	ctx context.Context,
	paymentId string,
	refundRequest *payments.RefundRequest,
	idempotencyKey *string,
) (*payments.RefundResponse, error) {
	paymentId, err := common.EncodePathSegment(paymentId)
	if err != nil {
		return nil, err
	}

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

// VoidPayment sends POST /payments/{id}/voids: void a payment.
//
// paymentId is the {id} path value: the payment identifier. voidRequest is sent as the JSON body.
// idempotencyKey, when not nil, is sent as the Cko-Idempotency-Key header.
//
// It returns the payments.VoidResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." paymentId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) VoidPayment(
	paymentId string,
	voidRequest *payments.VoidRequest,
	idempotencyKey *string,
) (*payments.VoidResponse, error) {
	return c.VoidPaymentWithContext(context.Background(), paymentId, voidRequest, idempotencyKey)
}

// VoidPaymentWithContext is VoidPayment with a caller-supplied context.
func (c *Client) VoidPaymentWithContext(
	ctx context.Context,
	paymentId string,
	voidRequest *payments.VoidRequest,
	idempotencyKey *string,
) (*payments.VoidResponse, error) {
	paymentId, err := common.EncodePathSegment(paymentId)
	if err != nil {
		return nil, err
	}

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

// CancelAScheduledRetry sends POST /payments/{id}/cancellations: cancel a scheduled retry.
//
// paymentId is the {id} path value: the unique payment identifier. request is sent as the JSON
// body. idempotencyKey, when not nil, is sent as the Cko-Idempotency-Key header.
//
// It returns the payments.CancellationResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." paymentId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) CancelAScheduledRetry(
	paymentId string,
	request *payments.CancellationRequest,
	idempotencyKey *string,
) (*payments.CancellationResponse, error) {
	return c.CancelAScheduledRetryWithContext(context.Background(), paymentId, request, idempotencyKey)
}

// CancelAScheduledRetryWithContext is CancelAScheduledRetry with a caller-supplied context.
func (c *Client) CancelAScheduledRetryWithContext(
	ctx context.Context,
	paymentId string,
	request *payments.CancellationRequest,
	idempotencyKey *string,
) (*payments.CancellationResponse, error) {
	paymentId, err := common.EncodePathSegment(paymentId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response payments.CancellationResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(payments.PathPayments, paymentId, "cancellations"),
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

// ReversePayment sends POST /payments/{id}/reversals: reverse a payment.
//
// paymentId is the {id} path value: the unique identifier for the payment. request is sent as the
// JSON body. idempotencyKey, when not nil, is sent as the Cko-Idempotency-Key header.
//
// It returns the payments.PaymentReversalResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." paymentId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) ReversePayment(
	paymentId string,
	request *payments.PaymentReversalRequest,
	idempotencyKey *string,
) (*payments.PaymentReversalResponse, error) {
	return c.ReversePaymentWithContext(context.Background(), paymentId, request, idempotencyKey)
}

// ReversePaymentWithContext is ReversePayment with a caller-supplied context.
func (c *Client) ReversePaymentWithContext(
	ctx context.Context,
	paymentId string,
	request *payments.PaymentReversalRequest,
	idempotencyKey *string,
) (*payments.PaymentReversalResponse, error) {
	paymentId, err := common.EncodePathSegment(paymentId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response payments.PaymentReversalResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(payments.PathPayments, paymentId, "reversals"),
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

func (c *Client) SearchPayments(request SearchPaymentsRequest) (*SearchPaymentsResponse, error) {
	return c.SearchPaymentsWithContext(context.Background(), request)
}

func (c *Client) SearchPaymentsWithContext(ctx context.Context, request SearchPaymentsRequest) (*SearchPaymentsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response SearchPaymentsResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(payments.PathPayments, "search"),
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
