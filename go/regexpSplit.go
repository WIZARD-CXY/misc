package main

import "fmt"
import "regexp"

func main() {
	s := "/Hello/World\tHello/\nGolang"
	reg := regexp.MustCompile("/")
	fmt.Printf("%q\n", reg.Split(s, 5)) // have leading ""!!
	// ["Hello" "World" "Hello" "Golang"]
}
