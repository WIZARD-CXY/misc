package main

import "fmt"
import "strconv"
import "sync"
import "os"
import "bytes"

type stack struct {
	i    int
	data [10]int
}

func (s *stack) push(k int) {
	if s.i+1 > 9 {
		return
	}
	s.data[s.i] = k
	s.i++
}
func (s *stack) pop() int {
	s.i--
	return s.data[s.i]
}

func (s stack) String() string {
	var str string
	for i := 0; i < s.i; i++ {
		str = str + "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
	}
	return str
}
func main() {
	fmt.Println(os.TempDir())

	fmt.Println("Hello World")
	var c complex64 = 5 + 5i
	fmt.Printf("Value is: %v\n", c)
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
	defer func(x int) {
		fmt.Printf("%d ", x)
	}(8900)

	a := func() {
		println("Hello")
	}

	//a()

	for i := 0; i < 10; i++ {
		fmt.Printf("%v\n", i)
	}
	//fmt.Printf("%v\n", i)

	fmt.Printf("%T\n", a)

	aa := [...]float64{1.2, 2.4, 3.6, 4.7, 5.9}
	s1 := aa[2:5]
	bbb := []int{1, 10, 900, 2}
	bbb = append(bbb, 1000)

	bb := make([]int, 10)
	_ = bb
	//fmt.Printf("%v", average(aa))
	fmt.Printf("average val is %v\n", average(s1))

	s := new(stack)
	s.push(25)
	s.push(90)
	s.pop()

	fmt.Printf("stack %v\n", s)
	prtall(1, 2, 3, 4, 5, 6, 7, 89, 90)

	m := []int{1, 4, 5, 89}
	f := func(i int) int {
		return i * i
	}

	fmt.Printf("%v\n", Map(f, m))
	fmt.Println("max is " + strconv.Itoa(max(bbb)))

	fmt.Println("---------------------------")
	n := []int{5, -1, 0, 12, 3, 5, -10}

	fmt.Printf("unsorted %v\n", n)
	bubbleSort(n)
	fmt.Printf("sorted %v\n", n)

	p3 := plusX(3)
	fmt.Printf("haha %v\n", p3(3))

	var p *int
	fmt.Printf("%v\n", p)

	var i int
	p = &i

	fmt.Printf("%v\n", p)
}

type SyncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}

func average(xs []float64) (avg float64) {
	sum := 0.0
	switch len(xs) {
	case 0:
		avg = 0
	default:
		for _, v := range xs {
			sum += v
		}
		avg = sum / float64(len(xs))
	}
	return
}

func prtall(numbers ...int) {
	for index, x := range numbers {
		fmt.Printf("%v %v \n", index+2, x)
	}
}
func Map(f func(int) int, l []int) []int {
	j := make([]int, len(l))

	for k, v := range l {
		j[k] = f(v)
	}
	return j
}
func max(l []int) (max int) {
	max = l[0]
	for _, v := range l {
		if v > max {
			max = v
		}
	}
	return
}

func bubbleSort(n []int) {
	for i := 0; i < len(n)-1; i++ {
		for j := i + 1; j < len(n); j++ {
			if n[j] < n[i] {
				n[i], n[j] = n[j], n[i]
			}
		}
	}
}
func plusX(x int) func(int) int {
	return func(y int) int { return x + y }
}
