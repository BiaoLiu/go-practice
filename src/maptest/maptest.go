package maptest

import "fmt"

func Mtest() {
	var m1 map[string]string = map[string]string{"a": "b"}

	value,ok:=m1["c"]
	fmt.Println(value,ok)

	fmt.Println(m1)
}
