package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "David",
		Last:  "Bujosa",
		Age:   21,
	}
	fmt.Println(p1)

	// uses marshal from json package
	bs, err := json.Marshal(p1)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))

	// uses unmarshal for eliminated tags
}
