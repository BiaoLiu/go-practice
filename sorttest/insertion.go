package main

import "fmt"

func insertion(a []int) {
	n := len(a)

	for i := 1; i < n; i++ {
		j := i - 1
		value := a[i]

		for ; j >= 0; j-- {
			if a[j] > value {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j+1] = value
	}
}

func main() {
	a := []int{2, 5, 1, 7, 3}

	insertion(a)

	fmt.Println(a)

}
