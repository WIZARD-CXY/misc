package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	os.Mkdir("astaxie", 0777)
	os.MkdirAll("astaxie/test1/test2", 0777)
	err := os.Remove("astaxie")
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Second * 10)

	defer os.RemoveAll("astaxie")
}
