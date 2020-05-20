package main

import "fmt"

func main() {
	array := [6]string{"D", "e", "e", "p", "a", "k"}
	slice := array[1:3]
	slice = append(slice, "x")
	slice = append(slice, "y", "z")
	for _, letter := range slice {
		fmt.Println(letter)
	}
}
