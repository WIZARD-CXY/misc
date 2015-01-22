package main

import (
	"flag"
	"fmt"
	//"strconv"
)

const (
	//_   = 1000 * iota
	ADD = 2 * iota
	SUB
	MUL
	DIV
	MAXPOS = 11
)

var mop = map[int]string{ADD: "+", SUB: "-", MUL: "*", DIV: "/"}

var (
	ok    bool
	value int
)

type Stack struct {
	i    int
	data [MAXPOS]int
}

func (s *Stack) Reset() { s.i = 0 }

func (s *Stack) Len() int { return s.i }

func (s *Stack) Push(k int) { s.data[s.i] = k; s.i++ }

func (s *Stack) Pop() int { s.i--; return s.data[s.i] }

var found int
var stack = new(Stack)

func main() {
	flag.Parse()
	list := []int{1, 6, 7, 8, 8, 75, ADD, SUB, MUL, DIV}
	fmt.Printf("%v", list)
}
