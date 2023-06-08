package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/configuration"
	issuingTesting "github.com/checkout/checkout-sdk-go/issuing/testing"
	"github.com/checkout/checkout-sdk-go/nas"

	cardholders "github.com/checkout/checkout-sdk-go/issuing/cardholders"
	cards "github.com/checkout/checkout-sdk-go/issuing/cards"
	controls "github.com/checkout/checkout-sdk-go/issuing/controls"
)

var (
	issuingClientApi *nas.Api

	cardholder          cardholders.CardholderRequest
	cardholderResponse  *cardholders.CardholderResponse
	virtualCardResponse *cards.CardResponse
	cardControlResponse *controls.CardControlResponse

	virtualCardId string
)

func TestSetupIssuing(t *testing.T) {
	t.Skip("Avoid creating cards all the time")

	cardholderResponse = cardholderRequest(t)
	virtualCardResponse = cardRequest(t, cardholderResponse)
	cardControlResponse = cardControlRequest(t, virtualCardId)

	virtualCardId = virtualCardResponse.Id
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

func cardholderRequest(t *testing.T) *cardholders.CardholderResponse {
	request := cardholders.CardholderRequest{
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

	response, err := buildIssuingClientApi().Issuing.CreateCardholder(request)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("error creating cardholder - %s", err.Error()))
	}

	return response
}

func cardRequest(t *testing.T, cardholderResponse *cardholders.CardholderResponse) *cards.CardResponse {
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

	response, err := buildIssuingClientApi().Issuing.CreateCard(virtualCard)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("error creating card - %s", err.Error()))
	}

	return response
}

func cardControlRequest(t *testing.T, virtualCardId string) *controls.CardControlResponse {
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
	response, err := buildIssuingClientApi().Issuing.CreateControl(velocityCardControl)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("error creating card control - %s", err.Error()))
	}

	return response
}

func cardSimulation(t *testing.T, card cards.CardResponse) *issuingTesting.CardAuthorizationResponse {
	request := issuingTesting.CardAuthorizationRequest{
		Card: issuingTesting.CardSimulation{
			Id:          card.Id,
			ExpiryMonth: card.ExpiryMonth,
			ExpiryYear:  card.ExpiryYear,
		},
		Transaction: issuingTesting.TransactionSimulation{
			Type:     issuingTesting.Purchase,
			Amount:   100,
			Currency: common.GBP,
		},
	}

	response, err := buildIssuingClientApi().Issuing.SimulateAuthorization(request)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("error authorizing transaction - %s", err.Error()))
	}

	return response
}
