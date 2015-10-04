/*
another aio
2015-10-4
*/
package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"runtime"
	"time"
)

// outside the function, every construct begins with a keyword (var, func, type)and := is not available
// the var statement declares a list of vars, and the type is coming last
var x, y, z int
var c, a, gogo bool

var x1, y1, z1 int = 1, 2, 3

//if an initializer is present, the type can be omitted
var x2, y2, z2 = 1, 2, "f"

//basic type
//bool
//string
//int int8 int16 int32 int64
//uint uint8 uint16 uint32 uint64 uintptr
//byte (alias uint8)
//rune (alias int32)
//float32 float64
//complex64 complex128

var (
	Tobe    bool   = false
	MaxUInt uint64 = 1<<64 - 1
)

//const is declarede like vars
const Pi = 3.14

type vertex struct {
	x int
	y int
}

type Vertex2 struct {
	latx, laty float64
}

var m map[string]Vertex2

var m2 = map[string]Vertex2{
	"chenxingyu": Vertex2{
		1, 2,
	},
	"haha": Vertex2{
		-3, -9,
	},
}

type MyFloat float64

type Abser interface {
	Abs() float64
}

func main() {
	fmt.Println("Hi~ Wiz")
	fmt.Println("The time is", time.Now())

	//random number generator, for different
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println("My favourite number is", r.Intn(7))
	fmt.Println("My favourite number is", r.Intn(7))
	fmt.Println("My favourite number is", r.Intn(7))

	fmt.Printf("Now you have %g problems\n", math.Nextafter(2, 3))

	// A name is exported if it begins with a capital letter
	fmt.Println(math.Pi)

	//type go after the variable name
	fmt.Println(add(42, 13))
	fmt.Println(add(42, 23))

	//multiple results
	a, b := swap("jaja", "haha")

	fmt.Println(a, b)

	//named return
	fmt.Println(split(17))
	fmt.Println(split2(17))

	// used var
	fmt.Println(x, y, z)
	// inside a function, the short := assignment can be used in place of a var
	c2, python, java := true, false, "aha"
	fmt.Println(c2, python, java)

	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	ival := 1
	for {
		if ival++; ival > 1000 {
			fmt.Println("leave the infinite loop", ival)
			break
		}
	}

	fmt.Println(sqrt(2), sqrt(-4))

	pv := new(vertex)
	fmt.Println(pv)

	as := []int{1, 2, 3, 4}
	for i := 0; i < len(as); i++ {
		fmt.Printf("as[%d] = %d \n", i, as[i])
	}

	//The slice can be resliced, creating a new slice that point to the same array
	// use expression s[lo:hi] through element s[lo] to s[hi-1]
	fmt.Println("as[1:4]", as[1:4])

	// slice can also be make, cap is default
	s1 := make([]int, 5)
	_ = s1
	m = make(map[string]Vertex2)

	m4 := make(map[string]int)

	m4["date"] = 20130506
	fmt.Println("m4[date]", m4["date"])
	m4["date"]++

	fmt.Println("m4[date]", m4["date"])
	delete(m4, "date")

	date2, ok := m4["date"]
	fmt.Println("the value is:", date2, "Present?", ok)

	//func values
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(3, 4))

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Mac OS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println(os)
	}

	f1 := MyFloat(-math.Sqrt2)
	fmt.Println(f1.Abs())

	// An interface type is defined by a set of methods
	// a value of interface type can hold any value that implement those methods
	var a_interface Abser

	a_interface = f1
	a_interface = Vertex2{1.3, 1.3}

	fmt.Println(a_interface.Abs())

	// Interface are satisfied
	var w Writer
	//os.Stdout implement Writer interface
	w = os.Stdout

	fmt.Fprintf(w, "Hello world\n")

}
func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	return y, x
}

func split2(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	return
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func (v Vertex2) Abs() float64 {
	return math.Sqrt(v.latx*v.latx + v.laty*v.laty)
}

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}
