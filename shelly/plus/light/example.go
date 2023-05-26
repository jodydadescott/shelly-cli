package light

func ExampleConfig() *Config {

	name := "Switch Name"
	initialState := "off, on, restore_last, match_input"

	// type LightConfig struct {
	// 	// ID Id of the Switch component instance
	// 	ID int `json:"id" yaml:"id"`
	// 	// Name of the switch instance
	// 	Name *string `json:"name" yaml:"name"`
	// 	// InitialState range of values: off, on, restore_last, match_input
	// 	InitialState string `json:"initial_state" yaml:"initial_state"`
	// 	// AutoOn True if the "Automatic ON" function is enabled, false otherwise
	// 	AutoOn bool `json:"auto_on" yaml:"auto_on"`
	// 	// AutoOnDelay Seconds to pass until the component is switched back on
	// 	AutoOnDelay float64 `json:"auto_on_delay" yaml:"auto_on_delay"`
	// 	// AutoOff True if the "Automatic OFF" function is enabled, false otherwise
	// 	AutoOff bool `json:"auto_off" yaml:"auto_off"`
	// 	// AutoOffDelay Seconds to pass until the component is switched back off
	// 	AutoOffDelay float64 `json:"auto_off_delay" yaml:"auto_off_delay"`
	// 	// DefaultBrightness brightness level (in percent) after power on
	// 	DefaultBrightness float64 `json:"default.brightness" yaml:"default.brightness"`
	// 	// NightModeEnable Enable or disable night mode
	// 	NightModeEnable bool `json:"night_mode.enable" yaml:"night_mode.enable"`
	// 	// NightModeBrightness brightness level limit when night mode is active
	// 	NightModeBrightness float64 `json:"night_mode.brightness" yaml:"night_mode.brightness"`
	// 	// NightModeActiveBetween containing 2 elements of type string, the first element indicates the start of
	// 	// the period during which the night mode will be active, the second indicates the end of that period.
	// 	// Both start and end are strings in the format HH:MM, where HH and MM are hours and minutes with optinal
	// 	// leading zeros
	// 	NightModeActiveBetween []string `json:"night_mode.active_between" yaml:"night_mode.active_between"`
	// }

	return &Config{
		ID:           0,
		Name:         &name,
		InitialState: initialState,
		AutoOn:       false,
		AutoOnDelay:  60,
		AutoOff:      false,
		AutoOffDelay: 60,
	}
}
