package main

import (
	"fmt"
	"go-pratice/links"
)

func main() {
	urls := []string{"https://www.qq.com", "https://2.robo2025.com"}

	worklist := make(chan []string)
	unseenLinks := make(chan string)

	go func() {
		worklist <- urls
	}()

	for i := 0; i < 5; i++ {
		go func() {
			for link := range unseenLinks {
				foundlinks, _ := links.Extract(link)
				fmt.Println(foundlinks)
				go func() {
					worklist <- foundlinks
				}()
			}
		}()
	}

	for list := range worklist {
		fmt.Println("work list:", list)
		for _, link := range list {
			unseenLinks <- link
		}
	}
}
