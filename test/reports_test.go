package test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/errors"
	"github.com/checkout/checkout-sdk-go/reports"
)

var (
	report *reports.ReportResponse
)

func TestSetupReports(t *testing.T) {
	report = getReportId(t)
}

func TestGetAllReports(t *testing.T) {
	var (
		layout  = "2006-01-02"
		now     = time.Now()
		from, _ = time.Parse(layout, now.AddDate(0, 0, -10).Format(layout))
		to, _   = time.Parse(layout, now.Format(layout))
	)

	cases := []struct {
		name    string
		request reports.QueryFilter
		checker func(*reports.QueryResponse, error)
	}{
		{
			name: "when reports match filters then return reports",
			request: reports.QueryFilter{
				Limit:         10,
				CreatedAfter:  &from,
				CreatedBefore: &to,
			},
			checker: func(response *reports.QueryResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.Data)
			},
		},
		{
			name: "when invalid filters then return error",
			request: reports.QueryFilter{
				Limit: 255,
			},
			checker: func(response *reports.QueryResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				errChk := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, errChk.StatusCode)
				assert.Equal(t, "invalid_request", errChk.Data.ErrorType)
			},
		},
	}

	client := DefaultApi().Reports

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetAllReports(tc.request))
		})
	}
}

func TestGetReportDetails(t *testing.T) {
	cases := []struct {
		name     string
		reportId string
		checker  func(*reports.ReportResponse, error)
	}{
		{
			name:     "when report exists then return report details",
			reportId: report.Id,
			checker: func(response *reports.ReportResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, report.Id, response.Id)
				assert.Equal(t, report.Type, response.Type)
				assert.Equal(t, report.CreatedOn, response.CreatedOn)
				assert.Equal(t, report.Description, response.Description)
				assert.Equal(t, report.Account, response.Account)
				assert.Equal(t, report.From, response.From)
			},
		},
		{
			name:     "when report doesn't exist then return error",
			reportId: "not_found",
			checker: func(response *reports.ReportResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := DefaultApi().Reports

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetReportDetails(tc.reportId))
		})
	}
}

func TestGetReportFile(t *testing.T) {
	cases := []struct {
		name     string
		reportId string
		fileId   string
		checker  func(*common.ContentResponse, error)
	}{
		{
			name:     "when report file exists then return file download link",
			reportId: report.Id,
			fileId:   report.Files[0].Id,
			checker: func(response *common.ContentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
			},
		},
		{
			name:     "when report doesn't exist then return error",
			reportId: "not_found",
			checker: func(response *common.ContentResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
		{
			name:     "when report file doesn't exist then return error",
			reportId: report.Id,
			fileId:   "not_found",
			checker: func(response *common.ContentResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := DefaultApi().Reports

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetReportFile(tc.reportId, tc.fileId))
		})
	}
}

func getReportId(t *testing.T) *reports.ReportResponse {
	var (
		layout  = "2006-01-02"
		now     = time.Now()
		from, _ = time.Parse(layout, now.AddDate(0, 0, -10).Format(layout))
		to, _   = time.Parse(layout, now.Format(layout))
	)

	query := reports.QueryFilter{
		CreatedAfter:  &from,
		CreatedBefore: &to,
		Limit:         10,
	}

	reportsList, err := DefaultApi().Reports.GetAllReports(query)
	if err != nil || reportsList.Data == nil {
		assert.Fail(t, fmt.Sprintf("error fetching reports - %s", err.Error()))
	}

	return &reportsList.Data[0]
}
