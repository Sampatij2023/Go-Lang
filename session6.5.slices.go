package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter the size of Slice")
	fmt.Scanln(&n)
	mySlice := make([]int, n)
	fmt.Println(" Enter the Slice elements ")
	for i := 0; i < len(mySlice); i++ {

		fmt.Scanln(&mySlice[i])
	}
	var del int
	fmt.Println("Enter the index which is to be deleted")
	fmt.Scanln(&indexOfElement)
	if(indexOfElement<=n && indexOfElement>=0){
		newArray :=deleteElementInSlice(mySlice, indexOfElement)
		fmt.Println("My array after deleting the index element")
		fmt.Println(newArray)
	}else{
		fmt.Println("index is out of range of the array")}

}

func deleteElementInSlice(A []int, k int)[]int {
	A = append(A[:k], A[(k+1):]...)
	return A;
}
