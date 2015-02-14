package main

import "fmt"
import "time"

func f(from string){
    var input string
    fmt.Println(from)
    fmt.Scanln(&input)
    for i:=0;i<3;i++{
        fmt.Println(from,":",i,input)
    }
}

func main(){
    go f("direct")

    go f("goroutine")

    go func(msg string){
        fmt.Println(msg)
    }("going")
    
    var input string

    fmt.Scanln(&input)

    time.Sleep(2*time.Second)
}
