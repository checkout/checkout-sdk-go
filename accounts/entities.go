package accounts

import "github.com/checkout/checkout-sdk-go/common"

type CompanyVerificationType string

const (
	IncorporationDocument CompanyVerificationType = "incorporation_document"
	ArticlesOfAssociation CompanyVerificationType = "articles_of_association"
)

type TaxVerificationType string

const (
	EinLetter TaxVerificationType = "ein_letter"
)

type (
	CompanyVerification struct {
		Type  CompanyVerificationType `json:"type,omitempty"`
		Front string                  `json:"front,omitempty"`
	}

	TaxVerification struct {
		Type  TaxVerificationType `json:"type,omitempty"`
		Front string              `json:"front,omitempty"`
	}

	OnboardSubEntityDocuments struct {
		IdentityVerification *Document            `json:"identity_verification,omitempty"`
		CompanyVerification  *CompanyVerification `json:"company_verification,omitempty"`
		TaxVerification      *TaxVerification     `json:"tax_verification,omitempty"`
	}

	OnboardEntityRequest struct {
		Reference      string                     `json:"reference,omitempty"`
		ContactDetails *ContactDetails            `json:"contact_details,omitempty"`
		Profile        *Profile                   `json:"profile,omitempty"`
		Company        *Company                   `json:"company,omitempty"`
		Individual     *Individual                `json:"individual,omitempty"`
		Documents      *OnboardSubEntityDocuments `json:",omitempty"`
	}

	OnboardEntityResponse struct {
		HttpMetadata    common.HttpMetadata `json:"http_metadata,omitempty"`
		Id              string              `json:"id,omitempty"`
		Reference       string              `json:"reference,omitempty"`
		Status          OnboardingStatus    `json:"status,omitempty"`
		Capabilities    *Capabilities       `json:"capabilities,omitempty"`
		RequirementsDue []RequirementsDue   `json:"requirements_due,omitempty"`
	}

	OnboardEntityDetails struct {
		HttpMetadata    common.HttpMetadata `json:"http_metadata,omitempty"`
		Id              string              `json:"id,omitempty"`
		Reference       string              `json:"reference,omitempty"`
		Capabilities    *Capabilities       `json:"capabilities,omitempty"`
		Status          OnboardingStatus    `json:"status,omitempty"`
		RequirementsDue []RequirementsDue   `json:"requirements_due,omitempty"`
		ContactDetails  *ContactDetails     `json:"contact_details,omitempty"`
		Profile         *Profile            `json:"profile,omitempty"`
		Company         *Company            `json:"company,omitempty"`
		Individual      *Individual         `json:"individual,omitempty"`
		Instruments     []Instrument        `json:"instruments,omitempty"`
	}
)
