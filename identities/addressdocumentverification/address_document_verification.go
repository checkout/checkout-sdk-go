package addressdocumentverification

import (
	"time"

	"github.com/checkout/checkout-sdk-go/v2/common"
	"github.com/checkout/checkout-sdk-go/v2/identities"
)

const (
	addressDocumentVerificationsPath = "address-document-verifications"
	anonymizePath                    = "anonymize"
	attemptsPath                     = "attempts"
	reportPath                       = "pdf-report"
)

type CreateAddressDocumentVerificationRequest struct {
	ApplicantId   string                   `json:"applicant_id"`
	UserJourneyId string                   `json:"user_journey_id"`
	DeclaredData  *identities.DeclaredData `json:"declared_data,omitempty"`
}

type CreateAddressDocumentVerificationAttemptRequest struct {
	Document string `json:"document"`
}

// Address is the address extracted from the document.
type Address struct {
	AddressLine1 string `json:"address_line1,omitempty"`
	AddressLine2 string `json:"address_line2,omitempty"`
	City         string `json:"city,omitempty"`
	State        string `json:"state,omitempty"`
	Zip          string `json:"zip,omitempty"`
	Country      string `json:"country,omitempty"`
}

// AddressDocumentResult is the result of the address document check.
type AddressDocumentResult struct {
	DocumentType string   `json:"document_type,omitempty"`
	Issuer       string   `json:"issuer,omitempty"`
	FullNames    []string `json:"full_names,omitempty"`
	IssueDate    string   `json:"issue_date,omitempty"`
	Address      *Address `json:"address,omitempty"`
}

// Links holds the HAL links related to the resource (self and, for verifications, applicant).
type Links struct {
	Self      *common.Link `json:"self,omitempty"`
	Applicant *common.Link `json:"applicant,omitempty"`
}

// addressDocumentVerificationBase holds fields common to all response types.
type addressDocumentVerificationBase struct {
	HttpMetadata  common.HttpMetadata
	Id            string                    `json:"id,omitempty"`
	CreatedOn     *time.Time                `json:"created_on,omitempty"`
	ModifiedOn    *time.Time                `json:"modified_on,omitempty"`
	ResponseCodes []identities.ResponseCode `json:"response_codes,omitempty"`
	Links         *Links                    `json:"_links,omitempty"`
}

type AddressDocumentVerificationResponse struct {
	addressDocumentVerificationBase
	UserJourneyId   string                                       `json:"user_journey_id,omitempty"`
	ApplicantId     string                                       `json:"applicant_id,omitempty"`
	Status          identities.AddressDocumentVerificationStatus `json:"status,omitempty"`
	AddressDocument *AddressDocumentResult                       `json:"address_document,omitempty"`
}

type AddressDocumentVerificationAttemptResponse struct {
	addressDocumentVerificationBase
	Status identities.AddressDocumentVerificationAttemptStatus `json:"status,omitempty"`
}

type AddressDocumentVerificationAttemptsResponse struct {
	HttpMetadata common.HttpMetadata
	TotalCount   int                                          `json:"total_count,omitempty"`
	Skip         int                                          `json:"skip,omitempty"`
	Limit        int                                          `json:"limit,omitempty"`
	Data         []AddressDocumentVerificationAttemptResponse `json:"data,omitempty"`
	Links        *Links                                       `json:"_links,omitempty"`
}

type AddressDocumentVerificationReportResponse struct {
	HttpMetadata common.HttpMetadata
	SignedUrl    string `json:"signed_url,omitempty"`
}
