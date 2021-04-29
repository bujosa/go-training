package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 100; i++{
		fmt.Println(i)
	}

	// one condintion
	x := 100
	for x > 0 {
		fmt.Println(x)
		x -= 2
	}

	for {
		if x > 100 {
			break
		}
		fmt.Println(x)
		x++
	}
}