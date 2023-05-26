package mqtt

func ExampleConfig() *Config {

	clientID := "If left empty will be set automatically based on device ID"
	server := "10.10.10.10"
	topicPrefix := "shelly"
	rpcNtf := true
	statusNtf := true
	useClientCert := false
	enableControl := true

	return &Config{
		Enable:        true,
		Server:        &server,
		TopicPrefix:   &topicPrefix,
		ClientID:      &clientID,
		RPCNtf:        rpcNtf,
		StatusNtf:     statusNtf,
		UseClientCert: useClientCert,
		EnableControl: enableControl,
	}
}
