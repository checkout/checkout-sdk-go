package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/errors"
	"github.com/checkout/checkout-sdk-go/sessions"
	"github.com/checkout/checkout-sdk-go/sessions/channels"
	"github.com/checkout/checkout-sdk-go/sessions/completion"
	"github.com/checkout/checkout-sdk-go/sessions/sources"
)

func TestRequestSession(t *testing.T) {
	cases := []struct {
		name    string
		request sessions.SessionRequest
		checker func(*sessions.SessionResponse, error)
	}{
		{
			name: "when browser request is valid then create a browser session",
			request: getNonHostedSession(
				getBrowserChannel(),
				sessions.Payment,
				common.NoPreference,
				sessions.GoodsService,
			),
			checker: func(response *sessions.SessionResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusCreated, response.Created.HttpMetadata.StatusCode)
				assert.Nil(t, response.Accepted)
				assert.NotNil(t, response.Created)
				assert.NotEmpty(t, response.Created.Id)
				assert.NotEmpty(t, response.Created.SessionSecret)
			},
		},
		{
			name: "when app request is valid then create an app session",
			request: getNonHostedSession(
				getAppChannel(),
				sessions.Payment,
				common.NoPreference,
				sessions.GoodsService,
			),
			checker: func(response *sessions.SessionResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusCreated, response.Created.HttpMetadata.StatusCode)
				assert.Nil(t, response.Accepted)
				assert.NotNil(t, response.Created)
				assert.NotEmpty(t, response.Created.Id)
				assert.NotEmpty(t, response.Created.SessionSecret)
			},
		},
	}

	client := OAuthApi().Sessions

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.RequestSession(tc.request))
		})
	}
}

func TestGetSessionDetails(t *testing.T) {
	session := createSession(t, getNonHostedSession(getBrowserChannel(),
		sessions.Payment,
		common.NoPreference,
		sessions.GoodsService))

	cases := []struct {
		name          string
		sessionId     string
		sessionSecret string
		checker       func(*sessions.GetSessionResponse, error)
	}{
		{
			name:      "when session exists return session details - with OAuth",
			sessionId: session.Created.Id,
			checker: func(response *sessions.GetSessionResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, session.Created.Id, response.Id)
				assert.Equal(t, session.Created.SessionSecret, response.SessionSecret)
			},
		},
		{
			name:          "when session exists return session details - with session secret",
			sessionId:     session.Created.Id,
			sessionSecret: session.Created.SessionSecret,
			checker: func(response *sessions.GetSessionResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, session.Created.Id, response.Id)
			},
		},
		{
			name:      "when session is inexistent then return error",
			sessionId: "not_found",
			checker: func(response *sessions.GetSessionResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				errChk := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, errChk.StatusCode)
			},
		},
	}

	client := OAuthApi().Sessions

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetSessionDetails(tc.sessionId, tc.sessionSecret))
		})
	}
}

func TestUpdateSession(t *testing.T) {
	session := createSession(t, getHostedSession()).Accepted

	cases := []struct {
		name          string
		sessionId     string
		request       channels.Channel
		sessionSecret string
		checker       func(*sessions.GetSessionResponse, error)
	}{
		{
			name:      "when updating to browser session with valid data then update correctly - with OAuth",
			sessionId: createSession(t, getHostedSession()).Accepted.Id,
			request:   getBrowserChannel(),
			checker: func(response *sessions.GetSessionResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotEmpty(t, response.Id)
				assert.NotEmpty(t, response.SessionSecret)
			},
		},
		{
			name:          "when updating to browser session with valid data then update correctly - with Session Secret",
			sessionId:     session.Id,
			request:       getBrowserChannel(),
			sessionSecret: session.SessionSecret,
			checker: func(response *sessions.GetSessionResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotEmpty(t, response.Id)
			},
		},
		{
			name:      "when updating inexistent session then return error",
			sessionId: "not_found",
			request:   getBrowserChannel(),
			checker: func(response *sessions.GetSessionResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				errChk := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, errChk.StatusCode)
			},
		},
	}

	client := OAuthApi().Sessions

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.UpdateSession(tc.sessionId, tc.request, tc.sessionSecret))
		})
	}
}

func TestUpdate3dsMethodCompletion(t *testing.T) {
	session := createSession(t, getHostedSession()).Accepted

	cases := []struct {
		name          string
		sessionId     string
		request       sessions.ThreeDsMethodCompletionRequest
		sessionSecret string
		checker       func(*sessions.Update3dsMethodCompletionResponse, error)
	}{
		{
			name:      "when updating 3ds method completion valid data then update correctly - with OAuth",
			sessionId: createSession(t, getHostedSession()).Accepted.Id,
			request: sessions.ThreeDsMethodCompletionRequest{
				ThreeDsMethodCompletion: common.Y,
			},
			checker: func(response *sessions.Update3dsMethodCompletionResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotEmpty(t, response.Id)
				assert.NotEmpty(t, response.SessionSecret)
			},
		},
		{
			name:      "when updating to browser session with valid data then update correctly - with Session Secret",
			sessionId: session.Id,
			request: sessions.ThreeDsMethodCompletionRequest{
				ThreeDsMethodCompletion: common.Y,
			},
			sessionSecret: session.SessionSecret,
			checker: func(response *sessions.Update3dsMethodCompletionResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotEmpty(t, response.Id)
			},
		},
		{
			name:      "when updating inexistent session then return error",
			sessionId: "not_found",
			request: sessions.ThreeDsMethodCompletionRequest{
				ThreeDsMethodCompletion: common.Y,
			},
			checker: func(response *sessions.Update3dsMethodCompletionResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				errChk := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, errChk.StatusCode)
			},
		},
	}

	client := OAuthApi().Sessions

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.Update3dsMethodCompletion(tc.sessionId, tc.request, tc.sessionSecret))
		})
	}
}

func TestCompleteSession(t *testing.T) {
	session := createSession(t, getHostedSession()).Accepted

	cases := []struct {
		name          string
		sessionId     string
		sessionSecret string
		checker       func(*common.MetadataResponse, error)
	}{
		{
			name:      "when trying to complete existing hosted session it returns error - with OAuth",
			sessionId: session.Id,
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				errChk := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusForbidden, errChk.StatusCode)
				assert.Equal(t, "update_not_allowed_due_to_state", errChk.Data.ErrorCodes[0])
			},
		},
		{
			name:          "when trying to complete existing hosted session it returns error - with Session Secret",
			sessionId:     session.Id,
			sessionSecret: session.SessionSecret,
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				errChk := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusForbidden, errChk.StatusCode)
				assert.Equal(t, "update_not_allowed_due_to_state", errChk.Data.ErrorCodes[0])
			},
		},
	}

	client := OAuthApi().Sessions

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.CompleteSession(tc.sessionId, tc.sessionSecret))
		})
	}
}

func getBrowserChannel() channels.Channel {
	c := channels.NewBrowserSession()
	c.AcceptHeader = "Accept:  *.*, q=0.1"
	c.JavaEnabled = true
	c.Language = "FR-fr"
	c.ColorDepth = "16"
	c.ScreenWidth = "1920"
	c.ScreenHeight = "1080"
	c.Timezone = "60"
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0 Win64 x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36 "
	c.ThreeDsMethodCompletion = common.Y
	c.IpAddress = "1.12.123.255"

	return c
}

func getAppChannel() channels.Channel {
	c := channels.NewAppSession()
	c.SdkAppId = "dbd64fcb-c19a-4728-8849-e3d50bfdde39"
	c.SdkMaxTimeout = 5
	c.SdkEncryptedData = "{}"
	c.SdkEphemPubKey = &channels.SdkEphemeralPublicKey{
		Kty: "EC",
		Crv: "P-256",
		X:   "f83OJ3D2xF1Bg8vub9tLe1gHMzV76e8Tus9uPHvRVEU",
		Y:   "x_FEzRu9m36HLN_tue659LNpXW6pCyStikYjKIWI5a0",
	}
	c.SdkReferenceNumber = "3DS_LOA_SDK_PPFU_020100_00007"
	c.SdkTransactionId = "b2385523-a66c-4907-ac3c-91848e8c0067"
	c.SdkInterfaceType = channels.Both
	c.SdkUiElements = []channels.UIElements{channels.SingleSelect, channels.HtmlOther}

	return c
}

func getNonHostedSession(
	channel channels.Channel,
	category sessions.Category,
	indicator common.ChallengeIndicator,
	transaction sessions.TransactionType,
) sessions.SessionRequest {
	sessionAddress := &sources.SessionAddress{
		Address: common.Address{
			AddressLine1: "Checkout.com",
			AddressLine2: "ABC building",
			City:         "London",
			State:        "ENG",
			Zip:          "WIT 4JT",
			Country:      common.GB,
		},
		AddressLine3: "14 Wells Mews",
	}

	phone := &common.Phone{
		CountryCode: "44",
		Number:      "020222333",
	}

	cardSource := sources.NewSessionCardSource()
	cardSource.BillingAddress = sessionAddress
	cardSource.Number = CardNumber
	cardSource.ExpiryMonth = ExpiryMonth
	cardSource.ExpiryYear = ExpiryYear
	cardSource.Name = "John Doe"
	cardSource.Email = GenerateRandomEmail()
	cardSource.HomePhone = phone
	cardSource.WorkPhone = phone
	cardSource.MobilePhone = phone

	completionInfo := completion.NewNonHostedCompletion()
	completionInfo.CallbackUrl = "https://merchant.com/callback"

	s := sessions.SessionRequest{
		Source:              cardSource,
		Amount:              6540,
		Currency:            common.USD,
		ProcessingChannelId: "pc_5jp2az55l3cuths25t5p3xhwru",
		Marketplace: &sessions.SessionMarketplaceData{
			SubEntityId: "ent_ocw5i74vowfg2edpy66izhts2u",
		},
		AuthenticationCategory: category,
		ChallengeIndicator:     indicator,
		BillingDescriptor: &sessions.SessionsBillingDescriptor{
			Name: "SUPERHEROES.COM",
		},
		Reference:       "ORD-5023-4E89",
		TransactionType: transaction,
		ShippingAddress: sessionAddress,
		Completion:      completionInfo,
		ChannelData:     channel,
	}

	return s
}

func getHostedSession() sessions.SessionRequest {
	sessionAddress := &sources.SessionAddress{
		Address: common.Address{
			AddressLine1: "Checkout.com",
			AddressLine2: "ABC building",
			City:         "London",
			State:        "ENG",
			Zip:          "WIT 4JT",
			Country:      common.GB,
		},
		AddressLine3: "14 Wells Mews",
	}

	phone := &common.Phone{
		CountryCode: "44",
		Number:      "020222333",
	}

	cardSource := sources.NewSessionCardSource()
	cardSource.BillingAddress = sessionAddress
	cardSource.Number = CardNumber
	cardSource.ExpiryMonth = ExpiryMonth
	cardSource.ExpiryYear = ExpiryYear
	cardSource.Name = "John Doe"
	cardSource.Email = GenerateRandomEmail()
	cardSource.HomePhone = phone
	cardSource.WorkPhone = phone
	cardSource.MobilePhone = phone

	hostedCompletionInfo := completion.NewHostedCompletion()
	hostedCompletionInfo.SuccessUrl = "https://example.com/sessions/success"
	hostedCompletionInfo.FailureUrl = "https://example.com/sessions/fail"

	sessionRequest := sessions.SessionRequest{}
	sessionRequest.Source = cardSource
	sessionRequest.Amount = 100
	sessionRequest.Currency = common.USD
	sessionRequest.ProcessingChannelId = "pc_5jp2az55l3cuths25t5p3xhwru"
	sessionRequest.AuthenticationType = sessions.RegularAuthType
	sessionRequest.AuthenticationCategory = sessions.Payment
	sessionRequest.ChallengeIndicator = common.NoPreference
	sessionRequest.Reference = "ORD-5023-4E89"
	sessionRequest.TransactionType = sessions.GoodsService
	sessionRequest.ShippingAddress = sessionAddress
	sessionRequest.Completion = hostedCompletionInfo

	return sessionRequest
}

func createSession(t *testing.T, req sessions.SessionRequest) *sessions.SessionResponse {
	s, err := OAuthApi().Sessions.RequestSession(req)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("Error creating session - %s", err.Error()))
	}

	return s
}
