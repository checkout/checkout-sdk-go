package client

import (
	"context"
	"crypto/x509"
	"errors"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"time"

	"github.com/checkout/checkout-sdk-go/v2/configuration"
)

// Patterns used to recognise non-transient transport failures that must not be
// retried.
var (
	redirectsErrorRe = regexp.MustCompile(`stopped after \d+ redirects\z`)
	schemeErrorRe    = regexp.MustCompile(`unsupported protocol scheme`)
)

// drainCap bounds how much of a discarded response body is read before a retry
// so the underlying connection can be reused without buffering large payloads.
const drainCap = 4 << 10

// shouldRetry reports whether the attempt that produced resp/err should be
// retried under cfg. It retries connection-level failures and 409, 429 and 5xx
// responses.
func shouldRetry(resp *http.Response, err error, attempt int, cfg *configuration.RetryConfiguration) bool {
	if attempt >= cfg.MaxRetries {
		return false
	}
	if err != nil {
		return isRetryableError(err)
	}
	if resp == nil {
		return false
	}
	if resp.StatusCode == http.StatusConflict || resp.StatusCode == http.StatusTooManyRequests {
		return true
	}
	return resp.StatusCode >= http.StatusInternalServerError
}

// isRetryableError distinguishes transient transport errors from permanent ones.
// Context cancellation/deadline, TLS trust failures, redirect loops and
// unsupported schemes are treated as permanent.
func isRetryableError(err error) bool {
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return false
	}
	var authErr x509.UnknownAuthorityError
	if errors.As(err, &authErr) {
		return false
	}
	var urlErr *url.Error
	if errors.As(err, &urlErr) {
		msg := urlErr.Error()
		if redirectsErrorRe.MatchString(msg) || schemeErrorRe.MatchString(msg) {
			return false
		}
	}
	return true
}

// backoff returns the delay to wait before the next attempt. A server-provided
// Retry-After hint takes precedence (capped at MaxBackoff); otherwise it applies
// exponential growth (min + min*attempt²) capped at MaxBackoff, jittered to
// 75-100% to avoid synchronised retries. Jitter is not security-sensitive, so
// math/rand is appropriate here.
func backoff(attempt int, cfg *configuration.RetryConfiguration, retryAfter time.Duration) time.Duration {
	if retryAfter > 0 {
		if retryAfter > cfg.MaxBackoff {
			return cfg.MaxBackoff
		}
		return retryAfter
	}
	d := cfg.MinBackoff + cfg.MinBackoff*time.Duration(attempt*attempt)
	if d > cfg.MaxBackoff {
		d = cfg.MaxBackoff
	}
	jitter := 0.75 + rand.Float64()*0.25
	return time.Duration(float64(d) * jitter)
}

// retryAfterDelay parses the Retry-After response header, supporting both the
// delta-seconds and HTTP-date forms. It returns 0 when the header is absent or
// unparseable.
func retryAfterDelay(resp *http.Response) time.Duration {
	if resp == nil {
		return 0
	}
	v := resp.Header.Get("Retry-After")
	if v == "" {
		return 0
	}
	if secs, err := strconv.Atoi(v); err == nil && secs >= 0 {
		return time.Duration(secs) * time.Second
	}
	if t, err := http.ParseTime(v); err == nil {
		if d := time.Until(t); d > 0 {
			return d
		}
	}
	return 0
}

// drainAndClose consumes and closes a response body that is about to be
// discarded so the connection can be kept alive for the retry.
func drainAndClose(resp *http.Response) {
	if resp == nil || resp.Body == nil {
		return
	}
	_, _ = io.Copy(ioutil.Discard, io.LimitReader(resp.Body, drainCap))
	_ = resp.Body.Close()
}

// sleep waits for d or until ctx is cancelled, whichever comes first, so backoff
// respects request deadlines.
func sleep(ctx context.Context, d time.Duration) error {
	if d <= 0 {
		return nil
	}
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-t.C:
		return nil
	}
}
