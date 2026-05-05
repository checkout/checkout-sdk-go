package issuing

type CardType string

const (
	Physical CardType = "physical"
	Virtual  CardType = "virtual"
)

type CardStatus string

const (
	CardActive    CardStatus = "active"
	CardInactive  CardStatus = "inactive"
	CardRevoked   CardStatus = "revoked"
	CardSuspended CardStatus = "suspended"
)

type CardScheme string

const (
	MastercardScheme CardScheme = "mastercard"
	VisaScheme       CardScheme = "visa"
)
