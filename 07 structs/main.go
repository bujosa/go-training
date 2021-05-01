package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

func main() {
	p1 := person{
		first: "David",
		last:  "Bujosa",
	}

	p2 := person{
		first: "Sebastiano",
		last:  "Faiella",
	}

	fmt.Println(p1)
	fmt.Println(p2)
}