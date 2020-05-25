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
	if(del<=n && del>=0){
		newArray :=deleteArray(mySlice, del)
		fmt.Println("My array after deleting the index element")
		fmt.Println(newArray)
	}else{
		fmt.Println("index is out of range of the array")}

}

func deleteArray(A []int, k int)[]int {
	A = append(A[:k], A[(k+1):]...)
	return A;
}
