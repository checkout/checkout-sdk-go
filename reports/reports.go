package reports

import (
	"time"

	"github.com/checkout/checkout-sdk-go/common"
)

const (
	reports = "reports"
	files   = "files"
)

type (
	QueryFilter struct {
		CreatedAfter    *time.Time `url:"created_after,omitempty" layout:"2006-01-02"`
		CreatedBefore   *time.Time `url:"created_before,omitempty" layout:"2006-01-02"`
		EntityId        string     `url:"entity_id,omitempty"`
		Limit           int        `url:"limit,omitempty"`
		PaginationToken string     `url:"pagination_token,omitempty"`
	}
)

type (
	ReportResponse struct {
		HttpMetadata   common.HttpMetadata
		Id             string                 `json:"id,omitempty"`
		CreatedOn      string                 `json:"created_on,omitempty"`
		LastModifiedOn string                 `json:"last_modified_on,omitempty"`
		Type           string                 `json:"type,omitempty"`
		Description    string                 `json:"description,omitempty"`
		Account        *Account               `json:"account,omitempty"`
		Tags           []string               `json:"tags,omitempty"`
		From           *time.Time             `json:"from,omitempty"`
		To             *time.Time             `json:"to,omitempty"`
		Files          []File                 `json:"files,omitempty"`
		Links          map[string]common.Link `json:"_links"`
	}

	QueryResponse struct {
		HttpMetadata common.HttpMetadata
		Count        int                    `json:"count,omitempty"`
		Limit        uint8                  `json:"limit,omitempty"`
		Data         []ReportResponse       `json:"data,omitempty"`
		Links        map[string]common.Link `json:"_links"`
	}
)

type Account struct {
	ClientId string `json:"client_id,omitempty"`
	EntityId string `json:"entity_id,omitempty"`
}

type File struct {
	Id       string                 `json:"id,omitempty"`
	Filename string                 `json:"filename,omitempty"`
	Format   string                 `json:"format,omitempty"`
	Links    map[string]common.Link `json:"_links"`
}
