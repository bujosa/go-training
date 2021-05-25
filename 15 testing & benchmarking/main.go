package main

import (
	"fmt"
)

func main() {
	fmt.Println("2 + 3 =",mySum(2, 3))
	fmt.Println("10 + 5 =",mySum(10, 5))
	fmt.Println("7 + 7 =",mySum(7, 7))
}

func mySum(xi ...int) int {
	sum := 0

	for _, v := range xi {
		sum += v
	}

	return sum
}