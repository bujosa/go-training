package mymathtwo

import (
	"fmt"
	"testing"

	"github.com/GoesToEleven/go-programming/code_samples/007-documentation/01/mymath"
)

func TestSum(t *testing.T) {
	type test struct {
		data []int
		answer int
	}

	tests := []test{
		test{[]int{21, 21}, 42},
		test{[]int{3, 4, 21}, 28},
	}

	for _, v := range tests {
		x :=mymath.Sum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}
}

func ExampleSum() {
	fmt.Println(mymath.Sum(2, 3))
}