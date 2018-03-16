package structtest

import "fmt"

type foo1 struct {
	name string
	age  int
}

type foo2 struct {
	name string
	f    *foo1
}

func Sstructtest10() {
	f := foo2{"lbi", &foo1{"lbi1", 20}}
	fmt.Println(f)
}

func (f *foo2) test() {
	foo1 := &foo1{"lbi1", 10}
	f.f = foo1
}
