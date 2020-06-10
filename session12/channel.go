package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go writeToChannel(ch)
	value := <-ch
	fmt.Println("Data in our channel is:", value)
}

func writeToChannel(ch chan int) {
	fmt.Println("Inside Go routine")
	ch <- 10
	time.Sleep(time.Second)
	fmt.Println("data has been written to channel")
}
