package functest

func Ftest20(a int, b int) (func(c int) int) {
	d := a + b

	return func(c int) int {
		return c + d
	}
}
