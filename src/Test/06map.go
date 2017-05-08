package main

import (
	"fmt"
	"sort"
)

func map1() {
	var m map[int]string //map  哈希表、字典
	m = map[int]string{}

	fmt.Println(m)

	m1 := make(map[int]string)
	fmt.Println(m1)

	m1[1] = "OK"
	fmt.Println(m1)

	delete(m1, 1) //删除键值对
	fmt.Println(m1)
}

func map2() {
	var m2 map[int]string
	m3 := make(map[int]string)

	m4 := map[int]string{1: "ok", 2: "test"}

	m3[1] = "test1"
	m3[2] = "test2"
	m3[3] = "test3"
	m3[4] = "test4"
	m3[5] = "test5"

	//delete(m3,1)

	fmt.Println(m2, m3, m4)

	for k, v := range m3 { //for range的使用
		fmt.Println(k, v)
	}

	a := []int{1, 2, 3, 4}
	for i, v := range a { //遍历数组
		fmt.Println(i, v)
	}
}

func map3() {
	var m = make([]map[int]string, 10)
	for _, v := range m {
		v = make(map[int]string)
		v[1] = "OK" //value 是 副本
		fmt.Println(v)
	}
	fmt.Println(m) //[map[] map[] map[] map[] map[] map[] map[] map[] map[] map[]]

}

func main() {
	map1()

	map2()

	map3()

	map4()
}

func map4(){
	var m =[]int{4,5,2,1,8}
	sort.Ints(m)
	fmt.Println(m)   //[1 2 4 5 8]
}