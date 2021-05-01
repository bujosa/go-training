package main

import (
	"fmt"
)

func main(){
	// defer is uses to close files, process, etc ...
	defer foo(1, 2, 3, 4)
	bar()
}

func foo(x ...int) {
	fmt.Println(x, "Second")
}

func bar() {
	fmt.Println("First")
}