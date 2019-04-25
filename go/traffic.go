package main

import (
    "fmt"
    "time"
    "golang.org/x/net/context"
    "golang.org/x/time/rate"
)


type pool struct {
   limiter *rate.Limiter
   cnt     int
}

func (p *pool) inc() {
   p.cnt++
}

func main() {
   p := pool{
           limiter: rate.NewLimiter(10,9),
   }


   go func(){
     for {
       ctx := context.Background()
       p.limiter.Wait(ctx)
       p.inc()
     }
   }()
   

   go func(){
     for {
       ctx := context.Background()
       p.limiter.Wait(ctx)
       p.inc()
     }
   }()
   time.Sleep(30*time.Second)
   fmt.Println(p.cnt)
}
