package main

//import "sync"
import "fmt"
import "time"
import "github.com/golang/glog"

var (
	wahaha     int = 9
	interfacea     = []a{}
)

type a interface {
	num2string(num int) string
	string2num(s string) int
}

func getInt() int {
	return 1
}

func getInt2() int {
	return 2
}

//var s string

func main() {
	// lock1 := sync.Mutex{}
	// lock2 := sync.Mutex{}
	// //lock3 := sync.Mutex{}
	// //lock3.Lock()
	// //
	// // go func() {
	// // lock1.Lock()
	// // lock2.Lock()
	// // }()

	// //lock3.Unlock()
	// lock2.Lock()
	// lock1.Lock()

	go func() {
		fmt.Println("in first goroutine")
		go func() {
			fmt.Println("in second goroutine")
		}()
	}()

	s := "abc"

	println(&s)

	s, y := "test", 20

	println(&s, y)

	{
		s, z := 1000, 30
		// 定义新同名变量: 不在同一一层次代码块。
		println(&s, z)
	}

	time.Sleep(10 * time.Second)
	glog.Fatalf("Error creating root directory: ")
	fmt.Println("haha")

}
