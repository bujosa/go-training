// Package sbtt asks if you are ready to rock
package sbtt

// Subtraction subtract an unlimited numbers of values of type int
func Subtraction(xi ...int) int {
	s := 0
	for _, v := range xi {
		s -= v
	}
	return s
}