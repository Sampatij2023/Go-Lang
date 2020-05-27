package main

import "fmt"

func main(){
  // key ---> value
  // fmt.Scanln()
   var myMap = make(map[string]int)
   myMap["a"]=10
   myMap["b"]=30
   myMap["a"]=20
  // // a-->20  and b-->10
  fmt.Println(myMap)
  
 // delete here do not delete anything because "ac" is not present in the myMap
  delete(myMap, "ac")

  // ok is boolean if it true then num1 value gets updated
  var num1,ok =myMap["a"]
  fmt.Println("Ok value is", ok)
  fmt.Println("num1 is",num1)

 //printing after mapping the corresponding elements
  for key,value := range myMap{
      fmt.Println("Key is :", key, "value is :", value)
  }

// fmt.Println("hello")
}
