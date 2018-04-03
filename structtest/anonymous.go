package main

import "fmt"

type Room struct {
	number int
}

type Student struct {
	name string
	age  int
	room []int
	Room
}

func main() {
	//r := Room{number: 10}

	int1 := make([]int, 10, 20)
	int1[1] = 1
	int1[2] = 0

	s := Student{name: "lbi", age: 10, room: int1}

	s.room[2] = 3
	fmt.Println(s)

	s.number=10

	fmt.Println(s.number)


	fmt.Println(int1[2])
}
