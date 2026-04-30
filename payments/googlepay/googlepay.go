package googlepay

import (
	"time"

	"github.com/checkout/checkout-sdk-go/v2/common"
)

const (
	GooglePayEnrollmentsPath = "googlepay/enrollments"
	domainPath               = "domain"
	domainsPath              = "domains"
	statePath                = "state"
)

type EnrollmentState string

const (
	Active   EnrollmentState = "ACTIVE"
	Inactive EnrollmentState = "INACTIVE"
)

type CreateEnrollmentRequest struct {
	EntityId             string `json:"entity_id"`
	EmailAddress         string `json:"email_address"`
	AcceptTermsOfService bool   `json:"accept_terms_of_service"`
}

type RegisterDomainRequest struct {
	WebDomain string `json:"web_domain"`
}

type CreateEnrollmentResponse struct {
	HttpMetadata    common.HttpMetadata
	TosAcceptedTime *time.Time      `json:"tosAcceptedTime"`
	State           EnrollmentState `json:"state"`
}

type DomainListResponse struct {
	HttpMetadata common.HttpMetadata
	Domains      []string `json:"domains"`
}

type EnrollmentStateResponse struct {
	HttpMetadata common.HttpMetadata
	State        EnrollmentState `json:"state"`
}
