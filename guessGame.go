package main

import (
	"fmt"
	"math/rand"
)

func main() {

	randomNumber := rand.Intn(50)
	for i := 0; i < 5; i++ {
		fmt.Println("guess the number")
		inputNumber, _ := fmt.Scanln(int)
		if inputNumber > randomNumber {
			fmt.Println("Your number is greater than the random number")
		} else if inputNumber < randomNumber {
			fmt.Println("Your number is smaller than the random number")
		} else {
			fmt.Println("great you have found the random number")
			break
		}
	}
}
