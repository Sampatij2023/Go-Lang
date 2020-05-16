package main

import "fmt"

func main() {
	cgpa := 8
	switch cgpa {
	case 10:
		fmt.Println("Excellent")
	case 9, 8:
		fmt.Println("very good")
	case 1000:
		fmt.Println("test")
	default:
		fmt.Println("inside Default")
	}
}
