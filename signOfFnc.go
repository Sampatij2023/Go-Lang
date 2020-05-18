package main

import "fmt"

func main() {
	fmt.Println(add(3, 2))
	fmt.Println(sub(6, 3))
	fmt.Println(sub(3, 6))
	// mathOper(4, 9, add)
	// mathOper(4, 9, sub)
	fmt.Println(mathOper(add))
	fmt.Println(mathOper(sub))

}

func sub(a, b int32) int {
	if a > b {
		return int(a - b)
	}

	return int(b - a)

}
func add(a, b int32) int {
	return int(a + b)
}

func mathOper(testFnc func(int32, int32) int) int {
	var a, b int32
	a = 10
	b = 5

	return int(testFnc(a, b))

}
