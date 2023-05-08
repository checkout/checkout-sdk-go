package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/errors"
	"github.com/checkout/checkout-sdk-go/payments"
	"github.com/checkout/checkout-sdk-go/payments/nas"
	"github.com/checkout/checkout-sdk-go/payments/nas/sources/apm"
)

func TestRequestPaymentsAPM(t *testing.T) {
	var (
		customer = common.CustomerRequest{
			Email: Email,
			Name:  Name,
			Phone: Phone(),
		}
	)

	cases := []struct {
		name                   string
		request                nas.PaymentRequest
		checkForPaymentRequest func(*nas.PaymentResponse, error)
		checkForPaymentInfo    func(*nas.GetPaymentResponse, error)
	}{
		{
			name: "test AliPay source for request payment",
			request: nas.PaymentRequest{
				Source:              apm.NewRequestAlipayPlusSource(),
				Amount:              10,
				Currency:            common.EUR,
				Reference:           Reference,
				ProcessingChannelId: "pc_5jp2az55l3cuths25t5p3xhwru",
				SuccessUrl:          SuccessUrl,
				FailureUrl:          FailureUrl,
			},
			checkForPaymentRequest: func(response *nas.PaymentResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				ckoErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, ckoErr.StatusCode)
				assert.Equal(t, "apm_service_unavailable", ckoErr.Data.ErrorCodes[0])
			},
		},
		{
			name: "test CvConnect source for request payment",
			request: nas.PaymentRequest{
				Source:      apm.NewRequestCvConnectSource(),
				Amount:      100,
				Currency:    common.EUR,
				Capture:     true,
				Reference:   Reference,
				Description: Description,
				SuccessUrl:  SuccessUrl,
				FailureUrl:  FailureUrl,
			},
			checkForPaymentRequest: func(response *nas.PaymentResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				ckoErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, ckoErr.StatusCode)
				assert.Equal(t, "payee_not_onboarded", ckoErr.Data.ErrorCodes[0])
			},
		},
		{
			name: "test Ideal source for request payment",
			request: nas.PaymentRequest{
				Source:      getIdealSourceRequest(),
				Amount:      1000,
				Currency:    common.EUR,
				Reference:   Reference,
				Description: Description,
				Customer:    &customer,
				SuccessUrl:  SuccessUrl,
				FailureUrl:  FailureUrl,
			},
			checkForPaymentRequest: func(response *nas.PaymentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			},
			checkForPaymentInfo: func(response *nas.GetPaymentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.Amount)
				assert.Equal(t, int64(1000), response.Amount)
				assert.NotNil(t, response.Currency)
				assert.Equal(t, common.EUR, response.Currency)
				assert.NotNil(t, response.Reference)
				assert.Equal(t, Reference, response.Reference)
				assert.NotNil(t, response.Description)
				assert.Equal(t, Description, response.Description)
				assert.NotNil(t, response.Customer)
			},
		},
		{
			name: "test Illicado source for request payment",
			request: nas.PaymentRequest{
				Source:      apm.NewRequestIllicadoSource(),
				Amount:      100,
				Currency:    common.EUR,
				Capture:     true,
				Reference:   Reference,
				Description: Description,
				SuccessUrl:  SuccessUrl,
				FailureUrl:  FailureUrl,
			},
			checkForPaymentRequest: func(response *nas.PaymentResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				ckoErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, ckoErr.StatusCode)
				assert.Equal(t, "payee_not_onboarded", ckoErr.Data.ErrorCodes[0])
			},
		},
		{
			name: "test PayPal source for request payment",
			request: nas.PaymentRequest{
				Source:      apm.NewRequestPayPalSource(),
				Amount:      1000,
				Currency:    common.EUR,
				Reference:   Reference,
				Description: Description,
				Customer:    &customer,
				Items: []payments.Product{
					{
						Name:      "test item",
						Quantity:  1,
						UnitPrice: 1000,
					},
				},
				SuccessUrl: SuccessUrl,
				FailureUrl: FailureUrl,
			},
			checkForPaymentRequest: func(response *nas.PaymentResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				ckoErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, ckoErr.StatusCode)
				assert.Equal(t, "payee_not_onboarded", ckoErr.Data.ErrorCodes[0])
			},
		},
		{
			name: "test Sofort source for request payment",
			request: nas.PaymentRequest{
				Source:      apm.NewRequestSofortSource(),
				Amount:      100,
				Currency:    common.EUR,
				Reference:   Reference,
				Description: Description,
				Customer:    &customer,
				SuccessUrl:  SuccessUrl,
				FailureUrl:  FailureUrl,
			},
			checkForPaymentRequest: func(response *nas.PaymentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			},
			checkForPaymentInfo: func(response *nas.GetPaymentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.Amount)
				assert.Equal(t, int64(100), response.Amount)
				assert.NotNil(t, response.Currency)
				assert.Equal(t, common.EUR, response.Currency)
				assert.NotNil(t, response.Reference)
				assert.Equal(t, Reference, response.Reference)
				assert.NotNil(t, response.Description)
				assert.Equal(t, Description, response.Description)
				assert.NotNil(t, response.Customer)
			},
		},
		{
			name: "test Tamara source for request payment",
			request: nas.PaymentRequest{
				Source:      getTamaraSourceRequest(),
				Amount:      1000,
				Currency:    common.SAR,
				Reference:   Reference,
				Description: Description,
				Customer:    &customer,
				Items: []payments.Product{
					{
						Name:      "test item",
						Quantity:  1,
						UnitPrice: 1000,
					},
				},
				SuccessUrl: SuccessUrl,
				FailureUrl: FailureUrl,
			},
			checkForPaymentRequest: func(response *nas.PaymentResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				ckoErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, ckoErr.StatusCode)
				assert.Equal(t, "payee_not_onboarded", ckoErr.Data.ErrorCodes[0])
			},
		},
		{
			name: "test Trustly source for request payment",
			request: nas.PaymentRequest{
				Source:      apm.NewRequestTrustlySource(),
				Amount:      100,
				Currency:    common.EUR,
				Capture:     true,
				Reference:   Reference,
				Description: Description,
				SuccessUrl:  SuccessUrl,
				FailureUrl:  FailureUrl,
			},
			checkForPaymentRequest: func(response *nas.PaymentResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				ckoErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, ckoErr.StatusCode)
				assert.Equal(t, "payee_not_onboarded", ckoErr.Data.ErrorCodes[0])
			},
		},
		{
			name: "test AfterPay source for request payment",
			request: nas.PaymentRequest{
				Source:     getAfterPaySourceRequest(),
				Amount:     10,
				Currency:   common.GBP,
				Reference:  Reference,
				SuccessUrl: SuccessUrl,
				FailureUrl: FailureUrl,
			},
			checkForPaymentRequest: func(response *nas.PaymentResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				ckoErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, ckoErr.StatusCode)
				assert.Equal(t, "account_holder_birth_date_required", ckoErr.Data.ErrorCodes[0])
			},
		},
		{
			name: "test Benefit source for request payment",
			request: nas.PaymentRequest{
				Source:     apm.NewRequestBenefitSource(),
				Amount:     10,
				Currency:   common.BHD,
				Reference:  Reference,
				SuccessUrl: SuccessUrl,
				FailureUrl: FailureUrl,
			},
			checkForPaymentRequest: func(response *nas.PaymentResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				ckoErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, ckoErr.StatusCode)
				assert.Equal(t, "payee_not_onboarded", ckoErr.Data.ErrorCodes[0])
			},
		},
		{
			name: "test QPay source for request payment",
			request: nas.PaymentRequest{
				Source:     getQPaySource(),
				Amount:     10,
				Currency:   common.QAR,
				Reference:  Reference,
				SuccessUrl: SuccessUrl,
				FailureUrl: FailureUrl,
			},
			checkForPaymentRequest: func(response *nas.PaymentResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				ckoErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, ckoErr.StatusCode)
				assert.Equal(t, "payee_not_onboarded", ckoErr.Data.ErrorCodes[0])
			},
		},
		{
			name: "test MBway source for request payment",
			request: nas.PaymentRequest{
				Source:              apm.NewRequestMbwaySource(),
				Amount:              10,
				Currency:            common.GBP,
				Reference:           Reference,
				ProcessingChannelId: "pc_5jp2az55l3cuths25t5p3xhwru",
				SuccessUrl:          SuccessUrl,
				FailureUrl:          FailureUrl,
			},
			checkForPaymentRequest: func(response *nas.PaymentResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				ckoErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, ckoErr.StatusCode)
				assert.Equal(t, "cko_processing_channel_id_invalid", ckoErr.Data.ErrorCodes[0])
			},
		},
		{
			name: "test Eps source for request payment",
			request: nas.PaymentRequest{
				Source:     getEpsSource(),
				Amount:     10,
				Currency:   common.EUR,
				Reference:  Reference,
				SuccessUrl: SuccessUrl,
				FailureUrl: FailureUrl,
			},
			checkForPaymentRequest: func(response *nas.PaymentResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				ckoErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, ckoErr.StatusCode)
				assert.Equal(t, "payee_not_onboarded", ckoErr.Data.ErrorCodes[0])
			},
		},
		{
			name: "test GiroPay source for request payment",
			request: nas.PaymentRequest{
				Source:          getGiropaySource(),
				Amount:          10,
				Currency:        common.EUR,
				Reference:       Reference,
				Description:     Description,
				ShippingDetails: &payments.ShippingDetails{Address: Address(), Phone: Phone()},
				SuccessUrl:      SuccessUrl,
				FailureUrl:      FailureUrl,
			},
			checkForPaymentRequest: func(response *nas.PaymentResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				ckoErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, ckoErr.StatusCode)
				assert.Equal(t, "payee_not_onboarded", ckoErr.Data.ErrorCodes[0])
			},
		},
		{
			name: "test P24 source for request payment",
			request: nas.PaymentRequest{
				Source:     getP24Source(),
				Amount:     10,
				Currency:   common.PLN,
				Reference:  Reference,
				SuccessUrl: SuccessUrl,
				FailureUrl: FailureUrl,
			},
			checkForPaymentRequest: func(response *nas.PaymentResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				ckoErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, ckoErr.StatusCode)
				assert.Equal(t, "payee_not_onboarded", ckoErr.Data.ErrorCodes[0])
			},
		},
		{
			name: "test KNET source for request payment",
			request: nas.PaymentRequest{
				Source:     getKnetSource(),
				Amount:     10,
				Currency:   common.KWD,
				Reference:  Reference,
				SuccessUrl: SuccessUrl,
				FailureUrl: FailureUrl,
			},
			checkForPaymentRequest: func(response *nas.PaymentResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				ckoErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, ckoErr.StatusCode)
				assert.Equal(t, "payee_not_onboarded", ckoErr.Data.ErrorCodes[0])
			},
		},
		{
			name: "test Bancontact source for request payment",
			request: nas.PaymentRequest{
				Source:     getBancontactSource(),
				Amount:     10,
				Currency:   common.EUR,
				Reference:  Reference,
				SuccessUrl: SuccessUrl,
				FailureUrl: FailureUrl,
			},
			checkForPaymentRequest: func(response *nas.PaymentResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				ckoErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, ckoErr.StatusCode)
				assert.Equal(t, "payee_not_onboarded", ckoErr.Data.ErrorCodes[0])
			},
		},
		{
			name: "test Multibanco source for request payment",
			request: nas.PaymentRequest{
				Source:     getMultiBancoSource(),
				Amount:     10,
				Currency:   common.EUR,
				Reference:  Reference,
				SuccessUrl: SuccessUrl,
				FailureUrl: FailureUrl,
			},
			checkForPaymentRequest: func(response *nas.PaymentResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				ckoErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, ckoErr.StatusCode)
				assert.Equal(t, "payee_not_onboarded", ckoErr.Data.ErrorCodes[0])
			},
		},
		{
			name: "test Postfinance source for request payment",
			request: nas.PaymentRequest{
				Source:     getPostFinanceSource(),
				Amount:     10,
				Currency:   common.EUR,
				Reference:  Reference,
				SuccessUrl: SuccessUrl,
				FailureUrl: FailureUrl,
			},
			checkForPaymentRequest: func(response *nas.PaymentResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				ckoErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, ckoErr.StatusCode)
				assert.Equal(t, "payee_not_onboarded", ckoErr.Data.ErrorCodes[0])
			},
		},
		{
			name: "test STC source for request payment",
			request: nas.PaymentRequest{
				Source:     apm.NewRequestStcPaySource(),
				Amount:     10,
				Currency:   common.SAR,
				Customer:   &customer,
				Reference:  Reference,
				SuccessUrl: SuccessUrl,
				FailureUrl: FailureUrl,
			},
			checkForPaymentRequest: func(response *nas.PaymentResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				ckoErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, ckoErr.StatusCode)
				assert.Equal(t, "currency_not_supported", ckoErr.Data.ErrorCodes[0])
			},
		},
		{
			name: "test Alma source for request payment",
			request: nas.PaymentRequest{
				Source:     getAlmaSource(),
				Amount:     10,
				Currency:   common.SAR,
				Customer:   &customer,
				Reference:  Reference,
				SuccessUrl: SuccessUrl,
				FailureUrl: FailureUrl,
			},
			checkForPaymentRequest: func(response *nas.PaymentResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				ckoErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, ckoErr.StatusCode)
				assert.Equal(t, "payee_not_onboarded", ckoErr.Data.ErrorCodes[0])
			},
		},
		{
			name: "test Klarna source for request payment",
			request: nas.PaymentRequest{
				Source:     getKlarnaSource(),
				Amount:     10,
				Currency:   common.SAR,
				Customer:   &customer,
				Reference:  Reference,
				SuccessUrl: SuccessUrl,
				FailureUrl: FailureUrl,
			},
			checkForPaymentRequest: func(response *nas.PaymentResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				ckoErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, ckoErr.StatusCode)
				assert.Equal(t, "apm_service_unavailable", ckoErr.Data.ErrorCodes[0])
			},
		},
		{
			name: "test Fawry source for request payment",
			request: nas.PaymentRequest{
				Source:     getFawrySource(),
				Amount:     10,
				Currency:   common.EGP,
				Reference:  Reference,
				SuccessUrl: SuccessUrl,
				FailureUrl: FailureUrl,
			},
			checkForPaymentRequest: func(response *nas.PaymentResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				ckoErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, ckoErr.StatusCode)
				assert.Equal(t, "payee_not_onboarded", ckoErr.Data.ErrorCodes[0])
			},
		},
		{
			name: "test Sepa source for request payment",
			request: nas.PaymentRequest{
				Source:     getSepaSource(),
				Amount:     10,
				Currency:   common.EUR,
				Reference:  Reference,
				SuccessUrl: SuccessUrl,
				FailureUrl: FailureUrl,
			},
			checkForPaymentRequest: func(response *nas.PaymentResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				ckoErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, ckoErr.StatusCode)
				assert.Equal(t, "apm_service_unavailable", ckoErr.Data.ErrorCodes[0])
			},
		},
	}

	client := DefaultApi().Payments

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			response, err := client.RequestPayment(tc.request, nil)

			tc.checkForPaymentRequest(response, err)

			if response != nil {
				tc.checkForPaymentInfo(client.GetPaymentDetails(response.Id))
			}
		})
	}
}

func getIdealSourceRequest() payments.PaymentSource {
	source := apm.NewRequestIdealSource()
	source.Description = "ORD50234E89"
	source.Bic = "INGBNL2A"
	source.Language = "nl"

	return source
}

func getAfterPaySourceRequest() payments.PaymentSource {
	source := apm.NewRequestAfterPaySource()
	source.AccountHolder = AccountHolder()

	return source
}

func getTamaraSourceRequest() payments.PaymentSource {
	source := apm.NewRequestTamaraSource()
	source.BillingAddress = Address()

	return source
}

func getQPaySource() payments.PaymentSource {
	source := apm.NewRequestQPaySource()
	source.Description = "QPay Demo Payment"
	source.Language = "en"
	source.Quantity = 1
	source.NationalId = "070AYY010BU234M"

	return source
}

func getEpsSource() payments.PaymentSource {
	source := apm.NewRequestEpsSource()
	source.Purpose = "Mens black t-shirt L"

	return source
}

func getGiropaySource() payments.PaymentSource {
	source := apm.NewRequestGiropaySource()

	return source
}

func getP24Source() payments.PaymentSource {
	source := apm.NewRequestP24Source()
	source.PaymentCountry = common.PL
	source.AccountHolderName = "Bruce Wayne"
	source.AccountHolderEmail = "bruce@wayne-enterprises.com"
	source.BillingDescriptor = "P24 Demo Payment"

	return source
}

func getKnetSource() payments.PaymentSource {
	source := apm.NewRequestKnetSource()
	source.Language = "en"

	return source
}

func getBancontactSource() payments.PaymentSource {
	source := apm.NewRequestBancontactSource()
	source.PaymentCountry = common.BE
	source.AccountHolderName = "Bruce Wayne"
	source.BillingDescriptor = "Bancontact Demo Payment"

	return source
}

func getMultiBancoSource() payments.PaymentSource {
	source := apm.NewRequestMultiBancoSource()
	source.PaymentCountry = common.PT
	source.AccountHolderName = "Bruce Wayne"
	source.BillingDescriptor = "Multibanco Demo Payment"

	return source
}

func getPostFinanceSource() payments.PaymentSource {
	source := apm.NewRequestPostFinanceSource()
	source.PaymentCountry = common.CH
	source.AccountHolderName = "Bruce Wayne"
	source.BillingDescriptor = "Postfinance Demo Payment"

	return source
}

func getAlmaSource() payments.PaymentSource {
	source := apm.NewRequestAlmaSource()
	source.BillingAddress = Address()

	return source
}

func getKlarnaSource() payments.PaymentSource {
	source := apm.NewRequestKlarnaSource()
	source.AccountHolder = AccountHolder()

	return source
}

func getFawrySource() payments.PaymentSource {
	source := apm.NewRequestFawrySource()
	source.Description = "Fawry Demo Payment"
	source.CustomerMobile = "01058375055"
	source.CustomerEmail = "bruce@wayne-enterprises.com"
	source.Products = []apm.FawryProduct{{
		ProductId:   "0123456789",
		Quantity:    1,
		Price:       10,
		Description: "Fawry Demo Product",
	}}

	return source
}

func getSepaSource() payments.PaymentSource {
	source := apm.NewRequestSepaSource()
	source.Currency = common.EUR
	source.Country = common.ES
	source.AccountNumber = "HU93116000060000000012345676"
	source.BankCode = "37040044"
	source.MandateId = "man_12321233211"
	source.DateOfSignature = "2023-01-01"
	source.AccountHolder = AccountHolder()

	return source
}
