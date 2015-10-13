package main

import "fmt"

func gen(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}
func main() {
	chan1 := gen(2, 3, 4)
	chan2 := square(square(square(chan1)))

	for res := range chan2 {
		fmt.Println(res)
	}
}
