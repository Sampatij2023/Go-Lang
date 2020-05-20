package main

import "fmt"

func main() {
	outer := s()
	fmt.Println("Break the Glass ")
	outer()
}

func s() func() {
	fmt.Println("Oh.., I have been revoked")
	var c = func() {
		fmt.Println("cool, now I have been called")
	}
	return c
}
