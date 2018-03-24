package main

import "fmt"

func main() {
	//有缓存跟无缓存的区别
	//c := make(chan int, 1)  //有缓存 异步
	c := make(chan int) //无缓存 同步阻塞

	go func() {
		d := <-c
		fmt.Println(d)
	}()

	c <- 100

	fmt.Println("读取完成")
}
