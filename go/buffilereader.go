package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	buf := make([]byte, 1024)
	f, _ := os.Open("/etc/passwd")
	defer f.Close()

	r := bufio.NewReader(f)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	for {
		n, _ := r.Read(buf)
		if n == 0 {
			break
		}
		w.Write(buf[0:n])
	}

	cmd := exec.Command("top")
	buf2, err := cmd.Output()
	if err != nil {
		fmt.Println("err")
	} else {
		fmt.Println("good", buf2)
	}
}
