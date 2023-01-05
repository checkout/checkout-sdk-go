package configuration

const (
	PreviousSecretKeyPattern string = "^sk_(test_)?(\\w{8})-(\\w{4})-(\\w{4})-(\\w{4})-(\\w{12})$"
	PreviousPublicKeyPattern string = "^pk_(test_)?(\\w{8})-(\\w{4})-(\\w{4})-(\\w{4})-(\\w{12})$"
	DefaultSecretKeyPattern  string = "^sk_(sbox_)?[a-z2-7]{26}[a-z2-7*#$=]$"
	DefaultPublicKeyPattern  string = "^pk_(sbox_)?[a-z2-7]{26}[a-z2-7*#$=]$"
)
