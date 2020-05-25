package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter the size of array")
	fmt.Scanln(&n)
	mySlice := make([]int, n)
	fmt.Println(" Enter the array elements ")
	for i := 0; i < len(mySlice); i++ {

		fmt.Scanln(&mySlice[i])
	}
	var del int
	fmt.Println("Enter the index which is to be deleted")
	fmt.Scanln(&del)

	updateArray(mySlice, del)

}

func updateArray(A []int, k int) {
	A = append(A[:k], A[(k+1):]...)
	fmt.Println(" my slice after removing the element is :", A)
}
