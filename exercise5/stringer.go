package main

import (
	"fmt"
)

// IPAddr is an IP address.
type IPAddr [4]byte

// String() for IP address.
// {0, 0, 0, 0} => "0.0.0.0"
func (ipaddr IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ipaddr[0], ipaddr[1], ipaddr[2], ipaddr[3])

	// The following statement doesn't work but I don't know the reason...
	// return string(ipaddr[0]) + "." + string(ipaddr[1]) + "." + string(ipaddr[2]) + "." + string(ipaddr[3])
	// Result:
	// loopback: ...
	// googleDNS:.
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
