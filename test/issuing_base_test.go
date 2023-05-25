package test

import (
	"os"

	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/checkout/checkout-sdk-go/issuing"
	"github.com/checkout/checkout-sdk-go/nas"
)

var (
	issuingClientApi *nas.Api
	cardholder       = issuing.CardholderRequest{
		Type:             issuing.Individual,
		Reference:        "X-123456-N11",
		EntityId:         "ent_mujh2nia2ypezmw5fo2fofk7ka",
		FirstName:        "John",
		MiddleName:       "Fitzgerald",
		LastName:         "Kennedy",
		Email:            "john.kennedy@myemaildomain.com",
		PhoneNumber:      Phone(),
		DateOfBirth:      "1985-05-15",
		BillingAddress:   Address(),
		ResidencyAddress: Address(),
		Document: &issuing.CardholderDocument{
			Type:            "national_identity_card",
			FrontDocumentId: "file_6lbss42ezvoufcb2beo76rvwly",
			BackDocumentId:  "file_aaz5pemp6326zbuvevp6qroqu4",
		},
	}
	cardholderResponse = cardholderRequest(cardholder)

	virtualCardResponse = cardRequest(cardholderResponse)

	virtualCardId = virtualCardResponse.Id
)

func buildIssuingClientApi() *nas.Api {
	if issuingClientApi == nil {
		issuingClientApi, _ = checkout.Builder().OAuth().
			WithClientCredentials(
				os.Getenv("CHECKOUT_DEFAULT_OAUTH_ISSUING_CLIENT_ID"),
				os.Getenv("CHECKOUT_DEFAULT_OAUTH_ISSUING_CLIENT_SECRET")).
			WithEnvironment(configuration.Sandbox()).
			WithScopes([]string{
				configuration.Vault,
				configuration.IssuingClient,
				configuration.IssuingCardMgmt,
				configuration.IssuingControlsRead,
				configuration.IssuingControlsWrite}).
			Build()
	}

	return issuingClientApi
}

func cardholderRequest(request issuing.CardholderRequest) *issuing.CardholderResponse {
	response, _ := buildIssuingClientApi().Issuing.CreateCardholder(request)
	return response
}

func cardRequest(cardholderResponse *issuing.CardholderResponse) *issuing.CardResponse {
	virtualCard := issuing.NewVirtualCardTypeRequest()
	virtualCard.CardDetailsRequest = issuing.CardDetailsRequest{
		Type:         issuing.Virtual,
		CardholderId: cardholderResponse.Id,
		Lifetime: issuing.CardLifetime{
			Unit:  issuing.Months,
			Value: 6,
		},
		Reference:     "X-123456-N11",
		CardProductId: "pro_3fn6pv2ikshurn36dbd3iysyha",
		DisplayName:   "John Kennedy",
		ActivateCard:  false,
	}
	virtualCard.IsSingleUse = true

	response, _ := buildIssuingClientApi().Issuing.CreateCard(virtualCard)
	return response
}
