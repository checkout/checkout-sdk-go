package cardholdertokens

import (
	"bytes"
	stderrors "errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/checkout/checkout-sdk-go/v3/client"
	"github.com/checkout/checkout-sdk-go/v3/configuration"
	"github.com/checkout/checkout-sdk-go/v3/errors"
)

// cardholderMarker is in the request form; the test checks that it reaches the API host only and
// is not in the error or the log.
const cardholderMarker = "cardholder-form-marker"

// TestRequestCardholderTokenResponsePolicy: the request goes through the API client policy: the
// form reaches the API host only, the second host receives no request, and the error is a
// CheckoutAPIError with the status code only.
func TestRequestCardholderTokenResponsePolicy(t *testing.T) {
	for _, code := range []int{http.StatusTemporaryRedirect, http.StatusPermanentRedirect} {
		t.Run(fmt.Sprint(code), func(t *testing.T) {
			var mu sync.Mutex
			var targetHits int
			var targetBodies, apiBodies []string
			target := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				body, _ := ioutil.ReadAll(r.Body)
				mu.Lock()
				targetHits++
				targetBodies = append(targetBodies, string(body))
				mu.Unlock()
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{"access_token":"from-target","token_type":"Bearer"}`))
			}))
			t.Cleanup(target.Close)
			api := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				body, _ := ioutil.ReadAll(r.Body)
				mu.Lock()
				apiBodies = append(apiBodies, string(body))
				mu.Unlock()
				w.Header().Set("Cko-Request-Id", "req-cardholder")
				w.Header().Set("Location", target.URL+"/issuing/access/connect/token")
				w.WriteHeader(code)
			}))
			t.Cleanup(api.Close)

			telemetry := false
			var logs bytes.Buffer
			config := configuration.NewConfiguration(
				configuration.NewDefaultKeysSdkCredentials("sk_sbox_cardholder_test", "pk_sbox_cardholder_test"),
				&telemetry, nil, nil, log.New(&logs, "", 0))
			tokens := NewClient(config, client.NewApiClient(config, api.URL))

			response, err := tokens.RequestCardholderToken(CardholderTokenRequest{
				GrantType:    "client_credentials",
				ClientId:     "ack_redirect_test",
				ClientSecret: cardholderMarker,
				CardholderId: "crh_redirect_test",
			})

			assert.Nil(t, response)
			require.Error(t, err)
			var apiErr errors.CheckoutAPIError
			require.True(t, stderrors.As(err, &apiErr), "want a CheckoutAPIError, got %T: %v", err, err)
			assert.Equal(t, code, apiErr.StatusCode)
			require.NotNil(t, apiErr.Data)
			assert.Equal(t, "req-cardholder", apiErr.Data.RequestID)
			assert.NotContains(t, fmt.Sprintf("%v %+v %+v", err, apiErr, *apiErr.Data), cardholderMarker)
			assert.NotContains(t, fmt.Sprintf("%v %+v", err, apiErr), target.URL)
			assert.NotContains(t, logs.String(), cardholderMarker)

			mu.Lock()
			defer mu.Unlock()
			require.Len(t, apiBodies, 1)
			assert.True(t, strings.Contains(apiBodies[0], cardholderMarker), "the API receives the form")
			assert.Equal(t, 0, targetHits, "the other host receives no request")
			assert.Empty(t, targetBodies)
		})
	}
}
