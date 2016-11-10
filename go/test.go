package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

func foo() error {
	return errors.New("xixi")
}
func main() {
	now := time.Now()
	//
	go func() {
		if err := foo(); err != nil {
			log.Fatal("xixi")
		}
	}()
	time.Sleep(1 * time.Second)
	for i := 0; i < 1000; i++ {
		fmt.Println(i, "hahah")
	}

	fmt.Println(time.Since(now).String())
}
