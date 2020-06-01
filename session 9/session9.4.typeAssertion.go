package main

import "fmt"

func main() {
	var i interface{} = "Helloo"
	storString := i.(string)
	fmt.Println(i)
	var num float32 = 17
	checkType(num)
}

func checkType(anyType interface{}) {
	switch anyType.(type) {
	case string:
		fmt.Println("Argument is a string")
	case float64:
		fmt.Println("Argument is a float64")
	case float32:
		fmt.Println("Argument is a float32")
	case int:
		fmt.Println("Argument is a int")
	default:
		fmt.Println("Argument is default")
	}
}
