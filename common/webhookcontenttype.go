package common

// WebhookContentType ...
type WebhookContentType string

const (
	// JSON ...
	JSON WebhookContentType = "json"
	// XML ...
	XML WebhookContentType = "xml"
)

func (c WebhookContentType) String() string {
	return string(c)
}
