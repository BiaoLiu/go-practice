package main

import "fmt"

func main() {

	c := make(chan int)

	var c1 <-chan int = c  //只读
	var c2 chan<- int = c  //只写


	c1<-1


	fmt.Println(c1, c2)
}
