package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data := []byte("Hello from go")
	err := ioutil.WriteFile("outfile.txt", data, 0644)
	checkError(err)
	fmt.Println("dONE Writing")

	f, er := os.Create("outFile2.txt")
	checkError(er)
	defer f.Close()

	bytesWritten, e := f.WriteString("Hello")
	checkError(e)
	fmt.Println("Bytes written are", bytesWritten)
	fmt.Println("done Writing")

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error Occured", err)
	}
}
