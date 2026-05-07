package standaloneaccountupdater

import "github.com/checkout/checkout-sdk-go/v2/common"

const accountUpdaterPath = "account-updater/cards"

type AccountUpdateStatus string

const (
	CardUpdated       AccountUpdateStatus = "CARD_UPDATED"
	CardExpiryUpdated AccountUpdateStatus = "CARD_EXPIRY_UPDATED"
	CardClosed        AccountUpdateStatus = "CARD_CLOSED"
	UpdateFailed      AccountUpdateStatus = "UPDATE_FAILED"
)

type AccountUpdateFailureCode string

const (
	CardholderOptOut    AccountUpdateFailureCode = "CARDHOLDER_OPT_OUT"
	UpToDate            AccountUpdateFailureCode = "UP_TO_DATE"
	NonParticipatingBin AccountUpdateFailureCode = "NON_PARTICIPATING_BIN"
	Unknown             AccountUpdateFailureCode = "UNKNOWN"
)

type AccountUpdaterCard struct {
	Number      string `json:"number"`
	ExpiryMonth int    `json:"expiry_month"`
	ExpiryYear  int    `json:"expiry_year"`
}

type AccountUpdaterInstrument struct {
	Id string `json:"id"`
}

type AccountUpdaterSourceOptions struct {
	Card       *AccountUpdaterCard       `json:"card,omitempty"`
	Instrument *AccountUpdaterInstrument `json:"instrument,omitempty"`
}

type GetUpdatedCardCredentialsRequest struct {
	SourceOptions AccountUpdaterSourceOptions `json:"source_options"`
}

type AccountUpdaterCardDetails struct {
	EncryptedCardNumber string `json:"encrypted_card_number,omitempty"`
	Bin                 string `json:"bin,omitempty"`
	Last4               string `json:"last4,omitempty"`
	ExpiryMonth         int    `json:"expiry_month,omitempty"`
	ExpiryYear          int    `json:"expiry_year,omitempty"`
	Fingerprint         string `json:"fingerprint,omitempty"`
}

type GetUpdatedCardCredentialsResponse struct {
	HttpMetadata             common.HttpMetadata
	AccountUpdateStatus      AccountUpdateStatus      `json:"account_update_status,omitempty"`
	AccountUpdateFailureCode AccountUpdateFailureCode `json:"account_update_failure_code,omitempty"`
	Card                     *AccountUpdaterCardDetails `json:"card,omitempty"`
}
