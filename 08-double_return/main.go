package main

import (
	"fmt"
)

//Write a function which takes an integer.
//The function will have two returns.
//The first return should be the argument divided by 2.
//The second return should be a bool that let’s us know whether or not the argument was even.
//For example:
//half(1) returns (0, false)
//half(2) returns (1, true)

func half(x int) (int, bool) {
	y := x / 2
	var z bool
	if x%2 == 0 {
		z = true
	} else {
		z = false
	}
	return y, z
}

func main() {
	fmt.Println(half(2))
}
