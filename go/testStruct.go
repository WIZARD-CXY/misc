package main

import "fmt"

type A struct{
	B
	C
}

type B struct{
	i int
}

type C struct{
	i int
}

func (c *C) GetI() int {
	return c.i
}

func (b *B) GetI() int {
	return b.i
}

func main(){
	a := A{B{1},C{2}}
	fmt.Println(a.B.i)

}