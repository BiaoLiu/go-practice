package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		a int
		b int
		c int
	}{
		{3, 5, 8},
		{1, 5, 6},
		{5, 10, 8},
	}

	for _, test := range tests {
		if test.c != Add(test.a, test.b) {
			t.Error("result error")
		}
	}
}
