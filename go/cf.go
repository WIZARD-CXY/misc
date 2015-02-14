package main

import "fmt"
import "strings"

func Index(s []string,t string){
    for index,val := range s{
        if val == t{
            return index
        }
    }
    return -1
}

func  Include (vs []string,t string) bool{
    return Index(vs,t) >= 0
}

func Any
