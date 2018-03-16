package method

type Tag struct {
	name string
	age  int
}

func maintest() {
	t := new(Tag)
	t.Metest2()
}

func (t Tag) Metest1() {
	t.name = "1"
	t.age = 2
}

func (t *Tag) Metest2() {
	t.name = "2"
	t.age = 50
}
