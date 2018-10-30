package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func g11() {
	fmt.Println("执行g1")
	wg.Done()
}

func g21() {
	fmt.Println("执行g2")
	wg.Done()
}

func g31() {
	fmt.Println("执行g3")
	wg.Done()
}

func main() {
	wg.Add(3)

	go g11()
	go g21()
	go g31()

	wg.Wait()
}
