package main

import (
	"math/rand"
	"fmt"
)

func main() {
	num := rand.Intn(9)
	fmt.Println(num)

	var s []int= make([]int,3,10)
	fmt.Println(len(s),cap(s))

	var c chan int=make(chan int)
	fmt.Println(c)

}
