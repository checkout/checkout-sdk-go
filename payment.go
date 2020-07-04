package checkout

import (
	"fmt"
	"time"
)

// public const string Authorized = "Authorized";
// public const string Canceled = "Canceled";
// public const string Captured = "Captured";
// public const string Declined = "Declined";
// public const string Expired = "Expired";
// public const string PartiallyCaptured = "Partially Captured";
// public const string PartiallyRefunded = "Partially Refunded";
// public const string Pending = "Pending";
// public const string Refunded = "Refunded";
// public const string Voided = "Voided";
// public const string CardVerified = "Card Verified";
// public const string Chargeback = "Chargeback";

const (
	// Authorized ...
	Authorized string = "Authorized"
	// Canceled ...
	Canceled string = "Canceled"
	// Captured ...
	Captured string = "Captured"
	// Declined ...
	Declined string = "Declined"
	// Expired ...
	Expired string = "Expired"
	// PartiallyCaptured ...
	PartiallyCaptured string = "Partially Captured"
	// PartiallyRefunded ...
	PartiallyRefunded string = "Partially Refunded"
	// Pending ...
	Pending string = "Pending"
	// Refunded ...
	Refunded string = "Refunded"
	// Voided ...
	Voided string = "Voided"
	// CardVerified ...
	CardVerified string = "Card Verified"
	// Chargeback ...
	Chargeback string = "Chargeback"
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

// ThreeDSEnrollment : 3D-Secure Enrollment Data
type ThreeDSEnrollment struct {
	Downgraded             bool   `json:"downgraded,omitempty"`
	Enrolled               string `json:"enrolled,omitempty"`
	SignatureValid         string `json:"signature_valid,omitempty"`
	AuthenticationResponse string `json:"authentication_response,omitempty"`
	Cryptogram             string `json:"cryptogram,omitempty"`
	XID                    string `json:"xid,omitempty"`
}

// ActionSummary ...
type ActionSummary struct {
	ID              string  `json:"id,omitempty"`
	Type            *string `json:"type,omitempty"`
	ResponseCode    *string `json:"response_code,omitempty"`
	ResponseSummary *string `json:"response_summary,omitempty"`
}

// Risk ...
type Risk struct {
	Enabled bool `json:"enabled,omitempty"`
}

// RiskAssessment ...
type RiskAssessment struct {
	Flagged bool `json:"flagged,omitempty"`
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

// Source ...
type Source struct {
	Type           string   `json:"type" binding:"required"`
	ID             string   `json:"id" binding:"required"`
	BillingAddress *Address `json:"billing_address,omitempty"`
	Phone          *Phone   `json:"phone,omitempty"`
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
	Stored         bool     `json:"stored,omitempty"`
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
	Processed   *PaymentProcessed
	Pending     *PaymentPending
}

// PaymentResponse ...
type PaymentResponse struct {
	APIResponse *APIResponse
	Payment     *Payment
}

// Payment ...
type Payment struct {
	ID                string             `json:"id,omitempty"`
	RequestedOn       time.Time          `json:"requested_on,omitempty"`
	Source            *Source            `json:"source,omitempty"`
	Amount            int                `json:"amount,omitempty"`
	Currency          string             `json:"currency,omitempty"`
	PaymentType       string             `json:"payment_type,omitempty"`
	Reference         string             `json:"reference,omitempty"`
	Description       string             `json:"description,omitempty"`
	Approved          bool               `json:"approved,omitempty"`
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

// PaymentPending ...
type PaymentPending struct {
	ID        string             `json:"id,omitempty"`
	Status    string             `json:"status,omitempty"`
	Reference string             `json:"reference,omitempty"`
	Customer  *Customer          `json:"customer,omitempty"`
	ThreeDS   *ThreeDSEnrollment `json:"3ds,omitempty"`
	Links     map[string]Link    `json:"_links,omitempty"`
}

// PaymentProcessed ...
type PaymentProcessed struct {
	ID              string             `json:"id,omitempty"`
	ActionID        string             `json:"action_id,omitempty"`
	Amount          uint64             `json:"amount,omitempty"`
	Currency        string             `json:"currency,omitempty"`
	Approved        bool               `json:"approved,omitempty"`
	Status          string             `json:"status,omitempty"`
	AuthCode        string             `json:"auth_code,omitempty"`
	ResponseCode    string             `json:"response_code,omitempty"`
	ResponseSummary string             `json:"response_summary,omitempty"`
	ThreeDS         *ThreeDSEnrollment `json:"3ds,omitempty"`
	Risk            *RiskAssessment    `json:"risk,omitempty"`
	Source          *Source            `json:"source,omitempty"`
	Customer        *Customer          `json:"customer,omitempty"`
	ProcessedOn     time.Time          `json:"processed_on,omitempty"`
	Reference       string             `json:"reference,omitempty"`
	Processing      *Processing        `json:"processing,omitempty"`
	ECI             string             `json:"eci,omitempty"`
	SchemeID        string             `json:"scheme_id,omitempty"`
	Links           map[string]Link    `json:"_links,omitempty"`
}
