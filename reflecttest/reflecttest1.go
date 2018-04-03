package main

import (
	"reflect"
	"fmt"
)

type Student struct {
	name string
	age  int
}

type Any interface{}

func main() {
	s := Student{name: "lbi", age: 10}
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	fmt.Println(t, v)
	fmt.Println("kind: ", t.Kind(), "value: ", v.Kind())

	fmt.Println(reflect.TypeOf(v.Interface()))
	fmt.Println(reflect.TypeOf(v.Interface().(float32)))

	s1:=new(Student)
	s1.name="lbi2"

	var val Any
	val=s1

	switch val.(type) {
	case *Student:
		fmt.Println("")
	}

}
