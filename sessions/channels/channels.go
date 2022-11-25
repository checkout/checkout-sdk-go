package channels

import (
	"github.com/checkout/checkout-sdk-go-beta/common"
)

type ChannelType string

const (
	Browser ChannelType = "browser"
	App     ChannelType = "app"
)

type SdkInterfaceType string

const (
	Native SdkInterfaceType = "native"
	Html   SdkInterfaceType = "html"
	Both   SdkInterfaceType = "both"
)

type UIElements string

const (
	Text         UIElements = "text"
	SingleSelect UIElements = "single_select"
	MultiSelect  UIElements = "multi_select"
	Oob          UIElements = "oob"
	HtmlOther    UIElements = "html_other"
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
		ThreeDsMethodCompletion common.ThreeDsMethodCompletion `json:"three_ds_method_completion,omitempty"`
		AcceptHeader            string                         `json:"accept_header,omitempty"`
		JavaEnabled             bool                           `json:"java_enabled,omitempty"`
		Language                string                         `json:"language,omitempty"`
		ColorDepth              string                         `json:"color_depth,omitempty"`
		ScreenHeight            string                         `json:"screen_height,omitempty"`
		ScreenWidth             string                         `json:"screen_width,omitempty"`
		Timezone                string                         `json:"timezone,omitempty"`
		UserAgent               string                         `json:"user_agent,omitempty"`
		IpAddress               string                         `json:"ip_address,omitempty"`
	}
)

func NewAppSession() *appSession {
	return &appSession{ChannelData: ChannelData{Channel: App}}
}

func NewBrowserSession() *browserSession {
	return &browserSession{ChannelData: ChannelData{Channel: Browser}}
}

func (s *appSession) GetType() ChannelType {
	return s.Channel
}

func (s *browserSession) GetType() ChannelType {
	return s.Channel
}

type SdkEphemeralPublicKey struct {
	Kty string `json:"kty,omitempty"`
	Crv string `json:"crv,omitempty"`
	X   string `json:"x,omitempty"`
	Y   string `json:"y,omitempty"`
}
