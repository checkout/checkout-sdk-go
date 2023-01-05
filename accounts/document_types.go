package accounts

type DocumentType string

const (
	Passport             DocumentType = "passport"
	NationalIdentityCard DocumentType = "national_identity_card"
	DrivingLicense       DocumentType = "driving_license"
	CitizenCard          DocumentType = "citizen_card"
	ResidencePermit      DocumentType = "residence_permit"
	ElectoralId          DocumentType = "electoral_id"
)
