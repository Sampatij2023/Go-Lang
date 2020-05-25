package main

import "fmt"

func main() {
	data := []string{"d", "e", "e", "p", "a", "k"}
	counts := make(map[string]int)
	for _, item := range data {
		counts[item]++
	}

	letters := []string{"s", "u", "t", "p", "a", "k"}
	for _, letter := range letters {
		count, ok := counts[letter]
		if !ok {
			fmt.Printf("%s :not found\n", letter)

		} else {
			fmt.Printf("%s :%d\n ", letter, count)
		}

	}
}
