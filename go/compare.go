package main

import "fmt"
import "reflect"

type data struct {
	num    int               //comparable
	checks [10]func() bool   //not comparable
	doit   func() bool       //not comparable
	m      map[string]string //not comparable
	bytes  []byte            //not comparable but bytes [10]byte is comparable
}

func main() {
	v1 := data{}
	v2 := data{}
	fmt.Println("v1==v2:", reflect.DeepEqual(v1, v2))

	m1 := map[string]string{"one": "a", "two": "b"}
	m2 := map[string]string{"two": "b", "one": "a"}

	fmt.Println("m1==m2:", reflect.DeepEqual(m1, m2))
}
