package types

import (
	"github.com/jinzhu/copier"
)

// EthernetStatus Ethernet component top level status
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Eth#status
type EthernetStatus struct {
	// IP of the device in the network
	IP *string `json:"ip,omitempty" yaml:"ip,omitempty"`
}

// Clone return copy
func (t *EthernetStatus) Clone() *EthernetStatus {
	c := &EthernetStatus{}
	copier.Copy(&c, &t)
	return c
}

// EthernetConfig Ethernet component top level config
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Eth#configuration
type EthernetConfig struct {
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
func (t *EthernetConfig) Clone() *EthernetConfig {
	c := &EthernetConfig{}
	copier.Copy(&c, &t)
	return c
}
