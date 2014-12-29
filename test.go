package main

//import "sync"
import "fmt"
import "time"
import "github.com/golang/glog"

var (
	wahaha int = 9
)

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

	time.Sleep(10 * time.Second)
	glog.Fatalf("Error creating root directory: ")
	fmt.Println("haha")

}
