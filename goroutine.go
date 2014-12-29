package main

import (
	"fmt"

	"time"
)

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready!")
	c <- 1
}

var c chan int

func main() {
	c = make(chan int, 1)
	go ready("Tea", 2)
	ready("Coffee", 1)
	fmt.Println("I'm waiting")
	time.Sleep(5 * time.Second)

	fmt.Println(<-c)
	//<-c
	//<-c
}
