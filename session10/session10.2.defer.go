package main

import "fmt"

func main() {
	//defer close
	// panic
	// open resource

	var number = 10
	// these defer functions are stored in stack
	F1()
	defer F2(number)
	number = 100
	// stack follows last in - first o !! LIFO
	fmt.Println("end of main. number is: ", number)
}

func F1() {
	defer fmt.Println("hello from F1")
	fmt.Println("end of F1")
}
func F2(number int) {
	fmt.Println("end of F2. number is ", number)
}
