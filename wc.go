package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var chars, words, lines int

	r := bufio.NewReader(os.Stdin)
	for {
		switch s, err := r.ReadString('\n'); {
		case err != nil:
			fmt.Printf("%d %d %d\n", chars, words, lines)
			return
		default:
			chars += len(s)
			words += len(strings.Fields(s))
			lines++

		}

	}
}
