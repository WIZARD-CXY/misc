package main

import "fmt"
import (
	"os"
)

func main() {
	fmt.Println(os.Getenv("DOCKER_HOST"))
}
