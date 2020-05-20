package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 20
	fmt.Println("Value of a :", a)
	fmt.Println("Memory address of a :", &a)
	p := &a
	fmt.Println("Value of p :", p)
	fmt.Println("Value of p :", *p)
	fmt.Println("Type of p :", reflect.TypeOf(p))
	fmt.Println("Memory address of p :", p)

}
