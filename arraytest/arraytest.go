package arraytest

import (
	"fmt"
)

func m3() {
	//a := [...]int{1, 2, 3}
	//fmt.Println(a)
	//
	//var a [3]int=[...]int{1,2,3}
	//数组值类型 切片引用类型
	a := [...]int{1, 2, 3}
	changeArray(a)

	fmt.Println(a[0]) //值传递
	var a2 [3]int = [3]int{4, 5, 6}
	a2[2] = 77
	fmt.Println(a2[2])

	var s []int = []int{1, 2, 3, 4}
	s = append(s, 5, 6, 7)
	fmt.Println(s)
	for i := 0; i < len(s)-1; i++ {
		fmt.Println(s[i])
	}
	fmt.Println(len(s), cap(s))
}

func M4() {
	a1 := []int{1, 2, 3, 4, 5, 6, 7}
	//a11:=[]int{}
	//fmt.Println(a11)

	a2 := a1[:2]
	a2[0], a2[1] = 10, 11
	fmt.Println(a1, len(a1), cap(a1))

	slice := []int32{1, 2, 3, 4, 5, 6}
	slice2 := slice[:2]
	fmt.Println(slice, slice2)
	slice3 := append(slice2, 50, 60, 70, 80, 90) //全新数组
	fmt.Println(slice, slice2, slice3)

	_ = append(slice2, 50, 60) //替换原数组数据
	fmt.Printf("slice为：%v\n", slice)
	fmt.Printf("操作的切片：%v\n", slice2)

	for index, elem := range slice {
		fmt.Println(index, elem)
	}
	for _, elem := range slice {
		fmt.Println(elem)
	}
	for index := range slice {
		fmt.Println(index)
	}

}

func M5() {
	a := make([]int, 10, 20)
	fmt.Println(a)
}

func changeArray(a [3]int) {
	a[0] = 100
}
func changeSlice(s []int) {
	s[0] = 100
}
