//wrong version
package main

import "fmt"
import "time"

func main() {
	vals := []string{"a", "b", "c", "d"}

	for i := 0; i < len(vals); i++ {
		go func() {
			println(i)
			fmt.Println(vals[i])
		}()
	}

	time.Sleep(1 * time.Second)

}
//right version
package main

import "fmt"
import "time"

func main() {
     vals :=[]string{"a","b","c","d"}

     for i:=0; i<len(vals); i++{
        go func(j int){
           fmt.Println(j,vals[j])
        }(i)
     }

     time.Sleep(1*time.Second)
  
}
