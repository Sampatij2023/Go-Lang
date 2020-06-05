package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

	f, er := os.Open("outfile.txt")
	defer f.Close()
	checkError(er)
	bucket := make([]byte, 5)
	bytesRead, e := f.Read(bucket)
	checkError(e)
	fmt.Println("Content of the file (limited) is :", string(bucket[:bytesRead]))
	fmt.Println("bytes Read :", bytesRead)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error Occurred:", err)
	}
}
