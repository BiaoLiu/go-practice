package main

type Itest interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func test() {
	//var p Itest
	//p = &Square{10}
	//
	//p.Area()
}
