package main

import "net"
import "fmt"
import "encoding/binary"
import "bytes"

func main() {
	//string convert to net.IP ,then convert to uint32
	cidr := "255.0.0.0/11"
	_, ipNet, _ := net.ParseCIDR(cidr)

	var num uint32

	buf := bytes.NewBuffer(ipNet.IP)

	err := binary.Read(buf, binary.BigEndian, &num)
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
