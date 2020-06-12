package main

import (
	"fmt"
)

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
	Address
}

type Employee struct {
	Name   string
	Salary float64
	Address
}
type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}

// YOUR CODE HERE:

// Define a struct type named Address that has Street, City, State,
// and PostalCode fields, each with a type of "string".
// Then embed the Address type within the Subscriber and Employee
// types using anonymous fields, so that the code in "main" will
// compile, run, and produce the output shown.

func main() {
	var subscriber Subscriber
	subscriber.Name = "Aman Singh"
	subscriber.AddressDetails.Street = "123 Oak St"
	subscriber.AddressDetails.City = "Omaha"
	subscriber.AddressDetails.State = "NE"
	subscriber.AddressDetails.PostalCode = "68111"
	fmt.Println("Name:", subscriber.Name)                             // => Name: Aman Singh
	fmt.Println("Street:", subscriber.AddressDetails.Street)          // => Street: 123 Oak St
	fmt.Println("City:", subscriber.AddressDetails.City)              // => City: Omaha
	fmt.Println("State:", subscriber.AddressDetails.State)            // => State: NE
	fmt.Println("Postal Code:", subscriber.AddressDetails.PostalCode) // => Postal Code: 68111

	var employee Employee
	employee.Name = "Joy Carr"
	employee.AddressDetails.Street = "456 Elm St"
	employee.AddressDetails.City = "Portland"
	employee.AddressDetails.State = "OR"
	employee.AddressDetails.PostalCode = "97222"
	fmt.Println("Name:", employee.Name)                             // => Name: Joy Carr
	fmt.Println("Street:", employee.AddressDetails.Street)          // => Street: 456 Elm St
	fmt.Println("City:", employee.AddressDetails.City)              // => City: Portland
	fmt.Println("State:", employee.AddressDetails.State)            // => State: OR
	fmt.Println("Postal Code:", employee.AddressDetails.PostalCode) // => Postal Code: 97222
}
