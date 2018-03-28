package structtest

import (
	"reflect"
	"fmt"
)

type Tag struct {
	name string "名称"
	age  int    "年龄"
}

type Tag2 struct {
	address string
	Tag
}

func Tagtest() {
	t := Tag{"lbi", 18}
	ttype := reflect.TypeOf(t)
	for i := 0; i < 2; i++ {
		tfield := ttype.Field(i)
		fmt.Println(tfield.Type, tfield.Name, tfield.Tag)
	}

	tag2 := Tag2{}
	tag2.name = "lbi"
	tag2.address = "test"
}
