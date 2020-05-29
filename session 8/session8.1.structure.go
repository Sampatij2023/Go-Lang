package main

import (
	"fmt"
	"./structure"
)

func  main(){

  var person1 = structure.NewPerson("Ravi", "Prakash", 40)  
  fmt.Println("Person 1 is :", person1.FirstName)

}