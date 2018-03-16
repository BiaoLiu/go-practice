package main

import (
	"reflect"
	"fmt"
)

type Element interface{}



func main(){
	var a=5
	a1:=reflect.TypeOf(a)
	v:=reflect.ValueOf(a)


	fmt.Println(a1,v)
}