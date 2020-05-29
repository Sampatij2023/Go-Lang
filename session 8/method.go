package main

import "fmt"

func main() {
	fmt.Scanln()
	fmt.Scanln()
	fmt.Scanln()
	fmt.Scanln()
	var person1 = person{
		firstName: "Vidya",
		lastName:  "sutaria",
		age:       20,
	}
	var person2 = person{
		firstName: "Vidyam",
		lastName:  "sutariam",
		age:       20,
	}
	fmt.Println("person 1 is :", person1)
	fmt.Println("person 1 full name is :", person1.getFullName())
	fmt.Println("person 1 full name is :", person2.getFullName())
}

func (p *person) getFullName() string {
	return p.firstName + " " + p.lastName
}

type person struct {
	firstName string
	lastName  string
	age       int
}
