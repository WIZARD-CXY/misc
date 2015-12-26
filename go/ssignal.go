package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c := make(chan os.Signal, 22)
	signal.Notify(c, os.Interrupt, syscall.SIGKILL)

	s := <-c
	fmt.Println("Got signal:", s)
}
