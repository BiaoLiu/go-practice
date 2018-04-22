package main

import (
	"fmt"
)

func doWork(id int, w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n", id, n)
		w.done <- true
	}
}

type worker struct {
	in   chan int
	done chan bool
}

func createWorker(id int) worker {
	w := worker{in: make(chan int), done: make(chan bool)}
	go doWork(id, w)
	return w
}

func chanDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
		<-worker.done
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
		<-worker.done
	}
}

func main() {
	chanDemo()
}
