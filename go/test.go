package main

import "fmt"
import (
	"even"
	"github.com/GoogleCloudPlatform/kubernetes/test/e2e"
	"os"
)

func main() {
	fmt.Println(even.Even(2))
	fmt.Println(os.Getenv("DOCKER_HOST"))

	a := 1
	b := e2e.TestContextType{}
	fmt.Println("haha", b)
}
