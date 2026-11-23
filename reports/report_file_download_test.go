package reports

import (
	"bytes"
	"context"
	stderrors "errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/checkout/checkout-sdk-go/v3/client"
	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/configuration"
	"github.com/checkout/checkout-sdk-go/v3/errors"
)

// reportMarker is in the download URL; the tests check that the error and the log keep the status
// only.
const reportMarker = "report-marker"

const reportKey = "sk_sbox_report_test"

type downloadHost struct {
	mu sync.Mutex
	// apiAuth is the Authorization header the API server received.
	apiAuth []string
	hits    int
	auth    []string
	bodies  []string
	methods []string
}

func (d *downloadHost) record(r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	d.mu.Lock()
	defer d.mu.Unlock()
	d.hits++
	d.auth = append(d.auth, r.Header.Get("Authorization"))
	d.bodies = append(d.bodies, string(body))
	d.methods = append(d.methods, r.Method)
}

// newReportServers starts an API server whose report file path answers code with a Location on the
// same host (crossHost false) or on a download host (crossHost true). The download path answers
// 200 with CSV content, or one more redirect when secondHop is set.
func newReportServers(t *testing.T, code int, crossHost, secondHop bool) (*httptest.Server, *downloadHost) {
	download := &downloadHost{}
	downloadHandler := func(w http.ResponseWriter, r *http.Request) {
		download.record(r)
		if secondHop && r.URL.Path == "/download" {
			w.Header().Set("Cko-Request-Id", "chosen-by-the-download-host")
			w.Header().Set("Location", "/download-again?q="+reportMarker)
			w.WriteHeader(http.StatusFound)
			return
		}
		w.Header().Set("Content-Type", "text/csv")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("a,b\n1,2\n"))
	}

	var downloadBase string
	if crossHost {
		other := httptest.NewServer(http.HandlerFunc(downloadHandler))
		t.Cleanup(other.Close)
		downloadBase = other.URL
	}

	api := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/download") {
			downloadHandler(w, r)
			return
		}
		download.mu.Lock()
		download.apiAuth = append(download.apiAuth, r.Header.Get("Authorization"))
		download.mu.Unlock()
		w.Header().Set("Cko-Request-Id", "req-report")
		w.Header().Set("Location", downloadBase+"/download?q="+reportMarker)
		w.WriteHeader(code)
	}))
	t.Cleanup(api.Close)

	return api, download
}

func newReportClient(baseUri string, logs *bytes.Buffer) *Client {
	telemetry := false
	config := configuration.NewConfiguration(
		configuration.NewDefaultKeysSdkCredentials(reportKey, "pk_sbox_report_test"),
		&telemetry, nil, nil, log.New(logs, "", 0))
	return NewClient(config, client.NewApiClient(config, baseUri))
}

func assertReportRedirectRefused(t *testing.T, err error, code int, logs string) {
	t.Helper()
	require.Error(t, err)
	apiErr, ok := err.(errors.CheckoutAPIError)
	require.True(t, ok, "want a CheckoutAPIError, got %T: %v", err, err)
	assert.Equal(t, code, apiErr.StatusCode)
	rendered := fmt.Sprintf("%v %+v", err, apiErr)
	if apiErr.Data != nil {
		rendered += fmt.Sprintf(" %+v", *apiErr.Data)
	}
	assert.NotContains(t, rendered, reportMarker)
	assert.NotContains(t, rendered, "/download")
	assert.NotContains(t, logs, reportMarker)
}

// TestGetReportFileFollowsTheDownloadRedirectWithoutAuthorization: the 302 of GET
// /reports/{id}/files/{fileId} is followed once, as a GET without Authorization and without a body,
// to another host and to the same host.
func TestGetReportFileFollowsTheDownloadRedirectWithoutAuthorization(t *testing.T) {
	for _, crossHost := range []bool{true, false} {
		for _, withContext := range []bool{false, true} {
			crossHost, withContext := crossHost, withContext
			t.Run(fmt.Sprintf("crossHost=%v/withContext=%v", crossHost, withContext), func(t *testing.T) {
				api, download := newReportServers(t, http.StatusFound, crossHost, false)
				var logs bytes.Buffer
				reports := newReportClient(api.URL, &logs)

				var response *common.ContentResponse
				var err error
				if withContext {
					response, err = reports.GetReportFileWithContext(context.Background(), "rpt_1", "file_1")
				} else {
					response, err = reports.GetReportFile("rpt_1", "file_1")
				}

				require.NoError(t, err)
				require.NotNil(t, response)
				assert.Equal(t, "a,b\n1,2\n", response.Content, "the content of the download URL")
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, []string{"Bearer " + reportKey}, download.apiAuth, "the API receives the exact Authorization once")
				require.Equal(t, 1, download.hits)
				assert.Equal(t, http.MethodGet, download.methods[0])
				assert.Equal(t, "", download.auth[0], "the download request is sent without Authorization")
				assert.Equal(t, "", download.bodies[0])
				assert.NotContains(t, logs.String(), reportMarker)
				assert.NotContains(t, logs.String(), reportKey)
			})
		}
	}
}

// TestGetReportFileRefusesOtherRedirects: only a 302 is followed; 301, 303, 307 and 308 are returned
// as a CheckoutAPIError and the download host is not reached.
func TestGetReportFileRefusesOtherRedirects(t *testing.T) {
	for _, code := range []int{http.StatusMovedPermanently, http.StatusSeeOther, http.StatusTemporaryRedirect, http.StatusPermanentRedirect} {
		for _, crossHost := range []bool{true, false} {
			code, crossHost := code, crossHost
			t.Run(fmt.Sprintf("%d/crossHost=%v", code, crossHost), func(t *testing.T) {
				api, download := newReportServers(t, code, crossHost, false)
				var logs bytes.Buffer

				_, err := newReportClient(api.URL, &logs).GetReportFile("rpt_1", "file_1")

				assertReportRedirectRefused(t, err, code, logs.String())
				assert.Equal(t, 0, download.hits)
			})
		}
	}
}

// TestGetReportFileRefusesASecondRedirect: the download URL itself redirecting again is returned as
// a CheckoutAPIError; the second Location is not requested.
func TestGetReportFileRefusesASecondRedirect(t *testing.T) {
	api, download := newReportServers(t, http.StatusFound, true, true)
	var logs bytes.Buffer

	_, err := newReportClient(api.URL, &logs).GetReportFile("rpt_1", "file_1")

	assertReportRedirectRefused(t, err, http.StatusFound, logs.String())
	apiErr := err.(errors.CheckoutAPIError)
	require.NotNil(t, apiErr.Data)
	assert.Equal(t, "req-report", apiErr.Data.RequestID, "the request id of the API response, not of the download host")
	assert.Equal(t, 1, download.hits, "only the first download URL is requested")
	assert.Equal(t, "", download.auth[0])
}

// TestOtherReportCallsDoNotFollowRedirects: the exception is limited to GetReportFile.
func TestOtherReportCallsDoNotFollowRedirects(t *testing.T) {
	api, download := newReportServers(t, http.StatusFound, true, false)
	var logs bytes.Buffer

	_, err := newReportClient(api.URL, &logs).GetReportDetails("rpt_1")

	assertReportRedirectRefused(t, err, http.StatusFound, logs.String())
	assert.Equal(t, 0, download.hits)
}

// TestGetReportFileErrorsKeepTheDownloadUrlOut: a failed or timed out download request returns an
// error that keeps its type and Timeout, without the query of the download URL.
func TestGetReportFileErrorsKeepTheDownloadUrlOut(t *testing.T) {
	closed := httptest.NewServer(http.NotFoundHandler())
	closedUrl := closed.URL
	closed.Close()
	slow := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(400 * time.Millisecond)
	}))
	t.Cleanup(slow.Close)

	for name, target := range map[string]string{"closed host": closedUrl, "timeout": slow.URL} {
		target := target
		t.Run(name, func(t *testing.T) {
			api := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Location", target+"/file?q="+reportMarker+"#"+reportMarker)
				w.WriteHeader(http.StatusFound)
			}))
			t.Cleanup(api.Close)
			telemetry := false
			var logs bytes.Buffer
			config := configuration.NewConfiguration(
				configuration.NewDefaultKeysSdkCredentials(reportKey, "pk_sbox_report_test"),
				&telemetry, nil, &http.Client{Timeout: 150 * time.Millisecond}, log.New(&logs, "", 0))

			_, err := NewClient(config, client.NewApiClient(config, api.URL)).GetReportFile("rpt_1", "file_1")

			require.Error(t, err)
			assert.NotContains(t, err.Error(), reportMarker)
			assert.NotContains(t, fmt.Sprintf("%+v", err), reportMarker)
			assert.NotContains(t, logs.String(), reportMarker)
			var urlErr *url.Error
			require.True(t, stderrors.As(err, &urlErr), "the error keeps the *url.Error, got %T", err)
			assert.Contains(t, urlErr.URL, "/file")
			if name == "timeout" {
				timeout, ok := err.(interface{ Timeout() bool })
				require.True(t, ok)
				assert.True(t, timeout.Timeout())
			}
		})
	}
}

// TestGetReportFileSendsNoUserInformationToTheDownloadUrl: user information in the download URL is
// not sent, so the download request carries no Authorization at all.
func TestGetReportFileSendsNoUserInformationToTheDownloadUrl(t *testing.T) {
	download := &downloadHost{}
	host := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		download.record(r)
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(host.Close)
	withUser := strings.Replace(host.URL, "http://", "http://user:"+reportMarker+"@", 1)
	api := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", withUser+"/download")
		w.WriteHeader(http.StatusFound)
	}))
	t.Cleanup(api.Close)
	var logs bytes.Buffer

	_, err := newReportClient(api.URL, &logs).GetReportFile("rpt_1", "file_1")

	require.NoError(t, err)
	require.Equal(t, 1, download.hits)
	assert.Equal(t, "", download.auth[0], "the download request is sent without Authorization")
}

// TestGetReportFileDownloadErrorsKeepTheStatusOnly: a download URL answering an error returns a
// CheckoutAPIError with its status code and the standard status text, without its body.
func TestGetReportFileDownloadErrorsKeepTheStatusOnly(t *testing.T) {
	for _, code := range []int{http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError} {
		code := code
		t.Run(fmt.Sprint(code), func(t *testing.T) {
			host := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/xml")
				w.Header().Set("Cko-Request-Id", "chosen-by-the-download-host")
				w.WriteHeader(code)
				_, _ = w.Write([]byte("<Error><Url>/download?q=" + reportMarker + "</Url></Error>"))
			}))
			t.Cleanup(host.Close)
			api := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Cko-Request-Id", "req-report")
				w.Header().Set("Location", host.URL+"/download")
				w.WriteHeader(http.StatusFound)
			}))
			t.Cleanup(api.Close)
			var logs bytes.Buffer

			_, err := newReportClient(api.URL, &logs).GetReportFile("rpt_1", "file_1")

			require.Error(t, err)
			apiErr, ok := err.(errors.CheckoutAPIError)
			require.True(t, ok, "want a CheckoutAPIError, got %T: %v", err, err)
			assert.Equal(t, code, apiErr.StatusCode)
			assert.Equal(t, fmt.Sprintf("%d %s", code, http.StatusText(code)), apiErr.Status)
			require.NotNil(t, apiErr.Data)
			assert.Equal(t, "req-report", apiErr.Data.RequestID, "the request id of the API response, not of the download host")
			assertReportRedirectRefused(t, err, code, logs.String())
		})
	}
}

// TestGetReportFileDecodingErrorForOtherMediaTypes: a download that is neither text/csv nor
// application/pdf is decoded as JSON; a body that is not JSON returns a CheckoutDecodingError with
// the status code and the media type only, and a JSON body fills the response.
func TestGetReportFileDecodingErrorForOtherMediaTypes(t *testing.T) {
	t.Run("json body", func(t *testing.T) {
		host := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/octet-stream")
			_, _ = w.Write([]byte(`{"content":"x"}`))
		}))
		t.Cleanup(host.Close)
		api := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Location", host.URL+"/download")
			w.WriteHeader(http.StatusFound)
		}))
		t.Cleanup(api.Close)

		response, err := newReportClient(api.URL, &bytes.Buffer{}).GetReportFile("rpt_1", "file_1")

		require.NoError(t, err)
		assert.Equal(t, "x", response.Content)
	})
	for _, mediaType := range []string{"application/octet-stream", "binary/octet-stream"} {
		mediaType := mediaType
		t.Run(mediaType, func(t *testing.T) {
			host := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", mediaType)
				_, _ = w.Write([]byte("a,b\n"))
			}))
			t.Cleanup(host.Close)
			api := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Location", host.URL+"/download?q="+reportMarker)
				w.WriteHeader(http.StatusFound)
			}))
			t.Cleanup(api.Close)
			var logs bytes.Buffer

			_, err := newReportClient(api.URL, &logs).GetReportFile("rpt_1", "file_1")

			require.Error(t, err)
			decodingErr, ok := err.(errors.CheckoutDecodingError)
			require.True(t, ok, "want a CheckoutDecodingError, got %T: %v", err, err)
			assert.Equal(t, http.StatusOK, decodingErr.StatusCode)
			assert.Equal(t, mediaType, decodingErr.ContentType)
			assert.NotContains(t, fmt.Sprintf("%v %+v", err, decodingErr), reportMarker)
			assert.NotContains(t, logs.String(), reportMarker)
		})
	}
}

// TestGetReportFileRefusesAnHttpDownloadUrl: a download URL over http to a host that is not loopback
// is not requested; the 302 returns a CheckoutAPIError without the Location.
func TestGetReportFileRefusesAnHttpDownloadUrl(t *testing.T) {
	for _, location := range []string{
		"http://download.example/download?q=" + reportMarker,
		"http://localhost.download.example/download?q=" + reportMarker,
		"ftp://download.example/download?q=" + reportMarker,
	} {
		location := location
		t.Run(location, func(t *testing.T) {
			sent := 0
			transport := roundTripFunc(func(r *http.Request) (*http.Response, error) {
				sent++
				if r.URL.Host != "api.example" {
					return nil, fmt.Errorf("unexpected request to %s", r.URL.Host)
				}
				header := http.Header{}
				header.Set("Location", location)
				header.Set("Cko-Request-Id", "req-report")
				return &http.Response{StatusCode: http.StatusFound, Header: header,
					Body: ioutil.NopCloser(strings.NewReader("")), Request: r}, nil
			})
			telemetry := false
			var logs bytes.Buffer
			config := configuration.NewConfiguration(
				configuration.NewDefaultKeysSdkCredentials(reportKey, "pk_sbox_report_test"),
				&telemetry, nil, &http.Client{Transport: transport}, log.New(&logs, "", 0))

			_, err := NewClient(config, client.NewApiClient(config, "https://api.example")).GetReportFile("rpt_1", "file_1")

			assertReportRedirectRefused(t, err, http.StatusFound, logs.String())
			assert.Equal(t, 1, sent, "the download URL is not requested")
			assert.NotContains(t, fmt.Sprintf("%v %+v", err, err), "download.example")
			assert.Equal(t, "req-report", err.(errors.CheckoutAPIError).Data.RequestID)
		})
	}
}

// TestGetReportFileSendsNoSdkHeadersToTheDownloadUrl: with telemetry on, the API receives the
// telemetry header and the https download host receives no Cko- header, no Authorization and no
// Referer.
func TestGetReportFileSendsNoSdkHeadersToTheDownloadUrl(t *testing.T) {
	var mu sync.Mutex
	var downloadCko, downloadAuth, downloadReferer, apiTelemetry []string
	host := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		for name := range r.Header {
			if strings.HasPrefix(strings.ToLower(name), "cko-") {
				downloadCko = append(downloadCko, name)
			}
		}
		downloadAuth = append(downloadAuth, r.Header.Get("Authorization"))
		downloadReferer = append(downloadReferer, r.Header.Get("Referer"))
		mu.Unlock()
		w.Header().Set("Content-Type", "text/csv")
		_, _ = w.Write([]byte("a,b\n"))
	}))
	t.Cleanup(host.Close)
	api := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		apiTelemetry = append(apiTelemetry, r.Header.Get(client.CkoTelemetryHeader))
		mu.Unlock()
		w.Header().Set("Cko-Request-Id", "req-report")
		w.Header().Set("Location", host.URL+"/download?q="+reportMarker)
		w.WriteHeader(http.StatusFound)
	}))
	t.Cleanup(api.Close)
	telemetry := true
	var logs bytes.Buffer
	config := configuration.NewConfiguration(
		configuration.NewDefaultKeysSdkCredentials(reportKey, "pk_sbox_report_test"),
		&telemetry, nil, host.Client(), log.New(&logs, "", 0))
	reports := NewClient(config, client.NewApiClient(config, api.URL))

	for i := 0; i < 3; i++ {
		response, err := reports.GetReportFile("rpt_1", "file_1")
		require.NoError(t, err)
		require.NotNil(t, response)
		assert.Equal(t, "a,b\n", response.Content, "the content of the download URL")
	}

	require.Len(t, apiTelemetry, 3)
	assert.Equal(t, "", apiTelemetry[0], "no metric before the first request")
	assert.NotEqual(t, "", apiTelemetry[1], "the API receives the telemetry header")
	assert.NotEqual(t, "", apiTelemetry[2], "the API receives the telemetry header")
	assert.Equal(t, []string{"", "", ""}, downloadAuth, "three downloads, without Authorization")
	assert.Empty(t, downloadCko, "the download host receives no Cko- header")
	assert.Equal(t, []string{"", "", ""}, downloadReferer, "the download host receives no Referer")
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }
