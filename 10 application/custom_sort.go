package main

import (
	"fmt"
	"sort"
)


type person struct {
	first string
	age   int
}

type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }

func main() {
	p1 := person{"Bujosa", 21}
	p2 := person{"Sebastiano", 24}
	p3 := person{"Raichu", 4}
	p4 := person{"Pichu", 18}

	people := []person{p3, p2, p1, p4}
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
}