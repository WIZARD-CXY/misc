package main

import "fmt"
import "sync"

type I1 interface {
	a()
	b()
}

type I2 interface {
	d()
	c()
}

type S struct {
}

func (s S) a() {
	fmt.Println("sa")
}

func (s S) b() {
	fmt.Println("sb")
}

func (s S) c() {
	fmt.Println("sc")
}

func (s S) d() {
	fmt.Println("sd")
}

func main() {
	nums := []int{1, 2, 3, 4}

	var wg sync.WaitGroup
	wg.Add(4)
	for _, x := range nums {
		// new a new variable x, just have the same name with x, but it
		// is a different x
		x := x
		foo := func() {
			fmt.Println(x)
			wg.Done()
		}

		go foo()
	}
	wg.Wait()

	var i1 I1 = S{}
	i1.a()

	if i2, ok := i1.(I2); ok {
		i2.c()
	}
}
