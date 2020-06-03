package main

import "fmt"

func main() {
	func() {
		fmt.Println("Inside anonymous function")
	}() //Immediately invoked function expression. IIFE
	anonyFnc()
}

func anonyFnc() {
	var number = func(s string) {
		fmt.Println(s, "Inside non-anonymous function")
		return 100
	}("Heelo Rajesh,")
	// return number
	fmt.Println("Return value is: ", number)

}
