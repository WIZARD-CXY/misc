package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	array := make([]int, 10)

	for i := 0; i < 10; i++ {
		array[i] = i
	}

	rselect(array, 10, 3)

}

// random select m element from array
func rselect(array []int, n int, m int) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		j := rand.Intn(n - i)

		if j < m {
			fmt.Println(array[i])
			m--
		}
	}
}
