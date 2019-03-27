package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	context.Background()

	stop := make(chan int)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("监控退出了...")
				return
			default:
				fmt.Println("监控中...")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	time.Sleep(5 * time.Second)
	stop <- 1

	fmt.Println("监控结束...")
}
