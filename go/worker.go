package main

import "fmt"

type Worker struct {
	x, y, z int
}

func worker(in <-chan *Work, out chan<- *Work) {
	for w := range in {
		w.z = w.x * w.y
		time.Sleep(w.z * time.Second)
		out <- w
	}
}

func Run() {
	in, out := make(chan *Work), make(chan *Work)

	for i := 0; i < 5; i++ {
		go worker(in, out)
	}

	go sendLotsOfWork(in)
	receiveLotsOfwork(out)
}
