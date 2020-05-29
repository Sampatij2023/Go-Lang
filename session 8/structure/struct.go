package structure

// person struct is created general from can be used any where

// func main() {
// 	fmt.Scanln()
// 	person1 := person{
// 		// intialising with the struct sub items if the we do not intialise something
// 		firstName: "Deepak",

// 		age: 19,
// 	}
// 	fmt.Scanln()
// 	// . (dot) helps to print the inner layes items
// 	fmt.Println(person1.firstName)
// 	//printing address
// 	fmt.Println("address is", &person1.firstName)

// 	//creating new person
// 	var person2 = newPerson("Deepak", "Suther", 50)
// 	fmt.Println("new person is", person2)

// 	var employee1 = employee{
// 		employeeID: 101,
// 		person: person{
// 			firstName: "Deepak",
// 			lastName:  "Suthar",
// 			age:       30,
// 		},
// 	}
// 	fmt.Println("Employee 1 is :", employee1.employeeID)
// }

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func NewPerson(fName, lastName string, age int) *Person {
	var personTest = &Person{
		FirstName: fName,
		LastName:  lastName,
		Age:       age,
	}
	return personTest
}

type employee struct {
	employeeID int
	Person
}
