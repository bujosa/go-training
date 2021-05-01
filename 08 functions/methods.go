package main

import (
	"fmt"
)

type person struct {
	first string
}

func (s person) speak(){
	fmt.Println("I am", s.first)
}

func main(){
	s := person{
		first: "David",
	}
	s.speak()
}
