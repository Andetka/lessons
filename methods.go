package main

import (
	"fmt"
)

type MyType string

func (m MyType) sayHi() {
	fmt.Println("Hi")
}

type Number int
type Number1 int

func (n Number) Add(otherNumber int) {
	fmt.Println(n, "plus", otherNumber, "is", int(n)+otherNumber)
}

func (n Number) Subtract(otherNumber int) {
	fmt.Println(n, "minus", otherNumber, "is", int(n)-otherNumber)
}

func (n *Number1) Double() {
	*n *= 2
}

func main() {
	value := MyType("A myType value")
	value.sayHi()
	anotherValue := MyType("another value")
	anotherValue.sayHi()

	ten := Number(10)
	ten.Add(4)
	ten.Subtract(5)
	four := Number(4)
	four.Add(3)
	four.Subtract(2)

	number := Number1(4)
	fmt.Println("Original value of number", number)
	number.Double()
	fmt.Println("number after calling Double:", number)
}
