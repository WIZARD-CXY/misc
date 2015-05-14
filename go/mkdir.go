package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	os.Mkdir("astaxie", 0777)
	os.MkdirAll("astaxie/test1/test2", 0777)

	defer os.RemoveAll("astaxie")
	err := os.Remove("astaxie")
	if err != nil {
		fmt.Println(err)
	}

	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		<-sigs
		log.Println("program killed")
		os.RemoveAll("astaxie")

		os.Exit(0)
	}()

	time.Sleep(time.Second * 10)
}
