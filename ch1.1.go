package main

import "fmt"

func main() {
	var price int = 100
	fmt.Println("Price is", price, " rupees.")

	var taxRate float64 = 0.08
	var tax float64 = float64(price) * taxRate
	fmt.Println("Tax is ", tax, "rupees.")

	var total float64 = tax + float64(price)
	fmt.Println("Total bill is ", total, "rupees")

	var availableFunds int = 120
	fmt.Println("Available Funds :", availableFunds, "rupees")
	fmt.Println("Within Budget ?", total <= float64(availableFunds))

}
