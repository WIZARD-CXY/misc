package main

import (
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

var errEmpty = errors.New("List is empty")

type Node struct {
	val        int
	prev, next *Node
}
type List struct {
	head, tail *Node
}

func (l *List) Front() *Node {
	return l.head
}

func (l *List) PushBack(v int) *List {
	newNode := &Node{val: v}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		//insert it to the tail
		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode
	}

	return l
}

func (l *List) Pop() (v int, err error) {
	if l.tail == nil {
		err = errEmpty
	} else {
		v = l.tail.val
		l.tail = l.tail.prev
		if l.tail == nil {
			l.head = nil
		}
	}
	return
}

func (n *Node) Next() *Node {
	return n.next
}

func main() {
	a := new(NameAge)
	a.name = "Peter"
	a.age = 42
	fmt.Printf("%#v\n", a)

	bb := 4
	Set(&bb)
	fmt.Println(bb)

	//l := list.New()

	l := new(List)

	l.PushBack(1)

	value, err := l.Pop()
	if err == nil {
		fmt.Printf("value %v\n", value)
	} else {
		fmt.Printf("error %v\n", err)
	}

	//l.PushBack(3)
	//l.PushBack(3)
	value, err = l.Pop()
	if err == nil {
		fmt.Printf("value %v\n", value)
	} else {
		fmt.Printf("error %v\n", err)
	}
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)
	l.PushBack(6)
	l.PushBack(7)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v\n", e.val)
	}
}
