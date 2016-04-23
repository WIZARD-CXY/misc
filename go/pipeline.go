package main

import (
	"fmt"
	"sync"
)

func gen(done chan struct{}, in ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range in {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}()

	return out
}

func sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}

		}

		close(out)
	}()

	return out
}

func merge(done chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	out := make(chan int)

	wg.Add(len(cs))

	for _, c := range cs {
		go func(c <-chan int) {
			defer wg.Done()
			for n := range c {
				select {
				case out <- n:
				case <-done:
					return
				}
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
func main() {
	// done channel will be passed to every stage function
	done := make(chan struct{}, 2)
	defer close(done)
	c := gen(done, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)

	c1 := sq(done, c)
	c2 := sq(done, c)

	out := merge(done, c1, c2)

	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
	// just got enough, do not need anymore
}
