package main

import "fmt"
type a struct {
   c int
   d *int
}
func main() {
     structa := a{}  
     dpointer := structa.d
     bpointer := dpointer
     aint :=1
     structa.d = &aint 
     fmt.Println(dpointer, bpointer)
     if bpointer != nil{
        fmt.Print("haha")
     }
}

