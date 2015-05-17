package main

import "fmt"

func main() {
	ch := make(chan string)
	//done := make(chan bool)

	go func() {
		for m := range ch {
			fmt.Println("processed:", m)
		}

		//done <- true
	}()

	ch <- "cmd.1"
	ch <- "cmd.2" //won't be processed
	close(ch)
	//<-done
}

/*buggy code the output will either be

cmd.1

or

cmd.1
cmd.2

It is not determistic*/
