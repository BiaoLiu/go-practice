package main

import (
	"fmt"
)

func Mtest() {
	var m1 map[string]string = map[string]string{"a": "b"}

	var m2 map[string]string = map[string]string{"tet": "wetwe"}
	fmt.Println(m2)

	value, ok := m1["c"]
	fmt.Println(value, ok)

	fmt.Println(m1)
}

type test struct {
	Name string
	Age  int
}

func main() {
	t := new(test)

	t2 := new([]int)
	if *t2 == nil {
		fmt.Println("true")
	}
	fmt.Println(t2)

	m1 := make(map[string]int)

	t3 := make([]int, 3)
	fmt.Println(t3)

	fmt.Println(t.Name, t.Age)
	fmt.Println(m1)
}
