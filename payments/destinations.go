package payments

type DestinationType string

const (
	BankAccountDestination  DestinationType = "bank_account"
	CardDestination         DestinationType = "card"
	IdDestination           DestinationType = "id"
	TokenDestination        DestinationType = "token"
	NetworkTokenDestination DestinationType = "network_token"
)

type Destination interface {
	GetType() DestinationType
}
