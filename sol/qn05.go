// Fill in the blanks so this program compiles and produces
// the output shown.
package main

import "fmt"

func main() {
	// Declare a variable that holds a map with strings for keys
	// and booleans for values.
	var isVovel = make(map[string]bool)
	isVovel["a"] = true
	isVovel["b"] = false
	isVovel["c"] = false
	// a-->20  and b-->10
	fmt.Println(isVovel)

	// Call make to create the map.
	// => map[string]bool{"a":true, "b":false, "c":false}

	// Make a map with ints as keys and strings as values. The 1 key
	// should have the value "H", 2 should have the value "He", and
	// 3 should have the value "Li".
	elements := make(map[int]string)
	elements[1] = "Li"
	elements[2] = "H"
	elements[3] = "He"

	// Print all the keys and corresponding values in the slice.
	// Order doesn't matter.
	for atomicNumber, symbol := range elements {
		fmt.Println(atomicNumber, symbol)
	}
	// => 3 Li
	// => 1 H
	// => 2 He
}
