package arraytest

import "fmt"

func Atest() {
	var arr1 [3]int = [3]int{}
	arr1[0] = 1
	arr1[1] = 2
	for index, elem := range arr1 {
		fmt.Println(index, elem)
	}

	if arr2 := [3]int{1, 2, 3}; arr2[0] == 1 {
		fmt.Println("arr2[0]为1")
	}

	//arr2:=append(arr1, 3)
}
