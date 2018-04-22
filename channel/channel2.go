package main

import (
	"fmt"
	"time"
)

func worker(i int, c chan int) {
	for {
		fmt.Printf("this is channel %d ,the result is %d \n", i, <-c)
	}
}

func main() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go worker(i, channels[i])
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	time.Sleep(20 * time.Second)
}
