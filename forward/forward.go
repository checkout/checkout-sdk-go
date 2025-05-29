package forward

import (
	"github.com/checkout/checkout-sdk-go/common"
	"time"
)

const (
	forward = "forward"
)

type SourceType string

const (
	IdST    SourceType = "id"
	TokenST SourceType = "token"
)

type (
	AbstractSource interface {
		GetType() SourceType
	}
)

type MethodType string

const (
	GetMT     MethodType = "GET"
	PostMT    MethodType = "POST"
	PutMT     MethodType = "PUT"
	DeleteMT  MethodType = "DELETE"
	PatchMT   MethodType = "PATCH"
	HeadMT    MethodType = "HEAD"
	OptionsMT MethodType = "OPTIONS"
	TraceMT   MethodType = "TRACE"
)

type (
	idSource struct {
		// The payment source type (Required)
		Type SourceType `json:"type"`
		// The unique identifier of the payment instrument (Required, pattern ^(src)_(\w{26})$)
		Id string `json:"id"`
		// The unique token for the card's security code. Checkout.com does not store a card's Card Verification Value
		// (CVV) with its associated payment instrument. To pass a CVV with your forward request, use the Frames SDK for
		// Android or iOS to collect and tokenize the CVV and pass the value in this field. The token will replace the
		// placeholder {{card_cvv}} value in destination_request.body (Optional, pattern ^(tok)_(\w{26})$)
		CvvToken string `json:"cvv_token,omitempty"`
	}

	tokenSource struct {
		// The payment source type (Required)
		Type SourceType `json:"type"`
		// The unique Checkout.com token (Required, pattern ^(tok)_(\w{26})$)
		Token string `json:"token"`
	}

	NetworkToken struct {
		// Specifies whether to use a network token (Optional)
		Enabled bool `json:"enabled,omitempty"`
		// Specifies whether to generate a cryptogram. For example, for customer-initiated transactions (CITs). If you
		// set network_token.enabled to true, you must provide this field (Optional)
		RequestCryptogram bool `json:"request_cryptogram,omitempty"`
	}

	Headers struct {
		// The raw headers to include in the forward request (Required, max 16 characters)
		Raw map[string]string `json:"raw"`
		// The encrypted headers to include in the forward request, as a JSON object with string values encrypted
		// with JSON Web Encryption (JWE) (Optional, max 8192 characters)
		Encrypted string `json:"encrypted,omitempty"`
	}

	DestinationRequest struct {
		// Url is the URL of the forward request. (Required: true)
		Url string `json:"url"`
		// Method is the HTTP method of the forward request. (Required: true)
		Method MethodType `json:"method"`
		// Headers are the HTTP headers of the forward request. (Required: true)
		Headers *Headers `json:"headers"`
		// Body is the HTTP message body of the forward request. (Required: true)
		Body string `json:"body"`
	}

	ForwardRequest struct {
		// The payment source to enrich the forward request with. You can provide placeholder values in
		// destination_request.body. The request will be enriched with the respective payment credentials from the token or
		// payment instrument you specified. For example, {{card_number}} (Required)
		Source AbstractSource `json:"source"`
		// The parameters of the forward request (Required)
		DestinationRequest *DestinationRequest `json:"destination_request"`
		// The unique reference for the forward request (Optional, max 80 characters)
		Reference string `json:"reference,omitempty"`
		// The processing channel ID to associate the billing for the forward request with (Optional,
		// pattern ^(pc)_(\w{26})$)
		ProcessingChannelId string `json:"processing_channel_id,omitempty"`
		// Specifies if and how a network token should be used in the forward request (Optional)
		NetworkToken *NetworkToken `json:"network_token,omitempty"`
	}
)

type (
	DestinationResponse struct {
		// The HTTP status code of the destination response (Required)
		Status int `json:"status"`
		// The destination response's HTTP headers. (Required)
		Headers map[string][]string `json:"headers"`
		// The destination response's HTTP message body (Required)
		Body string `json:"body"`
	}
)

type (
	ForwardAnApiResponse struct {
		HttpMetadata common.HttpMetadata `json:"http_metadata,omitempty"`
		// The unique identifier for the forward request (Required)
		RequestId string `json:"request_id"`
		// The HTTP response received from the destination, if the forward request completed successfully. Sensitive PCI
		// data will be removed from the response (Optional)
		DestinationResponse *DestinationResponse `json:"destination_response,omitempty"`
	}
)

type (
	DestinationRequestResponse struct {
		// Url is the URL of the forward request. (Required: true)
		Url string `json:"url"`
		// Method is the HTTP method of the forward request. (Required: true)
		Method string `json:"method"`
		// Headers are the HTTP headers of the forward request. Encrypted and sensitive header values are redacted. (Required: true)
		Headers map[string]string `json:"headers"`
		// Body is the HTTP message body of the forward request. This is the original value used to initiate the request, with placeholder value text included. For example, {{card_number}} is not replaced with an actual card number. (Required: true)
		Body string `json:"body"`
	}

	GetForwardResponse struct {
		HttpMetadata common.HttpMetadata `json:"http_metadata,omitempty"`
		// The unique identifier for the forward request (Required)
		RequestId string `json:"request_id"`
		// The client entity linked to the forward request (Required)
		EntityId string `json:"entity_id"`
		// The parameters of the HTTP request forwarded to the destination (Required)
		DestinationRequest *DestinationRequestResponse `json:"destination_request"`
		// The date and time the forward request was created, in UTC (Required)
		CreatedOn time.Time `json:"created_on"`
		// The unique reference for the forward request (Optional)
		Reference string `json:"reference,omitempty"`
		// The HTTP response received from the destination. Sensitive PCI data is not included in the response
		// (Optional)
		DestinationResponse *DestinationResponse `json:"destination_response,omitempty"`
	}
)

func NewIdSource() *idSource {
	return &idSource{Type: IdST}
}

func NewTokenSource() *tokenSource {
	return &tokenSource{Type: TokenST}
}

func (s *idSource) GetType() SourceType {
	return s.Type
}

func (s *tokenSource) GetType() SourceType {
	return s.Type
}
