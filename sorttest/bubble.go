package main

import "fmt"

func bubble(l []int) {
	for i := 0; i < len(l); i++ {
		for j := 0; j < len(l)-1-i; j++ {
			if l[j] > l[j+1] {
				temp := l[j+1]
				l[j+1] = l[j]
				l[j] = temp
			}
		}
	}
}

func main() {
	v := []int{2, 1, 4, 6, 3}
	bubble(v)

	fmt.Println(v)
}
