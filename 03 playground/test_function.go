package main 

import "fmt"

// global variable
var z = 100

// you can identifier w is of TYPE int
var w int

func main() {
	fmt.Println("Am testing a foo function")

	// call a foo function
	foo()

	// simple loop
	for i := 0; i < 100; i++ {
		if i % 50 == 0 {
			fmt.Println(i)
		}
	}

	// Short declaration operator
	// Declare a variable and assing a value
	x := 42
	fmt.Println(x)
	x = 99
	fmt.Println(x)
	y := 100 + 24 
	fmt.Println(y)
}

// this is foo function
func foo(){
	fmt.Println("Running foo function: ", z)
	fmt.Printf("%T\n", z)
	fmt.Printf("%b\n", z) // binario
	fmt.Printf("%x\n", z) // base 16 lowe-case 
	// %X is upper case
	fmt.Printf("%#x\n", z)  // for hex

	s := fmt.Sprintf("%b", z)
	fmt.Println(s)
}