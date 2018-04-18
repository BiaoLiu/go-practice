package main

import "fmt"

type ICar interface {
	getBrand() string
}

type IPeople interface {
	getUserName() string
}

//实现了ICar 与 IPeople
type man struct {
	name    string
	carName string
}

func (m man) getBrand() string {
	return m.carName
}

func (m *man) getUserName() string {
	return m.name
}

func m1(car ICar) {
	fmt.Println(car.getBrand())
}

func m2(people IPeople) {
	fmt.Println(people.getUserName())
}

func main() {
	m := man{"lbi", "Toyota"}
	m1(m)
	m2(&m)

	var mm *man = new(man)
	mm.name = "lbi2"
	mm.carName = "toyota2"

	m1(mm)
	m2(mm)

	//指针方法可以通过指针调用
	//值方法可以通过值调用
	//接收者是值的方法可以通过指针调用，因为指针会首先被解引用
	//接收者是指针的方法不可以通过值调用，因为存储在接口中的值没有地址
}
