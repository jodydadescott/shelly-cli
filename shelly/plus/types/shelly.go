package types

import (
	"github.com/jinzhu/copier"
)

// ShellyStatus status of all the components of the device.
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Shelly
type ShellyStatus struct {
	Bluetooth *BluetoothStatus `json:"ble,omitempty" yaml:"ble,omitempty"`
	Cloud     *CloudStatus     `json:"cloud,omitempty" yaml:"cloud,omitempty"`
	Mqtt      *MqttStatus      `json:"mqtt,omitempty" yaml:"mqtt,omitempty"`
	Ethernet  *EthernetStatus  `json:"eth,omitempty" yaml:"eth,omitempty"`
	System    *SystemStatus    `json:"sys,omitempty" yaml:"sys,omitempty"`
	Wifi      *WifiStatus      `json:"wifi,omitempty" yaml:"wifi,omitempty"`
	Websocket *WebsocketStatus `json:"ws,omitempty" yaml:"ws,omitempty"`
	Light0    *LightStatus     `json:"light:0,omitempty" yaml:"light:0,omitempty"`
	Light1    *LightStatus     `json:"light:1,omitempty" yaml:"light:1,omitempty"`
	Light2    *LightStatus     `json:"light:2,omitempty" yaml:"light:2,omitempty"`
	Light3    *LightStatus     `json:"light:3,omitempty" yaml:"light:3,omitempty"`
	Light4    *LightStatus     `json:"light:4,omitempty" yaml:"light:4,omitempty"`
	Light5    *LightStatus     `json:"light:5,omitempty" yaml:"light:5,omitempty"`
	Light6    *LightStatus     `json:"light:6,omitempty" yaml:"light:6,omitempty"`
	Light7    *LightStatus     `json:"light:7,omitempty" yaml:"light:7,omitempty"`
	Input0    *InputStatus     `json:"input:0,omitempty" yaml:"input:0,omitempty"`
	Input1    *InputStatus     `json:"input:1,omitempty" yaml:"input:1,omitempty"`
	Input2    *InputStatus     `json:"input:2,omitempty" yaml:"input:2,omitempty"`
	Input3    *InputStatus     `json:"input:3,omitempty" yaml:"input:3,omitempty"`
	Input4    *InputStatus     `json:"input:4,omitempty" yaml:"input:4,omitempty"`
	Input5    *InputStatus     `json:"input:5,omitempty" yaml:"input:5,omitempty"`
	Input6    *InputStatus     `json:"input:6,omitempty" yaml:"input:6,omitempty"`
	Input7    *InputStatus     `json:"input:7,omitempty" yaml:"input:7,omitempty"`
	Switch0   *SwitchStatus    `json:"switch:0,omitempty" yaml:"switch:0,omitempty"`
	Switch1   *SwitchStatus    `json:"switch:1,omitempty" yaml:"switch:1,omitempty"`
	Switch2   *SwitchStatus    `json:"switch:2,omitempty" yaml:"switch:2,omitempty"`
	Switch3   *SwitchStatus    `json:"switch:3,omitempty" yaml:"switch:3,omitempty"`
	Switch4   *SwitchStatus    `json:"switch:4,omitempty" yaml:"switch:4,omitempty"`
	Switch5   *SwitchStatus    `json:"switch:5,omitempty" yaml:"switch:5,omitempty"`
	Switch6   *SwitchStatus    `json:"switch:6,omitempty" yaml:"switch:6,omitempty"`
	Switch7   *SwitchStatus    `json:"switch:7,omitempty" yaml:"switch:7,omitempty"`
}

// ShellyRPCMethods lists of all available RPC methods. It takes into account both ACL and authentication
// restrictions and only lists the methods allowed for the particular user/channel that's making the request.
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Shelly#shellylistmethods
type ShellyRPCMethods struct {
	// Methods names of the methods allowed
	Methods []string `json:"methods" yaml:"methods"`
}

// Clone return copy
func (t *ShellyRPCMethods) Clone() *ShellyRPCMethods {
	c := &ShellyRPCMethods{}
	copier.Copy(&c, &t)
	return c
}

// ShellyConfig Shelly component config. The config is composed of each components config.
// Shelly devices can have zero or more 'Light', 'Input' and 'Switch' types. Because these
// are explicity named and not members of a JSON array we have statically created them.
// This seemed to be a cleaner solution then a customized JSON/YAML encoder/decoder. We have
// created 8 for each which is currently more then enough as the max for any Shelly product as
// of today is 4.
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Shelly#configuration
type ShellyConfig struct {
	Auth      *AuthConfig      `json:"auth,omitempty" yaml:"auth,omitempty"`
	Bluetooth *BluetoothConfig `json:"ble,omitempty" yaml:"ble,omitempty"`
	Cloud     *CloudConfig     `json:"cloud,omitempty" yaml:"cloud,omitempty"`
	Mqtt      *MqttConfig      `json:"mqtt,omitempty" yaml:"mqtt,omitempty"`
	Ethernet  *EthernetConfig  `json:"eth,omitempty" yaml:"eth,omitempty"`
	System    *SystemConfig    `json:"sys,omitempty" yaml:"sys,omitempty"`
	Wifi      *WifiConfig      `json:"wifi,omitempty" yaml:"wifi,omitempty"`
	Websocket *WebsocketConfig `json:"ws,omitempty" yaml:"ws,omitempty"`
	Light0    *LightConfig     `json:"light:0,omitempty" yaml:"light:0,omitempty"`
	Light1    *LightConfig     `json:"light:1,omitempty" yaml:"light:1,omitempty"`
	Light2    *LightConfig     `json:"light:2,omitempty" yaml:"light:2,omitempty"`
	Light3    *LightConfig     `json:"light:3,omitempty" yaml:"light:3,omitempty"`
	Light4    *LightConfig     `json:"light:4,omitempty" yaml:"light:4,omitempty"`
	Light5    *LightConfig     `json:"light:5,omitempty" yaml:"light:5,omitempty"`
	Light6    *LightConfig     `json:"light:6,omitempty" yaml:"light:6,omitempty"`
	Light7    *LightConfig     `json:"light:7,omitempty" yaml:"light:7,omitempty"`
	Input0    *InputConfig     `json:"input:0,omitempty" yaml:"input:0,omitempty"`
	Input1    *InputConfig     `json:"input:1,omitempty" yaml:"input:1,omitempty"`
	Input2    *InputConfig     `json:"input:2,omitempty" yaml:"input:2,omitempty"`
	Input3    *InputConfig     `json:"input:3,omitempty" yaml:"input:3,omitempty"`
	Input4    *InputConfig     `json:"input:4,omitempty" yaml:"input:4,omitempty"`
	Input5    *InputConfig     `json:"input:5,omitempty" yaml:"input:5,omitempty"`
	Input6    *InputConfig     `json:"input:6,omitempty" yaml:"input:6,omitempty"`
	Input7    *InputConfig     `json:"input:7,omitempty" yaml:"input:7,omitempty"`
	Switch0   *SwitchConfig    `json:"switch:0,omitempty" yaml:"switch:0,omitempty"`
	Switch1   *SwitchConfig    `json:"switch:1,omitempty" yaml:"switch:1,omitempty"`
	Switch2   *SwitchConfig    `json:"switch:2,omitempty" yaml:"switch:2,omitempty"`
	Switch3   *SwitchConfig    `json:"switch:3,omitempty" yaml:"switch:3,omitempty"`
	Switch4   *SwitchConfig    `json:"switch:4,omitempty" yaml:"switch:4,omitempty"`
	Switch5   *SwitchConfig    `json:"switch:5,omitempty" yaml:"switch:5,omitempty"`
	Switch6   *SwitchConfig    `json:"switch:6,omitempty" yaml:"switch:6,omitempty"`
	Switch7   *SwitchConfig    `json:"switch:7,omitempty" yaml:"switch:7,omitempty"`
}

// Clone return copy
func (t *ShellyConfig) Clone() *ShellyConfig {
	c := &ShellyConfig{}
	copier.Copy(&c, &t)
	return c
}

// DeviceInfo Shelly component top level device info
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Shelly#shellygetdeviceinfo
type DeviceInfo struct {
	Name *string `json:"name" yaml:"name"`
	// ID Id of the device
	ID string `json:"id" yaml:"id"`
	// MAC address of the device
	MAC string `json:"mac" yaml:"mac"`
	// Model of the device
	Model string `json:"model" yaml:"model"`
	// Generation of the device
	Generation float32 `json:"gen" yaml:"gen"`
	// FirmwareID Id of the firmware of the device
	FirmwareID string `json:"fw_id" yaml:"fw_id"`
	// Version of the firmware of the device
	Version string `json:"ver" yaml:"ver"`
	// App name
	App string `json:"app" yaml:"app"`
	// Profile name of the device profile (only applicable for multi-profile devices)
	Profile *string `json:"profile" yaml:"profile"`
	// AuthEnabled true if authentication is enabled, false otherwise
	AuthEnabled bool `json:"auth_en" yaml:"auth_en"`
	// AuthDomain name of the domain (null if authentication is not enabled)
	AuthDomain *string `json:"auth_domain" yaml:"auth_domain"`
	// Discoverable present only when false. If true, device is shown in 'Discovered devices'. If false, the device is hidden.
	Discoverable bool `json:"discoverable" yaml:"discoverable"`
	// Key cloud key of the device (see note below), present only when the ident parameter is set to true
	Key string `json:"key" yaml:"key"`
	// Batch used to provision the device, present only when the ident parameter is set to true
	Batch string `json:"batch" yaml:"batch"`
	// FwSbits Shelly internal flags, present only when the ident parameter is set to true
	FwSbits string `json:"fw_sbits" yaml:"fw_sbits"`
}

// Clone return copy
func (t *DeviceInfo) Clone() *DeviceInfo {
	c := &DeviceInfo{}
	copier.Copy(&c, &t)
	return c
}

// AuthConfig holds the user and clear text password. It is not part of the official Shelly Config API but it should be convenient
type AuthConfig struct {
	User string  `json:"user" yaml:"user"`
	Pass *string `json:"pass" yaml:"pass"`
}

// Clone return copy
func (t *AuthConfig) Clone() *AuthConfig {
	c := &AuthConfig{}
	copier.Copy(&c, &t)
	return c
}

// ShellyParams Shelly config parameters
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Shelly#configuration
type ShellyParams struct {
	// Stage is used by the following methods:
	// Update : The type of the new version - either stable or beta. By default updates to stable version. Optional
	Stage *string `json:"stage" yaml:"stage"`
	// Url is used by the following methods:
	// Update : Url address of the update. Optional
	Url *string `json:"url" yaml:"url"`
	// User is used by the following methods:
	// SetAuth: Must be set to admin. Only one user is supported. Required
	User *string `json:"user" yaml:"user"`
	// Realm is used by the following methods:
	// SetAuth : Must be the id of the device. Only one realm is supported. Required
	Realm *string `json:"realm" yaml:"realm"`
	// Ha1 is used by the following methods:
	// SetAuth : "user:realm:password" encoded in SHA256 (null to disable authentication). Required
	Ha1 *string `json:"ha1" yaml:"ha1"`
	// Data is used by the following methods:
	// PutUserCA : Contents of the PEM file (null if you want to delete the existing data). Required
	// PutTLSClientCert : Contents of the client.crt file (null if you want to delete the existing data). Required
	// PutTLSClientKey : Contents of the client.key file (null if you want to delete the existing data). Required
	Data *string `json:"data" yaml:"data"`
	// Append is used by the following methods:
	// PutUserCA : true if more data will be appended afterwards, default false.
	// PutTLSClientCert : true if more data will be appended afterwards, default false
	// PutTLSClientKey : true if more data will be appended afterwards, default false
	Append *bool `json:"append" yaml:"append"`
}

// Clone return copy
func (t *ShellyParams) Clone() *ShellyParams {
	c := &ShellyParams{}
	copier.Copy(&c, &t)
	return c
}

// ShellyReport is the report returned by Shelly.SetConfig
type ShellyReport struct {
	Bluetooth       *SetReport `json:"ble,omitempty" yaml:"ble,omitempty"`
	Cloud           *SetReport `json:"cloud,omitempty" yaml:"cloud,omitempty"`
	Mqtt            *SetReport `json:"mqtt,omitempty" yaml:"mqtt,omitempty"`
	Ethernet        *SetReport `json:"eth,omitempty" yaml:"eth,omitempty"`
	System          *SetReport `json:"sys,omitempty" yaml:"sys,omitempty"`
	Wifi            *SetReport `json:"wifi,omitempty" yaml:"wifi,omitempty"`
	Websocket       *SetReport `json:"ws,omitempty" yaml:"ws,omitempty"`
	Light0          *SetReport `json:"light:0,omitempty" yaml:"light:0,omitempty"`
	Light1          *SetReport `json:"light:1,omitempty" yaml:"light:1,omitempty"`
	Light2          *SetReport `json:"light:2,omitempty" yaml:"light:2,omitempty"`
	Light3          *SetReport `json:"light:3,omitempty" yaml:"light:3,omitempty"`
	Light4          *SetReport `json:"light:4,omitempty" yaml:"light:4,omitempty"`
	Light5          *SetReport `json:"light:5,omitempty" yaml:"light:5,omitempty"`
	Light6          *SetReport `json:"light:6,omitempty" yaml:"light:6,omitempty"`
	Light7          *SetReport `json:"light:7,omitempty" yaml:"light:7,omitempty"`
	Input0          *SetReport `json:"input:0,omitempty" yaml:"input:0,omitempty"`
	Input1          *SetReport `json:"input:1,omitempty" yaml:"input:1,omitempty"`
	Input2          *SetReport `json:"input:2,omitempty" yaml:"input:2,omitempty"`
	Input3          *SetReport `json:"input:3,omitempty" yaml:"input:3,omitempty"`
	Input4          *SetReport `json:"input:4,omitempty" yaml:"input:4,omitempty"`
	Input5          *SetReport `json:"input:5,omitempty" yaml:"input:5,omitempty"`
	Input6          *SetReport `json:"input:6,omitempty" yaml:"input:6,omitempty"`
	Input7          *SetReport `json:"input:7,omitempty" yaml:"input:7,omitempty"`
	Switch0         *SetReport `json:"switch:0,omitempty" yaml:"switch:0,omitempty"`
	Switch1         *SetReport `json:"switch:1,omitempty" yaml:"switch:1,omitempty"`
	Switch2         *SetReport `json:"switch:2,omitempty" yaml:"switch:2,omitempty"`
	Switch3         *SetReport `json:"switch:3,omitempty" yaml:"switch:3,omitempty"`
	Switch4         *SetReport `json:"switch:4,omitempty" yaml:"switch:4,omitempty"`
	Switch5         *SetReport `json:"switch:5,omitempty" yaml:"switch:5,omitempty"`
	Switch6         *SetReport `json:"switch:6,omitempty" yaml:"switch:6,omitempty"`
	Switch7         *SetReport `json:"switch:7,omitempty" yaml:"switch:7,omitempty"`
	RestartRequired bool       `json:"restart_required" yaml:"restart_required"`
}

// Clone return copy
func (t *ShellyReport) Clone() *ShellyReport {
	c := &ShellyReport{}
	copier.Copy(&c, &t)
	return c
}

// UpdatesReport checks for new firmware version for the device and returns information about it.
// If no update is available returns empty JSON object as result.
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Shelly#shellycheckforupdate
type UpdatesReport struct {
	Src              string                  `json:"src,omitempty" yaml:"src"`
	AvailableUpdates *SystemAvailableUpdates `json:"available_updates,omitempty" yaml:"available_updates,omitempty"`
}

// Clone return copy
func (t *UpdatesReport) Clone() *UpdatesReport {
	c := &UpdatesReport{}
	copier.Copy(&c, &t)
	return c
}
