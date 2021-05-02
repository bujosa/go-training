package main

import (
	"fmt"
	"sort"
)

func main() {
	
	s := []int{5, 100, 3, 0, 1, 390, 2000}
	sort.Ints(s)
	fmt.Println(s)
	xs := []string{"David", "Sebastiano", "Adames", "Bujosa", "Faiella"}
	sort.Strings(xs)
	fmt.Println(xs)
}