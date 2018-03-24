package maptest

import (
	"fmt"
	"math/rand"
)

func Mtest() {
	var m1 map[string]string = map[string]string{"a": "b"}

	var m2 map[string]string = map[string]string{"tet": "wetwe"}
	fmt.Println(m2)

	value, ok := m1["c"]
	fmt.Println(value, ok)

	fmt.Println(m1)
}
