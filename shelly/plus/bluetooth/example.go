package bluetooth

func ExampleConfig() *Config {
	return &Config{
		Enable: true,
		RPC: &RPC{
			Enable: true,
		},
	}
}
