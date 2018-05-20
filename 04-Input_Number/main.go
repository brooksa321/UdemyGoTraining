package main

import (
	"fmt"
)

//Create a program that prints to the terminal asking for a user to enter a small number and a larger number.
//Print the remainder of the larger number divided by the smaller number.
func main() {
	var smallNum int
	var largeNum int
	fmt.Print("Please enter a number: ")
	fmt.Scan(&smallNum)
	fmt.Print("Pease enter a number that is larger than the one you just entered: ")
	fmt.Scan(&largeNum)
	fmt.Println(largeNum % smallNum)
}
