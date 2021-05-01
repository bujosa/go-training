package main

import "fmt"

func main() {
	x := sum(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(x)
}

func sum(xi ...int) int {
	total := 0

	for _, v := range xi {
		total += v
	}

	return total
}