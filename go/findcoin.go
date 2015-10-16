package main

import "fmt"

func main() {
	v := []int{1, 2, 5, 10, 50, 100}

	res := 900756

	Min := make([]int, res+1)

	for i := 1; i <= res; i++ {
		Min[i] = 0x7fffffff
	}
	//Min[i]=j means the min coin need to fit i value is j
	Min[0] = 0

	for i := 1; i <= res; i++ {
		for j := 0; j < len(v); j++ {
			if v[j] <= i && Min[i-v[j]]+1 < Min[i] {
				Min[i] = Min[i-v[j]] + 1
			}
		}
	}

	fmt.Println(Min[res])
}
