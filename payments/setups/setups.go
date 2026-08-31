package setups

import (
	"time"

	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/payments"
)

const (
	PaymentSetupsPath = "payments/setups"
	ConfirmPath       = "confirm"
)

type PaymentMethodInitialization string

const (
	PaymentMethodInitializationDisabled PaymentMethodInitialization = "disabled"
	PaymentMethodInitializationEnabled  PaymentMethodInitialization = "enabled"
)

// ===== Main Request/Response Structs =====

type PaymentSetupRequest struct {
	ProcessingChannelId       string                                 `json:"processing_channel_id"`
	Amount                    int64                                  `json:"amount"`
	Currency                  common.Currency                        `json:"currency"`
	PaymentType               payments.PaymentType                   `json:"payment_type,omitempty"`
	Reference                 string                                 `json:"reference,omitempty"`
	Description               string                                 `json:"description,omitempty"`
	PaymentMethods            *PaymentMethods                        `json:"payment_methods,omitempty"`
	Settings                  *PaymentSetupSettings                  `json:"settings,omitempty"`
	Customer                  *PaymentSetupCustomer                  `json:"customer,omitempty"`
	Order                     *PaymentSetupOrder                     `json:"order,omitempty"`
	Billing                   *PaymentSetupBilling                   `json:"billing,omitempty"`
	Industry                  *PaymentSetupIndustry                  `json:"industry,omitempty"`
	AccountFundingTransaction *PaymentSetupAccountFundingTransaction `json:"account_funding_transaction,omitempty"`
	BillingDescriptor         *PaymentSetupBillingDescriptor         `json:"billing_descriptor,omitempty"`
	PresentmentDetails        *PaymentSetupPresentmentDetails        `json:"presentment_details,omitempty"`
	Terminal                  *PaymentSetupTerminal                  `json:"terminal,omitempty"`
}

type PaymentSetupResponse struct {
	HttpMetadata              common.HttpMetadata
	Id                        string                                 `json:"id,omitempty"`
	ProcessingChannelId       string                                 `json:"processing_channel_id"`
	Amount                    int64                                  `json:"amount"`
	Currency                  common.Currency                        `json:"currency"`
	PaymentType               payments.PaymentType                   `json:"payment_type,omitempty"`
	Reference                 string                                 `json:"reference,omitempty"`
	Description               string                                 `json:"description,omitempty"`
	PaymentMethods            *PaymentMethods                        `json:"payment_methods,omitempty"`
	AvailablePaymentMethods   []string                               `json:"available_payment_methods,omitempty"`
	Settings                  *PaymentSetupSettings                  `json:"settings,omitempty"`
	Customer                  *PaymentSetupCustomer                  `json:"customer,omitempty"`
	Order                     *PaymentSetupOrder                     `json:"order,omitempty"`
	Billing                   *PaymentSetupBilling                   `json:"billing,omitempty"`
	Industry                  *PaymentSetupIndustry                  `json:"industry,omitempty"`
	AccountFundingTransaction *PaymentSetupAccountFundingTransaction `json:"account_funding_transaction,omitempty"`
	BillingDescriptor         *PaymentSetupBillingDescriptor         `json:"billing_descriptor,omitempty"`
	PresentmentDetails        *PaymentSetupPresentmentDetails        `json:"presentment_details,omitempty"`
	Terminal                  *PaymentSetupTerminal                  `json:"terminal,omitempty"`
	LatestPayment             map[string]interface{}                 `json:"latest_payment,omitempty"`
}

// PaymentSetupBilling is the billing details for the payment.
type PaymentSetupBilling struct {
	Address *common.Address `json:"address,omitempty"`
}

// PaymentSetupBillingDescriptor is the billing descriptor for the payment.
type PaymentSetupBillingDescriptor struct {
	Name      string `json:"name,omitempty"`
	City      string `json:"city,omitempty"`
	Reference string `json:"reference,omitempty"`
}

// PaymentSetupPresentmentDetails is the amount and currency to present to the customer, when the
// settlement currency differs from the customer-facing currency.
type PaymentSetupPresentmentDetails struct {
	Amount   int64           `json:"amount,omitempty"`
	Currency common.Currency `json:"currency,omitempty"`
}

// PaymentSetupTerminal holds terminal details.
type PaymentSetupTerminal struct {
	Id            string     `json:"id,omitempty"`
	LocalDateTime *time.Time `json:"local_date_time,omitempty"`
}

// ===== Customer Structs =====

type PaymentSetupCustomer struct {
	Id              string                      `json:"id,omitempty"`
	Country         common.Country              `json:"country,omitempty"`
	Email           *PaymentSetupCustomerEmail  `json:"email,omitempty"`
	Name            string                      `json:"name,omitempty"`
	TaxNumber       string                      `json:"tax_number,omitempty"`
	Phone           *common.Phone               `json:"phone,omitempty"`
	Device          *PaymentSetupCustomerDevice `json:"device,omitempty"`
	MerchantAccount *CustomerMerchantAccount    `json:"merchant_account,omitempty"`
}

type PaymentSetupCustomerEmail struct {
	Address  string `json:"address,omitempty"`
	Verified *bool  `json:"verified,omitempty"`
}

type PaymentSetupCustomerDevice struct {
	Locale string `json:"locale,omitempty"`
}

type CustomerMerchantAccount struct {
	Id                   string     `json:"id,omitempty"`
	RegistrationDate     *time.Time `json:"registration_date,omitempty"`
	LastModified         *time.Time `json:"last_modified,omitempty"`
	ReturningCustomer    *bool      `json:"returning_customer,omitempty"`
	FirstTransactionDate *time.Time `json:"first_transaction_date,omitempty"`
	LastTransactionDate  *time.Time `json:"last_transaction_date,omitempty"`
	TotalOrderCount      int        `json:"total_order_count,omitempty"`
	LastPaymentAmount    int64      `json:"last_payment_amount,omitempty"`
}

// ===== Payment Methods Structs =====

type PaymentMethods struct {
	Instrument  *InstrumentPaymentMethod  `json:"instrument,omitempty"`
	Klarna      *KlarnaPaymentMethod      `json:"klarna,omitempty"`
	Stcpay      *StcpayPaymentMethod      `json:"stcpay,omitempty"`
	Tabby       *TabbyPaymentMethod       `json:"tabby,omitempty"`
	Bizum       *BizumPaymentMethod       `json:"bizum,omitempty"`
	Paynow      *SimplePaymentMethod      `json:"paynow,omitempty"`
	Qpay        *SimplePaymentMethod      `json:"qpay,omitempty"`
	Eps         *SimplePaymentMethod      `json:"eps,omitempty"`
	Ideal       *SimplePaymentMethod      `json:"ideal,omitempty"`
	Knet        *SimplePaymentMethod      `json:"knet,omitempty"`
	Bancontact  *SimplePaymentMethod      `json:"bancontact,omitempty"`
	Benefit     *SimplePaymentMethod      `json:"benefit,omitempty"`
	Vipps       *SimplePaymentMethod      `json:"vipps,omitempty"`
	Twint       *SimplePaymentMethod      `json:"twint,omitempty"`
	AlipayCn    *SimplePaymentMethod      `json:"alipay_cn,omitempty"`
	AlipayHk    *SimplePaymentMethod      `json:"alipay_hk,omitempty"`
	Gcash       *SimplePaymentMethod      `json:"gcash,omitempty"`
	Tng         *SimplePaymentMethod      `json:"tng,omitempty"`
	Dana        *SimplePaymentMethod      `json:"dana,omitempty"`
	Mobilepay   *SimplePaymentMethod      `json:"mobilepay,omitempty"`
	Tamara      *SimplePaymentMethod      `json:"tamara,omitempty"`
	Mbway       *SimplePaymentMethod      `json:"mbway,omitempty"`
	Multibanco  *MultibancoPaymentMethod  `json:"multibanco,omitempty"`
	Wechatpay   *SimplePaymentMethod      `json:"wechatpay,omitempty"`
	Kakaopay    *SimplePaymentMethod      `json:"kakaopay,omitempty"`
	Truemoney   *SimplePaymentMethod      `json:"truemoney,omitempty"`
	Octopus     *SimplePaymentMethod      `json:"octopus,omitempty"`
	P24         *P24PaymentMethod         `json:"p24,omitempty"`
	Alma        *SimplePaymentMethod      `json:"alma,omitempty"`
	Swish       *SwishPaymentMethod       `json:"swish,omitempty"`
	Sequra      *SimplePaymentMethod      `json:"sequra,omitempty"`
	Ach         *AchPaymentMethod         `json:"ach,omitempty"`
	Sepa        *SepaPaymentMethod        `json:"sepa,omitempty"`
	Paypal      *PaypalPaymentMethod      `json:"paypal,omitempty"`
	Googlepay   *SimplePaymentMethod      `json:"googlepay,omitempty"`
	Applepay    *SimplePaymentMethod      `json:"applepay,omitempty"`
	Card        *SimplePaymentMethod      `json:"card,omitempty"`
	Blik        *BlikPaymentMethod        `json:"blik,omitempty"`
	Bacs        *BacsPaymentMethod        `json:"bacs,omitempty"`
	CardPresent *CardPresentPaymentMethod `json:"card_present,omitempty"`
	PayByBank   *PayByBankPaymentMethod   `json:"pay_by_bank,omitempty"`
	Stablecoin  *StablecoinPaymentMethod  `json:"stablecoin,omitempty"`
}

type PaymentMethodBase struct {
	Status         string                      `json:"status,omitempty"`
	Flags          []string                    `json:"flags,omitempty"`
	Initialization PaymentMethodInitialization `json:"initialization,omitempty"`
}

type PaymentMethodOption struct {
	Id     string               `json:"id,omitempty"`
	Status string               `json:"status,omitempty"`
	Flags  []string             `json:"flags,omitempty"`
	Action *PaymentMethodAction `json:"action,omitempty"`
}

type PaymentMethodAction struct {
	Type        string `json:"type,omitempty"`
	ClientToken string `json:"client_token,omitempty"`
	SessionId   string `json:"session_id,omitempty"`
}

type PaymentMethodOptions struct {
	Sdk          *PaymentMethodOption `json:"sdk,omitempty"`
	PayInFull    *PaymentMethodOption `json:"pay_in_full,omitempty"`
	Installments *PaymentMethodOption `json:"installments,omitempty"`
	PayNow       *PaymentMethodOption `json:"pay_now,omitempty"`
}

type KlarnaPaymentMethod struct {
	PaymentMethodBase
	AccountHolder        *KlarnaAccountHolder  `json:"account_holder,omitempty"`
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}

type KlarnaAccountHolder struct {
	Name string `json:"name,omitempty"`
}

type StcpayPaymentMethod struct {
	PaymentMethodBase
	Otp                  string                `json:"otp,omitempty"`
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}

type TabbyPaymentMethod struct {
	PaymentMethodBase
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}

type BizumPaymentMethod struct {
	PaymentMethodBase
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}

type SimplePaymentMethod struct {
	PaymentMethodBase
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}

type InstrumentPaymentMethod struct {
	PaymentMethodBase
	Id                   string                `json:"id,omitempty"`
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}

type MultibancoPaymentMethod struct {
	PaymentMethodBase
	AccountHolderName    string                `json:"account_holder_name,omitempty"`
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}

type P24AccountHolder struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type P24PaymentMethod struct {
	PaymentMethodBase
	AccountHolder        *P24AccountHolder     `json:"account_holder,omitempty"`
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}

type SwishAccountHolder struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

type SwishPaymentMethod struct {
	PaymentMethodBase
	BillingDescriptor    string                `json:"billing_descriptor,omitempty"`
	AccountHolder        *SwishAccountHolder   `json:"account_holder,omitempty"`
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}

type AchAccountHolder struct {
	Type      string `json:"type,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

type AchPaymentMethod struct {
	PaymentMethodBase
	AccountType          string                `json:"account_type,omitempty"`
	AccountHolder        *AchAccountHolder     `json:"account_holder,omitempty"`
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}

type SepaAccountHolder struct {
	Type        string `json:"type,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	CompanyName string `json:"company_name,omitempty"`
}

type SepaPaymentMethod struct {
	PaymentMethodBase
	AccountHolder        *SepaAccountHolder    `json:"account_holder,omitempty"`
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}

type PaypalPaymentMethod struct {
	PaymentMethodBase
	UserAction           string                `json:"user_action,omitempty"`
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}

// BacsAccountHolderType is the type of Bacs account holder.
type BacsAccountHolderType string

const (
	BacsAccountHolderIndividual BacsAccountHolderType = "individual"
	BacsAccountHolderCorporate  BacsAccountHolderType = "corporate"
)

type BacsAccountHolder struct {
	Type        BacsAccountHolderType `json:"type,omitempty"`
	FirstName   string                `json:"first_name,omitempty"`
	LastName    string                `json:"last_name,omitempty"`
	CompanyName string                `json:"company_name,omitempty"`
	Email       string                `json:"email,omitempty"`
}

// BacsPaymentMethod is the Bacs payment method's details and configuration.
type BacsPaymentMethod struct {
	PaymentMethodBase
	InstrumentId         string                `json:"instrument_id,omitempty"`
	AccountHolder        *BacsAccountHolder    `json:"account_holder,omitempty"`
	AccountNumber        string                `json:"account_number,omitempty"`
	BankCode             string                `json:"bank_code,omitempty"`
	Country              common.Country        `json:"country,omitempty"`
	Currency             string                `json:"currency,omitempty"`
	AllowPartialMatch    bool                  `json:"allow_partial_match,omitempty"`
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}

type CardPresentPin struct {
	KeySetId    string `json:"key_set_id,omitempty"`
	Block       string `json:"block,omitempty"`
	BlockFormat string `json:"block_format,omitempty"`
}

// CardPresentPaymentMethod is the Card Present payment method's details and configuration.
type CardPresentPaymentMethod struct {
	PaymentMethodBase
	Track2            string          `json:"track2,omitempty"`
	Emv               string          `json:"emv,omitempty"`
	EntryMode         string          `json:"entry_mode,omitempty"`
	Pin               *CardPresentPin `json:"pin,omitempty"`
	StoreForFutureUse bool            `json:"store_for_future_use,omitempty"`
	Name              string          `json:"name,omitempty"`
}

type PayByBankBank struct {
	BankId      string `json:"bank_id,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	LogoUrl     string `json:"logo_url,omitempty"`
	Available   bool   `json:"available,omitempty"`
}

type PayByBankAction struct {
	Type  string          `json:"type,omitempty"`
	Banks []PayByBankBank `json:"banks,omitempty"`
}

// PayByBankPaymentMethod is the Pay by Bank (Open Banking) payment method's details and configuration.
type PayByBankPaymentMethod struct {
	PaymentMethodBase
	BankId string           `json:"bank_id,omitempty"`
	Action *PayByBankAction `json:"action,omitempty"`
}

// StablecoinPaymentMethod is the Stablecoin payment method's details and configuration.
type StablecoinPaymentMethod struct {
	PaymentMethodBase
}

// ===== Support Structs =====

type PaymentSetupSettings struct {
	SuccessUrl             string   `json:"success_url,omitempty"`
	FailureUrl             string   `json:"failure_url,omitempty"`
	Capture                bool     `json:"capture,omitempty"`
	ExcludedPaymentMethods []string `json:"excluded_payment_methods,omitempty"`
}

type PaymentSetupOrder struct {
	Items             []payments.Product             `json:"items,omitempty"`
	Shipping          *payments.ShippingDetails      `json:"shipping,omitempty"`
	SubMerchants      []OrderSubMerchant             `json:"sub_merchants,omitempty"`
	InvoiceId         string                         `json:"invoice_id,omitempty"`
	ShippingAmount    int                            `json:"shipping_amount,omitempty"`
	DiscountAmount    int                            `json:"discount_amount,omitempty"`
	SurchargeAmount   int                            `json:"surcharge_amount,omitempty"`
	TaxAmount         int                            `json:"tax_amount,omitempty"`
	TippingAmount     int                            `json:"tipping_amount,omitempty"`
	AmountAllocations []PaymentSetupAmountAllocation `json:"amount_allocations,omitempty"`
}

// AmountAllocationCommission is the commission to collect from an amount allocation split.
type AmountAllocationCommission struct {
	Amount     int64   `json:"amount,omitempty"`
	Percentage float64 `json:"percentage,omitempty"`
}

// PaymentSetupAmountAllocation represents a sub-entity on whose behalf the payment is processed.
type PaymentSetupAmountAllocation struct {
	Id         string                      `json:"id,omitempty"`
	Amount     int64                       `json:"amount,omitempty"`
	Reference  string                      `json:"reference,omitempty"`
	Commission *AmountAllocationCommission `json:"commission,omitempty"`
}

type OrderSubMerchant struct {
	Id               string     `json:"id,omitempty"`
	ProductCategory  string     `json:"product_category,omitempty"`
	NumberOfSales    int        `json:"number_of_sales,omitempty"`
	RegistrationDate *time.Time `json:"registration_date,omitempty"`
}

type PaymentSetupIndustry struct {
	AirlineData       *AirlineData                 `json:"airline_data,omitempty"`
	AccommodationData []payments.AccommodationData `json:"accommodation_data,omitempty"`
}

type AirlineData struct {
	Ticket           *payments.Ticket            `json:"ticket,omitempty"`
	Passengers       []payments.Passenger        `json:"passengers,omitempty"`
	FlightLegDetails []payments.FlightLegDetails `json:"flight_leg_details,omitempty"`
}

type AccountFundingTransactionPurpose string

const (
	AFTPurposeDonations         AccountFundingTransactionPurpose = "donations"
	AFTPurposeEducation         AccountFundingTransactionPurpose = "education"
	AFTPurposeEmergencyNeed     AccountFundingTransactionPurpose = "emergency_need"
	AFTPurposeExpatriation      AccountFundingTransactionPurpose = "expatriation"
	AFTPurposeFamilySupport     AccountFundingTransactionPurpose = "family_support"
	AFTPurposeFinancialServices AccountFundingTransactionPurpose = "financial_services"
	AFTPurposeGifts             AccountFundingTransactionPurpose = "gifts"
	AFTPurposeIncome            AccountFundingTransactionPurpose = "income"
	AFTPurposeInsurance         AccountFundingTransactionPurpose = "insurance"
	AFTPurposeInvestment        AccountFundingTransactionPurpose = "investment"
	AFTPurposeItServices        AccountFundingTransactionPurpose = "it_services"
	AFTPurposeLeisure           AccountFundingTransactionPurpose = "leisure"
	AFTPurposeLoanPayment       AccountFundingTransactionPurpose = "loan_payment"
	AFTPurposeMedicalTreatment  AccountFundingTransactionPurpose = "medical_treatment"
	AFTPurposeOther             AccountFundingTransactionPurpose = "other"
	AFTPurposePension           AccountFundingTransactionPurpose = "pension"
	AFTPurposeRoyalties         AccountFundingTransactionPurpose = "royalties"
	AFTPurposeSavings           AccountFundingTransactionPurpose = "savings"
	AFTPurposeTravelAndTourism  AccountFundingTransactionPurpose = "travel_and_tourism"
)

type AccountFundingTransactionIdentificationType string

const (
	AFTIdentificationPassport       AccountFundingTransactionIdentificationType = "passport"
	AFTIdentificationDrivingLicense AccountFundingTransactionIdentificationType = "driving_license"
	AFTIdentificationNationalId     AccountFundingTransactionIdentificationType = "national_id"
)

type AccountFundingTransactionIdentification struct {
	Type           AccountFundingTransactionIdentificationType `json:"type,omitempty"`
	Number         string                                      `json:"number,omitempty"`
	IssuingCountry string                                      `json:"issuing_country,omitempty"`
}

type AccountFundingTransactionSender struct {
	DateOfBirth    *time.Time                               `json:"date_of_birth,omitempty"`
	Reference      string                                   `json:"reference,omitempty"`
	Identification *AccountFundingTransactionIdentification `json:"identification,omitempty"`
}

type AccountFundingTransactionRecipient struct {
	DateOfBirth   *time.Time      `json:"date_of_birth,omitempty"`
	AccountNumber string          `json:"account_number,omitempty"`
	FirstName     string          `json:"first_name,omitempty"`
	LastName      string          `json:"last_name,omitempty"`
	Address       *common.Address `json:"address,omitempty"`
}

type PaymentSetupAccountFundingTransaction struct {
	Enabled   *bool                               `json:"enabled,omitempty"`
	Purpose   AccountFundingTransactionPurpose    `json:"purpose,omitempty"`
	Sender    *AccountFundingTransactionSender    `json:"sender,omitempty"`
	Recipient *AccountFundingTransactionRecipient `json:"recipient,omitempty"`
}

type BlikPaymentMethod struct {
	PaymentMethodBase
	PartnerCode          string                `json:"partner_code,omitempty"`
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options,omitempty"`
}
