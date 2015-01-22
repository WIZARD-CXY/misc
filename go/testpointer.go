//testpointer.go
package main

import (
        "fmt"
)

var p *int
var err error

func foo() (*int, error) {
        var i int = 5
        return &i, nil
}

func bar() {
        //use p
        fmt.Println(*p)
}

func main() {

	    //below one is wrong
        // p, err := foo()
	    
	    // below one is correct
        p,err = foo()
        if err != nil {
                fmt.Println(err)
                return
        }
        bar()
        fmt.Println(*p)
}