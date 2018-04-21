package main

import "fmt"

var ch1 = make(chan string)
var ch2 = make(chan bool)

func sample1() {
	fmt.Println("goroutine1 start")
	ch1 <- "lbi goroutine"
}

func sample2() {
	fmt.Println("goroutine2 start")
	str1 := <-ch1
	fmt.Println(str1)
	ch2 <- true
}

func main() {
	go sample1()
	go sample2()

	isEnd := <-ch2
	fmt.Println(isEnd)

	fmt.Println("the end")
}
