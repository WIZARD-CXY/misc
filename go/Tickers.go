package main

import "time"
import "fmt"

func main(){

    ticker := time.NewTicker(time.Millisecond*500)

    go func(){
        for  {
            fmt.Println("Ticker at", <-ticker.C)
        }
    }()

    time.Sleep(time.Second*1500)
    ticker.Stop()
    fmt.Println("Ticker stopped")
}
