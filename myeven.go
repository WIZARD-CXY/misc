package main

import (
	"even"
	"fmt"
	"github.com/kr/fs/walk"
)

func main() {
	i := 6
	fmt.Printf("Is %d even? %v\n", i, even.Even(i))
}
