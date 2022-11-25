package ideal

import "github.com/checkout/checkout-sdk-go-beta/common"

const (
	idealExternalPath = "ideal-external"
	issuersPath       = "issuers"
)

type (
	IdealInfo struct {
		HttpMetadata   common.HttpMetadata `json:"http_metadata,omitempty"`
		IdealInfoLinks InfoLinks           `json:"_links,omitempty"`
	}

	InfoLinks struct {
		Self    common.Link  `json:"self,omitempty"`
		Curies  []CuriesLink `json:"curies,omitempty"`
		Issuers common.Link  `json:"ideal:issuers,omitempty"`
	}

	CuriesLink struct {
		Name      string `json:"name,omitempty"`
		Href      string `json:"href,omitempty"`
		Templated bool   `json:"templated,omitempty"`
	}
)

type (
	IssuerResponse struct {
		HttpMetadata common.HttpMetadata    `json:"http_metadata,omitempty"`
		Countries    []IdealCountry         `json:"countries,omitempty"`
		Links        map[string]common.Link `json:"_links,omitempty"`
	}

	IdealCountry struct {
		Name    string   `json:"name,omitempty"`
		Issuers []Issuer `json:"issuers,omitempty"`
	}

	Issuer struct {
		Bic  string `json:"bic,omitempty"`
		Name string `json:"name,omitempty"`
	}
)
