package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
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

	sa := secretAgent{
		person: p1,
		ltk: true,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(sa)

	// anonymous struct

	p3 := struct {
		first string
		last string 
		age int 
	}{
		first: "David",
		last: "Bujosa",
		age: 21,
	}

	fmt.Println(p3)
}