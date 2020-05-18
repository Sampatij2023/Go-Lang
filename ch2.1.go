package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	var now time.Time = time.Now()

	var day int = now.Day()

	var year int = now.Year()

	var hour int = now.Hour()

	fmt.Println(day, year)

	fmt.Println(hour)

	songErrStr := "lq lqlq lq lq lqlqlq"
	replacer := strings.NewReplacer("q", "a")
	songStr := replacer.Replace(songErrStr)
	fmt.Println(songStr)

	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(input)

}
