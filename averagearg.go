package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"awesomeProject/average"
)

func main() {
	arguments := os.Args[1:]
	var numbers []float64
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	fmt.Printf("Average: %0.2f\n", average.Average(numbers...))
}
