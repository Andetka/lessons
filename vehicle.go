package main

import "fmt"

type Car string

func (c Car) Accelerate() {
	fmt.Println("Ускоряемся!")
}

func (c Car) Brake() {
	fmt.Println("Тормозим!")
}

func (c Car) Steer(direction string) {
	fmt.Println("Поворот", direction)
}

type Truck string

func (t Truck) Accelerate() {
	fmt.Println("Ускоряемся!")
}
func (t Truck) Brake() {
	fmt.Println("Тормозим!")
}

func (t Truck) Steer(direction string) {
	fmt.Println("Поворачиваем", direction)
}

func (t Truck) LoadCargo(cargo string) {
	fmt.Println("Грузим", cargo)
}

type Vehicle interface {
	Accelerate()
	Brake()
	Steer(direction string)
}

func TryVehicle(vehicle Vehicle) {
	vehicle.Accelerate()
	vehicle.Steer("направо!")
	vehicle.Steer("налево!")
	vehicle.Brake()
	truck, ok := vehicle.(Truck)
	if ok {
		truck.LoadCargo("контейнеры!")
	}
}

func main() {
	TryVehicle(Truck("Sitrak"))
}
