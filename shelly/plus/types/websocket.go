package types

import (
	"github.com/jinzhu/copier"
)

// WebsocketStatus status
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Ws#status
type WebsocketStatus struct {
	// Connected true if device is connected to a websocket outbound connection or false otherwise.
	Connected bool `json:"connected" yaml:"connected"`
}

// Clone return copy
func (t *WebsocketStatus) Clone() *WebsocketStatus {
	c := &WebsocketStatus{}
	copier.Copy(&c, &t)
	return c
}

// WebsocketConfig configuration
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Ws#configuration
type WebsocketConfig struct {
	// Enable true if websocket outbound connection is enabled, false otherwise
	Enable bool `json:"enable" yaml:"enable"`
	// Server name of the server to which the device is connected. When prefixed with wss:// a TLS socket will be used
	Server string `json:"server" yaml:"server"`
	// SslCa type of the TCP sockets
	SslCa *string `json:"ssl_ca" yaml:"ssl_ca"`
}

// Clone return copy
func (t *WebsocketConfig) Clone() *WebsocketConfig {
	c := &WebsocketConfig{}
	copier.Copy(&c, &t)
	return c
}
