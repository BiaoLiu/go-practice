package main

import (
	"fmt"
	"encoding/json"
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
}
