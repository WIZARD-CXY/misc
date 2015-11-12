package main

import "fmt"

func m2(a string) string {
	return a + "haha"
}
func main() {
	var m = func(a int) int {
		return a + 1
	}

	res := m(10)      //m defined as a function value
	res2 := m2("wiz") // m2 defined as old way
	res3 := func(a int) int {
		return a + 1
	}(45) //closure func

	fmt.Println("Hello, playground", res, res2, res3)
}
