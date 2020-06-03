package main

import "fmt"

type Car string

func (c Car) Accelerate() {
	fmt.Println("Speeding up")
}

func (c Car) Brake() {
	fmt.Println("Stopping")
}

func (c Car) Steer(direction string) {
	fmt.Println("Turning", direction)
}

type Trunk string

func (t Trunk) Accelerate() {
	fmt.Println("Speeding up")
}

func (t Trunk) Brake() {
	fmt.Println("Stopping")
}

func (t Trunk) Steer(direction string) {
	fmt.Println("Turning", direction)
}

func (t Trunk) LoadCargo(cargo string) {
	fmt.Println("Loading", cargo)
}

type Vehical interface {
	Accelerate()
	Steer(string)
	Brake()
}

func main() {
	var vehical Vehical = Car("Toyoda Yarvic")
	vehical.Accelerate()
	vehical.Steer("left")

	vehical = Trunk("Fnord F180")
	vehical.Brake()
	vehical.Steer("right")
}
