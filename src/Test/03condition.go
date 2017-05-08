package main

import "fmt"

func main() {
	fmt.Println(1 << 2)

	m2()
}

//指针
//
// Go虽然保留了指针，但与其它编程语言不同的是，在Go当中不
//支持指针运算以及”->”运算符，而直接采用”.”选择符来操作指针
//目标对象的成员
//
//操作符”&”取变量地址，使用”*”通过指针间接访问目标对象
//默认值为 nil 而非 NULL

func m2() {
	a := 1
	var p *int = &a //指向a的 指针  & 取变量地址

	fmt.Println(p)  //0xc04200e240
	fmt.Println(*p) // * 通过指针取得对象

	var result = m3()
	fmt.Println(result)

	m4()

	m5()
}

func m3() bool {
	if a := 1; a < 2 {
		return true
	}
	return false
}
func m4() {
	a := 0
	for {
		a++
		if a > 10 {
			break
		}
		fmt.Println(a)
	}

	for a < 10 {
		a++
		fmt.Println(a)
	}

Label1:
	for a := 0; a < 10; a++ {
		fmt.Println(a)
		break Label1  //默认情况下 break 只会跳出最近的一层循环，使用label后可以跳出到指定的循环
	}

}

//不需要写break，一旦条件符合自动终止
//如希望继续执行下一个case，需使用fallthrough语句

func m5() {
	a := 1
	switch a {
	case 0:
		fmt.Println(0)
		fallthrough //继续检查后面的表达式
	case 1:
		fmt.Println(1)
	default:
		fmt.Println("default")
	}

	switch {
	case a > 1:  //表达式
		fmt.Println("小于a")
	case a == 1:
		fmt.Println(1)
	}
}
