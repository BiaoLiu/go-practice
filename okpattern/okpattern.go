package main

import "fmt"

type People interface {
	getName() string
}

type People2 interface {
	getName2() string
}

type Student struct {
	name string
	age  int
}

type Student2 struct {
	address string
}

func (s Student) getName() string {
	return s.name
}



func main() {
	var s People = Student{"lbi", 20}

	s1, ok := s.(Student)
	fmt.Println(s1, ok)

	var s2 People =new(Student2)
	if v, ok := s2.(People2); ok {
		fmt.Println(v, ok)
	} else {
		fmt.Println("没有实现 People2接口")
	}
}
