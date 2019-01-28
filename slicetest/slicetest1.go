package main

import "fmt"

func main(){
	var v1 []int

	v1=append(v1, 1)

	m1(v1)

	fmt.Println(v1)

}

func m1(v []int){
	v[0]=5
	v=append(v, 2)
	fmt.Println(v)
}
