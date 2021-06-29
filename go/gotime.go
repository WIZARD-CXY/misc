package main

import "time"
import "fmt"
import "math/rand"
func main() {
  b := 100*time.Millisecond
  // randomize the batchIntercal
  for i:=0; i<10; i++{
  r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
    interval := time.Duration(b+time.Duration(r1.Intn(10))*time.Duration(int64(b)/10))
    fmt.Println("interval", interval)
    time.Sleep(interval)
  }
}
