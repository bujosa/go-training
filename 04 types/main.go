package main 

import (
	"fmt"
)

var a int = 42
type bujosa int
var b bujosa

var x bool 
var y float64

// use constants
const ww = 42
const www = 42.28
const w = "Bujosa"

// Other format 

// const (
//  ww = 42
//  www = 42.28
//  w = "Bujosa"
// )


// Use own type
func main() {
	b = 1020
	y = 42.334324
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	a = int(b)

	// boolean
	fmt.Println(x)
	x = true
	fmt.Println(x)

	// float64
	fmt.Println(y)

	// string
	s := "Hello Bujosa"
	fmt.Printf("%T\n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n",bs)

	for i, v := range s {
		fmt.Println(i, v)
	}

	test()
}

func test(){
	fmt.Println(w)
	fmt.Println(ww)
	fmt.Println(www)
	fmt.Printf("%T\n",w)
	fmt.Printf("%T\n",ww)
	fmt.Printf("%T\n",www)
}

// iota increment