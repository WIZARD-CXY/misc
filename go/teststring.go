package main

import "fmt"
import "strings"
import "regexp"

func main() {
	s := "/containers/abcdedf/start"
	ss := regexp.MustCompile("/").Split(s, 5)

	if strings.HasPrefix(s, "/v") {
		//cid := ss[3]
		fmt.Printf("%q", ss)

	} else {
		//cid := ss[2]
		fmt.Printf("%q", ss)
		fmt.Println(len(ss))
	}

}
