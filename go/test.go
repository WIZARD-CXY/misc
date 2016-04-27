package main

import (
	"fmt"
)

type nb interface {
	Dohaha()
}
type Point struct {
	x, y float64
	haha func()
}

func (p *Point) ScaleBy(factor float64) {
	p.x *= factor
	p.y *= factor
}

func (p Point) Dohaha() {
	p.haha()
}

func test(a nb) {
	a.Dohaha()
}
func main() {
	//Point{1, 2}.ScaleBy(2)

	a := Point{x: 1, y: 2}
	a.haha = func() {
		fmt.Println("haha")
	}

	test(a)

}
