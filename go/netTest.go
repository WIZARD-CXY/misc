package main

import "net"
import "math"

func main() {
	cidr := "192.168.0.1/11"
	_, addr, _ := net.ParseCIDR(cidr)
	mask, _ := addr.Mask.Size()
	println(mask)
	if addr.IP.To4() != nil {
		print(math.Pow(2, float64(32-mask)))
	} else {
		print(math.Pow(2, float64(128-mask)))
	}
}
