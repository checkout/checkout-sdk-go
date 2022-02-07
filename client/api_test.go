package client

import (
	"github.com/checkout/checkout-sdk-go"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCreateCheckoutApi(t *testing.T) {

	pk := os.Getenv("CHECKOUT_PUBLIC_KEY")
	sk := os.Getenv("CHECKOUT_SECRET_KEY")

	checkoutApi := CheckoutApi(&sk, &pk, checkout.Sandbox) // or Production
	assert.NotNil(t, checkoutApi)
	assert.NotNil(t, checkoutApi.Payments)
	assert.NotNil(t, checkoutApi.Sources)
	assert.NotNil(t, checkoutApi.Tokens)
	assert.NotNil(t, checkoutApi.Events)
	assert.NotNil(t, checkoutApi.Webhooks)
	assert.NotNil(t, checkoutApi.Disputes)
	assert.NotNil(t, checkoutApi.Files)
	assert.NotNil(t, checkoutApi.Reconciliation)
	assert.NotNil(t, checkoutApi.Instruments)

}
