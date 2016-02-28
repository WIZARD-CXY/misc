package main

import "fmt"
import (
	"even"
	"os"
)

func main() {
	fmt.Println(even.Even(2))
	fmt.Println(os.Getenv("DOCKER_HOST"))

	var z float64

	fmt.Println(z, -z, 1/z, -1/z, z/z)
}
