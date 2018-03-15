package functest

import "fmt"

func Ftest8(a []int) {
	b := &a
	fmt.Println(b)
}

func Ftest9(a *[]int) {
	var a1 = []int{}
	a1 = append(a1, 1, 2, 3, 4)

	*a = append(*a, 5)

	b := a
	fmt.Println(b)
}

func Ftest10(a, b int, c string) (d int, e string) {
	d = a
	e = c
	return
}

// this function changes reply:
func Multiply(a, b int, reply *int) {
	*reply = a * b
}

func Ftest11() {
	c := func(a int, b int) int { return a + b }
	fmt.Println("匿名函数",c)
}
