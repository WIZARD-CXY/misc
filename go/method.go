package main

import "fmt"
import "reflect"

type rect struct {
	width, height int
}

func (r *rect) Area() int {
	return r.width * r.height
}

func (r rect) Perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	mtv := reflect.ValueOf(&r).Elem()
	fmt.Println("Before Perim:", mtv.MethodByName("Perim").Call(nil)[0].Interface().(int))

	fmt.Println("..............")
	fmt.Println("area: ", r.Area())
	fmt.Println("perim: ", r.Perim())

	rp := &r

	fmt.Println("area: ", rp.Area())
	fmt.Println("perim: ", rp.Perim())
}
