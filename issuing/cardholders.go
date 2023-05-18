package issuing

import (
	"time"

	"github.com/checkout/checkout-sdk-go/common"
)

const (
	issuing     = "issuing"
	cardholders = "cardholders"
	cards       = "cards"
)

type CardholderType string

const (
	Individual CardholderType = "individual"
)

type CardholderStatus string

const (
	CardholderActive          CardholderStatus = "active"
	CardholderPending         CardholderStatus = "pending"
	CardholderRestricted      CardholderStatus = "restricted"
	CardholderRequirementsDue CardholderStatus = "requirements_due"
	CardholderInactive        CardholderStatus = "inactive"
	CardholderRejected        CardholderStatus = "rejected"
)

type (
	CardholderDocument struct {
		Type            common.DocumentType `json:"type,omitempty"`
		FrontDocumentId string              `json:"front_document_id,omitempty"`
		BackDocumentId  string              `json:"back_document_id,omitempty"`
	}

	CardholderRequest struct {
		Type             CardholderType      `json:"type,omitempty"`
		Reference        string              `json:"reference,omitempty"`
		EntityId         string              `json:"entity_id,omitempty"`
		FirstName        string              `json:"first_name,omitempty"`
		MiddleName       string              `json:"middle_name,omitempty"`
		LastName         string              `json:"last_name,omitempty"`
		Email            string              `json:"email,omitempty"`
		PhoneNumber      *common.Phone       `json:"phone_number,omitempty"`
		DateOfBirth      string              `json:"date_of_birth,omitempty"`
		BillingAddress   *common.Address     `json:"billing_address,omitempty"`
		ResidencyAddress *common.Address     `json:"residency_address,omitempty"`
		Document         *CardholderDocument `json:"document,omitempty"`
	}
)

type (
	CardholderResponse struct {
		HttpMetadata     common.HttpMetadata
		Id               string                 `json:"id,omitempty"`
		Type             CardholderType         `json:"type,omitempty"`
		Status           CardholderStatus       `json:"status,omitempty"`
		Reference        string                 `json:"reference,omitempty"`
		CreatedDate      *time.Time             `json:"created_date,omitempty"`
		LastModifiedDate *time.Time             `json:"last_modified_date,omitempty"`
		Links            map[string]common.Link `json:"links,omitempty"`
	}

	CardholderDetailsResponse struct {
		HttpMetadata      common.HttpMetadata
		Id                string                 `json:"id,omitempty"`
		Type              CardholderType         `json:"type,omitempty"`
		FirstName         string                 `json:"first_name,omitempty"`
		MiddleName        string                 `json:"middle_name,omitempty"`
		LastName          string                 `json:"last_name,omitempty"`
		Email             string                 `json:"email,omitempty"`
		PhoneNumber       *common.Phone          `json:"phone_number,omitempty"`
		DateOfBirth       string                 `json:"date_of_birth,omitempty"`
		BillingAddress    *common.Address        `json:"billing_address,omitempty"`
		ResidencyAddress  *common.Address        `json:"residency_address,omitempty"`
		Reference         string                 `json:"reference,omitempty"`
		AccountEntityId   string                 `json:"account_entity_id,omitempty"`
		ParentSubEntityId string                 `json:"parent_sub_entity_id,omitempty"`
		EntityId          string                 `json:"entity_id,omitempty"`
		CreatedDate       *time.Time             `json:"created_date,omitempty"`
		LastModifiedDate  *time.Time             `json:"last_modified_date,omitempty"`
		Links             map[string]common.Link `json:"links,omitempty"`
	}

	CardholderCardsResponse struct {
		HttpMetadata common.HttpMetadata
		Cards        []CardDetailsResponse `json:"cards,omitempty"`
	}
)
