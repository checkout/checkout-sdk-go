package common

// AccountHolder -
type AccountHolder struct {
	BillingAddress *Address `json:"billing_address,omitempty"`
	Phone          *Phone   `json:"phone,omitempty"`
}
