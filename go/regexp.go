package main

import "regexp"
import "fmt"

func main() {
	s1 := "/containers/<cid>/start"
	s2 := regexp.MustCompile("/").Split(s1, 5)
	fmt.Printf("%v", s2)

}
