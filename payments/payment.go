package payments

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/common"
)

type (
	// Request ...
	Request struct {
		Source            interface{}        `json:"source,omitempty"`
		Destination       interface{}        `json:"destination,omitempty"`
		Amount            uint64             `json:"amount,omitempty"`
		Currency          string             `json:"currency"`
		Reference         string             `json:"reference,omitempty"`
		PaymentType       common.PaymentType `json:"payment_type,omitempty"`
		MerchantInitiated *bool              `json:"merchant_initiated,omitempty"`
		Description       string             `json:"description,omitempty"`
		Capture           *bool              `json:"capture,omitempty"`
		CaptureOn         *time.Time         `json:"capture_on,omitempty"`
		Customer          *Customer          `json:"customer,omitempty"`
		BillingDescriptor *BillingDescriptor `json:"billing_descriptor,omitempty"`
		Shipping          *Shipping          `json:"shipping,omitempty"`
		ThreeDS           *ThreeDS           `json:"3ds,omitempty"`
		PreviousPaymentID string             `json:"previous_payment_id,omitempty"`
		Risk              *Risk              `json:"risk,omitempty"`
		SuccessURL        string             `json:"success_url,omitempty,omitempty"`
		FailureURL        string             `json:"failure_url,omitempty,omitempty"`
		PaymentIP         string             `json:"payment_ip,omitempty"`
		Recipient         *Recipient         `json:"recipient,omitempty"`
		Destinations      []*Destination     `json:"destinations,omitempty"`
		Processing        *Processing        `json:"processing,omitempty"`
		FundTransferType  string             `json:"fund_transfer_type,omitempty"`
		Metadata          map[string]string  `json:"metadata,omitempty"`
		// FOUR only
		AuthorizationType   string `json:"authorization_type,omitempty"`
		ProcessingChannelId string `json:"processing_channel_id,omitempty"`
	}

	// IDSource ...
	IDSource struct {
		Type string `json:"type" binding:"required"`
		ID   string `json:"id" binding:"required"`
		CVV  string `json:"cvv,omitempty"`
	}

	// CardSource ...
	CardSource struct {
		Type           string          `json:"type" binding:"required"`
		Number         string          `json:"number" binding:"required"`
		ExpiryMonth    uint64          `json:"expiry_month" binding:"required"`
		ExpiryYear     uint64          `json:"expiry_year" binding:"required"`
		Name           string          `json:"name,omitempty"`
		CVV            string          `json:"cvv,omitempty"`
		Stored         *bool           `json:"stored,omitempty"`
		BillingAddress *common.Address `json:"billing_address,omitempty"`
		Phone          *common.Phone   `json:"phone,omitempty"`
	}

	// TokenSource ...
	TokenSource struct {
		Type           string          `json:"type" binding:"required"`
		Token          string          `json:"token" binding:"required"`
		BillingAddress *common.Address `json:"billing_address,omitempty"`
		Phone          *common.Phone   `json:"phone,omitempty"`
	}

	// CustomerSource ...
	CustomerSource struct {
		Type  string `json:"type" binding:"required"`
		ID    string `json:"id,omitempty"`
		Email string `json:"email,omitempty"`
	}

	// NetworkTokenSource ...
	NetworkTokenSource struct {
		Type           string          `json:"type" binding:"required"`
		Token          string          `json:"token" binding:"required"`
		ExpiryMonth    uint64          `json:"expiry_month" binding:"required"`
		ExpiryYear     uint64          `json:"expiry_year" binding:"required"`
		TokenType      string          `json:"token_type" binding:"required"`
		Cryptogram     string          `json:"cryptogram" binding:"required"`
		ECI            string          `json:"eci" binding:"required"`
		Stored         *bool           `json:"stored,omitempty"`
		Name           string          `json:"name,omitempty"`
		CVV            string          `json:"cvv,omitempty"`
		BillingAddress *common.Address `json:"billing_address,omitempty"`
		Phone          *common.Phone   `json:"phone,omitempty"`
	}

	// AlipaySource ...
	AlipaySource struct {
		Type string `json:"type" binding:"required"`
	}

	// BenefitpaySource ...
	BenefitpaySource struct {
		Type            string `json:"type" binding:"required"`
		IntegrationType string `json:"integration_type" binding:"required"`
	}

	// BalotoSource ...
	BalotoSource struct {
		Type            string `json:"type" binding:"required"`
		IntegrationType string `json:"integration_type" binding:"required"`
		Country         string `json:"country" binding:"required"`
		Description     string `json:"description,omitempty"`
		Payer           *Payer `json:"payer,omitempty"`
	}

	// BoletoSource ...
	BoletoSource struct {
		Type            string `json:"type" binding:"required"`
		IntegrationType string `json:"integration_type" binding:"required"`
		Country         string `json:"country" binding:"required"`
		Description     string `json:"description,omitempty"`
		Payer           *Payer `json:"payer,omitempty"`
	}

	// Payer -
	Payer struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Document string `json:"document" binding:"required"`
	}

	// EPSSource -
	EPSSource struct {
		Type    string `json:"type" binding:"required"`
		Purpose string `json:"purpose" binding:"required"`
		BIC     string `json:"bic" binding:"required"`
	}

	// GiropaySource ...
	GiropaySource struct {
		Type       string      `json:"type" binding:"required"`
		Purpose    string      `json:"purpose" binding:"required"`
		BIC        string      `json:"bic,omitempty"`
		InfoFields []InfoField `json:"info_fields,omitempty"`
	}

	// InfoField ...
	InfoField struct {
		Label string `json:"label,omitempty"`
		Text  string `json:"text,omitempty"`
	}

	// IDealSource ...
	IDealSource struct {
		Type        string `json:"type" binding:"required"`
		Description string `json:"description,omitempty"`
		BIC         string `json:"bic" binding:"required"`
		Language    string `json:"language,omitempty"`
	}

	// KlarnaSource ...
	KlarnaSource struct {
		Type               string            `json:"type" binding:"required"`
		AuthorizationToken string            `json:"authorization_token" binding:"required"`
		Locale             string            `json:"locale" binding:"required"`
		PurchaseCountry    string            `json:"purchase_country" binding:"required"`
		AutoCapture        string            `json:"auto_capture,omitempty"`
		BillingAddress     *KlarnaAddress    `json:"billing_address,omitempty"`
		ShippingAddress    *KlarnaAddress    `json:"shipping_address,omitempty"`
		TaxAmount          uint64            `json:"tax_amount,omitempty"`
		Product            *KlarnaProduct    `json:"products,omitempty"`
		Customer           *KlarnaCustomer   `json:"customer,omitempty"`
		MerchantReference1 string            `json:"merchant_reference1,omitempty"`
		MerchantReference2 string            `json:"merchant_reference2,omitempty"`
		MerchantData       string            `json:"merchant_data,omitempty"`
		Attachment         *KlarnaAttachment `json:"attachment,omitempty"`
	}

	// KlarnaAddress ...
	KlarnaAddress struct {
		Attention        string `json:"attention,omitempty"`
		City             string `json:"city,omitempty"`
		Country          string `json:"country,omitempty"`
		Email            string `json:"email,omitempty"`
		FamilyName       string `json:"family_name,omitempty"`
		GivenName        string `json:"given_name,omitempty"`
		OrganizationName string `json:"organization_name,omitempty"`
		Phone            string `json:"phone,omitempty"`
		PostalCode       string `json:"postal_code,omitempty"`
		Region           string `json:"region,omitempty"`
		StreetAddress    string `json:"street_address,omitempty"`
		StreetAddress2   string `json:"street_address2,omitempty"`
		Title            string `json:"title,omitempty"`
	}

	// KlarnaCustomer ...
	KlarnaCustomer struct {
		DateOfBirth                  string `json:"date_of_birth,omitempty"`
		Gender                       string `json:"gender,omitempty"`
		LastFourSSN                  string `json:"last_four_ssn,omitempty"`
		NationalIdentificationNumber string `json:"national_identification_number,omitempty"`
		OrganizationEntityType       string `json:"organization_entity_type,omitempty"`
		OrganizationRegistrationID   string `json:"organization_registration_id,omitempty"`
		Title                        string `json:"title,omitempty"`
		Type                         string `json:"type,omitempty"`
		VatID                        string `json:"vat_id,omitempty"`
	}

	// KlarnaProduct ...
	KlarnaProduct struct {
		ImageURL            string                    `json:"image_url,omitempty"`
		MerchantData        string                    `json:"merchant_data,omitempty"`
		Name                string                    `json:"name,omitempty"`
		ProductIdentifiers  *KlarnaProductIdentifiers `json:"product_identifiers,omitempty"`
		ProductURL          string                    `json:"product_url,omitempty"`
		Quantity            uint64                    `json:"quantity,omitempty"`
		QuantityUnit        string                    `json:"quantity_unit,omitempty"`
		Reference           string                    `json:"reference,omitempty"`
		TaxRate             uint64                    `json:"tax_rate,omitempty"`
		TotalAmount         uint64                    `json:"total_amount,omitempty"`
		TotalDiscountAmount uint64                    `json:"total_discount_amount,omitempty"`
		TotalTaxAmount      uint64                    `json:"total_tax_amount,omitempty"`
		Type                string                    `json:"type,omitempty"`
		UnitPrice           uint64                    `json:"unit_price,omitempty"`
	}

	// KlarnaProductIdentifiers ...
	KlarnaProductIdentifiers struct {
		Brand                  string `json:"brand,omitempty"`
		CategoryPath           string `json:"category_path,omitempty"`
		GlobalTradeItemNumber  string `json:"global_trade_item_number,omitempty"`
		ManufacturerPartNumber string `json:"manufacturer_part_number,omitempty"`
	}

	// KlarnaAttachment ...
	KlarnaAttachment struct {
		Body        string `json:"body,omitempty"`
		ContentType string `json:"content_type,omitempty"`
	}
	// KNetSource ...
	KNetSource struct {
		Type              string `json:"type" binding:"required"`
		Language          string `json:"language" binding:"required"`
		UserDefinedField1 string `json:"user_defined_field1,omitempty"`
		UserDefinedField2 string `json:"user_defined_field2,omitempty"`
		UserDefinedField3 string `json:"user_defined_field3,omitempty"`
		UserDefinedField4 string `json:"user_defined_field4,omitempty"`
		UserDefinedField5 string `json:"user_defined_field5,omitempty"`
		CardToken         string `json:"card_token,omitempty"`
		PTLF              string `json:"ptlf,omitempty"`
	}

	// OxxoSource ...
	OxxoSource struct {
		Type            string `json:"type" binding:"required"`
		IntegrationType string `json:"integration_type" binding:"required"`
		Country         string `json:"country,omitempty"`
		Description     string `json:"description,omitempty"`
		Payer           *Payer `json:"payer,omitempty"`
	}

	// P24Source ...
	P24Source struct {
		Type               string `json:"type" binding:"required"`
		PaymentCountry     string `json:"payment_country" binding:"required"`
		AccountHolderName  string `json:"account_holder_name,omitempty"`
		AccountHolderEmail string `json:"account_holder_email,omitempty"`
		BillingDescriptor  string `json:"billing_descriptor,omitempty"`
	}

	// PagofacilSource ...
	PagofacilSource struct {
		Type            string `json:"type" binding:"required"`
		IntegrationType string `json:"integration_type" binding:"required"`
		Country         string `json:"country,omitempty"`
		Description     string `json:"description,omitempty"`
		Payer           *Payer `json:"payer,omitempty"`
	}

	// PayPalSource ...
	PayPalSource struct {
		Type string       `json:"type" binding:"required"`
		Plan *common.Plan `json:"plan,omitempty"`
	}

	// PoliSource ...
	PoliSource struct {
		Type string `json:"type" binding:"required"`
	}

	// RapipagoSource ...
	RapipagoSource struct {
		Type            string `json:"type" binding:"required"`
		IntegrationType string `json:"integration_type" binding:"required"`
		Country         string `json:"country,omitempty"`
		Description     string `json:"description,omitempty"`
		Payer           *Payer `json:"payer,omitempty"`
	}

	// SofortSource ...
	SofortSource struct {
		Type string `json:"type" binding:"required"`
	}

	// BancontactSource ...
	BancontactSource struct {
		Type              string `json:"type" binding:"required"`
		PaymentCountry    string `json:"payment_country,omitempty"`
		AccountHolderName string `json:"account_holder_name,omitempty"`
		BillingDescriptor string `json:"billing_descriptor,omitempty"`
		Language          string `json:"language,omitempty"`
	}

	// FawrySource ...
	FawrySource struct {
		Type              string          `json:"type" binding:"required"`
		Description       string          `json:"description,omitempty"`
		CustomerProfileID string          `json:"customer_profile_id,omitempty"`
		CustomerEmail     string          `json:"customer_email,omitempty"`
		CustomerMobile    string          `json:"customer_mobile,omitempty"`
		ExpiresOn         time.Time       `json:"expires_on,omitempty"`
		Products          *[]FawryProduct `json:"products,omitempty"`
	}

	// FawryProduct ...
	FawryProduct struct {
		ProductID   string `json:"product_id,omitempty"`
		Quantity    uint64 `json:"quantity,omitempty"`
		Price       uint64 `json:"price,omitempty"`
		Description string `json:"description,omitempty"`
	}

	// QPaySource ...
	QPaySource struct {
		Type        string `json:"type" binding:"required"`
		Quantity    uint64 `json:"quantity,omitempty"`
		Description string `json:"description,omitempty"`
		Language    string `json:"language,omitempty"`
		NationalID  string `json:"national_id,omitempty"`
	}

	// MultibancoSource ...
	MultibancoSource struct {
		Type              string `json:"type" binding:"required"`
		PaymentCountry    string `json:"payment_country,omitempty"`
		AccountHolderName string `json:"account_holder_name,omitempty"`
		BillingDescriptor string `json:"billing_descriptor,omitempty"`
	}

	// TokenDestination ...
	TokenDestination struct {
		Type           string          `json:"type" binding:"required"`
		Token          string          `json:"token" binding:"required"`
		FirstName      string          `json:"first_name" binding:"required"`
		LastName       string          `json:"last_name" binding:"required"`
		BillingAddress *common.Address `json:"billing_address,omitempty"`
		Phone          *common.Phone   `json:"phone,omitempty"`
	}

	// IDDestination ...
	IDDestination struct {
		Type      string `json:"type" binding:"required"`
		ID        string `json:"id" binding:"required"`
		FirstName string `json:"first_name,required"`
		LastName  string `json:"last_name,required"`
	}

	// CardDestination ...
	CardDestination struct {
		Type           string          `json:"type" binding:"required"`
		Number         string          `json:"number" binding:"required"`
		ExpiryMonth    uint64          `json:"expiry_month" binding:"required"`
		ExpiryYear     uint64          `json:"expiry_year" binding:"required"`
		FirstName      string          `json:"first_name,required"`
		LastName       string          `json:"last_name,required"`
		Name           string          `json:"name,omitempty"`
		BillingAddress *common.Address `json:"billing_address,omitempty"`
		Phone          *common.Phone   `json:"phone,omitempty"`
	}

	// Customer ...
	Customer struct {
		ID    string `json:"id,omitempty"`
		Email string `json:"email,omitempty"`
		Name  string `json:"name,omitempty"`
	}

	// BillingDescriptor ...
	BillingDescriptor struct {
		Name string `json:"name,omitempty"`
		City string `json:"city,omitempty"`
	}

	// Shipping ...
	Shipping struct {
		Address *common.Address `json:"address,omitempty"`
		Phone   *common.Phone   `json:"phone,omitempty"`
	}

	// Risk ...
	Risk struct {
		Enabled *bool `json:"enabled,omitempty"`
	}

	// RiskAssessment ...
	RiskAssessment struct {
		Flagged *bool `json:"flagged,omitempty"`
	}

	// Recipient ...
	Recipient struct {
		DOB           string `json:"dob"`
		AccountNumber string `json:"account_number"`
		ZIP           string `json:"zip"`
		LastName      string `json:"last_name"`
	}

	// Destination ...
	Destination struct {
		ID     string `json:"id"`
		Amount uint64 `json:"amount"`
	}

	// Processing - Use the processing object to influence or
	// override the data sent during card processing
	Processing struct {
		Mid                      string             `json:"mid,omitempty"`
		Aft                      *bool              `json:"aft,omitempty"`
		DLocal                   *DLocal            `json:"dlocal,omitempty"`
		AcquirerTransactionID    string             `json:"acquirer_transaction_id,omitempty"`
		AcquirerReferenceNumber  string             `json:"acquirer_reference_number,omitempty"`
		RetrievalReferenceNumber string             `json:"retrieval_reference_number,omitempty"`
		SenderInformation        *SenderInformation `json:"senderInformation,omitempty"`
	}

	// SenderInformation -
	SenderInformation struct {
		FirstName     string `json:"firstName" binding:"required"`
		LastName      string `json:"lastName" binding:"required"`
		Address       string `json:"address" binding:"required"`
		City          string `json:"city,omitempty"`
		State         string `json:"state,omitempty"`
		PostalCode    string `json:"postalCode" binding:"required"`
		Country       string `json:"country" binding:"required"`
		SourceOfFunds string `json:"sourceOfFunds" binding:"required"`
		AccountNumber string `json:"accountNumber" binding:"required"`
		Reference     string `json:"reference" binding:"required"`
	}

	// DLocal - Processing information required for dLocal payments.
	DLocal struct {
		Country      string        `json:"country,omitempty"`
		Payer        *Customer     `json:"payer,omitempty"`
		Installments *Installments `json:"installments,omitempty"`
	}

	// Installments - Details about the installments.
	Installments struct {
		Count string `json:"count,omitempty"`
	}
)

// SetSource ...
func (r *Request) SetSource(s interface{}) error {
	var err error
	switch p := s.(type) {
	case *IDSource:
	case *CardSource:
	case *TokenSource:
	case *NetworkTokenSource:
	case *AlipaySource:
	case *BenefitpaySource:
	case *BalotoSource:
	case *BoletoSource:
	case *EPSSource:
	case *GiropaySource:
	case *IDealSource:
	case *KlarnaSource:
	case *KNetSource:
	case *OxxoSource:
	case *P24Source:
	case *PagofacilSource:
	case *PayPalSource:
	case *PoliSource:
	case *RapipagoSource:
	case SofortSource:
	case BancontactSource:
	case FawrySource:
	case QPaySource:
	case MultibancoSource:
	case map[string]string:
	default:
		err = fmt.Errorf("Unsupported source type %T", p)
	}
	if err == nil {
		r.Source = s
	}
	return err
}

// SetDestination ...
func (r *Request) SetDestination(d interface{}) error {
	var err error
	switch p := d.(type) {
	case *IDDestination:
	case *CardDestination:
	case *TokenDestination:
	case map[string]string:
	default:
		err = fmt.Errorf("Unsupported source type %T", p)
	}
	if err == nil {
		r.Destination = d
	}
	return err
}

type (
	// Response ...
	Response struct {
		StatusResponse *checkout.StatusResponse `json:"api_response,omitempty"`
		Processed      *Processed               `json:"processed,omitempty"`
		Pending        *PaymentPending          `json:"pending,omitempty"`
	}
	// Processed ...
	Processed struct {
		ID                string                 `json:"id,omitempty"`
		ActionID          string                 `json:"action_id,omitempty"`
		Amount            uint64                 `json:"amount,omitempty"`
		Currency          string                 `json:"currency,omitempty"`
		Approved          *bool                  `json:"approved,omitempty"`
		Status            common.PaymentAction   `json:"status,omitempty"`
		AuthCode          string                 `json:"auth_code,omitempty"`
		ResponseCode      string                 `json:"response_code,omitempty"`
		ResponseSummary   string                 `json:"response_summary,omitempty"`
		ThreeDSEnrollment *ThreeDSEnrollment     `json:"3ds,omitempty"`
		Flagged           *bool                  `json:"flagged,omitempty"`
		RiskAssessment    *RiskAssessment        `json:"risk,omitempty"`
		Source            *SourceResponse        `json:"source,omitempty"`
		Destination       *DestinationResponse   `json:"destination,omitempty"`
		Customer          *Customer              `json:"customer,omitempty"`
		ProcessedOn       time.Time              `json:"processed_on,omitempty"`
		Reference         string                 `json:"reference,omitempty"`
		Processing        *Processing            `json:"processing,omitempty"`
		ECI               string                 `json:"eci,omitempty"`
		SchemeID          string                 `json:"scheme_id,omitempty"`
		Links             map[string]common.Link `json:"_links,omitempty"`
	}
	// PaymentPending ...
	PaymentPending struct {
		ID        string                 `json:"id,omitempty"`
		Status    common.PaymentAction   `json:"status,omitempty"`
		Reference string                 `json:"reference,omitempty"`
		Customer  *Customer              `json:"customer,omitempty"`
		ThreeDS   *ThreeDSEnrollment     `json:"3ds,omitempty"`
		Links     map[string]common.Link `json:"_links,omitempty"`
	}
	// PaymentResponse ...
	PaymentResponse struct {
		StatusResponse *checkout.StatusResponse `json:"api_response,omitempty"`
		Payment        *Payment                 `json:"payment,omitempty"`
	}
	// Payment ...
	Payment struct {
		ID                string               `json:"id,omitempty"`
		RequestedOn       time.Time            `json:"requested_on,omitempty"`
		Source            *SourceResponse      `json:"source,omitempty"`
		Amount            uint64               `json:"amount,omitempty"`
		Currency          string               `json:"currency,omitempty"`
		PaymentType       common.PaymentType   `json:"payment_type,omitempty"`
		Reference         string               `json:"reference,omitempty"`
		Description       string               `json:"description,omitempty"`
		Approved          *bool                `json:"approved,omitempty"`
		Status            common.PaymentAction `json:"status,omitempty"`
		ThreeDS           *ThreeDSEnrollment   `json:"3ds,omitempty"`
		Risk              *RiskAssessment      `json:"risk,omitempty"`
		Customer          *Customer            `json:"customer,omitempty"`
		BillingDescriptor *BillingDescriptor   `json:"billing_descriptor,omitempty"`
		Shipping          *Shipping            `json:"shipping,omitempty"`
		PaymentIP         string               `json:"payment_ip,omitempty"`
		Recipient         *Recipient           `json:"recipient,omitempty"`
		Metadata          map[string]string    `json:"metadata,omitempty"`
		ECI               string               `json:"eci,omitempty"`
		Actions           []ActionSummary      `json:"actions,omitempty"`
		SchemeID          string               `json:"scheme_id,omitempty"`
	}
	// SourceResponse ...
	SourceResponse struct {
		*CardSourceResponse
		*AlternativePaymentSourceResponse
	}
	// CardSourceResponse ...
	CardSourceResponse struct {
		ID                      string              `json:"id,omitempty"`
		Type                    string              `json:"type,omitempty"`
		BillingAddress          *common.Address     `json:"billing_address,omitempty"`
		Phone                   *common.Phone       `json:"phone,omitempty"`
		ExpiryMonth             uint64              `json:"expiry_month,omitempty"`
		ExpiryYear              uint64              `json:"expiry_year,omitempty"`
		Name                    string              `json:"name,omitempty"`
		Scheme                  string              `json:"scheme,omitempty"`
		Last4                   string              `json:"last4,omitempty"`
		Fingerprint             string              `json:"fingerprint,omitempty"`
		Bin                     string              `json:"bin,omitempty"`
		CardType                common.CardType     `json:"card_type,omitempty"`
		CardCategory            common.CardCategory `json:"card_category,omitempty"`
		Issuer                  string              `json:"issuer,omitempty"`
		IssuerCountry           string              `json:"issuer_country,omitempty"`
		ProductID               string              `json:"product_id,omitempty"`
		ProductType             string              `json:"product_type,omitempty"`
		AVSCheck                string              `json:"avs_check,omitempty"`
		CVVCheck                string              `json:"cvv_check,omitempty"`
		PaymentAccountReference string              `json:"payment_account_reference,omitempty"`
		Payouts                 *bool               `json:"payouts,omitempty"`
		FastFunds               string              `json:"fast_funds,omitempty"`
	}

	// AlternativePaymentSourceResponse ...
	AlternativePaymentSourceResponse struct {
		ID             string          `json:"id"`
		Type           string          `json:"type"`
		BillingAddress *common.Address `json:"billing_address,omitempty"`
		Phone          *common.Phone   `json:"phone,omitempty"`
	}

	// DestinationResponse -
	DestinationResponse struct {
		ID            string              `json:"id,omitempty"`
		Type          string              `json:"type,omitempty"`
		ExpiryMonth   uint64              `json:"expiry_month,omitempty"`
		ExpiryYear    uint64              `json:"expiry_year,omitempty"`
		Scheme        string              `json:"scheme,omitempty"`
		Last4         string              `json:"last4,omitempty"`
		Fingerprint   string              `json:"fingerprint,omitempty"`
		Bin           string              `json:"bin,omitempty"`
		CardType      common.CardType     `json:"card_type,omitempty"`
		CardCategory  common.CardCategory `json:"card_category,omitempty"`
		Issuer        string              `json:"issuer,omitempty"`
		IssuerCountry string              `json:"issuer_country,omitempty"`
		ProductID     string              `json:"product_id,omitempty"`
		ProductType   string              `json:"product_type,omitempty"`
	}
)

// MarshalJSON ...
func (s *SourceResponse) MarshalJSON() ([]byte, error) {
	if s.CardSourceResponse != nil {
		return json.Marshal(s.CardSourceResponse)
	} else if s.AlternativePaymentSourceResponse != nil {
		return json.Marshal(s.AlternativePaymentSourceResponse)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON ...
func (s *SourceResponse) UnmarshalJSON(data []byte) error {
	temp := &struct {
		Type string `json:"type"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	switch temp.Type {
	case "card":
		var source CardSourceResponse
		if err := json.Unmarshal(data, &source); err != nil {
			return err
		}
		s.CardSourceResponse = &source
		s.AlternativePaymentSourceResponse = nil
	default:
		var source AlternativePaymentSourceResponse
		if err := json.Unmarshal(data, &source); err != nil {
			return err
		}
		s.AlternativePaymentSourceResponse = &source
		s.CardSourceResponse = nil
	}
	return nil
}

type (
	// Accepted ...
	Accepted struct {
		ActionID  string                 `json:"action_id"`
		Reference string                 `json:"reference"`
		Links     map[string]common.Link `json:"_links"`
	}
)
