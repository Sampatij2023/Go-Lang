package main
import "fmt"

func main(){
	fmt.Scanln()
	fmt.Scanln()
	fmt.Scanln()
	fmt.Scanln()
	fmt.Scanln()

	person1:=person{
		firstName: "Pawan",
		lastName: "chaman",
		age: 34,
		address: []string{"Mumbai","pune"},
	}
	person2:=person{
		firstName: "Pawn",
		lastName: "chamn",
		age: 34,
		address: []string{"Mumbai","pune"},
	}
	person3:=person{
		firstName: "Pan",
		lastName: "chan",
		age: 34,
		address: []string{"Mumbai","pune"},
	}
	fmt.Println("Before calling:" , person1)
	modifyPerson(person1)
	fmt.Println("after calling:" , person1)
	fmt.Println("Before calling:" , person2)
	modifyPerson(person1)
	fmt.Println("after calling:" , person2)
	fmt.Println("Before calling:" , person3)
	modifyPerson(person1)
	fmt.Println("after calling:" , person3)
}
var allPerson[]string
allPerson=[]person{person1,person2,person3}
// fmt.Println(allPerson)
for _,singlePerson:=range allPerson{
	fmt.Println("First name is :",singlePerson.firstName)
	fmt.Println("last name is :",singlePerson.lastName)
	fmt.Println("age is :",singlePerson.age)
	fmt.Println("address is :",singlePerson.address)
}

func modifyPerson(p person){
	p.firstName="NewName"
	p.age = 101
	fmt.Println("Inside modify: ",p)

}

type person struct{
	firstName  string
	lastName   string
	age 	   int
	address    []string
}
