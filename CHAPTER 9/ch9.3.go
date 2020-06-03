package main

import "fmt"

type Popltn int

func main() {
	var population Popltn
	population = Popltn(572)
	fmt.Println("Sleepy Creek Country population:", population)
	fmt.Println("Congratulations, Kevin and Anna! It's a girl !")
	population += 1
	fmt.Println("Sleepy Creek Country population:", population)
}
