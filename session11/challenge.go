package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var folderName string
	_, erro := fmt.Scan(&folderName)
	// reader := bufio.NewReader(os.Stdin)
	// input, erro := reader.ReadString('\n')
	// // dir := "" + input
	if erro != nil {
		log.Fatal("Error Occured ", erro)
	}
	// folderName = ("" + input)
	// files, err := ioutil.ReadDir("movie")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, f := range files {
	// 	fmt.Println(f.Name())
	// }
	err := filepath.Walk(folderName,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fmt.Println(path)
			return nil
		})
	if err != nil {
		log.Println(err)
	}

}
