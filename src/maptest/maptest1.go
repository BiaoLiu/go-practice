package maptest

import (
	"fmt"
	"sort"
)

func Maptest2() {
	var m2 map[int]string = map[int]string{1: "1"}
	fmt.Println(m2)
}

func Maptest3() {
	m1 := map[int]string{}
	m1[0] = "str1"
	m1[1] = "str2"
	m1[2] = "str3"

	if m2,ok:=m1[8];ok {
		fmt.Println(m2, ok)
	}else{
		fmt.Println("not in")
	}
	delete(m1,2)

	for key,value:=range m1{
		fmt.Println(key,value)
	}

	s1:= []string{"z","a","f","c","d","b","k"}
	sort.Strings(s1)

	fmt.Println(s1)
}
