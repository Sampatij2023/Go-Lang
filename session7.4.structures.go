package main
import "fmt"

// person struct is created general from can be used any where
type person struct {
  firstName string
  lastName  string
  age       int

}
func main(){
  fmt.Scanln()
  person1:= person{
    // intialising with the struct sub items if the we do not intialise something
    firstName : "Deepak",

    age       : 19,
  }
  // . (dot) helps to print the inner layes items
  fmt.Println(person1.firstName)
  //printing address
  fmt.Println("address is", &person1.firstName)

  //creating new person
  var person2 = newPerson("Deepak", "Suther", 50)
  fmt.Println("new person is", person2  )
}
func newPerson(fName,lastName string, age int)*person{
  var personTest = &person{
    firstName: fName,
    lastName: lastName,
    age:age,
  }
  return personTest
}
