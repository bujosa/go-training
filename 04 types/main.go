package main 

import (
	"fmt"
)

var a int = 42
type bujosa int
var b bujosa

var x bool 

// Use own type
func main() {
	b = 1020
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	a = int(b)
	fmt.Println(x)
	x = true
	fmt.Println(x)
}