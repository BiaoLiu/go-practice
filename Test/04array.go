package Test

import "fmt"

func main() {
	t1()
	t3()
}

func t1() {
	//var test []int
	//test := []int{}
	test := [3]int{1, 2, 3}             //[1 2 3]
	test2 := [...]int{19: 1}            //[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1]
	test3 := [...]int{1: 1, 2: 2, 3: 3} //[0 1 2 3]

	fmt.Println(test)
	fmt.Println(test2)
	fmt.Println(test3)

	var test4 *[3]int = &test  //指向数组的指针
	test5 := &test

	fmt.Println(*test4)
	fmt.Println(test5)

	x, y := 1, 2

	fmt.Println(x, y)
	fmt.Println(&x, &y)
}

func t2() {
	x, y := 1, 2
	a := [...]*int{&x, &y} //指针数组

	fmt.Println(x, y)
	fmt.Println(a)
}

func t3() {
	a := [2]int{1, 2}
	b := [2]int{1, 2}
	a[1] = 1
	fmt.Println(a == b)

	p := new([10]int) //&[0 0 0 0 0 0 0 0 0 0]  指针  使用new是生成的指针
	p[2] = 2          //&[0 0 2 0 0 0 0 0 0 0]

	fmt.Println(p)
}
