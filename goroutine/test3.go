package main

import (
	"fmt"
	"time"
)

func main() {
	names := []string{"test1", "test2", "test3", "test4", "test5"}
	for _, name := range names {
		go func(who string) {
			fmt.Println(who)
		}(name)
	}

	time.Sleep(time.Microsecond)
}
