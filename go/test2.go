package main

import (
 "encoding/json"
 "fmt"
)


type A struct {
     B int
     C *int
}
   
func main(){
    //x := 0
    a := A{B:1,}
    d,_ := json.Marshal(a)
    fmt.Println(string(d))
   
    dd := A{}
    _ = json.Unmarshal(d,&dd)
    fmt.Println(dd.C)
    // print nil
    data := `{"b":10}` 
    d2 := A{}
    _ = json.Unmarshal([]byte(data),&d2)
    fmt.Println(d2)
}    
