package completion

type CompletionType string

const (
	Hosted    CompletionType = "hosted"
	NonHosted CompletionType = "non_hosted"
)

type (
	Completion interface {
		GetType() CompletionType
	}

	CompletionInfo struct {
		Type CompletionType `json:"type,omitempty"`
	}

	hostedCompletion struct {
		CompletionInfo
		CallbackUrl string `json:"callback_url,omitempty"`
		SuccessUrl  string `json:"success_url,omitempty"`
		FailureUrl  string `json:"failure_url,omitempty"`
	}

	nonHostedCompletion struct {
		CompletionInfo
		CallbackUrl              string `json:"callback_url,omitempty"`
		ChallengeNotificationUrl string `json:"challenge_notification_url,omitempty"`
	}
)

func NewHostedCompletion() *hostedCompletion {
	return &hostedCompletion{CompletionInfo: CompletionInfo{Type: Hosted}}
}

func NewNonHostedCompletion() *nonHostedCompletion {
	return &nonHostedCompletion{CompletionInfo: CompletionInfo{Type: NonHosted}}
}

func (c *hostedCompletion) GetType() CompletionType {
	return c.Type
}

func (c *nonHostedCompletion) GetType() CompletionType {
	return c.Type
}
