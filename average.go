package main

import (
	"fmt"
	"log"

	"awesomeProject/datafile"
)

func main() {
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64
	for _, number := range numbers {
		sum += number
	}

	sampleCount := float64(len(numbers))
	avg := sum / sampleCount
	fmt.Printf("Average: %.2f\n", avg)
}
