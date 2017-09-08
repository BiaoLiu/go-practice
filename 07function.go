package go_pratice

import "fmt"

func f1() (int, int, int) {
	a, b, c := 1, 2, 3
	return a, b, c
}

func f2(a, b, c int) (d, e, f int) { //命名返回参数
	d, e, f = 1, 2, 3
	//return
	return d, e, f
}

func f3(a ...int) { //不定长参数 必须作为参数列表中的最后一个参数 slice  与python的可变长参数一致
	fmt.Println(a) //[1 2 3 4 5]
}

func main() {
	f3(1, 2, 3, 4, 5)

	f5 := func() {
		fmt.Println("匿名函数")
	}
	f5()

	f6 := closure(5)
	f6(3)
}


func closure(a int) func(int) int { //函数闭包
	fmt.Printf("%p\n", &a) //0xc042036230

	return func(y int) int {
		fmt.Printf("%p\n", &a) //0xc042036230
		return a + y
	}
}


//string,int,float,array,slice,map