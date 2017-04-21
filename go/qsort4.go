package main

import (
	"fmt"
	"math/rand"
        "time"
        "sort"
)

const cutoff=0

func qsort4(a []int, l, u int){
    if u-l < cutoff{
       return
    }
    
    // get a random index from l to u
    idx := rand.Intn(u-l+1)+l

    //swap a[l] with a[idx]

    tmp := a[idx]
    a[idx]=a[l]
    a[l]=tmp
    
    i:=l
    j:=u+1
    t:=a[l]

    for ;;{
       for ;;{
          i++
          if i>u || a[i]>= t {
             break
          }
       }

       for ;; {
         j--
         if j<l || a[j]<=t {
             break
         }
       }

       
       if i>j {
         break;
       }
       temp := a[j]
       a[j]= a[i]
       a[i]=temp
    }

    temp2 := a[l]
    a[l]= a[j]
    a[j]=temp2

	    qsort4(a,l,j-1)
	    qsort4(a,j+1,u)
}



func main() {
		rand.Seed(time.Now().UTC().UnixNano())
		n:=10000000
		a := make([]int,n)

		for i:=0; i<n; i++{
			a[i] = rand.Intn(n)
		}

	/*for i:=0; i<n; i++ {
		fmt.Print(a[i]," ")
	}
        
        fmt.Println("-----------------")*/

	qsort4(a, 0, n-1)

        /*		for i:=0; i<n; i++ {
			fmt.Print(a[i]," ")
		}*/
       fmt.Print(sort.IntsAreSorted(a))
}
