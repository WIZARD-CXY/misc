package main 

import "fmt"
import "sync"
func main(){
	var wg sync.WaitGroup

	for i:=0; i<100; i++{
		wg.Add(1)
	}
	i:=0
	for ; i<100; i++{
		go wg.Done()
	}
	fmt.Println("exit",i)
	//wg.Wait()
}