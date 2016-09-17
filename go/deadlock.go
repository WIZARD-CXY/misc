package main

import (
	"fmt"
	"sync"
)
      var mutex = &sync.Mutex{}

func foo(){
      mutex.Lock()
      foo()
       mutex.Unlock()
}
func main() {
        foo()
	fmt.Println("Hello, playground")
}
