package main

import (
	"fmt"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				//a[i]++   //协程无法交出控制权 死循环
				//协程阻塞交出控制权
				fmt.Println("hello goroutine", i)
			}
		}(i)
	}
	time.Sleep(2 * 1e9)
}
