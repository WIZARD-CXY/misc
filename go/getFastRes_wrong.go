package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	count := 4
	ch := make(chan string, count)
	var amt int

	for i := 1; i <= count; i++ {
		go func() {
			amt = rand.Intn(8)
			fmt.Println(amt)
			time.Sleep(time.Second * time.Duration(amt))

			ch <- fmt.Sprintf("message %d %d", i, amt)

		}()
	}

	fmt.Println(<-ch)
}
