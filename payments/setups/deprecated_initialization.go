package setups

import "encoding/json"

// The payment method types below embed PaymentMethodBase, whose Initialization is valid for some
// methods and not for these. Go marks a deprecation per field, not per embedding type, so each
// type declares its own deprecated Initialization that shadows the promoted one. Their JSON
// methods keep the wire as it was: a value set through PaymentMethodBase, as in
// T{PaymentMethodBase: PaymentMethodBase{Initialization: ...}}, is still sent, and a received
// value is readable through both fields.

// MarshalJSON sends Initialization, or PaymentMethodBase.Initialization when Initialization is
// empty.
func (m BizumPaymentMethod) MarshalJSON() ([]byte, error) {
	// The alias sheds the method set, so json.Marshal below does not re-enter this method.
	type alias BizumPaymentMethod
	if m.Initialization == "" {
		m.Initialization = m.PaymentMethodBase.Initialization
	}
	return json.Marshal(alias(m))
}

// UnmarshalJSON sets a received initialization on both Initialization and
// PaymentMethodBase.Initialization.
func (m *BizumPaymentMethod) UnmarshalJSON(data []byte) error {
	type alias BizumPaymentMethod
	if err := json.Unmarshal(data, (*alias)(m)); err != nil {
		return err
	}
	m.PaymentMethodBase.Initialization = m.Initialization
	return nil
}

// MarshalJSON sends Initialization, or PaymentMethodBase.Initialization when Initialization is
// empty.
func (m InstrumentPaymentMethod) MarshalJSON() ([]byte, error) {
	// The alias sheds the method set, so json.Marshal below does not re-enter this method.
	type alias InstrumentPaymentMethod
	if m.Initialization == "" {
		m.Initialization = m.PaymentMethodBase.Initialization
	}
	return json.Marshal(alias(m))
}

// UnmarshalJSON sets a received initialization on both Initialization and
// PaymentMethodBase.Initialization.
func (m *InstrumentPaymentMethod) UnmarshalJSON(data []byte) error {
	type alias InstrumentPaymentMethod
	if err := json.Unmarshal(data, (*alias)(m)); err != nil {
		return err
	}
	m.PaymentMethodBase.Initialization = m.Initialization
	return nil
}

// MarshalJSON sends Initialization, or PaymentMethodBase.Initialization when Initialization is
// empty.
func (m MultibancoPaymentMethod) MarshalJSON() ([]byte, error) {
	// The alias sheds the method set, so json.Marshal below does not re-enter this method.
	type alias MultibancoPaymentMethod
	if m.Initialization == "" {
		m.Initialization = m.PaymentMethodBase.Initialization
	}
	return json.Marshal(alias(m))
}

// UnmarshalJSON sets a received initialization on both Initialization and
// PaymentMethodBase.Initialization.
func (m *MultibancoPaymentMethod) UnmarshalJSON(data []byte) error {
	type alias MultibancoPaymentMethod
	if err := json.Unmarshal(data, (*alias)(m)); err != nil {
		return err
	}
	m.PaymentMethodBase.Initialization = m.Initialization
	return nil
}

// MarshalJSON sends Initialization, or PaymentMethodBase.Initialization when Initialization is
// empty.
func (m P24PaymentMethod) MarshalJSON() ([]byte, error) {
	// The alias sheds the method set, so json.Marshal below does not re-enter this method.
	type alias P24PaymentMethod
	if m.Initialization == "" {
		m.Initialization = m.PaymentMethodBase.Initialization
	}
	return json.Marshal(alias(m))
}

// UnmarshalJSON sets a received initialization on both Initialization and
// PaymentMethodBase.Initialization.
func (m *P24PaymentMethod) UnmarshalJSON(data []byte) error {
	type alias P24PaymentMethod
	if err := json.Unmarshal(data, (*alias)(m)); err != nil {
		return err
	}
	m.PaymentMethodBase.Initialization = m.Initialization
	return nil
}

// MarshalJSON sends Initialization, or PaymentMethodBase.Initialization when Initialization is
// empty.
func (m SwishPaymentMethod) MarshalJSON() ([]byte, error) {
	// The alias sheds the method set, so json.Marshal below does not re-enter this method.
	type alias SwishPaymentMethod
	if m.Initialization == "" {
		m.Initialization = m.PaymentMethodBase.Initialization
	}
	return json.Marshal(alias(m))
}

// UnmarshalJSON sets a received initialization on both Initialization and
// PaymentMethodBase.Initialization.
func (m *SwishPaymentMethod) UnmarshalJSON(data []byte) error {
	type alias SwishPaymentMethod
	if err := json.Unmarshal(data, (*alias)(m)); err != nil {
		return err
	}
	m.PaymentMethodBase.Initialization = m.Initialization
	return nil
}

// MarshalJSON sends Initialization, or PaymentMethodBase.Initialization when Initialization is
// empty.
func (m AchPaymentMethod) MarshalJSON() ([]byte, error) {
	// The alias sheds the method set, so json.Marshal below does not re-enter this method.
	type alias AchPaymentMethod
	if m.Initialization == "" {
		m.Initialization = m.PaymentMethodBase.Initialization
	}
	return json.Marshal(alias(m))
}

// UnmarshalJSON sets a received initialization on both Initialization and
// PaymentMethodBase.Initialization.
func (m *AchPaymentMethod) UnmarshalJSON(data []byte) error {
	type alias AchPaymentMethod
	if err := json.Unmarshal(data, (*alias)(m)); err != nil {
		return err
	}
	m.PaymentMethodBase.Initialization = m.Initialization
	return nil
}

// MarshalJSON sends Initialization, or PaymentMethodBase.Initialization when Initialization is
// empty.
func (m SepaPaymentMethod) MarshalJSON() ([]byte, error) {
	// The alias sheds the method set, so json.Marshal below does not re-enter this method.
	type alias SepaPaymentMethod
	if m.Initialization == "" {
		m.Initialization = m.PaymentMethodBase.Initialization
	}
	return json.Marshal(alias(m))
}

// UnmarshalJSON sets a received initialization on both Initialization and
// PaymentMethodBase.Initialization.
func (m *SepaPaymentMethod) UnmarshalJSON(data []byte) error {
	type alias SepaPaymentMethod
	if err := json.Unmarshal(data, (*alias)(m)); err != nil {
		return err
	}
	m.PaymentMethodBase.Initialization = m.Initialization
	return nil
}

// MarshalJSON sends Initialization, or PaymentMethodBase.Initialization when Initialization is
// empty.
func (m PayByBankPaymentMethod) MarshalJSON() ([]byte, error) {
	// The alias sheds the method set, so json.Marshal below does not re-enter this method.
	type alias PayByBankPaymentMethod
	if m.Initialization == "" {
		m.Initialization = m.PaymentMethodBase.Initialization
	}
	return json.Marshal(alias(m))
}

// UnmarshalJSON sets a received initialization on both Initialization and
// PaymentMethodBase.Initialization.
func (m *PayByBankPaymentMethod) UnmarshalJSON(data []byte) error {
	type alias PayByBankPaymentMethod
	if err := json.Unmarshal(data, (*alias)(m)); err != nil {
		return err
	}
	m.PaymentMethodBase.Initialization = m.Initialization
	return nil
}

// MarshalJSON sends Initialization, or PaymentMethodBase.Initialization when Initialization is
// empty.
func (m StablecoinPaymentMethod) MarshalJSON() ([]byte, error) {
	// The alias sheds the method set, so json.Marshal below does not re-enter this method.
	type alias StablecoinPaymentMethod
	if m.Initialization == "" {
		m.Initialization = m.PaymentMethodBase.Initialization
	}
	return json.Marshal(alias(m))
}

// UnmarshalJSON sets a received initialization on both Initialization and
// PaymentMethodBase.Initialization.
func (m *StablecoinPaymentMethod) UnmarshalJSON(data []byte) error {
	type alias StablecoinPaymentMethod
	if err := json.Unmarshal(data, (*alias)(m)); err != nil {
		return err
	}
	m.PaymentMethodBase.Initialization = m.Initialization
	return nil
}

// MarshalJSON sends Initialization, or PaymentMethodBase.Initialization when Initialization is
// empty.
func (m BlikPaymentMethod) MarshalJSON() ([]byte, error) {
	// The alias sheds the method set, so json.Marshal below does not re-enter this method.
	type alias BlikPaymentMethod
	if m.Initialization == "" {
		m.Initialization = m.PaymentMethodBase.Initialization
	}
	return json.Marshal(alias(m))
}

// UnmarshalJSON sets a received initialization on both Initialization and
// PaymentMethodBase.Initialization.
func (m *BlikPaymentMethod) UnmarshalJSON(data []byte) error {
	type alias BlikPaymentMethod
	if err := json.Unmarshal(data, (*alias)(m)); err != nil {
		return err
	}
	m.PaymentMethodBase.Initialization = m.Initialization
	return nil
}

// MarshalJSON sends Initialization, or PaymentMethodBase.Initialization when Initialization is
// empty.
func (m IdealPaymentMethod) MarshalJSON() ([]byte, error) {
	// The alias sheds the method set, so json.Marshal below does not re-enter this method.
	type alias IdealPaymentMethod
	if m.Initialization == "" {
		m.Initialization = m.PaymentMethodBase.Initialization
	}
	return json.Marshal(alias(m))
}

// UnmarshalJSON sets a received initialization on both Initialization and
// PaymentMethodBase.Initialization.
func (m *IdealPaymentMethod) UnmarshalJSON(data []byte) error {
	type alias IdealPaymentMethod
	if err := json.Unmarshal(data, (*alias)(m)); err != nil {
		return err
	}
	m.PaymentMethodBase.Initialization = m.Initialization
	return nil
}

// MarshalJSON sends Initialization, or PaymentMethodBase.Initialization when Initialization is
// empty.
func (m KnetPaymentMethod) MarshalJSON() ([]byte, error) {
	// The alias sheds the method set, so json.Marshal below does not re-enter this method.
	type alias KnetPaymentMethod
	if m.Initialization == "" {
		m.Initialization = m.PaymentMethodBase.Initialization
	}
	return json.Marshal(alias(m))
}

// UnmarshalJSON sets a received initialization on both Initialization and
// PaymentMethodBase.Initialization.
func (m *KnetPaymentMethod) UnmarshalJSON(data []byte) error {
	type alias KnetPaymentMethod
	if err := json.Unmarshal(data, (*alias)(m)); err != nil {
		return err
	}
	m.PaymentMethodBase.Initialization = m.Initialization
	return nil
}

// MarshalJSON sends Initialization, or PaymentMethodBase.Initialization when Initialization is
// empty.
func (m BancontactPaymentMethod) MarshalJSON() ([]byte, error) {
	// The alias sheds the method set, so json.Marshal below does not re-enter this method.
	type alias BancontactPaymentMethod
	if m.Initialization == "" {
		m.Initialization = m.PaymentMethodBase.Initialization
	}
	return json.Marshal(alias(m))
}

// UnmarshalJSON sets a received initialization on both Initialization and
// PaymentMethodBase.Initialization.
func (m *BancontactPaymentMethod) UnmarshalJSON(data []byte) error {
	type alias BancontactPaymentMethod
	if err := json.Unmarshal(data, (*alias)(m)); err != nil {
		return err
	}
	m.PaymentMethodBase.Initialization = m.Initialization
	return nil
}
