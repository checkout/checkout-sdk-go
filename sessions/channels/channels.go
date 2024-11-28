package channels

import (
	"github.com/checkout/checkout-sdk-go/common"
)

type ChannelType string

const (
	App               ChannelType = "app"
	Browser           ChannelType = "browser"
	MerchantInitiated ChannelType = "merchant_initiated"
)

type SdkInterfaceType string

const (
	Both   SdkInterfaceType = "both"
	Html   SdkInterfaceType = "html"
	Native SdkInterfaceType = "native"
)

type UIElements string

const (
	HtmlOther    UIElements = "html_other"
	MultiSelect  UIElements = "multi_select"
	Oob          UIElements = "oob"
	SingleSelect UIElements = "single_select"
	Text         UIElements = "text"
)

type RequestType string

const (
	AccountVerification     RequestType = "account_verification"
	AddCard                 RequestType = "add_card"
	InstallmentTransaction  RequestType = "installment_transaction"
	MailOrder               RequestType = "mail_order"
	MaintainCardInformation RequestType = "maintain_card_information"
	OtherPayment            RequestType = "other_payment"
	RecurringTransaction    RequestType = "recurring_transaction"
	SplitOrDelayedShipment  RequestType = "split_or_delayed_shipment"
	TelephoneOrder          RequestType = "telephone_order"
	TopUp                   RequestType = "top_up"
	WhitelistStatusCheck    RequestType = "whitelist_status_check"
)

type (
	Channel interface {
		GetType() ChannelType
	}

	ChannelData struct {
		Channel ChannelType `json:"channel,omitempty"`
	}

	appSession struct {
		ChannelData
		SdkAppId           string                 `json:"sdk_app_id,omitempty"`
		SdkMaxTimeout      int                    `json:"sdk_max_timeout,omitempty"`
		SdkEphemPubKey     *SdkEphemeralPublicKey `json:"sdk_ephem_pub_key,omitempty"`
		SdkReferenceNumber string                 `json:"sdk_reference_number,omitempty"`
		SdkEncryptedData   string                 `json:"sdk_encrypted_data,omitempty"`
		SdkTransactionId   string                 `json:"sdk_transaction_id,omitempty"`
		SdkInterfaceType   SdkInterfaceType       `json:"sdk_interface_type,omitempty"`
		SdkUiElements      []UIElements           `json:"sdk_ui_elements,omitempty"`
	}

	browserSession struct {
		ChannelData
		ThreeDsMethodCompletion common.ThreeDsMethodCompletion `json:"three_ds_method_completion,omitempty" default:"u"`
		AcceptHeader            string                         `json:"accept_header,omitempty"`
		JavaEnabled             bool                           `json:"java_enabled,omitempty"`
		JavascriptEnabled       bool                           `json:"javascript_enabled,omitempty"`
		Language                string                         `json:"language,omitempty"`
		ColorDepth              string                         `json:"color_depth,omitempty"`
		ScreenHeight            string                         `json:"screen_height,omitempty"`
		ScreenWidth             string                         `json:"screen_width,omitempty"`
		Timezone                string                         `json:"timezone,omitempty"`
		UserAgent               string                         `json:"user_agent,omitempty"`
		IpAddress               string                         `json:"ip_address,omitempty"`
	}

	merchantInitiatedSession struct {
		ChannelData
		RequestType RequestType `json:"request_type,omitempty"`
	}
)

func NewAppSession() *appSession {
	return &appSession{ChannelData: ChannelData{Channel: App}}
}

func NewBrowserSession() *browserSession {
	return &browserSession{ChannelData: ChannelData{Channel: Browser}}
}

func NewMerchantInitiatedSession() *merchantInitiatedSession {
	return &merchantInitiatedSession{ChannelData: ChannelData{Channel: MerchantInitiated}}
}

func (s *appSession) GetType() ChannelType {
	return s.Channel
}

func (s *browserSession) GetType() ChannelType {
	return s.Channel
}

func (s *merchantInitiatedSession) GetType() ChannelType {
	return s.Channel
}

type SdkEphemeralPublicKey struct {
	Kty string `json:"kty,omitempty"`
	Crv string `json:"crv,omitempty"`
	X   string `json:"x,omitempty"`
	Y   string `json:"y,omitempty"`
}
