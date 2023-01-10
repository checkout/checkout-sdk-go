package payments

type DestinationType string

const (
	BankAccountDestination DestinationType = "bank_account"
	CardDestination        DestinationType = "card"
	IdDestination          DestinationType = "id"
	TokenDestination       DestinationType = "token"
)

type Destination interface {
	GetType() DestinationType
}
