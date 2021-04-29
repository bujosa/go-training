package main

import (
	"fmt"
)

func main() {
	for i := 33; i <= 122; i++{
		fmt.Printf("%v\t%#U\n", i,i)
	}

	if x:= 100; x == 100 {
		fmt.Println("Limit scope")
	}
}