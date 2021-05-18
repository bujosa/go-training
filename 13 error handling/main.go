package main

import (
	"fmt"
	"log"
	"os"
)


func main() {
	n, err := fmt.Println("Hello")

	if err != nil {
       fmt.Println(err)
	}

	// file

	_, err = os.Open("no-file.txt")

	if err != nil{
	   log.Println("err happened", err)
	   panic(err)
	}

	fmt.Println(n)
}