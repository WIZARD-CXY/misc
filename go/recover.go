package main

import "fmt"
import "os"
import "runtime"

// recover from specific error
func main() {
        type bailout struct{}
        defer func(){
            switch p:= recover(); p{
               case nil:
                   //no panic
               case bailout{}:
                   // expected error
  
                   //printStack()
                   fmt.Println("caught you")
               default:
                   panic(p) // unexpected error, carry on panic
            }
        }()
 
	fmt.Println("Hello, playground")
	panic(bailout{})
}

func printStack(){
     var buf [500]byte
     n := runtime.Stack(buf[:],false)
     os.Stdout.Write(buf[0:n])
}
