package abc

import (
	"encoding/json"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/payments"
)

type (
	SourceResponse struct {
		ResponseCardSource  *ResponseCardSource
		AlternativeResponse *common.AlternativeResponse
	}

	ResponseCardSource struct {
		Type                    payments.SourceType `json:"type,omitempty"`
		Id                      string              `json:"id,omitempty"`
		BillingAddress          *common.Address     `json:"billing_address,omitempty"`
		Phone                   *common.Phone       `json:"phone,omitempty"`
		ExpiryMonth             int                 `json:"expiry_month,omitempty"`
		ExpiryYear              int                 `json:"expiry_year,omitempty"`
		Name                    string              `json:"name,omitempty"`
		Scheme                  string              `json:"scheme,omitempty"`
		SchemeLocal             string              `json:"scheme_local,omitempty"`
		Last4                   string              `json:"last4,omitempty"`
		Fingerprint             string              `json:"fingerprint,omitempty"`
		Bin                     string              `json:"bin,omitempty"`
		CardType                common.CardType     `json:"card_type,omitempty"`
		CardCategory            common.CardCategory `json:"card_category,omitempty"`
		Issuer                  string              `json:"issuer,omitempty"`
		IssuerCountry           common.Country      `json:"issuer_country,omitempty"`
		ProductId               string              `json:"product_id,omitempty"`
		ProductType             string              `json:"product_type,omitempty"`
		AvsCheck                string              `json:"avs_check,omitempty"`
		CvvCheck                string              `json:"cvv_check,omitempty"`
		Payouts                 bool                `json:"payouts,omitempty"`
		FastFunds               string              `json:"fast_funds,omitempty"`
		PaymentAccountReference string              `json:"payment_account_reference,omitempty"`
	}
)

func (s *SourceResponse) UnmarshalJSON(data []byte) error {
	var typeMapping common.TypeMapping
	if err := json.Unmarshal(data, &typeMapping); err != nil {
		return err
	}

	switch typeMapping.Type {
	case string(payments.CardSource):
		var typeMapping ResponseCardSource
		if err := json.Unmarshal(data, &typeMapping); err != nil {
			return err
		}
		s.ResponseCardSource = &typeMapping
	default:
		var typeMapping common.AlternativeResponse
		if err := json.Unmarshal(data, &typeMapping); err != nil {
			return err
		}
		s.AlternativeResponse = &typeMapping
	}

	return nil
}
