package main

import (
	"fmt"
	"sync"
	"time"
)

func chanDemo() {
	c := make(chan int)
	go func() {
		for {
			n := <-c
			fmt.Println(n)
		}
	}()
	c<-5
	c<-6

	time.Sleep(time.Second)
}

type Test struct {

}

func main() {
	var c chan int
	chanDemo()

	fmt.Println(c)



	var wg sync.WaitGroup
	wg.Add(10+1)

	go func() {
		wg.Done()
	}()

	for i:=0;i<10;i++{
		go func(i int) {
			wg.Done()
		}(i)
	}
	wg.Wait()
}
