package googlepay

import (
	"time"

	"github.com/checkout/checkout-sdk-go/v3/common"
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

// CreateEnrollmentResponse is the body returned by POST /googlepay/enrollments.
//
// The API returns merchant_id, tos_accepted_time and state. The spec declares only
// tosAcceptedTime (camelCase) and state, with additionalProperties false, so this struct was
// generated from a wrong schema: merchant_id was missing, and TosAcceptedTime carried the
// camelCase name, which never matches the real payload and so was always nil. Reported by a
// merchant. The spec is being fixed separately; until then this struct follows the live API.
type CreateEnrollmentResponse struct {
	HttpMetadata common.HttpMetadata
	// MerchantId is the Google Pay merchant identifier assigned to the entity, needed to
	// initialise Google Pay on the client.
	MerchantId      string          `json:"merchant_id"`
	TosAcceptedTime *time.Time      `json:"tos_accepted_time"`
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
