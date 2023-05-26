package types

import (
	"github.com/jinzhu/copier"
)

// LightStatus status of the Light component contains information about the brightness level and output state of the light instance.
// To obtain the status of the Light component its id must be specified.
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Light#status
type LightStatus struct {
	// ID Id of the Switch component instance
	ID int `json:"id" yaml:"id"`
	// Source of the last command, for example: init, WS_in, http, ...
	Source string `json:"source" yaml:"source"`
	// Output true if the output channel is currently on, false otherwise
	Output bool `json:"output,omitempty" yaml:"output"`
	// Brightness current brightness level (in percent)
	Brightness *float64 `json:"brightness,omitempty" yaml:"brightness"`
	// TimerStartedAt Unix timestamp, start time of the timer (in UTC) (shown if the timer is triggered)
	TimerStartedAt *float64 `json:"timer_started_at,omitempty" yaml:"timer_started_at,omitempty"`
	// TimerDuration duration of the timer in seconds (shown if the timer is triggered)
	TimerDuration *float64 `json:"timer_duration,omitempty" yaml:"timer_duration,omitempty"`
}

// Clone return copy
func (t *LightStatus) Clone() *LightStatus {
	c := &LightStatus{}
	copier.Copy(&c, &t)
	return c
}

type LightConfig struct {
	// ID Id of the Switch component instance
	ID int `json:"id" yaml:"id"`
	// Name of the switch instance
	Name *string `json:"name" yaml:"name"`
	// InitialState range of values: off, on, restore_last, match_input
	InitialState string `json:"initial_state" yaml:"initial_state"`
	// AutoOn True if the "Automatic ON" function is enabled, false otherwise
	AutoOn bool `json:"auto_on" yaml:"auto_on"`
	// AutoOnDelay Seconds to pass until the component is switched back on
	AutoOnDelay float64 `json:"auto_on_delay" yaml:"auto_on_delay"`
	// AutoOff True if the "Automatic OFF" function is enabled, false otherwise
	AutoOff bool `json:"auto_off" yaml:"auto_off"`
	// AutoOffDelay Seconds to pass until the component is switched back off
	AutoOffDelay float64 `json:"auto_off_delay" yaml:"auto_off_delay"`
	// DefaultBrightness brightness level (in percent) after power on
	DefaultBrightness float64 `json:"default.brightness" yaml:"default.brightness"`
	// NightModeEnable Enable or disable night mode
	NightModeEnable bool `json:"night_mode.enable" yaml:"night_mode.enable"`
	// NightModeBrightness brightness level limit when night mode is active
	NightModeBrightness float64 `json:"night_mode.brightness" yaml:"night_mode.brightness"`
	// NightModeActiveBetween containing 2 elements of type string, the first element indicates the start of
	// the period during which the night mode will be active, the second indicates the end of that period.
	// Both start and end are strings in the format HH:MM, where HH and MM are hours and minutes with optinal
	// leading zeros
	NightModeActiveBetween []string `json:"night_mode.active_between" yaml:"night_mode.active_between"`
}

// Clone return copy
func (t *LightConfig) Clone() *LightConfig {
	c := &LightConfig{}
	copier.Copy(&c, &t)
	return c
}

// LightParams ...
type LightParams struct {
	ID         int          `json:"id" yaml:"id"`
	Config     *LightConfig `json:"config,omitempty" yaml:"on,omitempty"`
	On         *bool        `json:"on,omitempty" yaml:"on,omitempty"`
	Brightness *float64     `json:"brightness,omitempty" yaml:"brightness,omitempty"`
}

// Clone return copy
func (t *LightParams) Clone() *LightParams {
	c := &LightParams{}
	copier.Copy(&c, &t)
	return c
}

type LightReport struct {
	Src string `json:"src,omitempty" yaml:"src"`
}

// Clone return copy
func (t *LightReport) Clone() *LightReport {
	c := &LightReport{}
	copier.Copy(&c, &t)
	return c
}
