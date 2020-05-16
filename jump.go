package main

import "fmt"

func main() {
	var userNumber int
test:
	fmt.Println("Enter a number")
	fmt.Scanln(&userNumber)
	fmt.Println("Name of the user is", userNumber)
	if userNumber > 200 {
		fmt.Println("Exiting")
		return
	}
	if userNumber > 100 {
		fmt.Println("Number is very high")
		goto test
	}
	fmt.Println("Good Number")
}
