package main

import (
	"fmt"
)

type student struct {
	name string
	age  int
}

type student2 struct{}

type any interface{}

func main() {
	s1 := new(student)
	s1.name = "lbi2"

	var val any
	val = s1

	switch val.(type) {
	case *student:
		fmt.Println("")
	}

	s2,ok:= val.(student2)
	fmt.Println(s2,ok)
}
