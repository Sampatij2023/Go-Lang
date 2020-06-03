package main

import "fmt"

func main() {

	// panic

	defer func() {
		fmt.Println("close resource")
		details := recover()
		fmt.Println("Details : ", details)
	}()
	fmt.Println("open resource")
	panic("panoicking: I don't know what to do")

}
