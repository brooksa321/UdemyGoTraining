package main

import (
	"fmt"
)

var highest int

//Write a function with one variadic parameter that finds the greatest number in a list of numbers.
func greatest(x ...int) int {
	for _, num := range x {
		if num > highest {
			highest = num
		}
	}
	return highest
}

func main() {
	n := []int{43, 56, 87, 12, 45, 57, 200}
	answer := greatest(n...)
	fmt.Println(answer)
}
