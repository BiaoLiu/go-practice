package main

import (
	"fmt"
)

func main() {
	ch1 := make([]chan int, 10)

	for i := 0; i < 10; i++ {
		ch1[i] = make(chan int)
		go Add(ch1[i])
	}

	for _, c := range ch1 {
		 <-c
	}

}

func Add(c chan int) {
	c <- 1
	fmt.Println("channel加入数据")
}
