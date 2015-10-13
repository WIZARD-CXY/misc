package main

import "fmt"
import "sync"

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

func seq(in <-chan int, i int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			fmt.Println(i, "#", n)
			out <- n * n
		}
		close(out)
	}()

	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	wg.Add(len(cs))

	// define an output func for each input channel in cs
	// copies the value from c to out until c is closed
	// then call wg.Done
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	for _, c := range cs {
		go output(c)
	}

	// start a gorountine to close out once all the ouput goroutines
	// are done

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	//set up the pipline
	in := gen(2, 3, 4, 5, 6, 7, 8, 9)

	// fan out, distribute the sq work to two goroutine
	c1 := seq(in, 1)
	c2 := seq(in, 2)
	c3 := seq(in, 3)

	//fan in, merge output from c1 and c2

	for n := range merge(c1, c2, c3) {
		fmt.Println(n)
	}
}
