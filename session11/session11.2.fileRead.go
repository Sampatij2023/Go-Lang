package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("outfile.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Content is", content)
	fmt.Println("Content is", string(content))
	result := string(content)
	fmt.Println("content is ", result)
}
