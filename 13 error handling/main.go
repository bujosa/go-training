package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	defer foo()

	// file
	_, err := os.Open("no-file.txt")

	if err != nil{
	   log.Println("err happened", err)
	   panic(err) // use panic
	}
}

func foo(){
	fmt.Println("When os.exit() is called, deffered functions dont run")
}