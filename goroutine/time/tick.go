package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTicker(10 * time.Second)

	fmt.Println("now: ", time.Now())

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	<-t.C

	fmt.Println("now: ", time.Now())
}
