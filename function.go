package main

import "fmt"

func main() {
	displayName()
	increment := numberOne(6)
	fmt.Println(increment)
	fmt.Println("end of main.")
}
func displayName() {
	fmt.Println("Deepak")
}
func numberOne(number int) int {
	number += 1
	return number
}
