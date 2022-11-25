package checkout

import (
	"github.com/checkout/checkout-sdk-go-beta/abc"
	"github.com/checkout/checkout-sdk-go-beta/nas"
)

type CheckoutSdkBuilder struct{}

func (b *CheckoutSdkBuilder) Previous() *abc.CheckoutPreviousSdkBuilder {
	return &abc.CheckoutPreviousSdkBuilder{}
}

func (b *CheckoutSdkBuilder) StaticKeys() *nas.CheckoutDefaultSdkBuilder {
	return &nas.CheckoutDefaultSdkBuilder{}
}

func (b *CheckoutSdkBuilder) OAuth() *nas.CheckoutOAuthSdkBuilder {
	return &nas.CheckoutOAuthSdkBuilder{}
}
