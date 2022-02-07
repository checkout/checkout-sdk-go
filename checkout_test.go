package checkout

import (
	"github.com/google/uuid"
	assert "github.com/stretchr/testify/require"
	"testing"
)

func TestCreateSdkConfigDifferentKeyPatterns_Sandbox(t *testing.T) {
	secretKeys := []string{"sk_test_fde517a8-3f01-41ef-b4bd-4282384b0a64", "sk_sbox_m73dzbpy7cf3gfd46xr4yj5xo4e", uuid.New().String()}
	publicKeys := []string{"pk_test_fde517a8-3f01-41ef-b4bd-4282384b0a64", "pk_sbox_m73dzbpy7cf3gfd46xr4yj5xo4e", uuid.New().String()}
	for _, sk := range secretKeys {
		for _, pk := range publicKeys {
			config, err := SdkConfig(&sk, &pk, Sandbox)
			assert.Nil(t, err)
			assert.NotNil(t, config)
			assert.Equal(t, sandboxURI, *config.URI)
			assert.Equal(t, sk, config.SecretKey)
			assert.Equal(t, pk, config.PublicKey)
			assert.NotNil(t, config.HTTPClient)
			assert.NotNil(t, config.LeveledLogger)
			assert.NotNil(t, config.MaxNetworkRetries)
		}
	}
}

func TestCreateSdkConfigDifferentKeyPatterns_Production(t *testing.T) {
	secretKeys := []string{"sk_fde517a8-3f01-41ef-b4bd-4282384b0a64", "sk_m73dzbpy7cf3gfd46xr4yj5xo4e", uuid.New().String()}
	publicKeys := []string{"pk_fde517a8-3f01-41ef-b4bd-4282384b0a64", "pk_m73dzbpy7cf3gfd46xr4yj5xo4e", uuid.New().String()}
	for _, sk := range secretKeys {
		for _, pk := range publicKeys {
			config, err := SdkConfig(&sk, &pk, Production)
			assert.Nil(t, err)
			assert.NotNil(t, config)
			assert.Equal(t, productionURI, *config.URI)
			assert.Equal(t, sk, config.SecretKey)
			assert.Equal(t, pk, config.PublicKey)
			assert.NotNil(t, config.HTTPClient)
			assert.NotNil(t, config.LeveledLogger)
			assert.NotNil(t, config.MaxNetworkRetries)
		}
	}
}

func TestCreateSdkConfigUnknownKeyPatterns_Sandbox(t *testing.T) {
	secretKeys := []string{"sk_fde517a8-3f01-41ef", "cko_123_m73dzbpy7cf3gfd46xr4yj5xo4e", uuid.New().String()}
	publicKeys := []string{"pk_-b4bd-4282384b0a64", "cko_m73dzbpy7cf3gfd46xr4yj5xo4e", uuid.New().String()}
	for _, sk := range secretKeys {
		for _, pk := range publicKeys {
			config, err := SdkConfig(&sk, &pk, Production)
			assert.Nil(t, err)
			assert.NotNil(t, config)
			assert.Equal(t, productionURI, *config.URI)
			assert.Equal(t, sk, config.SecretKey)
			assert.Equal(t, pk, config.PublicKey)
			assert.NotNil(t, config.HTTPClient)
			assert.NotNil(t, config.LeveledLogger)
			assert.NotNil(t, config.MaxNetworkRetries)
			assert.False(t, config.BearerAuthentication)
		}
	}
}

func TestAssignBearerAuthenticationDifferentKeyPatterns(t *testing.T) {
	validFourSecretKeys := []string{"sk_sbox_m73dzbpy7cf3gfd46xr4yj5xo4e", "sk_m73dzbpy7cf3gfd46xr4yj5xo4e",
		"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aWxkZXIiLCJpYXQiOjE2NDQ0MTYzMTYsImV4cCI6MTY3NTk1MjMxNiwiYXVkIjoid3d3LmV4YW1wbGUuY29tIiwic3ViIjoianJvY2tldEBleGFtcGxlLmNvbSIsIkdpdmVuTmFtZSI6IkpvaG5ueSIsIlN1cm5hbWUiOiJSb2NrZXQiLCJFbWFpbCI6Impyb2NrZXRAZXhhbXBsZS5jb20iLCJSb2xlIjpbIk1hbmFnZXIiLCJQcm9qZWN0IEFkbWluaXN0cmF0b3IiXX0.n53cqgfiUp8q9TTNCOt43EG5IXqaL9rhqblj63OKifU"}
	for _, sk := range validFourSecretKeys {
		config, err := SdkConfig(&sk, nil, Production)
		assert.Nil(t, err)
		assert.NotNil(t, config)
		assert.Equal(t, productionURI, *config.URI)
		assert.Equal(t, sk, config.SecretKey)
		assert.Equal(t, "", config.PublicKey)
		assert.NotNil(t, config.HTTPClient)
		assert.NotNil(t, config.LeveledLogger)
		assert.NotNil(t, config.MaxNetworkRetries)
		assert.True(t, config.BearerAuthentication)
	}
	validMbcPublicKeys := []string{"pk_fde517a8-3f01-41ef-b4bd-4282384b0a64", "pk_test_fde517a8-3f01-41ef-b4bd-4282384b0a64"}
	for _, pk := range validMbcPublicKeys {
		config, err := SdkConfig(nil, &pk, Sandbox)
		assert.Nil(t, err)
		assert.NotNil(t, config)
		assert.Equal(t, sandboxURI, *config.URI)
		assert.Equal(t, "", config.SecretKey)
		assert.Equal(t, pk, config.PublicKey)
		assert.NotNil(t, config.HTTPClient)
		assert.NotNil(t, config.LeveledLogger)
		assert.NotNil(t, config.MaxNetworkRetries)
		assert.False(t, config.BearerAuthentication)
	}
	invalidOAuthJwts := []string{"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJFkpva2NNCOt43EG5IXqaL9rhqblj63OKifU", "123123213.adasdsadasdasdasdas"}
	for _, sk := range invalidOAuthJwts {
		config, err := SdkConfig(&sk, nil, Sandbox)
		assert.Nil(t, err)
		assert.NotNil(t, config)
		assert.Equal(t, sandboxURI, *config.URI)
		assert.Equal(t, sk, config.SecretKey)
		assert.Equal(t, "", config.PublicKey)
		assert.NotNil(t, config.HTTPClient)
		assert.NotNil(t, config.LeveledLogger)
		assert.NotNil(t, config.MaxNetworkRetries)
		assert.False(t, config.BearerAuthentication)
	}
}
