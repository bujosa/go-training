package main

import (
	"fmt"
)

type person struct {
	first string
}

type secretAgent struct {
	person
	ltk bool
}

func (s person) speak(){
	fmt.Println("I am", s.first)
}

type human interface {
	speak()
}

func bar(h human) {
	switch  h.(type) {
	case person:
      fmt.Println("I was passed into bar", h.(person).first)
	case secretAgent:
      fmt.Println("I was passed into bar", h.(secretAgent).first)
	default:
	  fmt.Println("Unknown")		
	}	
}

func main() {
	p := person{
		first: "David",
	}

	s := secretAgent{
		person: p,
		ltk: true,
	}

	// Polymorphims
	s.speak()
	p.speak()

	bar(s)
	bar(p)
}