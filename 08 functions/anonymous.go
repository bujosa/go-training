package main

import (
	"fmt"
)

func main() {

	// anonymous func
	func() {
		fmt.Println("Anonymous func ran")
	}()

	// func expresion
	 f := func(){
		 fmt.Println("My first func expression")
	 }

	 f()

	 // return a func
	 s1 := foo()
	 fmt.Println(s1)
}

func foo() string {
	s := "Hello world"
	return s
}