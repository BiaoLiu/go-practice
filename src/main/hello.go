package main

import "fmt"
//import "arraytest"
//import "maptest"
import (
	//"functest"
	//"maptest"
	//"stringtest"
	"timetest"
)

//import "test2"

var (
	a       = 5
	b       = 6
	c       = 7
	h, e, f = 8, 9, 10
)

var d = 10
var m int = 100

const (
	t1 = 1
	t2 = 2
)

type (
	文本 string
)

//func main() {
//	//m3()
//	//arraytest.M4()
//	//maptest.Mtest()
//	//functest.Mtest2(1,2,3,4,5,6)
//	//stringtest.Stest1()
//	timetest.Ttest1()
//}

func m0() {
	//fmt.Println("hello world")
	//
	//a := "string"
	//
	//var f1 float32=1.11
	//i1:=int(f1)
	//fmt.Println(i1)
	//
	//var d string = "test"
	//fmt.Println(d,m)
	//
	//isSuccess := false
	//fmt.Println(isSuccess)
	//
	//fmt.Println("返回的结果:", test2.GetTest2())
	//
	//fmt.Println(b)
	//fmt.Println(c)
	//fmt.Println(a)
	//
	//fmt.Println("import B:",test2.B)
}

func m1() {
	var a int = 1
	var b *int = &a
	var f *int
	fmt.Println(f) //nil

	var c **int = &b
	fmt.Println(&b, c)

	fmt.Println(*c, **c)

	fmt.Println(a, b)

	if a == 1 {
		fmt.Println("a==1")
	} else {
		fmt.Println("a!=1")
	}
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
}

func m2() {
	var test string = "abc";
	out := ""
	switch test {
	case "abc":
		out = "abc"
		fallthrough //继续执行以下的case
	case "a":
		out = "a"
	default:
		out = "default"
	}
	fmt.Println(out)
}

func m3() {
	//a := [...]int{1, 2, 3}
	//fmt.Println(a)
	//
	//var a [3]int=[...]int{1,2,3}
	//数组值类型 切片引用类型
	a := [...]int{1, 2, 3}
	changeArray(a)

	fmt.Println(a[0]) //值传递
	var a2 [3]int = [3]int{4, 5, 6}
	a2[2] = 77
	fmt.Println(a2[2])

	var s []int=[]int{1,2,3,4}
	s=append(s, 5,6,7)
	fmt.Println(s)
	for i:=0;i<len(s)-1;i++{
		fmt.Println(s[i])
	}
	fmt.Println(len(s),cap(s))
}

func changeArray(a [3]int) {
	a[0] = 100
}
func changeSlice(s []int) {
	s[0] = 100
}



