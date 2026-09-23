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
			name: "attempts without a query sends no query string",
			call: func(c *addressdocumentverification.Client) error {
				_, err := c.GetAddressDocumentVerificationAttempts("adv_1")
				return err
			},
			expected: "/address-document-verifications/adv_1/attempts",
		},
		{
			name: "attempts with skip and limit reaches the query string",
			call: func(c *addressdocumentverification.Client) error {
				_, err := c.GetAddressDocumentVerificationAttemptsQuery(
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
				_, err := c.GetIdDocumentVerificationAttempts("iddv_1")
				return err
			},
			expected: "/id-document-verifications/iddv_1/attempts",
		},
		{
			name: "attempts with pagination",
			call: func(c *iddocumentverification.Client) error {
				_, err := c.GetIdDocumentVerificationAttemptsQuery(
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
				GetIdentityVerificationAttemptsQuery("idv_1", tc.query)

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
				GetFaceAuthenticationAttemptsQuery("fav_1", tc.query)

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
		GetIdentityVerificationAttemptsQuery("idv_1", identities.AttemptsQueryFilter{Skip: 0, Limit: 10})

	assert.Nil(t, err)
	assert.Equal(t, "/identity-verifications/idv_1/attempts?limit=10", captured)
}

// The 2026-09-02 row added pagination to these four endpoints. The query filter went onto a
// separate Query method rather than onto the existing method, so that callers written against
// the previous signature keep compiling. That is the same choice UpdateCardHeaders makes.
//
// These assert the contract that makes the split safe: the no-argument method has to produce
// exactly the path it produced before pagination existed, which is also what an empty filter
// produces. If BuildQueryPath ever started appending a bare "?" for an empty filter, or the
// delegation were wired to a non-empty default, this is what would catch it.
func TestAttemptsPath_NoQueryMatchesAnEmptyFilter(t *testing.T) {
	cases := []struct {
		name     string
		noQuery  func(*configuration.Configuration, *mocks.ApiClientMock) error
		withZero func(*configuration.Configuration, *mocks.ApiClientMock) error
		expected string
	}{
		{
			name: "address document verification",
			noQuery: func(c *configuration.Configuration, a *mocks.ApiClientMock) error {
				_, err := addressdocumentverification.NewClient(c, a).
					GetAddressDocumentVerificationAttempts("adv_1")
				return err
			},
			withZero: func(c *configuration.Configuration, a *mocks.ApiClientMock) error {
				_, err := addressdocumentverification.NewClient(c, a).
					GetAddressDocumentVerificationAttemptsQuery("adv_1", identities.AttemptsQueryFilter{})
				return err
			},
			expected: "/address-document-verifications/adv_1/attempts",
		},
		{
			name: "ID document verification",
			noQuery: func(c *configuration.Configuration, a *mocks.ApiClientMock) error {
				_, err := iddocumentverification.NewClient(c, a).
					GetIdDocumentVerificationAttempts("iddv_1")
				return err
			},
			withZero: func(c *configuration.Configuration, a *mocks.ApiClientMock) error {
				_, err := iddocumentverification.NewClient(c, a).
					GetIdDocumentVerificationAttemptsQuery("iddv_1", identities.AttemptsQueryFilter{})
				return err
			},
			expected: "/id-document-verifications/iddv_1/attempts",
		},
		{
			name: "identity verification",
			noQuery: func(c *configuration.Configuration, a *mocks.ApiClientMock) error {
				_, err := identityverification.NewClient(c, a).
					GetIdentityVerificationAttempts("idv_1")
				return err
			},
			withZero: func(c *configuration.Configuration, a *mocks.ApiClientMock) error {
				_, err := identityverification.NewClient(c, a).
					GetIdentityVerificationAttemptsQuery("idv_1", identities.AttemptsQueryFilter{})
				return err
			},
			expected: "/identity-verifications/idv_1/attempts",
		},
		{
			name: "face authentication",
			noQuery: func(c *configuration.Configuration, a *mocks.ApiClientMock) error {
				_, err := faceauthentication.NewClient(c, a).
					GetFaceAuthenticationAttempts("fav_1")
				return err
			},
			withZero: func(c *configuration.Configuration, a *mocks.ApiClientMock) error {
				_, err := faceauthentication.NewClient(c, a).
					GetFaceAuthenticationAttemptsQuery("fav_1", identities.AttemptsQueryFilter{})
				return err
			},
			expected: "/face-authentications/fav_1/attempts",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			config, apiClient := testConfig(t)
			var noQueryPath string
			expectGet(apiClient, &noQueryPath)
			assert.Nil(t, tc.noQuery(config, apiClient))

			config2, apiClient2 := testConfig(t)
			var emptyFilterPath string
			expectGet(apiClient2, &emptyFilterPath)
			assert.Nil(t, tc.withZero(config2, apiClient2))

			assert.Equal(t, tc.expected, noQueryPath)
			assert.Equal(t, noQueryPath, emptyFilterPath)
		})
	}
}
