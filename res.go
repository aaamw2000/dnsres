package dnsres

import (
	"fmt"
	"net"
)

func Resolve(ip string) error {
	addrs, err := net.LookupIP(ip)
	if err != nil {
		return err
	}

	fmt.Printf("IP address for %s is\n", ip)
	for i, addr := range addrs {
		fmt.Printf("%d => %s\n", i, addr.String())
	}
	return nil
}
