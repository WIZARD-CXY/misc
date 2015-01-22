package main

import "net"
import "fmt"

func main() {
	l, _ := net.Listen("tcp", ":0")
	defer l.Close()
	println(l.Addr().String())

	for {

		inConn, _ := l.Accept()

		fmt.Printf("Accepted TCP connection from %v to %v\n", inConn.RemoteAddr().String(), inConn.LocalAddr().String())
	}
}
