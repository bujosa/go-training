package main

import (
	"fmt"
)

func main(){
	foo()
	bar("Bujosa")
	x, y := compare("Bujosa", "Faiella")
	fmt.Println(x, y)
}

func foo() {
	fmt.Println("Hello I AM B U J O S A")
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