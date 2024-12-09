package main

import (
	"fmt"
)

type Liters float64
type Gallons float64

func main() {
	carFuel := Gallons(10.0)
	busFuel := Liters(240.0)

	fmt.Println(carFuel, busFuel)
}
