package main

import "fmt"

func main(){
	fmt.Println(sum(""))
	fmt.Println(sum("",10))
	fmt.Println(sum("",10, 20, 30, 40))
	var mySlice = []int{10,20,30,40}
	fmt.Println(sum("",mySlice...))
}
func sum(str string,num ...int) int {
	total := 0
	for _,val:= range num {
		total += val
	}
	return total
}