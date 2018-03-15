package functest

import "fmt"

func Mtest2(args ...int) {
	for index, elem := range args {
		fmt.Println(index, elem)
	}
}

func Mtest3(a int, b int) (int, int) {
	return b, a
}
