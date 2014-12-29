package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
}

func hello() {
	fmt.Println("Hello world!")
}

func main() {
	var x float64 = 3.14
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is floate64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())

	t := T{23, "hello world"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

	hl := hello
	fv := reflect.ValueOf(hl)
	fmt.Println("fv is reflect.Func ?", fv.Kind() == reflect.Func)
	fv.Call(nil)
}
