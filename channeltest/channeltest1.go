package main

import (
	"fmt"
	//"math/rand"
)

func main() {
	c := make(chan int)


	go func() {
		fmt.Println("go go go...")

		//i := 0
		//for {
		//	if i < 10 {
		//		i++
		//		r := rand.Intn(100)
		//		fmt.Println("add chan:", r)
		//		c <- r
		//	}
		//}

		c <- 10
		//close(c)  //关闭channel
	}()

	//i2 := <-c
	for item := range c {
		fmt.Println(item)
	}

	//fmt.Println(i2)
}
