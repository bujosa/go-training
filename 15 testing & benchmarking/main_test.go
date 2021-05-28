package main

import "testing"

func TestMySum(t *testing) {
	if mySum(2, 3) != 5{
		t.Error("")
	}
}