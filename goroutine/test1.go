package main

import (
	"fmt"
	"runtime"
)

func main() {
	go println()

	//暂定当前的goroutine
	runtime.Gosched()
}

func println() {
	fmt.Println("test")

}
