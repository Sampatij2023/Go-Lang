package main

import "fmt"

func main() {
	passInterface(&circle{
		radius: 10,
	})
}

type shape interface {
	area()
	perimeter()
}

type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return (22 / 7) * (c.radius * c.radius)
}
func (c *circle) perimeter() float64 {
	return (2 * 22 / 7) * (c.radius)
}
func passInterface(sh shape) {
	fmt.Println("Area is :", sh.area())
	fmt.Println("Perimeter is :", sh.perimeter())
}
