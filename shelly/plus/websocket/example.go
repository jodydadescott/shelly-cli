package websocket

func ExampleConfig() *Config {

	sslCa := "the optional SSL CA"

	return &Config{
		Enable: false,
		Server: "the server",
		SslCa:  &sslCa,
	}
}
