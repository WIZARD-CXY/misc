package main

import "fmt"
func main(){
    foo()
}

func foo(){
    bar()
}

func bar(){
    wiz()
}

func wiz(){
    defer func(){
        x := recover()
            fmt.Println("haha"+x.(string))
        }()
    panic("xixi")
    
    
}

type S struct {
    a int

}

func (s *S) print(){
}