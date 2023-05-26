package ethernet

func ExampleConfig() *Config {

	nullIfDhcpOrReq := "null if dhcp, required if static"
	nullIfDhcpOrOpt := "null if dhcp, optional if static"

	return &Config{
		Enable:     true,
		Ipv4Mode:   "dhcp",
		IP:         &nullIfDhcpOrReq,
		Netmask:    &nullIfDhcpOrReq,
		Gateway:    &nullIfDhcpOrOpt,
		Nameserver: &nullIfDhcpOrOpt,
	}

}
