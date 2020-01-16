package main

import "fmt"

type IPAddr [4]byte

func (p IPAddr) String() string {
	var str string
	for i, val := range p {
		if i != len(p) - 1 {
			str += fmt.Sprintf("%v.", val)
		} else {
			str += fmt.Sprintf("%v", val)
		}
	}
	return str
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