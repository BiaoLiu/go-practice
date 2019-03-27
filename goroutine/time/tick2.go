package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1000)
	sign := make(chan byte, 1)

	for i := 0; i < 1000; i++ {
		ch <- i
	}

	var t *time.Timer
	ok := true

	go func() {
		for {
			select {
			case c := <-ch:
				{
					fmt.Println(c)
				}
			case <-func() <-chan time.Time {
				if t == nil {
					t = time.NewTimer(time.Millisecond)
				} else {
					t.Reset(time.Millisecond)
				}
				return t.C
			}():
				fmt.Println("Timeout.")
				ok = false
				break
			}

			//终止for循环
			if !ok {
				sign <- 0
				break
			}
		}
	}()
	<-sign
}
