package switchx

func ExampleConfig() *Config {

	name := "Switch Name"
	inMode := "momentary, follow, flip, detached"
	initialState := "off, on, restore_last, match_input"

	var powerLimit float64 = 4480
	var undervoltageLimit float64 = 0
	var currentLimit float64 = 16

	return &Config{
		ID:                       0,
		Name:                     &name,
		InMode:                   inMode,
		InitialState:             initialState,
		AutoOn:                   false,
		AutoOnDelay:              60,
		AutoOff:                  false,
		AutoOffDelay:             60,
		AutorecoverVoltageErrors: false,
		PowerLimit:               &powerLimit,
		UndervoltageLimit:        &undervoltageLimit,
		CurrentLimit:             &currentLimit,
	}
}
