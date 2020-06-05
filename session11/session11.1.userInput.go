package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Type something: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Input is : ", input)

	fmt.Println("Type something")
	var number int
	_, err := fmt.Scan(&number)
	if err != nil {
		log.Fatal("Error Occured ", err)
	}
	fmt.Println("Input is :", number, "Error is :", err)

}
