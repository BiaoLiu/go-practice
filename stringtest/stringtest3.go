package main

import "fmt"

type student struct {
	Name string
	Age int
}

func main(){
	s:=student{"lbi",18}

	o:=fmt.Sprintf("student lbi is %+v",s)
	fmt.Println(o)
}
