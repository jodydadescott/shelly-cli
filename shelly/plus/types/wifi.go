package types

import (
	"strings"

	"github.com/jinzhu/copier"
)

// WifiStatus status of the WiFi component contains information about the state of the WiFi connection of the device.
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/WiFi#status
type WifiStatus struct {
	// StaIP Ip of the device in the network (null if disconnected)
	StaIP *string `json:"sta_ip,omitempty" yaml:"sta_ip,omitempty"`
	// Status of the connection. Range of values: disconnected, connecting, connected, got ip
	Status string `json:"status" yaml:"status"`
	// Ssid of the network (null if disconnected)
	SSID *string `json:"ssid,omitempty" yaml:"ssid,omitempty"`
	// Rssi Strength of the signal in dBms
	RSSI int `json:"rssi,omitempty" yaml:"rssi,omitempty"`
	// ApClientCount Number of clients connected to the access point. Present only when AP is
	// enabled and range extender functionality is present and enabled.
	ApClientCount int `json:"ap_client_count,omitempty" yaml:"ap_client_count,omitempty"`
}

// Clone return copy
func (t *WifiStatus) Clone() *WifiStatus {
	c := &WifiStatus{}
	copier.Copy(&c, &t)
	return c
}

func (t *WifiStatus) GetStatus() WifiStatusStatus {
	return WifiStatusStatusFromString(t.Status)
}

type WifiStatusStatus string

const (
	WifiStatusStatusInvalid      WifiStatusStatus = "invalid"
	WifiStatusStatusDisconnected WifiStatusStatus = "disconnected"
	WifiStatusStatusConnecting   WifiStatusStatus = "connecting"
	WifiStatusStatusConnected    WifiStatusStatus = "connected"
	WifiStatusStatusGotIP        WifiStatusStatus = "got ip"
)

func WifiStatusStatusFromString(s string) WifiStatusStatus {

	switch strings.ToLower(s) {

	case string(WifiStatusStatusDisconnected):
		return WifiStatusStatusDisconnected

	case string(WifiStatusStatusConnecting):
		return WifiStatusStatusConnecting

	case string(WifiStatusStatusConnected):
		return WifiStatusStatusConnected

	case string(WifiStatusStatusGotIP):
		return WifiStatusStatusGotIP

	}

	return WifiStatusStatusInvalid
}

// WifiConfig configuration of the WiFi component contains information about the access point of the device,
// the network stations and the roaming settings.
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/WiFi#configuration
type WifiConfig struct {
	// Ap Information about the access point
	Ap *WifiAP `json:"ap" yaml:"ap"`
	// Sta information about the sta configuration
	Sta *WifiSTA `json:"sta" yaml:"sta"`
	// Sta1 information about the sta configuration
	Sta1 *WifiSTA `json:"sta1" yaml:"sta1"`
	// Roam WiFi roaming configuration
	Roam *WifiRoam `json:"roam" yaml:"roam"`
}

// Clone return copy
func (t *WifiConfig) Clone() *WifiConfig {
	c := &WifiConfig{}
	copier.Copy(&c, &t)
	return c
}

// WifiAP WiFi component object
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/WiFi#configuration
type WifiAP struct {
	// SSID readonly SSID of the access point
	SSID *string `json:"ssid" yaml:"ssid"`
	// Pass password for the ssid, writeonly. Must be provided if you provide ssid
	Pass *string `json:"pass" yaml:"pass"`
	// IsOpen True if the access point is open, false otherwise
	IsOpen bool `json:"is_open" yaml:"is_open"`
	// Enable true if the access point is enabled, false otherwise
	Enable bool `json:"enable" yaml:"enable"`
	// RangeExtender range extender configuration object, available only when range extender functionality is present.
	RangeExtender *WifiRangeExtender `json:"range_extender" yaml:"range_extender"`
}

// Clone return copy
func (t *WifiAP) Clone() *WifiAP {
	c := &WifiAP{}
	copier.Copy(&c, &t)
	return c
}

// WifiRangeExtender Range extender configuration object, available only when range extender functionality is present.
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/WiFi#configuration
type WifiRangeExtender struct {
	Enable bool `json:"enable" yaml:"enable"`
}

// Clone return copy
func (t *WifiRangeExtender) Clone() *WifiRangeExtender {
	c := &WifiRangeExtender{}
	copier.Copy(&c, &t)
	return c
}

// WifiSTA WiFi component object
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/WiFi#configuration
type WifiSTA struct {
	// SSID of the network
	SSID *string `json:"ssid" yaml:"ssid"`
	// Password for the ssid, writeonly. Must be provided if you provide ssid
	Pass *string `json:"pass" yaml:"pass"`
	// IsOpen true if the network is open, i.e. no password is set, false otherwise, readonly
	IsOpen bool `json:"is_open" yaml:"is_open"`
	// Enable True if the configuration is enabled, false otherwise
	Enable bool `json:"enable" yaml:"enable"`
	// Ipv4Mode IPv4 mode. Range of values: dhcp, static
	Ipv4Mode string `json:"ipv4mode" yaml:"ipv4mode"`
	// IP Ip to use when ipv4mode is static
	IP *string `json:"ip" yaml:"ip"`
	// Netmask to use when ipv4mode is static
	Netmask *string `json:"netmask" yaml:"netmask"`
	// Gateway to use when ipv4mode is static
	Gateway *string `json:"gw" yaml:"gw"`
	// Nameserver to use when ipv4mode is static
	Nameserver *string `json:"nameserver" yaml:"nameserver"`
}

// Clone return copy
func (t *WifiSTA) Clone() *WifiSTA {
	c := &WifiSTA{}
	copier.Copy(&c, &t)
	return c
}

// WifiRoam WiFi roaming configuration
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/WiFi#configuration
type WifiRoam struct {
	// RSSIThreshold - when reached will trigger the access point roaming. Default value: -80
	RSSIThreshold int `json:"rssi_thr" yaml:"rssi_thr"`
	// Interval at which to scan for better access points. Enabled if set to positive number,
	// disabled if set to 0. Default value: 60
	Interval int `json:"interval" yaml:"interval"`
}

// Clone return copy
func (t *WifiRoam) Clone() *WifiRoam {
	c := &WifiRoam{}
	copier.Copy(&c, &t)
	return c
}

// Scan WiFi component object
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/WiFi
type WifiNet struct {
	SSID    *string `json:"ssid,omitempty"`
	BSSID   string  `json:"bssid"`
	Auth    int     `json:"auth,omitempty"`
	Channel int     `json:"channel,omitempty"`
	RSSI    int     `json:"rssi,omitempty"`
}

// Clone return copy
func (t *WifiNet) Clone() *WifiNet {
	c := &WifiNet{}
	copier.Copy(&c, &t)
	return c
}

// WifiAPClient WiFi component object
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/WiFi
type WifiAPClient struct {
	MAC      string `json:"mac"`
	IP       string `json:"ip"`
	IPStatic bool   `json:"ip_static"`
	Mport    int    `json:"mport"`
	Since    int    `json:"since"`
}

// Clone return copy
func (t *WifiAPClient) Clone() *WifiAPClient {
	c := &WifiAPClient{}
	copier.Copy(&c, &t)
	return c
}

// WifiScanResults internal use only
type WifiScanResults struct {
	Results []WifiNet `json:"results,omitempty"`
}

// Clone return copy
func (t *WifiScanResults) Clone() *WifiScanResults {
	c := &WifiScanResults{}
	copier.Copy(&c, &t)
	return c
}

// WifiAPClients internal use only
type WifiAPClients struct {
	Ts      *int           `json:"ts,omitempty"`
	Clients []WifiAPClient `json:"ap_clients,omitempty"`
}

// Clone return copy
func (t *WifiAPClients) Clone() *WifiAPClients {
	c := &WifiAPClients{}
	copier.Copy(&c, &t)
	return c
}
