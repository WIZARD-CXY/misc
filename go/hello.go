package main

import (
	"fmt"
	_ "os"
	"reflect"
)

func hello() {
	fmt.Println("Hello world")
}

func main() {
	hl := hello
	fv := reflect.ValueOf(hl)

	fmt.Println("fv is reflect.Func ?", fv.Kind() == reflect.Func)
	fv.Call(nil)

}
