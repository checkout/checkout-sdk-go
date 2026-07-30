package accounts

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/checkout/checkout-sdk-go/v2/configuration"
	"github.com/checkout/checkout-sdk-go/v2/mocks"
)

// acceptHeader extracts the resolved Accept header value from the Headers field of the
// payload/source struct passed to the mocked ApiClient.
func acceptHeader(arg interface{}) string {
	v := reflect.ValueOf(arg)
	h := v.FieldByName("Headers")
	if h.Kind() == reflect.Ptr {
		h = h.Elem()
	}
	return h.FieldByName("Accept").String()
}

func newSchemaVersionTestClient(apiClient *mocks.ApiClientMock) *Client {
	credentials := new(mocks.CredentialsMock)
	environment := new(mocks.EnvironmentMock)
	credentials.On("GetAuthorization", mock.Anything).Return(&configuration.SdkAuthorization{}, nil)
	enableTelemetry := true
	config := configuration.NewConfiguration(credentials, &enableTelemetry, environment, &http.Client{}, nil)
	return NewClient(config, apiClient, apiClient)
}

func TestCreateEntitySendsSchemaVersionHeader(t *testing.T) {
	apiClient := new(mocks.ApiClientMock)
	var captured interface{}
	apiClient.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil).
		Run(func(args mock.Arguments) { captured = args.Get(3) })

	_, _ = newSchemaVersionTestClient(apiClient).CreateEntity(OnboardEntityRequest{}, "3.0")

	assert.Equal(t, "application/json;schema_version=3.0", acceptHeader(captured))
}

func TestGetEntitySendsSchemaVersionHeader(t *testing.T) {
	apiClient := new(mocks.ApiClientMock)
	apiClient.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	_, _ = newSchemaVersionTestClient(apiClient).GetEntity("ent_1234", "3.0")

	assert.Equal(t, "application/json;schema_version=3.0", acceptHeader(apiClient.CapturedGetRequest))
}

func TestUpdateEntitySendsSchemaVersionHeader(t *testing.T) {
	apiClient := new(mocks.ApiClientMock)
	var captured interface{}
	apiClient.On("PutWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil).
		Run(func(args mock.Arguments) { captured = args.Get(3) })

	_, _ = newSchemaVersionTestClient(apiClient).UpdateEntity("ent_1234", OnboardEntityRequest{}, "3.0")

	assert.Equal(t, "application/json;schema_version=3.0", acceptHeader(captured))
}

func TestGetEntityRequirementsSendsSchemaVersionHeader(t *testing.T) {
	apiClient := new(mocks.ApiClientMock)
	apiClient.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	_, _ = newSchemaVersionTestClient(apiClient).GetEntityRequirements("ent_1234", "3.0")

	assert.Equal(t, "application/json;schema_version=3.0", acceptHeader(apiClient.CapturedGetRequest))
}

func TestGetEntityHonorsSchemaVersionOverride(t *testing.T) {
	apiClient := new(mocks.ApiClientMock)
	apiClient.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	_, _ = newSchemaVersionTestClient(apiClient).GetEntity("ent_1234", "2.0")

	assert.Equal(t, "application/json;schema_version=2.0", acceptHeader(apiClient.CapturedGetRequest))
}

func TestEmptySchemaVersionFallsBackToDefault(t *testing.T) {
	apiClient := new(mocks.ApiClientMock)
	var captured interface{}
	apiClient.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil).
		Run(func(args mock.Arguments) { captured = args.Get(3) })

	_, _ = newSchemaVersionTestClient(apiClient).CreateEntity(OnboardEntityRequest{}, "")

	assert.Equal(t, "application/json;schema_version="+DefaultSchemaVersion, acceptHeader(captured))
}
