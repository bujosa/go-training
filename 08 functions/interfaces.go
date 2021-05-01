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

type human interface {
	speak()
}

func main() {
	fmt.Println("Hello world")
}