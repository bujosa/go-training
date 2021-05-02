package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello")
	fmt.Fprintln(os.Stdout, "Hello world")
}