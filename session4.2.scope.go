package main

import "fmt"

var num int

func main() {

	display()
	num = 30
	fmt.Println("in main ", num)
}
func display() {
	num1 := 40
	fmt.Println(num1)
	fmt.Println("in display ", num)

}
