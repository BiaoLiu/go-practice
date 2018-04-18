package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 3)
	exit_ch := make(chan bool)

	ch1 <- 1
	ch1 <- 2
	ch1 <- 3

	go func() {
		for c := range ch1 {
			fmt.Println(c)
		}

		fmt.Println("读取完毕")
		exit_ch <- true
	}()

	fmt.Println("主协程...")

	ch1 <- 4
	ch1 <- 5

	fmt.Println("阻塞完成...")
	close(ch1)

	<-exit_ch
}
