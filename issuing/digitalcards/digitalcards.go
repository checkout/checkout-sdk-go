package issuing

import (
	"time"

	"github.com/checkout/checkout-sdk-go/v2/common"
)

type DigitalCardStatus string

const (
	DigitalCardInactive DigitalCardStatus = "inactive"
	DigitalCardActive   DigitalCardStatus = "active"
	DigitalCardDeleted  DigitalCardStatus = "deleted"
)

type DigitalCardType string

const (
	SecureElement    DigitalCardType = "secure_element"
	HostCardEmulation DigitalCardType = "host_card_emulation"
	CardOnFile       DigitalCardType = "card_on_file"
	ECommerce        DigitalCardType = "e_commerce"
	QrCode           DigitalCardType = "qr_code"
)

type DigitalCardDeviceType string

const (
	DeviceSamsungPhone        DigitalCardDeviceType = "samsung_phone"
	DeviceSamsungTablet       DigitalCardDeviceType = "samsung_tablet"
	DeviceSamsungWatch        DigitalCardDeviceType = "samsung_watch"
	DeviceSamsungTv           DigitalCardDeviceType = "samsung_tv"
	DeviceIphone              DigitalCardDeviceType = "iphone"
	DeviceIwatch              DigitalCardDeviceType = "iwatch"
	DeviceIpad                DigitalCardDeviceType = "ipad"
	DeviceMacBook             DigitalCardDeviceType = "mac_book"
	DeviceAndroidPhone        DigitalCardDeviceType = "android_phone"
	DeviceAndroidTablet       DigitalCardDeviceType = "android_tablet"
	DeviceAndroidWatch        DigitalCardDeviceType = "android_watch"
	DeviceMobilePhone         DigitalCardDeviceType = "mobile_phone"
	DeviceTablet              DigitalCardDeviceType = "tablet"
	DeviceWatch               DigitalCardDeviceType = "watch"
	DeviceMobilePhoneOrTablet DigitalCardDeviceType = "mobile_phone_or_tablet"
	DeviceBracelet            DigitalCardDeviceType = "bracelet"
	DeviceHostCardEmulation   DigitalCardDeviceType = "host_card_emulation"
	DeviceUnknown             DigitalCardDeviceType = "unknown"
)

type DigitalCardDeviceNetworkType string

const (
	NetworkCellular DigitalCardDeviceNetworkType = "cellular"
	NetworkWifi     DigitalCardDeviceNetworkType = "wifi"
)

type DigitalCardDeviceTimeZoneSetting string

const (
	TimeZoneNetworkSet  DigitalCardDeviceTimeZoneSetting = "network_set"
	TimeZoneConsumerSet DigitalCardDeviceTimeZoneSetting = "consumer_set"
)

type (
	DigitalCardRequestor struct {
		Id   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	}

	DigitalCardDevice struct {
		Id               string                           `json:"id,omitempty"`
		Type             DigitalCardDeviceType            `json:"type,omitempty"`
		Manufacturer     string                           `json:"manufacturer,omitempty"`
		Brand            string                           `json:"brand,omitempty"`
		Model            string                           `json:"model,omitempty"`
		OsVersion        string                           `json:"os_version,omitempty"`
		FirmwareVersion  string                           `json:"firmware_version,omitempty"`
		PhoneNumber      string                           `json:"phone_number,omitempty"`
		DeviceName       string                           `json:"device_name,omitempty"`
		DeviceParentId   string                           `json:"device_parent_id,omitempty"`
		Language         string                           `json:"language,omitempty"`
		SerialNumber     string                           `json:"serial_number,omitempty"`
		TimeZone         string                           `json:"time_zone,omitempty"`
		TimeZoneSetting  DigitalCardDeviceTimeZoneSetting `json:"time_zone_setting,omitempty"`
		SimSerialNumber  string                           `json:"sim_serial_number,omitempty"`
		Imei             string                           `json:"imei,omitempty"`
		NetworkOperator  string                           `json:"network_operator,omitempty"`
		NetworkType      DigitalCardDeviceNetworkType     `json:"network_type,omitempty"`
	}

	GetDigitalCardResponse struct {
		HttpMetadata  common.HttpMetadata
		Id            string                 `json:"id,omitempty"`
		CardId        string                 `json:"card_id,omitempty"`
		ClientId      string                 `json:"client_id,omitempty"`
		EntityId      string                 `json:"entity_id,omitempty"`
		LastFour      string                 `json:"last_four,omitempty"`
		Status        DigitalCardStatus      `json:"status,omitempty"`
		Type          DigitalCardType        `json:"type,omitempty"`
		SchemeCardId  string                 `json:"scheme_card_id,omitempty"`
		Requestor     *DigitalCardRequestor  `json:"requestor,omitempty"`
		Device        *DigitalCardDevice     `json:"device,omitempty"`
		ProvisionedOn *time.Time             `json:"provisioned_on,omitempty"`
		Links         map[string]common.Link `json:"_links,omitempty"`
	}
)
