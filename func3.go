package main

import "fmt"

func main() {
	var number1,number2 int
	fmt.Println("ENter first number")
	fmt.Scanln(&number1)
	fmt.Println("ENter second number")
	fmt.Scanln(&number2)
	
	sum := sumTwoNumbers(number1,number2)
	fmt.Println("sum of the numbers is",sum)
}
func sumTwoNumbers(number1 int, number2 int )int{
	return number1+number2
}