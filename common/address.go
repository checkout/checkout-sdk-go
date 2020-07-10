package common

// Address - Defines a postal address.
type Address struct {
	AddressLine1 string `json:"address_line1,omitempty"`
	AddressLine2 string `json:"address_line2,omitempty"`
	City         string `json:"city,omitempty"`
	State        string `json:"state,omitempty"`
	ZIP          string `json:"zip,omitempty"`
	Country      string `json:"country,omitempty"`
}
