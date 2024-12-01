package main

import (
	"fmt"
)

type Cat struct {
	Name string
	Age  int64
}

func main() {

	cat1 := Cat{
		Name: "Busia",
		Age:  5,
	}

	secondCat := Cat{
		Name: "Ficha",
		Age:  2,
	}

	cat1.Meow()
	secondCat.Meow()

}

func (cat Cat) Meow() {
	fmt.Println("meow:", "my name is", cat.Name, "I am", cat.Age, "years old")
}
