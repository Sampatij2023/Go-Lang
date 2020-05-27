package main
import "fmt"

func main(){
  fmt.Scanln()
  myMap := map[int]string{
    1 : "a",
    2: "c",
    3 : "b",
  }
  newMap := checkMap(myMap)
  fmt.Println("newMap is",newMap)
  fmt.Println("myMap is ",myMap)

}
func checkMap ( myMap map[int]string)map[int]string{
  myMap[1]="changed"
  return myMap
}
