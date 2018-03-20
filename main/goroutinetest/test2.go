package main

import "fmt"

func main() {
	var ch1 chan string

	ch1 = make(chan string)

	ch1<- "lbi"

	var  value string= <- ch1
	fmt.Println(value)


	fmt.Println(ch1)
}
