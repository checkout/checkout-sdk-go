package nas

import (
	"encoding/json"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/payments"
)

type (
	SourceResponse struct {
		ResponseCardSource            *ResponseCardSource
		ResponseCurrencyAccountSource *ResponseCurrencyAccountSource
		AlternativeResponse           *common.AlternativeResponse
	}

	ResponseCardSource struct {
		Type           payments.SourceType `json:"type,omitempty"`
		Id             string              `json:"id,omitempty"`
		BillingAddress *common.Address     `json:"billing_address,omitempty"`
		Phone          *common.Phone       `json:"phone,omitempty"`
		ExpiryMonth    int                 `json:"expiry_month,omitempty"`
		ExpiryYear     int                 `json:"expiry_year,omitempty"`
		Name           string              `json:"name,omitempty"`
		Scheme         string              `json:"scheme,omitempty"`
		// Deprecated: This property will be removed in the future, and should not be used. Use LocalSchemes instead.
		SchemeLocal             string              `json:"scheme_local,omitempty"`
		LocalSchemes            []string            `json:"local_schemes,omitempty"`
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
		PaymentAccountReference string              `json:"payment_account_reference,omitempty"`
	}

	ResponseCurrencyAccountSource struct {
		Type   payments.SourceType `json:"type,omitempty"`
		Amount int                 `json:"amount,omitempty"`
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
	case string(payments.CurrencyAccountSource):
		var typeMapping ResponseCurrencyAccountSource
		if err := json.Unmarshal(data, &typeMapping); err != nil {
			return err
		}
		s.ResponseCurrencyAccountSource = &typeMapping
	default:
		var typeMapping common.AlternativeResponse
		if err := json.Unmarshal(data, &typeMapping); err != nil {
			return err
		}
		s.AlternativeResponse = &typeMapping
	}

	return nil
}
