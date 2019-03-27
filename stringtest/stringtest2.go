package main

import "fmt"

func Stest2() {
	s1 := "sglwegjwlieg"

	b := []byte(s1)
	b[1] = 'y'

	s2 := string(b)

	fmt.Println(b, s2)
}
