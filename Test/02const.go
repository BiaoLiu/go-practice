package Test

import "fmt"

const ( //常量组
	a = 5
	b = 6
	c = 7
)

const (
	d = 8
	e  //8
	f  //8
)

const (
	g, h = 1, '2'
	l, m  //1,'2'
)

const (
	o = 'A'
	p        //A
	q = iota //0
	r        //1
)

const (
	s = iota
	t = iota
	u = iota
)

const (
	B float64 = 1 << (iota * 10)
	KB
)

func main() {
	var i, j int = 1, 2
	k := 3

	fmt.Println(i, j, k)
	fmt.Println("test")
	fmt.Println(a, b, c) //5 6 7
	fmt.Println(d, e, f) //8 8 8

	fmt.Println("s t u：", s, t, u)

	m1()
}

func m1() {
	var a int = 97
	b := string(a) //并不是将a转换成字符串

	fmt.Println(a, b) //97 a
}
