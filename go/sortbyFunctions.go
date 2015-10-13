package main

import "sort"
import "fmt"

type stringSlice []string

func (s stringSlice) Len() int {
	return len(s)
}

func (s stringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s stringSlice) Less(i, j int) bool {
	return (s[i][len(s[i])-1] < s[j][len(s[j])-1])
}

func helper(nums []int) {
	for i := 0; i < len(nums); i++ {
		nums[i] += 10
	}
}
func main() {
	fruits := []string{"mario", "peach", "banana", "kiwi"}
	nums := []int{10, 20, 30}

	helper(nums)

	for _, x := range nums {
		println(x)
	}
	for _, f := range fruits {
		f += "a"
	}

	for _, f := range fruits {
		println(f)
	}
	sort.Sort(stringSlice(fruits))
	fmt.Println(fruits)
}
