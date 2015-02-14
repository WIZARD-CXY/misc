package main

import "fmt"
import "time"

func worker(id int,jobs <-chan int, result chan<- int){
    for j := range jobs{
        fmt.Println("worker",id,"processing job",j)
        time.Sleep(time.Second)
        result <- j*2
    }
}
func main(){
    jobs := make(chan int,1000)
    results := make(chan int,1000)

    for w := 1; w <=1000; w++ {
        go worker(w,jobs,results)
    }

    for j := 1; j <= 900; j++ {
        jobs <- j;
    }

    close(jobs)
    for a := 1; a <= 899; a++ {
        fmt.Println(<-results)
    }
    close(results)
}
