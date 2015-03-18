package main

import (
	"fmt"
	"net/http"
	"strconv"
	"unsafe"
)

func main() {
	s := "abc"
	println(&s)
	s, y := "hello", 20 // 重新赋值: 与前 s 在同一一层次的代码块中,且有新的变量被定义。
	println(&s, y)      // 通常函数多返回值 err 会被重复使用用。
	{
		s, z := 1000, 30
		// 定义新同名变量: 不在同一一层次代码块。
		println(&s, z)
	}

	const (
		a byte = 100
		//b int  = 1e20
	)
	fmt.Println("Hello World")

	x := 0x12345678

	p := unsafe.Pointer(&x)
	n := (*[4]byte)(p)
	fmt.Printf("%T\n", n)

	for i := 0; i < len(n); i++ {
		fmt.Printf("%X ", n[i])
	}
	fmt.Println()
	beyondHello()
}

func beyondHello() {

	var x int
	x = 3
	data, i := [3]int{0, 1, 2}, 0
	i, data[i] = 2, 100

	fmt.Println(i, data)
	y := 4
	sum, prod := learnMultiple(x, y)
	fmt.Println("sum", sum, "prod", prod)
	learnTypes()
}

func learnMultiple(x, y int) (sum, prod int) {
	return x + y, x * y
}

func learnTypes() {
	s := "Learn Go!"

	g := "sigma"
	f := 3.14195
	c := 3 + 4i

	var u uint = 7
	var pi float32 = 22.0 / 7

	n := byte('\n')

	var a4 [4]int
	a3 := [...]int{3, 1, 5, 6}

	s3 := []int{4, 5, 9}

	s4 := make([]int, 4)

	var d2 [][]float64

	bs := []byte("a slice")

	p, q := learnMemory()
	fmt.Println(*p, *q)

	m := map[string]int{"three": 3, "four": 4}

	m["one"] = 1

	_, _, _, _, _, _, _, _ = g, f, u, pi, n, a3, s4, bs

	fmt.Println(s, c, a4, s3, d2, m)

	learnFlowControl()

}

func learnMemory() (p, q *int) {
	p = new(int)

	s := make([]int, 20)
	s[3] = 7
	r := -2
	return &s[3], &r //store on the heap
}

func expensiveCompute() int {
	return 1e6
}

func learnFlowControl() {
	if true {
		fmt.Println("told you")
	}

	if false {

	} else {

	}

	x := 1

	switch x {
	case 0:
	case 1:
		fmt.Println("case 1 reached")
	case 2:

	}

	for x := 0; x < 3; x++ {
		fmt.Println("iteration", x)
	}
	for {
		break
		continue
	}

	if y := expensiveCompute(); y > x {
		x = y
	}

	xBig := func() bool {
		return x > 100
	}

	fmt.Println("xBig:", xBig())

	x /= 1e5
	fmt.Println("xBig:", xBig())

	goto love
love:
	learnInterface()
}

type Stringer interface {
	String() string
}

type pair struct {
	x, y int
}

func (p pair) String() string {
	return fmt.Sprintf("(%d %d)", p.x, p.y)
}

func learnInterface() {

	p := pair{3, 4}

	fmt.Println(p.String())

	var i Stringer
	i = p
	fmt.Println(i.String())

	fmt.Println(p)
	fmt.Println(i)

	learnErrorHandling()

}

func learnErrorHandling() {
	m := map[int]string{3: "three", 4: "four"}
	if x, ok := m[1]; !ok {
		fmt.Println("no one there")
	} else {
		fmt.Print(x)
	}

	if _, err := strconv.Atoi("12"); err != nil {
		fmt.Println(err)
	}
	learnConcurrency()
}

func inc(i int, c chan int) {
	c <- i + 1
}

func learnConcurrency() {
	c := make(chan int, 3)

	go inc(0, c)
	go inc(10, c)
	go inc(-805, c)

	fmt.Println(<-c, <-c, <-c)

	cs := make(chan string)
	cc := make(chan chan string)
	fmt.Println("block")

	go func() { c <- 84 }()
	go func() { cs <- "wordy" }()

	select {
	case i := <-c:
		fmt.Println("it's a %T", i)
	case <-cs:
		fmt.Println("it's a string")
	case <-cc:
		fmt.Println("didn't happen")
	}

	learnWebProgramming()
}

func learnWebProgramming() {
	err := http.ListenAndServe(":8000", pair{})
	fmt.Println(err)
}

func (p pair) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(*r)
	w.Write([]byte("You learned Go in Y minutes!"))
}
