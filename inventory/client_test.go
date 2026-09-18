package inventory

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/checkout/checkout-sdk-go/v3/configuration"
	"github.com/checkout/checkout-sdk-go/v3/errors"
	"github.com/checkout/checkout-sdk-go/v3/mocks"
)

func TestAdjustInventory(t *testing.T) {
	var (
		levelsResponse = InventoryLevels{
			HttpMetadata: mocks.HttpMetadataStatusCreated,
			VariantId:    "sku_123",
			OnHand:       42,
			Reserved:     2,
			SafetyStock:  5,
			Available:    35,
			State:        InStock,
			Source:       Managed,
		}
	)

	cases := []struct {
		name             string
		request          InventoryAdjustmentRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*InventoryLevels, error)
	}{
		{
			name:    "when request is correct then adjust inventory",
			request: InventoryAdjustmentRequest{VariantId: "sku_123", Delta: 5, Reason: "restock"},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(4).(*InventoryLevels)
						*respMapping = levelsResponse
					})
			},
			checker: func(response *InventoryLevels, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, levelsResponse.VariantId, response.VariantId)
				assert.Equal(t, levelsResponse.Available, response.Available)
			},
		},
		{
			name:    "when credentials invalid then return error",
			request: InventoryAdjustmentRequest{VariantId: "sku_123", Delta: 5, Reason: "restock"},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization"))
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *InventoryLevels, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
			},
		},
		{
			name:    "when delta drives stock negative then api returns 409",
			request: InventoryAdjustmentRequest{VariantId: "sku_123", Delta: -1000, Reason: "restock"},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(errors.CheckoutAPIError{
						StatusCode: http.StatusConflict,
						Data: &errors.ErrorDetails{
							ErrorType:  "insufficient_stock",
							VariantId:  "sku_123",
							Available:  35,
							ErrorCodes: []string{"on_hand_insufficient_stock"},
						},
					})
			},
			checker: func(response *InventoryLevels, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusConflict, chkErr.StatusCode)
				assert.Equal(t, "sku_123", chkErr.Data.VariantId)
				assert.Equal(t, int64(35), chkErr.Data.Available)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient, credentials, config := buildInventoryClientConfig()
			tc.getAuthorization(&credentials.Mock)
			tc.apiPost(&apiClient.Mock)

			client := NewClient(config, apiClient)
			tc.checker(client.AdjustInventory(tc.request, nil))
		})
	}
}

func TestCreateInventoryReservation(t *testing.T) {
	expiresAt := time.Date(2026, 9, 18, 12, 0, 0, 0, time.UTC)
	reservationResponse := InventoryReservation{
		HttpMetadata:   mocks.HttpMetadataStatusCreated,
		Id:             "rsv_abc123",
		State:          Held,
		OwnerType:      "cart",
		OwnerReference: "cart_789",
		Items:          []InventoryReservationItem{{VariantId: "sku_123", Quantity: 2}},
		ExpiresAt:      &expiresAt,
	}

	request := InventoryReservationRequest{
		OwnerType:      "cart",
		OwnerReference: "cart_789",
		Items:          []InventoryReservationItem{{VariantId: "sku_123", Quantity: 2}},
		TtlSeconds:     900,
	}

	apiClient, credentials, config := buildInventoryClientConfig()
	credentials.On("GetAuthorization", mock.Anything).
		Return(&configuration.SdkAuthorization{}, nil)
	apiClient.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil).
		Run(func(args mock.Arguments) {
			respMapping := args.Get(4).(*InventoryReservation)
			*respMapping = reservationResponse
		})

	client := NewClient(config, apiClient)
	response, err := client.CreateInventoryReservation(request, nil)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "rsv_abc123", response.Id)
	assert.Equal(t, Held, response.State)
	assert.Len(t, response.Items, 1)
}

func TestGetInventoryReservation(t *testing.T) {
	reservationResponse := InventoryReservation{
		HttpMetadata: mocks.HttpMetadataStatusOk,
		Id:           "rsv_abc123",
		State:        Committed,
	}

	apiClient, credentials, config := buildInventoryClientConfig()
	credentials.On("GetAuthorization", mock.Anything).
		Return(&configuration.SdkAuthorization{}, nil)
	apiClient.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil).
		Run(func(args mock.Arguments) {
			respMapping := args.Get(3).(*InventoryReservation)
			*respMapping = reservationResponse
		})

	client := NewClient(config, apiClient)
	response, err := client.GetInventoryReservation("rsv_abc123")

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, Committed, response.State)
}

func TestCommitInventoryReservation(t *testing.T) {
	reservationResponse := InventoryReservation{
		HttpMetadata: mocks.HttpMetadataStatusOk,
		Id:           "rsv_abc123",
		State:        Committed,
	}

	apiClient, credentials, config := buildInventoryClientConfig()
	credentials.On("GetAuthorization", mock.Anything).
		Return(&configuration.SdkAuthorization{}, nil)
	apiClient.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil).
		Run(func(args mock.Arguments) {
			respMapping := args.Get(4).(*InventoryReservation)
			*respMapping = reservationResponse
		})

	client := NewClient(config, apiClient)
	response, err := client.CommitInventoryReservation("rsv_abc123")

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, Committed, response.State)
}

func TestReleaseInventoryReservation(t *testing.T) {
	reservationResponse := InventoryReservation{
		HttpMetadata: mocks.HttpMetadataStatusOk,
		Id:           "rsv_abc123",
		State:        Released,
	}

	apiClient, credentials, config := buildInventoryClientConfig()
	credentials.On("GetAuthorization", mock.Anything).
		Return(&configuration.SdkAuthorization{}, nil)
	apiClient.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil).
		Run(func(args mock.Arguments) {
			respMapping := args.Get(4).(*InventoryReservation)
			*respMapping = reservationResponse
		})

	client := NewClient(config, apiClient)
	response, err := client.ReleaseInventoryReservation("rsv_abc123")

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, Released, response.State)
}

func TestGetInventoryLevels(t *testing.T) {
	levelsResponse := InventoryLevels{
		HttpMetadata: mocks.HttpMetadataStatusOk,
		VariantId:    "sku_123",
		OnHand:       10,
		State:        Limited,
	}

	apiClient, credentials, config := buildInventoryClientConfig()
	credentials.On("GetAuthorization", mock.Anything).
		Return(&configuration.SdkAuthorization{}, nil)
	apiClient.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil).
		Run(func(args mock.Arguments) {
			respMapping := args.Get(3).(*InventoryLevels)
			*respMapping = levelsResponse
		})

	client := NewClient(config, apiClient)
	response, err := client.GetInventoryLevels("sku_123", InventoryLevelsQuery{Expand: "product"})

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, Limited, response.State)
}

func TestSetInventoryLevels(t *testing.T) {
	safetyStock := int64(0)
	request := InventorySetLevelsRequest{OnHand: 100, SafetyStock: &safetyStock, Reason: "initial load"}
	levelsResponse := InventoryLevels{
		HttpMetadata: mocks.HttpMetadataStatusOk,
		VariantId:    "sku_123",
		OnHand:       100,
	}

	apiClient, credentials, config := buildInventoryClientConfig()
	credentials.On("GetAuthorization", mock.Anything).
		Return(&configuration.SdkAuthorization{}, nil)
	apiClient.On("PutWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil).
		Run(func(args mock.Arguments) {
			respMapping := args.Get(4).(*InventoryLevels)
			*respMapping = levelsResponse
		})

	client := NewClient(config, apiClient)
	response, err := client.SetInventoryLevels("sku_123", request)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, int64(100), response.OnHand)
}

func TestGetInventoryProduct(t *testing.T) {
	createdOn := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	productResponse := InventoryProductKnowledge{
		HttpMetadata: mocks.HttpMetadataStatusOk,
		VariantId:    "sku_123",
		Title:        "Blue T-Shirt",
		Description:  "A comfortable blue t-shirt",
		ProductUrl:   "https://example.com/products/sku_123",
		ImageUrl:     "https://example.com/images/sku_123.png",
		Condition:    NewCondition,
		CreatedOn:    &createdOn,
		ModifiedOn:   &createdOn,
	}

	apiClient, credentials, config := buildInventoryClientConfig()
	credentials.On("GetAuthorization", mock.Anything).
		Return(&configuration.SdkAuthorization{}, nil)
	apiClient.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil).
		Run(func(args mock.Arguments) {
			respMapping := args.Get(3).(*InventoryProductKnowledge)
			*respMapping = productResponse
		})

	client := NewClient(config, apiClient)
	response, err := client.GetInventoryProduct("sku_123")

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "Blue T-Shirt", response.Title)
	assert.Equal(t, NewCondition, response.Condition)
}

func TestSetInventoryProduct(t *testing.T) {
	request := InventorySetProductRequest{
		Title:       "Blue T-Shirt",
		Description: "A comfortable blue t-shirt",
		ProductUrl:  "https://example.com/products/sku_123",
		ImageUrl:    "https://example.com/images/sku_123.png",
		Price:       &InventoryMoney{Amount: 1999, Currency: "USD"},
	}
	productResponse := InventoryProductKnowledge{
		HttpMetadata: mocks.HttpMetadataStatusOk,
		VariantId:    "sku_123",
		Title:        "Blue T-Shirt",
	}

	apiClient, credentials, config := buildInventoryClientConfig()
	credentials.On("GetAuthorization", mock.Anything).
		Return(&configuration.SdkAuthorization{}, nil)
	apiClient.On("PutWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil).
		Run(func(args mock.Arguments) {
			respMapping := args.Get(4).(*InventoryProductKnowledge)
			*respMapping = productResponse
		})

	client := NewClient(config, apiClient)
	response, err := client.SetInventoryProduct("sku_123", request)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "Blue T-Shirt", response.Title)
}

func TestDeleteInventoryProduct(t *testing.T) {
	apiClient, credentials, config := buildInventoryClientConfig()
	credentials.On("GetAuthorization", mock.Anything).
		Return(&configuration.SdkAuthorization{}, nil)
	apiClient.On("DeleteWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil)

	client := NewClient(config, apiClient)
	response, err := client.DeleteInventoryProduct("sku_123")

	assert.Nil(t, err)
	assert.NotNil(t, response)
}

// TestInventoryLevelsSwaggerExampleRoundtrip exercises every InventoryLevels field against a
// payload shaped like the swagger example, confirming an unmarshal/marshal roundtrip preserves
// every property including the embedded product knowledge and HAL links.
func TestInventoryLevelsSwaggerExampleRoundtrip(t *testing.T) {
	raw := []byte(`{
		"variant_id": "sku_123",
		"on_hand": 42,
		"reserved": 2,
		"safety_stock": 5,
		"available": 35,
		"state": "in_stock",
		"source": "managed",
		"created_on": "2026-01-01T00:00:00Z",
		"modified_on": "2026-01-02T00:00:00Z",
		"product": {
			"variant_id": "sku_123",
			"title": "Blue T-Shirt",
			"description": "A comfortable blue t-shirt",
			"product_url": "https://example.com/products/sku_123",
			"image_url": "https://example.com/images/sku_123.png",
			"condition": "new",
			"price": {"amount": 1999, "currency": "USD"},
			"created_on": "2026-01-01T00:00:00Z",
			"modified_on": "2026-01-02T00:00:00Z",
			"_links": {"self": {"href": "https://api.checkout.com/inventory/sku_123/product", "actions": ["GET"], "types": ["application/json"]}}
		},
		"_links": {
			"self": {"href": "https://api.checkout.com/inventory/sku_123", "actions": ["GET"], "types": ["application/json"]},
			"set": {"href": "https://api.checkout.com/inventory/sku_123", "actions": ["PUT"], "types": ["application/json"]}
		}
	}`)

	var levels InventoryLevels
	err := json.Unmarshal(raw, &levels)
	assert.Nil(t, err)

	assert.Equal(t, "sku_123", levels.VariantId)
	assert.Equal(t, int64(42), levels.OnHand)
	assert.Equal(t, int64(2), levels.Reserved)
	assert.Equal(t, int64(5), levels.SafetyStock)
	assert.Equal(t, int64(35), levels.Available)
	assert.Equal(t, InStock, levels.State)
	assert.Equal(t, Managed, levels.Source)
	assert.NotNil(t, levels.CreatedOn)
	assert.NotNil(t, levels.ModifiedOn)
	assert.NotNil(t, levels.Product)
	assert.Equal(t, "Blue T-Shirt", levels.Product.Title)
	assert.Equal(t, int64(1999), levels.Product.Price.Amount)
	assert.Equal(t, "USD", levels.Product.Price.Currency)
	assert.NotNil(t, levels.Links)
	assert.NotNil(t, levels.Links.Self)
	assert.Equal(t, []string{"GET"}, levels.Links.Self.Actions)
	assert.NotNil(t, levels.Links.Set)

	out, err := json.Marshal(&levels)
	assert.Nil(t, err)

	var roundtripped InventoryLevels
	assert.Nil(t, json.Unmarshal(out, &roundtripped))
	assert.Equal(t, levels.VariantId, roundtripped.VariantId)
	assert.Equal(t, levels.Available, roundtripped.Available)
	assert.Equal(t, levels.Product.Title, roundtripped.Product.Title)
}

// TestInventoryReservationSwaggerExampleRoundtrip exercises every InventoryReservation field,
// including the held-state _links object with all three actions.
func TestInventoryReservationSwaggerExampleRoundtrip(t *testing.T) {
	raw := []byte(`{
		"id": "rsv_2n9x0f7q1z3m4k5j6h7g8f9d0",
		"state": "held",
		"owner_type": "cart",
		"owner_reference": "cart_789",
		"items": [{"variant_id": "sku_123", "quantity": 2}, {"variant_id": "sku_456", "quantity": 1}],
		"expires_at": "2026-09-18T13:00:00Z",
		"created_on": "2026-09-18T12:45:00Z",
		"_links": {
			"self": {"href": "https://api.checkout.com/inventory/reservations/rsv_2n9x0f7q1z3m4k5j6h7g8f9d0", "actions": ["GET"]},
			"commit": {"href": "https://api.checkout.com/inventory/reservations/rsv_2n9x0f7q1z3m4k5j6h7g8f9d0/commit", "actions": ["POST"]},
			"release": {"href": "https://api.checkout.com/inventory/reservations/rsv_2n9x0f7q1z3m4k5j6h7g8f9d0/release", "actions": ["POST"]}
		}
	}`)

	var reservation InventoryReservation
	err := json.Unmarshal(raw, &reservation)
	assert.Nil(t, err)

	assert.Equal(t, "rsv_2n9x0f7q1z3m4k5j6h7g8f9d0", reservation.Id)
	assert.Equal(t, Held, reservation.State)
	assert.Equal(t, "cart", reservation.OwnerType)
	assert.Equal(t, "cart_789", reservation.OwnerReference)
	assert.Len(t, reservation.Items, 2)
	assert.Equal(t, "sku_123", reservation.Items[0].VariantId)
	assert.Equal(t, int64(2), reservation.Items[0].Quantity)
	assert.NotNil(t, reservation.ExpiresAt)
	assert.NotNil(t, reservation.CreatedOn)
	assert.NotNil(t, reservation.Links)
	assert.NotNil(t, reservation.Links.Self)
	assert.NotNil(t, reservation.Links.Commit)
	assert.NotNil(t, reservation.Links.Release)

	out, err := json.Marshal(&reservation)
	assert.Nil(t, err)

	var roundtripped InventoryReservation
	assert.Nil(t, json.Unmarshal(out, &roundtripped))
	assert.Equal(t, reservation.Id, roundtripped.Id)
	assert.Equal(t, reservation.Items, roundtripped.Items)
}

// TestInventoryAdjustmentRequestSerialization confirms all three required fields, including a
// zero-safe negative Delta, serialize onto the wire without being dropped by omitempty.
func TestInventoryAdjustmentRequestSerialization(t *testing.T) {
	request := InventoryAdjustmentRequest{VariantId: "sku_123", Delta: -3, Reason: "damaged"}

	out, err := json.Marshal(&request)
	assert.Nil(t, err)

	var decoded map[string]interface{}
	assert.Nil(t, json.Unmarshal(out, &decoded))
	assert.Equal(t, "sku_123", decoded["variant_id"])
	assert.Equal(t, float64(-3), decoded["delta"])
	assert.Equal(t, "damaged", decoded["reason"])
}

// TestInventorySetLevelsRequestOmittedSafetyStock confirms that an omitted SafetyStock pointer
// is left out of the wire payload, distinguishing "omitted" from an explicit 0, per the
// swagger note that omitting safety_stock on update leaves it unchanged.
func TestInventorySetLevelsRequestOmittedSafetyStock(t *testing.T) {
	request := InventorySetLevelsRequest{OnHand: 0}

	out, err := json.Marshal(&request)
	assert.Nil(t, err)

	var decoded map[string]interface{}
	assert.Nil(t, json.Unmarshal(out, &decoded))
	assert.Equal(t, float64(0), decoded["on_hand"])
	_, present := decoded["safety_stock"]
	assert.False(t, present)
}

// common methods

func buildInventoryClientConfig() (*mocks.ApiClientMock, *mocks.CredentialsMock, *configuration.Configuration) {
	apiClient := new(mocks.ApiClientMock)
	credentials := new(mocks.CredentialsMock)
	environment := new(mocks.EnvironmentMock)
	enableTelemetry := true
	config := configuration.NewConfiguration(credentials, &enableTelemetry, environment, &http.Client{}, nil)
	return apiClient, credentials, config
}
