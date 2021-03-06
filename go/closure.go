package main

import "fmt"

// This function `intSeq` returns another function, which
// we define anonymously in the body of `intSeq`. The
// returned function _closes over_ the variable `i` to
// form a closure.
func intSeq() func() int {
    i := 0
    return func() int {
        i += 1
        return i
    }
}

func main() {

    // We call `intSeq`, assigning the result (a function)
    // to `nextInt`. This function value captures its
    // own `i` value, which will be updated each time
    // we call `nextInt`.
    nextInt1 := intSeq()

    // See the effect of the closure by calling `nextInt`
    // a few times.
    fmt.Println(nextInt1())
    fmt.Println(nextInt1())
    fmt.Println(nextInt1())
    fmt.Println(nextInt1())

    // To confirm that the state is unique to that
    // particular function, create and test a new one.
    newInts2 := intSeq()
    fmt.Println(newInts2())
}
