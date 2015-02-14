package main

import "os"

func main(){
    panic("you should be panic")
    _,err := os.Create("/tmp/file")

    if err != nil {
        panic(err)
    }
}
