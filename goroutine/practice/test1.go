package main

import "time"

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(1 * time.Second)
			ch1 <- "test ch1"
		}()
	}

	for i := 0; i < 5; i++ {
		go func() {
			time.Sleep(2 * time.Second)
			ch2 <- "test ch2"
		}()
	}

	select {
	case <-ch1:

	}

}
