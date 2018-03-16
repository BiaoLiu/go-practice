package Test

import "fmt"

func main() {
	s1()
	s2()
}

func s1() {
	var slice1 []int //slice  切片
	fmt.Println(slice1)

	slice2 := make([]int, 5, 20) //使用make生成切片
	fmt.Println(slice2)
	fmt.Println(len(slice2), cap(slice2))

	a := [10]int{}
	b := a[0:3]

	fmt.Println(b) //[0 0 0]
	//fmt.Println(b[3])  //index out of range
	fmt.Println(cap(b)) //10  容量是底层数组的容量
}

//append
func s2() {
	s1 := make([]int, 3, 6)
	fmt.Print("%p\n", s1)

	s1 = append(s1, 1, 2, 3)
	fmt.Print("%v %p\n", s1, s1)

	s1 = append(s1, 1, 2, 3, 4)
	fmt.Print(&s1)

	array1 := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s2 := array1[1:3]
	s3 := array1[2:5]

	fmt.Println(s2, s3) //[1 2] [2 3 4]

	s2[1] = 5
	fmt.Println(s2, s3) //[1 5] [5 3 4]
}
