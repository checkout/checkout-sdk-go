module github.com/HelloRelai/checkout-sdk-go

go 1.18

require (
	github.com/checkout/checkout-sdk-go v0.0.22
	github.com/google/go-querystring v1.0.0
	github.com/google/uuid v1.3.0
	github.com/stretchr/testify v1.7.0
)

require (
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)

replace github.com/checkout/checkout-sdk-go => github.com/HelloRelai/checkout-sdk-go v0.0.28
