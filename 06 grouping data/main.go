package main

import (
	"fmt"
)

func main() {
	var x [5]int
	fmt.Println(x)
	x[3] = 100
	fmt.Println(x)
	fmt.Println(len(x))

	// slice

	y := []int{4, 5, 7, 8, 42}
	fmt.Println(y)

	for i, v := range y {
		fmt.Println(i, v)
	}

	fmt.Println(y[1:3])

	y = append(y, 77, 88)
	fmt.Println(y)
    z := []int{1, 2, 3}
	y = append(y, z...)
	fmt.Println(y)

	// uses make

	a := make([]int, 10, 100)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	// multi-dimensional slice
	jb:= []string{"Bujosa", "Faiella"}
	mp:= []string {"David", "Sebastiano"}
	xp:= [][]string{jb, mp}
	fmt.Println(xp)

	// uses map

	m := map[string]int{
		"Bujosa":21,
		"Faiella":24,
	}
	fmt.Println(m)
	fmt.Println(m["Faiella"])
	v, ok := m["Ogando"]
	fmt.Println(v)
	fmt.Println(ok)

	m["Ogando"] = 21
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}

	delete(m, "Ogando")
	fmt.Println(m)
}