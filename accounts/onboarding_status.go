package accounts

type OnboardingStatus string

const (
	Active         OnboardingStatus = "active"
	Pending        OnboardingStatus = "pending"
	Restricted     OnboardingStatus = "restricted"
	RequirementDue OnboardingStatus = "requirements_due"
	Inactive       OnboardingStatus = "inactive"
	Rejected       OnboardingStatus = "rejected"
)
