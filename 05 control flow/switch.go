package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("this should not print")
	
	case (3==3):
		fmt.Println("prints")
		fallthrough
	case (4!=3):
		fmt.Println("prints")
	default:
		fmt.Println("this is default")
	}
}