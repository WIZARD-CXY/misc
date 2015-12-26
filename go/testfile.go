package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("/proc/1/net/dev")

	ss := string(data)

	idx := strings.Index(ss, "eth1")

	lastline := ss[idx:]

	aa := strings.Split(lastline, " ")

	for idx, item := range aa {
		fmt.Printf("%d %q\n", idx, item)

	}
	fmt.Println(aa[1], aa[39])

}
