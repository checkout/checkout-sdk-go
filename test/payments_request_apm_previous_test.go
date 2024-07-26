package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/errors"
	"github.com/checkout/checkout-sdk-go/payments"
	"github.com/checkout/checkout-sdk-go/payments/abc"
	"github.com/checkout/checkout-sdk-go/payments/abc/sources/apm"
)

func TestRequestPaymentsAPMPrevious(t *testing.T) {
	var (
		customer = common.CustomerRequest{
			Id:    "cus_vtkefqy4pevebjhkdp5bkklncy",
			Email: Email,
			Name:  Name,
			Phone: Phone(),
		}
	)

	cases := []struct {
		name                   string
		request                abc.PaymentRequest
		checkForPaymentRequest func(*abc.PaymentResponse, error)
		checkForPaymentInfo    func(*abc.GetPaymentResponse, error)
	}{
		{
			name: "test AliPay source for request payment",
			request: abc.PaymentRequest{
				Source:      getAlipaySourceRequest(),
				Amount:      0,
				Currency:    common.USD,
				Reference:   Reference,
				Description: Description,
				Customer:    &customer,
			},
			checkForPaymentRequest: func(response *abc.PaymentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			},
			checkForPaymentInfo: func(response *abc.GetPaymentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.Amount)
				assert.Equal(t, int64(0), response.Amount)
				assert.NotNil(t, response.Currency)
				assert.Equal(t, common.USD, response.Currency)
				assert.NotNil(t, response.Reference)
				assert.Equal(t, Reference, response.Reference)
				assert.NotNil(t, response.Description)
				assert.Equal(t, Description, response.Description)
				assert.NotNil(t, response.Customer)
				assert.Equal(t, customer.Id, response.Customer.Id)
			},
		},
		{
			name: "test Bancontact source for request payment",
			request: abc.PaymentRequest{
				Source:      getBancontactSourceRequest(),
				Amount:      100,
				Currency:    common.EUR,
				Reference:   Reference,
				Description: Description,
				Customer:    &customer,
				Capture:     true,
			},
			checkForPaymentRequest: func(response *abc.PaymentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			},
			checkForPaymentInfo: func(response *abc.GetPaymentResponse, err error) {
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
				assert.Equal(t, customer.Id, response.Customer.Id)
			},
		},
		{
			name: "test BenefitPay source for request payment",
			request: abc.PaymentRequest{
				Source:      getBenefitPaySourceRequest(),
				Amount:      100,
				Currency:    common.USD,
				Reference:   Reference,
				Description: Description,
				Customer:    &customer,
			},
			checkForPaymentRequest: func(response *abc.PaymentResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Equal(t, "payment_method_not_supported", chkErr.Data.ErrorCodes[0])
			},
			checkForPaymentInfo: func(response *abc.GetPaymentResponse, err error) {
				/*TODO: BenefitPay is deprecated
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.Amount)
				assert.Equal(t, 10000, *response.Amount)
				assert.NotNil(t, response.Currency)
				assert.Equal(t, common.COP, *response.Currency)
				assert.NotNil(t, response.Reference)
				assert.Equal(t, Reference, *response.Reference)
				assert.NotNil(t, response.Description)
				assert.Equal(t, Description, *response.Description)
				assert.NotNil(t, response.Customer)
				assert.Equal(t, customer.Id, *response.Customer.Id)*/
			},
		},
		{
			name: "test Boleto source for request payment redirect",
			request: abc.PaymentRequest{
				Source:      getBoletoSourceRequest(apm.Redirect),
				Amount:      100,
				Currency:    common.BRL,
				Reference:   Reference,
				Description: Description,
				Customer:    &customer,
			},
			checkForPaymentRequest: func(response *abc.PaymentResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Equal(t, "business_not_onboarded", chkErr.Data.ErrorCodes[0])
			},
			checkForPaymentInfo: func(response *abc.GetPaymentResponse, err error) {
				/*TODO: uncomment when "busines_not_onboarded" error gets fixed
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.Amount)
				assert.Equal(t, 100, *response.Amount)
				assert.NotNil(t, response.Currency)
				assert.Equal(t, common.BRL, *response.Currency)
				assert.NotNil(t, response.Reference)
				assert.Equal(t, Reference, *response.Reference)
				assert.NotNil(t, response.Description)
				assert.Equal(t, Description, *response.Description)
				assert.NotNil(t, response.Customer)
				assert.Equal(t, customer.Id, *response.Customer.Id)*/
			},
		},
		{
			name: "test Boleto source for request payment direct",
			request: abc.PaymentRequest{
				Source:      getBoletoSourceRequest(apm.Direct),
				Amount:      100,
				Currency:    common.BRL,
				Reference:   Reference,
				Description: Description,
				Customer:    &customer,
			},
			checkForPaymentRequest: func(response *abc.PaymentResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Equal(t, "business_not_onboarded", chkErr.Data.ErrorCodes[0])
			},
			checkForPaymentInfo: func(response *abc.GetPaymentResponse, err error) {
				/*TODO: uncomment when "busines_not_onboarded" error gets fixed
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.Amount)
				assert.Equal(t, 100, *response.Amount)
				assert.NotNil(t, response.Currency)
				assert.Equal(t, common.BRL, *response.Currency)
				assert.NotNil(t, response.Reference)
				assert.Equal(t, Reference, *response.Reference)
				assert.NotNil(t, response.Description)
				assert.Equal(t, Description, *response.Description)
				assert.NotNil(t, response.Customer)
				assert.Equal(t, customer.Id, *response.Customer.Id)*/
			},
		},
		{
			name: "test EPS source for request payment",
			request: abc.PaymentRequest{
				Source:      getEpsSourceRequest(),
				Amount:      100,
				Currency:    common.EUR,
				Reference:   Reference,
				Description: Description,
				Customer:    &customer,
			},
			checkForPaymentRequest: func(response *abc.PaymentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			},
			checkForPaymentInfo: func(response *abc.GetPaymentResponse, err error) {
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
				assert.Equal(t, customer.Id, response.Customer.Id)
			},
		},
		{
			name: "test Fawry source for request payment",
			request: abc.PaymentRequest{
				Source:      getFawrySourceRequest(),
				Amount:      1000,
				Currency:    common.EGP,
				Reference:   Reference,
				Description: Description,
				Customer:    &customer,
			},
			checkForPaymentRequest: func(response *abc.PaymentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			},
			checkForPaymentInfo: func(response *abc.GetPaymentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.Amount)
				assert.Equal(t, int64(1000), response.Amount)
				assert.NotNil(t, response.Currency)
				assert.Equal(t, common.EGP, response.Currency)
				assert.NotNil(t, response.Reference)
				assert.Equal(t, Reference, response.Reference)
				assert.NotNil(t, response.Description)
				assert.Equal(t, Description, response.Description)
				assert.NotNil(t, response.Customer)
				assert.Equal(t, customer.Id, response.Customer.Id)
			},
		},
		{
			name: "test Giropay source for request payment",
			request: abc.PaymentRequest{
				Source:      getGiropaySourceRequest(),
				Amount:      1000,
				Currency:    common.EUR,
				Reference:   Reference,
				Description: Description,
				Customer:    &customer,
			},
			checkForPaymentRequest: func(response *abc.PaymentResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Equal(t, "payment_method_not_supported", chkErr.Data.ErrorCodes[0])
			},
			checkForPaymentInfo: func(response *abc.GetPaymentResponse, err error) {
				/*TODO: uncomment when "payment_method_not_supported" error gets fixed
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
				assert.Equal(t, customer.Id, response.Customer.Id)*/
			},
		},
		{
			name: "test Ideal source for request payment",
			request: abc.PaymentRequest{
				Source:      getIdealSourceRequestPrevious(),
				Amount:      1000,
				Currency:    common.EUR,
				Reference:   Reference,
				Description: Description,
				Customer:    &customer,
			},
			checkForPaymentRequest: func(response *abc.PaymentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			},
			checkForPaymentInfo: func(response *abc.GetPaymentResponse, err error) {
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
				assert.Equal(t, customer.Id, response.Customer.Id)
			},
		},
		//TODO klarna request payment
		{
			name: "test Knet source for request payment",
			request: abc.PaymentRequest{
				Source:      getKnetSourceRequest(),
				Amount:      100,
				Currency:    common.KWD,
				Reference:   Reference,
				Description: Description,
				Customer:    &customer,
				Capture:     true,
			},
			checkForPaymentRequest: func(response *abc.PaymentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			},
			checkForPaymentInfo: func(response *abc.GetPaymentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.Amount)
				assert.Equal(t, int64(100), response.Amount)
				assert.NotNil(t, response.Currency)
				assert.Equal(t, common.KWD, response.Currency)
				assert.NotNil(t, response.Reference)
				assert.Equal(t, Reference, response.Reference)
				assert.NotNil(t, response.Description)
				assert.Equal(t, Description, response.Description)
				assert.NotNil(t, response.Customer)
				assert.Equal(t, customer.Id, response.Customer.Id)
			},
		},
		{
			name: "test MultiBanco source for request payment",
			request: abc.PaymentRequest{
				Source:      getMultiBancoSourceRequest(),
				Amount:      100,
				Currency:    common.QAR,
				Reference:   Reference,
				Description: Description,
				Customer:    &customer,
				Capture:     true,
			},
			checkForPaymentRequest: func(response *abc.PaymentResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Equal(t, "payment_method_not_supported", chkErr.Data.ErrorCodes[0])
			},
			checkForPaymentInfo: func(response *abc.GetPaymentResponse, err error) {
				/*TODO: uncomment when "payment_method_not_supported" error gets fixed
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.Amount)
				assert.Equal(t, 100, *response.Amount)
				assert.NotNil(t, response.Currency)
				assert.Equal(t, common.QAR, *response.Currency)
				assert.NotNil(t, response.Reference)
				assert.Equal(t, Reference, *response.Reference)
				assert.NotNil(t, response.Description)
				assert.Equal(t, Description, *response.Description)
				assert.NotNil(t, response.Customer)
				assert.Equal(t, customer.Id, *response.Customer.Id)*/
			},
		},
		{
			name: "test Oxxo source for request payment redirect",
			request: abc.PaymentRequest{
				Source:      getOxxoSourceRequest(apm.Redirect),
				Amount:      100000,
				Currency:    common.MXN,
				Reference:   Reference,
				Description: Description,
				Customer:    &customer,
				Capture:     true,
			},
			checkForPaymentRequest: func(response *abc.PaymentResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Equal(t, "business_not_onboarded", chkErr.Data.ErrorCodes[0])
			},
			checkForPaymentInfo: func(response *abc.GetPaymentResponse, err error) {
				/*TODO: uncomment when "business_not_onboarded" error gets fixed
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.Amount)
				assert.Equal(t, 100000, *response.Amount)
				assert.NotNil(t, response.Currency)
				assert.Equal(t, common.MXN, *response.Currency)
				assert.NotNil(t, response.Reference)
				assert.Equal(t, Reference, *response.Reference)
				assert.NotNil(t, response.Description)
				assert.Equal(t, Description, *response.Description)
				assert.NotNil(t, response.Customer)
				assert.Equal(t, customer.Id, *response.Customer.Id)*/
			},
		},
		{
			name: "test Oxxo source for request payment direct",
			request: abc.PaymentRequest{
				Source:      getOxxoSourceRequest(apm.Direct),
				Amount:      100000,
				Currency:    common.MXN,
				Reference:   Reference,
				Description: Description,
				Customer:    &customer,
				Capture:     true,
			},
			checkForPaymentRequest: func(response *abc.PaymentResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Equal(t, "business_not_onboarded", chkErr.Data.ErrorCodes[0])
			},
			checkForPaymentInfo: func(response *abc.GetPaymentResponse, err error) {
				/*TODO: uncomment when "business_not_onboarded" error gets fixed
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.Amount)
				assert.Equal(t, 100000, *response.Amount)
				assert.NotNil(t, response.Currency)
				assert.Equal(t, common.MXN, *response.Currency)
				assert.NotNil(t, response.Reference)
				assert.Equal(t, Reference, *response.Reference)
				assert.NotNil(t, response.Description)
				assert.Equal(t, Description, *response.Description)
				assert.NotNil(t, response.Customer)
				assert.Equal(t, customer.Id, *response.Customer.Id)*/
			},
		},
		{
			name: "test P24 source for request payment",
			request: abc.PaymentRequest{
				Source:      getP24SourceRequest(),
				Amount:      100,
				Currency:    common.PLN,
				Reference:   Reference,
				Description: Description,
				Customer:    &customer,
				Capture:     true,
			},
			checkForPaymentRequest: func(response *abc.PaymentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			},
			checkForPaymentInfo: func(response *abc.GetPaymentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.Amount)
				assert.Equal(t, int64(100), response.Amount)
				assert.NotNil(t, response.Currency)
				assert.Equal(t, common.PLN, response.Currency)
				assert.NotNil(t, response.Reference)
				assert.Equal(t, Reference, response.Reference)
				assert.NotNil(t, response.Description)
				assert.Equal(t, Description, response.Description)
				assert.NotNil(t, response.Customer)
				assert.Equal(t, customer.Id, response.Customer.Id)
			},
		},
		{
			name: "test PagoFacil source for request payment redirect",
			request: abc.PaymentRequest{
				Source:      getPagoFacilSourceRequest(apm.Redirect),
				Amount:      100000,
				Currency:    common.ARS,
				Reference:   Reference,
				Description: Description,
				Customer:    &customer,
				Capture:     true,
			},
			checkForPaymentRequest: func(response *abc.PaymentResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Equal(t, "business_not_onboarded", chkErr.Data.ErrorCodes[0])
			},
			checkForPaymentInfo: func(response *abc.GetPaymentResponse, err error) {
				/*TODO: uncomment when "business_not_onboarded" error gets fixed
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.Amount)
				assert.Equal(t, 100000, *response.Amount)
				assert.NotNil(t, response.Currency)
				assert.Equal(t, common.ARS, *response.Currency)
				assert.NotNil(t, response.Reference)
				assert.Equal(t, Reference, *response.Reference)
				assert.NotNil(t, response.Description)
				assert.Equal(t, Description, *response.Description)
				assert.NotNil(t, response.Customer)
				assert.Equal(t, customer.Id, *response.Customer.Id)*/
			},
		},
		{
			name: "test PagoFacil source for request payment direct",
			request: abc.PaymentRequest{
				Source:      getPagoFacilSourceRequest(apm.Direct),
				Amount:      100000,
				Currency:    common.ARS,
				Reference:   Reference,
				Description: Description,
				Customer:    &customer,
				Capture:     true,
			},
			checkForPaymentRequest: func(response *abc.PaymentResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Equal(t, "business_not_onboarded", chkErr.Data.ErrorCodes[0])
			},
			checkForPaymentInfo: func(response *abc.GetPaymentResponse, err error) {
				/*TODO: uncomment when "business_not_onboarded" error gets fixed
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.Amount)
				assert.Equal(t, 100000, *response.Amount)
				assert.NotNil(t, response.Currency)
				assert.Equal(t, common.ARS, *response.Currency)
				assert.NotNil(t, response.Reference)
				assert.Equal(t, Reference, *response.Reference)
				assert.NotNil(t, response.Description)
				assert.Equal(t, Description, *response.Description)
				assert.NotNil(t, response.Customer)
				assert.Equal(t, customer.Id, *response.Customer.Id)*/
			},
		},
		{
			name: "test PayPal source for request payment",
			request: abc.PaymentRequest{
				Source:      getPayPalSourceRequestPrevious(),
				Amount:      100,
				Currency:    common.EUR,
				Reference:   Reference,
				Description: Description,
				Customer:    &customer,
				Capture:     true,
			},
			checkForPaymentRequest: func(response *abc.PaymentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			},
			checkForPaymentInfo: func(response *abc.GetPaymentResponse, err error) {
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
				assert.Equal(t, customer.Id, response.Customer.Id)
			},
		},
		{
			name: "test Poli source for request payment",
			request: abc.PaymentRequest{
				Source:      getPoliSourceRequest(),
				Amount:      100,
				Currency:    common.AUD,
				Reference:   Reference,
				Description: Description,
				Customer:    &customer,
				Capture:     true,
			},
			checkForPaymentRequest: func(response *abc.PaymentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			},
			checkForPaymentInfo: func(response *abc.GetPaymentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.Amount)
				assert.Equal(t, int64(100), response.Amount)
				assert.NotNil(t, response.Currency)
				assert.Equal(t, common.AUD, response.Currency)
				assert.NotNil(t, response.Reference)
				assert.Equal(t, Reference, response.Reference)
				assert.NotNil(t, response.Description)
				assert.Equal(t, Description, response.Description)
				assert.NotNil(t, response.Customer)
				assert.Equal(t, customer.Id, response.Customer.Id)
			},
		},
		{
			name: "test QPay source for request payment",
			request: abc.PaymentRequest{
				Source:      getQPaySourceRequest(),
				Amount:      100,
				Currency:    common.QAR,
				Reference:   Reference,
				Description: Description,
				Customer:    &customer,
				Capture:     true,
			},
			checkForPaymentRequest: func(response *abc.PaymentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			},
			checkForPaymentInfo: func(response *abc.GetPaymentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.Amount)
				assert.Equal(t, int64(100), response.Amount)
				assert.NotNil(t, response.Currency)
				assert.Equal(t, common.QAR, response.Currency)
				assert.NotNil(t, response.Reference)
				assert.Equal(t, Reference, response.Reference)
				assert.NotNil(t, response.Description)
				assert.Equal(t, Description, response.Description)
				assert.NotNil(t, response.Customer)
				assert.Equal(t, customer.Id, response.Customer.Id)
			},
		},
		{
			name: "test RapiPago source for request payment",
			request: abc.PaymentRequest{
				Source:      getRapiPagoSourceRequest(),
				Amount:      100000,
				Currency:    common.ARS,
				Reference:   Reference,
				Description: Description,
				Customer:    &customer,
				Capture:     true,
			},
			checkForPaymentRequest: func(response *abc.PaymentResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Equal(t, "business_not_onboarded", chkErr.Data.ErrorCodes[0])
			},
			checkForPaymentInfo: func(response *abc.GetPaymentResponse, err error) {
				/*TODO: uncomment when "business_not_onboarded" error gets fixed
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.Amount)
				assert.Equal(t, 100000, *response.Amount)
				assert.NotNil(t, response.Currency)
				assert.Equal(t, common.ARS, *response.Currency)
				assert.NotNil(t, response.Reference)
				assert.Equal(t, Reference, *response.Reference)
				assert.NotNil(t, response.Description)
				assert.Equal(t, Description, *response.Description)
				assert.NotNil(t, response.Customer)
				assert.Equal(t, customer.Id, *response.Customer.Id)*/
			},
		},
		//TODO Sepa request payment
		{
			name: "test Sofort source for request payment",
			request: abc.PaymentRequest{
				Source:      getSofortSourceRequestPrevious(),
				Amount:      100,
				Currency:    common.EUR,
				Reference:   Reference,
				Description: Description,
				Customer:    &customer,
				Capture:     true,
			},
			checkForPaymentRequest: func(response *abc.PaymentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			},
			checkForPaymentInfo: func(response *abc.GetPaymentResponse, err error) {
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
				assert.Equal(t, customer.Id, response.Customer.Id)
			},
		},
	}

	client := PreviousApi().Payments

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

var (
	payer = payments.Payer{
		Name:     "Bruce Wayne",
		Email:    "bruce@wayne-enterprises.com",
		Document: "53033315550",
	}
)

func getAlipaySourceRequest() payments.PaymentSource {
	return apm.NewRequestAlipaySource()
}

func getBancontactSourceRequest() payments.PaymentSource {
	source := apm.NewRequestBancontactSource()
	source.AccountHolderName = "Bruce Wayne"
	source.BillingDescriptor = "Bancontact test payment"
	source.PaymentCountry = common.BE

	return source
}

func getBenefitPaySourceRequest() payments.PaymentSource {
	return apm.NewRequestBenefitPaySource()
}

func getBoletoSourceRequest(i apm.IntegrationType) payments.PaymentSource {
	source := apm.NewRequestBoletoSource()
	source.IntegrationType = i
	source.Description = "boleto test payment"
	source.Country = common.BR
	source.Payer = &payer

	return source
}

func getEpsSourceRequest() payments.PaymentSource {
	source := apm.NewRequestEpsSource()
	source.Purpose = "purpose"

	return source
}

func getFawrySourceRequest() payments.PaymentSource {
	source := apm.NewRequestFawrySource()
	source.Description = "Fawry test payment"
	source.CustomerEmail = "bruce@wayne-enterprises.com"
	source.CustomerMobile = "01058375055"
	source.Products = []apm.FawryProduct{
		{
			ProductId:   "0123456789",
			Quantity:    1,
			Price:       1000,
			Description: "Fawry Demo Product",
		},
	}

	return source
}

func getGiropaySourceRequest() payments.PaymentSource {
	source := apm.NewRequestGiropaySource()
	source.Purpose = "Giropay test payment"

	return source
}

func getIdealSourceRequestPrevious() payments.PaymentSource {
	source := apm.NewRequestIdealSource()
	source.Description = "ORD50234E89"
	source.Bic = "INGBNL2A"
	source.Language = "nl"

	return source
}

func getKlarnaSourceRequest() payments.PaymentSource {
	source := apm.NewRequestKlarnaSource()
	source.AuthorizationToken = "b4bd3423-24e3"
	source.PurchaseCountry = string(common.GB)
	source.Locale = "en-GB"
	source.TaxAmount = 1
	source.BillingAddress = Address()
	source.Customer = map[string]interface{}{
		"DateOfBirth": "1970-01-01",
		"Gender":      "male",
	}
	source.Products = []map[string]interface{}{
		{
			"Name":           "test item",
			"Quantity":       1,
			"UnitPrice":      1000,
			"TaxRate":        0,
			"TotalAmount":    1000,
			"TotalTaxAmount": 0,
		},
	}

	return source
}

func getKnetSourceRequest() payments.PaymentSource {
	source := apm.NewRequestKnetSource()
	source.Language = "en"
	return source
}

func getMultiBancoSourceRequest() payments.PaymentSource {
	source := apm.NewRequestMultiBancoSource()
	source.AccountHolderName = "Bruce Wayne"
	source.BillingDescriptor = "MultiBanco test payment"
	source.PaymentCountry = common.PT

	return source
}

func getOxxoSourceRequest(i apm.IntegrationType) payments.PaymentSource {
	source := apm.NewRequestOxxoSource()
	source.IntegrationType = i
	source.Description = "ORD50234E89"
	source.Country = common.MX
	source.Payer = &payer

	return source
}

func getP24SourceRequest() payments.PaymentSource {
	source := apm.NewRequestP24Source()
	source.AccountHolderName = "Bruce Wayne"
	source.AccountHolderEmail = "bruce@wayne-enterprises.com"
	source.BillingDescriptor = "P24 test payment"
	source.PaymentCountry = common.PL

	return source
}

func getPagoFacilSourceRequest(i apm.IntegrationType) payments.PaymentSource {
	source := apm.NewRequestPagoFacilSource()
	source.IntegrationType = i
	source.Description = "PagoFacil test payment"
	source.Country = common.AR
	source.Payer = &payer

	return source
}

func getPayPalSourceRequestPrevious() payments.PaymentSource {
	source := apm.NewRequestPayPalSource()
	source.InvoiceNumber = "CKO00001"
	source.LogoUrl = "https://www.example.com/logo.jpg"

	return source
}

func getPoliSourceRequest() payments.PaymentSource {
	return apm.NewRequestPoliSource()
}

func getQPaySourceRequest() payments.PaymentSource {
	source := apm.NewRequestQPaySource()
	source.Description = "QPay test payment"
	source.NationalId = "070AYY010BU234M"
	source.Language = "en"
	source.Quantity = 1

	return source
}

func getRapiPagoSourceRequest() payments.PaymentSource {
	source := apm.NewRequestRapiPagoSource()
	source.Description = "RapiPago test payment"
	source.Country = common.AR
	source.Payer = &payer

	return source
}

func getSepaSourceRequest() payments.PaymentSource {
	source := apm.NewRequestSepaSource()
	source.Id = "1"

	return source
}

func getSofortSourceRequestPrevious() payments.PaymentSource {
	return apm.NewRequestSofortSource()
}
