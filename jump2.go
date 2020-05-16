package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 8 {
			break
		}

		if i == 5 {
			continue
		}

		fmt.Println("iterating number", i)

	}
}
