package main

import (
	"fmt"
	_ "os"
	"reflect"
	"time"
)

func hello() {
	fmt.Println("Hello world")
}

func main() {
	hl := hello
	fv := reflect.ValueOf(hl)

	fmt.Println("fv is reflect.Func ?", fv.Kind() == reflect.Func)
	fv.Call(nil)

	slc := []int{1, 2, 3, 4, 5}

	for i := 0; i < 5; i++ {
		go func(i int) {
			slc[i]++
		}(i)
	}

	time.Sleep(3 * time.Second)
	fmt.Println(slc)

}
