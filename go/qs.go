package main

import (
	"fmt"
	"math/rand"
	"time"
)

func swap(a []int, x int, y int) {
	tmp := a[x]
	a[x] = a[y]
	a[y] = tmp
}
func qs(array []int, k int) int {
	l := 0
	r := len(array) - 1
	fmt.Println("xixi", rand.Intn(10))

	for l < r {
		m := l
		// choose the random pivot
		swap(array, l, l+rand.Intn(r-l+1))

		for i := l + 1; i <= r; i++ {
			if array[i] < array[l] {
				m++
				swap(array, m, i)
			}
		}
		swap(array, m, l)

		if m < k {
			l = m + 1
		} else if m > k {
			r = m - 1
		} else {
			return array[m]
		}
	}
	return array[l]

}
func main() {
	X := []int{1, 3, -2, -10, -1, 100, 60, 30}
	rand.Seed(time.Now().UnixNano())

	fmt.Println("haha", qs(X, 0))
}
