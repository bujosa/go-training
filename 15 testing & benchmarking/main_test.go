package main

import "testing"

func TestMySum(t *testing.T) {

	type test struct {
		data []int
		answer int
	}

	tests := []test{
		test{[]int{21, 21}, 42},
		test{[]int{10, 11}, 21},
		test{[]int{2, 3}, 5},
	}

	for _, v:= range tests {
		x := mySum(v.data...)
		if  x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}
}