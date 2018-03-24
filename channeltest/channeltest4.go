package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpuNums := runtime.NumCPU()
	//runtime.GOMAXPROCS(cpuNums)
	fmt.Println(cpuNums)

	c := make(chan bool)
	for i := 0; i < 10; i++ {
		go Go(c, i)
	}
	<-c
	fmt.Println("执行结束")
}

func Go(c chan bool, index int) {
	a := 0
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println(index, a)

	if index == 9 {
		c <- true
	}
}
