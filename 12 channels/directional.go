package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 2)
	cr := make(<-chan int) // receive
	cs := make(chan<- int) // send
	fmt.Printf("%T\n",c)
	fmt.Printf("%T\n",cr)
	fmt.Printf("%T\n",cs)
}