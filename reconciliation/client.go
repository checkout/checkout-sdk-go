package reconciliation

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/httpclient"
	"github.com/checkout/checkout-sdk-go/internal/utils"
	"github.com/google/go-querystring/query"
)

const paymentsPath = "reporting/payments"
const statementsPath = "reporting/statements"

// Client ...
type Client struct {
	API checkout.HTTPClient
}

// NewClient ...
func NewClient(config checkout.Config) *Client {
	return &Client{
		API: httpclient.NewClient(config),
	}
}

// PaymentsReport -
func (c *Client) PaymentsReport(request *Request) (*Response, error) {

	value, _ := query.Values(request.PaymentsParameter)
	var query string = value.Encode()
	var urlPath string = "/" + paymentsPath + "?"
	resp, err := c.API.Get(urlPath + query)
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusOK {
		var paymentsReport PaymentsReport
		err = json.Unmarshal(resp.ResponseBody, &paymentsReport)
		response.PaymentsReport = &paymentsReport
		return response, err
	}
	return response, err
}

// PaymentsReportCSV -
func (c *Client) PaymentsReportCSV(request *Request) (*Response, error) {

	value, _ := query.Values(request.PaymentsParameter)
	var query string = value.Encode()
	var urlPath string = "/" + paymentsPath + "/" + "download" + "?"
	resp, err := c.API.Download(urlPath + query)
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusOK {
		response.CSV = resp.ResponseCSV
		return response, err
	}
	return response, err
}

// StatementsReportCSV -
func (c *Client) StatementsReportCSV(request *Request) (*Response, error) {

	value, _ := query.Values(request.StatementsParameter)
	var query string = value.Encode()
	var urlPath string = "/" + statementsPath + "/" + "download" + "?"
	resp, err := c.API.Download(urlPath + query)
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusOK {
		response.CSV = resp.ResponseCSV
		return response, err
	}
	return response, err
}

// StatementPaymentReportCSV -
func (c *Client) StatementPaymentReportCSV(statementID *string) (*Response, error) {

	resp, err := c.API.Download(fmt.Sprintf("/%v/%v/payments/download", statementsPath, utils.StringValue(statementID)))
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusOK {
		response.CSV = resp.ResponseCSV
		return response, err
	}
	return response, err
}

// PaymentReport -
func (c *Client) PaymentReport(paymentID *string, request *Request) (*Response, error) {

	if paymentID != nil {

		resp, err := c.API.Get(fmt.Sprintf("/%v/%v", paymentsPath, utils.StringValue(paymentID)))
		response := &Response{
			StatusResponse: resp,
		}
		if err != nil {
			return response, err
		}
		if resp.StatusCode == http.StatusOK {
			var paymentsReport PaymentsReport
			err = json.Unmarshal(resp.ResponseBody, &paymentsReport)
			response.PaymentsReport = &paymentsReport
			return response, err
		}
		return response, err
	}
	if request != nil {
		value, _ := query.Values(request.PaymentParameter)
		var query string = value.Encode()
		var urlPath string = "/" + paymentsPath + "?"
		resp, err := c.API.Get(urlPath + query)
		response := &Response{
			StatusResponse: resp,
		}
		if err != nil {
			return response, err
		}
		if resp.StatusCode == http.StatusOK {
			var paymentsReport PaymentsReport
			err = json.Unmarshal(resp.ResponseBody, &paymentsReport)
			response.PaymentsReport = &paymentsReport
			return response, err
		}
		return response, err
	}
	return nil, nil
}

// StatementsReport -
func (c *Client) StatementsReport(request *Request) (*Response, error) {

	value, _ := query.Values(request.StatementsParameter)
	var query string = value.Encode()
	var urlPath string = "/" + statementsPath + "?"
	resp, err := c.API.Get(urlPath + query)
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusOK {
		var statementsReport StatementsReport
		err = json.Unmarshal(resp.ResponseBody, &statementsReport)
		response.StatementsReport = &statementsReport
		return response, err
	}
	return response, err
}

// StatementPaymentReport -
func (c *Client) StatementPaymentReport(statementID string, request *Request) (*Response, error) {

	if request != nil {
		value, _ := query.Values(request.StatementParameter)
		var query string = value.Encode()
		var urlPath string = "/" + statementsPath + "/" + statementID + "/" + "payments" + "?"
		resp, err := c.API.Get(urlPath + query)
		response := &Response{
			StatusResponse: resp,
		}
		if err != nil {
			return response, err
		}
		if resp.StatusCode == http.StatusOK {
			var paymentsReport PaymentsReport
			err = json.Unmarshal(resp.ResponseBody, &paymentsReport)
			response.PaymentsReport = &paymentsReport
			return response, err
		}
		return response, err
	}
	resp, err := c.API.Get(fmt.Sprintf("/%v/%v/payments", statementsPath, statementID))
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusOK {
		var paymentsReport PaymentsReport
		err = json.Unmarshal(resp.ResponseBody, &paymentsReport)
		response.PaymentsReport = &paymentsReport
		return response, err
	}
	return response, err
}
