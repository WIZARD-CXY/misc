/* Go quine */
package main

import "fmt"

func main() {
	fmt.Printf("%s%x%s%x\n", q, 0x60, q, 0x60)
}

var q = `/* Go quine */
package main
import "fmt"
func main() {
    fmt.Printf("%s%x%s%x\n", q, 0x60, q, 0x60)
}
var q = `
