package identities_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/checkout/checkout-sdk-go/v3/configuration"
	"github.com/checkout/checkout-sdk-go/v3/identities"
	"github.com/checkout/checkout-sdk-go/v3/identities/addressdocumentverification"
	"github.com/checkout/checkout-sdk-go/v3/identities/faceauthentication"
	"github.com/checkout/checkout-sdk-go/v3/identities/iddocumentverification"
	"github.com/checkout/checkout-sdk-go/v3/identities/identityverification"
	"github.com/checkout/checkout-sdk-go/v3/mocks"
)

// The field tests cover the response shapes and the query filter encoding in isolation. Nothing
// covered the join: that each client method composes the right path, and that the pagination
// filter reaches the query string. These assert the exact path handed to the ApiClient, which is
// the last thing the SDK controls before the request leaves.
//
// Lives in identities_test so it can import all four sub-packages without an import cycle.

func testConfig(t *testing.T) (*configuration.Configuration, *mocks.ApiClientMock) {
	t.Helper()
	credentials := &mocks.CredentialsMock{}
	credentials.On("GetAuthorization", mock.Anything).
		Return(&configuration.SdkAuthorization{}, nil)

	enableTelemetry := false
	config := configuration.NewConfiguration(
		credentials, &enableTelemetry, &mocks.EnvironmentMock{}, &http.Client{}, nil)

	apiClient := &mocks.ApiClientMock{}
	return config, apiClient
}

// expectGet stubs GetWithContext and captures the path it was called with.
func expectGet(apiClient *mocks.ApiClientMock, captured *string) {
	apiClient.On("GetWithContext", mock.Anything, mock.MatchedBy(func(path string) bool {
		*captured = path
		return true
	}), mock.Anything, mock.Anything).Return(nil)
}

func TestAddressDocumentVerificationClient_BuildsTheAttemptPaths(t *testing.T) {
	cases := []struct {
		name     string
		call     func(*addressdocumentverification.Client) error
		expected string
	}{
		{
			name: "attempts without a filter sends no query string",
			call: func(c *addressdocumentverification.Client) error {
				_, err := c.GetAddressDocumentVerificationAttempts("adv_1", identities.AttemptsQueryFilter{})
				return err
			},
			expected: "/address-document-verifications/adv_1/attempts",
		},
		{
			name: "attempts with skip and limit reaches the query string",
			call: func(c *addressdocumentverification.Client) error {
				_, err := c.GetAddressDocumentVerificationAttempts(
					"adv_1", identities.AttemptsQueryFilter{Skip: 5, Limit: 25})
				return err
			},
			expected: "/address-document-verifications/adv_1/attempts?limit=25&skip=5",
		},
		{
			name: "assets without a filter sends no query string",
			call: func(c *addressdocumentverification.Client) error {
				_, err := c.GetAddressDocumentVerificationAttemptAssets(
					"adv_1", "adva_1", identities.AttemptAssetsQueryFilter{})
				return err
			},
			expected: "/address-document-verifications/adv_1/attempts/adva_1/assets",
		},
		{
			name: "assets with skip and limit reaches the query string",
			call: func(c *addressdocumentverification.Client) error {
				_, err := c.GetAddressDocumentVerificationAttemptAssets(
					"adv_1", "adva_1", identities.AttemptAssetsQueryFilter{Skip: 2, Limit: 50})
				return err
			},
			expected: "/address-document-verifications/adv_1/attempts/adva_1/assets?limit=50&skip=2",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			config, apiClient := testConfig(t)
			var captured string
			expectGet(apiClient, &captured)

			assert.Nil(t, tc.call(addressdocumentverification.NewClient(config, apiClient)))
			assert.Equal(t, tc.expected, captured)
		})
	}
}

func TestIdDocumentVerificationClient_BuildsTheAttemptPaths(t *testing.T) {
	cases := []struct {
		name     string
		call     func(*iddocumentverification.Client) error
		expected string
	}{
		{
			name: "attempts without a filter",
			call: func(c *iddocumentverification.Client) error {
				_, err := c.GetIdDocumentVerificationAttempts("iddv_1", identities.AttemptsQueryFilter{})
				return err
			},
			expected: "/id-document-verifications/iddv_1/attempts",
		},
		{
			name: "attempts with pagination",
			call: func(c *iddocumentverification.Client) error {
				_, err := c.GetIdDocumentVerificationAttempts(
					"iddv_1", identities.AttemptsQueryFilter{Skip: 5, Limit: 25})
				return err
			},
			expected: "/id-document-verifications/iddv_1/attempts?limit=25&skip=5",
		},
		{
			name: "assets without a filter",
			call: func(c *iddocumentverification.Client) error {
				_, err := c.GetIdDocumentVerificationAttemptAssets(
					"iddv_1", "datp_1", identities.AttemptAssetsQueryFilter{})
				return err
			},
			expected: "/id-document-verifications/iddv_1/attempts/datp_1/assets",
		},
		{
			name: "assets with pagination",
			call: func(c *iddocumentverification.Client) error {
				_, err := c.GetIdDocumentVerificationAttemptAssets(
					"iddv_1", "datp_1", identities.AttemptAssetsQueryFilter{Skip: 2, Limit: 50})
				return err
			},
			expected: "/id-document-verifications/iddv_1/attempts/datp_1/assets?limit=50&skip=2",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			config, apiClient := testConfig(t)
			var captured string
			expectGet(apiClient, &captured)

			assert.Nil(t, tc.call(iddocumentverification.NewClient(config, apiClient)))
			assert.Equal(t, tc.expected, captured)
		})
	}
}

func TestIdentityVerificationClient_BuildsTheAttemptsPath(t *testing.T) {
	cases := []struct {
		name     string
		query    identities.AttemptsQueryFilter
		expected string
	}{
		{
			name:     "no filter",
			query:    identities.AttemptsQueryFilter{},
			expected: "/identity-verifications/idv_1/attempts",
		},
		{
			name:     "skip and limit",
			query:    identities.AttemptsQueryFilter{Skip: 5, Limit: 25},
			expected: "/identity-verifications/idv_1/attempts?limit=25&skip=5",
		},
		{
			name:     "limit only",
			query:    identities.AttemptsQueryFilter{Limit: 25},
			expected: "/identity-verifications/idv_1/attempts?limit=25",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			config, apiClient := testConfig(t)
			var captured string
			expectGet(apiClient, &captured)

			_, err := identityverification.NewClient(config, apiClient).
				GetIdentityVerificationAttempts("idv_1", tc.query)

			assert.Nil(t, err)
			assert.Equal(t, tc.expected, captured)
		})
	}
}

func TestFaceAuthenticationClient_BuildsTheAttemptsPath(t *testing.T) {
	cases := []struct {
		name     string
		query    identities.AttemptsQueryFilter
		expected string
	}{
		{
			name:     "no filter",
			query:    identities.AttemptsQueryFilter{},
			expected: "/face-authentications/fav_1/attempts",
		},
		{
			name:     "skip and limit",
			query:    identities.AttemptsQueryFilter{Skip: 5, Limit: 25},
			expected: "/face-authentications/fav_1/attempts?limit=25&skip=5",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			config, apiClient := testConfig(t)
			var captured string
			expectGet(apiClient, &captured)

			_, err := faceauthentication.NewClient(config, apiClient).
				GetFaceAuthenticationAttempts("fav_1", tc.query)

			assert.Nil(t, err)
			assert.Equal(t, tc.expected, captured)
		})
	}
}

// A zero Skip is dropped by url omitempty, so it cannot be expressed. Asserted at the path level
// so the limitation is visible where a caller would hit it, not just in the filter unit test.
func TestAttemptsPath_ZeroSkipIsNotExpressible(t *testing.T) {
	config, apiClient := testConfig(t)
	var captured string
	expectGet(apiClient, &captured)

	_, err := identityverification.NewClient(config, apiClient).
		GetIdentityVerificationAttempts("idv_1", identities.AttemptsQueryFilter{Skip: 0, Limit: 10})

	assert.Nil(t, err)
	assert.Equal(t, "/identity-verifications/idv_1/attempts?limit=10", captured)
}
