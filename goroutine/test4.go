package main

import "fmt"

func main() {
	c1 := make(chan string)
	//c2:=make(chan<- string)
	//	c3:=make(<-chan string)

	t1 := <-c1
	fmt.Println(t1)

	fmt.Println(c1)
}
