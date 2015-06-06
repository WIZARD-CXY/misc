package main

import "net"
import "fmt"
import "encoding/binary"
import "bytes"

func main() {
	//string convert to net.IP ,then convert to uint32
	cidr := "192.168.0.0/16"
	_, ipNet, _ := net.ParseCIDR(cidr)
	ipString := "192.168.0.4"
	ipp := net.ParseIP(ipString)

	var nump uint32

	bufp := bytes.NewBuffer(ipp.To4())

	err := binary.Read(bufp, binary.BigEndian, &nump)
	if err == nil {
		fmt.Println("hahap", nump)
	} else {
		fmt.Println(err)
	}

	var num uint32

	buf := bytes.NewBuffer(ipNet.IP)

	err = binary.Read(buf, binary.BigEndian, &num)
	if err == nil {
		fmt.Println("haha", num)
	} else {
		fmt.Println(err)
	}
	num += 1

	fmt.Println("haha", num)

	buf2 := new(bytes.Buffer)
	err = binary.Write(buf2, binary.BigEndian, num)

	if err == nil {
		fmt.Printf("%#v", buf2.Bytes())
	} else {
		fmt.Println(err)
	}

	// []byte convert to net.IP then convert to string
	ip2 := net.IP(buf2.Bytes()).String()
	fmt.Println(ip2)

}
