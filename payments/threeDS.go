package payments

type Exemption string

const (
	// LowValue ...
	LowValue Exemption = "low_value"
	// SecureCorporatePayment ...
	SecureCorporatePayment Exemption = "secure_corporate_payment"
	// TrustedListing ...
	TrustedListing Exemption = "trusted_listing"
	// TransactionRiskAssessment ...
	TransactionRiskAssessment Exemption = "transaction_risk_assessment"
	// ThreeDSOutage ...
	ThreeDSOutage Exemption = "3ds_outage"
	// SCADelegation ...
	SCADelegation Exemption = "sca_delegation"
	// OutOfSCAScope ...
	OutOfSCAScope Exemption = "out_of_sca_scope"
	// Other ...
	Other Exemption = "other"
	// LowRiskProgram ...
	LowRiskProgram Exemption = "low_risk_program"
)

type ChallengeIndicator string

const (
	// NoPreference ...
	NoPreference ChallengeIndicator = "no_preference"
	// NoChallengeRequested ...
	NoChallengeRequested ChallengeIndicator = "no_challenge_requested"
	// ChallengeRequested ...
	ChallengeRequested ChallengeIndicator = "challenge_requested"
	// ChallengeRequestedMandate ...
	ChallengeRequestedMandate ChallengeIndicator = "challenge_requested_mandate"
)

func (c ChallengeIndicator) String() string {
	return string(c)
}

// ThreeDS ...
type ThreeDS struct {
	Enabled            *bool              `json:"enabled,omitempty"`
	AttemptN3d         *bool              `json:"attempt_n3d,omitempty"`
	ECI                string             `json:"eci,omitempty"`
	Cryptogram         string             `json:"cryptogram,omitempty"`
	XID                string             `json:"xid,omitempty"`
	Version            string             `json:"version,omitempty"`
	Exemption          Exemption          `json:"exemption,omitempty"`
	ChallengeIndicator ChallengeIndicator `json:"challenge_indicator,omitempty"`
}

// ThreeDSEnrollment : 3D-Secure Enrollment Data
type ThreeDSEnrollment struct {
	Downgraded             *bool  `json:"downgraded,omitempty"`
	Enrolled               string `json:"enrolled,omitempty"`
	SignatureValid         string `json:"signature_valid,omitempty"`
	AuthenticationResponse string `json:"authentication_response,omitempty"`
	Cryptogram             string `json:"cryptogram,omitempty"`
	XID                    string `json:"xid,omitempty"`
}
