package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (iPaddr IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", iPaddr[0], iPaddr[1], iPaddr[2], iPaddr[03])
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