package main

import "sort"
import "fmt"

type Bylength []string

func (s Bylength) Len() int {
    return len(s)
}

func (s Bylength) Swap(i,j int) {
    s[i],s[j] = s[j],s[i]
}

func (s Bylength) Less(i,j int) bool {
    return (s[i][len(s[i])-1] < s[j][len(s[j])-1])
}

func main() {
    fruits := []string{"mario","peach","banana","kiwi"}
    sort.Sort(Bylength(fruits))
    fmt.Println(fruits)
}
