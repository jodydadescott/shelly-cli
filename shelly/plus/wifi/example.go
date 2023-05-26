package wifi

func ExampleConfig() *Config {

	apPass := "My Pass"

	ssid1 := "MY_SSID_1"
	ssid2 := "MY_SSID_2"
	pass1 := "MY_PASS_1"

	nullIfDhcpOrReq := "null if dhcp, required if static"
	nullIfDhcpOrOpt := "null if dhcp, optional if static"

	return &Config{
		Ap: &APConfig{
			Pass:   &apPass,
			IsOpen: false,
			Enable: true,
			RangeExtender: &RangeExtenderConfig{
				Enable: false,
			},
		},
		Sta: &StaConfig{
			SSID:       &ssid1,
			Pass:       &pass1,
			IsOpen:     false,
			Enable:     true,
			Ipv4Mode:   "dhcp",
			IP:         &nullIfDhcpOrReq,
			Netmask:    &nullIfDhcpOrReq,
			Gateway:    &nullIfDhcpOrOpt,
			Nameserver: &nullIfDhcpOrOpt,
		},
		Sta1: &StaConfig{
			SSID:       &ssid2,
			Pass:       &pass1,
			IsOpen:     false,
			Enable:     true,
			Ipv4Mode:   "dhcp",
			IP:         &nullIfDhcpOrReq,
			Netmask:    &nullIfDhcpOrReq,
			Gateway:    &nullIfDhcpOrOpt,
			Nameserver: &nullIfDhcpOrOpt,
		},
		Roam: &RoamConfig{
			RSSIThreshold: -80,
			Interval:      60,
		},
	}
}
