package main

import (
	"calc"
	"fmt"
	"greeting"
)

func main() {
	greeting.Hlw()
	greeting.Hi()

	fmt.Println(calc.Add(12, 52))
	fmt.Println(calc.Subtract(7, 3))
}
