package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myArray = [5]string{"Hello", "from", "Go", "Student"}
	fmt.Println(reflect.TypeOf(myArray))
	case1(myArray)
	for index, val := range myArray {
		fmt.Println("index is", index, "value is", val)
	}
	case2(&myArray)
	for index, val := range myArray {
		fmt.Println("index is", index, "value is", val)
	}
}

func case1(arr [5]string) {
	arr[0] = "hey"
	fmt.Println("in case1", arr)
}
func case2(arr *[5]string) {
	arr[0] = "hey"
	fmt.Println("in case2 (deref)", *arr)
	fmt.Println("in case2", arr)
}
