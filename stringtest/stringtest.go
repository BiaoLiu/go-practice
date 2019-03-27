package main

import (
	"strings"
	"fmt"
	"strconv"
)

func main() {
	index := strings.Index("abc", "d")
	fmt.Println(index)

	if strings.Index("abc", "a") > -1 {
		fmt.Println("存在")
	}

	for _, elem := range strings.Fields("a b c d") {
		fmt.Println(elem)
	}
	fmt.Println("split test")
	for _, elem := range strings.Split("a,b,c,d", ",") {
		fmt.Println(elem)
	}

	fmt.Println("合并的slice")
	str3:= strings.Join([]string{"1","2"},",")
	fmt.Println(str3)

	fmt.Println(strconv.Itoa(100))

	t3:=fmt.Sprintf("test1:%s is %s","t1","t2")
	fmt.Println(t3)
}
