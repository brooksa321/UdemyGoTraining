package main

import (
	"fmt"
)

//Modify the previous program to use a func expression.
func main() {
	half := func(x int) (int, bool) {
		y := x / 2
		var z bool
		if x%2 == 0 {
			z = true
		} else {
			z = false
		}
		return y, z
	}
	fmt.Println(half(5))
}
