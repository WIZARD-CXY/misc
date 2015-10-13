package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(&person{name: "Alice", age: 30})

	s := person{
		name: "Sean",
	}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(&s)
	fmt.Println(sp.age)

	sp.name = "Wizard"
	fmt.Println(sp.name)
}
