package main

import "fmt"

func main() {
	var number1,number2 int32
	var sum float64
	fmt.Println("ENter first number")
	fmt.Scanln(&number1)
	fmt.Println("ENter second number")
	fmt.Scanln(&number2)
	var s,d float64
	sum = sumTwoNumbers(number1,number2)
	fmt.Println("sum of the numbers is",sum)
	s,d=operateTwoNumbers(number1,number2)
	// fmt.Println("sum and difference together is",s,",",d)
	fmt.Println("sum and difference together is",s,",",d)
}
func sumTwoNumbers(number1, number2 int32 )float64{
	return float64(number1+number2)
}

func operateTwoNumbers(number1, number2 int32 )( float64, float64){

	return float64(number1+number2),float64(number1-number2)
}
