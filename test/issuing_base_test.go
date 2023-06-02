package test

import (
	"github.com/checkout/checkout-sdk-go/common"
	"os"
	"testing"

	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/checkout/checkout-sdk-go/nas"

	cardholders "github.com/checkout/checkout-sdk-go/issuing/cardholders"
	cards "github.com/checkout/checkout-sdk-go/issuing/cards"
	controls "github.com/checkout/checkout-sdk-go/issuing/controls"
)

var issuingClientApi *nas.Api
var cardholder = cardholders.CardholderRequest{}
var cardholderResponse *cardholders.CardholderResponse
var virtualCardResponse *cards.CardResponse
var virtualCardId string
var cardControlResponse *controls.CardControlResponse

func TestSetupIssuing(t *testing.T) {
	t.Skip("Avoid creating cards all the time")
	cardholder = cardholders.CardholderRequest{
		Type:             cardholders.Individual,
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
		Document: &cardholders.CardholderDocument{
			Type:            common.NationalIdentityCard,
			FrontDocumentId: "file_6lbss42ezvoufcb2beo76rvwly",
			BackDocumentId:  "file_aaz5pemp6326zbuvevp6qroqu4",
		},
	}

	cardholderResponse = cardholderRequest(cardholder)
	virtualCardResponse = cardRequest(cardholderResponse)
	virtualCardId = virtualCardResponse.Id
	cardControlResponse = cardControlRequest(virtualCardId)
}

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

func cardholderRequest(request cardholders.CardholderRequest) *cardholders.CardholderResponse {
	response, _ := buildIssuingClientApi().Issuing.CreateCardholder(request)
	return response
}

func cardRequest(cardholderResponse *cardholders.CardholderResponse) *cards.CardResponse {
	virtualCard := cards.NewVirtualCardRequest()
	virtualCard.CardDetailsRequest = cards.CardDetailsRequest{
		Type:         cards.Virtual,
		CardholderId: cardholderResponse.Id,
		Lifetime: cards.CardLifetime{
			Unit:  cards.Months,
			Value: 6,
		},
		Reference:     "X-123456-N11",
		CardProductId: "pro_3fn6pv2ikshurn36dbd3iysyha",
		DisplayName:   "John Kennedy",
		ActivateCard:  false,
	}
	virtualCard.IsSingleUse = false

	response, _ := buildIssuingClientApi().Issuing.CreateCard(virtualCard)
	return response
}

func cardControlRequest(virtualCardId string) *controls.CardControlResponse {
	velocityCardControl := controls.NewVelocityCardControlRequest()
	velocityCardControl.ControlType = controls.VelocityLimitType
	velocityCardControl.Description = "Max spend of 500€ per week for restaurants"
	velocityCardControl.TargetId = virtualCardId
	velocityCardControl.VelocityLimit = controls.VelocityLimit{
		AmountLimit: 500,
		VelocityWindow: controls.VelocityWindow{
			Type: controls.Weekly,
		},
	}
	response, _ := buildIssuingClientApi().Issuing.CreateControl(velocityCardControl)
	return response
}
