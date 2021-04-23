package main 

import (
	"fmt"
)

var a int = 42
type bujosa int
var b bujosa

func main() {
	b = 1020
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	a = int(b)
}