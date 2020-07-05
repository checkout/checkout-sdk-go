package payments

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/shiuh-yaw-cko/checkout"
)

// Request ...
type Request struct {
	Source            interface{}        `json:"source"`
	Amount            uint64             `json:"amount,omitempty"`
	Currency          string             `json:"currency"`
	Reference         string             `json:"reference,omitempty"`
	PaymentType       string             `json:"payment_type,omitempty"`
	Description       string             `json:"description,omitempty"`
	Capture           *bool              `json:"capture,omitempty"`
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

// SetSource ...
func (r *Request) SetSource(s interface{}) error {
	var err error
	switch p := s.(type) {
	case *IDSource:
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

// IDSource ...
type IDSource struct {
	Type string `json:"type" binding:"required"`
	ID   string `json:"id" binding:"required"`
	CVV  string `json:"cvv,omitempty"`
}

// CardSource ...
type CardSource struct {
	Type           string   `json:"type" binding:"required"`
	Number         string   `json:"number" binding:"required"`
	ExpiryMonth    int      `json:"expiry_month" binding:"required"`
	ExpiryYear     int      `json:"expiry_year" binding:"required"`
	Name           string   `json:"name,omitempty"`
	CVV            string   `json:"cvv,omitempty"`
	Stored         *bool    `json:"stored,omitempty"`
	BillingAddress *Address `json:"billing_address,omitempty"`
	Phone          *Phone   `json:"phone,omitempty"`
}

// TokenSource ...
type TokenSource struct {
	Type           string   `json:"type" binding:"required"`
	Token          string   `json:"token" binding:"required"`
	BillingAddress *Address `json:"billing_address,omitempty"`
	Phone          *Phone   `json:"phone,omitempty"`
}

// CustomerSource ...
type CustomerSource struct {
	Type  string   `json:"type" binding:"required"`
	ID    *Address `json:"id,omitempty"`
	Email string   `json:"email,omitempty"`
}

// Customer ...
type Customer struct {
	ID    string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}

// BillingDescriptor ...
type BillingDescriptor struct {
	Name string `json:"name,omitempty"`
	City string `json:"city,omitempty"`
}

// Shipping ...
type Shipping struct {
	Address *Address `json:"address,omitempty"`
	Phione  *Phone   `json:"phone,omitempty"`
}

// Risk ...
type Risk struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// RiskAssessment ...
type RiskAssessment struct {
	Flagged *bool `json:"flagged,omitempty"`
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
	Aft                      *bool   `json:"aft,omitempty"`
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

// Response ...
type Response struct {
	StatusResponse *checkout.StatusResponse `json:"api_response,omitempty"`
	Processed      *Processed               `json:"processed,omitempty"`
	Pending        *PaymentPending          `json:"pending,omitempty"`
}

// Processed ...
type Processed struct {
	ID                string                   `json:"id,omitempty"`
	ActionID          string                   `json:"action_id,omitempty"`
	Amount            uint64                   `json:"amount,omitempty"`
	Currency          string                   `json:"currency,omitempty"`
	Approved          *bool                    `json:"approved,omitempty"`
	Status            string                   `json:"status,omitempty"`
	AuthCode          string                   `json:"auth_code,omitempty"`
	ResponseCode      string                   `json:"response_code,omitempty"`
	ResponseSummary   string                   `json:"response_summary,omitempty"`
	ThreeDSEnrollment *ThreeDSEnrollment       `json:"3ds,omitempty"`
	RiskAssessment    *RiskAssessment          `json:"risk,omitempty"`
	Source            *SourceResponse          `json:"source"`
	Customer          *Customer                `json:"customer,omitempty"`
	ProcessedOn       time.Time                `json:"processed_on,omitempty"`
	Reference         string                   `json:"reference,omitempty"`
	Processing        *Processing              `json:"processing,omitempty"`
	ECI               string                   `json:"eci,omitempty"`
	SchemeID          string                   `json:"scheme_id,omitempty"`
	Links             map[string]checkout.Link `json:"_links,omitempty"`
}

// PaymentPending ...
type PaymentPending struct {
	ID        string                   `json:"id,omitempty"`
	Status    string                   `json:"status,omitempty"`
	Reference string                   `json:"reference,omitempty"`
	Customer  *Customer                `json:"customer,omitempty"`
	ThreeDS   *ThreeDSEnrollment       `json:"3ds,omitempty"`
	Links     map[string]checkout.Link `json:"_links,omitempty"`
}

// SourceResponse ...
type SourceResponse struct {
	*CardSourceResponse
	*AlternativePaymentSourceResponse
}

// MarshalJSON ...
func (s *SourceResponse) MarshalJSON() ([]byte, error) {
	if s.CardSourceResponse != nil {
		return json.Marshal(s.CardSourceResponse)
	} else if s.AlternativePaymentSourceResponse != nil {
		return json.Marshal(s.AlternativePaymentSourceResponse)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON ...
func (s *SourceResponse) UnmarshalJSON(data []byte) error {
	temp := &struct {
		Type string `json:"type"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	switch temp.Type {
	case "card":
		var source CardSourceResponse
		if err := json.Unmarshal(data, &source); err != nil {
			return err
		}
		s.CardSourceResponse = &source
		s.AlternativePaymentSourceResponse = nil
	default:
		var source AlternativePaymentSourceResponse
		if err := json.Unmarshal(data, &source); err != nil {
			return err
		}
		s.AlternativePaymentSourceResponse = &source
		s.CardSourceResponse = nil
	}
	return nil
}

// CardSourceResponse ...
type CardSourceResponse struct {
	ID                      string   `json:"id,omitempty"`
	Type                    string   `json:"type,omitempty"`
	BillingAddress          *Address `json:"billing_address,omitempty"`
	Phone                   *Phone   `json:"phone,omitempty"`
	ExpiryMonth             int      `json:"expiry_month,omitempty"`
	ExpiryYear              int      `json:"expiry_year,omitempty"`
	Name                    string   `json:"name,omitempty"`
	Scheme                  string   `json:"scheme,omitempty"`
	Last4                   string   `json:"last4,omitempty"`
	Fingerprint             string   `json:"fingerprint,omitempty"`
	Bin                     string   `json:"bin,omitempty"`
	CardType                string   `json:"card_type,omitempty"`
	CardCategory            string   `json:"card_category,omitempty"`
	Issuer                  string   `json:"issuer,omitempty"`
	IssuerCountry           string   `json:"issuer_country,omitempty"`
	ProductID               string   `json:"product_id,omitempty"`
	ProductType             string   `json:"product_type,omitempty"`
	AVSCheck                string   `json:"avs_check,omitempty"`
	CVVCheck                string   `json:"cvv_check,omitempty"`
	PaymentAccountReference string   `json:"payment_account_reference,omitempty"`
	Payouts                 *bool    `json:"payouts,omitempty"`
	FastFunds               string   `json:"fast_funds,omitempty"`
}

// AlternativePaymentSourceResponse ...
type AlternativePaymentSourceResponse struct {
	ID             string   `json:"id"`
	Type           string   `json:"type"`
	BillingAddress *Address `json:"billing_address,omitempty"`
	Phone          *Phone   `json:"phone,omitempty"`
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

// PaymentResponse ...
type PaymentResponse struct {
	StatusResponse *checkout.StatusResponse `json:"api_response,omitempty"`
	Payment        *Payment                 `json:"payment,omitempty"`
}

// Payment ...
type Payment struct {
	ID                string             `json:"id,omitempty"`
	RequestedOn       time.Time          `json:"requested_on,omitempty"`
	Source            *SourceResponse    `json:"source,omitempty"`
	Amount            int                `json:"amount,omitempty"`
	Currency          string             `json:"currency,omitempty"`
	PaymentType       string             `json:"payment_type,omitempty"`
	Reference         string             `json:"reference,omitempty"`
	Description       string             `json:"description,omitempty"`
	Approved          *bool              `json:"approved,omitempty"`
	Status            string             `json:"status,omitempty"`
	ThreeDS           *ThreeDSEnrollment `json:"3ds,omitempty"`
	Risk              *RiskAssessment    `json:"risk,omitempty"`
	Customer          *Customer          `json:"customer,omitempty"`
	BillingDescriptor *BillingDescriptor `json:"billing_descriptor,omitempty"`
	Shipping          *Shipping          `json:"shipping,omitempty"`
	PaymentIP         string             `json:"payment_ip,omitempty"`
	Recipient         *Recipient         `json:"recipient,omitempty"`
	Metadata          map[string]string  `json:"metadata,omitempty"`
	ECI               string             `json:"eci,omitempty"`
	Actions           []ActionSummary    `json:"actions,omitempty"`
	SchemeID          string             `json:"scheme_id,omitempty"`
}

// Accepted ...
type Accepted struct {
	ActionID  string                   `json:"action_id"`
	Reference string                   `json:"reference"`
	Links     map[string]checkout.Link `json:"_links"`
}
