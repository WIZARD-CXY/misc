package main

import (
	"container/heap"
	"fmt"
	"math/rand"
)

type Request struct {
	fn func() int
	c  chan int
}

func workFn() {
	return rand.Intn(100)
}

func requester(work chan<- Request) {
	c := make(chan int)

	for {
		//Kill some time(fake load).
		Sleep(rand.Int63n(nWorker * 2 * Second))
		work <- Request{workFn, c}
		result := <-c

		//furtherProcess(Result)
	}
}

type Worker struct {
	requests chan Request
	pending  int
	index    int
}

func (w *Worker) Work(done chan *Worker) {
	for {
		req := <-w.requests
		req.c <- req.fn()
		done <- w
	}
}

type Pool []*Worker

type Balancer struct {
	pool Pool
	done chan *Worker
}

func (b *Balancer) balance(work chan Request) {
	for {
		select {
		case req := <-work:
			b.dispatch(req)
		case w := <-b.done:
			b.completed(w)
		}

	}
}

func (pool Pool) Less(i, j int) bool {
	return p[i].pending < p[j].pending
}

func (b *Balancer) dispatch(req Request) {
	//grab the least loaded worker
	w := heap.Pop(&b.pool).(*Worker)
	// send task to it
	w.requests <- req

	//one more in its work queue
	w.pending++
	//Put it back to the heap
	heap.Push(&b.pool, w)

}

//job is completed ; update heap
func (b *Balancer) completed(w *Worker) {
	//one fewer in the queue
	w.pending--
	//remove it from heap
	heap.Remove(&b.pool, w.index)
	heap.Push(&b.pool, w)
}
