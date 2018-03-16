package structtest

import "fmt"

type struct1 struct {
	a int
	b int
}

func Sstructtest() {
	/*****
	实例化结构体的三种方式
	1.使用new
	2.字面量
	3.字面量值
	*****/

	//s为指针
	s := new(struct1)
	s.a = 5
	s.b = 10

	s2 := &struct1{10, 16}

	var s1 struct1 = struct1{10, 15}

	fmt.Println(s, s1, s2)
}
