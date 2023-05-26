package input

func ExampleConfig() *Config {

	name := "string or null"
	typex := ""
	reportThreshold := 50.0

	return &Config{
		ID:              0,
		Name:            &name,
		Type:            typex,
		Invert:          false,
		FactoryReset:    false,
		ReportThreshold: &reportThreshold,
	}
}
