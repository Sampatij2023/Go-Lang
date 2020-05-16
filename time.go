package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().Hour()
	if now <= 12 {
		fmt.Println("It's morning, have you brushed your teeth?")
	} else if now <= 17 {
		fmt.Println("Have a nice afternoon!")
	} else {
		fmt.Println("yeAH! it's evening")
	}
}
