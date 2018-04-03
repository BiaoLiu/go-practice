package main

import (
	"fmt"
)

/********
	Go 语言规范定义了接口方法集的调用规则：

	类型 *T 的可调用方法集包含接受者为 *T 或 T 的所有方法集
	类型 T 的可调用方法集包含接受者为 T 的所有方法
	类型 T 的可调用方法集不包含接受者为 *T 的方法
*/

type Itest2 interface {
	Area2() float32
}

type Square2 struct {
	side float32
}

func (s *Square2) Area2() float32 {
	return s.side * s.side
}

func (s Square2) Area3() float32 {
	return s.side * s.side
}

func test2() {
	var p Itest2 = &Square2{5}
	fmt.Println(p)

	s1 := &Square2{side: 8}
	s1.Area2()
	s1.Area3()

	s2 := Square2{side: 5}
	s2.Area2()
	s2.Area3()

	r := []Itest2{&Square2{10}}
	fmt.Println(r)

	//接口是否是某个类型的实现
	if pm, ok := p.(*Square2); ok {
		pm.Area2()
		pm.Area3()
	}

	var s Itest2 = &Square2{10}

	switch p.(type) {
	case *Square2:
		fmt.Println("square2")
	}

	//某个类型是否实现了接口
	if sv, ok := s.(Itest2); ok {
		fmt.Println(sv, s)
	}
}
