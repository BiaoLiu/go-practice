package go_pratice

import "fmt"

const (
	a1 = 1
	b2 = 2
	c3 = iota
)

func main() {
	fmt.Println(a1, b2, c3)

}


func test1() {
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	i := 0
	for i < 10 {
		fmt.Print(i)
	}
}

func test2() bool {
	a := 10
	switch a {
	case a == 10:
		return true
	default:
		return false

	}
}

func test3() {
	var a []int    //切片
	var b [10]int  //数组
	var c [...]int //数组

	var p *[10]int = &b           //指向数组的指针
	var t *[10]int = new([10]int) //指向数组的指针

	a[0] = 1
	b[0] = 1
	p[0] = 1
	c[0] = 1
	t[0] = 1
}

func test4() {
	var a []int //切片
	var b []int = make([]int, 2, 10)

	fmt.Print(a, b)

	a = append(a, 1, 2, 3)
}
