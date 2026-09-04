package nas

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/abc"
	"github.com/checkout/checkout-sdk-go/v3/configuration"
	"github.com/checkout/checkout-sdk-go/v3/mocks"
)

func testConfiguration() *configuration.Configuration {
	credentials := new(mocks.CredentialsMock)
	environment := new(mocks.EnvironmentMock)
	enableTelemetry := true
	environment.On("BaseUri").Return("https://api.sandbox.checkout.com")
	environment.On("FilesUri").Return("https://files.sandbox.checkout.com")
	environment.On("BalancesUri").Return("https://balances.sandbox.checkout.com")
	environment.On("TransfersUri").Return("https://transfers.sandbox.checkout.com")
	environment.On("ForwardUri").Return("https://forward.sandbox.checkout.com")
	environment.On("IdentityUri").Return("https://identity.sandbox.checkout.com")
	return configuration.NewConfiguration(credentials, &enableTelemetry, environment, &http.Client{}, nil)
}

// The client surface is otherwise untested, so a missing registration would leave Api.Bacs nil with
// the whole suite green.
func TestCheckoutApiRegistersTheBacsClient(t *testing.T) {
	api := CheckoutApi(testConfiguration())

	assert.NotNil(t, api)
	assert.NotNil(t, api.Bacs)
	assert.NotNil(t, api.Instruments)
	assert.NotNil(t, api.Sepa)
}

// POST /apms/bacs/notifications is current-platform, secret-key-only, so the client must not be
// reachable through the previous-platform Api. In Go that is a compile-time guarantee because the
// two Api structs are independent; this pins it so a future refactor cannot merge them silently.
func TestPreviousApiDoesNotExposeTheBacsClient(t *testing.T) {
	_, found := reflect.TypeOf(abc.Api{}).FieldByName("Bacs")
	assert.False(t, found)

	_, found = reflect.TypeOf(Api{}).FieldByName("Bacs")
	assert.True(t, found)
}
