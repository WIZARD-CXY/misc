package main

import (
	"fmt"
	"github.com/docker/libcontainer/netlink"
	"log"
	"net"
)

func main() {
	// create a new bridge
	if err := netlink.CreateBridge("mydocker1", false); err != nil {
		log.Fatal(err)
	}
	// get the bridge
	bridge, err := net.InterfaceByName("mydocker1")
	if err != nil {
		log.Fatal(err)
	}

	ip, ipNet, err := net.ParseCIDR("10.0.41.1/16")
	if err != nil {
		log.Fatal(err)
	}

	// add an ip to the bridge
	if err := netlink.NetworkLinkAddIp(bridge, ip, ipNet); err != nil {
		log.Fatal(err)
	}
	// bring the interface up
	if err := netlink.NetworkLinkUp(bridge); err != nil {
		log.Fatal(err)
	}

	fmt.Println("ip 10.10.0.3 %+v", net.ParseIP("10.10.0.3"))
}
