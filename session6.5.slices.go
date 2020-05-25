package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter the size of array")
	fmt.Scanln(&n)
	mySlice := make([]int, n)
	for i := 0; i < len(mySlice); i++ {
		fmt.Println(" Enter the array elements ")
		fmt.Scanln(&mySlice[i])
	}
	var del int
	fmt.Println("Enter the index which is to be deleted")
	fmt.Scanln(&del)

	mySlice = append(mySlice[:del], mySlice[(del+1):]...)
	fmt.Println(" my slice after removing the element is :", mySlice)

}
