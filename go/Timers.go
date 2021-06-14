package main

import "time"
import "fmt"

func main(){

    timer1 := time.NewTimer(999*time.Hour)

    <-timer1.C //block here
    fmt.Println("Timer1 expired")

    timer2 := time.NewTimer(time.Second)

    go func(){
        <-timer2.C
        fmt.Println("Timer2 expired")
    }() // non-block approach

    if stop2 := timer2.Stop();stop2 {
        fmt.Println("Timer2 canceled")
    }
}
