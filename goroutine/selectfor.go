package main

import "fmt"

func main() {
	initChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		initChan <- i
	}
	close(initChan)

	syncChan := make(chan struct{}, 1)

	go func() {
	Loop:
		for {
			select {
			case elem, ok := <-initChan:
				if !ok {
					fmt.Println("End.")
					break Loop
				}
				fmt.Println(elem)
			}
		}
		syncChan <- struct{}{}
	}()
	<-syncChan

}

//关闭channel对接收方没有影响，关闭channel后，如果channel中还有元素，接收方仍然可以接收元素
