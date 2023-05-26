package types

import (
	"github.com/jinzhu/copier"
)

// SwitchStatus status of the Switch component contains information about the temperature, voltage, energy level and
// other physical characteristics of the switch instance. To obtain the status of the Switch component its id must be specified.
// For switches with power metering capabilities the status payload contains an additional set of properties with information
// about instantaneous power, supply voltage parameters and energy counters.
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Switch#status
type SwitchStatus struct {
	// ID Id of the Switch component instance
	ID int `json:"id" yaml:"id"`
	// Source of the last command, for example: init, WS_in, http, ...
	Source string `json:"source" yaml:"source"`
	// Output true if the output channel is currently on, false otherwise
	Output bool `json:"output,omitempty" yaml:"output"`
	// TimerStartedAt Unix timestamp, start time of the timer (in UTC) (shown if the timer is triggered)
	TimerStartedAt *float64 `json:"timer_started_at,omitempty" yaml:"timer_started_at,omitempty"`
	// TimerDuration duration of the timer in seconds (shown if the timer is triggered)
	TimerDuration *float64 `json:"timer_duration,omitempty" yaml:"timer_duration,omitempty"`
	// Apower last measured instantaneous active power (in Watts) delivered to the attached load (shown if applicable)
	Apower *float64 `json:"apower,omitempty" yaml:"apower,omitempty"`
	// Voltage last measured voltage in Volts (shown if applicable)
	Voltage *float64 `json:"voltage,omitempty" yaml:"voltage,omitempty"`
	// Current last measured current in Amperes (shown if applicable)
	Current *float64 `json:"current,omitempty" yaml:"current,omitempty"`
	// PowerFactor last measured power factor (shown if applicable)
	PowerFactor *float64 `json:"pf,omitempty" yaml:"pf,omitempty"`
	// Aenergy information about the active energy counter (shown if applicable)
	Aenergy *SwitchAenergy `json:"aenergy,omitempty" yaml:"aenergy,omitempty"`
	// Temperature information about the temperature
	Temperature *SwitchTemperature `json:"temperature,omitempty" yaml:"temperature,omitempty"`
	// Error conditions occurred. May contain overtemp, overpower, overvoltage, undervoltage, (shown if at least one error is present)
	Errors []string `json:"errors,omitempty" yaml:"errors,omitempty"`
}

// Clone return copy
func (t *SwitchStatus) Clone() *SwitchStatus {
	c := &SwitchStatus{}
	copier.Copy(&c, &t)
	return c
}

// SwitchAenergy information about the active energy counter (shown if applicable)
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Switch#status
type SwitchAenergy struct {
	// Total energy consumed in Watt-hours
	Total float64 `json:"total" yaml:"total"`
	// ByMinute energy consumption by minute (in Milliwatt-hours) for the last three minutes
	// (the lower the index of the element in the array, the closer to the current moment the minute)
	ByMinute []float64 `json:"by_minute" yaml:"by_minute"`
	// MinuteTs Unix timestamp of the first second of the last minute (in UTC)
	MinuteTs int `json:"minute_ts" yaml:"minute_ts"`
}

// Clone return copy
func (t *SwitchAenergy) Clone() *SwitchAenergy {
	c := &SwitchAenergy{}
	copier.Copy(&c, &t)
	return c
}

// SwitchTemperature System component object
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Sys#status
type SwitchTemperature struct {
	// TC temperature in Celsius (null if temperature is out of the measurement range)
	TC *float64 `json:"tC,omitempty" yaml:"tC,omitempty"`
	// TF temperature in Fahrenheit (null if temperature is out of the measurement
	TF *float64 `json:"tF,omitempty" yaml:"tF,omitempty"`
}

// Clone return copy
func (t *SwitchTemperature) Clone() *SwitchTemperature {
	c := &SwitchTemperature{}
	copier.Copy(&c, &t)
	return c
}

// SwitchConfig configuration of the Switch component contains information about the input mode, the timers and the protection
// settings of the chosen switch instance. To Get/Set the configuration of the Switch component its id must be specified.
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Switch#configuration
type SwitchConfig struct {
	// ID Id of the Switch component instance
	ID int `json:"id" yaml:"id"`
	// Name of the switch instance
	Name *string `json:"name" yaml:"name"`
	// InMode range of values: momentary, follow, flip, detached
	InMode string `json:"in_mode" yaml:"in_mode"`
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
	// AutorecoverVoltageErrors True if switch output state should be restored after over/undervoltage error is cleared, false otherwise (shown if applicable)
	AutorecoverVoltageErrors bool `json:"autorecover_voltage_errors" yaml:"autorecover_voltage_errors"`
	// InputID Id of the Input component which controls the Switch. Applicable only to Pro1 and Pro1PM devices. Valid values: 0, 1
	InputID int `json:"input_id" yaml:"input_id"`
	// PowerLimit Limit (in Watts) over which overpower condition occurs (shown if applicable)
	PowerLimit *float64 `json:"power_limit" yaml:"power_limit"`
	// VoltageLimit Limit (in Volts) over which overvoltage condition occurs (shown if applicable)
	VoltageLimit *float64 `json:"voltage_limit" yaml:"voltage_limit"`
	// UndervoltageLimit Limit (in Volts) under which undervoltage condition occurs (shown if applicable)
	UndervoltageLimit *float64 `json:"undervoltage_limit" yaml:"undervoltage_limit"`
	// CurrentLimit Number, limit (in Amperes) over which overcurrent condition occurs (shown if applicable)
	CurrentLimit *float64 `json:"current_limit" yaml:"current_limit"`
}

// Clone return copy
func (t *SwitchConfig) Clone() *SwitchConfig {
	c := &SwitchConfig{}
	copier.Copy(&c, &t)
	return c
}

// SwitchParams ...
type SwitchParams struct {
	ID     int           `json:"id" yaml:"id"`
	Config *SwitchConfig `json:"config,omitempty" yaml:"config,omitempty"`
	On     *bool         `json:"on" yaml:"on"`
}

// Clone return copy
func (t *SwitchParams) Clone() *SwitchParams {
	c := &SwitchParams{}
	copier.Copy(&c, &t)
	return c
}

type SwitchReport struct {
	Src   string `json:"src,omitempty" yaml:"src"`
	WasOn *bool  `json:"was_on,omitempty" yaml:"was_on,omitempty"`
}

// Clone return copy
func (t *SwitchReport) Clone() *SwitchReport {
	c := &SwitchReport{}
	copier.Copy(&c, &t)
	return c
}
