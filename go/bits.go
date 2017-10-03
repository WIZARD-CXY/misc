package main

import (
	"fmt"
	"math/bits"
)


func main(){
	var a uint =31

	fmt.Printf("bits.Len(%d) = %d\n",a , bits.Len(a))
	a++
	fmt.Printf("bits.Len(%d) = %d\n",a, bits.Len(a))

	fmt.Printf("bits.OnesCount(%d) = %d\n",a, bits.OnesCount(a))
}