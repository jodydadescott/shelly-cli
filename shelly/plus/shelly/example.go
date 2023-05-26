package shelly

import (
	"github.com/jodydadescott/shelly-manager/shelly/plus/bluetooth"
	"github.com/jodydadescott/shelly-manager/shelly/plus/cloud"
	"github.com/jodydadescott/shelly-manager/shelly/plus/mqtt"
	"github.com/jodydadescott/shelly-manager/shelly/plus/system"
	"github.com/jodydadescott/shelly-manager/shelly/plus/wifi"
)

// Auth      *AuthConfig      `json:"auth,omitempty" yaml:"auth,omitempty"`
// Bluetooth *BluetoothConfig `json:"ble,omitempty" yaml:"ble,omitempty"`
// Cloud     *CloudConfig     `json:"cloud,omitempty" yaml:"cloud,omitempty"`
// Mqtt      *MqttConfig      `json:"mqtt,omitempty" yaml:"mqtt,omitempty"`
// Ethernet  *EthernetConfig  `json:"eth,omitempty" yaml:"eth,omitempty"`
// System    *SystemConfig    `json:"sys,omitempty" yaml:"sys,omitempty"`
// Wifi      *WifiConfig      `json:"wifi,omitempty" yaml:"wifi,omitempty"`
// Websocket *WebsocketConfig `json:"ws,omitempty" yaml:"ws,omitempty"`
// Light0    *LightConfig     `json:"light:0,omitempty" yaml:"light:0,omitempty"`
// Light1    *LightConfig     `json:"light:1,omitempty" yaml:"light:1,omitempty"`
// Light2    *LightConfig     `json:"light:2,omitempty" yaml:"light:2,omitempty"`
// Light3    *LightConfig     `json:"light:3,omitempty" yaml:"light:3,omitempty"`
// Light4    *LightConfig     `json:"light:4,omitempty" yaml:"light:4,omitempty"`
// Light5    *LightConfig     `json:"light:5,omitempty" yaml:"light:5,omitempty"`
// Light6    *LightConfig     `json:"light:6,omitempty" yaml:"light:6,omitempty"`
// Light7    *LightConfig     `json:"light:7,omitempty" yaml:"light:7,omitempty"`
// Input0    *InputConfig     `json:"input:0,omitempty" yaml:"input:0,omitempty"`
// Input1    *InputConfig     `json:"input:1,omitempty" yaml:"input:1,omitempty"`
// Input2    *InputConfig     `json:"input:2,omitempty" yaml:"input:2,omitempty"`
// Input3    *InputConfig     `json:"input:3,omitempty" yaml:"input:3,omitempty"`
// Input4    *InputConfig     `json:"input:4,omitempty" yaml:"input:4,omitempty"`
// Input5    *InputConfig     `json:"input:5,omitempty" yaml:"input:5,omitempty"`
// Input6    *InputConfig     `json:"input:6,omitempty" yaml:"input:6,omitempty"`
// Input7    *InputConfig     `json:"input:7,omitempty" yaml:"input:7,omitempty"`
// Switch0   *SwitchConfig    `json:"switch:0,omitempty" yaml:"switch:0,omitempty"`
// Switch1   *SwitchConfig    `json:"switch:1,omitempty" yaml:"switch:1,omitempty"`
// Switch2   *SwitchConfig    `json:"switch:2,omitempty" yaml:"switch:2,omitempty"`
// Switch3   *SwitchConfig    `json:"switch:3,omitempty" yaml:"switch:3,omitempty"`
// Switch4   *SwitchConfig    `json:"switch:4,omitempty" yaml:"switch:4,omitempty"`
// Switch5   *SwitchConfig    `json:"switch:5,omitempty" yaml:"switch:5,omitempty"`
// Switch6   *SwitchConfig    `json:"switch:6,omitempty" yaml:"switch:6,omitempty"`
// Switch7   *SwitchConfig    `json:"switch:7,omitempty" yaml:"switch:7,omitempty"`

func ExampleConfig() *ShellyConfig {
	return &ShellyConfig{
		Bluetooth: bluetooth.ExampleConfig(),
		Cloud:     cloud.ExampleConfig(),
		Mqtt:      mqtt.ExampleConfig(),
		System:    system.ExampleConfig(),
		Wifi:      wifi.ExampleConfig(),
	}
}
