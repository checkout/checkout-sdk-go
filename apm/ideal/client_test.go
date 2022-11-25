package ideal

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/checkout/checkout-sdk-go-beta/common"
	"github.com/checkout/checkout-sdk-go-beta/configuration"
	"github.com/checkout/checkout-sdk-go-beta/mocks"
)

func TestGetInfo(t *testing.T) {
	var (
		httpMetadata = common.HttpMetadata{
			Status:     "200 OK",
			StatusCode: http.StatusOK,
		}

		infoLinks = InfoLinks{
			Curies: []CuriesLink{
				{
					Name:      "test link",
					Href:      "https://test-link.com",
					Templated: false,
				},
			},
		}

		response = IdealInfo{
			HttpMetadata:   httpMetadata,
			IdealInfoLinks: infoLinks,
		}
	)

	cases := []struct {
		name             string
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*IdealInfo, error)
	}{
		{
			name: "when auth is correct then return info links",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(2).(*IdealInfo)
						*respMapping = response
					})
			},
			checker: func(response *IdealInfo, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.IdealInfoLinks)
				assert.NotNil(t, response.IdealInfoLinks.Curies)
				assert.Equal(t, infoLinks.Curies, response.IdealInfoLinks.Curies)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient := new(mocks.ApiClientMock)
			credentials := new(mocks.CredentialsMock)
			environment := new(mocks.EnvironmentMock)

			tc.getAuthorization(&credentials.Mock)
			tc.apiGet(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, environment, &http.Client{})
			client := NewClient(configuration, apiClient)

			tc.checker(client.GetInfo())
		})
	}
}
