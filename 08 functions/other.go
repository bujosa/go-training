package main

import (
	"fmt"
)

func main(){
	foo(1, 2, 3, 4)
}

func foo(x ...int) {
	fmt.Println(x)
}