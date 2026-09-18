package inventory

import (
	"context"

	"github.com/checkout/checkout-sdk-go/v3/client"
	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/configuration"
)

// Client holds the dependencies for making Inventory API requests.
//
// Every method authorizes with configuration.OAuth (scope agentic:inventory). None of these
// endpoints accept ApiSecretKey/ApiPublicKey authorization.
type Client struct {
	configuration *configuration.Configuration
	apiClient     client.HttpClient
}

// NewClient creates an Inventory Client using the provided configuration and HTTP client.
func NewClient(configuration *configuration.Configuration, apiClient client.HttpClient) *Client {
	return &Client{
		configuration: configuration,
		apiClient:     apiClient,
	}
}

// AdjustInventory applies a signed stock adjustment to a variant and returns its resulting
// stock levels.
//
// POST /inventory/adjustments
//
// Returns 201 on a new adjustment, or 200 with a Cache-Control response header on an
// idempotent replay (inspect response.HttpMetadata.Headers.Header for Cache-Control); both
// share the InventoryLevels response schema. A negative delta that would drive on_hand below
// zero is rejected with a 409, surfaced as errors.CheckoutAPIError.
func (c *Client) AdjustInventory(
	request InventoryAdjustmentRequest,
	idempotencyKey *string,
) (*InventoryLevels, error) {
	return c.AdjustInventoryWithContext(context.Background(), request, idempotencyKey)
}

// AdjustInventoryWithContext is AdjustInventory with a caller-supplied context.
func (c *Client) AdjustInventoryWithContext(
	ctx context.Context,
	request InventoryAdjustmentRequest,
	idempotencyKey *string,
) (*InventoryLevels, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	var response InventoryLevels
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(inventoryPath, adjustments),
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

// CreateInventoryReservation atomically holds stock across one or more variants.
//
// POST /inventory/reservations
//
// Returns 201 on a new reservation, or 200 with a Cache-Control response header on an
// idempotent replay (inspect response.HttpMetadata.Headers.Header for Cache-Control); both
// share the InventoryReservation response schema.
func (c *Client) CreateInventoryReservation(
	request InventoryReservationRequest,
	idempotencyKey *string,
) (*InventoryReservation, error) {
	return c.CreateInventoryReservationWithContext(context.Background(), request, idempotencyKey)
}

// CreateInventoryReservationWithContext is CreateInventoryReservation with a caller-supplied
// context.
func (c *Client) CreateInventoryReservationWithContext(
	ctx context.Context,
	request InventoryReservationRequest,
	idempotencyKey *string,
) (*InventoryReservation, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	var response InventoryReservation
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(inventoryPath, reservations),
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

// GetInventoryReservation retrieves a reservation by id.
//
// GET /inventory/reservations/{id}
func (c *Client) GetInventoryReservation(id string) (*InventoryReservation, error) {
	return c.GetInventoryReservationWithContext(context.Background(), id)
}

// GetInventoryReservationWithContext is GetInventoryReservation with a caller-supplied context.
func (c *Client) GetInventoryReservationWithContext(
	ctx context.Context,
	id string,
) (*InventoryReservation, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	var response InventoryReservation
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(inventoryPath, reservations, id),
		auth,
		nil,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// CommitInventoryReservation converts a held reservation into a permanent stock deduction.
//
// POST /inventory/reservations/{id}/commit
//
// The request has no body.
func (c *Client) CommitInventoryReservation(id string) (*InventoryReservation, error) {
	return c.CommitInventoryReservationWithContext(context.Background(), id)
}

// CommitInventoryReservationWithContext is CommitInventoryReservation with a caller-supplied
// context.
func (c *Client) CommitInventoryReservationWithContext(
	ctx context.Context,
	id string,
) (*InventoryReservation, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	var response InventoryReservation
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(inventoryPath, reservations, id, commitPath),
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

// ReleaseInventoryReservation cancels a held reservation and returns its stock to availability.
//
// POST /inventory/reservations/{id}/release
//
// The request has no body.
func (c *Client) ReleaseInventoryReservation(id string) (*InventoryReservation, error) {
	return c.ReleaseInventoryReservationWithContext(context.Background(), id)
}

// ReleaseInventoryReservationWithContext is ReleaseInventoryReservation with a caller-supplied
// context.
func (c *Client) ReleaseInventoryReservationWithContext(
	ctx context.Context,
	id string,
) (*InventoryReservation, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	var response InventoryReservation
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(inventoryPath, reservations, id, releasePath),
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

// GetInventoryLevels retrieves a variant's stock levels. Pass a non-empty InventoryLevelsQuery
// with Expand set to "product" to embed the variant's product knowledge in the response.
//
// GET /inventory/{variant_id}
func (c *Client) GetInventoryLevels(variantId string, query InventoryLevelsQuery) (*InventoryLevels, error) {
	return c.GetInventoryLevelsWithContext(context.Background(), variantId, query)
}

// GetInventoryLevelsWithContext is GetInventoryLevels with a caller-supplied context.
func (c *Client) GetInventoryLevelsWithContext(
	ctx context.Context,
	variantId string,
	query InventoryLevelsQuery,
) (*InventoryLevels, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	url, err := common.BuildQueryPath(common.BuildPath(inventoryPath, variantId), query)
	if err != nil {
		return nil, err
	}

	var response InventoryLevels
	err = c.apiClient.GetWithContext(ctx, url, auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// SetInventoryLevels creates or updates a variant's stock levels.
//
// PUT /inventory/{variant_id}
func (c *Client) SetInventoryLevels(
	variantId string,
	request InventorySetLevelsRequest,
) (*InventoryLevels, error) {
	return c.SetInventoryLevelsWithContext(context.Background(), variantId, request)
}

// SetInventoryLevelsWithContext is SetInventoryLevels with a caller-supplied context.
func (c *Client) SetInventoryLevelsWithContext(
	ctx context.Context,
	variantId string,
	request InventorySetLevelsRequest,
) (*InventoryLevels, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	var response InventoryLevels
	err = c.apiClient.PutWithContext(
		ctx,
		common.BuildPath(inventoryPath, variantId),
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

// GetInventoryProduct retrieves a variant's product knowledge (merchandising metadata for AI
// agents).
//
// GET /inventory/{variant_id}/product
//
// Beta: this operation is marked Beta in the API specification.
func (c *Client) GetInventoryProduct(variantId string) (*InventoryProductKnowledge, error) {
	return c.GetInventoryProductWithContext(context.Background(), variantId)
}

// GetInventoryProductWithContext is GetInventoryProduct with a caller-supplied context.
func (c *Client) GetInventoryProductWithContext(
	ctx context.Context,
	variantId string,
) (*InventoryProductKnowledge, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	var response InventoryProductKnowledge
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(inventoryPath, variantId, productPath),
		auth,
		nil,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// SetInventoryProduct creates or replaces a variant's product knowledge.
//
// PUT /inventory/{variant_id}/product
//
// Beta: this operation is marked Beta in the API specification.
func (c *Client) SetInventoryProduct(
	variantId string,
	request InventorySetProductRequest,
) (*InventoryProductKnowledge, error) {
	return c.SetInventoryProductWithContext(context.Background(), variantId, request)
}

// SetInventoryProductWithContext is SetInventoryProduct with a caller-supplied context.
func (c *Client) SetInventoryProductWithContext(
	ctx context.Context,
	variantId string,
	request InventorySetProductRequest,
) (*InventoryProductKnowledge, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	var response InventoryProductKnowledge
	err = c.apiClient.PutWithContext(
		ctx,
		common.BuildPath(inventoryPath, variantId, productPath),
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

// DeleteInventoryProduct deletes a variant's product knowledge.
//
// DELETE /inventory/{variant_id}/product
//
// Beta: this operation is marked Beta in the API specification. Returns 204 with no body.
func (c *Client) DeleteInventoryProduct(variantId string) (*common.MetadataResponse, error) {
	return c.DeleteInventoryProductWithContext(context.Background(), variantId)
}

// DeleteInventoryProductWithContext is DeleteInventoryProduct with a caller-supplied context.
func (c *Client) DeleteInventoryProductWithContext(
	ctx context.Context,
	variantId string,
) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.DeleteWithContext(
		ctx,
		common.BuildPath(inventoryPath, variantId, productPath),
		auth,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
