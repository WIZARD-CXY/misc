package main

import (
	// "container/list"
	"errors"
	"fmt"
)

type NameAge struct {
	name string
	age  int
}
type Stringer interface {
	String() string
}

type Reader interface {
	Read([]byte) (int, error)
}

func Set(t *int) {
	x := t
	*x = 3
}
func Set2(t int) {
	x := &t
	*x = 3
}

type Value int

type Node struct {
	Value
	prev, next *Node
}
type List struct {
	head, tail *Node
}

func (l *List) Front() *Node {
	return l.head
}

func (n *Node) Next() *Node {
	return n.next
}

func (l *List) PushBack(v Value) *List {
	n := &Node{Value: v}

	if l.head == nil {
		l.head = n
	} else {
		l.tail.next = n
		n.prev = l.tail
	}
	l.tail = n

	return l
}

var errEmpty = errors.New("List is Empty")

func (l *List) Pop() (v Value, err error) {
	if l.tail == nil {
		err = errEmpty
	} else {
		v = l.tail.Value
		l.tail = l.tail.prev
		if l.tail == nil {
			l.head = nil
		}
	}
	return
}

func main() {
	a := new(NameAge)
	a.name = "Peter"
	a.age = 42
	bb := 4
	Set2(bb)
	fmt.Println(bb)

	fmt.Printf("%v\n", a)
	// l := list.New()

	l := new(List)

	l.PushBack(1)

	value, err := l.Pop()
	if err == nil {
		fmt.Printf("value %v\n", value)
	} else {
		fmt.Printf("error %v\n", err)
	}

	value, err = l.Pop()
	if err == nil {
		fmt.Printf("value %v\n", value)
	} else {
		fmt.Printf("error %v\n", err)
	}
	// l.PushBack(a)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v\n", e.Value)
	}
}
