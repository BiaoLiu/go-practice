package structtest

//工厂模式
type foo struct {
	name string
	age  int
}

func Foo(name string, age int) *foo {
	f := new(foo)
	f.name = "lbi"
	f.age = 20
	return f
}
