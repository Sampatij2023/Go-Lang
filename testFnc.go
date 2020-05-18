package main

import "fmt"

func main() {
	test(sum)
	// mathOper(6, 9, sum)
}

// func test(testFnc func(int, int) int, a, b int) {
// 	fmt.Println(testFnc(a, b))
// }
// func sum(a, b int) int {
// 	return a + b
// }
func test(testFnc func()) {
	testFnc()
}

func sum() {
	fmt.Println("Hi")
}
