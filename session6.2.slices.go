package main

import "fmt"

func main() {
	mySlice := make([]int, 5, 5)
	for i := 0; i < len(mySlice); i++ {
		mySlice[i] = (i + 1) * 10
	}
	fmt.Println("My slice is :",mySlice)
	fmt.Println("My slice length is :",len(mySlice))
	fmt.Println("My slice capacity is :",cap(mySlice))
	fmt.Println("address of first element is :",&(mySlice[0]))

	mySlice = append(mySlice, 60, 70, 80, 90)
	
	fmt.Println("My slice is :",mySlice)
	fmt.Println("My slice length is :",len(mySlice))
	fmt.Println("My slice capacity is :",cap(mySlice))
	fmt.Println("address of first element is :",&(mySlice[0]))
	// newSlice := mySlice
	// newSlice[0]=100
	// fmt.Println(mySlice)
	// var mySlice []int
	// fmt.Println(mySlice == nil)
}
