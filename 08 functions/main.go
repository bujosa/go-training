package main

import (
	"fmt"
)

func main(){
	foo(1, 2, 3, 4)
	bar("Bujosa")
	x, y := compare("Bujosa", "Faiella")
	fmt.Println(x, y)
}

func foo(x ...int) {
	fmt.Println(x)
}

// Everything in Go is "PASS BY VALUE"
func bar(s string){
	fmt.Println(s)
}

func compare(L string, R string) (string, bool) {
	if L == R {
		return "Is Equal", true
	}

	return "Is not Equal", false
}