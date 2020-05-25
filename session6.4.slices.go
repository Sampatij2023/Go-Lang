package main

import "fmt"

func main() {
	mySlice := make([]int, 5, 5)
	for i := 0; i < len(mySlice); i++ {
		mySlice[i] = (i + 1) * 10
	}
	fmt.Println("My slice is : ", mySlice)

	newSlice := mySlice[2:]
	fmt.Println("New slice is :", newSlice)

	newSlice[0] = 100
	fmt.Println("My slice is : ", mySlice)
	fmt.Println("New slice is :", newSlice)

	new2Slice := mySlice[:3]
	fmt.Println("New slice 2 is :", new2Slice)

	new3Slice := mySlice[1:3]
	fmt.Println("New slice 3 is :", new3Slice)

	mySlice = append(mySlice[:2], mySlice[3:]...)
	fmt.Println("the my slice after removing the element is :", mySlice)

}
