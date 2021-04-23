package main 

import "fmt"

func main() {
	fmt.Println("Am testing a foo function")
	foo()

	for i := 0; i < 100; i++ {
		if i %2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo(){
	fmt.Println("Running foo function")
}