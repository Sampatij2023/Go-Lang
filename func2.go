package main

import "fmt"

func main() {
	var number int
	fmt.Println("Enter a number of which you want fibonacci number")
	fmt.Scanln(&number)
	series := fibSeries(number)
	fmt.Println(series)
}

func fibSeries(number int) int {
	a := 1
	b := 1
	var c int

	if number == 1 {
		return 1
	}
	if number == 2 {
		return 1
	}

	for i := 3; i <= number; i++ {
		c = a + b
		a = b
		b = c
	}
	return c

}
