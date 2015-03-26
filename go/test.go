package main

import "fmt"
import (
	"even"
	"os"
)

func main() {
	fmt.Println(even.Even(2))
	fmt.Println(os.Getenv("DOCKER_HOST"))
}
