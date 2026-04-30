package nas

import (
	"encoding/json"
	"strconv"

	"github.com/checkout/checkout-sdk-go/v2/common"
	"github.com/checkout/checkout-sdk-go/v2/payments"
)

type (
	SourceResponse struct {
		ResponseCardSource                  *ResponseCardSource
		ResponseCurrencyAccountSource       *ResponseCurrencyAccountSource
		PaymentContextsPayPayResponseSource *PaymentContextsPayPayResponseSource
		AlternativeResponse                 *common.AlternativeResponse
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
		SchemeLocal             string                           `json:"scheme_local,omitempty"`
		LocalSchemes            []string                         `json:"local_schemes,omitempty"`
		Last4                   string                           `json:"last4,omitempty"`
		Fingerprint             string                           `json:"fingerprint,omitempty"`
		Bin                     string                           `json:"bin,omitempty"`
		CardType                common.CardType                  `json:"card_type,omitempty"`
		CardCategory            common.CardCategory              `json:"card_category,omitempty"`
		CardWalletType          common.CardWalletType            `json:"card_wallet_type,omitempty"`
		Issuer                  string                           `json:"issuer,omitempty"`
		IssuerCountry           common.Country                   `json:"issuer_country,omitempty"`
		ProductId               string                           `json:"product_id,omitempty"`
		ProductType             string                           `json:"product_type,omitempty"`
		AvsCheck                string                           `json:"avs_check,omitempty"`
		CvvCheck                string                           `json:"cvv_check,omitempty"`
		PaymentAccountReference string                           `json:"payment_account_reference,omitempty"`
		EncryptedCardNumber     string                           `json:"encrypted_card_number,omitempty"`
		AccountUpdateStatus     payments.AccountUpdateStatusType `json:"account_update_status,omitempty"`
		AccountHolder           *common.AccountHolderResponse    `json:"account_holder,omitempty"`
	}

	ResponseCurrencyAccountSource struct {
		Type   payments.SourceType `json:"type,omitempty"`
		Amount int                 `json:"amount,omitempty"`
	}

	PaymentContextsPayPayResponseSource struct {
		Type payments.SourceType `json:"type,omitempty"`
	}
)

// UnmarshalJSON handles expiry_month/expiry_year being returned as either a
// JSON number (regular payment endpoints) or a JSON string (search endpoint).
func (r *ResponseCardSource) UnmarshalJSON(data []byte) error {
	type Alias ResponseCardSource
	aux := struct {
		ExpiryMonth interface{} `json:"expiry_month"`
		ExpiryYear  interface{} `json:"expiry_year"`
		*Alias
	}{Alias: (*Alias)(r)}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	r.ExpiryMonth = toIntField(aux.ExpiryMonth)
	r.ExpiryYear = toIntField(aux.ExpiryYear)
	return nil
}

func toIntField(v interface{}) int {
	switch val := v.(type) {
	case float64:
		return int(val)
	case string:
		if i, err := strconv.Atoi(val); err == nil {
			return i
		}
	}
	return 0
}

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
	case string(payments.PayPalSource):
		var typeMapping PaymentContextsPayPayResponseSource
		if err := json.Unmarshal(data, &typeMapping); err != nil {
			return err
		}
		s.PaymentContextsPayPayResponseSource = &typeMapping
	default:
		var typeMapping common.AlternativeResponse
		if err := json.Unmarshal(data, &typeMapping); err != nil {
			return err
		}
		s.AlternativeResponse = &typeMapping
	}

	return nil
}
