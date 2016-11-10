package main

import (
	"fmt"
	"time"
)

func foo(t time.Time) time.Time{
     time.Sleep(1*time.Second)
     fmt.Println("haha",time.Now())
     fmt.Println("xime",t)
     return t
}
func main() {
	fmt.Println("Hello, playground")
        
        go foo(foo(time.Now()))
        fmt.Println("xixi",time.Now())
	
	time.Sleep(3*time.Second)
}
