package main

import "fmt"
import "time"

func main() {
	timeChan := make(chan time.Time)
	go func() {
		time.Sleep(1 * time.Second)
		timeChan <- time.Now()

	}()

	time.Sleep(1 * time.Minute)
	completeAt := <-timeChan
	fmt.Println(completeAt, "now is ", time.Now())
}
