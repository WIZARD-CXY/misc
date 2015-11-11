package main

func foo() {
	println("foo")
}

var funcname = foo

func main() {
	funcname()
}
