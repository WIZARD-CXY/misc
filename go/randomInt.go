package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan bool, 4)
	ch <- true
	ch <- true
	close(ch)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	println(len(ch), cap(ch))
	for i := 0; i < len(ch)+10; i++ {
		fmt.Println("rand", r.Intn(3))
		v, ok := <-ch
		fmt.Println(v, ok)
	}
}
