package main

import "fmt"

func main() {
	ch := make(chan int)
	ch2 := make(chan bool)

	go shower(ch, ch2)

	for i := 0; i < 10; i++ {
		ch <- i
	}
	ch2 <- true
}

func shower(c chan int, ch2 chan bool) {
	for {
		select {
		case j := <-c:
			fmt.Printf("%d\n", j)
		case <-ch2:
			fmt.Println("haha")

		}
	}

}
