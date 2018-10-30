package main

import (
	"fmt"
	"encoding/json"
	"net/url"
)

type Student5 struct {
	Name string
	Age  int
}


func main() {
	s := Student5{"lbi", 15}
	b, _ := json.Marshal(s)
	//fmt.Println("json format %s", v)
	fmt.Println(string(b))

	url.Parse("https://www.robo2025.com/login?next=http://login.robo2025.com")
	//fmt.Println(u.Query())
	//fmt.Println(u.Path)
	//fmt.Println(u.RawPath)
	//
	//fmt.Println(err)
}
