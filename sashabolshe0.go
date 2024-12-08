package main

import (
	"fmt"
)

// написать функцию которая принимает список целых чисел и возвращает в ответ список только с числами больше нуля
//
// перелаю список [-10 - 20 0 10 20]
//
// в ответ получаю [10 20]

func main() {
	numbers := []int{-20, -10, 0, 10, 20}
	result := OnlyPositive(numbers)
	fmt.Println(result) // [10, 20]
}

func OnlyPositive(numbers []int) []int {
	positiveSlice := []int{}
	for _, number := range numbers {
		if number > 0 {
			positiveSlice = append(positiveSlice, number)
		}

	}
	return positiveSlice

}
