package main

import "fmt"

func test() {
	defer func() {
		if err := recover(); err != nil {
			println("catch err")
			println(err.(string))
		}
	}()

	panic("panic err here")
}

func test2(a [2]int) {
	fmt.Printf("a: %p\n", &a)
	a[1] = 1000
}

type user struct{ name string }

func main() {
	fmt.Println("1")
	defer fmt.Println("2")
	fmt.Println("3")
	defer fmt.Println("4")
	fmt.Println("5")
	defer fmt.Println("6")
	test()
	a := [2]int{}
	fmt.Printf("a: %p\n", &a)
	test2(a)
	fmt.Println(a)

	s2 := make([]int, 6, 10)
	fmt.Println(s2, len(s2), cap(s2))

	m2 := map[int]*user{
		1: &user{"haha"},
	}

	m2[1].name = "jack"
	fmt.Println(*m2[1])

}
