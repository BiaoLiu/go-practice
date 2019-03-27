package main

import "fmt"

var syncChan = make(chan int, 3)

func g1() {
	fmt.Println("执行g1")
	syncChan <- 1
}

func g2() {
	fmt.Println("执行g2")
	syncChan <- 2
}

func g3() {
	fmt.Println("执行g3")
	syncChan <- 3
}

	func main() {
		go g1()
		go g2()
		go g3()

		for i := 0; i < 3; i++ {
			<-syncChan
		}

}
