package main

import (
	"errors"
	"reflect"
)

func foo() {
	println("foo")
}

func bar(a, b, c, d int) {
	println(a, b, c, d)
}

func Call(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])

	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is not adapted")
		println("haha")
		return
	}

	in := make([]reflect.Value, len(params))

	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}

	result = f.Call(in)

	return
}
func main() {
	funcs := map[string]interface{}{"foo": foo, "bar": bar}

	Call(funcs, "foo")
	Call(funcs, "bar", 1, 2, 3, 4)
}
