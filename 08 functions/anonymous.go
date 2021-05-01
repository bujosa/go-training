package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("Anonymous func ran")
	}()
}