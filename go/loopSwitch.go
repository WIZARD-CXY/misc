package main

import "fmt"

func main() {
	a := 1
	b := 2
	//loop:
	for {
		switch a {
		case 1:
			if b == 1 {
				fmt.Println("doA()")
				//break // like 'goto A'
				goto A
			}

			if b == 2 {
				fmt.Println("doB()")
				//break loop // like 'goto B'
				goto B
			}

			fmt.Println("doC()")
		case 2:
			// ...
		}
	A:
		fmt.Println("doX()")
		// ...
	}

B:
	fmt.Println("doY()")

}
