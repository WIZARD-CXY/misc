package main

import (
	"fmt"
	"sync"
)

func main() {
	p := &sync.Pool{
		New: func() interface{} {
			return 2
		},
	}

	a := p.Get().(int)
	p.Put(1)
	//runtime.GC()
	b := p.Get().(int)
	c := p.Get().(int)
	fmt.Println(a, b, c)
}
