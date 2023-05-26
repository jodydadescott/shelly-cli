package mdns

import (
	"fmt"
	"time"

	"github.com/hashicorp/mdns"
)

func Execute() {

	params := mdns.DefaultParams("_shelly._tcp")
	params.DisableIPv6 = true
	// params.Entries = make(chan *mdns.ServiceEntry, 4)

	// Make a channel for results and start listening
	entriesCh := make(chan *mdns.ServiceEntry, 4)

	params.Entries = entriesCh

	go func() {
		for entry := range entriesCh {
			fmt.Printf("Got new entry: %v\n", entry)
		}
	}()

	mdns.Query(params)

	time.Sleep(10 * time.Second)
	close(entriesCh)

	//	mdns.Lookup("shelly._tcp", entriesCh)

	// Start the lookup
	// _shelly._tcp

}
