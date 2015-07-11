package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	path, err := exec.LookPath("ifconfig")
	if err != nil {
		log.Fatal("installing fortune is in your future")
	}
	output, _ := exec.Command(path, "eth1").CombinedOutput()

	//ifconfig "$interface" | grep "TX bytes" | cut -d: -f3 | awk '{ print $1 }'
	ss := string(output[:len(output)-1])
	fmt.Println(ss)
	idx := strings.Index(ss, "RX bytes:")
	lastline := ss[idx:]
	aa := strings.Split(lastline, " ")
	rxbytes := strings.Split(aa[1], ":")[1]
	txbytes := strings.Split(aa[6], ":")[1]
	fmt.Println(strings.Split(aa[1], ":")[1], strings.Split(aa[6], ":")[1])
	rx, _ := strconv.ParseUint(rxbytes, 10, 64) //trim trailing '/n'
	tx, _ := strconv.ParseUint(txbytes, 10, 64) //trim trailing '/n'
	fmt.Println(rx, tx)

}
