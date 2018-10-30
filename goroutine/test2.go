package main

import (
	"fmt"
	"time"
)

func main() {
	//name := "Eric"
	//
	//go func() {
	//	fmt.Println("name is ", name)
	//}()
	//
	//time.Sleep(time.Microsecond)
	//name = "Harry"
	////time.Sleep(time.Microsecond)

	names := []string{"test1", "test2", "test3", "test4", "test5"}
	for _, name := range names {
		go func() {
			fmt.Println(name)
		}()
	}

	time.Sleep(time.Microsecond)
}
