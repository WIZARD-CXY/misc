package main

import (
	"fmt"
)

func decorator(f func(s string)) func(s string) error {
	return func(s string) error {
		fmt.Println("Start")
		f(s)
		fmt.Println("end")

		return nil
	}
}

func hello(s string) {
	fmt.Println("Hello", s)
}

func main() {
	f := decorator(hello)
	f("world")
}
