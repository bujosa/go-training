package main

import (
	"fmt"
)

func main() {
	for i := 33; i <= 122; i++{
		fmt.Printf("%v\t%#U\n", i,i)
		if i%2 == 0{
			fmt.Println(i)
		}
	}

	if x:= 100; x == 100 {
		fmt.Println("Limit scope")
	} else if x == 43 {
		fmt.Println("our value was 43")
	} else {
		fmt.Println("our value its not 43 and 100")
	}

	for i := 0; i <= 100; i++{
		if i%2 == 0{
			fmt.Println(i)
		}
	}
}