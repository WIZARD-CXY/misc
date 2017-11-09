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
type A struct {
	DD int
	CC string
}
func main() {
	hl := hello
	fv := reflect.ValueOf(hl)

	fmt.Println("fv is reflect.Func ?", fv.Kind() == reflect.Func)
	i := A{}
	iv := reflect.ValueOf(i)
	t:= iv.Type()
	fmt.Println(t)

	for i:=0; i < iv.NumField(); i++{
		f := iv.Field(i)
		fmt.Printf("%d: %s %s = %#v\n",i, t.Field(i).Name, f.Type(), f.Interface())
	}
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
