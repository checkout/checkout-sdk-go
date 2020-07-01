package checkout

import (
	"fmt"
	"time"
)

// PaymentRequest ...
type PaymentRequest struct {
	Source            interface{}        `json:"source"`
	Amount            uint64             `json:"amount,omitempty"`
	Currency          string             `json:"currency"`
	Reference         string             `json:"reference,omitempty"`
	PaymentType       string             `json:"payment_type,omitempty"`
	Description       string             `json:"description,omitempty"`
	Capture           bool               `json:"capture,omitempty"`
	CaptureOn         *time.Time         `json:"capture_on,omitempty"`
	Customer          *Customer          `json:"customer,omitempty"`
	BillingDescriptor *BillingDescriptor `json:"billing_descriptor,omitempty"`
	Shipping          *Shipping          `json:"shipping,omitempty"`
	ThreeDS           *ThreeDS           `json:"3ds,omitempty"`
	PreviousPaymentID string             `json:"previous_payment_id,omitempty"`
	Risk              *Risk              `json:"risk,omitempty"`
	SuccessURL        string             `json:"success_url,omitempty,omitempty"`
	FailureURL        string             `json:"failure_url,omitempty,omitempty"`
	PaymentIP         string             `json:"payment_ip,omitempty"`
	Recipient         *Recipient         `json:"recipient,omitempty"`
	Destinations      []*Destination     `json:"destinations,omitempty"`
	Processing        *Processing        `json:"processing,omitempty"`
	Metadata          map[string]string  `json:"metadata,omitempty"`
}

// Customer ...
type Customer struct {
	ID    string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}

// BillingDescriptor ...
type BillingDescriptor struct {
	Name string `json:"name"`
	City string `json:"city"`
}

// Shipping ...
type Shipping struct {
	Address *Address `json:"address,omitempty"`
	Phione  *Phone   `json:"phone,omitempty"`
}

// ThreeDS ...
type ThreeDS struct {
	Enabled    bool   `json:"enabled,omitempty"`
	AttemptN3d bool   `json:"attempt_n3d,omitempty"`
	ECI        string `json:"eci,omitempty"`
	Cryptogram string `json:"cryptogram,omitempty"`
	XID        string `json:"xid,omitempty"`
	Version    string `json:"version,omitempty"`
}

// Risk ...
type Risk struct {
	Flagged bool `json:"flagged,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
}

// Recipient ...
type Recipient struct {
	DOB           string `json:"dob"`
	AccountNumber string `json:"account_number"`
	ZIP           string `json:"zip"`
	LastName      string `json:"last_name"`
}

// Destination ...
type Destination struct {
	ID     string `json:"id"`
	Amount uint64 `json:"amount"`
}

// Processing - Use the processing object to influence or
// override the data sent during card processing
type Processing struct {
	Mid                      string  `json:"mid,omitempty"`
	Aft                      bool    `json:"aft,omitempty"`
	DLocal                   *DLocal `json:"dlocal,omitempty"`
	AcquirerTransactionID    string  `json:"acquirer_transaction_id,omitempty"`
	RetrievalReferenceNumber string  `json:"retrieval_reference_number,omitempty"`
}

// DLocal - Processing information required for dLocal payments.
type DLocal struct {
	Country     string       `json:"country,omitempty"`
	Payer       *Customer    `json:"payer,omitempty"`
	Installment *Installment `json:"installment,omitempty"`
}

// Installment - Details about the installments.
type Installment struct {
	Count string `json:"count,omitempty"`
}

// SetSource ...
func (r *PaymentRequest) SetSource(s interface{}) error {
	var err error
	switch p := s.(type) {
	case *CardSource:
	case *TokenSource:
	case *CustomerSource:
	case map[string]string:
	default:
		err = fmt.Errorf("Unsupported source type %T", p)
	}
	if err == nil {
		r.Source = s
	}
	return err
}

// CardSource ...
type CardSource struct {
	Type        string `json:"type"`
	Number      string `json:"number"`
	ExpiryMonth uint64 `json:"expiry_month"`
	ExpiryYear  uint64 `json:"expiry_year"`
	CVV         string `json:"cvv"`
}

// TokenSource ...
type TokenSource struct {
	Type           string   `json:"type"`
	Token          string   `json:"token"`
	BillingAddress *Address `json:"billing_address"`
	Phone          *Phone   `json:"phone"`
}

// CustomerSource ...
type CustomerSource struct {
	Type  string   `json:"type"`
	ID    *Address `json:"id,omitempty"`
	Email string   `json:"email,omitempty"`
}

// Address ...
type Address struct {
	AddressLine1 string `json:"address_line1,omitempty"`
	AddressLine2 string `json:"address_line2,omitempty"`
	City         string `json:"city,omitempty"`
	State        string `json:"state,omitempty"`
	ZIP          string `json:"zip,omitempty"`
	Country      string `json:"country,omitempty"`
}

// Phone ...
type Phone struct {
	CountryCode string `json:"country_code,omitempty"`
	Number      string `json:"number,omitempty"`
}

// Response ...
type Response struct {
	APIResponse *APIResponse
	Pending     *Pending
	Authorized  *Authorized
}

// PaymentResponse ...
type PaymentResponse struct {
	APIResponse *APIResponse
	Payment     *Payment
}

// Payment ...
type Payment struct {
}

// Authorized ...
type Authorized struct {
}

// Pending ...
type Pending struct {
}

// Processed ...
type Processed struct {
}
