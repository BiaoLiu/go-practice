package main

import "fmt"

var intChan1 chan int
var intChan2 chan int
var channels = []chan int{intChan1, intChan2}

var numbers = []int{1, 2, 3, 4, 5}

func main() {
	select {
	case getChan(0) <- getNumber(0):
		fmt.Println("1th case is selectd.")
	case getChan(1) <- getNumber(1):
		fmt.Println("2th case is selectd.")
	default:
		fmt.Println("Default")
	}

}

func getNumber(i int) int {
	fmt.Printf("numbers[%d]\n", i)
	return numbers[i]
}

func getChan(i int) chan int {
	fmt.Printf("channels[%d]\n", i)
	return channels[i]
}

// 首先，case后的表达式会先求值， 执行顺序，从左到右，从上到下
