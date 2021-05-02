package main

import (
	"fmt"
)

func main() {
	a := 42
	fmt.Println(&a)

	b := &a
	// & gives you the address

	fmt.Println(*b)
	// * gives you the value stored at an address when you have the address

    *b = 43
	fmt.Println(a)

	x := 0
	foo(&x)
	fmt.Println(x)

}

func foo(y *int){
  fmt.Println(*y)
  *y = 43
  fmt.Println(*y)
}