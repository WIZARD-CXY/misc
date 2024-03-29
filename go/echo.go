package main

import (
	"bufio"
	"fmt"
	"net"
	"test"
)

func main() {
	fmt.Println("wahaha", wahaha)
	l, err := net.Listen("tcp", "127.0.0.1:8053")
	if err != nil {
		fmt.Printf("Failure to listen: %s\n", err.Error())
	}

	for {
		if c, err := l.Accept(); err == nil {
			go Echo(c)
		}
	}
}

func Echo(c net.Conn) {
	defer c.Close()

	line, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		fmt.Printf("Fail to read: %s \n", err.Error())
		return
	}

	_, err = c.Write([]byte(line))
	if err != nil {
		fmt.Printf("Fail to write: %s\n", err.Error())
		return
	}
}
