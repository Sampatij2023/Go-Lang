package main

import (
	"fmt"
	"week2/geo"
)
func main() {
	location := geo.Landmark{}
	location.Name = "The Googleplex"
	location.Latitude = 34.75
	location.Longitude = -104.08
	fmt.Println(location)
}
