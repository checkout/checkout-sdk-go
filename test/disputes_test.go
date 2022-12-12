package test

import (
	"fmt"
	"net/http"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/disputes"
	"github.com/checkout/checkout-sdk-go/errors"
	"github.com/checkout/checkout-sdk-go/payments"
	"github.com/checkout/checkout-sdk-go/payments/nas"
	"github.com/checkout/checkout-sdk-go/payments/nas/sources"
)

var (
	disputeId string
	fileId    string

	fileUpload = common.File{
		File:    "./checkout.jpeg",
		Purpose: common.DisputesEvidence,
	}
)

func TestSetup(t *testing.T) {
	var (
		cardToken = RequestCardToken(t)
		payment   = getPaymentRequest(t, cardToken.Token)
	)

	Wait(time.Duration(5))

	disputeId = getDispute(payment.Id).Id
	fileId = uploadDisputeFile(t, fileUpload).Id
}

func TestQuery(t *testing.T) {
	var (
		layout  = "2006-01-02T15:04:05Z"
		now     = time.Now()
		from, _ = time.Parse(layout, "2018-08-12T00:00:00Z")
		to, _   = time.Parse(layout, now.Format(layout))
	)

	cases := []struct {
		name    string
		request disputes.QueryFilter
		checker func(*disputes.QueryResponse, error)
	}{
		{
			name: "when disputes match filters then return disputes",
			request: disputes.QueryFilter{
				Limit: 10,
				From:  from,
				To:    to,
			},
			checker: func(response *disputes.QueryResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.Data)
				assert.Equal(t, from, response.From)
				assert.Equal(t, to, response.To)
			},
		},
		{
			name: "when invalid filters then return error",
			request: disputes.QueryFilter{
				Limit: 255,
			},
			checker: func(response *disputes.QueryResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				errChk := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, errChk.StatusCode)
				assert.Equal(t, "request_invalid", errChk.Data.ErrorType)
			},
		},
	}

	client := DefaultApi().Disputes

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.Query(tc.request))
		})
	}
}

func TestGetDisputeDetails(t *testing.T) {
	cases := []struct {
		name      string
		disputeId string
		checker   func(*disputes.DisputeResponse, error)
	}{
		{
			name:      "when dispute exists then return dispute details",
			disputeId: disputeId,
			checker: func(response *disputes.DisputeResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, disputeId, response.Id)
			},
		},
		{
			name:      "when dispute doesn't exist then return error",
			disputeId: "not_found",
			checker: func(response *disputes.DisputeResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := DefaultApi().Disputes

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if disputeId != "" {
				tc.checker(client.GetDisputeDetails(tc.disputeId))
			}
		})
	}
}

func TestPutEvidence(t *testing.T) {
	t.Skip("Skipping tests because checkout is having timeouts")

	cases := []struct {
		name      string
		disputeId string
		request   disputes.Evidence
		checker   func(*common.MetadataResponse, error)
	}{
		{
			name:      "when request is valid then put evidence",
			disputeId: disputeId,
			request: disputes.Evidence{
				ProofOfDeliveryOrServiceFile:           fileId,
				ProofOfDeliveryOrServiceText:           "proof of delivery or service text",
				InvoiceOrReceiptFile:                   fileId,
				InvoiceOrReceiptText:                   "proof of receipt text",
				InvoiceShowingDistinctTransactionsFile: fileId,
				InvoiceShowingDistinctTransactionsText: "invoice showing distinct transactions text",
				CustomerCommunicationFile:              fileId,
				CustomerCommunicationText:              "customer communication text",
				RefundOrCancellationPolicyFile:         fileId,
				RefundOrCancellationPolicyText:         "refund or cancellation policy text",
				RecurringTransactionAgreementFile:      fileId,
				RecurringTransactionAgreementText:      "recurring transaction agreement text",
				AdditionalEvidenceFile:                 fileId,
				AdditionalEvidenceText:                 "additional evidence text",
				ProofOfDeliveryOrServiceDateFile:       fileId,
				ProofOfDeliveryOrServiceDateText:       "proof of delivery or service date text",
			},
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusNoContent, response.HttpMetadata.StatusCode)
			},
		},
		{
			name:      "when dispute doesn't exist then return error",
			disputeId: "disp_invalid",
			request:   disputes.Evidence{},
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
		{
			name:      "when file doesn't exist then return error",
			disputeId: disputeId,
			request: disputes.Evidence{
				RefundOrCancellationPolicyFile: "file_invalid",
				RefundOrCancellationPolicyText: "text",
			},
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
			},
		},
	}

	client := DefaultApi().Disputes

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if disputeId != "" {
				tc.checker(client.PutEvidence(tc.disputeId, tc.request))
			}
		})
	}
}

func TestSubmitEvidence(t *testing.T) {
	t.Skip("Skipping tests because checkout is having timeouts")

	cases := []struct {
		name      string
		disputeId string
		checker   func(*common.MetadataResponse, error)
	}{
		{
			name:      "when dispute has evidence attached then submit dispute",
			disputeId: disputeId,
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusNoContent, response.HttpMetadata.StatusCode)
			},
		},
		{
			name:      "when dispute doesn't exist then return error",
			disputeId: "disp_invalid",
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := DefaultApi().Disputes

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if disputeId != "" {
				tc.checker(client.SubmitEvidence(tc.disputeId))
			}
		})
	}
}

func TestGetEvidence(t *testing.T) {
	cases := []struct {
		name      string
		disputeId string
		checker   func(*disputes.EvidenceResponse, error)
	}{
		{
			name:      "when dispute evidence exists then return dispute evidence",
			disputeId: disputeId,
			checker: func(response *disputes.EvidenceResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
			},
		},
		{
			name:      "when dispute doesn't exist then return error",
			disputeId: "not_found",
			checker: func(response *disputes.EvidenceResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := DefaultApi().Disputes

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if disputeId != "" {
				tc.checker(client.GetEvidence(tc.disputeId))
			}
		})
	}
}

func TestUploadFile(t *testing.T) {
	cases := []struct {
		name        string
		fileRequest common.File
		checker     func(*common.IdResponse, error)
	}{
		{
			name: "when data is correct then return ID for uploaded file - IMAGE",
			fileRequest: common.File{
				File:    "./checkout.jpeg",
				Purpose: common.DisputesEvidence,
			},
			checker: func(response *common.IdResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.Links)
			},
		},
		{
			name: "when data is correct then return ID for uploaded file - PDF",
			fileRequest: common.File{
				File:    "./checkout.pdf",
				Purpose: common.DisputesEvidence,
			},
			checker: func(response *common.IdResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.Links)
			},
		},
		{
			name:        "when file path is missing then return error",
			fileRequest: common.File{},
			checker: func(response *common.IdResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				assert.Equal(t, "Invalid file name", err.Error())
			},
		},
		{
			name: "when purpose is missing then return error",
			fileRequest: common.File{
				File: "./checkout.pdf",
			},
			checker: func(response *common.IdResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				assert.Equal(t, "Invalid purpose", err.Error())
			},
		},
	}

	client := DefaultApi().Disputes

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.UploadFile(tc.fileRequest))
		})
	}
}

func TestGetFileDetails(t *testing.T) {
	var (
		req = common.File{
			File:    "./checkout.jpeg",
			Purpose: common.DisputesEvidence,
		}
		fileId = uploadDisputeFile(t, req).Id
	)

	cases := []struct {
		name    string
		fileId  string
		checker func(*common.FileResponse, error)
	}{
		{
			name:   "when file exists then return file details",
			fileId: fileId,
			checker: func(response *common.FileResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, fileId, response.Id)
				assert.Equal(t, req.Purpose, response.Purpose)
				assert.Equal(t, filepath.Base(req.File), response.Filename)
				assert.NotNil(t, response.Size)
				assert.NotNil(t, response.UploadedOn)
			},
		},
		{
			name:   "when file does not exist then return error",
			fileId: "not_found",
			checker: func(response *common.FileResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := DefaultApi().Disputes

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetFileDetails(tc.fileId))
		})
	}
}

func TestGetDisputeSchemeFiles(t *testing.T) {
	t.Skip("Skipping tests because this suite is unstable")
	dispute := getDisputes(t).Data[0]

	cases := []struct {
		name      string
		disputeId string
		checker   func(*disputes.SchemeFilesResponse, error)
	}{
		{
			name:      "when dispute has files then return scheme files",
			disputeId: dispute.Id,
			checker: func(response *disputes.SchemeFilesResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, dispute.Id, response.Id)
				assert.NotEmpty(t, response.Files)
				for _, file := range response.Files {
					assert.NotNil(t, file.File)
					assert.NotNil(t, file.DisputeStatus)
				}
			},
		},
		{
			name:      "when dispute does not exist then return error",
			disputeId: "not_found",
			checker: func(response *disputes.SchemeFilesResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := DefaultApi().Disputes

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetDisputeSchemeFiles(tc.disputeId))
		})
	}
}

func getPaymentRequest(t *testing.T, token string) *nas.PaymentResponse {
	tokenSource := sources.NewRequestTokenSource()
	tokenSource.Token = token

	paymentRequest := nas.PaymentRequest{
		Source:      tokenSource,
		Amount:      1040,
		Currency:    common.GBP,
		Reference:   Reference,
		Description: "description",
		Capture:     true,
		Customer: &common.CustomerRequest{
			Email: Email,
			Phone: &common.Phone{},
		},
		BillingDescriptor: &payments.BillingDescriptor{
			Name:      Name,
			City:      "London",
			Reference: Reference,
		},
	}

	response, err := DefaultApi().Payments.RequestPayment(paymentRequest, nil)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("error requesting payment - %s", err.Error()))
	}

	return response
}

func getDispute(paymentId string) (dispute disputes.DisputeSummary) {
	var (
		response *disputes.QueryResponse
		client   = DefaultApi().Disputes
		attempt  = 1
	)

	queryFilter := disputes.QueryFilter{
		Statuses:  string(disputes.EvidenceRequired),
		PaymentId: paymentId,
	}

	for attempt <= MaxRetryAttemps {
		response, _ = client.Query(queryFilter)
		if response != nil && len(response.Data) > 0 {
			return response.Data[0]
		}
		attempt++
		Wait(time.Duration(10))
	}

	return dispute
}

func uploadDisputeFile(t *testing.T, fileRequest common.File) *common.IdResponse {
	response, err := DefaultApi().Disputes.UploadFile(fileRequest)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("error uploading file - %s", err.Error()))
	}

	return response
}

func getDisputes(t *testing.T) *disputes.QueryResponse {
	layout := "2006-01-02T15:04:05Z"
	from, _ := time.Parse(layout, time.Now().AddDate(0, -3, 0).String())
	to, _ := time.Parse(layout, time.Now().Format(layout))

	query := disputes.QueryFilter{
		Limit:           1,
		Skip:            0,
		From:            from,
		To:              to,
		ThisChannelOnly: true,
	}

	process := func() (interface{}, error) {
		return DefaultApi().Disputes.Query(query)
	}
	predicate := func(data interface{}) bool {
		response := data.(*disputes.QueryResponse)
		return response.Data != nil && len(response.Data) >= 0
	}

	response, err := retriable(process, predicate, 2)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("error getting subject events - %s", err.Error()))
	}

	return response.(*disputes.QueryResponse)
}
