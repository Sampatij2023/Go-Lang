package main

import "fmt"
//structure defined
type part struct {
	description string
	count       int
}

//function to change the inner properties
func doublePack(p *part) {
	p.count *= 2
}

func main() {
	var fuses part
	fuses.description = "Fuses"
	fuses.count = 5
  //calling by address so it has changes at the address of the varialbe
	doublePack(&fuses)
	fmt.Println(fuses.description)
	fmt.Println(fuses.count)
}
