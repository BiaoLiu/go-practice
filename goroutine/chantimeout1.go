package main

import (
	"fmt"
	"time"
)

func main()  {
	intChan:=make(chan int)

	t:=time.NewTimer(time.Second)

	for{
		select {
		case elem,ok:=<-intChan:
			fmt.Println(elem,ok)
		case <-t.C:
			fmt.Println("time out")
		}
	}
}
