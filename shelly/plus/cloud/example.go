package cloud

func ExampleConfig() *Config {

	server := "the server"

	return &Config{
		Enable: false,
		Server: &server,
	}
}
