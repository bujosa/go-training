package main 

import (
	"fmt"
)

// uses iota
const (
	a = iota
	b
	c
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}