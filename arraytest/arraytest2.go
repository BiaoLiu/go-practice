package arraytest

import "fmt"

func Atest2() {
	arr1 := new([3]int)
	fmt.Println(arr1)

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)
	fmt.Println(&arr2)

	int3 := 5
	fmt.Println(&int3)

	arr3 := [3]int{}
	arr4 := arr3

	arr5 := *arr1
	arr5[0] = 1
	fmt.Println(*arr1, arr5)

	arr4[0] = 5

	fmt.Println(arr3, arr4)

	fmt.Println("arr1 value:", *arr1)
	fp(arr1)
	fmt.Println("arr1 value:", *arr1)

}

func f(a [3]int) {
	fmt.Println(a)
}

func fp(a *[3]int) {
	(*a)[2] = 5
	fmt.Println(a)
}

func fp2() {
	var arr [3]int
	fmt.Println(arr)

	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1: 0, 2: 1}
	fmt.Println(arr1, arr2)
}
