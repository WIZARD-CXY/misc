package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

/*

This code contains data races because the updateFooSlice() goroutine
is changing the same *FooSlice as is being served up by the
http handler without any sync protection.

Should a sync.mutex variable be added or is there a slicker way
to expose internal memory structures to an http handler?

How to reproduce:

    $ go run -race main.go
    $ httpie http://localhost:8080

*/

type Foo struct {
	content string
}

type FooSlice []*Foo

var mutex sync.Mutex

func updateFooSlice(fooSlice FooSlice) {

	for {
		mutex.Lock()
		foo := &Foo{content: "new"}
		fooSlice[0] = foo
		fooSlice[1] = nil
		mutex.Unlock()
		time.Sleep(time.Second)
	}

}

func installHttpHandler(fooSlice FooSlice) {

	handler := func(w http.ResponseWriter, r *http.Request) {
		mutex.Lock()
		defer mutex.Unlock()
		for _, foo := range fooSlice {
			if foo != nil {
				fmt.Fprintf(w, "foo: %v ", (*foo).content)
			}
		}

	}

	http.HandleFunc("/", handler)
}

func main() {

	foo1 := &Foo{content: "hey"}
	foo2 := &Foo{content: "yo"}

	fooSlice := FooSlice{foo1, foo2}

	installHttpHandler(fooSlice)

	go updateFooSlice(fooSlice)

	http.ListenAndServe(":8080", nil)

}
