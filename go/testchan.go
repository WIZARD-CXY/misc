package main

import (
   "fmt"
   "time"
)

func main(){
   chan1 := make(chan bool,100)

   for i:=0; i<10000000; i++{
     go func(i int){
        select {
        case chan1 <- true:
            fmt.Println("haha")
            time.Sleep(10*time.Microsecond)
            fmt.Println("ith",i)
            <-chan1
            fmt.Println("lenxxx:", len(chan1))
        default:
            fmt.Println("rate limited",i,time.Now())
        }
   }(i)

   }
   time.Sleep(10*time.Second)
}
