package main

import (
	"fmt"
)

type S struct {
	i int
}

func (p *S) Get() int {
	return p.i
}
func (p *S) Set(v int) {
	p.i = v
}

type I interface {
	Get() int
	Set(int)
}

type R struct{ i int }

func (p *R) Get() int  { return p.i }
func (p *R) Set(v int) { p.i = v }

func f(p I) {
	// fmt.Println(p.Get())
	// p.Set(122)
	// fmt.Println(p.Get())
	switch p.(type) {
	case *S:
		fmt.Println("*S")
		break
	case *R:
		fmt.Println("*R")
		break
	// case S:
	// 	fmt.Println("S")
	// 	break
	// case R:
	// 	fmt.Println("R")
	// 	break
	default:
	}
}
func g(sth interface{}) int {
	return sth.(I).Get()
}

func main() {
	var s S
	f(&s)
	fmt.Println(g(&s))
	i := 5
	fmt.Println(g(i))
	var r R
	f(&r)
}
