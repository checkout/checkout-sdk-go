package client

import (
	"bytes"
	"context"
	stderrors "errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/configuration"
	sdkerrors "github.com/checkout/checkout-sdk-go/v3/errors"
	"github.com/checkout/checkout-sdk-go/v3/internal/download"
)

// redirectMarker is in every request body sent by these tests, in the Location and in the 3xx
// response body; the tests check that the error and the log keep the status only.
const redirectMarker = "redirect-marker"

const redirectRequestId = "req-3xx"

var redirectCodes = []int{
	http.StatusMovedPermanently,
	http.StatusFound,
	http.StatusSeeOther,
	http.StatusTemporaryRedirect,
	http.StatusPermanentRedirect,
}

// redirectTarget records every request that reaches the redirect target.
type redirectTarget struct {
	mu       sync.Mutex
	hits     int
	auth     []string
	bodies   []string
	methods  []string
	finalUrl string
	// originAuth is the Authorization header the API server received.
	originAuth []string
}

func (r *redirectTarget) record(req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	r.mu.Lock()
	defer r.mu.Unlock()
	r.hits++
	r.auth = append(r.auth, req.Header.Get("Authorization"))
	r.bodies = append(r.bodies, string(body))
	r.methods = append(r.methods, req.Method)
}

// newRedirectServers starts an API server that answers every request with code and a Location on
// the same host (crossHost false) or on a second server (crossHost true). The target answers 200.
func newRedirectServers(t *testing.T, code int, crossHost bool) (*httptest.Server, *redirectTarget) {
	target := &redirectTarget{}
	targetHandler := func(w http.ResponseWriter, r *http.Request) {
		target.record(r)
		jsonOK(w)
	}

	var targetBase string
	if crossHost {
		other := httptest.NewServer(http.HandlerFunc(targetHandler))
		t.Cleanup(other.Close)
		targetBase = other.URL
	}

	origin := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/target") {
			targetHandler(w, r)
			return
		}
		target.mu.Lock()
		target.originAuth = append(target.originAuth, r.Header.Get("Authorization"))
		target.mu.Unlock()
		w.Header().Set(CkoRequestId, redirectRequestId)
		w.Header().Set("Location", targetBase+"/target?q="+redirectMarker)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		_, _ = w.Write([]byte(`{"value":"` + redirectMarker + `"}`))
	}))
	t.Cleanup(origin.Close)
	if !crossHost {
		target.finalUrl = origin.URL + "/target"
	} else {
		target.finalUrl = targetBase + "/target"
	}

	return origin, target
}

type redirectCall struct {
	name string
	do   func(ctx context.Context, a *ApiClient) error
}

func redirectCalls() []redirectCall {
	body := map[string]string{"value": redirectMarker}
	return []redirectCall{
		{"GET", func(ctx context.Context, a *ApiClient) error {
			return a.GetWithContext(ctx, "/resource", testAuth(), nil, &common.IdResponse{})
		}},
		{"POST", func(ctx context.Context, a *ApiClient) error {
			return a.PostWithContext(ctx, "/resource", testAuth(), body, &common.IdResponse{}, nil)
		}},
		{"PUT", func(ctx context.Context, a *ApiClient) error {
			return a.PutWithContext(ctx, "/resource", testAuth(), body, &common.IdResponse{}, nil)
		}},
		{"PATCH", func(ctx context.Context, a *ApiClient) error {
			return a.PatchWithContext(ctx, "/resource", testAuth(), body, &common.IdResponse{})
		}},
		{"DELETE", func(ctx context.Context, a *ApiClient) error {
			return a.DeleteWithContext(ctx, "/resource", testAuth(), &common.IdResponse{})
		}},
		{"POST form", func(ctx context.Context, a *ApiClient) error {
			form := url.Values{"value": {redirectMarker}}
			return a.PostFormWithContext(ctx, "/resource", testAuth(), form, &common.IdResponse{})
		}},
		{"POST upload", func(ctx context.Context, a *ApiClient) error {
			buffer := new(bytes.Buffer)
			writer := multipart.NewWriter(buffer)
			_ = writer.WriteField("value", redirectMarker)
			_ = writer.Close()
			request := &common.FileUploadRequest{W: writer, B: buffer}
			return a.UploadWithContext(ctx, "/resource", testAuth(), request, &common.IdResponse{})
		}},
	}
}

// allowAllRedirects is a caller policy that would follow every redirect.
func allowAllRedirects(*http.Request, []*http.Request) error { return nil }

type redirectClient struct {
	name  string
	build func(baseUri string, telemetry bool, logger configuration.StdLogger) *ApiClient
}

func redirectClients() []redirectClient {
	return []redirectClient{
		{"default client", func(baseUri string, telemetry bool, logger configuration.StdLogger) *ApiClient {
			return NewApiClient(configuration.NewConfiguration(nil, &telemetry, nil, nil, logger), baseUri)
		}},
		{"caller client following every redirect", func(baseUri string, telemetry bool, logger configuration.StdLogger) *ApiClient {
			caller := &http.Client{Timeout: 5 * time.Second, CheckRedirect: allowAllRedirects}
			return NewApiClient(configuration.NewConfiguration(nil, &telemetry, nil, caller, logger), baseUri)
		}},
		{"ApiClient literal", func(baseUri string, telemetry bool, logger configuration.StdLogger) *ApiClient {
			return &ApiClient{
				HttpClient:          http.Client{},
				BaseUri:             baseUri,
				EnableTelemetry:     telemetry,
				RequestMetricsQueue: *common.NewTelemetryQueue(),
				Log:                 logger,
			}
		}},
	}
}

func assertRedirectRefused(t *testing.T, err error, code int, target *redirectTarget, logs string) {
	t.Helper()
	assert.Equal(t, 0, target.hits, "a redirect is not followed (Authorization %v, bodies %v)", target.auth, target.bodies)

	require.Error(t, err)
	apiErr, ok := err.(sdkerrors.CheckoutAPIError)
	require.True(t, ok, "want a CheckoutAPIError, got %T: %v", err, err)
	assert.Equal(t, code, apiErr.StatusCode)
	assert.Equal(t, fmt.Sprintf("%d %s", code, http.StatusText(code)), apiErr.Status)
	require.NotNil(t, apiErr.Data)
	assert.Equal(t, redirectRequestId, apiErr.Data.RequestID)

	rendered := fmt.Sprintf("%v %+v %+v", err, apiErr, *apiErr.Data)
	assert.NotContains(t, rendered, redirectMarker, "the error keeps the status only")
	assert.NotContains(t, rendered, "/target", "the error keeps the status only")
	assert.NotContains(t, logs, redirectMarker, "the log keeps the request line only")
	assert.NotContains(t, logs, "/target", "the log keeps the request line only")
}

// TestApiClientDoesNotFollowRedirects covers every 3xx, same host and another host, every request
// method of the ApiClient, the default, a caller-supplied and a literal client, with and without
// telemetry (both Do calls in doRequest).
func TestApiClientDoesNotFollowRedirects(t *testing.T) {
	for _, code := range redirectCodes {
		for _, crossHost := range []bool{false, true} {
			for _, c := range redirectClients() {
				for _, telemetry := range []bool{false, true} {
					for _, call := range redirectCalls() {
						code, crossHost, c, telemetry, call := code, crossHost, c, telemetry, call
						name := fmt.Sprintf("%d/crossHost=%v/%s/telemetry=%v/%s", code, crossHost, c.name, telemetry, call.name)
						t.Run(name, func(t *testing.T) {
							origin, target := newRedirectServers(t, code, crossHost)
							var logs bytes.Buffer
							api := c.build(origin.URL, telemetry, log.New(&logs, "", 0))

							err := call.do(context.Background(), api)

							assertRedirectRefused(t, err, code, target, logs.String())
							assert.Equal(t, []string{"Bearer test-token"}, target.originAuth, "the API receives the exact Authorization once")
						})
					}
				}
			}
		}
	}
}

// TestApiClientRedirectPolicyLeavesTheCallerClientUnchanged checks that the caller's http.Client,
// the copy in the configuration and the ApiClient field keep their own CheckRedirect and Timeout,
// and that the caller policy is not consulted for an API call.
func TestApiClientRedirectPolicyLeavesTheCallerClientUnchanged(t *testing.T) {
	consulted := 0
	policy := func(*http.Request, []*http.Request) error {
		consulted++
		return nil
	}
	caller := &http.Client{Timeout: 7 * time.Second, CheckRedirect: policy}
	telemetry := false
	var logs bytes.Buffer
	config := configuration.NewConfiguration(nil, &telemetry, nil, caller, log.New(&logs, "", 0))

	origin, target := newRedirectServers(t, http.StatusTemporaryRedirect, true)
	api := NewApiClient(config, origin.URL)

	err := api.PostWithContext(context.Background(), "/resource", testAuth(), map[string]string{"value": redirectMarker}, &common.IdResponse{}, nil)

	assertRedirectRefused(t, err, http.StatusTemporaryRedirect, target, logs.String())
	assert.Equal(t, 0, consulted, "the SDK redirect policy decides, not the caller policy")
	policyPointer := reflect.ValueOf(policy).Pointer()
	assert.Equal(t, policyPointer, reflect.ValueOf(caller.CheckRedirect).Pointer())
	assert.Equal(t, 7*time.Second, caller.Timeout)
	assert.Equal(t, policyPointer, reflect.ValueOf(config.HttpClient.CheckRedirect).Pointer())
	assert.Equal(t, policyPointer, reflect.ValueOf(api.HttpClient.CheckRedirect).Pointer())
}

// TestApiClientRedirectPolicyLeavesTheDefaultClientUnchanged checks that BuildDefaultClient and the
// ApiClient field keep no CheckRedirect: the policy is applied on a copy for each request.
func TestApiClientRedirectPolicyLeavesTheDefaultClientUnchanged(t *testing.T) {
	assert.Nil(t, common.BuildDefaultClient().CheckRedirect)

	origin, _ := newRedirectServers(t, http.StatusFound, false)
	api := newTestClient(origin.URL)
	_ = api.GetWithContext(context.Background(), "/resource", testAuth(), nil, &common.IdResponse{})

	assert.Nil(t, api.HttpClient.CheckRedirect)
}

// TestFollowDownloadRedirectRefusesWhatItDoesNotFollow checks the download policy directly: one
// 302 without a body is followed with Authorization removed; a body, another status, a second hop
// or a missing redirect response are refused.
func TestFollowDownloadRedirectRefusesWhatItDoesNotFollow(t *testing.T) {
	newHop := func(status int, body string) *http.Request {
		var reader *strings.Reader
		req, _ := http.NewRequest(http.MethodGet, "https://download.example/file", nil)
		if body != "" {
			reader = strings.NewReader(body)
			req.Body = ioutil.NopCloser(reader)
		}
		req.Header.Set("Authorization", "Bearer "+redirectMarker)
		if status != 0 {
			req.Response = &http.Response{StatusCode: status}
		}
		return req
	}
	first, _ := http.NewRequest(http.MethodGet, "https://api.example/reports/r/files/f", nil)

	followed := newHop(http.StatusFound, "")
	assert.NoError(t, followDownloadRedirect(followed, []*http.Request{first}))
	assert.Equal(t, "", followed.Header.Get("Authorization"))

	assert.Equal(t, http.ErrUseLastResponse, followDownloadRedirect(newHop(http.StatusFound, redirectMarker), []*http.Request{first}), "a body")
	assert.Equal(t, http.ErrUseLastResponse, followDownloadRedirect(newHop(http.StatusTemporaryRedirect, ""), []*http.Request{first}), "a 307")
	assert.Equal(t, http.ErrUseLastResponse, followDownloadRedirect(newHop(http.StatusFound, ""), []*http.Request{first, first}), "a second hop")
	assert.Equal(t, http.ErrUseLastResponse, followDownloadRedirect(newHop(0, ""), []*http.Request{first}), "no redirect response")
	assert.Equal(t, http.ErrUseLastResponse, refuseRedirect(newHop(http.StatusFound, ""), []*http.Request{first}))
}

// TestFollowDownloadRedirectFollowsHttpsOrLoopbackOnly: the download URL is followed only over
// https, or over http to a loopback host; any other scheme or host is returned as the response.
func TestFollowDownloadRedirectFollowsHttpsOrLoopbackOnly(t *testing.T) {
	first, _ := http.NewRequest(http.MethodGet, "https://api.example/reports/r/files/f", nil)
	targets := map[string]bool{
		"https://download.example/file?sig=" + redirectMarker: true,
		"https://127.0.0.1/file":                              true,
		"http://127.0.0.1:8080/file":                          true,
		"http://localhost/file":                               true,
		"http://[::1]/file":                                   true,
		"http://download.example/file?sig=" + redirectMarker:  false,
		"http://localhost.evil.example/file":                  false,
		"http://127.0.0.1.nip.io/file":                        false,
		"http://0.0.0.0/file":                                 false,
		"http://10.0.0.1/file":                                false,
		"ftp://download.example/file":                         false,
		"file:///etc/passwd":                                  false,
		"ws://127.0.0.1/file":                                 false,
	}
	for target, follow := range targets {
		req, err := http.NewRequest(http.MethodGet, target, nil)
		require.NoError(t, err, target)
		req.Response = &http.Response{StatusCode: http.StatusFound}

		got := followDownloadRedirect(req, []*http.Request{first})

		if follow {
			assert.NoError(t, got, target)
		} else {
			assert.Equal(t, http.ErrUseLastResponse, got, target)
		}
	}
}

// TestFollowDownloadRedirectDropsTheSdkHeaders: the download request carries no Authorization, no
// Referer and no Cko- header, whatever its case; other headers are kept.
func TestFollowDownloadRedirectDropsTheSdkHeaders(t *testing.T) {
	first, _ := http.NewRequest(http.MethodGet, "https://api.example/reports/r/files/f", nil)
	req, _ := http.NewRequest(http.MethodGet, "https://download.example/file", nil)
	req.Response = &http.Response{StatusCode: http.StatusFound}
	req.Header.Set("Authorization", "Bearer "+redirectMarker)
	req.Header.Set(CkoTelemetryHeader, redirectMarker)
	req.Header.Set("Cko-Idempotency-Key", redirectMarker)
	req.Header["cko-lower-case"] = []string{redirectMarker}
	req.Header["CKO-UPPER-CASE"] = []string{redirectMarker}
	req.Header.Set("Accept", "text/csv")
	req.Header.Set("Referer", "https://api.example/reports/r/files/f?q="+redirectMarker)

	require.NoError(t, followDownloadRedirect(req, []*http.Request{first}))

	for name := range req.Header {
		assert.False(t, strings.HasPrefix(strings.ToLower(name), "cko-"), name)
	}
	assert.Equal(t, "", req.Header.Get("Authorization"))
	assert.Equal(t, "", req.Header.Get("Referer"))
	assert.Equal(t, "text/csv", req.Header.Get("Accept"))
}

// TestApiClientRefusesRedirectsWithAMalformedOrMissingLocation: a 3xx with a Location that does not
// parse, or with no Location, gives the same CheckoutAPIError as any other 3xx.
func TestApiClientRefusesRedirectsWithAMalformedOrMissingLocation(t *testing.T) {
	locations := map[string]string{
		"control character": "http://other.example/" + redirectMarker + "\x7f",
		"bad escape":        "http://other.example/%zz" + redirectMarker,
		"bad escape, query": "https://other.example/%zz?sig=" + redirectMarker,
		"bad port":          "http://other.example:" + redirectMarker + "/",
		"missing":           "",
	}
	for name, location := range locations {
		for _, code := range redirectCodes {
			for _, marked := range []bool{false, true} {
				name, location, code, marked := name, location, code, marked
				t.Run(fmt.Sprintf("%s/%d/download=%v", name, code, marked), func(t *testing.T) {
					target := &redirectTarget{}
					origin := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						w.Header().Set(CkoRequestId, redirectRequestId)
						if location != "" {
							w.Header()["Location"] = []string{location}
						}
						w.WriteHeader(code)
					}))
					t.Cleanup(origin.Close)
					var logs bytes.Buffer
					api := newTestClient(origin.URL)
					api.Log = log.New(&logs, "", 0)
					ctx := context.Background()
					if marked {
						ctx = download.Mark(ctx)
					}

					err := api.GetWithContext(ctx, "/resource", testAuth(), nil, &common.IdResponse{})

					assertRedirectRefused(t, err, code, target, logs.String())
				})
			}
		}
	}
}

// TestApiClientRefusesRedirectsFromACallerTransport: a 3xx answered through the caller's own
// Transport is refused the same way, and the caller's Transport is not replaced.
func TestApiClientRefusesRedirectsFromACallerTransport(t *testing.T) {
	sent := 0
	transport := roundTripFunc(func(r *http.Request) (*http.Response, error) {
		sent++
		header := http.Header{}
		header.Set("Location", "https://other.example/target?q="+redirectMarker)
		header.Set(CkoRequestId, redirectRequestId)
		return &http.Response{StatusCode: http.StatusPermanentRedirect, Header: header,
			Body: ioutil.NopCloser(strings.NewReader(redirectMarker)), Request: r}, nil
	})
	caller := &http.Client{Transport: transport}
	telemetry := false
	var logs bytes.Buffer
	api := NewApiClient(configuration.NewConfiguration(nil, &telemetry, nil, caller, log.New(&logs, "", 0)), "https://api.example")

	err := api.PostWithContext(context.Background(), "/resource", testAuth(), map[string]string{"value": redirectMarker}, &common.IdResponse{}, nil)

	assertRedirectRefused(t, err, http.StatusPermanentRedirect, &redirectTarget{}, logs.String())
	assert.Equal(t, 1, sent, "only the API request is sent")
	_, stillCaller := caller.Transport.(roundTripFunc)
	assert.True(t, stillCaller)
	_, stillApi := api.HttpClient.Transport.(roundTripFunc)
	assert.True(t, stillApi)
}

// TestHandleResponseTurnsA3xxIntoACheckoutAPIError: the last check, for a 3xx that would reach
// handleResponse, gives the same error without reading the body.
func TestHandleResponseTurnsA3xxIntoACheckoutAPIError(t *testing.T) {
	header := http.Header{}
	header.Set(CkoRequestId, redirectRequestId)
	header.Set("Location", "https://other.example/target?q="+redirectMarker)
	resp := &http.Response{StatusCode: http.StatusFound, Status: "302 " + redirectMarker, Header: header,
		Body: ioutil.NopCloser(strings.NewReader(redirectMarker))}

	err := newTestClient("https://api.example").handleResponse(context.Background(), resp, &common.IdResponse{})

	assertRedirectRefused(t, err, http.StatusFound, &redirectTarget{}, "")
}

// TestRedirectGuardHandlesEmptyTransportAnswers: a caller Transport answering a 3xx with a nil Body
// gives the CheckoutAPIError and no panic. (A nil response with a nil error is left to net/http,
// which panics on it in Go 1.14 as on master.)
func TestRedirectGuardHandlesEmptyTransportAnswers(t *testing.T) {
	telemetry := false
	build := func(transport http.RoundTripper) *ApiClient {
		caller := &http.Client{Transport: transport}
		return NewApiClient(configuration.NewConfiguration(nil, &telemetry, nil, caller, log.New(ioutil.Discard, "", 0)), "https://api.example")
	}

	nilBody := build(roundTripFunc(func(r *http.Request) (*http.Response, error) {
		header := http.Header{}
		header.Set(CkoRequestId, redirectRequestId)
		return &http.Response{StatusCode: http.StatusFound, Header: header, Request: r}, nil
	}))
	err := nilBody.GetWithContext(context.Background(), "/resource", testAuth(), nil, &common.IdResponse{})
	assertRedirectRefused(t, err, http.StatusFound, &redirectTarget{}, "")
}

// TestApiClientKeepsTheClientTimeout: the Timeout of the configured client still applies to a
// request, marked for the download or not.
func TestApiClientKeepsTheClientTimeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(400 * time.Millisecond)
		w.Header().Set("Location", "/target")
		w.WriteHeader(http.StatusFound)
	}))
	t.Cleanup(server.Close)
	telemetry := false
	caller := &http.Client{Timeout: 100 * time.Millisecond}
	api := NewApiClient(configuration.NewConfiguration(nil, &telemetry, nil, caller, log.New(ioutil.Discard, "", 0)), server.URL)

	for _, ctx := range []context.Context{context.Background(), download.Mark(context.Background())} {
		start := time.Now()
		err := api.GetWithContext(ctx, "/resource", testAuth(), nil, &common.IdResponse{})
		elapsed := time.Since(start)

		require.Error(t, err)
		timeout, ok := err.(interface{ Timeout() bool })
		require.True(t, ok, "want a timeout error, got %T: %v", err, err)
		assert.True(t, timeout.Timeout())
		assert.Less(t, int64(elapsed), int64(350*time.Millisecond))
	}
}

// TestApiClientLetsACustomTransportApplyItsOwnBehaviour: a response the caller Transport returns as
// a 2xx is decoded as a success; the SDK applies its policy to what the Transport returns.
func TestApiClientLetsACustomTransportApplyItsOwnBehaviour(t *testing.T) {
	transport := roundTripFunc(func(r *http.Request) (*http.Response, error) {
		header := http.Header{}
		header.Set("Content-Type", "application/json")
		return &http.Response{StatusCode: http.StatusOK, Header: header,
			Body: ioutil.NopCloser(strings.NewReader(`{"id":"from-transport"}`)), Request: r}, nil
	})
	telemetry := false
	api := NewApiClient(configuration.NewConfiguration(nil, &telemetry, nil, &http.Client{Transport: transport}, log.New(ioutil.Discard, "", 0)), "https://api.example")

	var response common.IdResponse
	err := api.GetWithContext(context.Background(), "/resource", testAuth(), nil, &response)

	require.NoError(t, err)
	assert.Equal(t, "from-transport", response.Id)
}

// TestStatusOnlyErrorHasNoTrailingSpace: a status without a standard text gives the code only.
func TestStatusOnlyErrorHasNoTrailingSpace(t *testing.T) {
	assert.Equal(t, "306", statusOnlyError(306, "").Status)
	assert.Equal(t, "399", statusOnlyError(399, "").Status)
	assert.Equal(t, "302 Found", statusOnlyError(http.StatusFound, "").Status)
}

type cancelRecorder struct {
	http.RoundTripper
	cancelled []*http.Request
}

func (c *cancelRecorder) CancelRequest(req *http.Request) { c.cancelled = append(c.cancelled, req) }

// TestRedirectGuardForwardsCancelRequest: a Transport that supports CancelRequest still receives it.
func TestRedirectGuardForwardsCancelRequest(t *testing.T) {
	recorder := &cancelRecorder{RoundTripper: http.DefaultTransport}
	req, _ := http.NewRequest(http.MethodGet, "https://api.example/resource", nil)

	(&redirectGuard{next: recorder}).CancelRequest(req)
	(&redirectGuard{next: roundTripFunc(func(*http.Request) (*http.Response, error) { return nil, nil })}).CancelRequest(req)

	assert.Equal(t, []*http.Request{req}, recorder.cancelled)
}

// TestDownloadErrorsKeepTheStatusOnlyWithACallerTransport: an error answered by the download URL
// gives the status only, also when the caller Transport returns a response without its Request.
func TestDownloadErrorsKeepTheStatusOnlyWithACallerTransport(t *testing.T) {
	transport := roundTripFunc(func(r *http.Request) (*http.Response, error) {
		header := http.Header{}
		if r.URL.Host == "api.example" {
			header.Set("Location", "https://download.example/file?q="+redirectMarker)
			return &http.Response{StatusCode: http.StatusFound, Header: header, Body: ioutil.NopCloser(strings.NewReader("")), Request: r}, nil
		}
		return &http.Response{StatusCode: http.StatusInternalServerError, Header: header,
			Body: ioutil.NopCloser(strings.NewReader("<Url>https://download.example/file?q=" + redirectMarker + "</Url>"))}, nil
	})
	telemetry := false
	var logs bytes.Buffer
	api := NewApiClient(configuration.NewConfiguration(nil, &telemetry, nil, &http.Client{Transport: transport}, log.New(&logs, "", 0)), "https://api.example")

	err := api.GetWithContext(download.Mark(context.Background()), "/reports/r/files/f", testAuth(), nil, &common.IdResponse{})

	require.Error(t, err)
	apiErr, ok := err.(sdkerrors.CheckoutAPIError)
	require.True(t, ok, "want a CheckoutAPIError, got %T: %v", err, err)
	assert.Equal(t, http.StatusInternalServerError, apiErr.StatusCode)
	assert.Equal(t, "500 Internal Server Error", apiErr.Status)
	assert.NotContains(t, fmt.Sprintf("%v %+v", err, apiErr), redirectMarker)
	assert.NotContains(t, logs.String(), redirectMarker)
}

// TestDownloadBodyReadErrorsKeepTheDownloadUrlOut: when reading the download response fails with an
// error that carries the download URL, the CheckoutTransportError keeps the URL without its query.
func TestDownloadBodyReadErrorsKeepTheDownloadUrlOut(t *testing.T) {
	for _, telemetry := range []bool{false, true} {
		transport := roundTripFunc(func(r *http.Request) (*http.Response, error) {
			readErr := &url.Error{Op: "Get", URL: "https://download.example/file?q=" + redirectMarker, Err: io.ErrUnexpectedEOF}
			return &http.Response{StatusCode: http.StatusOK, Status: "200 OK", Header: http.Header{},
				Body: urlErrorBody{err: readErr}, Request: r}, nil
		})
		var logs bytes.Buffer
		api := NewApiClient(configuration.NewConfiguration(nil, &telemetry, nil, &http.Client{Transport: transport}, log.New(&logs, "", 0)), "https://api.example")

		err := api.GetWithContext(download.Mark(context.Background()), "/reports/r/files/f", testAuth(), nil, &common.IdResponse{})

		require.Error(t, err, "telemetry=%v", telemetry)
		transportErr, ok := err.(sdkerrors.CheckoutTransportError)
		require.True(t, ok, "want a CheckoutTransportError, got %T: %v", err, err)
		assert.NotContains(t, fmt.Sprintf("%v %+v", err, transportErr), redirectMarker, "telemetry=%v", telemetry)
		assert.Contains(t, err.Error(), "https://download.example/file", "telemetry=%v", telemetry)
		assert.True(t, stderrors.Is(err, io.ErrUnexpectedEOF), "telemetry=%v", telemetry)
		assert.NotContains(t, logs.String(), redirectMarker)
	}
}

type urlErrorBody struct{ err error }

func (b urlErrorBody) Read([]byte) (int, error) { return 0, b.err }
func (urlErrorBody) Close() error               { return nil }
