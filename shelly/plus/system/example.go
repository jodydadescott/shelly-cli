package system

func ExampleConfig() *Config {

	name := "Name of device or null"
	cfgRev := 6
	addOnType := "null or 'sensor'. See https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Sys#configuration"
	tz := "America/Denver"
	lat := 39.00001
	lon := -104.00001
	listenPort := "9999"
	addr := "1.1.1.1"
	ntpServer := "time.google.com"

	return &Config{
		Device: &DeviceConfig{
			Name:         &name,
			Discoverable: true,
			EcoMode:      false,
			AddonType:    &addOnType,
		},
		Location: &LocationConfig{
			Tz:  &tz,
			Lat: &lat,
			Lon: &lon,
		},
		Debug: &DebugConfig{
			Mqtt: &MqttDebug{
				Enable: true,
			},
			Websocket: &WebsocketDebug{
				Enable: true,
			},
			UDP: &UDP{
				Addr: &addr,
			},
		},
		UIData: &UIDataConfig{},
		RPCUDP: &RPCUDPConfig{
			DstAddr:    addr,
			ListenPort: &listenPort,
		},
		Sntp: &SntpConfig{
			Server: ntpServer,
		},
		CfgRev: &cfgRev,
	}

}
