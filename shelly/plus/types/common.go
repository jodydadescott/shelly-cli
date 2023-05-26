package types

import (
	"github.com/jinzhu/copier"
)

// FirmwareStatus is common for components Sys and Shelly
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Sys#status &
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Shelly#status
type FirmwareStatus struct {
	// Version of the new firmware
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
	// BuildID Id of the new build
	BuildID string `json:"build_id,omitempty" yaml:"build_id,omitempty"`
}

type SetReport struct {
	Src             string `json:"src,omitempty" yaml:"src"`
	RestartRequired bool   `json:"restart_required" yaml:"restart_required"`
}

// Request generic request
type Request struct {
	ID     int           `json:"id"`
	Method string        `json:"method,omitempty"`
	Params interface{}   `json:"params,omitempty"`
	Auth   *AuthResponse `json:"auth,omitempty"`
}

// Clone return copy
func (t *Request) Clone() *Request {
	c := &Request{}
	copier.Copy(&c, &t)
	return c
}

// Response generic response
type Response struct {
	ID    int    `json:"id" yaml:"id"`
	Src   string `json:"src" yaml:"src"`
	Error *Error `json:"error,omitempty"`
}

// Clone return copy
func (t *Response) Clone() *Response {
	c := &Response{}
	copier.Copy(&c, &t)
	return c
}
