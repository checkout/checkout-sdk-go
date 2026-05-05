package cardholdertokens

import "github.com/checkout/checkout-sdk-go/v2/common"

const cardholderTokenPath = "issuing/access/connect/token"

type CardholderTokenRequest struct {
	GrantType    string `url:"grant_type"`
	ClientId     string `url:"client_id"`
	ClientSecret string `url:"client_secret"`
	CardholderId string `url:"cardholder_id"`
	SingleUse    bool   `url:"single_use,omitempty"`
}

type CardholderTokenResponse struct {
	HttpMetadata common.HttpMetadata
	AccessToken  string  `json:"access_token,omitempty"`
	TokenType    string  `json:"token_type,omitempty"`
	ExpiresIn    float64 `json:"expires_in,omitempty"`
	Scope        string  `json:"scope,omitempty"`
}
