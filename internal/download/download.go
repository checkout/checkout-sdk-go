// Package download marks, inside this module only, the request whose file download URL the
// ApiClient may request. No other request is marked.
package download

import "context"

type markKey struct{}

// Mark returns a copy of ctx that lets the ApiClient request, once, the file download URL the API
// answers with, without the SDK credentials and without a body. Only reports GetReportFile uses it.
func Mark(ctx context.Context) context.Context {
	return context.WithValue(ctx, markKey{}, true)
}

// IsMarked reports whether ctx was returned by Mark, or derives from such a context.
func IsMarked(ctx context.Context) bool {
	marked, _ := ctx.Value(markKey{}).(bool)
	return marked
}
