package main

import (
	"fmt"
	"time"
)

func chanDemo() {
	c := make(chan int)
	go func() {
		for {
			n := <-c
			fmt.Println(n)
		}
	}()
	c<-5
	c<-6

	time.Sleep(time.Second)
}

func main() {
	var c chan int
	chanDemo()

	fmt.Println(c)
}
