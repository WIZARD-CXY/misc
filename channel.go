package main

import (
	"fmt"
	"math/rand"
	"time"
)

//return a raw message channel of string
func getRawMessageChannel() chan string {
	c := make(chan string)

	go func() {
		for msg := range c {
			fmt.Println(msg)
		}
	}()

	return c

}

//return a message channel of string
func processmessage(v chan string) {
	i := 1

	go func() {
		for {
			v <- fmt.Sprintf("message %d", i)
			time.Sleep(time.Second * time.Duration((rand.Intn(5))))
			i++
		}
	}()

}

func main() {
	/*go func() {
		updates <- 1
		updates <- "hello"

	}()*/
	processmessage(getRawMessageChannel())
	select {}

}
