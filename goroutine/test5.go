package main

import "fmt"

func main() {

	ch1 := make(chan int32)
	ch2 := make(chan int32)

	go func() {
		for c := 'a'; c < 'a'+26; c++ {
			ch1 <- c
		}
		close(ch1)
	}()

	go func() {
		for r := range ch1 {
			//fmt.Println("%c",r)
			ch2 <- r
			ch2 <- (r - 32)
		}
		close(ch2)
	}()

	for r := range ch2 {
		fmt.Printf("%c", r)
	}
}
