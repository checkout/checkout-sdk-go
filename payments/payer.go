package payments

type (
	Payer struct {
		Name     string `json:"name,omitempty"`
		Email    string `json:"email,omitempty"`
		Document string `json:"document,omitempty"`
	}
)
