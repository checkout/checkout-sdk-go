package accounts

const (
	accountsPath           = "accounts"
	entitiesPath           = "entities"
	membersPath            = "members"
	filesPath              = "files"
	instrumentsPath        = "instruments"
	payoutSchedulesPath    = "payout-schedules"
	paymentInstrumentsPath = "payment-instruments"
)

type AccountHolderType string

const (
	IndividualType AccountHolderType = "individual"
	Corporate      AccountHolderType = "corporate"
	Government     AccountHolderType = "government"
)

type InstrumentStatus string

const (
	Verified          InstrumentStatus = "verified"
	Unverified        InstrumentStatus = "unverified"
	InstrumentPending InstrumentStatus = "pending"
)

type (
	InstrumentDocument struct {
		Type   string `json:"type,omitempty"`
		FileId string `json:"file_id,omitempty"`
	}
)
